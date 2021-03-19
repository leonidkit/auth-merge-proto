package zamio

import (
	modelsIO "auth-merge-api/internal/models/zamio"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PersonalDataRepo struct {
	conn *pgxpool.Pool
}

func (pr *PersonalDataRepo) CreatePersonalData(ctx context.Context, usr modelsIO.PersonalData) error {
	return nil
}

func (pr *PersonalDataRepo) UpdatePersonalDataByPhone(ctx context.Context, phone string, usr modelsIO.User) error {
	return nil
}

func (pr *PersonalDataRepo) UpdatePersonalDataByEmail(ctx context.Context, phone string, usr modelsIO.User) error {
	return nil
}

func (pr *PersonalDataRepo) UpdatePersonalDataByUserID(ctx context.Context, phone string, usr modelsIO.User) error {
	return nil
}
