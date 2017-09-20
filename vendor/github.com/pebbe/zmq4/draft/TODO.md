## TODO for ZMQ 4.2 DRAFT

see: https://github.com/zeromq/libzmq/releases/tag/v4.2.0

### New poller mechanism and APIs have been introduced in DRAFT state:

    zmq_poller_new
    zmq_poller_destroy
    zmq_poller_add
    zmq_poller_modify
    zmq_poller_remove
    zmq_poller_wait
    zmq_poller_wait_all
    zmq_poller_add_fd
    zmq_poller_modify_fd
    zmq_poller_remove_fd

and a new supporting struct typedef

    zmq_poller_event_t

They support existing socket type, new thread-safe socket types and file
descriptors (cross-platform).

Documentation will be made available in the future before these APIs are
declared stable.

**Status in Go**: Not implemented because there is no documentation.

### New cross-platform timers helper functions have been introduced in DRAFT state:

    zmq_timers_new
    zmq_timers_destroy
    zmq_timers_add
    zmq_timers_cancel
    zmq_timers_set_interval
    zmq_timers_reset
    zmq_timers_timeout
    zmq_timers_execute

and a new supporting callback typedef:

    zmq_timer_fn

**Status in Go:** Not implemented because there is no documentation.
Do we need this in Go?
