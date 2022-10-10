package controllers

import (
	"cat/models"
	"cat/utils"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	breedsUrl := "https://api.thecatapi.com/v1/breeds"
	breeds := &models.Breeds{}
	aBreed := &models.CatBreed{}

	breed_channel := make(chan utils.Response)

	go utils.HttpGetRequest(breedsUrl, breed_channel)
	breed_data := <-breed_channel
	fmt.Println(breed_data.Result)
	//take json data as byte and a struct. convert json to struct and set field to struct reference.
	parseError := json.Unmarshal(breed_data.Result, breeds)
	if parseError != nil {
		panic("something happend while parsing json data")
	}

	//check that weather breed is empty or not
	if len(*breeds) < 1 {
		panic("empty data")
	}
	//declare a map to contain breed name and id
	var breedsNameAndId = make(map[string]interface{})

	for i := 0; i < len(*breeds); i++ {
		breedsNameAndId[(*breeds)[i].Name] = (*breeds)[i].ID
	}
	//deaclare a string to hold first breed id
	firstBreedId := ""
	firstBreedId = (*breeds)[0].ID

	breedByIdUrl := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=%s&api_key=%s", firstBreedId, beego.AppConfig.String("api_key"))

	var dataOfaBreed []byte

	breed_by_id_channel := make(chan utils.Response)
	go utils.HttpGetRequest(breedByIdUrl, breed_by_id_channel)

	abc := <-breed_by_id_channel
	fmt.Println(abc)
	//convert json byte and store data to struct
	parseABreedError := json.Unmarshal(dataOfaBreed, aBreed)
	if parseABreedError != nil {
		panic("something happend while parsing json data")
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
