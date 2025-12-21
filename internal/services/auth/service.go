package auth

import "github.com/jackc/pgx/v5"

type Service struct {
	DB *pgx.Conn
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateSuperuser(fullName string, email string, password string) {

}
