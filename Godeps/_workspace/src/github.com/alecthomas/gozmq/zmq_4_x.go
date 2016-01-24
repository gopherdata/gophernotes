// +build zmq_4_x

/*
  Copyright 2010-2012 Alec Thomas

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package gozmq

/*
#cgo pkg-config: libzmq
#include <zmq.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

const (
	IPV6            = IntSocketOption(C.ZMQ_IPV6)
	IMMEDIATE       = IntSocketOption(C.ZMQ_IMMEDIATE)
	MECHANISM       = IntSocketOption(C.ZMQ_MECHANISM)
	PLAIN_SERVER    = IntSocketOption(C.ZMQ_PLAIN_SERVER)
	PLAIN_USERNAME  = StringSocketOption(C.ZMQ_PLAIN_USERNAME)
	PLAIN_PASSWORD  = StringSocketOption(C.ZMQ_PLAIN_PASSWORD)
	CURVE_PUBLICKEY = StringSocketOption(C.ZMQ_CURVE_PUBLICKEY)
	CURVE_SECRETKEY = StringSocketOption(C.ZMQ_CURVE_SECRETKEY)
	CURVE_SERVERKEY = StringSocketOption(C.ZMQ_CURVE_SERVERKEY)
	ZAP_DOMAIN      = StringSocketOption(C.ZMQ_ZAP_DOMAIN)
	ROUTER_RAW      = IntSocketOption(C.ZMQ_ROUTER_RAW)
	PROBE_ROUTER    = IntSocketOption(C.ZMQ_PROBE_ROUTER)
	REQ_CORRELATE   = IntSocketOption(C.ZMQ_REQ_CORRELATE)
	REQ_RELAXED     = IntSocketOption(C.ZMQ_REQ_RELAXED)
	CURVE_SERVER    = IntSocketOption(C.ZMQ_CURVE_SERVER)
	CONFLATE        = IntSocketOption(C.ZMQ_CONFLATE)
)
