package main

import (
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	r "github.com/dancannon/gorethink"
)

const (
	sessionLength  = 7 * 24 * time.Hour
	encryptionCost = 10
	authPrefix     = "Bearer "
	authPrefixLen  = len(authPrefix)
)

var (
	websiteScopes           = []string{"website"}
	errAuthenticationFailed = errors.New("Authentication Failed")
	errRegistrationFailed   = errors.New("Registration Failed")
	errTokenCreationFailed  = errors.New("Could not create login token")
	errUsernameExists       = errors.New("Username already exists")
)

type loginPost struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type registerPost struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type postHandler interface {
	handlePost() (string, error)
}

func (l registerPost) handlePost() (string, error) {
	if l.Password != l.PasswordConfirmation {
		return "", errRegistrationFailed
	}
	if err := usernameExists(l.Username); err != nil {
		return "", errUsernameExists
	}
	hashedPassword, err := encryptPassword(l.Password)
	if err != nil {
		return "", err
	}
	return createUser(l.Username, string(hashedPassword))
}

func (l loginPost) handlePost() (string, error) {
	c, err := r.Table("users").
		Filter(dbParams{"username": l.Username}).
		Pluck("id", "username", "password").
		Run(dbSession)
	if c != nil {
		defer c.Close()
	}
	if err != nil {
		return "", err
	}
	var user map[string]string
	err = c.One(&user)
	if err != nil {
		return "", err
	}
	userID, haveID := user["id"]
	password, havePassword := user["password"]
	if haveID && havePassword && validPassword([]byte(password), []byte(l.Password)) {
		return userID, nil
	}
	return "", errAuthenticationFailed
}

func usernameExists(username string) error {
	c, err := r.Table("users").
		Filter(dbParams{"username": username}).
		Pluck("id", "username").
		Run(dbSession)
	if c != nil {
		defer c.Close()
	}
	if err != nil && err != r.ErrEmptyResult {
		return err
	}
	var result map[string]string
	gotResult := c.Next(&result)
	if gotResult {
		return errUsernameExists
	}
	return nil
}

func createUser(username, password string) (string, error) {
	i, err := r.Table("users").Insert(dbParams{
		"username": username,
		"password": password,
	}).RunWrite(dbSession)
	if err != nil {
		return "", err
	}
	return i.GeneratedKeys[0], nil
}

func generateAuthToken(userID string) (string, error) {
	createdAt := time.Now().UTC()
	expiresAt := createdAt.Add(sessionLength).UTC()
	c, err := r.Table("tokens").
		Insert(dbParams{
			"userId":    userID,
			"createdAt": createdAt,
			"expiresAt": expiresAt,
			"scopes":    websiteScopes,
		}).
		RunWrite(dbSession)
	if err == nil && len(c.GeneratedKeys) > 0 && c.GeneratedKeys[0] != "" {
		return c.GeneratedKeys[0], nil
	}
	if err != nil {
		return "", err
	}
	return "", errTokenCreationFailed
}

func expireAuthToken(tokenID string) error {
	if tokenID == "" {
		return nil
	}
	_, err := r.Table("tokens").Get(tokenID).
		Update(dbParams{"expiresAt": time.Now().UTC()}).
		RunWrite(dbSession)
	return err
}

func extractBearerToken(values []string) string {
	for _, value := range values {
		length := len(value)
		prefix := value[0:min(authPrefixLen, length)]
		if strings.EqualFold(prefix, authPrefix) {
			return value[authPrefixLen:length]
		}
	}
	return ""
}

func validPassword(existing, attempt []byte) bool {
	err := bcrypt.CompareHashAndPassword(existing, attempt)
	return err == nil
}

func encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), encryptionCost)
}
