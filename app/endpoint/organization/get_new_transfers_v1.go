package organization

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) GetNewTransfersV1(ctx context.Context, req *organization.GetNewTransfersV1Request) (*organization.GetNewTransfersV1Response, error) {
	userID, err := o.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), o.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var usr *models.User
	usr, err = o.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var org *models.Organization
	org, err = o.services.OrganizationService.GetOrganizationByID(ctx, usr.OrganizationID)

	if org.OwnerID != userID {
		return nil, status.Error(codes.PermissionDenied, "Пользователь не владеет организацией")
	}

	var transactions []*models.Transaction
	transactions, err = o.services.TransactionService.GetNewTransferTransactionsByToWalletID(ctx, org.WalletID)
	if err != nil {
		return nil, err
	}

	var transactionsResponse = make([]*organization.NewTransfers, 0, len(transactions))
	for _, transactionMode := range transactions {
		var fromUser *models.User
		fromUser, err = o.services.UserService.GetUserByWalletID(ctx, transactionMode.FromWalletID)
		if err != nil {
			return nil, err
		}

		var fund *models.Fund
		fund, err = o.services.FundService.GetFundByID(ctx, fromUser.FundID)

		transactionsResponse = append(transactionsResponse, &organization.NewTransfers{
			Id:        transactionMode.ID,
			Type:      transactionMode.Type,
			FromEmail: fromUser.Email,
			FromName:  fromUser.Name,
			FundName:  fund.Name,
			Coins:     transactionMode.Coins,
			Rubles:    transactionMode.Rubles,
			Status:    transactionMode.Status,
			CreatedAt: transactionMode.CreatedAt.String(),
		})
	}

	return &organization.GetNewTransfersV1Response{NewTransfers: transactionsResponse}, nil
}
