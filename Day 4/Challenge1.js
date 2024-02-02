var Arr = [1,3,3,3,2,2,2,4];
var endTemp = [];
var countLen =0;
var countMax = 0;

function cekArr(a,b){
    for(i=0; i<Arr.length; i++){
        endTemp.push(0);
    }
    for(i=0; i < Arr.length; i++){
        for(j=0; j<i; j++){
            if(Arr[j] === Arr[i]){
                endTemp[i]++;
            }
        }
    }
    for(i=0; i<Arr.length; i++){
        if(endTemp[i]>countMax){
            countLen = i;
            countMax = endTemp[i]
        }
    }
    console.log(Arr[countMax], countLen)
}

cekArr();
