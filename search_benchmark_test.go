package binary

import (
	"testing"
)

func BenchmarkSearch(b *testing.B) {

	b.Run("Искомого числа нет в массиве", func(b *testing.B) {
		Arr := []int {2,4,6,8,10}
		desiredNum := 5

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Искомое число в середине.", func(b *testing.B) {
		Arr := []int {2,4,6,8,10}
		desiredNum := 6

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Искомое число крайнее.", func(b *testing.B) {
		Arr := []int {2,4,6,8,10}
		desiredNum := 10

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Массив не сортирован. Искомого числа нет", func(b *testing.B) {
		Arr := []int {2,6,4,8,10}
		desiredNum := 5

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Массив не сортирован. Искомое число в середине", func(b *testing.B) {
		Arr := []int {2,6,4,8,10}
		desiredNum := 6

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Массив не сортирован. Искомое число крайнее", func(b *testing.B) {
		Arr := []int {2,6,4,8,10}
		desiredNum := 10

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Массив из одного элемента", func(b *testing.B) {
		Arr := []int {6}
		desiredNum := 6

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Массив из отрицательных чисел", func(b *testing.B) {
		Arr := []int {-10,-8,-6,-4,-2}
		desiredNum := -10

		for i := 0; i < b.N; i++ {
			Search(desiredNum, Arr)
		}
	})

	b.Run("Большие массивы", func(b *testing.B) {

		b.Run("Массив на 100 элементов", func(b *testing.B) {
			Arr := makeArr(100)
			desiredNum := 50

			for i := 0; i < b.N; i++ {
				Search(desiredNum, Arr)
			}
		})

		b.Run("Массив на 1000000 элементов. Число в центре.", func(b *testing.B) {
			Arr := makeArr(1000000)
			desiredNum := 500000

			for i := 0; i < b.N; i++ {
				Search(desiredNum, Arr)
			}
		})

		b.Run("Массив на 1000000 элементов, число крайнее", func(b *testing.B) {
			Arr := makeArr(1000000)
			desiredNum := 1999998

			for i := 0; i < b.N; i++ {
				Search(desiredNum, Arr)
			}
		})
	})
}