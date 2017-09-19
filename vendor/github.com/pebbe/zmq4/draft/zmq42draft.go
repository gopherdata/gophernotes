package zmq4

/*
#include <zmq.h>
#include <stdlib.h>
#include <string.h>
#include "zmq4.h"
#include "zmq42draft.h"

#ifdef ZMQ42HASDRAFT
int zmq4has42draft = 1;
#else
int zmq4has42draft = 0;
// Version >= 4.2.0 with draft

int zmq_join (void *s, const char *group) { return 0; }
int zmq_leave (void *s, const char *group) { return 0; }
int zmq_msg_set_routing_id(zmq_msg_t *msg, uint32_t routing_id) { return 0; }
uint32_t zmq_msg_routing_id(zmq_msg_t *msg) { return 0; }
int zmq_msg_set_group(zmq_msg_t *msg, const char *group) { return 0; }
const char *zmq_msg_group(zmq_msg_t *msg) { return NULL; }

#endif // ZMQ42HASDRAFT

*/
import "C"

import (
	"unsafe"
)

type OptRoutingId uint32
type OptGroup string

var (
	has42draft bool
)

func init() {
	has42draft = (C.zmq4has42draft != 0)
}

func (soc *Socket) Join(group string) error {
	if !has42draft {
		return ErrorNotImplemented42draft
	}
	cs := C.CString(group)
	defer C.free(unsafe.Pointer(cs))
	n, err := C.zmq_join(soc.soc, cs)
	if n != 0 {
		return errget(err)
	}
	return nil
}

func (soc *Socket) Leave(group string) error {
	if !has42draft {
		return ErrorNotImplemented42draft
	}
	cs := C.CString(group)
	defer C.free(unsafe.Pointer(cs))
	n, err := C.zmq_leave(soc.soc, cs)
	if n != 0 {
		return errget(err)
	}
	return nil
}
