package main

import (
	"fmt"
	"strconv"
	"github.com/zqdlove/cachego"
)

func main() {
	cache := cachego.Cache("myCache")

	cache.SetDataLoader(func(key interface{}, args ...interface{}) *cachego.CacheItem {
		val := "This is test with key" + key.(string)
		item := cache3go.NewCacheItem(key , 0 , val)
		return item
	})
	for i := 0; i < 10; i++ {
		res, err := cache.Value("someKey_" + strconv.Itoa(i))
		if err == nil {
			fmt.Println("Found value in cache:", res.Data())
		}else {
			fmt.Println("Error retrieving value from cache:", err)
		}
	}
}
