.global fib_rec_s

# a0 = n (input)

fib_rec_s:
    addi sp, sp, -24    	#allocate stack space
    sd ra, 16(sp)       	#save return address
    sd a0, 8(sp)        	#save n
    sd t2, 0(sp)        	#save t2 

    li t0, 1            	#t0 = 1
    ble a0, t0, base_case   #if n <= 1, return n

# Recursive Case 1: fib(n-1)
    addi a0, a0, -1     	#a0 = n-1
    jal fib_rec_s       	#compute fib(n-1)
    mv t2, a0           	#store fib(n-1) in t2

# Recursive Case 2: fib(n-2)
    ld a0, 8(sp)        	#restore original n
    addi a0, a0, -2     	#a0 = n-2
    jal fib_rec_s       	#compute fib(n-2)

# Combine fib(n-1) + fib(n-2)
    add a0, t2, a0      	#a0 = fib(n-1) + fib(n-2)

    j done              

base_case:
    ld a0, 8(sp)        	#load original n
    j done              	#return n (fib(0) = 0, fib(1) = 1)

done:
    ld t2, 0(sp)        	#restore t2
    ld ra, 16(sp)       	#restore return address
    addi sp, sp, 24     	#deallocate stack space
    ret                 
