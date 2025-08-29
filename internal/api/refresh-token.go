package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenSvc interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var log = helpers.Logger

	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")
	if !ok {
		log.Error("Token claim not found in context")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Token claim not found", nil)
		return
	}

	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("Token claim is not of type helpers.ClaimToken")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Invalid token claim", nil)
		return
	}

	resp, err := api.RefreshTokenSvc.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("Failed to refresh token", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, "Failed to refresh token", nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, "Token refreshed successfully", resp)
}
