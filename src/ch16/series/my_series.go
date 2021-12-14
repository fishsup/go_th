package series

import "fmt"

func GetFibList(n int) ([]int, error) {
	if n < 1 || n > 100 {
		// errors.New("something..")
		return nil, fmt.Errorf("The parameter %d cannot be less than 1 or greater than 100", n)

	}
	if n == 1 {
		return []int{1}, nil
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

/* init方法
1.在main被执行前, 所有依赖的package的init方法都会被执行
2.不通包的init函数按照包导入的依赖关系决定执行顺序
3.每个包可以有多个init函数
4.包的每个源文件也可以有多个init函数, 这点比较特殊
*/
func init() {
	fmt.Println("series init1")
}

func init() {
	fmt.Println("series init2")
}
