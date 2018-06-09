package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/gmaclinuxer/goplayer/menu"
	"github.com/gmaclinuxer/goplayer/player"
	"strconv"
)

var id int
var m *menu.MusicManager

// handlePlayCmd play/stop/pause control
func handlePlayCmd(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}

	pm := m.Find(tokens[1])

	if pm == nil {
		fmt.Printf("music %s dose not exist!\n", tokens[1])
		return
	}

	player.Play(pm.Source, pm.Type)

}

// handleMenuCmd add/remove/play music
func handleMenuCmd(tokens []string) {
	switch tokens[1] {
	case "list":
		m.List()
	case "add":
		id++
		m.Add(&menu.MusicEntry{
			Id:     strconv.Itoa(id),
			Name:   tokens[2],
			Artist: tokens[3],
			Source: tokens[4],
			Type:   tokens[5],
		})
	case "remove":
		fi := m.FindIndex(tokens[1])
		if fi != -1 {
			rm := m.Remove(fi)
			fmt.Printf("%s removed success\n", rm)
		} else {
			fmt.Printf("%s does not exist\n", tokens[1])
		}
	default:
		fmt.Println("operation not support")
	}
}

func main() {
	fmt.Println(`Enter following commands to control player:
menu list -- view all musics in menu
menu add <name><artist><source><type> -- add a music to the menu
menu remove <name> -- remove a music from menu
play <name> -- play specified music in menu
`)

	m = menu.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter command -> ")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "menu" {
			handleMenuCmd(tokens)
		} else if tokens[0] == "play" {
			handlePlayCmd(tokens)
		} else {
			fmt.Println("Unrecognized command: ", tokens[0])
		}

	}
}
