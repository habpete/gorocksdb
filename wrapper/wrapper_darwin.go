//go:build darwin
// +build darwin

package wrapper

/*
#cgo LDFLAGS: -L../bin

#include <stdlib.h>
#include "../bin/include/rocksdb/c.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Env is a system call environment used by a database.
type Env struct {
	c *C.rocksdb_env_t
}

// Cache is a cache used to store data read from data in memory.
type Cache struct {
	c *C.rocksdb_cache_t
}

// BlockBasedTableOptions represents block-based table options.
type BlockBasedTableOptions struct {
	c *C.rocksdb_block_based_table_options_t

	// Hold references for GC.
	cache     *Cache
	compCache *Cache

	// We keep these so we can free their memory in Destroy.
	cFp *C.rocksdb_filterpolicy_t
}

// Options represent all of the available options when opening a database with Open.
type Options struct {
	c *C.rocksdb_options_t

	// Hold references for GC.
	env  *Env
	bbto *BlockBasedTableOptions

	// We keep these so we can free their memory in Destroy.
	ccmp *C.rocksdb_comparator_t
	cmo  *C.rocksdb_mergeoperator_t
	cst  *C.rocksdb_slicetransform_t
	ccf  *C.rocksdb_compactionfilter_t
}

// DB is a reusable handle to a RocksDB database on disk, created by Open.
type DB struct {
	c    *C.rocksdb_t
	name string
	opts *Options
}

// OpenDb opens a database with the specified options.
func OpenDb(opts *Options, name string) (*DB, error) {
	var (
		cErr  *C.char
		cName = C.CString(name)
	)
	defer C.free(unsafe.Pointer(cName))
	opts.c = C.rocksdb_options_create()
	db := C.rocksdb_open(opts.c, cName, &cErr)
	if cErr != nil {
		defer C.rocksdb_free(unsafe.Pointer(cErr))
		return nil, errors.New(C.GoString(cErr))
	}
	return &DB{
		name: name,
		c:    db,
		opts: opts,
	}, nil
}
