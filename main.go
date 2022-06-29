package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	fixed := commonFix(word)
	var ret string
	for _, letter := range fixed {
		r := table[letter]
		if r == "" {
			r = string(letter)
		}
		ret += r
	}
	return ret
}

func Usage() {
	fmt.Println(`Usage: пщ білд -> go build

Пщ версіон 1.0

Використовуйте транслітеровані команди, на кшталт:
 - пщ гет гітхаб.ком/діван/гофреш
 - пщ білд
 - пщ інсталл
 - пщ тест`)

	os.Exit(0)
}

var table = map[rune]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
        'ґ': "g",
	'д': "d",
	'е': "e",
	'є': "e",
	'ж': "zh",
	'з': "z",
	'и': "y",
	'і': "i",
	'ї': "yi",
	'й': "y",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ь': "'",
	'ю': "yu",
	'я': "ya",
}

func commonFix(word string) string {
	ret := strings.Replace(word, "гітхаб.ком", "github.com", -1)
	ret = strings.Replace(ret, "білд", "build", -1)
	return ret
}
