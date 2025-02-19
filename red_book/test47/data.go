/*
@author: sk
@date: 2023/5/14
*/
package main

var (
	posVs = []float32{
		-1, -1, 0.0, 1.0,
		0, -1, 0.0, 1.0,
		-1, 0, 0.0, 1.0,
		0, 0, 0.0, 1.0,
	}
	colVs = []float32{
		1.0, 1.0, 1.0, 1.0,
		1.0, 1.0, 0.0, 1.0,
		1.0, 0.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 1.0,
	}
	offset = []float32{
		0, 0,
		0.5, 0.5,
		1, 1,
	}
	rectIndex = []uint32{
		0, 1, 2,
	}
)
