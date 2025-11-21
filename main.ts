/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2025-11-18
* @fileoverview This program calculates the receipt information for a table of 5.
*/

// stores 0.13 into HST
const HST: number = 0.13;

// stores 0.15 into GRATUITY
const GRATUITY: number = 0.15;

// INPUT - number data type numberOfPeople
let numberOfPeople: number = 5;

// INPUT - number data type subtotalBillPrice
let subtotalBillPrice: number = 315.99;

// PROCESS
// calculate receipt details
let tax: number = subtotalBillPrice * HST;
let individualTax: number = tax / numberOfPeople;
let tip: number = subtotalBillPrice * GRATUITY;
let individualTip: number = tip / numberOfPeople;
let totalBillPrice: number = subtotalBillPrice + tax + tip;
let individualSubtotal: number = subtotalBillPrice / numberOfPeople;
let individualTotal: number = individualSubtotal + individualTax + individualTip;

// OUTPUT
// display receipt details
console.log("WHOLE BILL");
console.log("Subtotal: " + subtotalBillPrice);
console.log("Tax: " + tax);
console.log("Tip: " + tip);
console.log("Total: " + totalBillPrice);
console.log("SPLIT BILL");
console.log("Subtotal: " + individualSubtotal);
console.log("Tax: " + individualTax);
console.log("Tip: " + individualTip);
console.log("Total: " + individualTotal);

console.log("\nDone.");
