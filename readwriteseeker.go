package readwriteseeker

import (
	"io"
	"sync"

	"github.com/orcaman/writerseeker"
)

var _ io.ReadWriteSeeker = &ReadWriteSeeker{}

type ReadWriteSeeker struct {
	sync.Mutex
	ws *writerseeker.WriterSeeker
}

func New() *ReadWriteSeeker {
	return &ReadWriteSeeker{
		ws: &writerseeker.WriterSeeker{},
	}
}

func (rs *ReadWriteSeeker) Write(p []byte) (n int, err error) {
	rs.Lock()
	defer rs.Unlock()
	n, err = rs.ws.Write(p)
	return
}

func (rs *ReadWriteSeeker) Seek(offset int64, whence int) (n int64, err error) {
	rs.Lock()
	defer rs.Unlock()
	n, err = rs.ws.Seek(offset, whence)
	return
}

func (rs *ReadWriteSeeker) Read(p []byte) (n int, err error) {
	rs.Lock()
	defer rs.Unlock()
	n, err = rs.ws.Reader().Read(p)
	return
}
