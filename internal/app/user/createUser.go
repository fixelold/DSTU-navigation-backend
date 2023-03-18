package user

import (
	"log"
	"navigation/internal/logging"
	"navigation/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	logger     *logging.Logger
	repository Repository
}

func NewUser(logger *logging.Logger, repository Repository) *User {
	return &User{
		logger:     logger,
		repository: repository,
	}
}

func (u *User) Create() error {
	user := models.User{Login: "root"}

	userCount, err := u.repository.FindRoot()
	if err != nil {
		log.Fatalln("Error find root user")
		return err
	}

	if userCount.ID == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("root"), bcrypt.MinCost)
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
		return nil
	}
}
