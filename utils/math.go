/*
@author: sk
@date: 2023/6/23
*/
package utils

func Min[T int](a, b T) T {
	if a < b {
		return a
	}
	return b
}
