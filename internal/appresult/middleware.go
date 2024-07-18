package appresult

import (
	"errors"
	"fmt"
	"github.com/Hajymuhammet03/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(r.Method, " <---> ", r.URL.Path)

		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}
			fmt.Println("any")
			w.WriteHeader(http.StatusTeapot)
			fmt.Println("some")
			w.Write(systemError(err).Marshal())
			return
		}
	}
}

func MiddlewareVideo(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(r.Method, " <---> ", r.URL.Path)

		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}
			fmt.Println("any")
			//w.WriteHeader(http.StatusTeapot)
			fmt.Println("some")
			w.Write(systemError(err).Marshal())
			return
		}
	}
}

func HeaderContentTypeJson() (string, string) {
	return "Content-Type", "application/json"
}

func TokenClaims(token, secretKey string) (jwt.MapClaims, error) {
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("err", err)
		return nil, ErrMissingParam
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !ok {
		// TODO tokenin omrini test etmeli
		return nil, ErrInternalServer
	}

	return claims, nil
}

func MidTokenChkSupAdmin(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("authorization")
		cfg := config.GetConfig()
		claims, err := TokenClaims(token, cfg.JwtKey)

		if err != nil || fmt.Sprint(claims["uuid"]) == "" {
			fmt.Println("err on mid token dvd", err)
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write(ErrNotAcceptable.Marshal())
			return
		}

		r.Header.Add("UUID", fmt.Sprint(claims["uuid"]))
		r.Header.Add("ROLE", fmt.Sprint(claims["role"]))

		h(w, r)
		return

	}

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
