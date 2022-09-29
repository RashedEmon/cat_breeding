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
	breeds := &models.CatBreed{}
	fmt.Println("root hit")
	res, err := http.Get("https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=abys&api_key=live_dnADHUMZofKkKEslnylF3VJTmvafvStFpbjjOHxuhAXVIjSRAYN4uUJnR5JZcXT1")
	if err != nil {
		fmt.Print("Some error happend while loading data from server")
	}
	// res.Body.Close()
	data, errr := io.ReadAll(res.Body)
	if errr != nil {
		fmt.Print("error happend while reading...")
	}
	json.Unmarshal(data, breeds)
	fmt.Println(len(*breeds))
	for i := 0; i < len(*breeds); i++ {
		fmt.Println((*breeds)[i].Breeds[0].Name)
	}

	// fmt.Println(string(data))
	c.Data["Website"] = "rashedul.me"
	c.Data["Email"] = "rashedulb13@gmail.com"
	c.Data["breads"] = string(data)
	c.TplName = "index.tpl"
}
