#include "textflag.h"
#include "funcdata.h"

DATA text<>+0(SB)/8,$"Hello"
DATA text<>+8(SB)/8,$" debug"
GLOBL text<>(SB),NOPTR,$16


TEXT ·asmSay(SB),$16-0
        NO_LOCAL_POINTERS
        MOVQ $text<>+0(SB), AX
        MOVQ AX, (SP)
        MOVQ $16, 8(SP)
        CALL runtime·printstring(SB)
        RET