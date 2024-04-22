package main

import (
	//"fmt"
	"os"
	"io/ioutil"
	"strconv"
	"strings"
)

func index1(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+1:]...)
}

func index2(s []string, index int) []string {
	if index < 0 || index >= len(s) {
		return s
	}
	return append(s[:index], s[index+2:]...)
}

func UpFunc(s []string) []string {
	for i, val := range s {
		if val == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = index1(s, i)
		}
	}
	return s
}

func LowFunc(s []string) []string {
	for i, val := range s {
		if val == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = index1(s, i)
		}
	}
	return s
}

func CapFunc(s []string) []string {
	for i, val := range s {
		if val == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = index1(s, i)
		}
	}
	return s
}

func UpCapAndLowFunc(s []string) []string {
	for idx, val := range s {
		if val == "(up," && idx < len(s)-1 {
			valn := strings.TrimSuffix(s[idx+1], ")")
			number, _ := strconv.Atoi(valn)
			for i := 1; i <= number; i++ {
				a := (idx - i)
				if a >= 0 && a < len(s) {
					s[a] = strings.ToUpper(s[a])
				}
			}
			s = index2(s, idx)
		} else if val == "(low," {
			valn := strings.TrimSuffix(s[idx+1], ")")
			number, _ := strconv.Atoi(string(valn))
			for i := 1; i <= number; i++ {
				a := (idx - i)
				if a >= 0 && a < len(s) {
					s[a] = strings.ToLower(s[a])
				}
			}
			s = index2(s, idx)
		} else if val == "(cap," {
			valn := strings.TrimSuffix(s[idx+1], ")")
			number, _ := strconv.Atoi(string(valn))
			for i := 1; i <= number; i++ {
				a := (idx - i)
				if a >= 0 && a < len(s) {
					s[a] = strings.Title(s[a])
				}
			}
			s = index2(s, idx)
		}
	}
	return s
}

func HexDec(s []string) []string {
	for i, v := range s {
		if v == "(hex)" {
			decVal, _ := strconv.ParseInt(s[i-1], 16, 64)
			s[i-1] = strconv.Itoa(int(decVal))
			s = index1(s, i)
		}
	}
	return s
}

func BinDec(s []string) []string {
	for i, v := range s {
		if v == "(bin)" {
			decVal, _ := strconv.ParseInt(s[i-1], 2, 64)
			s[i-1] = strconv.Itoa(int(decVal))
			s = index1(s, i)
		}
	}
	return s
}

func Ac(s []string) []string { 
	for i, val := range s {
		if val == "a" && i > 0 && i+1 < len(s) {
			nw := s[i+1] // next word
			vw := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
			for _, char := range vw {
				if strings.HasPrefix(nw, char) { 
					s[i] = val + "n"
				}
			}
		}
	}
	return s
}

func tas7i7Func(s []string) []string { // ts7i7 kalima wa no9at
	vals := []string{".", ",", "!", "?", ":", ";"}
	for idx, val := range s {
		for _, char := range vals {
			if val == char {
				s[idx-1] = s[idx-1] + char
				s = index1(s, idx)
			} else if strings.HasPrefix(val, char) { // check is first char in val == char 
 				s[idx-1] = s[idx-1] + char
				s[idx] = s[idx][1:]
			}
		}
	}
	return s
}

func AllChar(s []string) []string {
	vals := []string{",", "!", "?", ":", ";", "."}
	for i, val := range s {
		for _, v := range vals {
			if len(val) > 1 && strings.HasPrefix(val, v) && strings.HasSuffix(val, v) { //  check first and last char in string
				s[i-1] = s[i-1] + val
				s = index1(s, i)
			}
		}
	}
	return tas7i7Func(s)
}

func SingelCot(s []string) []string {
	c := 0
	for i, v := range s {
		if v == "'" && c == 0 {
			s[i+1] = v + s[i+1]
			s = append(s[:i], s[i+1:]...)
			c++
		}
	}
	for i, w := range s {
		if w == "'" {
			s[i-1] += w
			c = 0
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	// get args from user
	file_input := os.Args[1]
	file_output := os.Args[2]
	// open file and read
	srtFile, err := ioutil.ReadFile(file_input)
	if err != nil {
		return
	}
	// splite spice
	newStr := strings.Split(string(srtFile), " ")
	// call func 
	newStr = UpCapAndLowFunc(newStr)
	newStr = UpFunc(newStr)
	newStr = LowFunc(newStr)
	newStr = CapFunc(newStr)
	newStr = tas7i7Func(newStr)
	newStr = AllChar(newStr)
	newStr = HexDec(newStr)
	newStr = BinDec(newStr)
	newStr = Ac(newStr)
	newStr = SingelCot(newStr)

	var strOut string // var of stoce out put
	strOut = strings.Join(newStr, " ") + "\n" // add spice 

	// create file out put
	file, err := os.Create(file_output)
	if err != nil {
		return
	}
	defer file.Close() // close file 
	// write in file out put new string
	_, err = file.WriteString(strOut)
	if err != nil {
		return
	}
}