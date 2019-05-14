package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"

	"github.com/axgle/mahonia"
)

const (
	batPath = "C:\\Users\\Tang\\Documents\\WXWork\\1688852916302708\\Cache\\File\\2019-03\\CallBat\\bin\\Debug\\ci\\build.cmd"
)

func main() {
	build(batPath)
}

func build(path string) {
	cmd := exec.Command("cmd.exe", "/C", path)
	var exitCode = 0
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			status := exitErr.Sys().(syscall.WaitStatus)
			switch {
			case status.Exited():
				exitCode = status.ExitStatus()
			}
		} else {
			fmt.Printf("Return other error: %s\n", err)
		}
	}
	if exitCode != 0 {

	}
	fmt.Printf("in all caps: %q\n", Encoding(out.String(), "gbk", "utf8"))
}

//编码转换
func Encoding(input string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(input)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
