package tukuru_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sinmetalcraft/tukuru"
)

func TestID_New(t *testing.T) {
	ctx := context.Background()

	idTukuru := tukuru.NewIDTukuru(ctx)
	id := idTukuru.New()
	if id == "" {
		t.Fatal("ID is empty")
	}
}

type testIDTukuru struct {
	id string
}

func (tukuru *testIDTukuru) New() string {
	return tukuru.id
}

func ExampleID_New() {
	ctx := context.Background()

	testIDTukuru := &testIDTukuru{
		"testID",
	}

	idTukuru := tukuru.NewIDTukuru(ctx, tukuru.WithIDFactory(testIDTukuru))
	id := idTukuru.New()
	fmt.Println(id)
	// Output: testID
}
