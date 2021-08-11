package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:255;not null;unique" json:"user_name"`
	AwsAccessKeyId     string    `gorm:"size:100;not null;unique" json:"aws_access_key_id"`
	AwsSecretAccessKey  string    `gorm:"size:100;not null;" json:"aws_secret_access_key"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func Hash(awsSecretAccessKey string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(awsSecretAccessKey), bcrypt.DefaultCost)
}

func VerifySecret(hashedAwsSecretAccessKey, awsSecretAccessKey string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedAwsSecretAccessKey), []byte(awsSecretAccessKey))
}

//func (u *User) BeforeSave() error {
//	hashedAwsSecretAccessKey, err := Hash(u.AwsSecretAccessKey)
//	if err != nil {
//		return err
//	}
//	u.AwsSecretAccessKey = string(hashedAwsSecretAccessKey)
//	return nil
//}

func (u *User) Prepare() {
	u.ID = 0
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.AwsAccessKeyId = html.EscapeString(strings.TrimSpace(u.AwsAccessKeyId))
	u.AwsSecretAccessKey = html.EscapeString(strings.TrimSpace(u.AwsSecretAccessKey))
	u.CreatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "authenticate":
		if u.AwsAccessKeyId == "" {
			return errors.New("required aws access key id")
		}
		if u.AwsSecretAccessKey == "" {
			return errors.New("required aws secret access key")
		}
		return nil
	default:
		if u.UserName == "" {
			return errors.New("required user name")
		}
		if u.AwsAccessKeyId == "" {
			return errors.New("required aws access key id")
		}
		if u.AwsSecretAccessKey == "" {
			return errors.New("required aws secret access key")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found ")
	}
	return u, err
}





