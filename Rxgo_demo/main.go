package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

type Circle struct {
	Radius int
	Color  string
}

type Square struct {
	Size  int
	Color string
}

func mapToSquare(ctx context.Context, c interface{}) (interface{}, error) {
	circle, ok := c.(Circle)
	if !ok {
		return nil, errors.New("unexpected input type")
	}
	return Square{Size: circle.Radius, Color: circle.Color}, nil
}

func filterYellowSquares(s interface{}) bool {
	return s.(Square).Color == "yellow"
}

func main() {
	observable := rxgo.Just(
		Circle{Radius: 10, Color: "red"},
		Circle{Radius: 10, Color: "yellow"},
		Circle{Radius: 10, Color: "green"},
		Square{
			Size:  0,
			Color: "",
		},
	)().Map(mapToSquare).Filter(filterYellowSquares)

	for item := range observable.Observe() {
		if item.Error() {
			fmt.Println("an error occurred: ", item.E)
		}
		fmt.Println("observed items are:", item.V)
	}


}