package main

import (
	"fmt"
	"time"
	"github.com/zqdlove/cache3go"
)

type myStruct struct {
	text string
	moreData []byte
}

func main() {
	cache := cachego.Cache("mycache")

	val := myStruct{"This is a test!!!", []byte{}}
	cache.Add("someKey", 5*time.Second, &val)

	res, err := cache.Value("someKey")
	if err == nil {
		fmt.Println("Found value in cache:", res.Data().(*myStruct).text)
	} else {
		fmt.Println("Error retrieving value from cache:", err)
	}

	time.Sleep(6*time.Second)
	res, err = cache.Value("someKey")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	cache.Add("someKey", 0 , &val)

	cache.SetAboutToDeleteItemCallback(func(e * cachego.CacheItem){
	fmt.Println("Deleting:", e.Key(), e.Data().(*myStruct).text, e.CreatedOn())
	})
	fmt.Println("1")
	cache.Delete("someKey")
	fmt.Println("2")
	cache.Flush()
}
