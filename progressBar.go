package main

import (
	"io"
	"sync"

	"github.com/cheggaaa/pb/v3"
)

type progressBar struct {
	sync.Mutex
	// bar *progressbar.ProgressBar
	bar *pb.ProgressBar
}

type readCloser struct {
	io.Reader
	close func() error
}

func (r *readCloser) Close() error {
	return r.close()
}

func barConfig(bar *pb.ProgressBar) {
	bar.Set(pb.Bytes, true)
	bar.Set(pb.SIBytesPrefix, true)
	bar.Set(pb.Color, true)
	bar.Set(pb.Terminal, true)
}

func (p *progressBar) TrackProgress(src string, current, total int64, stream io.ReadCloser) io.ReadCloser {
	p.Lock()
	defer p.Unlock()
	newBar := pb.New64(total)
	newBar.SetCurrent(current)
	barConfig(newBar)
	if p.bar == nil {
		p.bar = newBar
		p.bar.Start()
	}
	reader := newBar.NewProxyReader(stream)
	return &readCloser{
		Reader: reader,
		close: func() error {
			p.Lock()
			defer p.Unlock()
			p.bar.Finish()
			return nil
		},
	}
}
