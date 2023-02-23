package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert_readline(readline string) string {

	switch readline {

	case "LDA_im":
		readline = strings.Trim(readline, "LDA_im")
		readline = fmt.Sprint(readline, 0xA1)

	case "LDA_ab":
		readline = strings.Trim(readline, "LDA_ab")
		readline = fmt.Sprint(readline, 0xA2)

	case "LDA_abR1":
		readline = strings.Trim(readline, "LDA_abR1")
		readline = fmt.Sprint(readline, 0xA3)

	case "LDA_abR2":
		readline = strings.Trim(readline, "LDA_abR2")
		readline = fmt.Sprint(readline, 0xA4)

	case "LDr1_im":
		readline = strings.Trim(readline, "LDr1_im")
		readline = fmt.Sprint(readline, 0xB1)

	case "LDr1_ab":
		readline = strings.Trim(readline, "LDr1_ab")
		readline = fmt.Sprint(readline, 0xB2)

	case "LDr1_abR2":
		readline = strings.Trim(readline, "LDr1_abR2")
		readline = fmt.Sprint(readline, 0xB3)

	case "LDr2_im":
		readline = strings.Trim(readline, "LDr2_im")
		readline = fmt.Sprint(readline, 0xC1)

	case "LDr2_ab":
		readline = strings.Trim(readline, "LDr2_ab")
		readline = fmt.Sprint(readline, 0xC2)

	case "LDr2_abR1":
		readline = strings.Trim(readline, "LDr2_abR1")
		readline = fmt.Sprint(readline, 0xC3)

	case "sta_ab":
		readline = strings.Trim(readline, "sta_ab")
		readline = fmt.Sprint(readline, 0xD1)

	case "sta_abR1":
		readline = strings.Trim(readline, "sta_abR1")
		readline = fmt.Sprint(readline, 0xD2)

	case "sta_abR2":
		readline = strings.Trim(readline, "sta_abR2")
		readline = fmt.Sprint(readline, 0xD3)

	case "str1_ab":
		readline = strings.Trim(readline, "str1_ab")
		readline = fmt.Sprint(readline, 0xE1)

	case "str1_abR2":
		readline = strings.Trim(readline, "str1_abR2")
		readline = fmt.Sprint(readline, 0xE2)

	case "str2_ab":
		readline = strings.Trim(readline, "str2_ab")
		readline = fmt.Sprint(readline, 0xF1)

	case "str2_abR1":
		readline = strings.Trim(readline, "str2_abR1")
		readline = fmt.Sprint(readline, 0xF2)

	case "STO_ab":
		readline = strings.Trim(readline, "STO_ab")
		readline = fmt.Sprint(readline, 0xF3)

	case "JMP_ab":
		readline = strings.Trim(readline, "JMP_ab")
		readline = fmt.Sprint(readline, 0x4C)

	case "JE_ab":
		readline = strings.Trim(readline, "JE_ab")
		readline = fmt.Sprint(readline, 0x6F)

	case "CMP_ab":
		readline = strings.Trim(readline, "CMP_ab")
		readline = fmt.Sprint(readline, 0x6E)

	case "ADD_im":
		readline = strings.Trim(readline, "ADD_im")
		readline = fmt.Sprint(readline, 0x5C)

	case "ADD_ab":
		readline = strings.Trim(readline, "ADD_ab")
		readline = fmt.Sprint(readline, 0x5D)

	case "SUB_im":
		readline = strings.Trim(readline, "SUB_im")
		readline = fmt.Sprint(readline, 0x6C)

	case "SUB_ab":
		readline = strings.Trim(readline, "SUB_ab")
		readline = fmt.Sprint(readline, 0x6D)

	case "der1":
		readline = strings.Trim(readline, "der1")
		readline = fmt.Sprint(readline, 0x7C)

	case "der2":
		readline = strings.Trim(readline, "der2")
		readline = fmt.Sprint(readline, 0x7D)

	case "dea":
		readline = strings.Trim(readline, "dea")
		readline = fmt.Sprint(readline, 0x7E)

	case "inr1":
		readline = strings.Trim(readline, "inr1")
		readline = fmt.Sprint(readline, 0x7F)

	case "inr2":
		readline = strings.Trim(readline, "inr2")
		readline = fmt.Sprint(readline, 0x7A)

	case "ina":
		readline = strings.Trim(readline, "ina")
		readline = fmt.Sprint(readline, 0x7B)

	case "wrt_im":
		readline = strings.Trim(readline, "wrt_im")
		readline = fmt.Sprint(readline, 0x8B)

	case "wrt_ab":
		readline = strings.Trim(readline, "wrt_ab")
		readline = fmt.Sprint(readline, 0x8C)

	case "hlt":
		readline = strings.Trim(readline, "hlt")
		readline = fmt.Sprint(readline, 0x7A)

	case "LDCP_im":
		readline = strings.Trim(readline, "LDCP_im")
		readline = fmt.Sprint(readline, 0xA5)

	case "LDCP_ab":
		readline = strings.Trim(readline, "LDCP_ab")
		readline = fmt.Sprint(readline, 0xA6)

	}
	if strings.Contains(readline, "'") {
		//readline = strings.Trim(readline, "'A'")
		//readline = fmt.Sprint(readline, 65)
		readline = strings.Replace(readline, "'", "", -1)
		char := readline[0]
		ascii := uint8(char)
		//println(ascii)
		readline = strconv.FormatUint(uint64(ascii), 10)
		//readline = string(ascii)

		//readline = strings.Replace(readline, string(char), string(ascii), -1)

	}
	//println(readline)

	return readline
}
