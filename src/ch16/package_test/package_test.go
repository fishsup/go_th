package package_test

import (
	"go_ph/src/ch16/series"
	"testing"
)

/* package
1.基本复用模块单元, 以首字母大写来表明可被包外代码访问, 小写开头不允许被包外访问
2.代码的package可以和所在的目录不一致
3.同一目录里go代码的package要保持一致
*/
func TestPackage(t *testing.T) {
	t.Log(series.GetFibList(-1))
	t.Log(series.GetFibList(2))
}

