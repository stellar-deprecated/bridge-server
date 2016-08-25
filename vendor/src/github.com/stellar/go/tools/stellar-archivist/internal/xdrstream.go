// Copyright 2016 Stellar Development Foundation and contributors. Licensed
// under the Apache License, Version 2.0. See the COPYING file at the root
// of this distribution or at http://www.apache.org/licenses/LICENSE-2.0

package archivist

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/stellar/go/xdr"
)

type XdrStream struct {
	buf  bytes.Buffer
	rdr  io.ReadCloser
	rdr2 io.ReadCloser
}

func NewXdrStream(in io.ReadCloser) *XdrStream {
	return &XdrStream{rdr: bufReadCloser(in)}
}

func NewXdrGzStream(in io.ReadCloser) (*XdrStream, error) {
	rdr, err := gzip.NewReader(bufReadCloser(in))
	if err != nil {
		in.Close()
		return nil, err
	}
	return &XdrStream{rdr: bufReadCloser(rdr), rdr2: in}, nil
}

func (a *Archive) GetXdrStream(pth string) (*XdrStream, error) {
	if !strings.HasSuffix(pth, ".xdr.gz") {
		return nil, errors.New("File has non-.xdr.gz suffix: " + pth)
	}
	rdr, err := a.backend.GetFile(pth)
	if err != nil {
		return nil, err
	}
	return NewXdrGzStream(rdr)
}

func HashXdr(x interface{}) (Hash, error) {
	var msg bytes.Buffer
	_, err := xdr.Marshal(&msg, x)
	if err != nil {
		var zero Hash
		return zero, err
	}
	return Hash(sha256.Sum256(msg.Bytes())), nil
}

func (x *XdrStream) Close() {
	if x.rdr != nil {
		x.rdr.Close()
	}
	if x.rdr2 != nil {
		x.rdr2.Close()
	}
}

func (x *XdrStream) ReadOne(in interface{}) error {
	var nbytes uint32
	err := binary.Read(x.rdr, binary.BigEndian, &nbytes)
	if err != nil {
		x.rdr.Close()
		if err == io.ErrUnexpectedEOF {
			return io.EOF
		} else {
			return err
		}
	}
	nbytes &= 0x7fffffff
	x.buf.Reset()
	if nbytes == 0 {
		x.rdr.Close()
		return io.EOF
	}
	x.buf.Grow(int(nbytes))
	read, err := x.buf.ReadFrom(io.LimitReader(x.rdr, int64(nbytes)))
	if read != int64(nbytes) {
		x.rdr.Close()
		return errors.New("Read wrong number of bytes from XDR")
	}
	if err != nil {
		x.rdr.Close()
		return err
	}

	readi, err := xdr.Unmarshal(&x.buf, in)
	if int64(readi) != int64(nbytes) {
		return fmt.Errorf("Unmarshalled %d bytes from XDR, expected %d)",
			readi, nbytes)
	}
	return err
}

func WriteFramedXdr(out io.Writer, in interface{}) error {
	var tmp bytes.Buffer
	n, err := xdr.Marshal(&tmp, in)
	un := uint32(n)
	if un > 0x7fffffff {
		return fmt.Errorf("Overlong write: %d bytes", n)
	}

	un = un | 0x80000000
	binary.Write(out, binary.BigEndian, &un)
	k, err := tmp.WriteTo(out)
	if int64(n) != k {
		return fmt.Errorf("Mismatched write length: %d vs. %d", n, k)
	}
	return err
}
