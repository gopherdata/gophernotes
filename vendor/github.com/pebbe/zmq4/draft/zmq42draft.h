#if ZMQ_VERSION_MINOR == 2
#ifdef ZMQ_BUILD_DRAFT_API
#define ZMQ42HASDRAFT
#endif
#endif

#ifndef ZMQ42HASDRAFT
#define ZMQ_SERVER -12
#define ZMQ_CLIENT -13
#define ZMQ_RADIO -14
#define ZMQ_DISH -15
#define ZMQ_GATHER -16
#define ZMQ_SCATTER -17
#define ZMQ_DGRAM -18
#endif
