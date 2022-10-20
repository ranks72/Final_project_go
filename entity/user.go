package entity

import (
	"errors"
	"final_project_go/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar; unique;not null"`
	Email     string    `json:"email" gorm:"type:varchar; unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Age       int       `json:"age" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var secret_key = "RAHASIA"

func (u *User) HashPass() errs.MessageErr {
	salt := 8
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, salt)

	if err != nil {
		return errs.NewInternalServerErrorr("something went wrong")
	}

	u.Password = string(hash)

	return nil
}

func (u *User) ComparePassword(userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userPassword))

	return err == nil
}

func (u *User) GenerateToken() string {
	claims := jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(secret_key))

	return signedToken
}

func (u *User) VerifyToken(tokenStr string) error {

	if bearer := strings.HasPrefix(tokenStr, "Bearer"); !bearer {
		return errors.New("login to proceed")
	}

	stringToken := strings.Split(tokenStr, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("login to proceed")
		}

		return []byte(secret_key), nil
	})

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errors.New("login to proceed")
	} else {
		mapClaims = v
	}

	if exp, ok := mapClaims["exp"].(float64); !ok {
		return errors.New("login to proceed")
	} else {
		if int64(exp)-time.Now().Unix() <= 0 {
			return errors.New("login to proceed")
		}
	}

	if v, ok := mapClaims["id"].(float64); !ok {
		return errors.New("login to proceed")
	} else {
		u.ID = int(v)
	}

	if v, ok := mapClaims["email"].(string); !ok {
		return errors.New("login to proceed")
	} else {
		u.Email = v
	}

	return nil

}
