package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"navigation/internal/config"
	"navigation/internal/logging"
	"navigation/internal/models"
)

type User struct {
	logger     *logging.Logger
	repository Repository
	cfg config.AppConfig
}

func NewUser(logger *logging.Logger, repository Repository, cfg config.AppConfig) *User {
	return &User{
		logger:     logger,
		repository: repository,
		cfg: cfg,
	}
}
 
func (u *User) Create() error {
	login := u.cfg.User.Login
	password := u.cfg.User.Password
	user := models.User{Login: login}

	userCount, err := u.repository.FindRoot()
	if err != nil {
		log.Fatalln("Error find root user")
		return err
	}

	if userCount.ID == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		if err != nil {
			log.Fatalln("Can't bcrypt password")
		}
		user.Password = string(hashedPassword)

		_, err = u.repository.Create(models.User{
			Login:    user.Login,
			Password: user.Password,
		})

		if err != nil {
			return err
		}
		return nil
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		if err != nil {
			log.Fatalln("Can't bcrypt password")
		}
		user.Password = string(hashedPassword)

		err = u.repository.Update(models.User{
			ID: userCount.ID,
			Login:    login,
			Password: string(hashedPassword),
		})

		if err != nil {
			return err
		}
		return nil
	}
}
