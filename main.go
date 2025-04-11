package main

import(
	"fmt"
	//"os"
)

var RAM [255]uint8
var PC uint8 //special registers
var R1, R2, R3 uint8 //general porpouse registers

var ADDI uint8 = 0xA1
var LOAD uint8 = 0xC6

func reset_cpu() (PC uint8, R1 uint8, R2 uint8, R3 uint8, RAM [255]uint8) {
	PC = 0x00
	R1 = 0x00
	R2 = 0x00
	R3 = 0x00
	for i := uint8(0); i < 255; i++ {
		RAM[i] = 0
	}
	return
}


func fetch_byte(PC uint8) uint8{
	data := RAM[PC]
	//fmt.Println(data)
	return data
}

func cpu_exec(PC uint8, R1 uint8, R2 uint8, R3 uint8, RAM [255]uint8, instruction_count uint8) (uint8, uint8, uint8, uint8, [255]uint8){
	for instruction_count > PC{
		iadata := fetch_byte(PC)
		var acc uint8 = 0
		switch iadata{
			case ADDI: //Add immediate instruction
				switch fetch_byte(PC + 2){
					case 0xFF:
						acc = R1
					case 0xFE:
						acc = R2
					case 0xFD:
						acc = R3
				}

				add := acc + fetch_byte(PC + 3)

				switch fetch_byte(PC + 1){
					case 0xFF:
						R1 = add
					case 0xFE:
						R2 = add
					case 0xFD:
						R3 = add
				}
				PC += 3
			
			case LOAD: //Load from register instruction
				switch fetch_byte(PC + 1){
					case 0xFF:
						R1 = RAM[fetch_byte(PC + 2)]
					case 0xFE:
						R2 = RAM[fetch_byte(PC + 2)]
					case 0xFD:
						R3 = RAM[fetch_byte(PC + 2)]
				}
				PC += 2
				
		}
		
	PC++

	}

	return PC, R1, R2, R3, RAM
}
func main(){
	RAM[0] = 0xC6 //Load
	RAM[1] = 0xFF // to R1
	RAM[2] = 3 // The third value
	RAM[3] = 5 // The third value is 5

	RAM[4] = 0xA1 // Add immediate
	RAM[5] = 0xFE // Add result to R1
	RAM[6] = 0xFF // And R1 to //////////////////// Bellow
	RAM[7] = 5 // Number 5     //////////////////// Add R1 to 5


	//cpu_exec(PC, ACC, R1, RAM, 2)
	PC, R1, R2, R3, RAM = cpu_exec(PC, R1, R2, R3,  RAM, 8)
	fmt.Println(R2)

}