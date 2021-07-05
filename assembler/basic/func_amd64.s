#define SWAP(x,y,t) MOVQ x, t;MOVQ y, x;MOVQ t,y
TEXT ·Swap(SB),NOSPLIT,$0-32
       MOVQ a+0(FP), AX
       MOVQ b+8(FP), BX
       SWAP(AX,BX,CX)

       MOVQ AX,ret0+16(FP)
       MOVQ BX,ret1+24(FP)
       RET
