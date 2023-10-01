package organization

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/organization"
)

func (o *OrganizationEndpoint) GetTransactionsV1(ctx context.Context, req *organization.GetTransactionsV1Request) (*organization.GetTransactionsV1Response, error) {
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
	transactions, err = o.services.TransactionService.GetTransactionsByWalletID(ctx, org.WalletID, req.GetType(), req.GetStatus())
	if err != nil {
		return nil, err
	}

	var transactionsResponse = make([]*organization.Transaction, 0, len(transactions))
	for _, transactionModel := range transactions {
		var userEmail, userName, fundName string
		if transactionModel.Type == models.TransactionTransferType {
			var fromUser *models.User
			fromUser, err = o.services.UserService.GetUserByWalletID(ctx, transactionModel.FromWalletID)
			if err != nil {
				return nil, err
			}

			var fund *models.Fund
			fund, err = o.services.FundService.GetFundByID(ctx, fromUser.FundID)

			userEmail, userName, fundName = fromUser.Email, fromUser.Name, fund.Name
		}

		transactionsResponse = append(transactionsResponse, &organization.Transaction{
			Id:        transactionModel.ID,
			Type:      transactionModel.Type,
			Coins:     transactionModel.Coins,
			Rubles:    transactionModel.Rubles,
			Status:    transactionModel.Status,
			FromEmail: userEmail,
			FromName:  userName,
			FundName:  fundName,
			CreatedAt: transactionModel.CreatedAt.String(),
		})
	}

	return &organization.GetTransactionsV1Response{Transactions: transactionsResponse}, nil
}
