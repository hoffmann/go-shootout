import sys

def fib(n):
    if (n < 2):
        return(1)
    return( fib(n-2) + fib(n-1) )

def main():
    N = int(sys.argv[1])
    #sys.setrecursionlimit(3000)
    print fib(N)

main()

