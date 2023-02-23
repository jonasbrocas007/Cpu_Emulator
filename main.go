package main

import (
	"fmt"
	"os"
)

//The bad processor emulator

var PC, R1, R2, ACC, OR uint8
var mem_data [255]uint8
var Z, C, I, COMP uint8 //flag registers
var data, value, cycles uint8

// instruction data
// lda
var LDA_im uint8 = 0xA1
var LDA_ab uint8 = 0xA2
var LDA_abR1 uint8 = 0xA3
var LDA_abR2 uint8 = 0xA4
var LDCP_im uint8 = 0xA5
var LDCP_ab uint8 = 0xA6

// ldr1
var LDr1_im uint8 = 0xB1
var LDr1_ab uint8 = 0xB2
var LDr1_abR2 uint8 = 0xB3

// ldr2
var LDr2_im uint8 = 0xC1
var LDr2_ab uint8 = 0xC2
var LDr2_abR2 uint8 = 0xC3

// sta
var sta_ab uint8 = 0xD1
var sta_abR1 uint8 = 0xD2
var sta_abR2 uint8 = 0xD3

// str1
var str1_ab uint8 = 0xE1
var str1_abR2 uint8 = 0xE2

// str2
var str2_ab uint8 = 0xF1
var str2_abR1 uint8 = 0xF2

var STO_ab uint8 = 0xF3

var JMP_ab uint8 = 0x4C
var JE_ab uint8 = 0x6F
var CMP_ab uint8 = 0x6E

var ADD_im uint8 = 0x5C
var ADD_ab uint8 = 0x5D
var SUB_im uint8 = 0x6C
var SUB_ab uint8 = 0x6D

var der1 uint8 = 0x7C
var der2 uint8 = 0x7D
var dea uint8 = 0x7E
var inr1 uint8 = 0x7F
var inr2 uint8 = 0x7A
var ina uint8 = 0x7B

var writec_im uint8 = 0x8B
var writec_ab uint8 = 0x8C

var hlt uint8 = 0x7A
var current_adress uint8 = 0

func CPU_reset(PC uint8, R1 uint8, R2 uint8, mem_data [255]uint8) {
	PC = 0x00
	R1 = 0x00
	R2 = 0x00
	Z = 0x00
	C = 0x00
	I = 0x00
	for i := uint8(0); i < 255; i++ {
		mem_data[i] = 0
	}
}

func fetchbyte(mem_data [255]uint8) {
	data = mem_data[PC]
	PC++
	cycles--
	return
}

func CPU_execute(cycles uint8) {
	for cycles > 0 {
		fetchbyte(mem_data)
		//for i := 0; i < len(mem_data); i++ {
		//println(mem_data[i])
		//}

		switch data {

		case uint8(LDA_im): //load acc immediate
			fetchbyte(mem_data)
			ACC = data

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}

			break

		case uint8(LDA_ab): //load acc absolute

			fetchbyte(mem_data)

			value = data
			cycles--

			ACC = mem_data[value]
			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(LDA_abR1): //load acc from r1
			fetchbyte(mem_data)

			value = data
			cycles--

			ACC = mem_data[value] + R1
			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(LDA_abR2): //load acc from r2
			fetchbyte(mem_data)

			value = data
			cycles--

			ACC = value + R2
			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		///////////////////////////////////////////////////
		case uint8(LDr1_im): //load R1 immediate
			fetchbyte(mem_data)
			R1 = data

			if R1 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(LDr1_ab): //load R1 absolute
			fetchbyte(mem_data)

			value = data
			cycles--
			R1 = mem_data[value]
			cycles--

			if R1 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(LDr1_abR2): //load R1 from r2
			fetchbyte(mem_data)

			value = data
			cycles--

			R1 = mem_data[value] + R2
			cycles--

			if R1 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		///////////////////////////////////////////////////
		case uint8(LDr2_im): //load R2 immediate
			fetchbyte(mem_data)
			R2 = data

			if R2 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(LDr2_ab): //load R1 absolute
			fetchbyte(mem_data)

			value = data
			cycles--

			R2 = mem_data[value]
			cycles--

			if R2 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(LDr2_abR2): //load R2 from R1
			fetchbyte(mem_data)

			value = data
			cycles--

			R2 = mem_data[value] + R1
			cycles--

			if R2 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(sta_ab): //load acc absolute
			fetchbyte(mem_data)

			value = data
			cycles--

			mem_data[value] = ACC
			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(sta_abR1): //load acc from r1
			fetchbyte(mem_data)

			value = data
			cycles--

			mem_data[value] = ACC + R1
			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(sta_abR2): //load acc from r2
			fetchbyte(mem_data)
			if data == 0x00 {
				break
			}

			value = data
			cycles--

			mem_data[value] = ACC + R2
			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(str1_ab): //load R1 absolute
			fetchbyte(mem_data)

			value = data
			cycles--

			if value == 0 {
				PC--
				break

			} else {
				mem_data[value] = R1
			}
			cycles--

			if R1 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(str1_abR2): //store R1 from r2
			fetchbyte(mem_data)

			value = data
			cycles--

			if value == 0 {
				PC--
				break

			} else {
				mem_data[value] = R1 + R2
			}

			cycles--

			if R1 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(str2_ab): //load R2 absolute
			fetchbyte(mem_data)

			value = data
			cycles--
			if value == 0 {
				PC--
				break

			} else {
				mem_data[value] = R2
			}

			cycles--

			if R2 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(str2_abR1): //load R2 from r1
			fetchbyte(mem_data)

			value = data
			cycles--
			if value == 0 {
				PC--
				break

			} else {
				mem_data[value] = R2 + R1
			}

			cycles--

			if R2 == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(JMP_ab):
			fetchbyte(mem_data)
			PC = data
			current_adress = data
			cycles--
			break

		case uint8(LDCP_im):
			fetchbyte(mem_data)
			COMP = data

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}

			break

		case uint8(LDCP_ab):
			fetchbyte(mem_data)

			value = data
			cycles--

			COMP = mem_data[value]

			cycles--

			if ACC == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(CMP_ab):
			fetchbyte(mem_data)

			adress0 := mem_data[data]
			cycles--

			adress0 = mem_data[adress0]

			cycles--

			if adress0 == COMP {
				I = 1
			} else {
				I = 0
			}

			cycles--
			break

		case uint8(JE_ab):
			if I == 1 {
				fetchbyte(mem_data)

				PC = data
				current_adress = data
				cycles--
			}
			break

		case uint8(ADD_im): //add adds the memory and ACC
			fetchbyte(mem_data)

			mem_add := data
			cycles--

			acc_add := ACC
			OR = uint8(mem_add + acc_add)

			cycles--

			break

		case uint8(ADD_ab):
			fetchbyte(mem_data)
			adress := data
			cycles--

			mem_add := mem_data[adress]
			cycles--

			acc_add := ACC
			OR = uint8(mem_add + acc_add)
			cycles--

			break

		case uint8(SUB_im): //sub subtracts the ACC from memory
			fetchbyte(mem_data)
			mem_add := data
			cycles--

			acc_add := ACC
			OR = uint8(mem_add - acc_add)
			cycles--
			break

		case uint8(SUB_ab):
			fetchbyte(mem_data)
			adress := data
			cycles--

			mem_add := mem_data[adress]
			cycles--

			acc_add := ACC
			OR = uint8(mem_add - acc_add)
			cycles--
			break
		case uint8(hlt):
			PC = 0

		case uint8(STO_ab): //operation result
			fetchbyte(mem_data)

			value = data
			cycles--

			if value == 0 {
				PC--
				break

			} else {
				mem_data[value] = OR
			}
			cycles--

			if OR == 0 {
				Z = 1 //sets the zero flag to true or false
			}
			break

		case uint8(der1):
			fetchbyte(mem_data)
			R1--

		case uint8(der2):
			fetchbyte(mem_data)
			R2--

		case uint8(dea):
			fetchbyte(mem_data)
			ACC--

		case uint8(inr1):
			fetchbyte(mem_data)
			R1++

		case uint8(inr2):
			fetchbyte(mem_data)
			R2++

		case uint8(ina):
			fetchbyte(mem_data)
			ACC++

		case uint8(writec_im):
			fetchbyte(mem_data)

			character := data

			ascii_translation(character)
			cycles--
			write_character(translated_character)

		case uint8(writec_ab):
			fetchbyte(mem_data)
			value = data
			character := mem_data[value]
			//println(mem_data[3])
			ascii_translation(character)
			cycles--
			write_character(translated_character)

		}
		break
	}
	return

}

func main() {

	current_data := 0
	var memdata_table []uint8
	dir := "D:/programacao2/"
	file, err := os.Open(dir + "Cpu_emulator/binaries/example.bin")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := make([]byte, 1024)

	n, err := file.Read(data)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	data = data[:n]

	for current_data < len(data) {

		memdata_table = append(memdata_table, uint8(data[current_data]))
		current_data++

	}

	for current_adress < uint8(len(memdata_table)) {

		mem_data[current_adress] = memdata_table[current_adress]
		current_adress++

		if current_adress != 0 {
			if current_adress%2 == 0 {
				CPU_execute(4)
			}
		}

	}

	CPU_reset(PC, R1, R2, mem_data)
}
