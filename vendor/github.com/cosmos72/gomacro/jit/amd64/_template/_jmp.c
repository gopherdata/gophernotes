
int abs(int x) {
    return x < 0 ? -x : x;
}

unsigned collatz(unsigned n) {
    unsigned ret = 0;
    while (n > 1) {
	ret++;
	if (n & 1) {
	    n = (n * 3) + 1;
	}
	n >>= 1;
    }
    return ret;
}
