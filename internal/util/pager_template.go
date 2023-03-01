package util

const WalletPoolLackMessage = `
The available receiver wallet on %s not enough(less than 20%%) now. please add more.

Chain name: %s
Total wallet count: %.0f
Used wallet count: %.0f (%.2f%%)
`

const PoolRewardWithdrawFailedMessage = `
Pool reward automatic withdraw failed.

Chain name: %s
Is chain native token: %t
Currency address: %s
Withdraw amount (wei): %d
Tx hash: %s
`

const PoolRewardAddFailedMessage = `
Pool reward automatic add failed.

Chain name: %s
Is chain native token: %t
Currency address: %s
Withdraw amount (wei): %d
Tx hash: %s
`
