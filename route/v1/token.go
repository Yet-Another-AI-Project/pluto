package v1

import (
	"net/http"

	perror "pluto/datatype/pluto_error"
	"pluto/datatype/request"
	routeUtils "pluto/utils/route"
	"pluto/utils/rsa"
)

// RefreshToken godoc
// @Summary Refresh access token
// @Tags Token
// @Description Refresh access token
// @Accept  json
// @Produce  json
// @Param request body request.RefreshAccessToken true "Refresh access token"
// @Success 200 {object} response.Reponse{body=manage.GrantResult}
//
// @Router /v1/token/refresh [post]
func (router *Router) RefreshToken(w http.ResponseWriter, r *http.Request) *perror.PlutoError {
	rat := request.RefreshAccessToken{}
	if err := routeUtils.GetRequestData(r, &rat); err != nil {
		return err
	}

	res, err := router.manager.RefreshAccessToken(rat)

	if err != nil {
		return err
	}

	if err := routeUtils.ResponseOK(res, w); err != nil {
		return err
	}

	return nil
}

type PublicKeyResponse struct {
	PublicKey string `json:"public_key"`
}

// PublicKey godoc
// @Summary Get public key
// @Tags Token
// @Description Get public key
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Reponse{body=PublicKeyResponse}
//
// @Router /v1/token/publickey [get]
func (router *Router) PublicKey(w http.ResponseWriter, r *http.Request) *perror.PlutoError {
	pbkey, err := rsa.GetPublicKey()

	if err != nil {
		return perror.ServerError.Wrapper(err)
	}

	res := PublicKeyResponse{
		PublicKey: pbkey,
	}

	if err := routeUtils.ResponseOK(res, w); err != nil {
		return err
	}

	return nil
}

// VerifyAccessToken godoc
// @Summary Verify access token
// @Tags Token
// @Description Verify access token
// @Accept  json
// @Produce  json
// @Param request body request.VerifyAccessToken true "Verify access token"
// @Success 200 {object} response.Reponse{body=jwt.AccessPayload}
//
// @Router /v1/token/access/verify [post]
func (router *Router) VerifyAccessToken(w http.ResponseWriter, r *http.Request) *perror.PlutoError {
	accessToken := &request.VerifyAccessToken{}
	if err := routeUtils.GetRequestData(r, accessToken); err != nil {
		return err
	}

	res, err := router.manager.VerifyAccessToken(accessToken.Token)

	if err != nil {
		return err
	}

	if err := routeUtils.ResponseOK(res, w); err != nil {
		return err
	}

	return nil
}
