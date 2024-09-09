package pkg

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/ordinary-xiyv/qzipgo/src/internal"
)

// 声明一个全局变量 QatService
var QatService *internal.QatService = &internal.QatService{}

// 检查是否qat是否可用
func Available() bool {
	envAvailable := internal.CheckQATEnv(QatService)
	hwAvailable := internal.CheckQATHWState(QatService)
	if envAvailable && hwAvailable && hwIsAvailable(QatService.Hardwareses) && RunCompressTest() && RunDecompressTest() {
		log.Println("\033[1;32m ****** QAT服务可用 ******\033[0m")
	} else {
		log.Print("\033[1;31m ****** QAT服务不可用 ****** \033[0m")
		return false
	}

	//fmt.Print(QatService.String())
	return true
}

func hwIsAvailable(hws []*internal.Hardwarese) bool {
	var availableCount int = 0
	for _, hw := range hws {
		if hw.State == nil || *hw.State == "down" {
			log.Println("设备", *hw.Name, "不可用")
			continue
		} else if *hw.State == "up" {
			availableCount = availableCount + 1
		}
	}
	return availableCount != 0
}

// 进行简单的压缩测试
func RunCompressTest() bool {
	log.Println("***********************进行简单的压缩测试***********************")
	// 定义相对路径
	relPath := "../testfiles/test_15mb.json"

	// 获取当前工作目录
	absPath, err := filepath.Abs(relPath)
	log.Println(absPath)
	if err != nil {
		log.Printf("获取当前工作目录时出错: %s\n", err)
		return false
	}
	cmd := exec.Command("qzip", "-k", relPath)

	// 获取命令输出
	output, err := cmd.Output()
	if err != nil {
		// 获取退出状态
		if exitError, ok := err.(*exec.ExitError); ok {
			// 获取退出状态码
			exitCode := exitError.ExitCode()
			log.Printf("命令执行失败，退出状态码: %d\n", exitCode)
		} else {
			log.Printf("执行命令时出错: %s\n", err)
		}
		return false
	}

	// 打印命令输出
	log.Printf("命令输出:\n%s\n", output)
	log.Println("***********************压缩测试结束***********************")
	return true
}

// 进行简单的解压测试
func RunDecompressTest() bool {
	log.Println("***********************进行简单的解压测试***********************")
	// 定义相对路径
	relPath := "../testfiles/test_15mb.json.gz"

	// 获取当前工作目录
	absPath, err := filepath.Abs(relPath)
	log.Println(absPath)
	if err != nil {
		log.Printf("获取当前工作目录时出错: %s\n", err)
		return false
	}
	cmd := exec.Command("qzip", "-d", relPath)
	// 获取命令输出
	output, err := cmd.Output()
	if err != nil {
		// 获取退出状态
		if exitError, ok := err.(*exec.ExitError); ok {
			// 获取退出状态码
			exitCode := exitError.ExitCode()
			log.Printf("命令执行失败，退出状态码: %d\n", exitCode)
		} else {
			log.Printf("执行命令时出错: %s\n", err)
		}
		return false
	}

	// 打印命令输出
	log.Printf("命令输出:\n%s\n", output)
	log.Println("***********************解压测试结束***********************")
	return true
}
