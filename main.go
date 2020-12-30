package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type cursor struct {
	x int
	y int
	defx int
	defy int
	gmap []string
	win bool
}

func (c *cursor) move(x int, y int) {

	next := c.gmap[c.y - y][c.x + x]
	fmt.Println(next)

	if next == 69 {
		c.win = true
		return
	}

	if next != 35 {
		c.x += x
		c.y -= y
	}

}

func (c *cursor) reset() {
	c.x = c.defx
	c.y = c.defy
}

func (c cursor) showtile() string {

	return string(c.gmap[c.y][c.x])

}

func (c cursor) showmap() {

	for i, a := range c.gmap {

		for j, b := range a {

			if i == c.y && j == c.x {
				fmt.Printf("O")
			} else {
				fmt.Printf(string(b))
			}

		}
		fmt.Println()

	}

}

func newCursor(gmap []string) cursor {
	
	/* map reverse
    for i, j := 0, len(gmap)-1; i < j; i, j = i+1, j-1 {
        gmap[i], gmap[j] = gmap[j], gmap[i]
    }
	*/

	for i, a := range gmap {
		for j, b := range a {
			if b == 105 {
				return cursor{j, i, j, i, gmap, false}
			}
		}
	}

	return cursor{0, len(gmap)-1, 0, len(gmap)-1, gmap, false}
}

func clear() {
	fmt.Printf("\u001b[2J\u001b[0;0H")
}

func play(gmap []string) {

	player := newCursor(gmap)

	player.showmap()

	var choice string

	exit := false

	noclear := false

	for {

		fmt.Printf("> ")

		fmt.Scanln(&choice)

		for _, key := range choice {

			switch (string(key)) {
			case "w", "k":
				player.move(0, 1)

			case "s", "j":
				player.move(0, -1)

			case "a", "h":
				player.move(-1, 0)

			case "d", "l":
				player.move(1, 0)

			case "r":
				player.reset()
				
			case "q":
				exit = true

			default:
				clear()
				fmt.Printf(`
-------------
You can move with wasd or with hjkl
you can chain commands this way:
wdwas

this is the equivalent of pressing:
w <enter> d <enter> w <enter> a <enter> s <enter>


Pressing only <enter> without anything will repeat your last key, so

s <enter> <enter> <enter>
is equilavent to s <enter> s <enter> s <enter>


Press 'r' to return to the start

Press 'q' to quit

Press any other key for this help message
-------------

`)
				noclear = true
			}

		}


		if player.win {
			clear()
			fmt.Printf("\nCongrats! You won!\n\n")
			break
		}

		if exit {
			clear()
			fmt.Println("Bye")
			break
		}

		if noclear {
			noclear = false
		} else {
			clear()
		}

		player.showmap()

	}
	
}

func loadMap(mapname string) []string {

	var err error

	var data []byte

	if strings.HasSuffix(mapname, ".gmap") {
		data, err = ioutil.ReadFile("./maps/" + mapname)
	} else {
		data, err = ioutil.ReadFile("./maps/" + mapname + ".gmap")
	}

    if err != nil {
		return []string{
			"Default map:",
			"Press x for help",
			"",
			"  ############",
			"  #          #",
			"  # ######## #",
			"  # #        #",
			"  #i# ########",
			"#####        ########",
			"#   ######## #      #",
			"# #          #E#### #",
			"# ################# #",
			"#                   #",
			"#####################",
		}

    }

	return strings.Split(string(data), "\n")

}



func main() {

	clear()

	var choice string

	exit := false

	for {

		fmt.Printf(`(Default is selected if no valid map is found)
'list' to list maps
'q' or 'quit' to quit
Once in game, press 'x' for help

Load map: `)

		fmt.Scanln(&choice)

		clear()

		switch (choice) {
		case "q", "quit":
			clear()
			fmt.Println("Bye")
			exit = true

		case "list":
			files, err := ioutil.ReadDir("./maps/")
			if err != nil {
				panic(err)
			}

			fmt.Println("Reading files in './maps/':")
			for _, f := range files {
				fmt.Println(f.Name())
			}
			fmt.Println()

		default:
			exit = true
			play(loadMap(choice))

		}

		if exit {
			break
		}

	}

}
