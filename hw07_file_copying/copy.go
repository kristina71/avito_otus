package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrInvalidPath           = errors.New("invalid path")
	ErrSetOffset             = errors.New("unable to set offset")
	ErrFileStat              = errors.New("unable to get file's stat")
	ErrCopy                  = errors.New("copy failed")
)

func min(a, b int64) int64 {
	if a > b {
		return b
	}

	return a
}

func closeFile(f io.Closer) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func readFile(file *os.File) []byte {
	expData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return expData
}

func removeFile(tmpFile *os.File) {
	err := os.Remove(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
}

func tempFile(pattern string) *os.File {
	baseDir := "/tmp"
	tmpFile, err := ioutil.TempFile(baseDir, pattern)
	if err != nil {
		log.Fatal(err)
	}
	return tmpFile
}

func createDir(inputPath string) {
	err := os.MkdirAll(inputPath, 0777) //nolint
	if err != nil {
		log.Fatal(err)
	}
}

func Copy(fromPath string, toPath string, offset int64, limit int64) error {
	in, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("%w: open for reading failed %s", ErrInvalidPath, err)
	}
	defer in.Close()

	inFileInfo, err := in.Stat()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFileStat, err)
	}
	inFileSize := inFileInfo.Size()

	if offset > 0 {
		if offset > inFileSize {
			return fmt.Errorf(
				"%w: given offset %d to large, file %q has len %d",
				ErrOffsetExceedsFileSize,
				offset,
				fromPath,
				inFileSize,
			)
		}
		if _, err := in.Seek(offset, io.SeekStart); err != nil {
			return fmt.Errorf("%w: set position %d in file %q failed %s", ErrSetOffset, offset, from, err)
		}
	}

	var reader io.Reader = in
	needToCopyBytes := inFileSize

	if limit > 0 {
		reader = io.LimitReader(in, limit)
		needToCopyBytes = min(limit, inFileSize-offset)
	}

	writer, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("%w: unable to open file %q for writing %s", ErrInvalidPath, toPath, err)
	}
	defer closeFile(writer)

	bar := pb.Full.Start64(needToCopyBytes)
	defer bar.Finish()

	barReader := bar.NewProxyReader(reader)
	if _, err = io.Copy(writer, barReader); err != nil {
		return fmt.Errorf("%w: %s", ErrCopy, err)
	}

	return nil
}
