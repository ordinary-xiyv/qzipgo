package test

import (
	"fmt"
	"testing"

	"github.com/ordinary-xiyv/qzipgo/src/pkg"
)

func TestQzip(t *testing.T) {
	fmt.Println("进行简单的压缩测试")
	pkg.RunCompressTest()
	fmt.Println("进行简单的解压测试")
	pkg.RunDecompressTest()
}