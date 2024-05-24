/*
@author: sk
@date: 2023/5/27
*/
package main

import (
	"math/rand"
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

var (
	BrickData  [][]int
	BrickColor []mgl32.Vec4
)

func Init() {
	// BrickData
	rand.Seed(time.Now().UnixMilli())
	BrickData = make([][]int, 16)
	for i := 0; i < 16; i++ {
		BrickData[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			// <0 铁墙    0  空    其他 砖块
			BrickData[i][j] = rand.Intn(9) - 1
		}
	}
	// BrickColor  7个即可
	BrickColor = []mgl32.Vec4{{1, 0, 0, 1}, {0, 1, 0, 1}, {0, 0, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 1}}
}
