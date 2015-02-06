package main

import (
	"fmt"
)

var gopathstr []string

func gopaths() {
	gopathstr = GetGOPATHS()
	fmt.Printf("共有"+CL_CYAN+"%d"+CL_DEFAULT+"个GOPATH.\r\n", len(gopathstr))
}

func gopath(i int, path string) {
	fmt.Printf("开始分析第"+CL_CYAN+"%d"+CL_DEFAULT+"个GOPATH:%s.\r\n", i, path)
	gits := GetGitPath(path)
	igits := len(gits) - 1
	fmt.Printf("此GOPATH共有"+CL_CYAN+"%d"+CL_DEFAULT+"个Github库.\r\n", igits)
	for i := 0; i < igits; i++ {
		fmt.Printf("更新"+CL_CYAN+"(%d/%d)"+CL_DEFAULT, i+1, igits)

		b := GitPull(gits[i])
		if b {
			fmt.Printf("成功 [%d%%]\r\n", int((i+1)*100/igits))
		} else {
			fmt.Printf(CL_RED+"失败"+CL_DEFAULT+" [%d%%]\r\n", int((i+1)*100/igits))
		}
	}
	fmt.Printf("更新第"+CL_CYAN+"%d"+CL_DEFAULT+"个GOPATH完毕.\r\n", i)
}

func main() {
	fmt.Println("欢迎使用AutoPull。")
	fmt.Println("分析$GOPATH")
	gopaths()
	for i, v := range gopathstr {
		gopath(i+1, v)
	}
	fmt.Println("更新完毕!")
}
