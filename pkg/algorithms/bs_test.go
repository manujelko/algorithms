package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("present", func(t *testing.T) {
		if found := BinarySearch(4, []int{1, 2, 3, 4, 5}); !found {
			t.Error("expecting true")
		}
	})

	t.Run("missing", func(t *testing.T) {
		if found := BinarySearch(6, []int{1, 2, 3, 4, 5}); found {
			t.Error("expecting false")
		}
	})
}
