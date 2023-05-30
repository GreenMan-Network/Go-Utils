package blockbucket

import (
	"github.com/GreenMan-Network/Go-Utils/pkg/queue/block"
	"github.com/GreenMan-Network/Go-Utils/pkg/queue/fifo"
)

var bucket *BlockBucket

// New - Creates a new block bucket
func Init() {
	bucket = &BlockBucket{
		fifo: fifo.New(),
	}
}

// NewBlock - Creates a new block or get one from the bucket
func NewBlock(data interface{}) *block.Block {
	if bucket == nil {
		Init()
	}

	if bucket.fifo.GetNumBlocks() == 0 {
		return block.New(data)
	} else {
		return bucket.fifo.Pop()
	}
}

// ReturnBlock - Erases and returns a block to the bucket
func ReturnBlock(block *block.Block) {
	if bucket == nil {
		Init()
	}

	if block == nil {
		return
	}

	block.Erase()
	bucket.fifo.Push(block)
}

// GetNumBlocks - Returns the number of blocks in the bucket
func GetNumBlocks() int {
	return bucket.fifo.GetNumBlocks()
}
