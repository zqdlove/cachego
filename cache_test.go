package cachego

import (
//	"bytes"
//	"log"
//	"strconv"
//	"sync"
//	"sync/atomic"
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
