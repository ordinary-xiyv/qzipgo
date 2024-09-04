package internal

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

const (
	// ANSI 转义码设置绿色
	greenBold = "\033[1;32m"
	reset     = "\033[0m"
)

type QatService struct {
	IcpRoot     *string
	QzRoot      *string
	Hardwareses []*Hardwarese
}

type Hardwarese struct {
	Name   *string
	State  *string
	HwType *string
}

func CheckQATEnv(qat *QatService) {
	// 获取环境变量
	icpRoot := os.Getenv("ICP_ROOT")
	qzRoot := os.Getenv("QZ_ROOT")

	// 检查环境变量是否存在且不为空
	if icpRoot == "" {
		log.Fatalln("ICP_ROOT 环境变量不存在或为空")
	} else {
		log.Println(greenBold, "ICP_ROOT:", icpRoot, reset)

		qat.IcpRoot = &icpRoot
	}

	if qzRoot == "" {
		log.Fatalln("QZ_ROOT 环境变量不存在或为空")
	} else {
		log.Println(greenBold, "QZ_ROOT:", qzRoot, reset)

		qat.QzRoot = &qzRoot
	}
}

func CheckQATHWState(qat *QatService) {
	// 执行命令并获取输出
	// 创建命令
	cmd := exec.Command("service", "qat_service", "status")

	// 获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行命令时出错: %s\n", err)
		return
	}

	// output := `
	// Checking status of all devices.
	// There is 4 QAT acceleration device(s) in the system:
	//  qat_dev0 - type: 4xxx,  inst_id: 0,  node_id: 0,  bsf: 0000:76:00.0,  #accel: 1 #engines: 9 state: up
	//  qat_dev1 - type: 4xxx,  inst_id: 1,  node_id: 0,  bsf: 0000:7a:00.0,  #accel: 1 #engines: 9 state: up
	//  qat_dev2 - type: 4xxx,  inst_id: 2,  node_id: 1,  bsf: 0000:f3:00.0,  #accel: 1 #engines: 9 state: up
	//  qat_dev3 - type: 4xxx,  inst_id: 3,  node_id: 1,  bsf: 0000:f7:00.0,  #accel: 1 #engines: 9 state: up
	// `

	// 使用正则表达式提取所有字段
	re := regexp.MustCompile(`(qat_dev\d+) - type:\s*(\S+), .*state: (\S+)`)
	matches := re.FindAllStringSubmatch(string(output), -1)

	if len(matches) == 0 {
		log.Println("\033[1;31m未找到任何设备的信息\033[0m")
		return
	}

	// 遍历所有匹配项并输出
	var hardwareses = make([]*Hardwarese, 0)
	for _, match := range matches {
		log.Printf("%s%s type: %s, state: %s%s\n", greenBold, match[1], match[2], match[3], reset)

		hardwareses = append(hardwareses, &Hardwarese{
			Name:   &match[1],
			State:  &match[3],
			HwType: &match[2],
		})
	}

	qat.Hardwareses = hardwareses

}

func (qat *QatService) String() string {
	hwStrings := ""
	for _, hw := range qat.Hardwareses {
		hwStrings = hwStrings + hw.String()
	}
	QatService := fmt.Sprintf("ICP_ROOT:\n\t%s\nQZ_ROOT:\n\t%s\nHWs:\n%s\n", *qat.IcpRoot, *qat.QzRoot, hwStrings)
	return QatService
}

func (hw *Hardwarese) String() string {
	Hardwarese := fmt.Sprintf("\tName: %s HwType: %s State: %s\n", *hw.Name, *hw.HwType, *hw.State)
	return Hardwarese
}
