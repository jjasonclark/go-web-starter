package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

var (
	authHeader        = http.CanonicalHeaderKey("Authorization")
	contentTypeHeader = http.CanonicalHeaderKey("Content-Type")
	secretKey         = []byte("hello, world")
)

type loginSuccess struct {
	Token string `json:"token"`
	Error error  `json:"error,omitempty"`
}

func loginHandler() http.Handler {
	return http.HandlerFunc(loginHandlerFunc)
}

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.URL.Path == "/login" && r.Method == "POST" {
		handleLoginPost(w, r)
	} else if r.URL.Path == "/login" && r.Method == "DELETE" {
		handleLogout(w, r)
	} else if r.URL.Path == "/login/create" && r.Method == "POST" {
		handleRegisterPost(w, r)
	} else {
		notFoundHandlerFunc(w, r)
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	t := extractBearerToken(r.Header[authHeader])
	if t == "" {
		w.WriteHeader(http.StatusUnauthorized)
	} else if err := expireAuthToken(decodeToken(t)); err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleLoginPost(w http.ResponseWriter, r *http.Request) {
	var formData loginPost
	if err := decodeJSONRequest(r, &formData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commonAuthHandler(w, r, formData)
}

func handleRegisterPost(w http.ResponseWriter, r *http.Request) {
	var formData registerPost
	if err := decodeJSONRequest(r, &formData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commonAuthHandler(w, r, formData)
}

func getSignature(token []byte) []byte {
	mac := hmac.New(sha512.New, secretKey)
	mac.Write([]byte(token))
	return mac.Sum(nil)
}

func verifySignature(token, signature []byte) bool {
	return hmac.Equal(getSignature(token), signature)
}

func encodeToken(token string) (string, error) {
	tokenBytes := []byte(token)
	buf := bytes.NewBuffer(tokenBytes)
	buf.WriteRune('.')
	buf.Write(getSignature(tokenBytes))
	return base64.URLEncoding.EncodeToString(buf.Bytes()), nil
}

func decodeToken(token string) string {
	r, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return ""
	}
	data := bytes.SplitN(r, []byte{byte('.')}, 2)
	if len(data) < 2 {
		return ""
	}
	tokenID := data[0]
	signature := data[1]
	if !verifySignature(tokenID, signature) {
		return ""
	}
	return string(tokenID)
}

func commonAuthHandler(w http.ResponseWriter, r *http.Request, formData postHandler) {
	userID, err := formData.handlePost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tokenID, genError := generateAuthToken(userID)
	if genError != nil {
		http.Error(w, genError.Error(), http.StatusBadRequest)
		return
	}
	returnToken, signError := encodeToken(tokenID)
	if signError != nil {
		http.Error(w, signError.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Add(contentTypeHeader, JSONContentType)
	result, marshalError := json.Marshal(loginSuccess{
		Token: returnToken,
		Error: err,
	})
	if marshalError != nil {
		http.Error(w, marshalError.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(result))
}
