package newArr

import (
	"reflect"
	"testing"
)

func TestMakeArr(t *testing.T) {
	t.Run("количество элементов > 0", func(t *testing.T) {
		amountOfElements := 5

		want := []int {0,0,0,0,0}
		got := MakeArr(amountOfElements)

		if !reflect.DeepEqual(want, got) {
			t.Error("неправильно создает массив где кол-во элементов больше 0")
		}
	})

	t.Run("количество элементов 0", func(t *testing.T) {
		amountOfElements := 0

		want := make([]int, 0)
		got := MakeArr(amountOfElements)

		if !reflect.DeepEqual(want, got) {
			t.Error("неправильно создает массив где кол-во элементов 0")
		}
	})

	t.Run("количество элементов < 0", func(t *testing.T) {
		amountOfElements := -100

		want := make([]int, 0)
		got := MakeArr(amountOfElements)

		if !reflect.DeepEqual(want, got) {
			t.Error("неправильно создает массив где кол-во элементов < 0")
		}
	})
}

func TestFillArr(t *testing.T) {
	Arr := MakeArr(5)

	want := []int {1,2,3,4,5}
	got := FillArr(Arr)

	if !reflect.DeepEqual(want, got) {
		t.Error("неправильно заполняет массив")
	}
}