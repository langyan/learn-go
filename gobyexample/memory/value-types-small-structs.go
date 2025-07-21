package main

import (
	"fmt"
	"time"
)

// Using value types for small structs avoids heap allocations, as they’re stored on the stack instead.
type Payment struct {
	ID     int
	Amount float64
}

func processPayment(payment Payment) Payment {
	// Process payment (pass by value)
	time.Sleep(5 * time.Microsecond)

	payment.Amount *= 1.05
	fmt.Printf("Processed Payment ID  %d : $%.2f \n", payment.ID, payment.Amount)
	return payment
}

func processPayments(payments []Payment) []Payment {
	result := make([]Payment, len(payments))
	for i, payment := range payments {
		result[i] = processPayment(payment)
	}
	return result
}

func main() {
	payments := []Payment{
		{1, 100.0},
		{2, 200.0},
		{3, 300.0},
	}
	start := time.Now()
	processed := processPayments(payments)
	for _, p := range processed {
		fmt.Printf("Final Payment ID %d: $%.2f\n", p.ID, p.Amount)
	}
	fmt.Printf("Processed %d payments in %v\n", len(payments), time.Since(start))
}
