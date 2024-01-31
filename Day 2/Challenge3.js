function cetakSegitiga(){
    var bintang = "";
    for(var i = 5; i>0; i--){
        for(var j = 0; j<=5; j++){
            if(j >= i){
                bintang += "* ";
            }else{
                bintang += " ";
            }
        }
        bintang += "\n";
    }
    console.log(bintang);
}
cetakSegitiga();