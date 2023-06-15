package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/token"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users      []User
	sessions   map[string]*User // authToken user
	sessionsMu sync.RWMutex
	db         *Database
}

func NewAuthService(db *Database) (*AuthService, error) {
	users, err := loadUsersFromDB(db)
	if err != nil {
		return nil, err
	}

	return &AuthService{
		users:      users,
		sessions:   make(map[string]*User),
		sessionsMu: sync.RWMutex{},
		db:         db,
	}, nil
}

func (a *AuthService) Authenticate(username, password string) bool {
	for _, user := range a.users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err == nil {
				return true
			}
		}
	}
	return false
}

func (a *AuthService) StartSession(username string, password string) (authToken string, err error) {
	a.sessionsMu.Lock()
	defer a.sessionsMu.Unlock()

	authToken, err = token.Generate()
	if err != nil {
		return authToken, err
	}

	// Associate the session token with the user
	a.sessions[authToken] = &User{Username: username, Password: password, AuthToken: authToken}

	return authToken, nil
}

func (a *AuthService) GetSessionUser(token string) *User {
	a.sessionsMu.RLock()
	defer a.sessionsMu.RUnlock()
	return a.sessions[token]
}

// GetUsernameFromSessions retrieves the username associated with the session token
func (a *AuthService) GetUsernameFromSessions(sessionToken string) string {
	a.sessionsMu.RLock()
	defer a.sessionsMu.RUnlock()
	for _, user := range a.sessions {
		if user.AuthToken == sessionToken {
			return user.Username
		}
	}
	return ""
}

func HandleLogin(authService *AuthService) http.HandlerFunc {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}

			var loginReq LoginRequest
			err = json.Unmarshal(body, &loginReq)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Authenticate the user
			if authService.Authenticate(loginReq.Username, loginReq.Password) {
				token, err := authService.StartSession(loginReq.Username, loginReq.Password)
				if err != nil {
					// Internal server error
					http.Error(w, "Failed to start session", http.StatusInternalServerError)
					return
				}

				// Set the token in the response header
				w.Header().Set("Authorization", token)
				w.WriteHeader(http.StatusOK)
				return
			}

			// Authentication failed
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Return an empty response for GET requests
		w.WriteHeader(http.StatusOK)
	}
}

func HandleRegistration(authService *AuthService) http.HandlerFunc {
	type RegistrationRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Read the request body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}

			// Parse the request body into a RegistrationRequest struct
			var registrationReq RegistrationRequest
			err = json.Unmarshal(body, &registrationReq)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Register the user
			err = authService.Register(registrationReq.Username, registrationReq.Password)
			if err != nil {
				// Registration failed
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			// Registration successful
			w.WriteHeader(http.StatusOK)
			return
		}

		// Return an empty response for GET requests
		w.WriteHeader(http.StatusOK)
	}
}

// Register creates a new user with the provided username and password
func (a *AuthService) Register(username, password string) error {
	// Check if the username already exists in the database
	for _, user := range a.users {
		if user.Username == username {
			return errors.New("Username already exists")
		}
	}

	// Check for strong password
	ok, errMsg := utils.IsStrongPassword(password)
	if !ok {
		return errors.New(errMsg)
	}

	hashedPswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert the new user into the database
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = a.db.db.Exec(query, username, hashedPswd)
	if err != nil {
		return err
	}

	// Create a new user
	newUser := User{
		Username: username,
		Password: string(hashedPswd),
	}

	// Add the user to the list of users in memory
	a.users = append(a.users, newUser)

	return nil
}

func HandleChangePassword(authService *AuthService) http.HandlerFunc {
	type ChangePasswordRequest struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Check if the user is authenticated
			sessionToken := r.Header.Get("Authorization")
			if sessionToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Get the session user from the token
			sessionUser := authService.GetSessionUser(sessionToken)
			if sessionUser == nil {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			// Read the request body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}

			// Parse the request body into a ChangePasswordRequest struct
			var changePasswordReq ChangePasswordRequest
			err = json.Unmarshal(body, &changePasswordReq)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Check if the username matches the session user's username
			if sessionUser.Username != authService.GetUsernameFromSessions(sessionToken) {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			// Check for strong password
			ok, errMsg := utils.IsStrongPassword(changePasswordReq.NewPassword)
			if !ok {
				http.Error(w, errMsg, http.StatusForbidden)
				return
			}
			// Check if the old password matches
			if sessionUser.Password != changePasswordReq.OldPassword {
				http.Error(w, "Invalid old password", http.StatusBadRequest)
				return
			}

			// Change the user's password
			err = authService.ChangePassword(sessionUser.Username, changePasswordReq.NewPassword)
			if err != nil {
				// Password change failed
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Password change successful
			w.WriteHeader(http.StatusOK)
			return
		}

		// Return an empty response for GET requests
		w.WriteHeader(http.StatusOK)
	}
}

// ChangePassword changes the password of the user with the provided username

func (a *AuthService) ChangePassword(username, newPassword string) error {
	for _, user := range a.users {
		if user.Username == username {
			hashedPswd, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			user.Password = string(hashedPswd)

			// Update the password in the database
			query := "UPDATE users SET password = $1 WHERE username = $2"
			_, err = a.db.db.Exec(query, hashedPswd, username)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("User not found")
}

func HandleDeleteUser(authService *AuthService) http.HandlerFunc {
	type DeleteUserRequest struct {
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Check if the user is authenticated
			authToken := r.Header.Get("Authorization")
			if authToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Get the session user from the token
			sessionUser := authService.GetSessionUser(authToken)
			if sessionUser == nil {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			// Read the request body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}

			// Parse the request body into a DeleteUserRequest struct
			var deleteUserReq DeleteUserRequest
			err = json.Unmarshal(body, &deleteUserReq)
			if err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			// Verify the user's password
			if sessionUser.Password != deleteUserReq.Password {
				http.Error(w, "Invalid password", http.StatusUnauthorized)
				return
			}

			// Delete the user
			err = authService.DeleteUser(sessionUser.Username)
			if err != nil {
				// User deletion failed
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// User deletion successful
			w.WriteHeader(http.StatusOK)
			return
		}

		// Return an empty response for GET requests
		w.WriteHeader(http.StatusOK)
	}
}

func (a *AuthService) DeleteUser(username string) error {
	// Find the user index in the list of users
	userIndex := -1
	for i, user := range a.users {
		if user.Username == username {
			userIndex = i
			break
		}
	}

	// Check if the user exists
	if userIndex == -1 {
		return errors.New("User not found")
	}

	// Delete the user from the list of users
	a.users = append(a.users[:userIndex], a.users[userIndex+1:]...)

	// Delete the user from the sessions
	a.sessionsMu.Lock()
	delete(a.sessions, username)
	a.sessionsMu.Unlock()

	// Delete the user from the database
	query := "DELETE FROM users WHERE username = $1"
	_, err := a.db.db.Exec(query, username)
	if err != nil {
		return err
	}

	return nil
}

func HandleSignout(authService *AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Check if the user is authenticated
			authToken := r.Header.Get("Authorization")
			if authToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Get the session user from the token
			sessionUser := authService.GetSessionUser(authToken)
			if sessionUser == nil {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			// Remove the session token from the sessions map
			authService.RemoveSession(authToken)

			// User signout successful
			w.WriteHeader(http.StatusOK)
			return
		}

		// Return an empty response for GET requests
		w.WriteHeader(http.StatusOK)
	}
}

// RemoveSession removes the session with the given authToken from the sessions map
func (a *AuthService) RemoveSession(authToken string) {
	a.sessionsMu.Lock()
	defer a.sessionsMu.Unlock()

	delete(a.sessions, authToken)
}
