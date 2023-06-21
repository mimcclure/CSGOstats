package main

import( "github.com/joho/godotenv"
		"fmt"
		"log"
		"encoding/json"
		"io/ioutil"
		"net/http"
		"os"
)

type CSGOstats struct {
	playerStats struct{
		SteamID string `json:"steamID"`
		Stats []struct {
			Name string `json:"name"`
			Value string `json:"value"`
		} `json:"stats"`
	} `json:"playerstats"`
}

func main(){

	//Load KEY in .env file
	key := os.Getenv("KEY")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	//Enter steamID
	fmt.Print("Enter your SteamID: ")
	var steamID string
	fmt.Scanln(&steamID)

	//API endpoint
	userStatsURL := fmt.Sprintf("https://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v2/?key=%s&steamid=%s&appid=730", key, steamID)


	//GET HTTP request
	response, err := http.Get(userStatsURL)
	if err != nil{
		fmt.Println("Error during HTTP requests: ", err)
		return
	}
	defer response.Body.Close()

	//READ response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	//check response for errors
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error response from Steam API: ", response.Status)
		fmt.Println("Response body: ", string(body))
		return
	}

	//Parse response body
	var userStats CSGOstats
	err = json.Unmarshal(body, &userStats)
	if err != nil {
		fmt.Println("Error parsing response body: ", err)
		return
	}

	//Print player stats
	fmt.Println("Steam ID: %s\n", userStats.playerStats.SteamID)
	fmt.Println("Stats: ")
	for _, stat := range userStats.playerStats.Stats {
		fmt.Printf("- %s: %s\n", stat.Name, stat.Value)
	}
}