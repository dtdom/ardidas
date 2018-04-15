package main

import (
	"strings"
	"time"

	"github.com/dtdom/ardidas/config"
	"github.com/dtdom/ardidas/models"
)

func updateInstaImages() {
	for true {
		items := models.GetAllItems()

		for _, item := range items {
			images := ardidas_insta.GetValidImages(item.Hashtag)
			item.ArrayInsta = strings.Join(images[:], ",")
			item.Store()
		}

		time.Sleep(1 * time.Hour)
	}
}

func main() {
	go updateInstaImages()

	r := Runner{}
	r.Initialize()
	r.Run(":" + config.MainConfig.Server.Port)
}
