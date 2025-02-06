package routes

import (
	"context"
	"net/http"
	"social_web_server/utils"
)

type MiddlewareBody struct{
	Protected bool
	Handlerfunc http.Handler
}

func Protected(next http.HandlerFunc)http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token:= r.Header.Get("Authorization")

		claims,err:=utils.DecodeToken(token)

		if err!=nil{
			utils.ErrorResponse(w,r,http.StatusUnauthorized,"Not Authorizeed")
			return 
		}

		ctx:=context.WithValue(r.Context(),"userid",claims.UserID)
		next(w,r.WithContext(ctx))
	})
}