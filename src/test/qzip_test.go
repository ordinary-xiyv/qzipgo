package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/ordinary-xiyv/qzipgo/src/pkg"
)

func TestQzip(t *testing.T) {
	fmt.Println(pkg.Available())
}

// qzip -k filepath 测试压缩
// output:Executing command: /usr/local/bin/qzip -k /tmp/test.txt
func TestCompress(t *testing.T) {
	err := pkg.CompressFile("/tmp/test.txt")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

// qzip -d -k filepath 测试解压
// output:Executing command: /usr/local/bin/qzip -d -k /tmp/test.txt
func TestDecompress(t *testing.T) {

	err := pkg.DecompressFile("/tmp/test.txt.gz")

	if err != nil {
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

func TestCompressDictoryByTar(t *testing.T) {

	err := pkg.CompressDictoryByTar("/tmp/test", "/tmp/test.tgz")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

func TestDeCompressDictoryByTar(t *testing.T) {

	err := pkg.DecompressDictoryByTar("/tmp/test.tgz", "/tmp")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}

// 输出目录为空，则会在代码的执行目录下生成解压文件，请务必指定输出目录！！！！
func TestDeCompressDictoryByTarWithNoOutputDictory(t *testing.T) {

	err := pkg.DecompressDictoryByTar("/tmp/test.tgz", "")

	if err != nil {
		log.Fatalf("Failed to execute qzip command: %s", err)
	}
}
