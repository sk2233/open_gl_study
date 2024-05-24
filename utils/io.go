/*
@author: sk
@date: 2023/5/3
*/
package utils

import (
	"os"
)

const (
	BasePath = "/Users/bytedance/Documents/go/openGL/"
)

func ReadAll(name string) []byte {
	bs, err := os.ReadFile(BasePath + name)
	HandleErr(err)
	return bs
}

func Has(name string) bool {
	stat, err := os.Stat(BasePath + name)
	return err == nil && stat != nil
}

func OpenFile(name string) *os.File {
	file, err := os.Open(BasePath + name)
	HandleErr(err)
	return file
}
