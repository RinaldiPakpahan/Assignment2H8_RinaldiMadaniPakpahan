var Arr = [21,23,24,6,24,2,29];

function tukarBalik(a,b){
    for(i=0; i<Arr.length ; i++){
        for(j=0; j<i; j++){
            if(Arr[i] < Arr[j]){
                var temp = Arr[j];
                Arr[j] = Arr[i];
                Arr[i] = temp
            }
        }
    }
    for (let i = 0; i<Arr.length ; i++){
        console.log(Arr[i]);
    }
}

tukarBalik();