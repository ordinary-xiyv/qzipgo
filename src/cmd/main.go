package main

import (
	"fmt"
	"github.com/ordinary-xiyv/qzipgo/src/pkg"
)

func main() {

	fmt.Println("=========================")
	// 检查是否qat是否可用
	pkg.Available()

	// // 进行简单的压缩测试
	// fmt.Println("进行简单的压缩测试")
	// pkg.RunCompressTest()

	// // 进行简单的解压测试
	// fmt.Println("进行简单的解压测试")
	// pkg.RunDecompressTest()

	fmt.Println("=========================")
}
