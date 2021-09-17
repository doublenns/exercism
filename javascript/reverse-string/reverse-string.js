//
// This is only a SKELETON file for the 'Reverse String' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

import { join } from "path/posix";

export const reverseString = (str) => {
  // Built-in functions
  let splitString = str.split("");
  let reversedArray = splitString.reverse();
  let joinedArray = reversedArray.join("");
  return joinedArray;
};
