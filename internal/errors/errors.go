package errors

import "errors"

var (
	ErrInputNoASCII                 = errors.New("input no ASCII")
	ErrExpressionNotValid           = errors.New("expression is not valid")
	ErrDividedByZero                = errors.New("divided by zero")
	ErrUnknownOperator              = errors.New("unknown operator")
	ErrPopValueInvalidExpression    = errors.New("popValue: invalid expression")
	ErrPopOperatorInvalidExpression = errors.New("popOperator: invalid expression")
	ErrInvalidExpression            = errors.New("parse: invalid expression")
	ErrInvalidValue                 = errors.New("parse: invalid value")
	ErrInvalidExpressionValues      = errors.New("invalid expression's values")
)
