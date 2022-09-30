package controllers

import (
	"cat/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/astaxie/beego"
)

type SearchByIdController struct {
	beego.Controller
}

func (c *SearchByIdController) Get() {
	// urlByid := "https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=abys&api_key=live_dnADHUMZofKkKEslnylF3VJTmvafvStFpbjjOHxuhAXVIjSRAYN4uUJnR5JZcXT1"
	aBreed := &models.CatBreed{}
	fmt.Println("search by id controller hit")
	firstBreedId := c.Ctx.Input.Param(":id")
	fmt.Println(firstBreedId)
	resp, searchByIdError := http.Get("https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=" + firstBreedId + "&api_key=live_dnADHUMZofKkKEslnylF3VJTmvafvStFpbjjOHxuhAXVIjSRAYN4uUJnR5JZcXT1")

	if searchByIdError != nil {
		fmt.Println("Error while requesting search by id")
	}
	dataOfaBreed, errrrr := io.ReadAll(resp.Body)

	if errrrr != nil {
		fmt.Println("Error while reading search by id from body")
	}
	parseABreedError := json.Unmarshal(dataOfaBreed, aBreed)
	if parseABreedError != nil {
		fmt.Print("something happend while parsing json data")
	}
	if len(*aBreed) <= 0 {
		c.Data["json"] = "{'error':'not found'}"
		c.ServeJSON()
	}
	var breedInfo = make(map[string]interface{})
	breedInfo["Name"] = (*aBreed)[0].Breeds[0].Name
	breedInfo["ID"] = (*aBreed)[0].Breeds[0].ID
	breedInfo["Description"] = (*aBreed)[0].Breeds[0].Description
	breedInfo["Temperament"] = (*aBreed)[0].Breeds[0].Temperament
	breedInfo["Origin"] = (*aBreed)[0].Breeds[0].Origin
	breedInfo["Weight"] = (*aBreed)[0].Breeds[0].Weight.Metric
	breedInfo["LifeSpan"] = (*aBreed)[0].Breeds[0].LifeSpan
	breedInfo["Wikipedia"] = (*aBreed)[0].Breeds[0].WikipediaURL
	var images []string
	for i := 0; i < len(*aBreed); i++ {
		images = append(images, (*aBreed)[i].URL)
	}
	breedInfo["Images"] = images
	c.Data["json"] = &breedInfo
	// c.Data["json"] = "{'name':'hello'}"
	c.ServeJSON()
}
