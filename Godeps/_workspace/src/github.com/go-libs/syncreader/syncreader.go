package syncreader

import "io"

type SyncReader struct {
	readers []io.Reader
}

func (sr *SyncReader) Read(p []byte) (n int, err error) {
	if len(sr.readers) > 0 {
		n, err = sr.readers[0].Read(p)
		for _, r := range sr.readers[1:] {
			r.Read(p[:n])
		}
		if n > 0 || err != io.EOF {
			return
		}
		return 0, err
	}
	return 0, io.EOF
}

func New(readers ...io.Reader) io.Reader {
	r := make([]io.Reader, len(readers))
	copy(r, readers)
	return &SyncReader{r}
}