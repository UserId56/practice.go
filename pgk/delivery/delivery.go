package delivery

import (
	"fmt"
)

type Delivery interface {
	CalculateCost() float64
}

type CourierDelivery struct {
	Distance float64
}

func (cd CourierDelivery) CalculateCost() float64 {
	return cd.Distance * 2.5

}

type PostDelivery struct {
	Weight float64
}

func (pd PostDelivery) CalculateCost() float64 {
	return (pd.Weight * 0.5) + 10.0
}

type PickupDelivery struct{}

func (pd PickupDelivery) CalculateCost() float64 {
	return 0
}

func ProcessOrder(d Delivery) {
	fmt.Printf("Delivery cost: %.2f\n", d.CalculateCost())
}
