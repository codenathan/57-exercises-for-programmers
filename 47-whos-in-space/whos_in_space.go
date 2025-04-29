package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

type Astronaut struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

type AstronautData struct {
	Number  int         `json:"number"`
	People  []Astronaut `json:"people"`
	Message string      `json:"message"`
}

func main() {
	resp, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var data AstronautData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if data.Message != "success" {
		fmt.Println("API request failed:", data.Message)
		return
	}

	fmt.Printf("There are %d people in space right now:\n", data.Number)

	// Group astronauts by craft
	craftGroups := make(map[string][]Astronaut)
	for _, astronaut := range data.People {
		craftGroups[astronaut.Craft] = append(craftGroups[astronaut.Craft], astronaut)
	}

	// Sort crafts alphabetically
	var crafts []string
	for craft := range craftGroups {
		crafts = append(crafts, craft)
	}
	sort.Strings(crafts)

	// Calculate column widths
	maxNameLength := 4  // "Name" is 4 chars
	maxCraftLength := 5 // "Craft" is 5 chars
	for _, astronauts := range craftGroups {
		for _, astronaut := range astronauts {
			if len(astronaut.Name) > maxNameLength {
				maxNameLength = len(astronaut.Name)
			}
		}
		if len(astronauts[0].Craft) > maxCraftLength {
			maxCraftLength = len(astronauts[0].Craft)
		}
	}

	headerFormat := fmt.Sprintf("%%-%ds | %%-%ds\n", maxNameLength, maxCraftLength)
	fmt.Printf(headerFormat, "Name", "Craft")

	separator := strings.Repeat("-", maxNameLength) + "-|-" + strings.Repeat("-", maxCraftLength)
	fmt.Println(separator)

	for _, craft := range crafts {
		astronauts := craftGroups[craft]

		sort.Slice(astronauts, func(i, j int) bool {
			return getLastName(astronauts[i].Name) < getLastName(astronauts[j].Name)
		})

		for _, astronaut := range astronauts {
			fmt.Printf(headerFormat, astronaut.Name, astronaut.Craft)

		}
	}
}

func getLastName(fullName string) string {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
