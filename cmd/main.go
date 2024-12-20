package main

import "github.com/ivanov-nikolay/calculator/internal/app"

func main() {
	calculator := app.NewApplication()
	calculator.RunApplication()
}
