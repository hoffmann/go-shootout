import sys
sys.setrecursionlimit(5000000)

def Ack(M, N):
    if (not M):
        return( N + 1 )
    if (not N):
        return( Ack(M-1, 1) )
    return( Ack(M-1, Ack(M, N-1)) )

def main():
    NUM = int(sys.argv[1])
    sys.setrecursionlimit(10000)
    print "Ack(3,%d): %d" % (NUM, Ack(3, NUM))

main()

