# collatz conjecture
def collatz(n):
    while n > 1:
        if n&1 != 0:
            n = ((n * 3) + 1) / 2
        else:
            n = n / 2

i = 0
while i < 100000:
    collatz(837799)
    i+=1
