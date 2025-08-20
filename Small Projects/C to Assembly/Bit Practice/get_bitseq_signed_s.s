.global get_bitseq_signed_s
.global get_bitseq_s

#a0 = num
#a1 = start
#a2 = end

get_bitseq_signed_s:

	sub t2,a2,a1	#len = end - start
	addi t2,t2,1	#len = len + 1

	li t4,64		#creating temp register for 64 because a register is 64 bits
	sub t3,t4,t2	#shift_amt = 64 - len

	addi sp, sp,-16 #allocate stack space
    sd ra, (sp)     #save return address
    sd t3, 8(sp)

    jal get_bitseq_s
    mv t0,a0		#val = return value

    ld ra, (sp)     #restore return address
    ld t3, 8(sp)
    addi sp, sp,16  #deallocate stack space

    sll t0,t0,t3	#val = val << shift_amt

	slli t0,t0,0
    sra t1,t0,t3	#val_signed = ((int) val) >> shift_amt

   
    mv a0,t1		#assigning val into the return
    ret
