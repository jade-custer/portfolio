.global str_to_int_s

#a0 = *str
#a1 = base (never changes)

char_to_digit:
#finding char
	add a0,t3,a0		#finding arr[i]
	lb t3,(a0)			#finding ch

	li t0,2   			#checking if base 2
	beq a1,t0,base2		

	li t0,10			#checking if base 10
	beq a1,t0,base10

	li t0,16			#checking if base 16
	beq a1,t0,base16

	ret
		

base2:
	li t0,'0'			#checking if ch = 0
	
	beq t3,t0,base2_01	#branch if ch = 0

	li t0,'1'
	beq t3,t0,base2_01	#branch if ch = 1

	j not_valid

base2_01:
	li t0,'0'			
	sub a0,t3,t0		#return ch - '0'
	ret 
	
base10:
	li t0,'0'
	blt t3,t0,not_valid #if not valid return

	li t0,'9'
	bgt t3,t0,not_valid #if not valid return

	li t0,'0'			
	sub a0,t3,t0		#return ch - '0'

	ret
	
base16:
	li t0, '0'       
    li t1, '9'       
    blt t3, t0,check_lowercase
    bgt t3, t1,check_lowercase
    sub a0, t3, t0    	#ch - '0'
    ret              

check_lowercase:
    li t0, 'a'       
    li t1, 'f'       
    blt t3, t0,check_uppercase
    bgt t3, t1,check_uppercase
    sub t3,t3,t0    	#ch - 'a'
    addi a0,t3,0xa  	#add 0xa
    ret             

check_uppercase:
    li t0,'A'       
    li t1,'F'       
    blt t3,t0,not_valid 
    bgt t3,t1,not_valid 
    sub t3,t3,t0    	#ch - 'A'
    addi a0,t3,0xa  	#add 0xa
    ret               
	

not_valid:
	ret

str_to_int_s:
	li t0, 0 			#retval = 0
	li t1,1				#place_val =  1
	li t2,0				#digit = 0

	li t3,0				#i = 0
	mv t4,a0
	j strlen_loop

strlen_loop:
	lb t5,(t4)			#loading in first byte
	beqz t5,strlen_done #moving to done
	addi t3,t3,1 		#i++
	addi t4,t4,1 		#move onto next char
	j strlen_loop

strlen_done:
	addi t3,t3,-1 		#moving len back one
	#letting it fallthrough into loop

loop:
	blt t3,zero,done	#if branch less than or equal to 0 then done

	addi sp, sp, -48    #allocate space on stack for saving registers
    sd a0, 0(sp)        #save a0 (array pointer)
    sd t0, 8(sp)        #save t0 (retval)
    sd t1, 16(sp)       #save t1 (place_val)
    sd t2, 24(sp)       #save t2 (digit)
    sd t3, 32(sp)		#save t3 (i)
    sd ra, 40 (sp)		#save ra

	jal char_to_digit
	
	mv t4,a0			#grabbing return value 

	ld ra,40(sp)        #restore ra
	ld t3,32(sp)		#restore i 
	ld t2,24(sp)		#restore digit
	ld t1,16(sp)		#restore place_val
	ld t0,8(sp)			#restore retval
	ld a0,(sp)			#restore string
    addi sp, sp, 48     #dealloc stack space

    
	mul t2,t4,t1		#digit * place_val;
	add t0,t0,t2		#retval += digit * place_val;
	mul t1,t1,a1 		#place_val *= base;

	addi t3,t3,-1		#i--
	j loop

done:
	mv a0,t0			#moving return value into a0
    ret
