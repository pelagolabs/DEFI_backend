package multicall

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MultiCallMethod struct {
	ContractAddress  common.Address
	ContractMetadata *bind.MetaData
	MethodName       string
	Params           []any
	ResultProcess    func(result []any) error
}

func (m *MultiCallMethod) Data() (data []byte, err error) {
	abi, err := m.ContractMetadata.GetAbi()
	if err != nil {
		return nil, err
	}

	return abi.Pack(m.MethodName, m.Params...)
}

func (m *MultiCallMethod) returnData(data []byte) (err error) {
	if m.ResultProcess == nil {
		return nil
	}

	abi, err := m.ContractMetadata.GetAbi()
	if err != nil {
		return err
	}

	unpack, err := abi.Unpack(m.MethodName, data)
	if err != nil {
		return err
	}

	return m.ResultProcess(unpack)
}

type MultiCallFactory struct {
	address common.Address
	backend bind.ContractBackend

	pushMethodsData []MulticallCall
	pushMethods     []MultiCallMethod
}

func NewMultiCallFactory(address common.Address, backend bind.ContractBackend) *MultiCallFactory {
	return &MultiCallFactory{
		address:         address,
		backend:         backend,
		pushMethods:     make([]MultiCallMethod, 0, 10),
		pushMethodsData: make([]MulticallCall, 0, 10),
	}
}

func (f *MultiCallFactory) PushCall(call MultiCallMethod) error {
	data, err := call.Data()
	if err != nil {
		return err
	}

	f.pushMethods = append(f.pushMethods, call)
	f.pushMethodsData = append(f.pushMethodsData, MulticallCall{
		Target:   call.ContractAddress,
		CallData: data,
	})
	return nil
}

func (f *MultiCallFactory) PushCalls(calls ...MultiCallMethod) (err error) {
	for _, call := range calls {
		err = f.PushCall(call)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *MultiCallFactory) DoRead(opts *bind.CallOpts) error {
	multicall, err := NewMulticall(f.address, f.backend)
	if err != nil {
		return err
	}

	aggregate, err := multicall.Aggregate(opts, f.pushMethodsData)
	if err != nil {
		return err
	}

	for i, datum := range aggregate.ReturnData {
		err = f.pushMethods[i].returnData(datum)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *MultiCallFactory) DoWrite(opts *bind.TransactOpts) (*types.Transaction, error) {
	multicall, err := NewMulticallWriter(f.address, f.backend)
	if err != nil {
		return nil, err
	}

	methodData := make([]Multicall2Call, len(f.pushMethodsData))
	for i, data := range f.pushMethodsData {
		methodData[i] = Multicall2Call{
			Target:   data.Target,
			CallData: data.CallData,
		}
	}

	return multicall.Aggregate(opts, methodData)
}
