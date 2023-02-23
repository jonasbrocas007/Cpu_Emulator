package main

var translated_character string

func ascii_translation(character uint8) {

	translated_character = string(rune(character))

	return
}

func write_character(translated_character string) {
	print(translated_character)
	return
}
