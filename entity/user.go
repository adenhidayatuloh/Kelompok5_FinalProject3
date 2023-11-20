package entity

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/adenhidayatuloh/glng_ks08_Kelompok5_final_Project_3/pkg/errs"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = "SangatRahasia"

type User struct {
	ID        uint   `gorm:"not null;primary_key;autoIncrement"`
	FullName  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null;CHECK (role IN ('admin', 'member'))"`
	Tasks     []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) errs.MessageErr {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		fmt.Println("Password doesn't match:", err)
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": u.ID,
			"email":  u.Email,
		})

	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticated("Token method is not valid")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		fmt.Print("Hallo ", err)
		return nil, errs.NewUnauthenticated("Token is not valid")
	}

	return token, nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	if isBearer := strings.HasPrefix(bearerToken, "Bearer"); !isBearer {
		return errs.NewUnauthenticated("Token type should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)
	if len(splitToken) != 2 {
		return errs.NewUnauthenticated("Token is not valid")
	}

	tokenString := splitToken[1]
	token, err := u.ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mapClaims jwt.MapClaims

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {

		return errs.NewUnauthenticated("Token is not valid")
	}
	mapClaims = claims

	return u.bindTokenToUserEntity(mapClaims)
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	id, ok := claim["userId"].(float64)
	if !ok {
		return errs.NewUnauthenticated("Token doesn't contains userId")
	}
	u.ID = uint(id)

	email, ok2 := claim["email"].(string)

	if !ok2 {
		return errs.NewUnauthenticated("Token doesn't contains email")
	}
	u.Email = email

	return nil
}
