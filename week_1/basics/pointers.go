package main

import "fmt"

func main() {
	a := 2
	a2 := 4
	b := &a
	b2 := &a2
	b = b2
	*b = 3 // a = 3
	c := &a // новый указатель на переменную a
	fmt.Println(b == c)

	// получение указателя на переменнут типа int
	// инициализировано значением по-умолчанию
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13 // c и a не изменились

	c = d   // теперь с указывает туда же, куда d
	*c = 14 // с = 14 -> d = 14, a = 12

}
