package const

const pi = 3.141
const (
	hello = "Привет"
	e 	 = 2.718
	)
	const (
		zero = iota
		_  		// пустаяя переменная, пропуск iota
		three   // = 3
	)
	const (
		_ 			= iota 				// пропускаем первое значение
		KB uint64 = 1 << (10 * iota)	// 1024
		MB								// 1048576
	)