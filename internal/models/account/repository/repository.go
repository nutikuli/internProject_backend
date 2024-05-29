package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	//
	"github.com/nutikuli/internProject_backend/internal/models/account"

	"github.com/nutikuli/internProject_backend/internal/models/account/entities"
	"github.com/nutikuli/internProject_backend/pkg/utils"

	//"github.com/nutikuli/internProject_backend/internal/models/account/repository"
	"github.com/nutikuli/internProject_backend/internal/models/account/repository/repository_query"

	"github.com/nutikuli/internProject_backend/internal/models/account/dtos"
	adminstruct "github.com/nutikuli/internProject_backend/internal/models/admin/entities"
	customerstruct "github.com/nutikuli/internProject_backend/internal/models/customer/entities"
	storestruct "github.com/nutikuli/internProject_backend/internal/models/store/entities"
)

type AccountRepo struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) account.AccountRepository {
	return &AccountRepo{
		db: db,
	}
}

func (a *AccountRepo) GetAccountCustomers(ctx context.Context) ([]*customerstruct.Customer, error) {
	var customer []*customerstruct.Customer

	err := a.db.GetContext(ctx, &customer, repository_query.SQL_get_account_customer, "CUSTOMER")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return customer, nil
}

func (a *AccountRepo) GetAccountStores(ctx context.Context) ([]*storestruct.Store, error) {
	var store []*storestruct.Store

	err := a.db.GetContext(ctx, &store, repository_query.SQL_get_account_storeaccount, "STORE")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return store, nil
}

func (a *AccountRepo) GetAccountAdmins(ctx context.Context) ([]*adminstruct.Admin, error) {
	var admin []*adminstruct.Admin

	err := a.db.GetContext(ctx, &admin, repository_query.SQL_get_account_admin, "ADMIN")
	if err != nil {
		log.Info(err)
		return nil, err
	}

	return admin, nil
}

func (a *AccountRepo) FindUserAsPassport(ctx context.Context, email string) (*entities.UsersPassport, error) {
	// checking if user email was founded

	userData := &entities.Account{}

	err := a.db.QueryRowx(repository_query.SQL_find_email, email).StructScan(userData)
	log.Info(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("EmailNotFound")
		} else {
			log.Error(err)
			return nil, err

		}
	}

	userPassport := &entities.UsersPassport{
		Id:       userData.Id,
		Email:    userData.Email,
		Password: userData.Password,
	}

	return userPassport, nil
}

func (a *AccountRepo) SignUsersAccessToken(req *entities.UserSignToken) (*dtos.UserToken, error) {
	claims := entities.UsersClaims{
		Id:    req.Id,
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "access_token",
			Subject:   "users_access_token",
			ID:        uuid.NewString(),
			Audience:  []string{"users"},
		},
	}

	mySigningKey := viper.GetString("JWT_SECRET_TOKEN")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		return nil, err
	}
	return &dtos.UserToken{
		AccessToken: ss,
		TokenType:   "Authorization",
		ExpiresIn:   claims.ExpiresAt.String(),
		IssuedAt:    claims.IssuedAt.String(),
	}, nil
}

func (a *AccountRepo) CreateUser(ctx context.Context, req *entities.UserCreatedReq) (*int64, error) {

	args := utils.Array{
		req.Name,
		req.Password,
		req.Phone,
		req.Location,
		req.Email,
		"CUSTOMER",
		true,
	}

	log.Info(args)

	res, err := a.db.ExecContext(ctx, repository_query.SQL_insert_user, args...)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	userId, _ := res.RowsAffected()

	return &userId, nil
}
