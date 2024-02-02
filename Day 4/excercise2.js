let bioRinaldi = {
    nama_depan: "Rinaldi",
    nama_belakang: "Pakpahan",
    hobi: ["Olahraga"],
    angka_favorit: 23,
    memakai_kacamata: true,
};
console.log(bioRinaldi.nama_depan);
console.log(bioRinaldi.nama_belakang);
console.log(bioRinaldi.hobi);
console.log(bioRinaldi.angka_favorit);
console.log(bioRinaldi.memakai_kacamata);

console.log(bioRinaldi.nama_depan, bioRinaldi.nama_belakang);

bioRinaldi.angka_favorit = 9
console.log(bioRinaldi.angka_favorit)

bioRinaldi.hobi.push("ngoding")
console.log(bioRinaldi.hobi);

bioRinaldi.lulusan = "Hacktiv8"
console.log(bioRinaldi.lulusan)

for(var i=0; i<bioRinaldi.hobi.length; i++){
    console.log(bioRinaldi.hobi[i]);
}
