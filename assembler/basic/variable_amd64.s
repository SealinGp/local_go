GLOBL ·Id(SB),$8
DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00

GLOBL ·Num(SB),$16
DATA ·Num+0(SB)/8,$0
DATA ·Num+8(SB)/8,$3

GLOBL ·BoolValue(SB),$1

GLOBL ·TrueValue(SB),$1
DATA ·TrueValue(SB)/1,$1

GLOBL ·FalseValue(SB),$1
DATA ·FalseValue(SB)/1,$0

GLOBL ·Int32Value(SB),$4
DATA ·Int32Value+0(SB)/1,$0x01
DATA ·Int32Value+1(SB)/1,$0x02
DATA ·Int32Value+2(SB)/2,$0x03

GLOBL ·Uint32Value(SB),$4
DATA ·Uint32Value(SB)/4,$0x01020304

GLOBL ·Float32Value(SB),$4
DATA ·Float32Value(SB)/4,$1.5

GLOBL ·Float64Value(SB),$8
DATA ·Float64Value(SB)/8,$0x01020304

GLOBL ·M(SB),$8
DATA ·M+0(SB)/8,$0

GLOBL ·Ch(SB),$8
DATA ·Ch+0(SB)/8,$0


#include "textflag.h"
GLOBL ·ReadOnlyInt(SB),NOPTR|RODATA,$8
DATA ·ReadOnlyInt+0(SB)/8,$9527

GLOBL ·NameData(SB),$8
DATA ·NameData+0(SB)/8,$"gopher"

GLOBL ·Name(SB),$16
DATA ·Name+0(SB)/8,$·NameData(SB)
DATA ·Name+8(SB)/8,$6
