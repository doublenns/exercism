export const reverseString = (str) => {
  // Built-in functions
  // let splitString = str.split("");
  // let reversedArray = splitString.reverse();
  // let joinedArray = reversedArray.join("");
  // return joinedArray;

  // One-liner of above
  // return str.split("").reverse().join("");

  // Multiple assignment using array destructuring
  // let [i, j] = [0, str.length];
  // let splitString = str.split("");
  // while (i < j) {
  //   [splitString[i], splitString[j]] = [splitString[j], splitString[i]];
  //   i++;
  //   j--;
  // }
  // return splitString.join("");

  // Decrementing for loop
  let reversed = ""
  for (let i = str.length - 1; i >= 0; i--) {
    reversed += str[i];
  }
  return reversed;
}
