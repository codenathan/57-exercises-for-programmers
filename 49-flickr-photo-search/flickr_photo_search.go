package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type FlickrItem struct {
	Title string `json:"title"`
	Media struct {
		M string `json:"m"` // image URL
	} `json:"media"`
}

type FlickrFeed struct {
	Title string       `json:"title"`
	Items []FlickrItem `json:"items"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("frontend.html")
	})

	r.GET("/api/flickr", func(c *gin.Context) {
		tag := c.Query("tag")
		if tag == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tag is required"})
			return
		}

		url := "https://www.flickr.com/services/feeds/photos_public.gne?format=json&tags=" + tag + "&nojsoncallback=true"
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch from Flickr"})
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close Flickr response"})
				return
			}
		}(resp.Body)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Flickr response"})
			return
		}

		var feed FlickrFeed
		if err := json.Unmarshal(body, &feed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
			return
		}

		var images []string
		for _, item := range feed.Items {
			images = append(images, item.Media.M)
		}

		c.JSON(http.StatusOK, images)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
