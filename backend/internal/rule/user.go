package rule

import "backend/internal/entity"

type UserPolicy interface {
	CreateUser() *entity.User
}
