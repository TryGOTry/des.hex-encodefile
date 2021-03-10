package main

import (
	"desencode-shellcode/root"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入加密的文件名称!")
	} else {
		deskey := root.RandNewStr(8)
		filename := os.Args[1]
		key := []byte(deskey)
		code := root.Readcode(filename)
		ncode := hex.EncodeToString([]byte(code))  //hex编码
		descode, err := root.Encrypt(ncode, key) //前半段进行des加密
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("key:",deskey,"\ncode:",descode)
	}
}
