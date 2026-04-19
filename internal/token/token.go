package token

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Tokener struct{}

func New() Tokener {
	return Tokener{}
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func getKey() (*rsa.PrivateKey, error) {
	content := os.Getenv("WEB_CLIENT_KEY")

	decoded, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(decoded)
}

func getAssertion() (string, error) {
	key, err := getKey()
	if err != nil {
		return "", err
	}
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
	return tok.SignedString(key)
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
func (t Tokener) Get() string {

	assertion, err := getAssertion()
	if err != nil {
		return fmt.Sprintf("Error generating assertion: %v", err)
	}

	token, err := getTokenFromAPI(assertion)
	if err != nil {
		return fmt.Sprintf("Error calling API: %v", err)
	}

	return token
}

func (t Tokener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tok := t.Get()
	res := fmt.Sprintf("{\"token\":\"%s\",\"error\":\"\"}", tok)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, res)
}
