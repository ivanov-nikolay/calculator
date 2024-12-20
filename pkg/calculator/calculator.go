package calculator

import (
	"strconv"

	"github.com/ivanov-nikolay/calculator/internal/errors"
)

// Calculator вычисляет значение арифметического выражения
func Calculator(expression string) (float64, error) {
	var (
		cntByte int
		x       string
		number  string
	)
	resultString := make([]string, 0)
	cntByte = len([]rune(expression))

	if len(expression) != cntByte {
		return 0, errors.ErrInputNoASCII
	}

	for i := 0; i < len(expression); i++ {
		x = string(expression[i])
		_, err := strconv.Atoi(x)
		if err == nil || x == "." {
			number += x
			continue
		} else {
			if number != "" {
				resultString = append(resultString, number)
				number = ""
			}
			if x == "/" || x == "*" || x == "+" || x == "-" || x == "(" || x == ")" {
				resultString = append(resultString, x)
			} else if x != " " {
				return 0, errors.ErrExpressionNotValid
			}
		}
	}

	if number != "" {
		resultString = append(resultString, number)
	}

	result, err := parse(resultString)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func priority(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func calculator(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.ErrDividedByZero
		} else {
			return a / b, nil
		}
	default:
		return 0, errors.ErrUnknownOperator
	}
}

func pushValue(value float64, values []float64) []float64 {
	return append(values, value)
}

func popValue(values []float64) (float64, []float64, error) {
	if len(values) == 0 {
		return 0, nil, errors.ErrPopValueInvalidExpression
	}
	return values[len(values)-1], values[:len(values)-1], nil
}

func pushOperator(operator string, operators []string) []string {
	return append(operators, operator)
}

func popOperator(operators []string) (string, []string, error) {
	if len(operators) == 0 {
		return "", nil, errors.ErrPopOperatorInvalidExpression
	}
	return operators[len(operators)-1], operators[:len(operators)-1], nil
}

func estimation(values []float64, operators []string) ([]float64, []string, error) {
	operator, operators, err := popOperator(operators)
	if err != nil {
		return nil, nil, err
	}
	b, values, err := popValue(values)
	if err != nil {
		return nil, nil, err
	}
	a, values, err := popValue(values)
	if err != nil {
		return nil, nil, err
	}
	result, err := calculator(a, b, operator)
	if err != nil {
		return nil, nil, err
	}
	values = pushValue(result, values)
	return values, operators, nil
}

func parse(expression []string) (float64, error) {
	values := make([]float64, 0)
	operators := make([]string, 0)

	for _, value := range expression {
		if number, err := strconv.ParseFloat(value, 64); err == nil {
			values = pushValue(number, values)
		} else if value == "(" {
			operators = pushOperator(value, operators)
		} else if value == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				var err error
				values, operators, err = estimation(values, operators)
				if err != nil {
					return 0, err
				}
			}
			if len(operators) == 0 || operators[len(operators)-1] != "(" {
				return 0, errors.ErrInvalidExpression
			}
			operators = operators[:len(operators)-1]
		} else if value == "+" || value == "-" || value == "*" || value == "/" {
			for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(value) {
				var err error
				values, operators, err = estimation(values, operators)
				if err != nil {
					return 0, err
				}
			}
			operators = pushOperator(value, operators)
		} else {
			return 0, errors.ErrInvalidValue
		}
	}

	for len(operators) > 0 {
		var err error
		values, operators, err = estimation(values, operators)
		if err != nil {
			return 0, err
		}
	}

	if len(values) != 1 {
		return 0, errors.ErrInvalidExpressionValues
	}

	return values[0], nil
}
