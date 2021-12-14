package dependencymanagement

/*
1.同一环境下, 不同项目只能使用同一包的同一版本 (package get下来后都会放到gopath路径中, 依赖查找顺序gopath->goroot)
2.无法管理对包的特定版本的依赖


1.5之后vender目录
查找依赖包路径的解决方案如下:
	1.当前包下的vender目录
	2.向上级目录查找, 知道src下的vender目录
	3.gopath
	4.goroot


	godep glode dep

	glide init
	glide install
*/
