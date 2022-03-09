package user

import (
	"SmartBasket/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
)

var jwtSecretKey = helpers.GetEnvKey("JWT_SECRET_KEY")

type UserHandler interface {
	Login(*fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type userHandler struct {
	UserService IUserService
}

func (u userHandler) Register(ctx *fiber.Ctx) error {
	user := User{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return err
	}

	usernameExist, err := u.UserService.GetByUsername(user.Username)
	if usernameExist != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Error: "Username already exist."})

	}

	emailExist, err := u.UserService.GetByEmail(user.Email)
	if emailExist != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Error: "Email already exist."})
	}

	user.Password = helpers.Hash(user.Password)
	newUser, errSvc := u.UserService.Create(user)
	if errSvc != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(Response{Data: UserDtoFromModel(newUser)})
}

func (u userHandler) Login(ctx *fiber.Ctx) error {
	login := LoginDto{}
	err := ctx.BodyParser(&login)
	if err != nil {
		return err
	}

	user, err := u.UserService.GetByEmail(login.Email)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Error: "Email or Password is incorrect."})
	}

	verifyPass := helpers.VerifyPassword(user.Password, login.Password)
	if verifyPass == false {
		return ctx.Status(fiber.StatusInternalServerError).JSON(Response{Error: "Email or Password is incorrect."})
	}

	Jwt := jwt.New(jwt.SigningMethodHS256)
	claims := Jwt.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	s, err := Jwt.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(Response{Data: s})
}

var _ UserHandler = userHandler{}

func NewUserHandler(service IUserService) UserHandler {
	return userHandler{UserService: service}
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
