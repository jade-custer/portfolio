PROGS = pack_bytes unpack_bytes get_bitseq get_bitseq_signed int_to_str str_to_int

PACKBYTES_OBJS = pack_bytes.o pack_bytes_c.o pack_bytes_s.o
UNPACKBYTES_OBJS = unpack_bytes.o unpack_bytes_c.o unpack_bytes_s.o 
GET_BITSEQ_OBJS = get_bitseq.o get_bitseq_s.o get_bitseq_c.o
GET_BITSEQ_SIGNED_OBJS = get_bitseq_signed.o\
						 get_bitseq_signed_c.o get_bitseq_c.o\
                         get_bitseq_signed_s.o get_bitseq_s.o
INT_TO_STR_OBJS = int_to_str.o int_to_str_c.o int_to_str_s.o\
                  conv_util.o str_to_int_c.o
STR_TO_INT_OBJS = str_to_int.o str_to_int_c.o str_to_int_s.o\
                  conv_util.o

OBJS = $(PACKBYTES_OBJS) $(UNPACKBYTES_OBJS) \
	   $(GET_BITSEQ_OBJS) $(GET_BITSEQ_SIGNED_OBJS) \
       $(INT_TO_STR_OBJS) $(STR_TO_INT_OBJS)

%.o: %.c
	gcc -c -g -o $@ $<

%.o: %.s
	as -g -o $@ $<

all: $(PROGS)

pack_bytes: $(PACKBYTES_OBJS)
	gcc -g -o $@ $^

unpack_bytes: $(UNPACKBYTES_OBJS)
	gcc -g -o $@ $^


get_bitseq: $(GET_BITSEQ_OBJS)
	gcc -g -o $@ $^

get_bitseq_signed: $(GET_BITSEQ_SIGNED_OBJS)
	gcc -g -o $@ $^

int_to_str: $(INT_TO_STR_OBJS)
	gcc -g -o $@ $^

str_to_int: $(STR_TO_INT_OBJS)
	gcc -g -o $@ $^

clean:
	rm -rf $(PROGS) $(OBJS)
