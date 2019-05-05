/*
@Time : 2019/5/5 10:43 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
 该工具包依赖于User32.dll, 仅限Win平台下使用
*/
package utils

import (
	"testing"
)

func TestGetWindowResolution(t *testing.T){
	x, y := GetWindowResolution()
	t.Log(x,"---",y)
	return
}









