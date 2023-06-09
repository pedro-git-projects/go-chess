package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/pedro-git-projects/projeto-integrado-frontend/chessapi/pkg/token"
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
				response := struct {
					Token string `json:"token"`
				}{
					Token: token,
				}
				// Return the token in the response
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
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
