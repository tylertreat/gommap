// +build windows

package gommap

const (
	// RDONLY maps the memory read-only.
	// Attempts to write to the MMap object will result in undefined behavior.
	RDONLY = 0
	// RDWR maps the memory as read-write. Writes to the MMap object will update the
	// underlying file.
	RDWR = 1 << iota
	// COPY maps the memory as copy-on-write. Writes to the MMap object will affect
	// memory, but the underlying file will remain unchanged.
	COPY
	// If EXEC is set, the mapped memory is marked as executable.
	EXEC
)

type ProtFlags uint

const (
	// PROT_NONE  ProtFlags = 0x0
	// PROT_READ  ProtFlags = 0x1
	// PROT_WRITE ProtFlags = 0x2
	// PROT_EXEC  ProtFlags = 0x4

	// PROT_NONE  ProtFlags = 0x0	/* No PROT_NONE on windows */
	PROT_READ  ProtFlags = RDONLY
	PROT_WRITE ProtFlags = RDWR
	PROT_COPY  ProtFlags = COPY /* PROT_COPY Only on windows, use this flag to implement MAP_PRIVATE */
	PROT_EXEC  ProtFlags = EXEC
)

/* on win, use ProtFlags to simulate MapFlags */

type MapFlags uint

const (
	MAP_SHARED    MapFlags = 0x1
	MAP_PRIVATE   MapFlags = 0x2
	MAP_FIXED     MapFlags = 0x10
	MAP_ANONYMOUS MapFlags = 0x20
	MAP_GROWSDOWN MapFlags = 0x100
	MAP_LOCKED    MapFlags = 0x2000
	MAP_NONBLOCK  MapFlags = 0x10000
	MAP_NORESERVE MapFlags = 0x4000
	MAP_POPULATE  MapFlags = 0x8000
)

type SyncFlags uint

const (
	MS_SYNC       SyncFlags = 0x4
	MS_ASYNC      SyncFlags = 0x1
	MS_INVALIDATE SyncFlags = 0x2
)

type AdviseFlags uint

const (
	MADV_NORMAL     AdviseFlags = 0x0
	MADV_RANDOM     AdviseFlags = 0x1
	MADV_SEQUENTIAL AdviseFlags = 0x2
	MADV_WILLNEED   AdviseFlags = 0x3
	MADV_DONTNEED   AdviseFlags = 0x4
	MADV_REMOVE     AdviseFlags = 0x9
	MADV_DONTFORK   AdviseFlags = 0xa
	MADV_DOFORK     AdviseFlags = 0xb
)
