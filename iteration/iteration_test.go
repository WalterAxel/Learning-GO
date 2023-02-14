package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("Repeating 'a'", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q wanted %q", repeated, expected)
		}
	})

	t.Run("Repeating count", func(t *testing.T) {
		count := repeatCount
		expected := 5

		if count != expected {
			t.Errorf("%d does not match wanted repetCount %d", count, expected)
		}

	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
