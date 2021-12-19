package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func main() {

	// Create the input channel
	ch := make(chan rxgo.Item)
	// Data producer
	go producer(ch)

	observable := rxgo.FromChannel(ch).Filter(func(item interface{}) bool {
		// Filter operation
		customer := item.(Customer)
		return customer.Age >= 18
	}).
		Map(func(c context.Context, item interface{}) (interface{}, error) {
			// Enrich operation
			customer := item.(Customer)
			taxNumber, err := setTaxNumber(customer)
			if err != nil {
				return nil, err
			}
			customer.TaxNumber = taxNumber
			return customer, nil
		},
			// Create multiple instances of the map operator
			rxgo.WithCPUPool(),
			// Serialize the items emitted by their Customer.ID
			/*
				rxgo.Serialize(func(item interface{}) int {
					customer := item.(Customer)
					return customer.ID
				}),

			*/

			rxgo.WithBufferedChannel(3))

	for customer := range observable.Observe() {
		if customer.Error() {
			fmt.Println("error", customer.Error())
		}
		fmt.Println("observador ", customer.V)
	}

}

func setTaxNumber(customer Customer) (string, error) {
	if customer.TaxNumber == "" {
		return "", errors.New("cannot get tax number")
	}
	return "Tax:" + customer.TaxNumber, nil
}

func producer(ch chan rxgo.Item) {

	ch <- rxgo.Of(Customer{
		ID:        10,
		Name:      "miguel",
		LastName:  "",
		Age:       20,
		TaxNumber: "23459443",
	})
	ch <- rxgo.Of(Customer{
		ID:        20,
		Name:      "daniel",
		LastName:  "",
		Age:       10,
		TaxNumber: "234594234",
	})
	ch <- rxgo.Of(Customer{
		ID:        30,
		Name:      "Laura",
		LastName:  "",
		Age:       21,
		TaxNumber: "234594532",
	})
	ch <- rxgo.Of(Customer{
		ID:        32,
		Name:      "martin",
		LastName:  "",
		Age:       21,
		TaxNumber: "23459432432",
	})
	ch <- rxgo.Of(Customer{
		ID:        50,
		Name:      "Maria",
		LastName:  "",
		Age:       21,
		TaxNumber: "12594532",
	})
	ch <- rxgo.Of(2)
	close(ch)
}
