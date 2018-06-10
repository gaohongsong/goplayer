package player

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "mp3":
		// assign a struct inst to a interface by pointer
		p = &MP3Player{}
	case "wav":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unspported music type ", mtype)
		return
	}

	p.Play(source)
}