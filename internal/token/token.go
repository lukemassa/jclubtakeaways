package token

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Tokener struct {
	secretKey *rsa.PrivateKey
}

func New(rawSecretKey string) (Tokener, error) {
	secretKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(rawSecretKey))
	if err != nil {
		return Tokener{}, err
	}
	return Tokener{
		secretKey: secretKey,
	}, nil
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (t Tokener) getAssertion() (string, error) {
	alg := jwt.GetSigningMethod("RS256")
	tok := jwt.New(alg)
	start := time.Now()
	end := start.Add(50 * time.Minute)
	tok.Claims = jwt.MapClaims{
		"aud":   "https://www.googleapis.com/oauth2/v4/token",
		"scope": "https://www.googleapis.com/auth/spreadsheets",
		"exp":   end.Unix(),
		"iss":   "webclientaccount@calm-bliss-188620.iam.gserviceaccount.com",
		"iat":   start.Unix(),
	}
	return tok.SignedString(t.secretKey)
}

func getTokenFromAPI(assertion string) (string, error) {

	jsonValue, _ := json.Marshal(map[string]string{
		"grant_type": "urn:ietf:params:oauth:grant-type:jwt-bearer",
		"assertion":  assertion,
	},
	)

	resp, err := http.Post("https://www.googleapis.com/oauth2/v4/token", "appliction/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return "", err
	}

	tok := tokenResponse{}
	err = json.Unmarshal(body, &tok)
	if err != nil {
		return "", err
	}

	return tok.AccessToken, nil

}

// GetToken get a token
func (t Tokener) Get() (string, error) {

	assertion, err := t.getAssertion()
	if err != nil {
		return "", fmt.Errorf("generating assertion: %v", err)
	}

	token, err := getTokenFromAPI(assertion)
	if err != nil {
		return "", fmt.Errorf("calling API: %v", err)
	}

	return token, nil
}

func (t Tokener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tok, err := t.Get()
	res := fmt.Sprintf("{\"token\":\"%s\",\"error\":\"\"}", tok)
	if err != nil {
		res = fmt.Sprintf("{\"token\":\"\",\"error\":\"%s\"}", err.Error())
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, res)
}
