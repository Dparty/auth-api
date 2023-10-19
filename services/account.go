package authservices

import (
	"github.com/Dparty/dao/auth"
)

type AccountRole string

const (
	ROOT  AccountRole = "ROOT"
	ADMIN AccountRole = "ADMIN"
	USER  AccountRole = "USER"
)

type Account struct {
	entity auth.Account
}

func (a Account) ID() uint {
	return a.entity.ID()
}

func (a Account) Role() AccountRole {
	return AccountRole(a.entity.Role)
}
