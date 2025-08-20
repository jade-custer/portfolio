.global swap_s

# a0 is a pointer to a C int
# a1 is i
# a2 is j

swap_s:
#finding word size
   slli t1,a1,2 
   slli t2,a2,2 

#calculating actual memory location
   add t1,t1,a0 
   add t2,t2,a0 

#loading in word from memory address
   lw t3,0(t1) 
   lw t4,0(t2)

#storing word 
   sw t4,0(t1)
   sw t3,0(t2)

done:
    ret
   
