package request

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestLineParse(t *testing.T) {
	// Test: Good GET Request line
	reader := &chunkReader{
		data:            "GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 3,
	}
	r, err := RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Good GET Request line with path
	reader = &chunkReader{
		data:            "GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 1,
	}
	r, err = RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.RequestLine.Method)
	assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Good POST Request with path
	reader = &chunkReader{
		data:            "POST /submit HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 5,
	}
	r, err = RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "POST", r.RequestLine.Method)
	assert.Equal(t, "/submit", r.RequestLine.RequestTarget)
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

	// Test: Invalid number of parts in request line
	reader = &chunkReader{
		data:            "/coffee HTTP/1.1\r\nHost: localhost:42069\r\n\r\n",
		numBytesPerRead: 3,
	}
	_, err = RequestFromReader(reader)
	require.Error(t, err)

	// Test: Invalid method (out of order) Request line
	reader = &chunkReader{
		data:            "/coffee POST HTTP/1.1\r\nHost: localhost:42069\r\n\r\n",
		numBytesPerRead: 3,
	}
	_, err = RequestFromReader(reader)
	require.Error(t, err)

	// Test: Invalid version in Request line
	reader = &chunkReader{
		data:            "GET /coffee HTTP/1.2\r\nHost: localhost:42069\r\n\r\n",
		numBytesPerRead: 50,
	}
	_, err = RequestFromReader(reader)
	require.Error(t, err)
}

func TestRquestHeadersParse(t *testing.T) {
	// Test: Standard Headers
	reader := &chunkReader{
		data:            "GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 3,
	}
	r, err := RequestFromReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "localhost:42069", r.Headers["host"])
	assert.Equal(t, "curl/7.81.0", r.Headers["user-agent"])
	assert.Equal(t, "*/*", r.Headers["accept"])

	// Test: Malformed Header
	reader = &chunkReader{
		data:            "GET / HTTP/1.1\r\nHost localhost:42069\r\n\r\n",
		numBytesPerRead: 3,
	}
	r, err = RequestFromReader(reader)
	require.Error(t, err)
}

func TestRequestHeadersParse(t *testing.T) {
	tests := []struct {
		name            string
		request         string
		numBytesPerRead int
		expectedHeaders map[string]string
		expectError     bool
	}{
		{
			name:            "Standard Headers",
			request:         "GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{
				"host":       "localhost:42069",
				"user-agent": "curl/7.81.0",
				"accept":     "*/*",
			},
		},
		{
			name:            "Empty Headers",
			request:         "GET / HTTP/1.1\r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{},
		},
		{
			name:            "Malformed Header",
			request:         "GET / HTTP/1.1\r\nHost localhost:42069\r\n\r\n",
			numBytesPerRead: 3,
			expectError:     true,
		},
		{
			name:            "Duplicate Headers",
			request:         "GET / HTTP/1.1\r\nHost: localhost:42069\r\nHost: example.com\r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{"host": "localhost:42069, example.com"},
		},
		{
			name:            "Case Insensitive Headers",
			request:         "GET / HTTP/1.1\r\nHoSt: localhost:42069\r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{"host": "localhost:42069"},
		},
		{
			name:            "Missing End of Headers",
			request:         "GET / HTTP/1.1\r\nHost: localhost:42069\r\n",
			numBytesPerRead: 3,
			expectError:     true,
		},
		{
			name:            "Whitespace Around Value",
			request:         "GET / HTTP/1.1\r\nHost:    localhost   \r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{"host": "localhost"},
		},
		{
			name:            "Empty Value",
			request:         "GET / HTTP/1.1\r\nHost:\r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{"host": ""},
		},
		{
			name:            "Leading Space In Header Name",
			request:         "GET / HTTP/1.1\r\n Host: localhost\r\n\r\n",
			numBytesPerRead: 3,
			expectError:     true,
		},
		{
			name:            "Valid Token Punctuation",
			request:         "GET / HTTP/1.1\r\nX-Custom_Header.v2: ok\r\n\r\n",
			numBytesPerRead: 3,
			expectedHeaders: map[string]string{"x-custom_header.v2": "ok"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &chunkReader{data: tt.request, numBytesPerRead: tt.numBytesPerRead}
			r, err := RequestFromReader(reader)
			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, r)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, r)
			require.Len(t, r.Headers, len(tt.expectedHeaders))
			for key, want := range tt.expectedHeaders {
				got, ok := r.Headers[key]
				require.True(t, ok, "missing header %q", key)
				assert.Equal(t, want, got, "header %q", key)
			}
		})
	}
}

func TestRequestHeadersAnyChunkSize(t *testing.T) {
	request := "GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nAccept: */*\r\n\r\n"
	for size := 1; size <= len(request); size++ {
		reader := &chunkReader{data: request, numBytesPerRead: size}
		r, err := RequestFromReader(reader)
		require.NoError(t, err, "chunk size %d", size)
		require.NotNil(t, r, "chunk size %d", size)
		assert.Equal(t, "localhost:42069", r.Headers["host"], "chunk size %d", size)
		assert.Equal(t, "*/*", r.Headers["accept"], "chunk size %d", size)
	}
}

type chunkReader struct {
	data            string
	numBytesPerRead int
	pos             int
}

func (cr *chunkReader) Read(p []byte) (n int, err error) {
	if cr.pos >= len(cr.data) {
		return 0, io.EOF
	}
	end := min(cr.pos+cr.numBytesPerRead, len(cr.data))
	n = copy(p, cr.data[cr.pos:end])
	cr.pos += n
	return n, nil
}
