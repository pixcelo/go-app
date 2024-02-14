package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://graph.instagram.com/me?fields=id,username&access_token=YOUR_ACCESS_TOKEN"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result)
}
