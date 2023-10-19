package authservices

import (
	"fmt"
	"time"

	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/utils"
	"github.com/Dparty/dao/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthService(inject *gorm.DB) AuthService {
	return AuthService{accountRepository: auth.NewAccountRepository(inject)}
}

type AuthService struct {
	accountRepository auth.AccountRepository
}

func (a AuthService) CreateSession(email, password string) (string, error) {
	// fmt.Println("good")
	account := a.accountRepository.GetByEmail(email)
	fmt.Println(account)
	if account == nil {
		return "", fault.ErrUnauthorized
	}
	if !utils.PasswordsMatch(account.Password, password, account.Salt) {
		return "", fault.ErrUnauthorized
	}
	expiredAt := time.Now().AddDate(1, 0, 0).Unix()
	token, err := utils.SignJwt(
		utils.UintToString(account.ID()),
		account.Email,
		string(account.Role),
		expiredAt,
	)
	if err != nil {
		return "", fault.ErrUndefined
	}
	return token, nil
}

func (a AuthService) CreateAccount(email, password string) (Account, error) {
	account, err := a.accountRepository.Create(email, password, "USER")
	if err != nil {
		return Account{}, err
	}
	return Account{entity: *account}, nil
}

func (a AuthService) GetAccount(id uint) *Account {
	account := a.accountRepository.GetById(id)
	if account == nil {
		return nil
	}
	return &Account{*account}
}

func (a AuthService) VerifyToken(token string) (Account, error) {
	auth := AuthorizeByJWT(token)
	if auth.Status != Authorized {
		return Account{}, fault.ErrUndefined
	}
	account := a.GetAccount(auth.AccountId)
	if account == nil {
		return Account{}, fault.ErrUndefined
	}
	return *account, nil
}

func (a AuthService) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := Authorize(c)
		if auth.Status == Authorized {
			account := a.GetAccount(auth.AccountId)
			c.Set("account", account)
		}
		c.Next()
	}
}
