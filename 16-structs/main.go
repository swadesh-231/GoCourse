package main

import (
	"fmt"
	"time"
)

type Order struct{
	ID int
	amount float64
	status string
	created_at time.Time // nanorsecionds presigions
	updated_at time.Time

}

func printOrder(order Order) {
	fmt.Println("Order Details")
	fmt.Println("-------------------------")
	fmt.Printf("ID:         %d\n", order.ID)
	fmt.Printf("Amount:     %.2f\n", order.amount)
	fmt.Printf("Status:     %s\n", order.status)
	fmt.Printf("Created At: %s\n", order.created_at.Format(time.RFC3339Nano))
	fmt.Printf("Updated At: %s\n", order.updated_at.Format(time.RFC3339Nano))
	fmt.Println()
}

func main(){
	now := time.Now()

	order1 := Order{
		ID:         1,
		amount:     100.0,
		status:     "pending",
		created_at: now,
		updated_at: now,
	}

	order2 := Order{
		ID:         2,
		amount:     250.50,
		status:     "completed",
		created_at: now,
		updated_at: now,
	}

	printOrder(order1)
	printOrder(order2)
}