.global to_upper_s

# a0 is a pointer to a C int


to_upper_s:
#initializing variables
    li t1,97  #t1 = a
    li t2,122 #t2 = z
    li t3, 32 #difference between a and 0

loop: #has to be lb and sb because dealing with chars so byte by byte if not things are missing 
    lb t4,0(a0) #reading in current char
    beqz t4,done #if the loop hits null terminator 

    blt t4,t1,skip #if str[i] < 97 skip
    bge t4,t2,skip #if str[i] > 122 skip

    sub t4,t4,t3 #changing lowercase into upper by using ascii characters

    sb  t4,0(a0) #reassigning values
    j loop

skip:
    addi a0,a0,1
    j loop
   
done:
    ret
