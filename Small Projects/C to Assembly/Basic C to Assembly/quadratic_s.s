.global quadratic_s

quadratic_s:
    mul t0,a0,a0
    mul t0,t0,a1
    mul t1,a0,a2
    add t0,t0,t1
    add t0,t0,a3
    mv a0,t0       
    ret 
