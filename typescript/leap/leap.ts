export function isLeap(year: number = new Date().getFullYear()): boolean {
  // return year % 400 === 0 || (year % 4 === 0 && year % 100 !== 0);

  // Checks to see whether February has 29 days that year
  return new Date(year, 1, 29).getDate() === 29;
}
