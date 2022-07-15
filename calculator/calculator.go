package calculator

import "math/big"

type Calculator struct {
}

func (c Calculator) Add(a, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}

func (c Calculator) Minus(a, b *big.Float) *big.Float {
	return new(big.Float).Sub(a, b)
}

func (c Calculator) Multiply(a, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

func (c Calculator) Divide(a, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}
