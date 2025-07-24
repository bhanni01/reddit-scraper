package auth
import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"io"
	"net/http"
	"net/url"
	"os"
	"github.com/joho/godotenv"
)

type Result struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
    Scope       string `json:"scope"`
}

func GetAccessToken(){
	err := godotenv.Load()
	fmt.Println("Let' see err: ", err)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	username := os.Getenv("REDDIT_USERNAME")
	password := os.Getenv("REDDIT_PASSWORD")
// concatenates your reddit Oauth client id and client secret with a colon : in between 
// []byte converts that concatenated string into a byte slice as EncodeToString expects bytes
// base64.StdEncoding.EncodeToString() takes those bytes and encodes them into a base64string , base64 encoding converts binary data into ascii text that is safe for transmission in http headers
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
// type Values map[string][]string , it is a map type designed to hold URL-encoded form values 

	data := url.Values{}
	// fmt.Println("before adding anything:  ",data)
	data.Set("grant_type","password")
	data.Set("username",username)
	data.Set("password",password)
	fmt.Println("after adding values: ",data)
	fmt.Println("after adding values encoded: ",data.Encode())

	// understood till here 

	// 
	req,err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token",bytes.NewBufferString(data.Encode()))
	if err != nil{
		panic(err)
	}

	req.Header.Set("Authorization","Basic "+ auth)
	req.Header.Set("User-Agent","reddit-scrapper by /u/"+username)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")

	client := &http.Client{}
	resp,err := client.Do(req)
	fmt.Println(" Status code:", resp.StatusCode)
	
	bodyBytes,err := io.ReadAll(resp.Body)
	fmt.Println("Response body : ", string(bodyBytes))
	
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Access token:", result["access_token"])

	

}