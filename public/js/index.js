//get data from server
async function getCatByBreed(id){
  
  const url=`http://127.0.0.1:8080/cat/${id}`;
  try{
    const res= await fetch(url)
    // console.log(jsonData)
    return jsonData= await res.json()
  }catch{
    console.log("error while requesting")
  }
  return "error"
}
//create description div
function descriptionCreator(data){
    return `<div>
    <h1 class="text-dark text-2xl py-2">${data.Name}</h1>
    <h3 class="text-dark">id:${data.ID}</h3>
    <p class="text-left">${data.Description}</p>
    <i>${data.Temperament}</i>
    <p>${data.Origin}</p>
    <p>${data.Weight} kgs</p>
    <p>${data.LifeSpan} average life span</p>
    </div>
    <span class="p-2 text-left"> <a href="${data.Wikipedia}">Wikipedia</a> </span>`
}
//create courousel image
function courouselImageCreator(data){
    let res=''
    data.Images.forEach((element,idx) => {
        if(idx==0){
            res+=`<div class="carousel-item active" style="overflow: hidden;height:50vh">
                    <img class="d-block w-100" src="${element}" alt="${data.Name}" width="100%" height="100%" style="object-fit: cover;">
                </div>`
        }else{
            res+=`<div class="carousel-item" style="overflow: hidden;height:50vh">
                    <img class="d-block w-100" src="${element}" alt="${data.Name}" width="100%" height="100%" style="object-fit: cover;">
                </div>`
        }
    });
  return res
}
//create corousel indicator
function corouselIndicatorCreator(data){
    var res="";
    data.Images.forEach((element,idx) => {
        if(idx===0){
            res+=`<li data-target="#carouselExampleIndicators" data-slide-to="${idx}" class="active"></li>`
        }else{
            res+=`<li data-target="#carouselExampleIndicators" data-slide-to="${idx} class=""></li>`
        }
    });
    return res
}

//listen change of dropdown send request to server and set data to ui
document.getElementById("inputGroupSelect01").addEventListener('change',async (e)=>{    
    const data=await getCatByBreed(e.target.value)
    document.getElementById('corousel_indicator').innerHTML=corouselIndicatorCreator(data)
    document.getElementById('corousel_image').innerHTML=courouselImageCreator(data)
    document.getElementById('description').innerHTML=descriptionCreator(data)
})
    
