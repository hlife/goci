package goci

/*
#include <oci.h>
#include <stdlib.h>
#include <string.h>

#cgo CFLAGS: -I/home/oracle/app/oracle/product/11.2.0/client_1/rdbms/public
#cgo LDFLAGS: -lclntsh -L/home/oracle/app/oracle/product/11.2.0/client_1/lib

*/
import "C"
import (
	"unsafe"
)

type result struct {
	stmt *statement
}

func (r *result) LastInsertId() (int64, error) {
	var t C.ub4
	if C.OCIAttrGet(r.stmt.handle, C.OCI_HTYPE_STMT, unsafe.Pointer(&t), nil, C.OCI_ATTR_ROWID, (*C.OCIError)(r.stmt.conn.err)) != C.OCI_SUCCESS {
		return 0, ociGetError(r.stmt.conn.err)
	}
	return int64(t), nil
}

func (r *result) RowsAffected() (int64, error) {
	var t C.ub4
	if C.OCIAttrGet(r.stmt.handle, C.OCI_HTYPE_STMT, unsafe.Pointer(&t), nil, C.OCI_ATTR_ROW_COUNT, (*C.OCIError)(r.stmt.conn.err)) != C.OCI_SUCCESS {
		return 0, ociGetError(r.stmt.conn.err)
	}
	return int64(t), nil
}
