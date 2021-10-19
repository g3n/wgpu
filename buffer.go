package main

// Buffer Address
type BufferAddress uint64

// Buffer Size
type BufferSize uint64 // nonzero

// Buffer Binding Type
type BufferBindingType uint32

const (
	BufferBindingTypeUniform = BufferBindingType(iota)
	BufferBindingTypeStorage
	BufferBindingTypeStorageReadOnly
)

// Buffer Usage
type BufferUsage uint32

const (
	BufferUsageMapRead      = BufferUsage(1 << iota) // Allow a buffer to be mapped for reading
	BufferUsageMapWrite                              // Allow a buffer to be mapped for writing
	BufferUsageCopySrc                               // Allow a buffer to be the source buffer for a buffer copy operation
	BufferUsageCopyDst                               // Allow a buffer to be the destination buffer for a buffer copy operation
	BufferUsageIndex                                 // Allow a buffer to be the index buffer in a draw operation.
	BufferUsageVertex                                // Allow a buffer to be the vertex buffer in a draw operation.
	BufferUsageUniform                               // Allow a buffer to be a BufferBindingTypeUniform inside a bind group.
	BufferUsageStorage                               // Allow a buffer to be a BufferBindingTypeStorage inside a bind group.
	BufferUsageIndirect                              // Allow a buffer to be the indirect buffer in an indirect draw call.
	BufferUsageQueryResolve                          // Allow a buffer to be the indirect buffer in an indirect draw call.
)

type BufferDescriptor struct {
	Label            string
	Size             BufferAddress
	Usage            BufferUsage
	MappedAtCreation bool
}
