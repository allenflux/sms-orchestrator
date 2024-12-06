package service

import "github.com/tiger1103/gfast-cache/cache"

type (
	ICache interface {
		cache.IGCache
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
