package main

import (
 "fmt"
 "time"
)

type Transaction struct {
 ID     int
 Amount float64
}

func generateTransactions(count int) []Transaction {
 // Pre-allocate slice with exact capacity
 transactions := make([]Transaction, 0, count)
 
 for i := 1; i <= count; i++ {
  transaction := Transaction{
   ID:     i,
   Amount: float64(i) * 10.5,
  }
  transactions = append(transactions, transaction)
  // Simulate processing
  time.Sleep(2 * time.Millisecond)
 }
 return transactions
}

func main() {
 count := 1000
 start := time.Now()
 transactions := generateTransactions(count)
 for _, t := range transactions[:3] { // Print first 3 for brevity
  fmt.Printf("Transaction ID %d: $%.2f\n", t.ID, t.Amount)
 }
 fmt.Printf("Generated %d transactions in %v\n", len(transactions), time.Since(start))
}