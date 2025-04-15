package main

import (
	"Task/pgk/delivery"
	"fmt"
)

func main() {
	fmt.Println("Start")
	courier := delivery.CourierDelivery{
		Distance: 10.1,
	}
	post := delivery.PostDelivery{
		Weight: 2.0,
	}
	pickup := delivery.PickupDelivery{}
	delivery.ProcessOrder(courier)
	delivery.ProcessOrder(post)
	delivery.ProcessOrder(pickup)
}
