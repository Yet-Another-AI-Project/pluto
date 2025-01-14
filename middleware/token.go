package middleware

import (
	"net/http"
	"time"

	perror "pluto/datatype/pluto_error"
	"pluto/utils/jwt"

	"github.com/gorilla/context"
)

func AccessTokenAuth(w http.ResponseWriter, r *http.Request) *perror.PlutoError {
	accessToken, perr := getAccessToken(r)
	if perr != nil {
		return perr
	}

	jwtToken, perr := jwt.VerifyRS256JWT(accessToken)
	if perr != nil {
		return perr
	}

	accessPayload := &jwt.AccessPayload{}

	if perr := jwtToken.UnmarshalPayload(accessPayload); perr != nil {
		return perr
	}

	if accessPayload.Type != jwt.ACCESS {
		return perror.InvalidJWTToken
	}

	if time.Now().Unix() > accessPayload.Expire {
		return perror.JWTTokenExpired
	}

	context.Set(r, "payload", accessPayload)

	return nil
}
