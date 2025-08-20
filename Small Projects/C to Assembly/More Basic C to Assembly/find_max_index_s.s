.global find_max_index_s
# a0 is a pointer to a C int
# a1 is the len of the array

find_max_index_s:
    li t0,0 #max_index = 0
    lw t1,0(a0) #max_value = arr[0]
    li t2,1 #i = 1 (starting at 1 because arr[0] is always considered biggest at beginning)

loop:
    beq t2,a1,done #if no more values in *arr then exit

    slli t3, t2, 2 #finding offset
    add t4, a0, t3 #finding arr[i]
    lw t5, 0(t4) #load value

    bgt t5,t1,update_max #if new value is greater then update value

    addi t2,t2,1 #i++
    j loop 

update_max:
    mv t0,t2 #update new index value
    mv t1,t5 #update new max value
    addi t2,t2,1 #i++
    j loop

done:
    mv a0,t0 
    ret
