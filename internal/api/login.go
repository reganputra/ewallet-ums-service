package api

import (
	"ewallet-ums/constant"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log  = helpers.Logger
		req  models.LoginRequest
		resp models.LoginResponse
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to parse request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constant.ErrFailedParseRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Validation error", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constant.ErrFailedValidation, nil)
		return
	}

	resp, err := api.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("Login failed", err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, constant.ErrLoginFailed, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constant.Success, resp)
}
