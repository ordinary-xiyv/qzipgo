package main

import (
	"fmt"

	"qzipgo/internal"
)

// 声明一个全局变量 QatService
var QatService *internal.QatService = &internal.QatService{}

func main() {
	fmt.Println("Hello, World!")
	internal.CheckQATEnv(QatService)
	internal.CheckQATHWState(QatService)

	fmt.Print(QatService.String())
}
