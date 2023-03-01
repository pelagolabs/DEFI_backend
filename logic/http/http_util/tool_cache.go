package http_util

import (
	"bytes"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"net/http"
	"time"
	"veric-backend/internal/util"
)

type KeyProvider func(r *HTTPContext) []any

type cacheItem struct {
	recordWriter http.ResponseWriter

	expireTime time.Time
	data       bytes.Buffer
	header     http.Header
	statusCode int
}

func (i *cacheItem) Header() http.Header {
	return i.recordWriter.Header()
}

func (i *cacheItem) Write(bytes []byte) (int, error) {
	if i.header == nil {
		i.header = i.recordWriter.Header().Clone()
	}

	write, err := i.recordWriter.Write(bytes)
	i.data.Write(bytes[:write])
	return write, err
}

func (i *cacheItem) WriteHeader(statusCode int) {
	i.statusCode = statusCode
	if i.header == nil {
		i.header = i.recordWriter.Header().Clone()
	}

	i.recordWriter.WriteHeader(statusCode)
}

func (i *cacheItem) Replay(w http.ResponseWriter) {
	for k, vs := range i.header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	if i.statusCode != 0 {
		w.WriteHeader(i.statusCode)
	}

	_, _ = w.Write(i.data.Bytes())
}

func (i *cacheItem) Record(w http.ResponseWriter) http.ResponseWriter {
	i.recordWriter = w
	defer func() {
		i.recordWriter = nil
	}()

	return i
}

func (i *cacheItem) IsExpired() bool {
	return time.Now().After(i.expireTime)
}

type LruCacheTemplate struct {
	lru         *lru.Cache
	cacheTime   time.Duration
	keyProvider KeyProvider
}

func NewLruCacheTemplate(cacheCount, cacheTime time.Duration, keyProvider KeyProvider) *LruCacheTemplate {
	cache, _ := lru.New(int(cacheCount))
	return &LruCacheTemplate{lru: cache, cacheTime: cacheTime, keyProvider: keyProvider}
}

func (l *LruCacheTemplate) Cache(f RespFunc) RespFunc {
	randomId := util.RandString(16)
	return func(w http.ResponseWriter, r *http.Request) {
		keyProvider := l.keyProvider(NewHTTPContext(r, w))
		if len(keyProvider) == 0 {
			f(w, r)
		} else {
			key := fmt.Sprintf("%v-%s", keyProvider, randomId)
			if value, ok := l.lru.Get(key); ok {
				item := value.(*cacheItem)
				if !item.IsExpired() {
					value.(*cacheItem).Replay(w)
					return
				}
			}

			cache := &cacheItem{expireTime: time.Now().Add(l.cacheTime)}
			f(cache.Record(w), r)
			l.lru.Add(key, cache)
		}
	}
}
