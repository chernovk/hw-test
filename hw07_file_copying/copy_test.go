package main

import (
	"errors"
	"testing"
)

// func TestCopy(t *testing.T) {
// 	// Place your code here.
// 	t.Run(tt.name, func(t *testing.T) {
// 		err := Clone(tt.src, tt.dst)
// 		if err != tt.wantErr {
// 			t.Errorf("Clone() error = %v, wantErr %v", err, tt.wantErr)
// 		}
// 	})
// }

func TestCopy(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		dst     string
		offset  int64
		limit   int64
		wantErr error
	}{
		{"Unprocessible files src", "/dev/urandom", "out.txt", 0, 0, ErrUnsupportedFile},
		{"Empty dst", "testdata/input.txt", "", 0, 0, ErrCreate},
		{"Empty src", "", "destination", 0, 0, ErrCantGetFileInfo},
		{"Too big Offset", "testdata/input.txt", "out.txt", 7000, 1, ErrOffsetExceedsFileSize},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.src, tt.dst, tt.offset, tt.limit)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Clone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
