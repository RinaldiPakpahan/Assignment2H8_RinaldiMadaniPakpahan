
function balik(kata){
    var kataBaru = "";
    for(i = kata.length -1; i>=0; i--){
        kataBaru += kata[i];
    }
    if(kataBaru == kata){
        console.log("palingdrom")
    }else{
        console.log(kataBaru);
    }
}

balik("pakpahan");