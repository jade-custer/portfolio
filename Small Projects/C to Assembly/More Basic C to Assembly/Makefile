PROGS  = find_max_index sort swap to_upper
FIND_MAX_INDEX_OBJS = find_max_index.o find_max_index_c.o find_max_index_s.o 
SORT_OBJS = sort.o sort_c.o sort_s.o  find_max_index_c.o find_max_index_s.o 
SWAP_OBJS = swap.o swap_c.o swap_s.o
TO_UPPER_OBJS = to_upper.o to_upper_c.o to_upper_s.o

%.o: %.c
	gcc -c -g -o $@ $<

%.o: %.s
	as -g -o $@ $<

all: $(PROGS)

find_max_index: $(FIND_MAX_INDEX_OBJS)
	gcc -g -o $@ $^

sort: $(SORT_OBJS)
	gcc -g -o $@ $^

swap: $(SWAP_OBJS)
	gcc -g -o $@ $^

to_upper: $(TO_UPPER_OBJS)
	gcc -g -o $@ $^
	
clean:
	rm -rf $(PROGS) $(FIND_MAX_INDEX_OBJS) $(SORT_OBJS) $(SWAP_OBJS) $(TO_UPPER_OBJS)
