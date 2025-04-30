package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type MovieData struct {
	Title   string       `json:"Title"`
	Year    string       `json:"Year"`
	Rated   string       `json:"Rated"`
	Runtime string       `json:"Runtime"`
	Plot    string       `json:"Plot"`
	Ratings []RatingItem `json:"Ratings"`
}

type RatingItem struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func getMovieData(movieName string) (*MovieData, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	apiKey := os.Getenv("OMD_API_KEY")

	baseURL := "https://www.omdbapi.com/"
	query := url.Values{}
	query.Add("apikey", apiKey)
	query.Add("t", movieName)

	resp, err := http.Get(baseURL + "?" + query.Encode())
	if err != nil {
		fmt.Println("Error fetching movie data: %v", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("API request failed with status: %s", resp.Status)
		return nil, errors.New("API request failed with status:" + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var data MovieData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("error parsing movie data: %v", err)
	}

	return &data, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name of the movie: ")
	movieName, _ := reader.ReadString('\n')
	movieName = strings.TrimSpace(movieName)

	movie, err := getMovieData(movieName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Title:", movie.Title)
	fmt.Println("Year:", movie.Year)
	fmt.Println("Rating:", movie.Rated)
	fmt.Println("Runtime:", movie.Runtime)
	fmt.Println("Description:", movie.Plot)

	for _, rating := range movie.Ratings {
		if rating.Source == "Rotten Tomatoes" {
			// Parse the percentage value, e.g., "94%"
			percentStr := strings.TrimSuffix(rating.Value, "%")
			score, err := strconv.Atoi(percentStr)
			if err != nil {
				fmt.Println("Could not parse Rotten Tomatoes score")
				break
			}

			if score >= 80 {
				fmt.Println("You should watch this movie right now!")
			} else if score < 50 {
				fmt.Println("Avoid this movie at all costs.")
			} else {
				fmt.Println("This movie is okay — watch it if you’re curious.")
			}
			break
		}
	}

}
