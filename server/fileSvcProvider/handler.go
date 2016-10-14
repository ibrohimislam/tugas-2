package fileSvcProvider

import (
	"errors"
	"fmt"
	"github.com/ibrohimislam/tugas-2/services/tugas"
	"github.com/ibrohimislam/tugas-2/utils"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type uploadProgress struct {
	Target  string
	Handler string
	TempDir string
	Chunks  []string
	Hash    string
}

type FSHandler struct {
	sync.RWMutex
	UploadProgess map[string]*uploadProgress
}

func NewFSHandler() *FSHandler {
	return &FSHandler{UploadProgess: make(map[string]*uploadProgress)}
}

func (ch *FSHandler) Dir(path string) ([]*tugas.File, error) {
	var files []*tugas.File

	osFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, osFile := range osFiles {

		mtime := osFile.ModTime()
		stat := osFile.Sys().(*syscall.Stat_t)
		ctime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))

		file := &tugas.File{
			Filename:     osFile.Name(),
			Size:         strconv.FormatUint(uint64(osFile.Size()), 10),
			Mode:         osFile.Mode().String(),
			ModifiedTime: mtime.Format(time.RFC822),
			CreatedTime:  ctime.Format(time.RFC822),
			IsDir:        osFile.IsDir(),
		}
		files = append(files, file)

	}

	return files, nil
}

func (ch *FSHandler) CreateDir(path string, name string) error {
	return os.Mkdir(path+"/"+name, 0666)
}

func (ch *FSHandler) GetContent(path string, name string) ([]int8, error) {
	content, err := ioutil.ReadFile(path + "/" + name)
	return b2i(content), err
}

func (ch *FSHandler) PutContentOpen(path string, name string, hash string) (handler string, err error) {
	var tempDir string

	handler, err = utils.UUID()
	if err != nil {
		return
	}

	err = os.Mkdir(os.TempDir()+"/"+handler, 0777)
	if err != nil {
		return
	}

	ch.UploadProgess[handler] = &uploadProgress{
		Handler: handler,
		TempDir: tempDir,
		Hash:    hash,
		Target:  path + "/" + name,
		Chunks:  make([]string, 0),
	}
	return
}

func (ch *FSHandler) PutContentPartial(handler string, content []int8) (byteWritten int32, err error) {
	name, err := utils.UUID()
	ch.UploadProgess[handler].Chunks = append(ch.UploadProgess[handler].Chunks, name)

	contentByte := i2b(content)
	fullPath := os.TempDir() + "/" + ch.UploadProgess[handler].Handler + "/" + name

	var file *os.File

	file, err = os.Create(fullPath)
	if err != nil {
		return
	}

	defer file.Close()

	var _byteWritten int

	_byteWritten, err = file.Write(contentByte)
	byteWritten = int32(_byteWritten)

	return
}

func (ch *FSHandler) PutContentClose(handler string) (byteWritten int32, err error) {
	var targetFile *os.File

	cProgress := ch.UploadProgess[handler]
	tempPath := os.TempDir() + "/" + cProgress.Handler

	targetFile, err = os.Create(cProgress.Target)
	if err != nil {
		return
	}

	for _, chunk := range cProgress.Chunks {
		var chunkContent []byte
		var chunkByteWritten int

		fullPath := tempPath + "/" + chunk
		chunkContent, err = ioutil.ReadFile(fullPath)
		if err != nil {
			return
		}

		fmt.Printf("[%s] size: %d\n", chunk, len(chunkContent))

		chunkByteWritten, err = targetFile.Write(chunkContent)
		if err != nil {
			return
		}

		byteWritten += int32(chunkByteWritten)
	}

	targetFile.Close()

	// clean up
	os.RemoveAll(tempPath)
	delete(ch.UploadProgess, handler)

	var hash string
	hash = utils.Sha256file(cProgress.Target)

	if hash != cProgress.Hash {
		err = errors.New("hash doesn't match.")
		return
	}

	return
}

func i2b(bs []int8) []byte {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return b
}

func b2i(bs []byte) []int8 {
	b := make([]int8, len(bs))
	for i, v := range bs {
		b[i] = int8(v)
	}
	return b
}
