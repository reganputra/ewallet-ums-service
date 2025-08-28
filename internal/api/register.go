package api

import (
	"ewallet-ums/constant"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {

	var log = helpers.Logger

	req := models.User{}
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
	resp, err := api.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to register user", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constant.ErrFailedRegister, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, constant.Success, resp)
}
