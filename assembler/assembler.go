package main

import (
	"bufio"
	"encoding/binary"
	"os"
	"strconv"
	"strings"
)

var read []string
var data []uint8

func simplify_functions() {

	file, err := os.Open("D:/programacao2/Cpu_emulator/code.asm")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		read = strings.Split(line, ",")

		for line := 0; line < len(read); line++ {
			readline := read[line]
			//readline = strings.Trim(readline, " ") //RESOLVER DEPOIS
			readline = strings.Trim(readline, "\n")

			//if string(readline[0]) == " " {
			readline = strings.Trim(readline, " ")
			//}

			readline = convert_readline(readline)
			//println(readline)
			data_slice, _ := strconv.ParseUint(readline, 10, 8)

			data = append(data, uint8(data_slice))
			println(data_slice)

		}
	}
}
func main() {
	simplify_functions()

	dir := "D:/programacao2/"
	file, err := os.Create(dir + "Cpu_emulator/binaries/example.bin")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	//var data = []uint8{165, 1, 110, 1}

	for _, data_slice := range data {
		err := binary.Write(file, binary.LittleEndian, data_slice)
		if err != nil {
			panic(err)
		}
	}

}

// sta

// str1

// str2
