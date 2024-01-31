function getDataProfil(){
  let nama = document.getElementById("nama1").value;
  let role = document.getElementById("role1").value;
  let availability = document.getElementById("availability1").value;
  let age = document.getElementById("age1").value;
  let lokasi = document.getElementById("Lokasi1").value;
  let experience = document.getElementById("YearsExperience1").value;
  let email = document.getElementById("Email1").value;
 
  document.getElementById("name").innerHTML =nama;
  document.getElementById("rols").innerHTML= role;
  document.getElementById("avai").innerHTML = availability;
  document.getElementById("usai").innerHTML=age;
  document.getElementById("lok").innerHTML=lokasi;
  document.getElementById("expe").innerHTML=experience;
  document.getElementById("mail").innerHTML=email;
  console.log(nama, role, availability, age, lokasi, experience, email);
};
const form = document.getElementById("profileForm");
form.addEventListener("submit", (event) => {
    $("#form .close").click();
    event.preventDefault();
    getDataProfil()
});