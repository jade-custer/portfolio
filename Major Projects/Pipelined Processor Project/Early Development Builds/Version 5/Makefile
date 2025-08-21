PROG = analyze
OBJS = analyze.o part1_s.o part2_s.o

%.o: %.c
	gcc -c -g -o $@ $<

%.o: %.s
	as -g -o $@ $<

$(PROG): $(OBJS)
	gcc -g -o $@ $^
	
clean:
	rm -rf $(PROG) $(OBJS)