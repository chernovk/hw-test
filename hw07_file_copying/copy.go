package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrCantGetFileInfo       = errors.New("cannot get file info")
	ErrRead                  = errors.New("cannot read from file")
	ErrWrite                 = errors.New("cannot write to file")
	ErrClose                 = errors.New("cannot close file")
	ErrCreate                = errors.New("cannot create file")
)

func copyBunch(fromFile, toFile *os.File, buf []byte, iterLimit, iterOffset, offset int64) error {
	_, err := fromFile.ReadAt(buf[iterOffset:], offset+iterOffset)
	if errors.Is(err, io.EOF) {
		return err
	} else if err != nil {
		return ErrRead
	}

	_, err = toFile.WriteAt(buf[iterOffset:iterOffset+iterLimit], iterOffset)
	if err != nil {
		return ErrWrite
	}
	return nil
}

func closeFiles(fromFile, toFile *os.File) {
	err := fromFile.Close()
	if err != nil {
		log.Fatal(ErrClose)
	}
	err = toFile.Close()
	if err != nil {
		log.Fatal(ErrClose)
	}
}

func prepareFiles(fromPath, toPath string, offset int64) (fromFile, toFile *os.File, buf []byte, err error) {
	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return nil, nil, nil, ErrCantGetFileInfo
	}

	if fileInfo.Mode()&(os.ModeDevice|os.ModeCharDevice|os.ModeNamedPipe|os.ModeSocket) != 0 {
		return nil, nil, nil, ErrUnsupportedFile
	}

	fileSize := fileInfo.Size()

	if fileSize < offset {
		return nil, nil, nil, ErrOffsetExceedsFileSize
	}

	fromFile, err = os.Open(fromPath)
	if err != nil {
		return nil, nil, nil, ErrUnsupportedFile
	}
	toFile, err = os.Create(toPath)
	if err != nil {
		return nil, nil, nil, ErrCreate
	}

	buf = make([]byte, int(fileSize)-int(offset))
	return fromFile, toFile, buf, nil
}

func printProgressBar(progress, finish int64) {
	res := progress * 100 / finish
	fmt.Println(strings.Repeat("*", int(res))+strings.Repeat("-", int(100-res)), " ", res, "%")
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, toFile, buf, err := prepareFiles(fromPath, toPath, offset)
	if err != nil {
		return fmt.Errorf("error when trying to prepare the files: %w", err)
	}
	defer closeFiles(fromFile, toFile)

	bufferLen := int64(len(buf))

	var iterationOffset int64
	var iterationLimit int64 = 1024

	if limit == 0 || limit > bufferLen {
		limit = bufferLen
	}

	// bar := pb.StartNew(int(limit))

	for iterationOffset < limit {
		if iterationOffset+iterationLimit > limit {
			iterationLimit = limit - iterationOffset
		}

		err = copyBunch(fromFile, toFile, buf, iterationLimit, iterationOffset, offset)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return fmt.Errorf("error when trying to copy bunch: %w", err)
		}
		iterationOffset += iterationLimit
		// bar.Add(int(iterationLimit))
		printProgressBar(iterationOffset, limit)
	}
	// bar.Finish()

	return nil
}
