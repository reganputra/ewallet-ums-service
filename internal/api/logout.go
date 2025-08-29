package api

import (
	"ewallet-ums/constant"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var log = helpers.Logger

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("Logout failed", err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, constant.ErrLogoutFailed, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constant.Success, nil)
}
