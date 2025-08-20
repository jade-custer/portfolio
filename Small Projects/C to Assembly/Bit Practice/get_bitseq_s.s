#went over in clas so this code was just a follow along 
.global get_bitseq_s

get_bitseq_s:
    sub t0, a2, a1          # t0 is end - start
    addi t0, t0, 1          # t0 is len

    srl t1, a0, a1          # t1 is num >> start

    li t4, 64				#checking for if len == 64
    bne t4, t0, calc_mask	#if branch not equal then calculate amask 
    li t2, 0xFFFFFFFFFFFFFFFF
    j apply_mask

calc_mask:
                            # calculate mask
    li t2, 1                # t2 is 1
    sll t2, t2, t0          # t2 = 1 << len
    addi t2, t2, -1         # t2 is mask

apply_mask:
    and a0, t1, t2          # apply mask, isolating bits
    ret
