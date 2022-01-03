package readwriteseeker

import (
	"github.com/orcaman/writerseeker"
	"sync"
)

type ReadWriteSeeker struct {
	sync.Mutex
	rs *writerseeker.WriterSeeker
}

func New() *ReadWriteSeeker {
	return &ReadWriteSeeker{
		rs: &writerseeker.WriterSeeker{},
	}
}

func (rs *ReadWriteSeeker) Write(p []byte) (n int, err error) {
	rs.Lock()
	defer rs.Unlock()
	n, err = rs.Write(p)
	return
}

func (rs *ReadWriteSeeker) Seek(offset int64, whence int) (n int64, err error) {
	rs.Lock()
	defer rs.Unlock()
	n, err = rs.Seek(offset, whence)
	return
}

func (rs *ReadWriteSeeker) Read(p []byte) (n int, err error) {
	rs.Lock()
	defer rs.Unlock()
	n, err = rs.rs.Reader().Read(p)
	return
}
