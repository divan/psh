package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		Usage()
	}

	var args []string
	args = Transliterate(os.Args[1:])

	cmd := exec.Command("go", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()
}

func Transliterate(args []string) []string {
	var ret []string
	for _, arg := range args {
		tr := translit(arg)
		ret = append(ret, tr)
	}
	return ret
}

func translit(word string) string {
	var ret string
	for _, letter := range word {
		ret += table[letter]
	}
	return ret
}

func Usage() {
	fmt.Println(`Usage: пщ буилд -> go build

Пше гіт версіон 0.2

Используйте транслитерированные команды, например:
 - пщ гет гитхаб.ком/диван/гофреш
 - пщ буилд
 - пщ инсталл
 - пщ тест`)

	os.Exit(0)
}

var table = map[rune]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'е': "e",
	'ë': "e",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'u': "u",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ь': "'",
	'ы': "y",
	'ъ': "'",
	'э': "e",
	'ю': "yu",
	'я': "ya",
	'ї': "yi",
	'є': "e",
	'і': "i",
}
