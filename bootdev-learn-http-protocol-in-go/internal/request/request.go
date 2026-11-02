package request

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/umbra-codex/lab/bootdev-learn-http-protocol-in-go/internal/headers"
)

type Request struct {
	RequestLine    RequestLine
	Headers        headers.Headers
	state          parserState
	Body           []byte
	bodyLengthRead int
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type parserState int

const (
	stateInitialized parserState = iota
	stateParsingHeaders
	stateParsingBody
	stateDone
)

const (
	crlf       = "\r\n"
	bufferSize = 8
)

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := &Request{
		state:   stateInitialized,
		Headers: headers.NewHeaders(),
		Body:    make([]byte, 0),
	}
	buffer := make([]byte, bufferSize, bufferSize)
	readToIndex := 0

	for request.state != stateDone {
		numBytesParsed, err := request.parse(buffer[:readToIndex])
		if err != nil {
			return nil, err
		}
		copy(buffer, buffer[numBytesParsed:])
		readToIndex -= numBytesParsed
		if numBytesParsed > 0 {
			continue
		}

		if readToIndex >= len(buffer) {
			newBuffer := make([]byte, len(buffer)*2)
			copy(newBuffer, buffer)
			buffer = newBuffer
		}

		numBytesRead, err := reader.Read(buffer[readToIndex:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				if request.state != stateDone {
					return nil, fmt.Errorf("unexpected EOF: incomplete request")
				}
				break
			}
			return nil, err
		}
		readToIndex += numBytesRead
	}

	return request, nil
}

func parseRequestLine(data []byte) (*RequestLine, int, error) {
	index := bytes.Index(data, []byte(crlf))
	if index == -1 {
		return nil, 0, nil
	}

	requestLineText := string(data[:index])
	requestLine, err := requestLineFromString(requestLineText)
	if err != nil {
		return nil, 0, err
	}

	return requestLine, index + len(crlf), nil
}

func requestLineFromString(str string) (*RequestLine, error) {
	parts := strings.Split(str, " ")
	if len(parts) != 3 {
		return nil, fmt.Errorf("poorly formatted request-line: %s", str)
	}

	method := parts[0]
	for _, c := range method {
		if c < 'A' || c > 'Z' {
			return nil, fmt.Errorf("invalid method: %s", method)
		}
	}

	requestTarget := parts[1]

	versionParts := strings.Split(parts[2], "/")
	if len(versionParts) != 2 {
		return nil, fmt.Errorf("malformed start-line: %s", str)
	}

	httpPart := versionParts[0]
	if httpPart != "HTTP" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", httpPart)
	}
	version := versionParts[1]
	if version != "1.1" {
		return nil, fmt.Errorf("unrecognized HTTP-version: %s", version)
	}

	return &RequestLine{
		Method:        method,
		RequestTarget: requestTarget,
		HttpVersion:   versionParts[1],
	}, nil
}

func (r *Request) parse(data []byte) (int, error) {
	if r.state == stateDone {
		return 0, fmt.Errorf("parse called after request was already done")
	}

	switch r.state {
	case stateInitialized:
		requestLine, n, err := parseRequestLine(data)
		if err != nil {
			return 0, err
		}
		if n == 0 {
			return 0, nil
		}
		r.RequestLine = *requestLine
		r.state = stateParsingHeaders
		return n, nil

	case stateParsingHeaders:
		n, done, err := r.Headers.Parse(data)
		if err != nil {
			return 0, err
		}
		if done {
			r.state = stateParsingBody
		}
		return n, nil

	case stateParsingBody:
		contentLengthStr, ok := r.Headers.Get("content-length")
		if !ok {
			r.state = stateDone
			return 0, nil
		}

		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			return 0, fmt.Errorf("malformed Content-Length: %s", err)
		}

		r.Body = append(r.Body, data...)
		r.bodyLengthRead += len(data)
		if r.bodyLengthRead > contentLength {
			return 0, fmt.Errorf("Content-Length too large")
		}
		if r.bodyLengthRead == contentLength {
			r.state = stateDone
		}
		return len(data), nil

	default:
		return 0, fmt.Errorf("unknown parser state: %d", r.state)
	}
}
