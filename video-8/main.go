package main

import (
	"github.com/gofiber/fiber"
	"log"
	"strconv"
)

var d int

func main() {
	app := fiber.New()

	r := app.Group("/temperatures")

	r.Get("/local", func(c *fiber.Ctx) {
		c.JSON(buildDegrees())
	})

	r.Post("/local/:temp", func(c *fiber.Ctx) {
		temperature := c.Params("temp")
		_temperature, err := strconv.Atoi(temperature)

		if err != nil {
			c.Status(400)
			log.Printf(err.Error())
			return
		}
		log.Printf("the temperature is %d", _temperature)
		setD(_temperature)
	})

	log.Fatal(app.Listen(9192))
}

type Temperature struct {
	Degrees int `json:"degrees"`
}

func buildDegrees() Temperature {
	return Temperature{
		Degrees: d,
	}
}

func setD(_d int) {
	d = _d
}
