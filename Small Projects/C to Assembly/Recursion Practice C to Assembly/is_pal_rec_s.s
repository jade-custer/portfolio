.global pal_rec_s

#a0 is char array
#a1 is start
#a2 is end


pal_rec_s:
    addi sp, sp, -16     #alloc stack space
    sd ra, (sp)         #preserve ra
    sd a0, 8(sp)        #preserve original string

    bge a1,a2, true     #checking if start >= end

    add t0,a0,a1         #compute str[start]
    lbu t0,0(t0)         #getting value

    add t1,a0,a2         #compute str[end]
    lbu t1,0(t1)         #getting value

    bne t0,t1,false     #checking if the values equal

    addi a1,a1,1       #incrementing start
    addi a2,a2,-1       #incrementing end
    jal pal_rec_s
   
    j done

true:
    li a0,1             #setting bool to true
    j done

false:
    li a0,0             #setting bool to false

done:
    ld ra, (sp)         #restore ra
    addi sp, sp, 16     #dealloc stack space
    ret
