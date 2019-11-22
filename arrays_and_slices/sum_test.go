package sum

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Array(numbers)
	want := 15

	if got != want {
		t.Errorf("given numbers: %q, got %q and wanted %q", numbers, got, want)
	}
}

func TestSlice(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Slice(numbers)
		want := 15

		if got != want {
			t.Errorf("given numbers: %q, got %q and wanted %q", numbers, got, want)
		}
	})

	t.Run("collection of 7 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 1, 2}

		got := Slice(numbers)
		want := 18

		if got != want {
			t.Errorf("given numbers: %q, got %q and wanted %q", numbers, got, want)
		}
	})
}

func TestAllSingleSlice(t *testing.T) {
	slice := []int{1, 1}

	got := AllSingleSlice(slice)
	want := []int{2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("given slice: %q, got %q and wanted %q", slice, got, want)
	}

}

func TestAllMultipleSlices(t *testing.T) {
	t.Run("collection of 2 slices", func(t *testing.T) {
		sliceOne := []int{1, 1}
		sliceTwo := []int{1, 2, 3}

		got := AllMultipleSlices(sliceOne, sliceTwo)
		want := []int{2, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("given slice one: %v and clice two: %v, got %v and wanted %v", sliceOne, sliceTwo, got, want)
		}
	})

	t.Run("collection of 4 slices", func(t *testing.T) {
		sliceOne := []int{1, 1}
		sliceTwo := []int{1, 2, 3}
		sliceThree := []int{10, 0}
		sliceFour := []int{66, 4}

		got := AllMultipleSlices(sliceOne, sliceTwo, sliceThree, sliceFour)
		want := []int{2, 6, 10, 70}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("given slice one: %v, clice two: %v, slice three: %v, and slive four: %v, got %v and wanted %v", sliceOne, sliceTwo, sliceThree, sliceFour, got, want)
		}
	})
}

func TestBiggestsFromSingleSlice(t *testing.T) {
	slice := []int{3, 1}

	got := BiggestsFromSingleSlice(slice)
	want := []int{3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("given slice: %v, got %v and wanted %v", slice, got, want)
	}

}

func TestBiggestsFromMultipleSlices(t *testing.T) {
	sliceOne := []int{3, 1}
	sliceTwo := []int{3, 9}

	got := BiggestsFromMultipleSlices(sliceOne, sliceTwo)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("given slice one: %v and slice two: %v, got %v and wanted %v", sliceOne, sliceTwo, got, want)
	}

}
