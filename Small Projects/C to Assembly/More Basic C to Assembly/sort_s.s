.global sort_s

# a0 is a pointer to a C int
# a1 is the len of the array

sort_s:
    addi t0, a1, 0             		 #t0 = len (loop counter)
    addi t1, zero, 0           		 #t1 = i (loop index)
    addi sp, sp, -8            		 #allot space for 1 reg
    sd ra, (sp)                		 #preserve ra

loop:
    bge t1, t0, done   		 		 #if i >= len, end

    # Save registers before calling find_max_index_s
    addi sp, sp, -32          	 	 #alocate space on stack for saving registers
    sd a0, 0(sp)             	 	 #save a0 (array pointer)
    sd t0, 8(sp)             	 	 #save t0 (len)
    sd t1, 16(sp)            	 	 #save t1 (i)
    sd a1, 24(sp)            	 	 #save a1 (len - i)

    jal ra, find_max_index_s 

#saving return value
    mv t2,a0
    
#restore saved registers
    ld a0, 0(sp)             		 #restore a0 (array pointer)
    ld t0, 8(sp)             		 #restore t0 (len)
    ld t1, 16(sp)            		 #restore t1 (i)
    ld a1, 24(sp)            		 #restore a1 (len - i)
    addi sp, sp, 32           	 	 #deallocate space on stack

 #load arr[i] and arr[idx]
    lw t4, 0(a0)             		 #t4 = arr[i]

    slli t5, t2, 2             		 #t5 = byte offset
    add t5, a0, t5            		 #t5 = &arr[idx]
    lw t6, 0(t5)             		 #t6 = arr[idx]

    ble  t6, t4, continue	 		 #if swap not needed continue

 #swap arr[i] and arr[idx]
    sw t6, 0(a0)                	 #arr[i] = arr[idx]
    sw t4, 0(t5)                	 #arr[idx] = arr[i]

continue:
	addi a1,a1,-1
	addi a0,a0,4
    addi t1, t1, 1              	 # i++
    j loop                

done:
    ld ra, 0(sp)                	 #restore return address
    addi sp, sp, 8               	 #deallocate space for ra
    ret                            
