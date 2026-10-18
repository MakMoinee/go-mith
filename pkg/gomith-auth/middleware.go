package gomithauth

import (
	"context"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/response"
	"github.com/MakMoinee/go-mith/pkg/utils"
)

func (svc *mithAuthService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		tokenString, err := utils.GetTokenString(header)
		if err != nil {
			response.Error(w, response.ErrorResponse{
				ErrorMessage: err.Error(),
				ErrorStatus:  http.StatusUnauthorized,
			})
			return
		}

		if svc.ValidateToken(tokenString) {
			response.Error(w, response.ErrorResponse{
				ErrorMessage: "invalid token",
				ErrorStatus:  http.StatusUnauthorized,
			})
			return
		}

		claims := svc.getClaimsWithToken(tokenString)
		if claims == nil {
			response.Error(w, response.ErrorResponse{
				ErrorMessage: "invalid token",
				ErrorStatus:  http.StatusUnauthorized,
			})
			return
		}

		ctx := context.WithValue(
			r.Context(),
			svc.contextKey,
			claims.Subject,
		)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
