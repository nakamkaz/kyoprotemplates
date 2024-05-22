#N
#a1
#a2
#..
#aN
firstline=input()
N=int(firstline)
print(N)
sarr = [int(input()) for _ in range(N)]
for value in sarr:
	print((value*2)/2)
