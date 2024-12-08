package splicestrmode

import "testing"

func benchmark(b *testing.B, f func(int, string) string) {
	var str = randomString(10)
	for i := 0; i < b.N; i++ {
		f(10000, str)
	}
}
func BenchmarkPlusConcat(b *testing.B) {
	benchmark(b, PlusConcat)
}

func BenchmarkBuilderConcat(b *testing.B) {
	benchmark(b, BuilderConcat)
}
