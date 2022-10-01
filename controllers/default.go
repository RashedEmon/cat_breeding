package controllers

import (
	"cat/models"
	"encoding/json"
	"fmt"

	myutils "cat/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	breedsUrl := "https://api.thecatapi.com/v1/breeds"
	breeds := &models.Breeds{}
	aBreed := &models.CatBreed{}

	//get byte data from server by providing url
	data, err := myutils.HttpGetRequest(breedsUrl)
	if err != nil {
		fmt.Println("Something went wrong while getting data from server")
	}
	//take json data as byte and a struct. convert json to struct and set field to struct reference.
	parseError := json.Unmarshal(data, breeds)
	if parseError != nil {
		fmt.Print("something happend while parsing json data")
	}

	//check that weather breed is empty or not
	if len(*breeds) < 1 {
		fmt.Println("empty data")
	}
	//declare a map to contain breed name and id
	var breedsNameAndId = make(map[string]interface{})

	for i := 0; i < len(*breeds); i++ {
		breedsNameAndId[(*breeds)[i].Name] = (*breeds)[i].ID
	}
	//deaclare a string to hold first breed id
	firstBreedId := ""
	firstBreedId = (*breeds)[0].ID

	breedByIdUrl := "https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=" + firstBreedId + "&api_key=live_dnADHUMZofKkKEslnylF3VJTmvafvStFpbjjOHxuhAXVIjSRAYN4uUJnR5JZcXT1"
	dataOfaBreed, breedByIdError := myutils.HttpGetRequest(breedByIdUrl)

	if breedByIdError != nil {
		fmt.Println("Something went wrong while getting data from server")
	}
	
	//convert json byte and store data to struct
	parseABreedError := json.Unmarshal(dataOfaBreed, aBreed)
	if parseABreedError != nil {
		fmt.Print("something happend while parsing json data")
	}
	//declare a map to store breed information
	var breedInfo = make(map[string]interface{})

	//set data to map
	breedInfo["Name"] = (*aBreed)[0].Breeds[0].Name
	breedInfo["ID"] = (*aBreed)[0].Breeds[0].ID
	breedInfo["Description"] = (*aBreed)[0].Breeds[0].Description
	breedInfo["Temperament"] = (*aBreed)[0].Breeds[0].Temperament
	breedInfo["Origin"] = (*aBreed)[0].Breeds[0].Origin
	breedInfo["Weight"] = (*aBreed)[0].Breeds[0].Weight.Metric
	breedInfo["LifeSpan"] = (*aBreed)[0].Breeds[0].LifeSpan
	breedInfo["Wikipedia"] = (*aBreed)[0].Breeds[0].WikipediaURL
	//declare a slice to hold image url
	var images []string
	for i := 0; i < len(*aBreed); i++ {
		images = append(images, (*aBreed)[i].URL)
	}
	breedInfo["Images"] = images

	//set data to pass template
	c.Data["BreedInfo"] = breedInfo
	c.Data["breeds"] = breedsNameAndId
	c.TplName = "index.tpl"
}
