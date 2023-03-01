package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"veric-backend/internal/base58"
)

func TronAddress2Hex(tronAddress string) (hexString string, err error) {
	decode, err := base58.Decode(tronAddress, base58.BitcoinAlphabet)
	if err != nil {
		return "", err
	}
	if decode[0] != 0x41 || len(decode) < 21 {
		return "", errors.New("invalid address")
	}

	h256h0 := sha256.New()
	h256h0.Write(decode[:21])
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	if bytes.Compare(h1[:4], decode[len(decode)-4:]) == 0 {
		return hex.EncodeToString(decode[1:21]), nil
	}

	return "", errors.New("invalid address")
}

func Hex2TronAddress(hexString string) (tronAddress string, err error) {
	hexByte, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}

	if len(hexByte) != 20 {
		return "", errors.New("invalid hex")
	}

	hexByte = append([]byte{0x41}, hexByte...)

	h256h0 := sha256.New()
	h256h0.Write(hexByte)
	h0 := h256h0.Sum(nil)

	h256h1 := sha256.New()
	h256h1.Write(h0)
	h1 := h256h1.Sum(nil)

	hexByte = append(hexByte, h1[:4]...)
	return base58.Encode(hexByte, base58.BitcoinAlphabet), nil
}
