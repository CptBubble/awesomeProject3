package app

import (
	"awesomeProject3/internal/app/repository"
	"context"
)

type Application struct {
	ctx  context.Context
	repo *repository.Repository
}

func New(ctx context.Context) (*Application, error) {
	app := &Application{
		ctx: ctx,
	}
	repo, err := repository.New(ctx)
	if err != nil {
		return nil, err
	}
	app.repo = repo
	return app, nil
}

func (a *Application) Run(ctx context.Context) error {
	a.StartServer()
	return nil
}
