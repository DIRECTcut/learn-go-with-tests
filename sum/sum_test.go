package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10

		assertSum(got, want, t, numbers)
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		assertSum(got, want, t, numbers)
	})
}

func assertSum(got int, want int, t *testing.T, numbers []int) {
	if got != want {
		t.Errorf("Want %d got %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	inputLeft := []int{0, 9}
	inputRight := []int{1, 2}

	want := []int{9, 3}
	got := SumAll(inputLeft, inputRight)

	assertSumTails(got, want, t, inputLeft, inputRight)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum non empty slices", func(t *testing.T) {
		inputLeft := []int{0, 9}
		inputRight := []int{1, 2}
		got := SumAllTails(inputLeft, inputRight)
		want := []int{9, 2}

		assertSumTails(got, want, t, inputLeft, inputRight)
	})

	t.Run("Sum empty slices", func(t *testing.T) {
		inputLeft := []int{}
		inputRight := []int{1, 2}
		got := SumAllTails(inputLeft, inputRight)
		want := []int{0, 2}

		assertSumTails(got, want, t, inputLeft, inputRight)
	})

	t.Run("Sum empty input", func(t *testing.T) {
		got := SumAllTails()
		want := []int{}

		assertSumTails(got, want, t, nil, nil)
	})
}

func assertSumTails(got []int, want []int, t *testing.T, inputLeft []int, inputRight []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v got %v given %v and %v", want, got, inputLeft, inputRight)
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3}, []int{2, 3})
	}
}
