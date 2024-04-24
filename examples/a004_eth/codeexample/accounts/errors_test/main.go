package main
import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
)
func main(){
	if false{
		fmt.Println(accounts.ErrUnknownAccount)
		fmt.Println(accounts.ErrUnknownWallet)
		fmt.Println(accounts.ErrNotSupported)
		fmt.Println(accounts.ErrInvalidPassphrase)
		fmt.Println(accounts.ErrWalletAlreadyOpen)
		fmt.Println(accounts.ErrWalletClosed)
	}
	if true{
		fmt.Println(accounts.NewAuthNeededError("needed"))
	}
}