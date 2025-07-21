package main

import (
	"fmt"
	"sync"
	"time"
)

// Use to Reuse Objects and Slash Allocationssync.Pool
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) // pre-allocate 1024 bytes
	},
}

type PaymentLog struct {
	PaymentID int
	Amount    float64
	Timestamp time.Time
}

func logPayment(payment PaymentLog) string {
	// get a buffer from the pool
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf)

	// reset buffer and write log entry

	buf = buf[:0]
	buf = append(buf, fmt.Sprintf("PaymentID: %d, Amount: $%.2f, Time: %s",
		payment.PaymentID, payment.Amount, payment.Timestamp.Format(time.RFC3339))...)

	time.Sleep(10 * time.Millisecond)
	return string(buf)

}

func main() {
	payments := []PaymentLog{
		{1, 99.99, time.Now()},
		{2, 150.50, time.Now()},
		{3, 200.75, time.Now()},
	}
	start := time.Now()
	for _, payment := range payments {
		result := logPayment(payment)
		fmt.Println(result)
	}
	fmt.Printf("Logged %d payments in %v\n", len(payments), time.Since(start))
}
