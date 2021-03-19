package zamio

import (
	modelsIO "auth-merge-api/internal/models/zamio"
	"context"
)

type IUsers interface {
	CreateUser(ctx context.Context, usr modelsIO.User) error
	UpdateUserByPhone(ctx context.Context, phone string, usr modelsIO.User) error
}

type IPersonalData interface {
	CreatePersonalData(ctx context.Context, usr modelsIO.PersonalData) error
	UpdatePersonalDataByPhone(ctx context.Context, phone string, usr modelsIO.User) error
	UpdatePersonalDataByEmail(ctx context.Context, phone string, usr modelsIO.User) error
	UpdatePersonalDataByUserID(ctx context.Context, phone string, usr modelsIO.User) error
}

type ZamioRepo interface {
	IUsers
	IPersonalData
}
