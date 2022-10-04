package repository

import (
	"awesomeProject3/internal/app/ds"
	"awesomeProject3/internal/app/dsn"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func New(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetPromoByID(id uint) (*ds.UsersBuy, error) {
	users := &ds.UsersBuy{}

	err := r.db.First(users, id).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) NewRandRecords() error {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	age := rand.Intn(90)

	booksList := []string{"Алиса в стране чудес", "361 по Фаренгейту", "Вий", "Замок"}
	bookRandom := rand.Intn(len(booksList))
	magazinbook := booksList[bookRandom]

	nameList := []string{"Алиса", "Иван", "Даниил", "Константин"}
	nameRandom := rand.Intn(len(nameList))
	magazinName := nameList[nameRandom]

	new := ds.UsersBuy{
		Code: uint(code), // код от 100000 до 999999
		Book: magazinbook,
		Name: magazinName,
		Age:  age,
	}
	err := r.db.Create(&new).Error
	if err != nil {
		return err
	}
	return nil
}
