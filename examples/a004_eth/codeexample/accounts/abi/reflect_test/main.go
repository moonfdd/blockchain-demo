package main

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// ConvertType 未完成
	if true {
		// Test Basic Struct
		type T struct {
			X *big.Int
			Y *big.Int
		}
		// Create on-the-fly structure
		var fields []reflect.StructField
		fields = append(fields, reflect.StructField{
			Name: "X",
			Type: reflect.TypeOf(new(big.Int)),
			Tag:  "json:\"" + "x" + "\"",
		})
		fields = append(fields, reflect.StructField{
			Name: "Y",
			Type: reflect.TypeOf(new(big.Int)),
			Tag:  "json:\"" + "y" + "\"",
		})
		val := reflect.New(reflect.StructOf(fields))
		val.Elem().Field(0).Set(reflect.ValueOf(big.NewInt(1)))
		val.Elem().Field(1).Set(reflect.ValueOf(big.NewInt(2)))
		// ConvertType
		out := *abi.ConvertType(val.Interface(), new(T)).(*T)
		fmt.Println(out)
	}
	fmt.Println("")
}
