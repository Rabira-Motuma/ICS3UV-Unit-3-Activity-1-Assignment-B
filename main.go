/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date 2025-11-18
 * This program calculates the receipt information for a table of 5.
 */

package main

import "fmt"

func main() {
	// stores 0.13 into HST
	const HST float32 = 0.13

	// stores 0.15 into GRATUITY
	const GRATUITY float32 = 0.15

	// INPUT - float32 data type numberOfPeople
	var numberOfPeople float32 = 5.0

	// Input - float32 data type subtotalBillPrice
	var subtotalBillPrice float32 = 315.99

	// PROCESS
	// calculate receipt details
	var tax float32 = subtotalBillPrice * HST
	var individualTax float32 = tax / numberOfPeople
	var tip float32 = subtotalBillPrice * GRATUITY
	var individualTip float32 = tip / numberOfPeople
	var totalBillPrice float32 = subtotalBillPrice + tax + tip
	var individualSubtotal float32 = subtotalBillPrice / numberOfPeople
	var individualTotal float32 = individualSubtotal + individualTax + individualTip

	// OUTPUT
	// display receipt details
	fmt.Println("WHOLE BILL")
	fmt.Println("Subtotal:", subtotalBillPrice)
	fmt.Println("Tax:", tax)
	fmt.Println("Tip:", tip)
	fmt.Println("Total:", totalBillPrice)
	fmt.Println("SPLIT BILL")
	fmt.Println("Subtotal:", individualSubtotal)
	fmt.Println("Tax:", individualTax)
	fmt.Println("Tip:", individualTip)
	fmt.Println("Total:", individualTotal)

	fmt.Println("\nDone.")
}
