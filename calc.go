package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		a    int //первое число
		b    int //второе число
		res  int //результат в арабской СС
		oper string
		rim  [10]string = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
		fl1  int        = 0
		fl2  int        = 0
	)

	m := map[string]int{
		"I": 1, "1": 1,
		"II": 2, "2": 2,
		"III": 3, "3": 3,
		"IV": 4, "4": 4,
		"V": 5, "5": 5,
		"VI": 6, "6": 6,
		"VII": 7, "7": 7,
		"VIII": 8, "8": 8,
		"IX": 9, "9": 9,
		"X": 10, "10": 10,
	} //перевод в арабскую СС

	r := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	} //перевод в римскую СС

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	M := strings.Fields(input)

	if len(M) != 3 {
		fmt.Println("ПАНИКА: формат математической операции не удовлетворяет заданию")
		os.Exit(0)
	}

	oper = M[1]
	a = m[M[0]]
	b = m[M[2]]

	if oper == "+" {
		res = a + b
	} else if oper == "-" {
		res = a - b
	} else if oper == "*" {
		res = a * b
	} else if oper == "/" {
		res = a / b
	} else {
		fmt.Println("ОШИБКА: я не знаю такой операции")
	}

	for i := 0; i < 10; i++ {
		if M[0] == rim[i] {
			fl1 = 1
		}
		if M[2] == rim[i] {
			fl2 = 1
		}
	}

	if fl1 == 0 && fl2 == 0 {
		fmt.Println(res) //выводим арабские числа
	} else if fl1 == 1 && fl2 == 1 {
		if res <= 0 {
			fmt.Println("ПАНИКА: в римской системе нет отрицательных чисел и нуля.")
		} else {
			fmt.Println(r[res])
		}
	} else {
		fmt.Println("ПАНИКА: используются одновременно разные системы счисления.")
	}

}
