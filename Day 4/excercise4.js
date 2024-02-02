function ganjilGenap(bilangan, hasil){
    if (typeof bilangan !== 'number' || isNaN(bilangan)) {
        hasil = "Invalid Data"
    }else {
        if(bilangan % 2 == 0){
        hasil = "Genap"
        }else{
            hasil = "Ganjil"
        }
    }
    return hasil
}

console.log(ganjilGenap(10));