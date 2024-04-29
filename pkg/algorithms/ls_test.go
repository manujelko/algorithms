package algorithms

import "testing"

func LinearSearchTest(t *testing.T) {
	t.Run("present", func(t *testing.T) {
		if found := LinearSearch(4, []int{1, 2, 3, 4, 5}); !found {
			t.Error("true expected")
		}
	})

	t.Run("missing", func(t *testing.T) {
		if found := LinearSearch(6, []int{1, 2, 3, 4, 5}); found {
			t.Error("false expected")
		}
	})
}
