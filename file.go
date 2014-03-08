package common

import (
	"os"
)

/* 判断文件是否存在
 * os.Stat(name) 获取文件名name的文件信息
 * os.IsExist(err) 判断err是否是文件或目录已存在的错误
 */
func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}
