package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sadreddinov/distributed_calculator/orchestrator/pkg/models"
	"github.com/sirupsen/logrus"
)

// SignUp godoc
// @Summary      Sign up
// @Description  Create accout
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user   body      models.User  true  "User info"
// @Success      200  {array}  integer	
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /auth/sign-up [post]
func (h *Handler) signUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User

		if err := c.BindJSON(&input); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := h.services.Authorization.CreateUser(input)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignIn godoc
// @Summary      Sign in
// @Description  Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user   body      signInInput  true  "credentials"
// @Success      200  {array}  string "token"	
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /auth/sign-in [post]
func (h *Handler) signIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input signInInput

		if err := c.ShouldBindJSON(&input); err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
		if err != nil {
			logrus.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
	}
}