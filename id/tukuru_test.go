package id_test

import (
	"context"
	"fmt"
	"testing"

	idtukuru "github.com/sinmetalcraft/tukuru/id"
)

func TestTukuru_New(t *testing.T) {
	ctx := context.Background()

	idTukuru := idtukuru.NewTukuru(ctx)
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

func ExampleTukuru_New() {
	ctx := context.Background()

	testIDTukuru := &testIDTukuru{
		"testID",
	}

	idTukuru := idtukuru.NewTukuru(ctx, idtukuru.WithFactory(testIDTukuru))
	id := idTukuru.New()
	fmt.Println(id)
	// Output: testID
}
