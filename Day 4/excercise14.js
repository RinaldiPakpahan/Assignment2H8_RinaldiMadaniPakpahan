const fetch = require('node-fetch');
function getUserData(username) {
    return new Promise((resolve, reject) => {
        fetch('https://api.github.com/users/' + username)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Gagal mendapatkan data pengguna');
                }
                return response.json();
            })
            .then(data => {
                resolve(data);
            })
            .catch(error => {
                reject(error);
            });
    });
}

getUserData("RinaldiPakpahan")
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error(error.message);
    });