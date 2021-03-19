package service

import (
	modelsIO "auth-merge-api/internal/models/zamio"
	modelsME "auth-merge-api/internal/models/zamme"
	"context"
)

type ZamIO interface {
	// Уточнит какаие у них витрины, от этого и передавать.
	FromMEToIO(ctx context.Context, usr modelsME.User) error
}

type ZamME interface {
	FromIOToMe(ctx context.Context, usr modelsIO.User, pd modelsIO.PersonalData) error
}

type Interface interface {
	ZamIO
	ZamME
}
