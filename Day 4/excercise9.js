function threeStepAB(num) {
    for (let i = 0; i < num.length; i++) {
      if (num[i] === "a" && num[i + 4] === "b") {
        return true;
      } else if (num[i] === "b" && num[i + 4] === "a") {
        return true;
      }
    }    
    return false;
  }
  
  // TEST CASES
console.log(threeStepAB('lane borrowed')); // true
console.log(threeStepAB('i am sick')); // false
console.log(threeStepAB('you are boring')); // true
console.log(threeStepAB('barbarian')); // true
console.log(threeStepAB('bacon and meat')); // false