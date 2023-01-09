package util

import (
	"runtime/debug"
	"fmt"
	"os"
)

func Fatal(v interface{}) {
	fmt.Printf("gorvld:\033[0;1;31m fatal:\033[0m %v\n", v)  // 加粗的红色 恢复默认 
	debug.PrintStack()  // 打印堆栈信息
	os.Exit(1)
}
