package controllers

import (
	"DemoProject/api/auth"
	"DemoProject/api/models"
	"DemoProject/api/responses"
	"DemoProject/api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (s *Server) Authenticate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("authenticate")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := s.AccessToken(user.AwsAccessKeyId, user.AwsSecretAccessKey)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (s *Server) AccessToken(awsAccessKeyId, awsSecretAccessKey string) (string, error) {

	var err error

	user := models.User{}

	err = s.DB.Debug().Model(models.User{}).Where("aws_access_key_id = ?", awsAccessKeyId).Take(&user).Error
	if err != nil {
		return "", err
	}
	//err = models.VerifySecret(user.AwsSecretAccessKey, awsSecretAccessKey)
	//if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
	//	return "", err
	//}
	return auth.CreateToken(user.ID)
}
