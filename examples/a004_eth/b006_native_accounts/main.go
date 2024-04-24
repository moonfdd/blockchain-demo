package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	//创建账户管理器
	ks := keystore.NewKeyStore("test/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	data, _ := json.MarshalIndent(am, "", "  ")
	fmt.Println(string(data))

	//创建账户
	// Create a new account with the specified encryption passphrase.
	password1 := "password1"
	password2 := "password2"
	password3 := "password3"
	password4 := "password4"
	newAcc, err := ks.NewAccount(password1)
	if err != nil {
		fmt.Println("NewAccount失败", err)
		return
	}
	fmt.Println("NewAccount成功")
	fmt.Println("NewAccount = ", newAcc)

	// Export the newly created account with a different passphrase. The returned
	// data from this method invocation is a JSON encoded, encrypted key-file.
	// 这段注释描述的是在导出一个使用不同口令（passphrase）创建的新账户。调用该方法后返回的数据是一个 JSON 编码、加密的密钥文件。
	jsonAcc, err := ks.Export(newAcc, password1, password2)
	if err != nil {
		fmt.Println("Export失败", err)
		return
	}
	fmt.Println("Export = ", string(jsonAcc))
	err = ks.Unlock(newAcc, password1)
	if err != nil {
		fmt.Println("Unloc失败", err)
		return
	}
	fmt.Println("Unlock成功")

	// Update the passphrase on the account created above inside the local keystore.
	err = ks.Update(newAcc, password1, password3)
	if err != nil {
		fmt.Println("Update失败", err)
		return
	}
	fmt.Println("Update成功")
	err = ks.Unlock(newAcc, password3)
	if err != nil {
		fmt.Println("Unlock2失败", err)
		return
	}
	fmt.Println("Unlock2成功")

	// Delete the account updated above from the local keystore.
	err = ks.Delete(newAcc, password3)
	if err != nil {
		fmt.Println("Delete失败", err)
		return
	}
	fmt.Println("Delete成功")

	// Import back the account we've exported (and then deleted) above with yet
	// again a fresh passphrase.
	// 这段注释描述的是重新导入先前导出（然后删除）的账户，并使用新的口令（passphrase）。
	impAcc, err := ks.Import(jsonAcc, password2, password4)
	if err != nil {
		fmt.Println("Import失败", err)
		return
	}
	fmt.Println("Import成功")
	data, _ = json.MarshalIndent(impAcc, "", "  ")
	fmt.Println(string(data))

	// Create a new account to sign transactions with
	signer, err := ks.NewAccount(password1)
	if err != nil {
		fmt.Println("NewAccount2失败", err)
		return
	}
	fmt.Println("NewAccount2成功")
	txHash := common.HexToHash("0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")

	// Sign a transaction with a single authorization
	signature, err := ks.SignHashWithPassphrase(signer, password1, txHash.Bytes())
	if err != nil {
		fmt.Println("SignHashWithPassphrase失败", err)
		return
	}
	fmt.Println("SignHashWithPassphrase成功")

	// Sign a transaction with multiple manually cancelled authorizations
	err = ks.Unlock(signer, password1)
	if err != nil {
		fmt.Println("Unlock失败", err)
		return
	}
	fmt.Println("Unlock成功")
	signature, err = ks.SignHash(signer, txHash.Bytes())
	if err != nil {
		fmt.Println("SignHash失败", err)
		return
	}
	fmt.Println("SignHash成功")
	err = ks.Lock(signer.Address)
	if err != nil {
		fmt.Println("Lock失败", err)
		return
	}
	fmt.Println("Lock成功")

	// Sign a transaction with multiple automatically cancelled authorizations
	err = ks.TimedUnlock(signer, password1, time.Second)
	if err != nil {
		fmt.Println("TimedUnlock失败", err)
		return
	}
	fmt.Println("TimedUnlock成功")
	signature, err = ks.SignHash(signer, txHash.Bytes())
	if err != nil {
		fmt.Println("SignHash失败", err)
		return
	}
	fmt.Println("SignHash成功")
	_ = signature
}
