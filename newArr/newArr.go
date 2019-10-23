package newArr

func NewArr (amountOfElements int) []int {
	Arr := MakeArr(amountOfElements)
	Arr = FillArr(Arr)

	return Arr
}

func MakeArr (amountOfElements int) []int {
	//проверяем корректность кол-ва элементов
	if amountOfElements < 0 {
		amountOfElements = 0
	}

	Arr := make([]int, amountOfElements)

	return Arr
}

func FillArr (Arr []int) []int {
	for i := 0; i < len(Arr); i++ {
		Arr[i] = i+1
	}

	return Arr
}