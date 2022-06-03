package main

import "fmt"


const pi = 3.1415

func main() {

	printRadiusCircle()
	Rectangle()
		
} 
		func printRadiusCircle()  {

			radius := 4

			circleArea:= float32(radius) * float32(radius) * pi

			fmt.Println("Радиус круга:", circleArea, "\n")
		}


		func Rectangle() {
			x := 10
			y := 5

			SquareRectangle := x * y
			PerimetrRectangle := 2 * (x + y)

			fmt.Println("Площадь прямоугольника: ", SquareRectangle)
			fmt.Println("Периметр прямоугольника: ", PerimetrRectangle)
		}