package main

import (
	"fmt"
	"time"
	"github.com/zqdlove/cachego"
)

func main() {
	cache := cache3go.Cache("myCache")

	cache.SetAddedItemCallback(func(entry *cachego.CacheItem){
		fmt.Println("Added:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.SetAboutToDeleteItemCallback(func(entry * cachego.CacheItem) {
		fmt.Println("Deleting:", entry.Key(), entry.Data(), entry.CreatedOn())
	})

	cache.Add("someKey", 0, "This is a test!!")
	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data())
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	cache.Delete("someKey")

	res = cache.Add("anotherKey", 3*time.Second, "This is anothertest")

	res.SetAboutToExpireCallback(func(key interface{}) {
		fmt.Println("Ablout to expire:", key.(string))
	})
	time.Sleep(5 * time.Second)
}
