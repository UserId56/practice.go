package delivery

import (
	"fmt"
)

//### Задача 1: Система обработки заказов
//**ТЗ: Система обработки заказов в интернет-магазине**
//
//**Описание**:
//Вам нужно разработать систему для обработки заказов в интернет-магазине. Каждый заказ может быть доставлен разными способами (курьером, почтой, самовывозом), и для каждого способа доставки нужно рассчитать стоимость.
//
//**Требования**:
//1. Создайте интерфейс Delivery с методом CalculateCost() float64, который возвращает стоимость доставки.
//2. Реализуйте три структуры, представляющие способы доставки:
//- CourierDelivery (доставка курьером, стоимость зависит от расстояния: 2.5 за километр).
//- PostDelivery (почтовая доставка, фиксированная стоимость 10.0 плюс 0.5 за каждый килограмм веса посылки).
//- PickupDelivery (самовывоз, стоимость всегда 0).
//3. Каждая структура должна содержать необходимые поля (например, расстояние для CourierDelivery, вес для PostDelivery).
//4. Создайте функцию ProcessOrder(d Delivery), которая принимает способ доставки и выводит сообщение вида: Delivery cost: <стоимость>.
//5. В функции main создайте по одному экземпляру каждого способа доставки и вызовите ProcessOrder для каждого.
//
//**Ограничения**:
//- Используйте только интерфейсы и структуры.
//- Стоимость должна быть неотрицательной.
//- Параметры доставки (расстояние, вес) задаются при создании структур.
//
//**Пример вывода**:
//Delivery cost: 25.0
//Delivery cost: 12.5
//Delivery cost: 0.0

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
