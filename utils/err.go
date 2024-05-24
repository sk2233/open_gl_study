/*
@author: sk
@date: 2023/5/3
*/
package utils

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
