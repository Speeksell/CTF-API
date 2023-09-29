package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	baseURL       = "http://10.49.122.144:"
	startPort     = 1
	endPort       = 65535
	timeout       = 1 * time.Second
	concurrent    = 1000 // nombre de goroutines concurrentes
	username      = "Louis"
	maxIterations = 100
)

var secret string

type RequestBody struct {
	User string `json:"User"`
}

type RequestSecret struct {
	User   string `json:"User"`
	Secret string `json:"Secret"`
}

type Challenge struct {
	Username string `json:"Username"`
	Secret   string `json:"Secret"`
	Points   int    `json:"Points"`
}

type Content struct {
	Level     int       `json:"Level"`
	Challenge Challenge `json:"Challenge"`
}

type UserSecretContent struct {
	User      string `json:"User"`
	Secret    string `json:"Secret"`
	Protocol  string `json:"Protocol"`
	SecretKey string `json:"SecretKey"`
}

// type RequestSubmition struct {
// 	Username string `json:"Username"`
// 	Secret  string  `json:"Secret"`
// 	Content map[string]string{
// 		Level     string       `json:"Level"`
// 		Challenge map[string]string {
// 			Secret   string `json:"Secret"`
// 			Points   string    `json:"Points"`
// 			User string `json:"User"`
// 		}
// 	}
// 	Protocol  string    `json:"Protocol"`
// 	SecretKey string    `json:"SecretKey"`
// }

func handleSignUp(client *http.Client, port int, username string) {
	url := fmt.Sprintf("%s%d/signup", baseURL, port)
	body := RequestBody{User: username}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /signup:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Réponse reçue depuis /signup sur le port %d: %s\n", port, string(responseBody))
}

func handleCheck(client *http.Client, port int, username string) {
	url := fmt.Sprintf("%s%d/check", baseURL, port)
	body := RequestBody{User: username}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /check:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Réponse reçue depuis /check sur le port %d: %s\n", port, string(responseBody))
}

func handleGetUserSecret(client *http.Client, port int, username string) string {
	url := fmt.Sprintf("%s%d/getUserSecret", baseURL, port)
	body := RequestBody{User: username}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /getUserSecret:", err)
		return ""
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Réponse reçue depuis /getUserSecret sur le port %d: %s\n", port, string(responseBody))
	return string(responseBody)
}

func handleGetUserLevel(client *http.Client, port int, username, secret string) {
	url := fmt.Sprintf("%s%d/getUserLevel", baseURL, port)
	body := RequestSecret{User: username, Secret: secret}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /getUserLevel:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Réponse reçue depuis /getUserLevel sur le port %d: %s\n", port, string(responseBody))
}

func handleGetUserPoints(client *http.Client, port int, username, secret string) {
	url := fmt.Sprintf("%s%d/getUserPoints", baseURL, port)
	body := RequestSecret{User: username, Secret: secret}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /getUserPoints:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Réponse reçue depuis /getUserPoints sur le port %d: %s\n", port, string(responseBody))
}

func handleINeedAHint(client *http.Client, port int, username, secret string) {
	url := fmt.Sprintf("%s%d/iNeedAHint", baseURL, port)
	body := RequestSecret{User: username, Secret: secret}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes)) // sending request without body
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /iNeedAHint:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Réponse reçue depuis /iNeedAHint sur le port %d: %s\n", port, string(responseBody))
}

func handleEnterChallenge(client *http.Client, port int, username, secret string) {
	url := fmt.Sprintf("%s%d/enterChallenge", baseURL, port)
	body := RequestSecret{User: username, Secret: secret}
	bodyBytes, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		fmt.Println("Erreur lors de la requête POST /enterChallenge:", err)
		return
	}
	defer resp.Body.Close()
	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Réponse reçue depuis /enterChallenge sur le port %d: %s\n", port, string(responseBody))
}

// func handleSubmitSolution(client *http.Client, port int, username, secret string, protocole string) {
// 	url := fmt.Sprintf("%s%d/submitSolution", baseURL, port)
// 	body := RequestSubmition{Content: Content, UserSecretContent: UserSecretContent}
// 	bodyBytes, _ := json.Marshal(body)
// 	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
// 	if err != nil {
// 		fmt.Println("Erreur lors de la requête POST /submitSolution:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	responseBody, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Printf("Réponse reçue depuis /submitSolution sur le port %d: %s\n", port, string(responseBody))
// }

func main() {
	client := &http.Client{
		Timeout: timeout,
	}

	var wg sync.WaitGroup
	ports := make(chan int, concurrent)

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				resp, err := client.Get(fmt.Sprintf("%s%d/ping", baseURL, port))
				if err == nil && resp.StatusCode == http.StatusOK {
					defer resp.Body.Close()
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Printf("Réponse reçue depuis le port %d: %s\n", port, string(body))

					handleSignUp(client, port, username)
					handleCheck(client, port, username)
					// Boucle sur handleGetUserSecret
					previousResponse := handleGetUserSecret(client, port, username)
					for i := 0; i < maxIterations; i++ {
						response := handleGetUserSecret(client, port, username)
						//fmt.Printf("Iteration %d, Réponse reçue depuis /getUserSecret sur le port %d: %s\n", i+1, port, response)
						if response != "" && response != previousResponse {
							secret = response
							fmt.Printf("Réponse reçue depuis /getUserSecret sur le port %d: %s\n", port, response)
							break
						}
					}
					fmt.Println(secret)
					trimmedSecretKey := strings.TrimSpace(strings.TrimPrefix(secret, "User secret:"))
					//userData["secret"] = trimmedSecretKey
					handleGetUserLevel(client, port, username, trimmedSecretKey)
					handleGetUserPoints(client, port, username, trimmedSecretKey)
					// Nouveau : Appel de handleINeedAHint
					handleINeedAHint(client, port, username, trimmedSecretKey)
					handleEnterChallenge(client, port, username, trimmedSecretKey)
					//handleSubmitSolution(client, port, username, trimmedSecretKey, protocole)

					return // On arrête dès qu'on trouve un port valide
				}
			}
		}()
	}

	// Remplissage du canal avec les numéros de port
	for port := startPort; port <= endPort; port++ {
		ports <- port
	}
	close(ports)

	wg.Wait()

}
