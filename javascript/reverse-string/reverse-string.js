export const reverseString = (str) => {
  // Built-in functions
  let splitString = str.split("");
  let reversedArray = splitString.reverse();
  let joinedArray = reversedArray.join("");
  return joinedArray;

  // Multiple assignment using array destructuring
  // let [i, j] = [0, str.length];
  // let splitString = str.split("");
  // while (i < j) {
  //   [splitString[i], splitString[j]] = [splitString[j], splitString[i]];
  //   i++;
  //   j--;
  // }
  // return splitString.join("");
}
