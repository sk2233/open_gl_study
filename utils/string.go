/*
@author: sk
@date: 2023/5/8
*/
package utils

import "strconv"

func ToInt(val string) int {
	res, err := strconv.Atoi(val)
	HandleErr(err)
	return res
}

func ToFloat(val string) float32 {
	res, err := strconv.ParseFloat(val, 32)
	HandleErr(err)
	return float32(res)
}
