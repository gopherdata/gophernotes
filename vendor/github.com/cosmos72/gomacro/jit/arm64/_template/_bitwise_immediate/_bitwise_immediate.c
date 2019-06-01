// taken from
// https://stackoverflow.com/questions/30904718/range-of-immediate-values-in-armv8-a64-assembly/33265035#33265035

#include <stdint.h>
#include <stdio.h>

// Dumps all legal bitmask immediates for ARM64
// Total number of unique 64-bit patterns:
//   1*2 + 3*4 + 7*8 + 15*16 + 31*32 + 63*64 = 5334

const char *uint64_to_binary(uint64_t x) {
  static char b[65];
  unsigned i;
  for (i = 0; i < 64; i++, x <<= 1)
    b[i] = (0x8000000000000000ULL & x) ? '1' : '0';
  b[64] = '\0';
  return b;
}

int main() {
  uint64_t result;
  unsigned size, length, rotation, e;
  for (size = 2; size <= 64; size *= 2) {
    for (length = 1; length < size; ++length) {
      result = 0xffffffffffffffffULL >> (64 - length);
      for (e = size; e < 64; e *= 2)
        result |= result << e;
      for (rotation = 0; rotation < size; ++rotation) {
        printf("0x%016llx %s (size=%u, length=%u, rotation=%u)\n",
               (unsigned long long)result, uint64_to_binary(result), size,
               length, rotation);
        result = (result >> 63) | (result << 1);
      }
    }
  }
  return 0;
}
