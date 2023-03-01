package types

import (
	"database/sql/driver"
	"errors"
	"math/big"
)

const defaultPrecision = 128

type BigFloat struct {
	value *big.Float
}

func NewBigFloatZero() *BigFloat {
	return NewBigFloatFast(0)
}

func NewBigFloatFast(value float64) *BigFloat {
	return &BigFloat{value: big.NewFloat(value).SetPrec(defaultPrecision)}
}

func NewBigFloatUseBigInt(value *BigInt) *BigFloat {
	return &BigFloat{value: big.NewFloat(0).SetPrec(defaultPrecision).SetInt(value.RawBigInt())}
}

func NewBigFloat(value *big.Float) *BigFloat {
	return &BigFloat{value: value.SetPrec(defaultPrecision)}
}

func (*BigFloat) GormDataType() string {
	return "string"
}

func (b *BigFloat) Value() (driver.Value, error) {
	if b == nil {
		return "0.0", nil
	}

	return b.value.String(), nil
}

func (b *BigFloat) Scan(src any) error {
	b.value = big.NewFloat(0)

	switch t := src.(type) {
	default:
		if src != nil {
			return errors.New("types.BigFloat only accept string or bytes")
		}
	case []byte:
		if _, ok := b.value.SetString(string(t)); !ok {
			return errors.New("parse data to types.BigFloat error")
		}
	case string:
		if _, ok := b.value.SetString(t); !ok {
			return errors.New("parse data to types.BigFloat error")
		}
	}

	return nil
}

func (b *BigFloat) RawBigFloat() *big.Float {
	return big.NewFloat(0).Set(b.value)
}

func (b *BigFloat) Copy() *BigFloat {
	return NewBigFloat(b.RawBigFloat())
}

func (b *BigFloat) Add(y *BigFloat) *BigFloat {
	b.value.Add(b.value, y.value)
	return b
}

func (b *BigFloat) Sub(y *BigFloat) *BigFloat {
	b.value.Sub(b.value, y.value)
	return b
}

func (b *BigFloat) Mul(y *BigFloat) *BigFloat {
	b.value.Mul(b.value, y.value)
	return b
}

func (b *BigFloat) Div(y *BigFloat) *BigFloat {
	b.value.Quo(b.value, y.value)
	return b
}

func (b *BigFloat) ToInt() *BigInt {
	intVal, _ := b.value.Int(big.NewInt(0))
	return NewBigInt(intVal)
}

func (b *BigFloat) RoundToInt() *BigInt {
	ob := b.Copy()
	delta := 0.5
	if ob.value.Sign() < 0 {
		delta = -0.5
	}

	ob.value.Add(ob.value, new(big.Float).SetFloat64(delta))
	intVal, _ := ob.value.Int(big.NewInt(0))
	return NewBigInt(intVal)
}

func (b *BigFloat) Cmp(y *BigFloat) int {
	return b.value.Cmp(y.value)
}

func (b *BigFloat) String() string {
	return b.value.String()
}

func (b *BigFloat) MarshalJSON() ([]byte, error) {
	return b.value.MarshalText()
}

func (b *BigFloat) UnmarshalJSON(text []byte) error {
	if b.value == nil {
		b.value = big.NewFloat(0)
	}
	return b.value.UnmarshalText(text)
}
