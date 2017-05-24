package cachego

import (
//	"bytes"
//	"log"
//	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	k = "testkey"
	v = "testvalue"
)

func TestCache(t *testing.T){
	table := Cache("testCache")
	table.Add(k+"_1", 0*time.Second, v)
	table.Add(k+"_2", 1*time.Second, v)

	p, err := table.Value(k + "_1")
	if  err != nil || p == nil || p.Data().(string) != v {
		t.Error("Erroe retrieving non expiring data from cache", err)
	}
	p, err = table.Value(k + "_2")
	if err != nil || p == nil || p.Data().(string) != v {
		t.Error("Error retrieving data from cache", err)
	}
	if p.AccessCount() != 1 {
		t.Error("Error getting correct access count")
	}
	if p.LifeSpan() != 1*time.Second {
		t.Error("Error getting correct life-span")
	}
	if p.AccessedOn().Unix() == 0 {
		t.Error("Error getting access time")
	}
	if p.CreatedOn().Unix() == 0 {
		t.Error("Error getting creation time")
	}
}

func TestCacheExpire(t *testing.T) {
	table := Cache("testCache")

	table.Add(k+"_1", 100*time.Millisecond, v+"_1")
	table.Add(k+"_2", 125*time.Millisecond, v+"_2")

	time.Sleep(75 * time.Millisecond)

	_, err := table.Value(k + "_1")
	if err != nil {
		t.Error("Error retrieving value from cache:", err)
	}

	time.Sleep(75 * time.Millisecond)

	_, err = table.Value(k + "_1")
	if err != nil  {
		t.Error("Error retrieving value from cache:", err)
	}
	_, err = table.Value(k + "_2")
	if err == nil {
		t.Error("Found key whitch should have been expired by now")
	}
}

func TestExists(t *testing.T) {
	table := Cache("testExists")
	table.Add(k, 0, v)
	if !table.Exists(k) {
		t.Error("Error verifying data in cache")
	}
}

func TestNotFoundAdd(t *testing.T) {
	table := Cache("testNotFoundAdd")

	if !table.NotFoundAdd(k, 0, v) {
		t.Error("Error verifying NotFound, data not in cache")
	}

	if table.NotFoundAdd(k, 0, v) {
		t.Error("Error verifying NotFound data in cache")
	}
}

func TestNotFoundAddConcurrency(t *testing.T) {
	table := Cache("testNotFoundAddConcurrency")
	var finish sync.WaitGroup
	var added int32
	var idle int32

	fn := func(id int) {
		for i := 0; i < 100; i++ {
			if table.NotFoundAdd(i, 0, i+id) {
				atomic.AddInt32(&added, 1)
			} else {
				atomic.AddInt32(&idle, 1)
			}
			time.Sleep(0)
		}
		finish.Done()
	}
	finish.Add(10)
	go fn(0x0000)
	go fn(0x1100)
	go fn(0x2200)
	go fn(0x3300)
	go fn(0x4400)
	go fn(0x5500)
	go fn(0x6600)
	go fn(0x7700)
	go fn(0x8800)
	go fn(0x9900)
	finish.Wait()

	t.Log(added, idle)

	table.Foreach(func(key interface{}, item *CacheItem) {
		v, _ := item.Data().(int)
		k, _ := key.(int)
		t.Logf("%02x %04x\n", k, v)
	})
}

func TestCacheKeepAlive(t *testing.T) {
	table := Cache("testKeepAlive")
	p := table.Add(k, 100*time.Millisecond, v)

	time.Sleep(50 * time.Millisecond)
	p.KeepAlive()
	time.Sleep(75 * time.Millisecond)
	if !table.Exists(k) {
		t.Error("Error keeping item alive")
	}
	time.Sleep(75 * time.Millisecond)
	if table.Exists(k) {
		t.Error("Error expiring item after keeping it alive")
	}
}
