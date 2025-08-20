.global unpack_bytes_s

#a0 = val
#a1 = bytes[]

unpack_bytes_s:
	li t0,0			#i = 0
	li t1,0			#b = 0
	li t3,4			#highest value in for loop

start:
	blt t0,t3,loop	#go into the loop if 
	
	j done

loop:
	and  t1,a0,0xFF # b = val & 0xFF;

	slli t2,t0,2	#finding byte offset
	add t2,a1,t2	#findind address of bytes[i]
	sw t1,0(t2)		#setting bytes[i] as b

	srl a0,a0,8 	# val = val >> 8;

	addi t0,t0,1	#i++
	
	j start			#jump to for condition
	
	
done:
    ret
