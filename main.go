package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./frontend/build", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		urlQuery, err := url.QueryUnescape(c.Query("url"))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString("Invalid URL")
		}

		resp, err := http.Get(urlQuery)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error fetching URL")
		}

		defer resp.Body.Close()

		var result map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return c.Status(http.StatusInternalServerError).SendString("Error decoding response")
		}

		inverseObjects(result)
		inverseKeyValue(result)

		processedResult := map[string]interface{}{
			"originalResult":  result,
			"processedResult": result,
		}

		return c.JSON(processedResult)
	})

	log.Fatal(app.Listen(":8080"))
}

func inverseObjects(data map[string]interface{}) {
	for _, value := range data {
		if child, ok := value.(map[string]interface{}); ok {
			inverseObjects(child)
		}

		if array, ok := value.([]interface{}); ok {
			for _, item := range array {
				if child, ok := item.(map[string]interface{}); ok {
					inverseObjects(child)
				}
			}
		}
	}

	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}

	newData := make(map[string]interface{}, len(data))
	for _, key := range keys {
		newData[key] = data[key]
	}

	for key, value := range newData {
		data[key] = value
	}
}

func inverseKeyValue(data map[string]interface{}) {
	for key, value := range data {
		if child, ok := value.(map[string]interface{}); ok {
			inverseKeyValue(child)
		} else if array, ok := value.([]interface{}); ok {
			for _, item := range array {
				if child, ok := item.(map[string]interface{}); ok {
					inverseKeyValue(child)
				}
			}
		} else if str, ok := value.(string); ok {
			inverseString := inverse(str)
			data[key] = inverseString
		}
	}
}

func inverse(str string) string {
	var sb strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteByte(str[i])
	}
	return sb.String()
}
