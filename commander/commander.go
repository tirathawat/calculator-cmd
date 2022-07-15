package commander

import (
	"math/big"
	"pair/calculator"
)

type CommandFunc func(a, b *big.Float) *big.Float

type Commander struct {
	calculator calculator.Calculator
}

func (c *Commander) Add() CommandFunc {
	return func(a, b *big.Float) *big.Float {
		return c.calculator.Add(a, b)
	}
}

func (c *Commander) Minus() CommandFunc {
	return func(a, b *big.Float) *big.Float {
		return c.calculator.Minus(a, b)
	}
}

func (c *Commander) Multiply() CommandFunc {
	return func(a, b *big.Float) *big.Float {
		return c.calculator.Multiply(a, b)
	}
}

func (c *Commander) Divide() CommandFunc {
	return func(a, b *big.Float) *big.Float {
		return c.calculator.Divide(a, b)
	}
}

func ProvideCommander(calculator calculator.Calculator) *Commander {
	return &Commander{
		calculator: calculator,
	}
}
