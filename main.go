package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/trewanek-org/ddd-practice/di"
	"github.com/trewanek-org/ddd-practice/interface/controller"
)

var scanner = bufio.NewScanner(os.Stdin)

func scan() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func main() {
	con := controller.NewUser(di.GetUserUseCase())
	http.HandleFunc("/", con.Post) // ハンドラを登録してウェブページを表示させる
	log.Fatalln(http.ListenAndServe(":8080", nil))

	// CLIでの実装例
	//for {
	//	fmt.Println("Input user name")
	//	fmt.Println(">")
	//
	//	input := scan()
	//
	//	use := di.GetUserUseCase()
	//
	//	ctx := context.Background()
	//	err := use.Register(ctx, input)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	fmt.Println("--------------------------------------")
	//	fmt.Println("user created:")
	//	fmt.Println("--------------")
	//	fmt.Println("user name:")
	//	fmt.Println("-", input)
	//	fmt.Println("--------------------------------------")
	//
	//	var yesOrNo string
	//	for yesOrNo != "y" && yesOrNo != "n" {
	//		fmt.Println("continue? (y/n)")
	//		fmt.Println(">")
	//		yesOrNo = scan()
	//	}
	//	if yesOrNo == "n" {
	//		break
	//	}
	//}
}
