package helpers

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(id uint, email, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return signedToken, err
}

func VerifyToken(tokenStr string) (interface{}, error) {

	errResp := errors.New("token invalid")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}

	return token.Claims.(jwt.MapClaims), nil
}

// func SaveTokenToDB(token, email string) {
// 	tokensByte, err := ioutil.ReadFile("./token_db.json")
// 	errorhandler.Check(err)
// 	tokens := []models.TokenDB{}

// 	err = json.Unmarshal(tokensByte, &tokens)
// 	errorhandler.Check(err)

// 	newToken := models.TokenDB{
// 		Email: email,
// 		Token: token,
// 	}

// 	tokens = append(tokens, newToken)

// 	content, err := json.Marshal(tokens)
// 	errorhandler.Check(err)

// 	err = ioutil.WriteFile("./token_db.json", content, 0644)
// 	errorhandler.Check(err)
// }
