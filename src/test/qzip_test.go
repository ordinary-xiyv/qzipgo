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
	cmd.InputFile = append(cmd.InputFile, "/tmp/test.txt")

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
	cmd.InputFile = append(cmd.InputFile, "/tmp/test.txt.gz")

	if err := internal.ExecuteQzipCommand(cmd); err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

// qzip -k -o filepath filepath 测试压缩
// output:Executing command: /usr/local/bin/qzip -k -o /tmp/test1111111.txt /tmp/test.txt
func TestCompressWithOutputFile(t *testing.T) {
	err := pkg.CompressWithOutputFile("/tmp/test.txt", "/tmp/test_with_output_file.txt")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

// TestCompressDictory tests compressing a directory with the CompressDictory function.
func TestCompressDictory(t *testing.T) {

	err := pkg.CompressDictoryByEveryFile("/tmp/test")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

func TestDeCompressDictory(t *testing.T) {

	err := pkg.DecompressDictoryByEveryFile("/tmp/test")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

// TestCompressFiles tests compressing multiple files with the CompressFiles function.
func TestCompressFiles(t *testing.T) {

	err := pkg.CompressFiles("/tmp/test/1.txt", "/tmp/test/2.txt")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

func TestDeCompressFiles(t *testing.T) {

	err := pkg.DecompressFiles("/tmp/test/1.txt.gz", "/tmp/test/2.txt.gz")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

func TestCompressFilesWithBusyPoll(t *testing.T) {

	err := pkg.CompressDictoryWithBusyPoll("/tmp/test")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

func TestDeCompressFilesWithBusyPoll(t *testing.T) {

	err := pkg.DecompressDictoryWithBusyPoll("/tmp/test")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}
