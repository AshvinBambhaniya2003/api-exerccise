package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Title struct {
	ID                  string   `json:"id"`
	Title               string   `json:"title"`
	Type                string   `json:"type"`
	Description         string   `json:"description"`
	ReleaseYear         int      `json:"release_year"`
	AgeCertification    string   `json:"age_certification"`
	Runtime             int      `json:"runtime"`
	Genres              []string `json:"genres"`
	ProductionCountries []string `json:"production_countries"`
	Seasons             int      `json:"seasons"`
	IMDbID              string   `json:"imdb_id"`
	IMDbScore           float64  `json:"imdb_score"`
	IMDbVotes           int      `json:"imdb_votes"`
	TmdbPopularity      float64  `json:"tmdb_popularity"`
	TmdbScore           float64  `json:"tmdb_score"`
}

var titles []Title

func CreateTitle(c *fiber.Ctx) error {
	var newTitle Title
	if err := c.BodyParser(&newTitle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	titles = append(titles, newTitle)

	return c.JSON(fiber.Map{
		"message": "Title created successfully",
		"title":   newTitle,
	})
}

func ListTitle(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"titles": titles,
	})
}

func UpdateTitle(c *fiber.Ctx) error {

	id := c.Params("id")

	// Find the title with the given ID
	var found bool
	for i, title := range titles {
		if title.ID == id {
			var updatedTitle Title
			if err := c.BodyParser(&updatedTitle); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Failed to parse request body",
				})
			}
			titles[i] = updatedTitle
			found = true
			break
		}
	}

	if !found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Title not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Title updated successfully",
	})
}

func DeleteTitle(c *fiber.Ctx) error {
	id := c.Params("id")

	// Find the index of the title with the given ID
	index := -1
	for i, title := range titles {
		if title.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Title not found",
		})
	}

	// Remove the title from the titles slice
	titles = append(titles[:index], titles[index+1:]...)

	return c.JSON(fiber.Map{
		"message": "Title deleted successfully",
	})
}
