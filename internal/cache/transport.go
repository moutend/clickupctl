package cache

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Transport struct {
	debug   *log.Logger
	buffers map[*url.URL]*bytes.Buffer
	cache   *FileCache
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	now := time.Now().UTC()

	t.debug.Printf("Request: Sent at: %s\n", now.Format(time.RFC3339))
	t.debug.Printf("Request: URL: %s %s\n", req.Method, req.URL)

	for k, v := range req.Header {
		t.debug.Printf("Request: Header: %s: %s", k, strings.Join(v, ";"))
	}

	res, err := t.cache.Load(ctx, req, now)

	if err != nil {
		return nil, err
	}
	if res != nil {
		t.debug.Printf("Response: read from cache")

		return res, nil
	}

	res, err = http.DefaultTransport.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buffer := &bytes.Buffer{}

	body, err := io.ReadAll(io.TeeReader(res.Body, buffer))

	if err != nil {
		return nil, fmt.Errorf("cache: %w", err)
	}

	res.Body = io.NopCloser(buffer)

	t.debug.Printf("Response Status: %s\n", res.Status)

	for k, v := range res.Header {
		t.debug.Printf("Response Header: %s: %s", k, strings.Join(v, ";"))
	}

	t.debug.Printf("Response Body: %s\n", body)

	if err := t.cache.Save(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (t *Transport) SetLogger(l *log.Logger) {
	if l == nil {
		return
	}

	t.debug = l
}

func NewTransport(ctx context.Context) (*Transport, error) {
	cache, err := NewFileCache(ctx)

	if err != nil {
		return nil, err
	}

	transport := &Transport{
		debug:   log.New(io.Discard, "", 0),
		buffers: make(map[*url.URL]*bytes.Buffer),
		cache:   cache,
	}

	return transport, nil
}
