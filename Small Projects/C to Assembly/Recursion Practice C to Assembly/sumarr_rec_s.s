.global sumarr_rec_s
#used example from in class

#a0 = arr 
#a1 = start 
#a2 = end 

sumarr_rec_s:
	slli t0, a1, 2        #t0 = start * 4 
    add t0, a0, t0        #t0 = arr[start]
    lw t2, (t0)           #load arr[start] into t1
 
    bne  a1, a2,recursion #base case
	mv a0,t2			  #add sum into return
	j done
  
recursion:
	addi sp, sp, -16  	  #allocate stack space
    sd ra, (sp)           #save return address
    sd t2, 8(sp)
    
 	addi a1, a1, 1        #increment start
    jal sumarr_rec_s 
    
	ld ra, (sp)           #restore return address
	ld t2,8(sp)			  #restore t2
    addi sp, sp, 16       #deallocate stack space

    add a0, a0, t2        #add arr[start] to the result

done:
   ret
