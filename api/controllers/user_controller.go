package controllers

import (
	"fmt"
	"jikkosoft/api/responses"
	"jikkosoft/models"
	"net/http"
	"strings"

	"github.com/jaswdr/faker"
)

type UserResponse struct {
	Email    string      `json:"email"`
	UserName string      `json:"userName"`
	Contact  ContactInfo `json:"contactInfo"`
}

type ContactInfo struct {
	MobileNumber string `json:"mobileNumber"`
	Country      string `json:"country"`
}

type GetUsersResponse struct {
	Success bool         `json:"success"`
	Users   UserResponse `json:"users"`
}

type CreateUserResponse struct {
	Success bool `json:"success"`
}

// CreateUser godoc
// @Summary Seed One user at a time using Faker Library for fake data and Gorm as ORM
// @Description This API Endpoint is used to insert user data and contact Details in the Postgres Database
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} CreateUserResponse
// @Router /user [post]
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	contact := models.Contact{}
	faker := faker.New()
	email := fmt.Sprintf("%v%v@%v",
		faker.Person().FirstName(),
		faker.Company().JobTitle(),
		faker.Internet().FreeEmailDomain())

	user.Email = strings.ReplaceAll(email, " ", "")
	user.Password = faker.Internet().Password()
	user.UserName = faker.Internet().User()
	contact.Country = faker.Address().Country()
	contact.MobileNumber = faker.Person().Contact().Phone

	user.Prepare()
	err := user.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = user.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	contact.UserID = int64(user.ID)

	contact.Prepare()
	_, err = contact.Save(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, map[string]bool{
		"success": true,
	})
}

// GetUsers godoc
// @Summary Get All Users of the Database
// @Description This API will get all users and contacts details and merge the two details together and send the response back
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} GetUsersResponse
// @Router /user [get]
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	contact := models.Contact{}

	var userResponse []UserResponse

	users, err := user.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	contacts, err := contact.FindAll(server.DB)
	contactMap := make(map[int64]ContactInfo)
	for contactIndex := range contacts {
		contactData := contacts[contactIndex]
		contactMap[contactData.UserID] = ContactInfo{
			MobileNumber: contactData.MobileNumber,
			Country:      contactData.Country,
		}
	}

	for userIndex := range users {
		userData := users[userIndex]

		userResponse = append(userResponse, UserResponse{
			Email:    userData.Email,
			UserName: userData.UserName,
			Contact:  contactMap[int64(userData.ID)],
		})

	}
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"users":   userResponse,
	})
}
