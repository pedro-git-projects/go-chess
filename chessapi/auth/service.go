package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/token"
	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/utils"
)

type AuthService struct {
	users      []User
	sessions   map[string]*User
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
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func (a *AuthService) StartSession(username string) (authToken string, err error) {
	a.sessionsMu.Lock()
	defer a.sessionsMu.Unlock()

	authToken, err = token.Generate()
	if err != nil {
		return authToken, err
	}

	// Associate the session token with the user
	a.sessions[authToken] = &User{Username: username}

	return authToken, nil
}

func (a *AuthService) GetSessionUser(token string) *User {
	a.sessionsMu.RLock()
	defer a.sessionsMu.RUnlock()
	return a.sessions[token]
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
				token, err := authService.StartSession(loginReq.Username)
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

	// Insert the new user into the database
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := a.db.db.Exec(query, username, password)
	if err != nil {
		return err
	}

	// Create a new user
	newUser := User{
		Username: username,
		Password: password,
	}

	// Add the user to the list of users in memory
	a.users = append(a.users, newUser)

	return nil
}

func HandleChangePassword(authService *AuthService) http.HandlerFunc {
	type ChangePasswordRequest struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
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

			// Check if the user is authenticated
			sessionToken := r.Header.Get("Authorization")
			if sessionToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Get the session user from the token
			sessionUser := authService.GetSessionUser(sessionToken)
			if sessionUser == nil || sessionUser.Username != changePasswordReq.Username {
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}

			// Check if the old password matches
			if sessionUser.Password != changePasswordReq.OldPassword {
				http.Error(w, "Invalid old password", http.StatusBadRequest)
				return
			}

			// Change the user's password
			err = authService.ChangePassword(changePasswordReq.Username, changePasswordReq.NewPassword)
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
	// Find the user in the list of users
	for _, user := range a.users {
		if user.Username == username {
			// Update the user's password
			user.Password = newPassword

			// Update the password in the database
			query := "UPDATE users SET password = $1 WHERE username = $2"
			_, err := a.db.db.Exec(query, newPassword, username)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("User not found")
}
