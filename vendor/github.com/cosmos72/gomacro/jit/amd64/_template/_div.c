#include <stdint.h>

#define DIV(T) \
   T div_##T(T a, T b) { \
       return a/b; \
   }

#define PDIV(T) \
   T pdiv_##T(T a, const T * b) { \
       return a / *b; \
   }

DIV(int8_t)
DIV(int16_t)
DIV(int32_t)
DIV(int64_t)

PDIV(int8_t)
PDIV(int16_t)
PDIV(int32_t)
PDIV(int64_t)

DIV(uint8_t)
DIV(uint16_t)
DIV(uint32_t)
DIV(uint64_t)

PDIV(uint8_t)
PDIV(uint16_t)
PDIV(uint32_t)
PDIV(uint64_t)
