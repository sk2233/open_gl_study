/*
@author: sk
@date: 2023/5/27
*/
package frame3

import (
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Music struct {
	player oto.Player
}

func (m *Music) initData(path string) {
	file := OpenFile(path)
	defer file.Close()
	decoder, err := mp3.NewDecoder(file)
	HandleErr(err)
	ctx, ready, err := oto.NewContext(decoder.SampleRate(), 2, 2)
	HandleErr(err)
	<-ready
	m.player = ctx.NewPlayer(decoder)
}

func (m *Music) Play() {
	m.player.Play()
}

func NewMusic(path string) *Music {
	res := &Music{}
	res.initData(path)
	return res
}
