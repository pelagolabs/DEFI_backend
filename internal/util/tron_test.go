package util

import "testing"

func TestTronAddress2Hex(t *testing.T) {
	check := map[string]string{
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t": "a614f803b6fd780986a42c78ec9c7f77e6ded13c",
		"TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8": "3487b63d30b5b2c87fb7ffa8bcfade38eaac1abe",
		"TXpw8XeWYeTUd4quDskoUqeQPowRh4jY65": "efc230e125c24de35f6290afcafa28d50b436536",
		"THb4CqiFdwNHsWsQCs4JhzwjMWys4aqCbF": "53908308f4aa220fb10d778b5d1b34489cd6edfc",
	}

	for k, v := range check {
		if val, _ := TronAddress2Hex(k); val != v {
			t.Fatalf("TronAddress2Hex %s error: %s", k, val)
		}
	}

	for k, v := range check {
		if val, _ := Hex2TronAddress(v); val != k {
			t.Fatalf("Hex2TronAddress %s error: %s", k, val)
		}
	}
}
