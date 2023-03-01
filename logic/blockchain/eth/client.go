package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type Client struct {
	client *ethclient.Client
}

func NewClient(rawurl string) (*Client, error) {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.client.SuggestGasPrice(ctx)
}

func (c *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return c.client.EstimateGas(ctx, msg)
}

func (c *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.client.HeaderByNumber(ctx, number)
}

func (c *Client) BlockNumber(ctx context.Context) (uint64, error) {
	return c.client.BlockNumber(ctx)
}

func (c *Client) ChainID(ctx context.Context) (*big.Int, error) {
	return c.client.ChainID(ctx)
}

func (c *Client) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return c.client.FilterLogs(ctx, query)
}

func (c *Client) Nonce(ctx context.Context, fromAddress common.Address) (uint64, error) {
	return c.client.PendingNonceAt(context.Background(), fromAddress)
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.client.NetworkID(context.Background())
}

func (c *Client) TransactionReceipt(txHash common.Hash) (*types.Receipt, error) {
	return c.client.TransactionReceipt(context.Background(), txHash)
}

func (c *Client) TransactionByHash(txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return c.client.TransactionByHash(context.Background(), txHash)
}

func (c *Client) BlockByHash(hash common.Hash) (*types.Block, error) {
	return c.client.BlockByHash(context.Background(), hash)
}

func (c *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return c.client.SubscribeFilterLogs(ctx, q, ch)
}

func (c *Client) TransferringETH(ctx context.Context, pk *PrivateKey, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	nonce, err := c.Nonce(ctx, pk.Address())
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	chainID, err := c.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      uint64(21000),
		GasPrice: gasPrice,
		Data:     data,
	})

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk.key)
	if err != nil {
		return nil, err
	}

	err = c.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (c *Client) SendRAWTransaction(ctx context.Context, pk *PrivateKey, rawByte []byte) (*types.Transaction, error) {
	tx := &types.Transaction{}

	err := rlp.DecodeBytes(rawByte, &tx)
	if err != nil {
		return nil, err
	}

	err = c.client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) Client() *ethclient.Client {
	return c.client
}
