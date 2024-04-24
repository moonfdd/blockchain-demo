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
	"encoding/binary"
	"fmt"
	"testing"
	"unsafe"

	"github.com/ethereum/go-ethereum/common/lru"
)

type testKey [8]byte

func mkKey(i int) (key testKey) {
	binary.LittleEndian.PutUint64(key[:], uint64(i))
	return key
}

func getSize(lru *lru.SizeConstrainedCache[testKey, []byte]) uint64 {
	return *(*uint64)(unsafe.Pointer(lru))
}

// 容量固定的LRU
func TestSizeConstrainedCache(t *testing.T) {
	lru := lru.NewSizeConstrainedCache[testKey, []byte](100)

	// Add 11 items of 10 byte each. First item should be swapped out
	for i := 0; i < 109; i++ {
		k := mkKey(i)
		v := fmt.Sprintf("value-%d", i)
		lru.Add(k, []byte(v))

	}
	if true {
		k := mkKey(109)
		v := fmt.Sprintf("value-%d,value-%d,value-%d", 109, 110, 111)
		lru.Add(k, []byte(v))
	}
	for i := 0; i < 110; i++ {
		fmt.Println(lru.Get(mkKey(i)))
	}
	fmt.Println("字节数", getSize(lru))

	// Zero:th should be evicted
	{
		k := mkKey(0)
		fmt.Println(lru.Get(k))
		if _, ok := lru.Get(k); ok {
			t.Fatalf("should be evicted: %v", k)
		}
	}

	// Elems 1-11 should be present
	for i := 102; i < 110; i++ {
		k := mkKey(i)
		// want := fmt.Sprintf("value-%04d", i)
		have, ok := lru.Get(k)
		if !ok {
			t.Fatalf("missing key %v", k)
		}
		fmt.Println(have, ok)
		// if string(have) != want {
		// 	t.Fatalf("wrong value, have %v want %v", have, want)
		// }
	}
}

// This test adds inserting an element exceeding the max size.
// 容量只有100，插入字节是200的值，插入成功
func TestSizeConstrainedCacheOverflow(t *testing.T) {
	lru := lru.NewSizeConstrainedCache[testKey, []byte](100)

	// Add 10 items of 10 byte each, filling the cache
	for i := 0; i < 10; i++ {
		k := mkKey(i)
		v := fmt.Sprintf("value-%04d", i)
		lru.Add(k, []byte(v))
	}
	fmt.Println(getSize(lru))
	// Add one single large elem. We expect it to swap out all entries.
	{
		k := mkKey(1337)
		v := make([]byte, 200)
		lru.Add(k, v) //容量只有100，插入字节是200的值，插入成功
	}
	fmt.Println(lru.Get(mkKey(1337)))

	// Elems 0-9 should be missing
	for i := 1; i < 10; i++ {
		k := mkKey(i)
		if _, ok := lru.Get(k); ok {
			t.Fatalf("should be evicted: %v", k)
		}
		fmt.Println(lru.Get(k))
	}
	return

}

// 多次插入key相同，值不同
func TestSizeConstrainedCacheSameItem(t *testing.T) {
	lru := lru.NewSizeConstrainedCache[testKey, []byte](100)

	// Add one 10 byte-item 10 times.
	k := mkKey(0)
	// v := fmt.Sprintf("value-%04d", 0)
	for i := 0; i < 10; i++ {
		v := fmt.Sprintf("value-%04d", i)
		fmt.Println(lru.Get(k))
		lru.Add(k, []byte(v))
	}

	return

	// The size should be accurate.
	if have, want := getSize(lru), uint64(10); have != want {
		t.Fatalf("size wrong, have %d want %d", have, want)
	}
}

// 添加空或者nil
func TestSizeConstrainedCacheEmpties(t *testing.T) {
	lru := lru.NewSizeConstrainedCache[testKey, []byte](100)

	// This test abuses the lru a bit, using different keys for identical value(s).
	for i := 0; i < 10; i++ {
		lru.Add(testKey{byte(i)}, []byte{})  //添加空
		lru.Add(testKey{byte(255 - i)}, nil) //添加nil
	}

	for i := 0; i < 10; i++ {
		fmt.Print(byte(i))
		fmt.Println(lru.Get(testKey{byte(i)}))
		fmt.Print(byte(255 - i))
		fmt.Println(lru.Get(testKey{byte(255 - i)}))
	}

	fmt.Println("getSize(lru) = ", getSize(lru))
	return

	// The size should not count, only the values count. So this could be a DoS
	// since it basically has no cap, and it is intentionally overloaded with
	// different-keyed 0-length values.

	if have, want := getSize(lru), uint64(0); have != want {
		t.Fatalf("size wrong, have %d want %d", have, want)
	}

	for i := 0; i < 10; i++ {
		if v, ok := lru.Get(testKey{byte(i)}); !ok {
			t.Fatalf("test %d: expected presence", i)
		} else if v == nil {
			t.Fatalf("test %d, v is nil", i)
		}

		if v, ok := lru.Get(testKey{byte(255 - i)}); !ok {
			t.Fatalf("test %d: expected presence", i)
		} else if v != nil {
			t.Fatalf("test %d, v is not nil", i)
		}
	}
}
