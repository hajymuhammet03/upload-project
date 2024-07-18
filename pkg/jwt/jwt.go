package jwt

import (
	"github.com/Hajymuhammet03/pkg/config"
	"github.com/golang-jwt/jwt"
	"time"

	"fmt"
)

func GetJWT(login, uuid string) (string, error) {
	cfg := config.GetConfig()
	MySigningKey := []byte(cfg.JwtKey)
	token := jwt.New(jwt.SigningMethodHS256)
	//---refresh
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["jwt"] = login
	claims["uuid"] = uuid

	claims["exp"] = time.Now().Add(time.Minute * 3600).Unix()

	tokenString, err := token.SignedString(MySigningKey)
	fmt.Println("TOKEN", tokenString, "GVMVMNBVNMBV")

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenPair(email, uuid string) (interface{}, error) {
	fmt.Println("id ------>>>0000", email)

	type Token struct {
		T  string `json:"t"`
		Rt string `json:"rt"`
	}
	var list Token
	cfg := config.GetConfig()
	MySigningKey := []byte(cfg.JwtKey)
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get dvd etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["jwt"] = email
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	fmt.Println("id ------>>>", email)

	fmt.Println("uuid  ----->>", uuid)

	var err error
	list.T, err = token.SignedString(MySigningKey)
	if err != nil {
		return "string", err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["authorized"] = true
	rtClaims["jwt"] = email
	rtClaims["uuid"] = uuid
	//rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 240).Unix()

	list.Rt, err = refreshToken.SignedString(MySigningKey)
	if err != nil {
		return "string", err
	}

	return list, nil
}

func RefreshToken(RToken string) (interface{}, error) {
	f := GenerateTokenPair

	type Token struct {
		T  string `json:"t"`
		Rt string `json:"rt"`
	}
	var list Token
	fmt.Println("RTTTTTTT________1   ", list.Rt)
	fmt.Println("RTTTTTTT________ 2  ", RToken)
	if list.Rt == RToken {

		f = GenerateTokenPair

	}

	return f, nil
}
