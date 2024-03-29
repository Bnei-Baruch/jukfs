/*
Copyright 2011 The Perkeep Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package blobstore

import (
	"context"
	"errors"
	"io"
	"os"

	"github.com/Bnei-Baruch/jukfs/pkg/blob"
)

// NoImplStorage is an implementation of Storage that returns a not
// implemented error for all operations.
type NoImplStorage struct{}

var _ Storage = NoImplStorage{}

func (NoImplStorage) Fetch(context.Context, blob.Ref) (file io.ReadCloser, size uint32, err error) {
	return nil, 0, os.ErrNotExist
}

func (NoImplStorage) ReceiveBlob(context.Context, blob.Ref, io.Reader) (sb blob.SizedRef, err error) {
	err = errors.New("ReceiveBlob not implemented")
	return
}

func (NoImplStorage) StatBlobs(ctx context.Context, blobs []blob.Ref, fn func(blob.SizedRef) error) error {
	return errors.New("Stat not implemented")
}

func (NoImplStorage) EnumerateBlobs(ctx context.Context, dest chan<- blob.SizedRef, after string, limit int) error {
	close(dest)
	return errors.New("EnumerateBlobs not implemented")
}

func (NoImplStorage) RemoveBlobs(ctx context.Context, blobs []blob.Ref) error {
	return errors.New("Remove not implemented")
}
