// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).

package prque

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/prque"
)

// Prque类似于zset
func TestPrque(t *testing.T) {
	// Generate a batch of random data and a specific priority order
	size := 16 * blockSize
	// prio := rand.Perm(size)
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	queue := prque.New[int, int](nil)
	queue.Push(3, 3)
	queue.Push(1, 1)
	queue.Push(2, 20)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	// for rep := 0; rep < 2; rep++ {
	// 	// Fill a priority queue with the above data
	// 	for i := 0; i < size; i++ {
	// 		queue.Push(data[i], prio[i])
	// 		if queue.Size() != i+1 {
	// 			t.Errorf("queue size mismatch: have %v, want %v.", queue.Size(), i+1)
	// 		}
	// 	}
	// 	// Create a map the values to the priorities for easier verification
	// 	dict := make(map[int]int)
	// 	for i := 0; i < size; i++ {
	// 		dict[prio[i]] = data[i]
	// 	}

	// 	// Pop out the elements in priority order and verify them
	// 	prevPrio := size + 1
	// 	for !queue.Empty() {
	// 		val, prio := queue.Pop()
	// 		if prio > prevPrio {
	// 			t.Errorf("invalid priority order: %v after %v.", prio, prevPrio)
	// 		}
	// 		prevPrio = prio
	// 		if val != dict[prio] {
	// 			t.Errorf("push/pop mismatch: have %v, want %v.", val, dict[prio])
	// 		}
	// 		delete(dict, prio)
	// 	}
	// }
}

const blockSize = 4096

// 清空
func TestReset(t *testing.T) {
	// Generate a batch of random data and a specific priority order
	size := 16 * blockSize
	// prio := rand.Perm(size)
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	queue := prque.New[int, int](nil)
	queue.Push(3, 3)
	queue.Push(1, 1)
	queue.Push(2, 20)
	queue.Reset() //清空
	if queue.Size() != 0 {
		fmt.Println("不空")
	} else {
		fmt.Println("空")
	}
	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Pop())
	// fmt.Println(queue.Pop())

	// // Generate a batch of random data and a specific priority order
	// size := 16 * blockSize
	// prio := rand.Perm(size)
	// data := make([]int, size)
	// for i := 0; i < size; i++ {
	// 	data[i] = rand.Int()
	// }
	// queue := prque.New[int, int](nil)

	// for rep := 0; rep < 2; rep++ {
	// 	// Fill a priority queue with the above data
	// 	for i := 0; i < size; i++ {
	// 		queue.Push(data[i], prio[i])
	// 		if queue.Size() != i+1 {
	// 			t.Errorf("queue size mismatch: have %v, want %v.", queue.Size(), i+1)
	// 		}
	// 	}
	// 	// Create a map the values to the priorities for easier verification
	// 	dict := make(map[int]int)
	// 	for i := 0; i < size; i++ {
	// 		dict[prio[i]] = data[i]
	// 	}
	// 	// Pop out half the elements in priority order and verify them
	// 	prevPrio := size + 1
	// 	for i := 0; i < size/2; i++ {
	// 		val, prio := queue.Pop()
	// 		if prio > prevPrio {
	// 			t.Errorf("invalid priority order: %v after %v.", prio, prevPrio)
	// 		}
	// 		prevPrio = prio
	// 		if val != dict[prio] {
	// 			t.Errorf("push/pop mismatch: have %v, want %v.", val, dict[prio])
	// 		}
	// 		delete(dict, prio)
	// 	}
	// 	// Reset and ensure it's empty
	// 	queue.Reset()
	// 	if !queue.Empty() {
	// 		t.Errorf("priority queue not empty after reset: %v", queue)
	// 	}
	// }
}
