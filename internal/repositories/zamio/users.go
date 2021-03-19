package zamio

import (
	modelsIO "auth-merge-api/internal/models/zamio"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersRepo struct {
	conn *pgxpool.Pool
}

func (ur *UsersRepo) CreateUser(ctx context.Context, usr modelsIO.User) error {
	return nil
}

func (ur *UsersRepo) UpdateUserByPhone(ctx context.Context, phone string, usr modelsIO.User) error {
	return nil
}
