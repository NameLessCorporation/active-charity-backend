package app

import (
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
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
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

	app.logger.Info("emmet-server successfully started",
		zap.String("addr", app.config.Server.IP+":"+app.config.Server.Port),
		zap.Int64("duration", time.Now().UnixNano()-startTime),
	)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
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
