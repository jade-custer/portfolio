.global pack_bytes_s

#a0 = b3
#a1 = b2
#a2 = b1
#a3 = b0


pack_bytes_s:
	li t0,0			#t0 = 0 (val)

	mv t0,a0		#val = b3
	
	slli t0,t0,8	#val << 8
	or t0,t0,a1		#(val << 8) | b2

	slli t0,t0,8	#val << 8
	or t0,t0,a2		#(val << 8) | b1

	slli t0,t0,8	#val << 8
	or t0,t0,a3		#(val << 8) | b0

	mv a0,t0		#moving val into a0 for return
    ret
