function bolean(bilangan1, bilangan2, hasil){
    if(bilangan1 < bilangan2){
        hasil = true;
    }else if(bilangan2 < bilangan1){
        hasil = false;
    }else{
        hasil = -1
    }
    return hasil
}

console.log(bolean(5,5));