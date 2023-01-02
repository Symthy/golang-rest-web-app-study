package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return Run(ctx)
	})

	in := "world"
	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		assert.Fail(t, "failed to read response body: %v", err)
	}

	want := fmt.Sprintf("Hello, %s!", in)
	assert.Equal(t, want, string(got))

	cancel()
	assert.NoError(t, eg.Wait())
}
