package userrepo

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/url"
	Config "usermvc/config"
	"usermvc/entity"
)

type UserRepo interface {
	Create(context context.Context, user entity.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() *userRepo {
	newDb, err := newDb()
	if err != nil {
		panic(err)
	}
	newDb.AutoMigrate(&entity.User{})
	return &userRepo{
		db: newDb,
	}
}

func newDb() (*gorm.DB, error) {
	conf := Config.NewDbConfig()
	dsn := url.URL{
		User:     url.UserPassword(conf.User, conf.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (u userRepo) Create(context context.Context, user entity.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
