package fdlimit

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/fdlimit"
)

// TestFileDescriptorLimits simply tests whether the file descriptor allowance
// per this process can be retrieved.
// TestFileDescriptorLimits 函数简单地测试了当前进程是否能够检索其文件描述符的允许数量。
func TestFileDescriptorLimits(t *testing.T) {
	if true {
		hardlimit, err := fdlimit.Maximum() // 返回该进程可以为自身请求的文件描述符的最大数量，实际上调用了 Current 函数。
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(hardlimit)
	}
	if true {
		hardlimit, err := fdlimit.Current() // 用于获取当前进程允许打开的文件描述符数量，返回的是 hardlimit 的值。
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(hardlimit)
	}
	if true {
		hardlimit, err := fdlimit.Raise(1) // 旨在尝试将该进程的文件描述符分配量最大化到操作系统允许的最大硬限制。如果传入的 max 超过了 hardlimit，则会返回 hardlimit 并附带一个错误消息。
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(hardlimit)
	}
	if false {
		target := 4096
		hardlimit, err := fdlimit.Maximum() // 返回该进程可以为自身请求的文件描述符的最大数量，实际上调用了 Current 函数。
		if err != nil {
			t.Fatal(err)
		}
		if hardlimit < target {
			t.Skipf("system limit is less than desired test target: %d < %d", hardlimit, target)
		}

		if limit, err := fdlimit.Current(); err != nil || limit <= 0 { // 用于获取当前进程允许打开的文件描述符数量，返回的是 hardlimit 的值。
			t.Fatalf("failed to retrieve file descriptor limit (%d): %v", limit, err)
		}
		if _, err := fdlimit.Raise(uint64(target)); err != nil { // 旨在尝试将该进程的文件描述符分配量最大化到操作系统允许的最大硬限制。如果传入的 max 超过了 hardlimit，则会返回 hardlimit 并附带一个错误消息。
			t.Fatalf("failed to raise file allowance")
		}
		if limit, err := fdlimit.Current(); err != nil || limit < target {
			t.Fatalf("failed to retrieve raised descriptor limit (have %v, want %v): %v", limit, target, err)
		}
	}
}
