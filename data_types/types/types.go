package types

type UserID int 

func main() {
	idx := 1
	var uid UserID = 42

	// Даже если базовый тип одинаковый, разные типы несовместимы
	// cannot use uid (type UserID) as type int in assignment
	// myID := idx

	myID := UserID(idx)

	println(uid, myID)
}