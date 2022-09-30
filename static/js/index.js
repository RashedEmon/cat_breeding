async function getCatByBreed(id){
  
  const url=`http://127.0.0.1:8080/cat/${id}`;
  console.log(url)
  try{
    const res= await fetch(url)
    // console.log(jsonData)
    return jsonData= await res.json()
  }catch{
    console.log("error while requesting")
  }
  return "error"
}
function descriptionCreator(data){
    return `<div>
    <h1>${data.Name}</h1>
    <h3>id:${data.ID}</h3>
    <p class="text-left">${data.Description}</p>
    <i>${data.Temperament}</i>
    <p>${data.Origin}</p>
    <p>${data.Weight}</p>
    <p>${data.LifeSpan}</p>
    </div>
    <span class="p-2 text-left"> <a href="${data.Wikipedia}">Wikipedia</a> </span>`
}

function courouselImageCreator(data){
    res=''
    data.Images.forEach((element,idx) => {
        if(idx==0){
            res+=`<div class="carousel-item h-[50vh] object-cover active">
                <img class="d-block w-100" src="${element}" alt="${data.Name}" width="100%" height="100%" style="width:100%;height:100%;">
            </div>`
        }else{
            res+=`<div class="carousel-item h-[50vh] object-cover">
                <img class="d-block w-100" src="${element}" alt="${data.Name}" width="100%" height="100%" style="width:100%;height:100%;">
                </div>`
        }
    });
  return res
}

function corouselIndicatorCreator(data){
    res=''
    data.Images.forEach((element,idx) => {
        if(idx==0){
            res+=`<li data-target="#carouselExampleIndicators" data-slide-to="${idx}" class="active"></li>`
        }else{
            res+=`<li data-target="#carouselExampleIndicators" data-slide-to="${idx}></li>`
        }
    });
    return res
}


document.getElementById("inputGroupSelect01").addEventListener('change',async (e)=>{
    console.log(e.target.value)
    
    const data=await getCatByBreed(e.target.value)
    document.getElementById('corousel_indicator').innerHTML=corouselIndicatorCreator(data)
    document.getElementById('corousel_image').innerHTML=courouselImageCreator(data)
    document.getElementById('description').innerHTML=descriptionCreator(data)
})
    
