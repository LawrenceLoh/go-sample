package service

import (
	"encoding/json"
	"sample/exception"
	"sample/model"
	"sample/repository"
	"time"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetUsers godoc
// @Summary      Get User
// @Description  Get all User
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.User
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Retrieve all users from the database
	var users []model.User
	repository.DbInstance.Find(&users)
	recordNotFound := exception.PanicType1{Message: "Lawrence Trigger error test"}
	panic(recordNotFound)
	// Return the users as JSON response
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Generate a unique ID for the user
	newUser.ID = uuid.New().String()
	// Set the current timestamp for CreatedAt and UpdatedAt fields
	now := time.Now()
	newUser.CreatedAt = now
	newUser.UpdatedAt = now
	// Save the new user to the database
	repository.DbInstance.Create(&newUser)
	// Return the newly created user as JSON response
	json.NewEncoder(w).Encode(newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	var updatedUser model.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Find the user with the given ID
	var user model.User
	result := repository.DbInstance.First(&user, "id = ?", userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	// Update the user's fields
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Photo = updatedUser.Photo
	user.Verified = updatedUser.Verified
	user.Password = updatedUser.Password
	user.Role = updatedUser.Role
	user.UpdatedAt = time.Now()
	// Save the updated user to the database
	repository.DbInstance.Save(&user)
	// Return the updated user as JSON response
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]
	// Find the user with the given ID
	var user model.User
	result := repository.DbInstance.First(&user, "id = ?", userID)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	// Delete the user from the database
	repository.DbInstance.Delete(&user)
	// Return a success message
	response := map[string]string{"message": "User deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
