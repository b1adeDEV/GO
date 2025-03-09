package main

import "fmt"

func main() {
	flowers := []string{"Rose", "Tulip", "Lily", "Blade"}
	msg := flower(flowers)
	fmt.Println(msg)
}

func flower(flowers []string) string {
	for _, flav := range flowers {
		switch flav {
		case "Rose":
			fmt.Println("Это розы", flav)
		case "Tulip":
			fmt.Println("Это тюльпаны", flav)
		case "Lily":
			fmt.Println("Это лилии", flav)
		default:
			fmt.Println("Это не понятные цветы", flav)
		}
	}
	return "Перебор цветов закончен"
}
