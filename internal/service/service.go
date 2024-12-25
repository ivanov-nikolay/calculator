package service

import "github.com/ivanov-nikolay/calculator/pkg/calculator"

// Calculation вызывает функцию вычисления арифметического выражения в слой бизнес-логики
func Calculation(expression string) (float64, error) {

	return calculator.Calculator(expression)
}
