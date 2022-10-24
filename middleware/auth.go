package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/Maulidito/personal_project_go/dataservice"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type JwtCustomClaim struct {
	*jwt.RegisteredClaims
	Name string
}

type JwtAuthMiddleware struct {
	userDataService *dataservice.DataServiceUser
	db              *gorm.DB
}

func NewJwtAuthMiddleware(userDataService *dataservice.DataServiceUser, db *gorm.DB) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{userDataService, db}
}

func (mid *JwtAuthMiddleware) Authentication(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return fiber.ErrUnauthorized
	}
	tokenRaw := strings.ReplaceAll(strings.TrimPrefix(token, "Bearer"), " ", "")
	fmt.Println(tokenRaw)
	claim := JwtCustomClaim{}
	tokenAfterParse, err := jwt.ParseWithClaims(tokenRaw, &claim, func(t *jwt.Token) (interface{}, error) {

		if err := t.Claims.Valid(); err != nil {
			return nil, err
		}

		return []byte(os.Getenv("SECRET_KEY_AUTH")), nil
	})

	if !tokenAfterParse.Valid || err != nil {
		return fiber.ErrBadRequest
	}

	user, err := mid.userDataService.GetOneByName(c, mid.db, claim.Name)

	if err != nil {
		return fiber.ErrNotFound
	}

	c.Locals("user", user)

	return c.Next()
}
