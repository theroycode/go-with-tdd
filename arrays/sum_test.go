package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	t.Run("for defined size", func(t *testing.T) {
		got := SumAllTails([]int{2, 4}, []int{3, 4, 8}, []int{7, 8, 1})
		want := []int{4, 12, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("when empty array is passed", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{}, []int{3, 8})
		want := []int{3, 0, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4})
	want := []int{3, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v expected %v", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum for any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("expected %q got %q given %v", want, got, numbers)
		}
	})
	t.Run("sum for size of any number", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("expected %q got %q given %v", want, got, numbers)
		}
	})
}
