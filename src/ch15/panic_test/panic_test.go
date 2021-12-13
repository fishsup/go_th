package panic_test

import (
	"errors"
	"fmt"
	"testing"
)

/* panic
1.用于不可恢复的错误
2.panic退出前会执行defer指定的内容

panic vs os.Exit
1.os.Exit退出时不会调用defer指定的函数
2.os.Exit退出时不输出大年调用栈信息
*/
func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	fmt.Println("finally")
	// }()
	//错误恢复
	defer func() {
		//单纯打印 异常恢复有几率导致僵尸进程,例如资源消耗光时的异常 -> Let it crash
		if err := recover(); err != nil {
			fmt.Println("revocer from ", err)
		}
	}()

	fmt.Println("start")

	panic(errors.New("panic error"))
}
