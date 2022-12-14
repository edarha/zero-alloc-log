package normallog

import "testing"

func BenchmarkEventWithInterface(b *testing.B) {
	stream := &testStream{}
	LogWriter = stream
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			InforByInterface("test")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatal("Log write count")
	}
}
