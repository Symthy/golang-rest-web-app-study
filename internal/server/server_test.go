package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func serverStartup(t *testing.T) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	eg, ctx := errgroup.WithContext(ctx)
	l, err := net.Listen("tcp", "localhost:0")
	require.NoError(t, err)
	eg.Go(func() error {
		return Run(ctx, l)
	})

	t.Cleanup(func() {
		cancel()
		require.NoError(t, eg.Wait())
	})

	baseUri := fmt.Sprintf("http://%s/", l.Addr().String())
	return baseUri
}

func TestRun(t *testing.T) {
	baseUri := serverStartup(t)

	in := "world"
	rsp, err := http.Get(baseUri + in)
	require.NoError(t, err)
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	require.NoError(t, err)
	want := fmt.Sprintf("Hello, %s!", in)
	assert.Equal(t, want, string(got))
}
