## N
## a1 a2 a3 a4 a5 ... aN
firstline = input()
N=int(firstline)
print(N)
secondline=input()
sarr = list (map(int, secondline.split() ) ) 
for v in sarr:
	print((v*2)/2)
