package did

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/multiformats/go-multibase"
	"testing"
	"veric-backend/logic/blockchain/eth"
)

var issuePryKey *eth.PrivateKey

func init() {
	var err error
	issuePryKey, err = eth.NewPrivateKey("d7d35e67ffe13a47033f888d3a9a9b5b734f4da8d5e88bdd45b39b7a2eb53a89")
	if err != nil {
		panic(err)
	}
}

func randomUserKey() *eth.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	return eth.NewPrivateKeyFromECDSA(key)
}

func randomUserDID() (*eth.PrivateKey, *DIDDocument) {
	key := randomUserKey()
	did, err := CreateUserDID(key.Address().String(), issuePryKey)
	if err != nil {
		panic(err)
	}

	return key, did
}

func TestCreateUserDID(t *testing.T) {
	_, oriDID := randomUserDID()

	toByte := oriDID.ToJson()

	println(string(toByte))

	toDID, err := ParseUserDIDFromJsonStr(toByte)
	if err != nil {
		panic(err)
	}

	if toDID.ID != oriDID.ID {
		panic("toDID.ID != oriDID.ID")
	}
}

func TestGetPubKey(t *testing.T) {
	fmt.Printf("ori pubkey: %s\n", issuePryKey.PublicKey().HexString())
	fmt.Printf("ori address: %s\n", issuePryKey.PublicKey().Address().String())
	_, oriDID := randomUserDID()

	toByte := oriDID.ToJson()

	println("============\n")

	toDID, err := ParseUserDIDFromJsonStr(toByte)
	if err != nil {
		panic(err)
	}

	fmt.Printf("import multibaseKey: %s\n", toDID.VerificationMethod[0].MultibaseKey)
	keyType, pubRaw, err := multibase.Decode(toDID.VerificationMethod[0].MultibaseKey)
	if err != nil {
		panic(err)
	}
	pubKey, err := eth.NewPublicKeyFromByte(pubRaw)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decoded keyType: %d\n", keyType)
	fmt.Printf("decoded pubKey: %s\n", pubKey.HexString())
	fmt.Printf("decoded address: %s\n", pubKey.Address().String())
}

func Test222(t *testing.T) {
	key := randomUserKey()
	t.Log(key.HexString())
	signature, err := key.ETHSignature([]byte("sign in"))
	if err != nil {
		panic(err)
	}

	sign, err := eth.GetPublicKeyUseEthSign("sign in", signature)
	if err != nil {
		panic(err)
	}

	if sign.Address().String() != key.Address().String() {
		panic("in")
	}

	t.Log(key.PublicKey().Address().String(), signature)
}

func Test333(t *testing.T) {
	//did, err := ParseUserDIDFromJsonStr([]byte("{\"ID\":\"did:veric:0x08830907f2e2d20a5cb37ee9e0a8bdf8c06e3508\",\"Context\":[\"https://ns.did.ai/suites/secp256k1-2019/v1/\",\"https://w3id.org/did/v1\"],\"Created\":\"2022-11-20T12:53:49+08:00\",\"Updated\":\"2022-11-20T12:53:49+08:00\",\"Version\":1,\"Authentication\":\"did:veric:0x08830907f2e2d20a5cb37ee9e0a8bdf8c06e3508#verification\",\"Address\":\"0x08830907f2e2d20a5cb37ee9e0a8bdf8c06e3508\",\"VerificationMethod\":[{\"ID\":\"did:veric:0x08830907f2e2d20a5cb37ee9e0a8bdf8c06e3508#verification\",\"Controller\":\"did:veric:0x08830907f2e2d20a5cb37ee9e0a8bdf8c06e3508\",\"MethodType\":\"EcdsaSecp256k1VerificationKey2019\",\"MultibaseKey\":\"zQxDgxC3FpGcazoQ2NguXkbm1KQzGBDXuc6bjmd67AUuzZFxcABiReNZS1CZXqkggwN5MuDpvwaPLqek9yA4W6SMG\"}]}"))
	//if err != nil {
	//	panic(err)
	//}

	//key := randomUserKey()
	//did, err := CreateUserDID(key.Address().String(), issuePryKey)
	//if err != nil {
	//	panic(err)
	//}
	//
	//t.Logf("%s", did.ToJson())
	//
	//method, err := did.RetrieveVerificationMethod("did:veric:0x08830907f2e2d20a5cb37ee9e0a8bdf8c06e3508#verification")
	//if err != nil {
	//	panic(err)
	//}

	encoding, data, err := multibase.Decode("zQxDgxC3FpGcazoQ2NguXkbm1KQzGBDXuc6bjmd67AUuzZFxcABiReNZS1CZXqkggwN5MuDpvwaPLqek9yA4W6SMG")
	if err != nil {
		panic(err)
	}

	fromByte, err := eth.NewPublicKeyFromByte(data)
	if err != nil {
		panic(err)
	}

	t.Logf("%v, %v", encoding, fromByte.Address())
}
