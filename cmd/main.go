package main

import (
	"Task/pgk/Shape"
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
	fmt.Println("------------------------------------------Задание 1 завершено-----------------------------------")

	box := Shape.Rectangle{
		Width:  32.0,
		Height: 12.2,
	}
	circle := Shape.Circle{
		Radius: 12,
	}

	Shape.DescribeShape(box)
	Shape.DescribeShape(circle)

	fmt.Println("------------------------------------------Задание 2 завершено-----------------------------------")
}
