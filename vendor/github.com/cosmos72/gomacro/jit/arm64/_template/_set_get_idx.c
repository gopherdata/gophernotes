#include <stdint.h>

uint8_t getidx_8(uint8_t *addr, uint64_t offset) {
	return addr[offset];
}

uint16_t getidx_16(uint16_t *addr, uint64_t offset) {
	return addr[offset];
}

uint32_t getidx_32(uint32_t *addr, uint64_t offset) {
	return addr[offset];
}

uint64_t getidx_64(uint64_t *addr, uint64_t offset) {
	return addr[offset];
}

uint64_t getidx_64_c(uint64_t *addr) {
	return addr[32];
}


void setidx_8(uint8_t *addr, uint64_t offset, uint8_t value) {
	addr[offset] = value;
}

void setidx_16(uint16_t *addr, uint64_t offset, uint16_t value) {
	addr[offset] = value;
}

void setidx_32(uint32_t *addr, uint64_t offset, uint32_t value) {
	addr[offset] = value;
}

void setidx_64(uint64_t *addr, uint64_t offset, uint64_t value) {
	addr[offset] = value;
}

void setidx_64_c(uint64_t *addr, uint64_t value) {
	addr[32] = value;
}

