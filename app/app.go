package app

import (
	"context"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NameLessCorporation/active-charity-backend/app/endpoint"
	"github.com/NameLessCorporation/active-charity-backend/app/endpoint/auth"
	"github.com/NameLessCorporation/active-charity-backend/app/endpoint/role"
	"github.com/NameLessCorporation/active-charity-backend/app/endpoint/user"
	"github.com/NameLessCorporation/active-charity-backend/app/repository"
	"github.com/NameLessCorporation/active-charity-backend/app/service"
	"github.com/NameLessCorporation/active-charity-backend/config"
	"github.com/NameLessCorporation/active-charity-backend/dependers/database"
	app_auth "github.com/NameLessCorporation/active-charity-backend/extra/auth"
	app_role "github.com/NameLessCorporation/active-charity-backend/extra/role"
	app_user "github.com/NameLessCorporation/active-charity-backend/extra/user"
	gateway_tools "github.com/NameLessCorporation/active-charity-backend/tools/gateway"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

type App struct {
	logger *zap.Logger
	config *config.Config
}

func NewApp(logger *zap.Logger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (app *App) StartApp(certPath string) error {
	startTime := time.Now().UnixNano()

	connectionAddr := &database.Connection{
		Host:     app.config.DB.Host,
		Port:     app.config.DB.Port,
		User:     app.config.DB.User,
		Password: app.config.DB.Password,
		DBName:   app.config.DB.DBName,
		SSLMode:  app.config.DB.SSLMode,
		CertPath: certPath,
	}

	connectionAddrStr := database.GenerateAddr(connectionAddr)

	db, err := database.NewDB(connectionAddrStr)
	if err != nil {
		return err
	}

	app.logger.Info(" successfully connected to database",
		zap.String("addr", app.config.DB.Host+":"+app.config.DB.Port),
		zap.String("db_name", app.config.DB.DBName),
		zap.String("user", app.config.DB.User),
		zap.Int64("duration", time.Now().UnixNano()-startTime),
	)

	store := repository.NewRepository(db)

	service := service.NewService(store, app.config)
	service.InitServices()

	endpointContainer := app.InitEndpointContainer(service.Services)

	listener, err := net.Listen("tcp", ":"+app.config.Server.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	app_auth.RegisterAuthServer(grpcServer, endpointContainer.AuthService)
	app_user.RegisterUserServer(grpcServer, endpointContainer.UserService)
	app_role.RegisterRoleServer(grpcServer, endpointContainer.RoleService)

	app.logger.Info("active-charity-backend successfully started",
		zap.String("addr", app.config.Server.IP+":"+app.config.Server.Port),
		zap.Int64("duration", time.Now().UnixNano()-startTime),
	)

	go func() {
		app.logger.Fatal("listen grpc server error", zap.Error(grpcServer.Serve(listener)))
	}()

	gwmux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gateway_tools.CustomMatcher),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   false,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "X-Remote-Address", "X-Access-Token"},
		AllowCredentials: true,
	}).Handler(gwmux)

	gwconn, err := grpc.DialContext(
		context.Background(),
		app.config.Server.IP+app.config.Server.Port,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	if err = app_auth.RegisterAuthHandler(context.Background(), gwmux, gwconn); err != nil {
		return err
	}

	if err = app_role.RegisterRoleHandler(context.Background(), gwmux, gwconn); err != nil {
		return err
	}

	if err = app_user.RegisterUserHandler(context.Background(), gwmux, gwconn); err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    app.config.Gateway.IP + app.config.Gateway.Port,
		Handler: gwmux,
	}

	return http.ListenAndServe(gwServer.Addr, wsproxy.WebsocketProxy(handler))
}

func (app *App) InitEndpointContainer(service *service.Services) *endpoint.EndpointContainer {
	authServices := auth.NewAuthEndpoint(service, app.config)
	userServices := user.NewUserEndpoint(service, app.config)
	roleServices := role.NewRoleEndpoint(service, app.config)

	serviceContainer := endpoint.NewEndpointContainer(
		authServices,
		userServices,
		roleServices,
	)

	return serviceContainer
}
