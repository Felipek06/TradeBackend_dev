package api

import (
	"github.com/Felipek06/TradeBackend_dev.git/services"
	"github.com/gofiber/fiber/v2"
)

type NewHandler struct {
	userService *services.NewUserService
	authService *services.NewAuthService
}

// CreateUser godoc
// @Summary      Cria um novo usuário
// @Description  Endpoint para criar usuário
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        userRequest  body  UserRequest  true  "Dados do usuário"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/users [post]
func (h *NewHandler) CreateUser(ctx *fiber.Ctx) error {
	var userRequest UserRequest
	err := isValid(ctx, &userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.userService.CreateNewUser(userRequest.Email, userRequest.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created succesfully"})
}

// CreateUser godoc
// @Summary      Autentica usuário
// @Description  Endpoint para autenticar usuário
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        loginRequest  body  LoginRequest  true  "Dados de login"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router       /api/login [post]
func (h *NewHandler) Login(ctx *fiber.Ctx) error {
	var loginRequest LoginRequest
	err := isValid(ctx, &loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	tokenJWT, err := h.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Login succesfully", "token": tokenJWT})
}
