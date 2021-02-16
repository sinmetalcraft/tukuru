package now_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	nowtukuru "github.com/sinmetalcraft/tukuru/now"
)

func TestTukuru_New(t *testing.T) {
	ctx := context.Background()

	nowTukuru := nowtukuru.NewTukuru(ctx)
	now := nowTukuru.New()
	if now.IsZero() {
		t.Fatal("Now is Zero")
	}
}

type testNowTukuru struct {
	now time.Time
}

func (tukuru *testNowTukuru) New() time.Time {
	return tukuru.now
}

func ExampleNowTukuru_New() {
	ctx := context.Background()

	testNowTukuru := &testNowTukuru{
		time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC),
	}
	nowTukuru := nowtukuru.NewTukuru(ctx, nowtukuru.WithFactory(testNowTukuru))
	now := nowTukuru.New()
	fmt.Println(now)
	// Output: 2021-02-15 00:00:00 +0000 UTC
}
