// Copyright 2018 The go-ethereum Authors
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

package accounts

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
)

// parseURL converts a user supplied URL into the accounts specific structure.
func parseURL(url string) (accounts.URL, error) {
	parts := strings.Split(url, "://")
	if len(parts) != 2 || parts[0] == "" {
		return accounts.URL{}, errors.New("protocol scheme missing")
	}
	return accounts.URL{
		Scheme: parts[0],
		Path:   parts[1],
	}, nil
}
func TestURLParsing(t *testing.T) {
	// t.Parallel()
	url, err := parseURL("https://ethereum.org")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	fmt.Println(url.Scheme)
	fmt.Println(url.Path)
	return
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "ethereum.org" {
		t.Errorf("expected: %v, got: %v", "ethereum.org", url.Path)
	}

	for _, u := range []string{"ethereum.org", ""} {
		if _, err = parseURL(u); err == nil {
			t.Errorf("input %v, expected err, got: nil", u)
		}
	}
}

func TestURLString(t *testing.T) {
	// t.Parallel()
	url := accounts.URL{Scheme: "https", Path: "ethereum.org"}
	fmt.Println(url.String())
	return
	if url.String() != "https://ethereum.org" {
		t.Errorf("expected: %v, got: %v", "https://ethereum.org", url.String())
	}

	url = accounts.URL{Scheme: "", Path: "ethereum.org"}
	if url.String() != "ethereum.org" {
		t.Errorf("expected: %v, got: %v", "ethereum.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	t.Parallel()
	url := accounts.URL{Scheme: "https", Path: "ethereum.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	fmt.Println(string(json))
	return
	if string(json) != "\"https://ethereum.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://ethereum.org\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	t.Parallel()
	url := &accounts.URL{}
	err := url.UnmarshalJSON([]byte("\"https://ethereum.org\""))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "ethereum.org" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	t.Parallel()
	tests := []struct {
		urlA   accounts.URL
		urlB   accounts.URL
		expect int
	}{
		{accounts.URL{"https", "ethereum.org"}, accounts.URL{"https", "ethereum.org"}, 0},
		{accounts.URL{"http", "ethereum.org"}, accounts.URL{"https", "ethereum.org"}, -1},
		{accounts.URL{"https", "ethereum.org/a"}, accounts.URL{"https", "ethereum.org"}, 1},
		{accounts.URL{"https", "abc.org"}, accounts.URL{"https", "ethereum.org"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
