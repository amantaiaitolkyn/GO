package unsplash

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"github.com/go-telegram/bot"
  	"github.com/go-telegram/bot/models"
	// "github.com/hashicorp/go-hclog"
)
const (
	unsplashAPIBaseURL = "https://api.unsplash.com"
	unsplashRandomPath = "/photos/random"
)

func GetRandomUnsplashImageURL() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, unsplashAPIBaseURL+unsplashRandomPath, nil)
	if err != nil {
	  return "", err
	}
	req.Header.Add("Authorization", "Client-ID APhyMWrOZfZbKgVkgLk9QNU-CgVXmtX43WqD88BBG8M")
  
	resp, err := client.Do(req)
	if err != nil {
	  return "", err
	}
	defer resp.Body.Close()
  
	var unsplashResp UnsplashResponse
	err = json.NewDecoder(resp.Body).Decode(&unsplashResp)
	if err != nil {
	  return "", err
	}
	return unsplashResp.URL.Regular, nil
}


// func (s service) GetRandomPhoto() (*RandomPhoto, error) {
// 	url := "https://api.unsplash.com/photos/random?client_id=" + s.credentials.ClientID

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		s.log.Error("failed making request", "error", err.Error())
// 		return nil, err
// 	}

// 	// making request
// 	resp, err := s.client.Do(req)
// 	if err != nil {
// 		s.log.Error("failed do request", "error", err.Error())
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	// handling response and unmarshalling
// 	randPhoto := &RandomPhoto{}
// 	response, err := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(response, randPhoto)
// 	if err != nil {
// 		s.log.Error("failed on unmarshalling response", "error", err.Error())
// 		return nil, err
// 	}

// 	return randPhoto, nil
// }