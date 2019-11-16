package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var TeslaURL = "https://owner-api.teslamotors.com/oauth/token"
var TeslaClientID = "e4a9949fcfa04068f59abb5a658f2bac0a3428e4652315490b659d5ab3f35a9e"
var TeslaClientSecret = "c75f14bbadc8bee3a7594412c31416f8300256d7668ea7e6e7f06727bfb9d220"

type Request struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type Response struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
}

func main() {
	if os.Getenv("TESLA_EMAIL") == "" {
		log.Fatalln("Please set the TESLA_EMAIL environment variable")
	}
	if os.Getenv("TESLA_PASSWORD") == "" {
		log.Fatalln("Please set the TESLA_PASSWORD environment variable")
	}

	req := new(bytes.Buffer)
	if err := json.NewEncoder(req).Encode(&Request{
		GrantType:    "password",
		ClientID:     TeslaClientID,
		ClientSecret: TeslaClientSecret,
		Email:        os.Getenv("TESLA_EMAIL"),
		Password:     os.Getenv("TESLA_PASSWORD"),
	}); err != nil {
		log.Fatalln("Failed to encode token request:", err)
	}

	resp, err := http.Post(TeslaURL, "application/json", req)
	if err != nil {
		log.Fatalln("Failed to make token request:", err)
	}
	switch resp.StatusCode {
	case 401:
		log.Println("Authentication failed")
	case 404:
		log.Println("URL not found")
	case 500:
		log.Println("Server error:", resp.Status)
	case 200:
		processResponse(resp.Body)
	}
	resp.Body.Close()
}

func processResponse(r io.Reader) {
	resp := new(Response)
	if err := json.NewDecoder(r).Decode(resp); err != nil {
		log.Fatalln("Failed to decode response:", err)
	}
	log.Println("Token: ", resp.AccessToken)
}
