package main

type UserID float64

func main() {
	idx := 1
	var uid UserID = 42

	// даже если базовый тип одинаковый, разные типы несовместимы
	// cannot use uid (type UserID) as type int64 in assignment
	// myID := idx

	myID := UserID(idx)

	println(uid, myID)
}