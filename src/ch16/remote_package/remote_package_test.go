package remotepackage

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)
/* package
1.通过go get来获取远程依赖  go get -u强制从网络更新远程依赖
2. 注意代码在远程仓库上的组织形式以适应go get (直接以代码路径开始, 不要有src)*/
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
