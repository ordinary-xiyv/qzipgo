package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ordinary-xiyv/qzipgo/src/internal"
	"github.com/ordinary-xiyv/qzipgo/src/pkg"
)

func TestQzip(t *testing.T) {
	fmt.Println("进行简单的压缩测试")
	pkg.RunCompressTest()
	fmt.Println("进行简单的解压测试")
	pkg.RunDecompressTest()
}

// qzip -k filepath 测试压缩
// output:Executing command: /usr/local/bin/qzip -k /tmp/test.txt
func TestCompress(t *testing.T) {
	cmd := internal.GetDefaultQzipCommand()
	cmd.KeepSource = true
	cmd.InputFile = "/tmp/test.txt"

	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

// qzip -d -k filepath 测试解压
// output:Executing command: /usr/local/bin/qzip -d -k /tmp/test.txt
func TestDecompress(t *testing.T) {
	cmd := internal.GetDefaultQzipCommand()
	cmd.KeepSource = true
	cmd.Compression = false
	cmd.InputFile = "/tmp/test.txt.gz"

	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}
