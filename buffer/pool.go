/*
####### sdk.base (c) 2024 Archivage Numérique ###################################################### MIT License #######
''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''''
*/

package buffer

import "sync"

type (
	Pool struct {
		sp sync.Pool
	}
)

func NewPool(size int) *Pool {
	return &Pool{
		sp: sync.Pool{
			New: func() any {
				return &Buffer{
					bs: make([]byte, 0, size),
				}
			},
		},
	}
}

func (p *Pool) Get() *Buffer {
	buf := p.sp.Get().(*Buffer)

	buf.Reset()
	buf.pool = p

	return buf
}

func (p *Pool) put(buf *Buffer) {
	p.sp.Put(buf)
}

/*
####### END ############################################################################################################
*/
