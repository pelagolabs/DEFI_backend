package cron

type walletPoolItem struct {
	Address string
	Chain   string
	IsStat  bool
}

type walletPoolStatItem struct {
	Total float64
	Used  float64
}
