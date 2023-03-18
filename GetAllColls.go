package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	base_url := "https://api.nftport.xyz/v0/nfts?chain=ethereum&page_size=50&include=all"
	req, _ := http.NewRequest("GET", base_url, nil)
	apiKey := goDotEnvVariable("APIKEY")

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func goDotEnvVariable(key string) string {

	err := godotenv.Load("nftxyz.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
