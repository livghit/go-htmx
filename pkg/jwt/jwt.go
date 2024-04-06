package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/livghit/go-htmx/pkg/config"
)

func GenerateJWT(userId int) (string, error) {
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Printf("%v", err)
	}
	log.Println(env.JWT_SECRET)

	var jwtSecret []byte = []byte(env.JWT_SECRET)
	// create a token and endcode it with the jwtSecret

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userId,
			"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
	s, err := t.SignedString(jwtSecret)
	if err != nil {
		log.Fatal(err)
		return "no token", err
	}

	// responsible for Token generation on Successfull login
	return s, nil
}

func SetToken(t string, c *fiber.Ctx) {
	jwtCookie := new(fiber.Cookie)
	jwtCookie.Name = "Authorization"
	jwtCookie.Value = "Bearer" + t
	c.Cookie(jwtCookie)
}

// needs rewrite to handles error on validation
func ValidateToken(t string) {
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Printf("%v", err)
	}

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(env.JWT_SECRET), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}
