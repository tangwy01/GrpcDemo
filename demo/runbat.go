package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"

	"github.com/axgle/mahonia"
)

const (
	batPath = `D:\WorkSpace\源代码\70\Test\ERP352SP4\ci\build.bat`
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
		}
	}

	if exitCode != 0 {

	}
	fmt.Printf("in all caps: %q\n", Encoding(out.String(), "gbk", "utf8"))
	time.Sleep(5 * time.Minute)
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

// @echo off
// rem 准备编译......

// %~d0
// cd %~dp0
// C:\Windows\Microsoft.NET\Framework64\v3.5\msbuild.exe ../分支2/明源整体解决方案/明源整体解决方案.sln  /t:Build /p:Configuration="Debug"
// if errorlevel 1 goto errorDone

// exit /b 0

// :errorDone
// exit /b 2
