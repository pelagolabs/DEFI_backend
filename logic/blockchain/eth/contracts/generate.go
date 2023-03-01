package contracts

// download abigen from https://geth.ethereum.org/downloads/
//go:generate abigen --abi=erc20/sol/erc20.abi --pkg=erc20 --out=erc20/ERC20.go
//go:generate abigen --abi=payment_vault/sol/payment_vault.abi --pkg=payment_vault --out=payment_vault/payment_vault.go
//go:generate abigen --abi=address_pool/sol/address_pool.abi --pkg=address_pool --out=address_pool/address_pool.go
//go:generate abigen --abi=multicall/sol/multicall.abi --pkg=multicall --out=multicall/multicall.go
//go:generate abigen --abi=multicall/sol/multicall_write.abi --pkg=multicall --type=MulticallWriter --out=multicall/multicall_write.go
//go:generate abigen --abi=did_registry/sol/did_registry.abi --pkg=did_registry --out=did_registry/did_registry.go
//go:generate abigen --abi=lp_manager/sol/lp_manager.abi --pkg=lp_manager --out=lp_manager/lp_manager.go
//go:generate abigen --abi=token_stake/sol/token_stake.abi --pkg=token_stake --out=token_stake/token_stake.go
