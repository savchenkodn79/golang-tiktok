package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var tiktok = [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func paint() {
	fmt.Println("|===|===|===|")
	fmt.Printf("| %v | %v | %v |\n", tiktok[0][0], tiktok[0][1], tiktok[0][2])
	fmt.Println("|===|===|===|")
	fmt.Printf("| %v | %v | %v |\n", tiktok[1][0], tiktok[1][1], tiktok[1][2])
	fmt.Println("|===|===|===|")
	fmt.Printf("| %v | %v | %v |\n", tiktok[2][0], tiktok[2][1], tiktok[2][2])
	fmt.Println("|===|===|===|")
}

func gameOver() bool {
	gameover := true
	for i := 0; i < 3; i++ {
		for a := 0; a < 3; a++ {
			if tiktok[i][a] == "1" ||
				tiktok[i][a] == "2" ||
				tiktok[i][a] == "3" ||
				tiktok[i][a] == "4" ||
				tiktok[i][a] == "5" ||
				tiktok[i][a] == "6" ||
				tiktok[i][a] == "7" ||
				tiktok[i][a] == "8" ||
				tiktok[i][a] == "9" {
				gameover = false
			}
		}
	}
	return gameover
}

func areYouWin() string {
	if (tiktok[0][0] == tiktok[0][1]) && (tiktok[0][1] == tiktok[0][2]) && (tiktok[0][0] == tiktok[0][2]) {
		return tiktok[0][0]
	}
	if (tiktok[1][0] == tiktok[1][1]) && (tiktok[1][1] == tiktok[1][2]) && (tiktok[1][0] == tiktok[1][2]) {
		return tiktok[1][0]
	}
	if (tiktok[2][0] == tiktok[2][1]) && (tiktok[2][1] == tiktok[2][2]) && (tiktok[2][0] == tiktok[2][2]) {
		return tiktok[2][0]
	}
	if (tiktok[0][0] == tiktok[1][0]) && (tiktok[1][0] == tiktok[2][0]) && (tiktok[0][0] == tiktok[2][0]) {
		return tiktok[0][0]
	}
	if (tiktok[0][1] == tiktok[1][1]) && (tiktok[1][1] == tiktok[2][1]) && (tiktok[0][1] == tiktok[2][1]) {
		return tiktok[0][1]
	}
	if (tiktok[0][2] == tiktok[1][2]) && (tiktok[1][2] == tiktok[2][2]) && (tiktok[0][2] == tiktok[2][2]) {
		return tiktok[0][2]
	}
	if (tiktok[0][0] == tiktok[1][1]) && (tiktok[1][1] == tiktok[2][2]) && (tiktok[0][0] == tiktok[2][2]) {
		return tiktok[0][0]
	}
	if (tiktok[0][2] == tiktok[1][1]) && (tiktok[1][1] == tiktok[2][0]) && (tiktok[0][2] == tiktok[2][0]) {
		return tiktok[0][2]
	}

	return ""
}

func main() {

	clear()
	paint()

	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	step := 1
	for scanner.Scan() {

		cmd := scanner.Text()
		switch cmd {
		case "exit":
			os.Exit(0)
		default:
			mod := step % 2
			num, err := strconv.Atoi(cmd)
			if err != nil {
				fmt.Println("Please put number")
			}
			switch num {
			case 1:
				if mod == 1 {
					tiktok[0][0] = "X"
				} else {
					tiktok[0][0] = "O"
				}
			case 2:
				if mod == 1 {
					tiktok[0][1] = "X"
				} else {
					tiktok[0][1] = "O"
				}
			case 3:
				if mod == 1 {
					tiktok[0][2] = "X"
				} else {
					tiktok[0][2] = "O"
				}
			case 4:
				if mod == 1 {
					tiktok[1][0] = "X"
				} else {
					tiktok[1][0] = "O"
				}
			case 5:
				if mod == 1 {
					tiktok[1][1] = "X"
				} else {
					tiktok[1][1] = "O"
				}
			case 6:
				if mod == 1 {
					tiktok[1][2] = "X"
				} else {
					tiktok[1][2] = "O"
				}
			case 7:
				if mod == 1 {
					tiktok[2][0] = "X"
				} else {
					tiktok[2][0] = "O"
				}
			case 8:
				if mod == 1 {
					tiktok[2][1] = "X"
				} else {
					tiktok[2][1] = "O"
				}
			case 9:
				if mod == 1 {
					tiktok[2][2] = "X"
				} else {
					tiktok[2][2] = "O"
				}
			}
			clear()
			paint()
			step++
			winer := areYouWin()
			if winer != "" {
				fmt.Println("We have a winer! Congratilations", winer)
				os.Exit(0)
			}
			if gameOver() {
				fmt.Println("Game Over")
				os.Exit(0)
			}

		}
	}
}
