package main 
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

)
// Post represent the JSON structure we'll get 

type Post struct{
	UserID int  `json:"userId`
	ID int 	`json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func main(){
	url := "https://jsonplaceholder.typicode.com/posts/1"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	} 
	// schedule closing this network connection to happen after the current function (i.e. main) finishes
	defer resp.Body.Close()
	// defer is useful because even after my code hits an error later , go still ensures resp.Body.Close() runs 

	var post Post 
	// API returns data in JSON format and resp.Body is that,  .NewDecoder creates a decoder that reads the JSON stream , .Decode(&post) tells go to take the JSON butes and covert them into Go post structure i defined

	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	fmt.Printf("Post ID: %d\nTitle: %s\nBody: %s\n",post.ID, post.Title,post.Body)
}