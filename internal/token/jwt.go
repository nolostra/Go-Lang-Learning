package token

import(
	"time"
	"github.com/dgrijalva/jwt-go"
)
// Your JWT secret key, you should keep this secret and not expose it in your code
var jwtSecret = []byte("B42A72139CF6F")

func GenerateJWTToken(username string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	// Set expiration time if needed
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Sign the token with your secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func VerifyJWTToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with your secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}