package types

import (
	"database/sql/driver"
	"errors"
	"math/big"
)

type BigInt struct {
	value *big.Int
}

func NewBigIntZero() *BigInt {
	return NewBigIntFast(0)
}

func NewBigIntFast(value int64) *BigInt {
	return &BigInt{value: big.NewInt(value)}
}

func NewBigInt(value *big.Int) *BigInt {
	if value == nil {
		value = big.NewInt(0)
	}
	return &BigInt{value: value}
}

func NewBigIntString(value string) (*BigInt, error) {
	if bi, ok := big.NewInt(0).SetString(value, 10); ok {
		return NewBigInt(bi), nil
	} else {
		return nil, errors.New("parse error")
	}
}
func (*BigInt) GormDataType() string {
	return "string"
}

func (b *BigInt) Value() (driver.Value, error) {
	if b == nil {
		return "0", nil
	}
	return b.value.String(), nil
}

func (b *BigInt) Scan(src any) error {
	b.value = big.NewInt(0)

	switch t := src.(type) {
	default:
		if src != nil {
			return errors.New("types.BigInt only accept string or bytes")
		}
	case []byte:
		if _, ok := b.value.SetString(string(t), 10); !ok {
			return errors.New("parse data to types.BigInt error")
		}
	case string:
		if _, ok := b.value.SetString(t, 10); !ok {
			return errors.New("parse data to types.BigInt error")
		}
	}

	return nil
}

func (b *BigInt) RawBigInt() *big.Int {
	return big.NewInt(0).Set(b.value)
}

func (b *BigInt) Copy() *BigInt {
	return NewBigInt(b.RawBigInt())
}

func (b *BigInt) Add(y *BigInt) *BigInt {
	b.value.Add(b.value, y.value)
	return b
}

func (b *BigInt) Sub(y *BigInt) *BigInt {
	b.value.Sub(b.value, y.value)
	return b
}

func (b *BigInt) Mul(y *BigInt) *BigInt {
	b.value.Mul(b.value, y.value)
	return b
}

func (b *BigInt) Div(y *BigInt) *BigInt {
	b.value.Div(b.value, y.value)
	return b
}

func (b *BigInt) Pow(y *BigInt) *BigInt {
	b.value.Exp(b.value, y.value, nil)
	return b
}

func (b *BigInt) Mod(y *BigInt) *BigInt {
	b.value.Mod(b.value, y.value)
	return b
}

func (b *BigInt) SetZero() *BigInt {
	b.SetInt64(0)
	return b
}

func (b *BigInt) Set(y *BigInt) *BigInt {
	b.value.Set(y.value)
	return b
}

func (b *BigInt) SetInt64(y int64) *BigInt {
	b.value.SetInt64(y)
	return b
}

func (b *BigInt) Cmp(y *BigInt) int {
	return b.value.Cmp(y.value)
}

func (b *BigInt) String() string {
	return b.value.String()
}

func (b *BigInt) Uint64() uint64 {
	return b.value.Uint64()
}

func (b *BigInt) Ceil(digits int) *BigInt {
	precision := NewBigIntFast(10).Pow(NewBigIntFast(int64(digits)))
	mantissa := b.Copy().Mod(precision)

	if !mantissa.IsZero() {
		return b.Sub(mantissa).Add(precision)
	} else {
		return b
	}
}

func (b *BigInt) Floor(digits int) *BigInt {
	precision := NewBigIntFast(10).Pow(NewBigIntFast(int64(digits)))
	mantissa := b.Copy().Mod(precision)

	if !mantissa.IsZero() {
		return b.Sub(mantissa)
	} else {
		return b
	}
}

func (b *BigInt) Round(digits int) *BigInt {
	precision := NewBigIntFast(10).Pow(NewBigIntFast(int64(digits)))
	middleNum := precision.Copy().Div(NewBigIntFast(2))
	mantissa := b.Copy().Mod(precision)

	if mantissa.IsZero() {
		return b
	} else if mantissa.Cmp(middleNum) >= 0 {
		return b.Sub(mantissa).Add(precision)
	} else {
		return b.Sub(mantissa)
	}
}

func (b *BigInt) IsZero() bool {
	return b.value.Cmp(big.NewInt(0)) == 0
}

func (b *BigInt) MarshalJSON() ([]byte, error) {
	return b.value.MarshalJSON()
}

func (b *BigInt) UnmarshalJSON(text []byte) error {
	if b.value == nil {
		b.value = big.NewInt(0)
	}
	return b.value.UnmarshalJSON(text)
}
