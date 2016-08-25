// Copyright 2016 Stellar Development Foundation and contributors. Licensed
// under the Apache License, Version 2.0. See the COPYING file at the root
// of this distribution or at http://www.apache.org/licenses/LICENSE-2.0

package archivist

import (
	"io"
	"path"
	"fmt"
	"net/http"
	"net/url"
	"errors"
)

type HttpArchiveBackend struct {
	client http.Client
	base url.URL
}

func checkResp(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 400 {
		return nil
	} else {
		return fmt.Errorf("Bad HTTP response '%s' for GET '%s'",
			r.Status, r.Request.URL.String())
	}
}

func (b *HttpArchiveBackend) GetFile(pth string) (io.ReadCloser, error) {
	var derived url.URL = b.base
	derived.Path = path.Join(derived.Path, pth)
	resp, err := b.client.Get(derived.String())
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	err = checkResp(resp)
	if err != nil {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
		return nil, err
	}
	return resp.Body, nil
}

func (b *HttpArchiveBackend) Exists(pth string) bool {
	var derived url.URL = b.base
	derived.Path = path.Join(derived.Path, pth)
	resp, err := b.client.Head(derived.String())
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return err == nil && resp != nil && checkResp(resp) == nil
}

func (b *HttpArchiveBackend) PutFile(pth string, in io.ReadCloser) error {
	in.Close()
	return errors.New("PutFile not available over HTTP")
}

func (b *HttpArchiveBackend) ListFiles(pth string) (chan string, chan error) {
	ch := make(chan string)
	er := make(chan error)
	close(ch)
	er <- errors.New("ListFiles not available over HTTP")
	close(er)
	return ch, er
}

func (b *HttpArchiveBackend) CanListFiles() bool {
	return false
}

func MakeHttpBackend(base *url.URL, opts *ConnectOptions) ArchiveBackend {
	return &HttpArchiveBackend{
		base: *base,
	}
}
