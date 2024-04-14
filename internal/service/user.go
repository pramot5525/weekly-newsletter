package service

import (
	"fmt"
	"strconv"
	"weekly-newsletter/internal/model"
	"weekly-newsletter/internal/repository"
	"weekly-newsletter/pkg/config"

	"gopkg.in/gomail.v2"
)

type UserService interface {
	CreateUser(model.UserRequest) (model.User, error)
	DeleteUser(model.UserRequest) error
	GetUsers() ([]model.User, error)

	SendEmail(email string) error
}

type Service struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(req model.UserRequest) (model.User, error) {
	user := model.User{
		Email: req.Email,
	}

	return s.repo.Create(user)
}

func (s *Service) DeleteUser(req model.UserRequest) error {
	user := model.User{
		Email: req.Email,
	}

	return s.repo.Delete(user)
}

func (s *Service) GetUsers() ([]model.User, error) {
	return s.repo.GetUsers()
}

func (s *Service) SendEmail(toEmail string) error {
	// https://app.brevo.com/settings/keys/smtp
	from := config.GetString(config.EmailSender)

	host := config.GetString(config.EmailHost)
	port, _ := strconv.Atoi(config.GetString(config.EmailPort))

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", "News letter")

	// msg.SetBody("text/plain", "Welcome to the news letter!")
	msg.SetBody("text/html", fmt.Sprintf(`
		<h1>Welcome to the news letter!</h1>
		<a href="http://localhost:9000/api/v1/user/unsubscribe?email=%s">Unsubscribe</a>
	`, toEmail))

	n := gomail.NewDialer(host, port, from, config.GetString(config.EmailAPIKey))

	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil

}
