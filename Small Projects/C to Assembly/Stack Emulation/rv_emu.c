#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "rv_emu.h"
#include "bits.h"

#define DEBUG 0

static void unsupported(char *s, uint32_t n) {
    printf("unsupported %s 0x%x\n", s, n);
    exit(-1);
}

void emu_r_type( rv_state *state, uint32_t iw) {
    uint32_t rd = (iw >> 7) & 0b11111;
    uint32_t rs1 = (iw >> 15) & 0b11111;
    uint32_t rs2 = (iw >> 20) & 0b11111;
    uint32_t funct3 = (iw >> 12) & 0b111;
    uint32_t funct7 = (iw >> 25) & 0b1111111;

	if (funct3 == 0b000 && funct7 == 0b0000000) { //add
		state->regs[rd] = state->regs[rs1] + state->regs[rs2];
	} else if (funct3 == 0b000 && funct7 == 0b0000001){ //mul
		state->regs[rd] = state->regs[rs1] * state->regs[rs2];
	} else if (funct3 == 0b000 && funct7 == 0b0100000){ //sub
		state->regs[rd] = state->regs[rs1] - state->regs[rs2];
	} else if (funct3 == 0b101 && funct7 == 0b0000000) {  // srl
		state->regs[rd] = state->regs[rs1] >> state->regs[rs2];
 	} else if (funct3 == 0b001 && funct7 == 0b0000000) {  // sll
		state->regs[rd] = state->regs[rs1] << state->regs[rs2];
	} else if (funct3 == 0b111 && funct7 == 0b0000000) {  // and
		state->regs[rd] = state->regs[rs1] & state->regs[rs2];
	} else if (funct3 == 0b110  && funct7 == 0b0000001) {  // rem
   		state->regs[rd] = state->regs[rs1] %  state->regs[rs2];
	} else if (funct3 == 0b100  && funct7 == 0b0000001) {  // div
	   		state->regs[rd] = state->regs[rs1] /  state->regs[rs2];
	} else {
		unsupported("R-type funct3", funct3);
	}
	
    state->pc += 4; // Next instruction
}

void emu_i_type( rv_state *state, uint32_t iw) {
    uint32_t rd = (iw >> 7) & 0b11111;
    uint32_t rs1 = (iw >> 15) & 0b11111;
    uint32_t funct3 = (iw >> 12) & 0b111;
    int32_t imm = (int32_t)(iw >> 20); 

	//used for when imm is within bits 1-11
    if (imm & (1 << 11)) {  
        imm |= 0xFFFFF000; 
    }
    
    if (funct3 == 0b101) { //srli
        state->regs[rd] = state->regs[rs1] >> imm;
    } else if (funct3 == 0b000) {  //addi
        state->regs[rd] = state->regs[rs1] + (int32_t) imm;
    } else if (funct3 == 0b001) {  //slli
            state->regs[rd] = state->regs[rs1] <<  (int32_t) imm;
    } else {
        unsupported("I-type funct3", funct3);
    }
    
      state->pc += 4; // Next instruction
}

void emu_b_type( rv_state *state, uint32_t iw) {
    uint32_t rs1 = (iw >> 15) & 0b11111;
    uint32_t rs2 = (iw >> 20) & 0b11111;
    uint32_t funct3 = (iw >> 12) & 0b111;

   int32_t imm = ((iw >> 31) << 12) |  
                  ((iw >> 7) & 0b1) << 11 |
                  ((iw >> 25) & 0b111111) << 5 |
                  ((iw >> 8) & 0b1111) << 1;


    if (funct3 == 0b100) { //BLT (used for BGT by flipping rs1 and rs2 since its a pseudoinstruction)
        if ((int64_t) state->regs[rs1] < (int64_t) state->regs[rs2]) {
        	state -> analysis.b_taken+= 1;
            state->pc += imm;
            return;
  		}
	} else if (funct3 == 0b001) { // bne
		 if ((int64_t)state->regs[rs1] != (int64_t)state->regs[rs2]) {
		 	state -> analysis.b_taken+= 1;
			state->pc += imm;
			return;
		}
	} else if (funct3 == 0b000) { // beq
		 if ((int64_t)state->regs[rs1] == (int64_t)state->regs[rs2]) {
		 	state -> analysis.b_taken+= 1;
			state->pc += imm;
			return;
		}
	} else if (funct3 == 0b101) { // bge
		 if ((int64_t)state->regs[rs1] >= (int64_t)state->regs[rs2]) {
		 	state -> analysis.b_taken+= 1;
			state->pc += imm;
			return;
		}
	} else {
	    unsupported("B-type funct3", funct3);
	}
	state -> analysis.b_not_taken+= 1;
	state->pc += 4;
}

void emu_jalr(rv_state *state, uint32_t iw) {
    uint32_t rs1 = (iw >> 15) & 0b1111;  // Will be ra (aka x1)
    uint64_t val = state->regs[rs1];  // Value of regs[1]

    state->pc = val;  // PC = return address
}

void emu_jal(rv_state *state, uint32_t iw) {
    uint32_t rd = (iw >> 7) & 0b11111;
    int32_t imm = ((iw >> 31) << 20) |  
                  ((iw >> 12) & 0xFF) << 12 |  
                  ((iw >> 20) & 0b1) << 11 |  
                  ((iw >> 21) & 0x3FF) << 1; 
                  
    //sign extend 
    if (imm & (1 << 20)) {  
        imm |= 0xFFFFF000; 
    }

 	if (rd == 0){
 		state -> pc += imm;
 		return;	
 	}
 	
    state -> regs[rd] = state -> pc + 4;
    
    state->pc += imm;
}

void emu_load(rv_state *state, uint32_t iw) {
	uint32_t rd = (iw >> 7) & 0b11111;  
    uint32_t rs1 = (iw >> 15) & 0b11111;  
    int32_t imm = (int32_t)(iw >> 20); 
    uint32_t func3 = (iw >> 12) & 0b111; 
    uint64_t mta = state -> regs[rs1] + imm;

    //sign extending imm 
    if (imm & (1 << 11)) {
        imm |= 0xFFFFF000; 
    }

    if (func3 == 0b000){ // lb
    	state -> regs[rd] = *(uint8_t*)  mta;
    }else if (func3 == 0b001){ //lh
    	state -> regs[rd] = *(uint16_t*) mta;	
    }else if (func3 == 0b010){ //lw
    	state -> regs[rd] = *(uint32_t*)  mta;
    }else if (func3 == 0b011){ //ld	
    	state -> regs[rd] = *(uint64_t*) mta   ;    	
    }

    state->pc += 4; 
    
}

void emu_save(rv_state *state, uint32_t iw) {
 	uint32_t rs1 = (iw >> 15) & 0b11111; 
    uint32_t rs2 = (iw >> 20) & 0b11111; 
    int32_t imm = ((iw >> 25 & 0x7F ) << 5) | ((iw >> 7)& 0x1F);
    uint32_t func3 = (iw >> 12) & 0b111;
    uint64_t mta = state -> regs[rs1] + imm;

    //sign extending imm
    if (imm & (1 << 11)) {
        imm |= 0xFFFFF000;
    }

    if (func3 == 0b000){ // sb
    	*(uint8_t*) mta = state -> regs[rs2];
    }else if (func3 == 0b001){ //sh
    	*(uint16_t*) mta = state -> regs[rs2];
    }else if (func3 == 0b010){ //sw
    	 *(uint32_t*) mta = state -> regs[rs2];
    }else if (func3 == 0b011){ //sd
    	*(uint64_t*) mta = state -> regs[rs2];       	
    }

    state->pc += 4;
   
}

static void rv_one(rv_state *state) {
    uint32_t iw  = *((uint32_t*) state->pc);
    iw = cache_lookup(&state->i_cache, (uint64_t) state->pc);
    

    uint32_t opcode = get_bits(iw, 0, 7);


#if DEBUG
    printf("iw: %x\n", iw);
#endif
	state -> analysis.i_count+= 1;

    switch (opcode) {
            case 0b0110011:
                // R-type instructions have two register operands
                state -> analysis.ir_count+= 1;
                emu_r_type(state, iw);
                break;
            case 0b0010011:
                //I-type instructions
                state -> analysis.ir_count+= 1;
                emu_i_type(state,iw);
                break;
            case 0b1100011:
                //B-type instructions
                emu_b_type(state, iw);
                break;
            case 0b1100111:
                // JALR (aka RET) is a variant of I-type instructions
                state -> analysis.j_count+= 1;
                emu_jalr(state, iw);
                break;
            case 0b1101111:
                // JAL 
                state -> analysis.j_count+= 1;
                emu_jal(state, iw);
                break;
            case 0b0000011:
            	// LD
            	state -> analysis.ld_count+= 1;
            	emu_load(state, iw);
            	 break;
            case 0b0100011 :
            	//SD
            	state -> analysis.st_count+= 1;
            	emu_save(state, iw);
            	break;
            default:
                unsupported("Unknown opcode: ", opcode);       
        }
}

void rv_init(rv_state *state, uint32_t *target, 
             uint64_t a0, uint64_t a1, uint64_t a2, uint64_t a3) {
    state->pc = (uint64_t) target;
    state->regs[RV_A0] = a0;
    state->regs[RV_A1] = a1;
    state->regs[RV_A2] = a2;
    state->regs[RV_A3] = a3;

    state->regs[RV_ZERO] = 0;  // zero is always 0  (:
    state->regs[RV_RA] = RV_STOP;
    state->regs[RV_SP] = (uint64_t) &state->stack[STACK_SIZE];

    memset(&state->analysis, 0, sizeof(rv_analysis));
    cache_init(&state->i_cache);
}

uint64_t rv_emulate(rv_state *state) {
    while (state->pc != RV_STOP) {
        rv_one(state);
    }
    return state->regs[RV_A0];
}

static void print_pct(char *fmt, int numer, int denom) {
    double pct = 0.0;

    if (denom)
        pct = (double) numer / (double) denom * 100.0;
    printf(fmt, numer, pct);
}

void rv_print(rv_analysis *a) {
    int b_total = a->b_taken + a->b_not_taken;

    printf("=== Analysis\n");
    print_pct("Instructions Executed  = %d\n", a->i_count, a->i_count);
    print_pct("R-type + I-type        = %d (%.2f%%)\n", a->ir_count, a->i_count);
    print_pct("Loads                  = %d (%.2f%%)\n", a->ld_count, a->i_count);
    print_pct("Stores                 = %d (%.2f%%)\n", a->st_count, a->i_count);    
    print_pct("Jumps/JAL/JALR         = %d (%.2f%%)\n", a->j_count, a->i_count);
    print_pct("Conditional branches   = %d (%.2f%%)\n", b_total, a->i_count);
    print_pct("  Branches taken       = %d (%.2f%%)\n", a->b_taken, b_total);
    print_pct("  Branches not taken   = %d (%.2f%%)\n", a->b_not_taken, b_total);
}
