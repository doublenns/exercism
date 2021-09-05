// @ts-check

const HOURS_IN_WORKDAY = 8;
const MONTHLY_BILLABLE_DAYS = 22;


/**
 * The day rate, given a rate per hour
 *
 * @param {number} ratePerHour
 * @returns {number} the rate per day
 */
export function dayRate(ratePerHour) {
  return ratePerHour * HOURS_IN_WORKDAY;
}

/**
 * Calculates the rate per month
 *
 * @param {number} ratePerHour
 * @param {number} discount for example 20% written as 0.2
 * @returns {number} the rounded up monthly rate
 */
export function monthRate(ratePerHour, discount) {
  let rawMonthRate = applyDiscount(dayRate(ratePerHour), discount) * MONTHLY_BILLABLE_DAYS;
  return Math.ceil(rawMonthRate);
}

/**
 * Calculates the number of days in a budget, rounded down
 *
 * @param {number} budget the total budget
 * @param {number} ratePerHour the rate per hour
 * @param {number} discount to apply, example 20% written as 0.2
 * @returns {number} the number of days
 */
export function daysInBudget(budget, ratePerHour, discount) {
  let rawDaysInBudget =  budget / applyDiscount(dayRate(ratePerHour), discount);
  return Math.floor(rawDaysInBudget);
}

/**
 * Applies a discount to the value
 *
 * @param {number} value
 * @param {number} percentage for example 20% written as 0.2
 * @returns {number} the discounted value
 */
function applyDiscount(value, percentage) {
  if (percentage > 0 && percentage <= 0.9) { 
    return value * (1 - percentage);
  } else {
    return value;
  }
}
