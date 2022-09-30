package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	jwtToken "technicaltest/pkg/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GetToken(w http.ResponseWriter, r *http.Request) {

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}
	w.WriteHeader(http.StatusOK)
	response := token
	json.NewEncoder(w).Encode(response)
}
