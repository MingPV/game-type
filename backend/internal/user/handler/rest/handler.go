package handler

import (
	"fmt"
	"time"

	"github.com/MingPV/clean-go-template/internal/user/dto"
	"github.com/MingPV/clean-go-template/internal/user/usecase"
	response "github.com/MingPV/clean-go-template/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewHttpUserHandler(useCase usecase.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: useCase}
}

// Register godoc
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body entities.User true "User registration payload"
// @Success 201 {object} entities.User
// @Router /auth/signup [post]
func (h *HttpUserHandler) Register(c *fiber.Ctx) error {
	req := new(dto.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	userEntity := dto.ToUserEntity(req)
	if err := h.userUseCase.Register(userEntity); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToUserResponse(userEntity))
}

// Login godoc
// @Summary Authenticate user and return token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Login credentials (email & password)"
// @Success 200 {object} map[string]interface{} "Authenticated user and JWT token"
// @Router /auth/signin [post]
func (h *HttpUserHandler) Login(c *fiber.Ctx) error {
	loginReq := new(dto.LoginRequest)
	if err := c.BodyParser(loginReq); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	token, userEntity, err := h.userUseCase.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "invalid email or password")
	}

	// Set JWT token in HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false, // must be true on production
		SameSite: "Strict",
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"user": dto.ToUserResponse(userEntity),
	})
}

// LoginWithUsername godoc
// @Summary Authenticate user and return token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Login credentials (username & password)"
// @Success 200 {object} map[string]interface{} "Authenticated user and JWT token"
// @Router /auth/signin [post]
func (h *HttpUserHandler) LoginWithUsername(c *fiber.Ctx) error {
	loginReq := new(dto.LoginWithUsernameRequest)
	if err := c.BodyParser(loginReq); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request")
	}

	token, userEntity, err := h.userUseCase.LoginWithUsername(loginReq.Username, loginReq.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "invalid username or password")
	}

	// Set JWT token in HTTP-only cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false, // must be true on production
		SameSite: "Strict",
		Path:     "/",
		Expires:  time.Now().Add(72 * time.Hour),
	})

	return c.JSON(fiber.Map{
		"user": dto.ToUserResponse(userEntity),
	})
}

// GetUser godoc
// @Summary Get currently authenticated user
// @Tags users
// @Produce json
// @Success 200 {object} entities.User
// @Router /users/me [get]
func (h *HttpUserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return response.Error(c, fiber.StatusUnauthorized, "invalid user id")
	}

	fmt.Println(fmt.Sprint(userID))

	userEntity, err := h.userUseCase.FindUserByID(fmt.Sprint(userID))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user not found")
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// FindUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entities.User
// @Router /users/{id} [get]
func (h *HttpUserHandler) FindUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.Error(c, fiber.StatusBadRequest, "id is required")
	}

	userEntity, err := h.userUseCase.FindUserByID(id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user not found")
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// FindUserByEmail godoc
// @Summary Get user by Email
// @Tags users
// @Produce json
// @Param email path string true "User Email"
// @Success 200 {object} entities.User
// @Router /users/email/{email} [get]
func (h *HttpUserHandler) FindUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return response.Error(c, fiber.StatusBadRequest, "email is required")
	}

	userEntity, err := h.userUseCase.FindUserByEmail(email)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user not found")
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// FindUserByUsername godoc
// @Summary Get user by Username
// @Tags users
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} entities.User
// @Router /users/username/{username} [get]
func (h *HttpUserHandler) FindUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return response.Error(c, fiber.StatusBadRequest, "username is required")
	}

	userEntity, err := h.userUseCase.FindUserByUsername(username)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user not found")
	}

	return c.JSON(dto.ToUserResponse(userEntity))
}

// FindAllUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {array} entities.User
// @Router /users [get]
func (h *HttpUserHandler) FindAllUsers(c *fiber.Ctx) error {
	users, err := h.userUseCase.FindAllUsers()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "internal server error")
	}

	return c.JSON(dto.ToUserResponseList(users))
}
