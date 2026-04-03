package concurrency_test

import (
	"concurrency"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return url != "waat://furhurterwe.geds"
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := concurrency.CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want: %v, got: %v", want, got)
	}
}

// Let's use a benchmark to test the speed of CheckWebsites so that we can see the effect of our changes.
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a dummy url"
	}

	for b.Loop() { // loop returns true as long as the benchmark should keep running
		concurrency.CheckWebsites(mockWebsiteChecker, urls)
	}
}
