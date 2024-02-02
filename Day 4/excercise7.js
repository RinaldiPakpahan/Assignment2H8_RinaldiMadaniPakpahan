function urutHuruf(text){
    return text.split("").sort().join("");
}

console.log(urutHuruf("halo"))
console.log(urutHuruf("qwerty"))
console.log(urutHuruf("qwertyuiopasdfghjklzxcvbnm"))