package controllers

import (
	"cat/models"
	"cat/utils"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

type SearchByIdController struct {
	beego.Controller
}

func (c *SearchByIdController) Get() {

	//declare a variable to store struct reference
	aBreed := &models.CatBreed{}
	//get data from url params
	firstBreedId := c.Ctx.Input.Param(":id")

	breedByIdUrl := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=%s&api_key=%s", firstBreedId, beego.AppConfig.String("api_key"))

	var dataOfaBreed []byte

	breed_by_id_channel := make(chan utils.Response)

	go utils.HttpGetRequest(breedByIdUrl, breed_by_id_channel)

	breed_by_id_data := <-breed_by_id_channel

	if breed_by_id_data.Err != nil {
		c.Data["json"] = "{'message':'Error while requesting search by id'}"
		c.ServeJSON()
	}

	fmt.Println(breed_by_id_data.Result)
	//convert json byte and store data to struct
	parseABreedError := json.Unmarshal(dataOfaBreed, aBreed)
	if parseABreedError != nil {
		c.Data["json"] = "{'message':'Something happend while parsing json data'}"
		c.ServeJSON()
	}
	//check that weather breed is empty or not
	if len(*aBreed) < 1 {
		c.Data["json"] = "{'message':'Empty Data'}"
		c.ServeJSON()
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
	var images []string
	for i := 0; i < len(*aBreed); i++ {
		images = append(images, (*aBreed)[i].URL)
	}
	breedInfo["Images"] = images

	c.Data["json"] = &breedInfo
	c.ServeJSON()
}
