package day4_24

import (
	"io"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	data := GetData()
	xmas_counter := 0

	for i, arr := range data {
		for j := 0; j < len(arr); j++ {
			if arr[j] == 'X' {
				//UP
				if i-3 >= 0 {
					if data[i-1][j] == 'M' && data[i-2][j] == 'A' && data[i-3][j] == 'S' {
						xmas_counter++
					}
					//LEFT
					if j-3 >= 0 {
						if data[i-1][j-1] == 'M' && data[i-2][j-2] == 'A' && data[i-3][j-3] == 'S' {
							xmas_counter++
						}
					}
					//RIGHT
					if j+3 < len(arr) {
						if data[i-1][j+1] == 'M' && data[i-2][j+2] == 'A' && data[i-3][j+3] == 'S' {
							xmas_counter++
						}

					}
				}

				//DOWN
				if i+3 < len(data) {
					if data[i+1][j] == 'M' && data[i+2][j] == 'A' && data[i+3][j] == 'S' {
						xmas_counter++
					}
					//LEFT
					if j-3 >= 0 {
						if data[i+1][j-1] == 'M' && data[i+2][j-2] == 'A' && data[i+3][j-3] == 'S' {
							xmas_counter++
						}
					}
					//RIGHT
					if j+3 < len(arr) {
						if data[i+1][j+1] == 'M' && data[i+2][j+2] == 'A' && data[i+3][j+3] == 'S' {
							xmas_counter++
						}
					}
				}

				//LEFT
				if j-3 >= 0 {
					if arr[j-1] == 'M' && arr[j-2] == 'A' && arr[j-3] == 'S' {
						xmas_counter++
					}
				}

				//RIGHT
				if j+3 < len(arr) {
					if arr[j+1] == 'M' && arr[j+2] == 'A' && arr[j+3] == 'S' {
						xmas_counter++
					}
				}
			}
		}
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(xmas_counter))
	shared.ErrCheck(err)
	out.Sync()
}

func GetData() [][]byte {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	out := [][]byte{}
	arr := []byte{}

	for i < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		i++

		if string(char[0]) == "\n" {
			out = append(out, arr)
			arr = []byte{}
			continue
		}
		arr = append(arr, char[0])
	}

	return out
}
