.global int_to_str_s


#a0 = value
#a1 = result_str
#a2 = obase


int_to_str_s:
	addi sp, sp, -64	#tmp
	mv a3,sp			#converting sp into a register
	mv t0, a0			#t0 = val
	mv t1,zero			#t1 = len
	mv t2,zero			#t2 = j
	

	bnez  t0,loop		#if value != 0 branch
	li t3, '0'			#t3 = '0'
	sb t3,(a3)			#save value
	addi t1,t1,1		#len ++
	j base
	

loop:
	beqz t0,base		#if while loop needs to terminate
	div t3,t0,a2		#div = value / obase
	rem t4,t0,a2		#rem = value % obase;

	li t5,10			#t5 = 10
	blt t4,t5,store_val	#if (rem <= 9)

	li t6,87			#ascii offset for correct value 

	add t4,t4,t5		#getting ascii value for letter
	
	j store_done
	
store_val:
	li t5, '0'			#used for storing  '0'
	add t4,t5,t4		#'0' + rem

store_done:
	add t6,a3,t1		#memory address 
	sb t4, (t6)         # tmp[len] = converted char
	addi t1, t1, 1      # len++
	mv t0, t3           # value = div
	j loop


base:
	li t3,10			#base 10
	beq t3,a2,reverse_tmp
	
	li t3,2				#base 2
	bne t3,a2,check_hex #branch to hex

	#obase = 2
	li t3,48			#t3 = 0
	add t0,a1,t2		#getting address
	sb t3,(t0)			#savng 0 into arr
	addi t2,t2,1		#j + 1
	add t0,a1,t2		#getting address
	li t3,98			#t3 = b
	sb t3,(t0)			#saving b 

	addi t2,t2,1		#j +=1 since j got incremented once 

	j reverse_tmp		
	

check_hex:
	li t3,48			#t3 = 0
	add t0,a1,t2		#getting address
	sb t3,(t0)			#savng 0 into arr
	addi t2,t2,1		#j + 1
	add t0,a1,t2		#getting address
	li t3,120			#t3 = x
	sb t3,(t0)			#saving x 

	addi t2,t2,1		#j += 1 since j got incremented once 

	j reverse_tmp		

reverse_tmp:
	mv t0,t1			#i = len
	li t1,0				#end condition

for_loop:
	blt t0,t1,done		#if loop is done then move on 
	
	addi t0,t0,-1 		# i-1
	add t4,a3,t0		#getting memory address
	lb t3,(t4)			#getting value out of the memory address

	add a1,a1,t2		#getting memory address for result_str[j]
	sb t3,(a1)			#saving tmp[i-1] into result_str[j]

	addi t2,t2,1 		#j++

	j for_loop

	
done:
	add a1,a1,t2		#getting memory address for result_str[j]
	li t3, '\0'
	sb t3,(a1)			#saving tmp[i-1] into result_str[j]
	
	addi sp,sp,64
    ret
