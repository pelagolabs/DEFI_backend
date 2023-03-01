package harmony

import "github.com/ethereum/go-ethereum/rlp"

type DidInfo struct {
	Did    string `json:"did"`
	PubKey string `json:"pub_key"`
}

func (e *DidInfo) MarshalBinary() (data []byte, err error) {
	return rlp.EncodeToBytes(e)
}

func (e *DidInfo) UnmarshalBinary(data []byte) error {
	return rlp.DecodeBytes(data, &e)
}
