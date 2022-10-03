<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="src/style.css">
    <script src="https://cdn.tailwindcss.com"></script>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
  </head>
  <body>
    
    <div class="container h-auto">
           <div class="row">
              <div class="col-12 offset-0 col-lg-6 offset-lg-3">
                  <div class="card text-center">
                    <div class="card-header">
                        <div class="input-group mb-3">
                        <div class="input-group-prepend">
                          <label class="input-group-text" for="inputGroupSelect01">Breeds</label>
                        </div>
                        <select class="custom-select" id="inputGroupSelect01">
                            {{range $key,$val := .breeds}}
                                <option value={{$val}}>{{$key}}</option>
                            {{end}}
                        </select>
                      </div>
                    </div>
                    <div class="card-body">
                        <div id="carouselExampleIndicators" class="carousel slide" data-ride="carousel">
                                <ol class="carousel-indicators" id="corousel_indicator">
                                  {{range $idx,$val := .BreedInfo.Images}}
                                    <li data-target="#carouselExampleIndicators" data-slide-to="{{$idx}}" {{if eq $idx 0}}class="active"{{end}}></li>
                                  {{end}}
                                </ol>
                                <div class="carousel-inner" id="corousel_image">
                                  {{range $idx,$val := .BreedInfo.Images}}
                                
                                    <div class="carousel-item h-[50vh] object-cover {{if eq $idx 0}}active{{end}}">
                                      <img class="d-block w-100" src="{{$val}}" alt="cat image" width="100%" height="100%" style="width:100%;height:100%;">
                                    </div>
                                  {{end}}
                    
                                </div>
                                <a class="carousel-control-prev" href="#carouselExampleIndicators" role="button" data-slide="prev">
                                  <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                                  <span class="sr-only">Previous</span>
                                </a>
                                <a class="carousel-control-next" href="#carouselExampleIndicators" role="button" data-slide="next">
                                  <span class="carousel-control-next-icon" aria-hidden="true"></span>
                                  <span class="sr-only">Next</span>
                                </a>
                        </div>
                    </div>
                    <div class="card-footer text-muted" id="description">
                      <div >
                        <h1 class="text-dark text-2">{{.BreedInfo.Name}}</h1>
                        <h3 class="text-dark">id:{{.BreedInfo.ID}}</h3>
                        <p class="text-left">{{.BreedInfo.Description}}</p>
                        <i>{{.BreedInfo.Temperament}}</i>
                        <p>{{.BreedInfo.Origin}}</p>
                        <p>{{.BreedInfo.Weight}} kgs</p>
                        <p>{{.BreedInfo.LifeSpan}} average life span</p>
                      </div>
                      <span class="p-2 text-left"> <a href="{{.BreedInfo.Wikipedia}}">Wikipedia</a> </span>
                    </div>
                      
                  </div>
              </div>
           </div>
    </div>
      <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
      <script src="https://cdn.jsdelivr.net/npm/popper.js@1.12.9/dist/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
      <script src="https://cdn.jsdelivr.net/npm/bootstrap@4.0.0/dist/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
      <script src="../static/js/index.js"></script>
    </body>
</html>