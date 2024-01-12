package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("****Original:")
	original()

	fmt.Println("***Masked:")
	masked()

	fmt.Println("***Compressed:")
	compressed()
}

func original() {
	r := newUserRepository(newStdLogger(os.Stdout))
	r.register(1, "bob", "bob@email.com")
	fmt.Println("")
	r.register(2, "alice", "alice@email.com")
	fmt.Println("")
	r.allDisplay()
}

// ユーザー情報を保存する際に、emaillはログで出すときに見えないようにする仕様が追加された
func masked() {
	r := newUserRepository(newMaskedLogger(newStdLogger(os.Stdout)))
	r.register(1, "bob", "bob@email.com")
	fmt.Println("")
	r.register(2, "alice", "alice@email.com")
	fmt.Println("")
	r.allDisplay()
}

// ユーザー情報を保存する際に、圧縮してログに出す仕様が追加された
func compressed() {
	r := newUserRepository(newCompressLogger(newStdLogger(os.Stdout)))
	r.register(1, "bob", "bob@email.com")
	fmt.Println("")
	r.register(2, "alice", "alice@email.com")
	fmt.Println("")
	r.allDisplay()
}
