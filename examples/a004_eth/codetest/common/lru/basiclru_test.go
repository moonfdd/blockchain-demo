// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package lru

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/lru"
)

// Some of these test cases were adapted
// from https://github.com/hashicorp/golang-lru/blob/master/simplelru/lru_test.go

func TestBasicLRU(t *testing.T) {
	cache := lru.NewBasicLRU[int, int](128)

	for i := 0; i < 256; i++ {
		r := cache.Add(i, i) //如果返回true，说明有删除元素
		fmt.Println(i, "新增", r)
	}
	if cache.Len() != 128 {
		t.Fatalf("bad len: %v", cache.Len())
	}

	// Check that Keys returns least-recent key first.
	keys := cache.Keys()
	if len(keys) != 128 {
		t.Fatal("wrong Keys() length", len(keys))
	}
	for i, k := range keys {
		v, ok := cache.Peek(k)
		fmt.Println("v = ", v)
		if !ok {
			t.Fatalf("expected key %d be present", i)
		}
		if v != k {
			t.Fatalf("expected %d == %d", k, v)
		}
		if v != i+128 {
			t.Fatalf("wrong value at key %d: %d, want %d", i, v, i+128)
		}
	}

	for i := 0; i < 128; i++ {
		v, ok := cache.Get(i)
		fmt.Println("v1", v, ok)
		if ok {
			t.Fatalf("%d should be evicted", i)
		}
	}
	for i := 128; i < 256; i++ {
		v, ok := cache.Get(i)
		fmt.Println("v2", v, ok)
		if !ok {
			t.Fatalf("%d should not be evicted", i)
		}
	}

	for i := 128; i < 192; i++ {
		ok := cache.Remove(i)
		if !ok {
			t.Fatalf("%d should be in cache", i)
		}
		ok = cache.Remove(i)
		if ok {
			t.Fatalf("%d should not be in cache", i)
		}
		_, ok = cache.Get(i)
		if ok {
			t.Fatalf("%d should be deleted", i)
		}
	}

	// Request item 192.
	cache.Get(192)
	// It should be the last item returned by Keys().
	for i, k := range cache.Keys() {
		if (i < 63 && k != i+193) || (i == 63 && k != 192) {
			t.Fatalf("out of order key: %v", k)
		}
	}

	cache.Purge() //清空缓存
	if cache.Len() != 0 {
		t.Fatalf("bad len: %v", cache.Len())
	}
	fmt.Println(cache.Get(200))
	if _, ok := cache.Get(200); ok {
		t.Fatalf("should contain nothing")
	}
}

// 会覆盖
func TestBasicLRUAddExistingKey(t *testing.T) {
	cache := lru.NewBasicLRU[int, int](1)
	cache.Add(1, 1)
	cache.Add(1, 2)

	v, ok := cache.Get(1)
	if v != 2 {
		t.Fatal("wrong value:", v)
	}
	fmt.Println("v", v, ok)
}

// This test checks GetOldest and RemoveOldest.
func TestBasicLRUGetOldest(t *testing.T) {
	cache := lru.NewBasicLRU[int, int](128)
	for i := 0; i < 256; i++ {
		cache.Add(i, i)
	}

	k, _, ok := cache.GetOldest()
	if !ok {
		t.Fatalf("missing")
	}
	if k != 128 {
		t.Fatalf("bad: %v", k)
	}

	k, _, ok = cache.RemoveOldest()
	if !ok {
		t.Fatalf("missing")
	}
	if k != 128 {
		t.Fatalf("bad: %v", k)
	}

	k, _, ok = cache.RemoveOldest()
	if !ok {
		t.Fatalf("missing oldest item")
	}
	if k != 129 {
		t.Fatalf("wrong oldest item: %v", k)
	}
}

// Test that Add returns true/false if an eviction occurred
func TestBasicLRUAddReturnValue(t *testing.T) {
	cache := lru.NewBasicLRU[int, int](1)
	if cache.Add(1, 1) {
		t.Errorf("first add shouldn't have evicted")
	}
	if !cache.Add(2, 2) {
		t.Errorf("second add should have evicted")
	}
}

// This test verifies that Contains doesn't change item recency.
func TestBasicLRUContains(t *testing.T) {
	cache := lru.NewBasicLRU[int, int](2)
	cache.Add(1, 1)
	cache.Add(2, 2)
	if !cache.Contains(1) {
		t.Errorf("1 should be in the cache")
	}
	cache.Add(3, 3)
	if cache.Contains(1) {
		t.Errorf("Contains should not have updated recency of 1")
	}
}

// Test that Peek doesn't update recent-ness
func TestBasicLRUPeek(t *testing.T) {
	cache := lru.NewBasicLRU[int, int](2)
	cache.Add(1, 1)
	cache.Add(2, 2)
	if v, ok := cache.Peek(1); !ok || v != 1 {
		t.Errorf("1 should be set to 1")
	}
	cache.Add(3, 3)
	if cache.Contains(1) {
		t.Errorf("should not have updated recent-ness of 1")
	}
}
