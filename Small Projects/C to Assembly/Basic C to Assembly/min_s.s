.global min_s

min_s:
    sub t0, a0, a1       
    bge t0, x0, set_t2    
    mv a0, a0           
    j done

set_t2:
    mv a0, a1            

done:
    ret
