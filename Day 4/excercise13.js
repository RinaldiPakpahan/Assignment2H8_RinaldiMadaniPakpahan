const fetch = require('node-fetch');
function getUserData(username, fn){
    fetch('https://api.github.com/users/' + username)
      .then(response => {
        //handle response            
        console.log(response);
        return response.json()
      })
      .then(data => {
        fn(data)
        
      })
      .catch(error => {
        
      });    
}

getUserData("RinaldiPakpahan", console.log)