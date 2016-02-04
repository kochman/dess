package dess

import (
	"math/big"
)

type Decimal struct {
	coefficient *big.Int
	exponent    int64
}

// pow returns an int64 of base raised to power. The math package
// returns floats, which could lead to accuracy issues here.
func pow(base, power int64) int64 {
	if power == 0 {
		return 1
	} else if power == 1 {
		return base
	} else {
		return base * pow(base, power-1)
	}
}

// NewDecimal allocates and returns a new Decimal set to x, with
// places indicating the number of places after the decimal point.
func NewDecimal(x int64, places int64) *Decimal {
	decimal := new(Decimal)

	if x == 0 {
		decimal.coefficient = big.NewInt(0)
		decimal.exponent = 1
	} else {
		decimal.coefficient = big.NewInt(x)
		decimal.exponent = places
	}
	return decimal
}

func (d1 *Decimal) Add(d2 *Decimal) *Decimal {
	decimal := new(Decimal)
	if d1.exponent > d2.exponent {
		diff := d1.exponent - d2.exponent
		alignedExponent := pow(10, diff)
		alignedCoefficient := &big.Int{}
		alignedCoefficient.Mul(d2.coefficient, big.NewInt(alignedExponent))

		coef := &big.Int{}
		decimal.coefficient = coef.Add(alignedCoefficient, d1.coefficient)
		decimal.exponent = d2.exponent
	} else if d1.exponent < d2.exponent {
		diff := d2.exponent - d1.exponent
		alignedExponent := pow(10, diff)
		alignedCoefficient := &big.Int{}
		alignedCoefficient.Mul(d1.coefficient, big.NewInt(alignedExponent))

		coef := &big.Int{}
		decimal.coefficient = coef.Add(alignedCoefficient, d2.coefficient)
		decimal.exponent = d1.exponent
	} else {
		coef := &big.Int{}
		decimal.coefficient = coef.Add(d1.coefficient, d2.coefficient)
		decimal.exponent = d1.exponent
	}
	return decimal
}

func (d1 *Decimal) Equals(d2 *Decimal) bool {
	if d1.coefficient.Cmp(d2.coefficient) != 0 {
		return false
	}
	if d1.exponent != d2.exponent {
		return false
	}
	return true
}

func (d1 *Decimal) Sub(d2 *Decimal) *Decimal {
	inverted := &Decimal{&big.Int{}, d2.exponent}
	inverted.coefficient.Neg(d2.coefficient)
	return d1.Add(inverted)
}

func (d *Decimal) String() string {
	coef := d.coefficient.String()
	if coef == "0" {
		return "0.0"
	}

	neg := false
	if coef[0] == '-' {
		neg = true
		coef = coef[1:]
	}

	var i int64
	for i = 0; i < d.exponent; i += 1 {
		if i >= int64(len(coef)) {
			coef = "0" + coef
		}
	}

	var before string
	var after string
	if coef[0] == '0' {
		before = "0"
		after = coef[:i]
	} else {
		before = coef[:i]
		after = coef[i:]
	}

	if neg {
		before = "-" + before
	}
	return before + "." + after
}
