package user

import (
	"context"

	"github.com/NameLessCorporation/active-charity-backend/app/models"
	"github.com/NameLessCorporation/active-charity-backend/extra/user"
)

func (u *UserEndpoint) GetTransactionsV1(ctx context.Context, req *user.GetTransactionsV1Request) (*user.GetTransactionsV1Response, error) {
	userID, err := u.services.TokenService.GetUserIDByAccessToken(ctx, req.GetAccessToken(), u.config.JWT.SecretKey)
	if err != nil {
		return nil, err
	}

	var usr *models.User
	usr, err = u.services.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var wallet *models.Wallet
	wallet, err = u.services.WalletService.GetWalletByID(ctx, usr.WalletID)
	if err != nil {
		return nil, err
	}

	var transactions []*models.Transaction
	transactions, err = u.services.TransactionService.GetTransactionsByWalletID(ctx, wallet.ID, req.GetType(), req.GetStatus())
	if err != nil {
		return nil, err
	}

	var transactionsResponse = make([]*user.Transaction, 0, len(transactions))
	for _, transactionModel := range transactions {
		var toOrganizationName string
		if transactionModel.Type == models.TransactionTransferType {
			var org *models.Organization
			org, err = u.services.OrganizationService.GetOrganizationByWalletID(ctx, transactionModel.ToWalletID)
			if err != nil {
				return nil, err
			}

			toOrganizationName = org.Name
		}

		transactionsResponse = append(transactionsResponse, &user.Transaction{
			Id:                 transactionModel.ID,
			Type:               transactionModel.Type,
			ToOrganizationName: toOrganizationName,
			Coins:              transactionModel.Coins,
			Rubles:             transactionModel.Rubles,
			Status:             transactionModel.Status,
			CreatedAt:          transactionModel.CreatedAt.String(),
			UpdatedAt:          transactionModel.UpdatedAt.String(),
		})
	}

	return &user.GetTransactionsV1Response{Transactions: transactionsResponse}, nil
}
