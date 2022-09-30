package controllers

import (
	"cat/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// urlByid := "https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=abys&api_key=live_dnADHUMZofKkKEslnylF3VJTmvafvStFpbjjOHxuhAXVIjSRAYN4uUJnR5JZcXT1"
	breedsUrl := "https://api.thecatapi.com/v1/breeds"
	breeds := &models.Breeds{}
	aBreed := &models.CatBreed{}
	fmt.Println("root hit")
	res, err := http.Get(breedsUrl)
	if err != nil {
		fmt.Print("Some error happend while loading data from server")
	}
	// res.Body.Close()
	data, errr := io.ReadAll(res.Body)
	if errr != nil {
		fmt.Print("error happend while reading...")
	}
	parseError := json.Unmarshal(data, breeds)
	if parseError != nil {
		fmt.Print("something happend while parsing json data")
	}
	// fmt.Println(len(*breeds))
	numberOfBreeds := 10
	if len(*breeds) < numberOfBreeds {
		numberOfBreeds = len(*breeds)
	}
	var breedsNameAndId = make(map[string]interface{})
	for i := 0; i < numberOfBreeds; i++ {
		breedsNameAndId[(*breeds)[i].Name] = (*breeds)[i].ID
	}
	firstBreedId := ""
	if len(*breeds) > 0 {
		firstBreedId = (*breeds)[0].ID
	}

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

	c.Data["BreedInfo"] = breedInfo
	// fmt.Print(breedsNameAndId)
	// fmt.Println(string(data))
	c.Data["Website"] = "rashedul.me"
	c.Data["Email"] = "rashedulb13@gmail.com"
	c.Data["breeds"] = breedsNameAndId
	c.TplName = "index.tpl"
}
