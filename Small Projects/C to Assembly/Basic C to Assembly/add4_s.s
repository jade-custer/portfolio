.global add4_s

add4_s:
    add t0, a0, a1      # t0 = a0 + a1
    add t0,t0,a2
    add t0,t0,a3
    mv a0, t0           # set up ret val in a0
    ret
