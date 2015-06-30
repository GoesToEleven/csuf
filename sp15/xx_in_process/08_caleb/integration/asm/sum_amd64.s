#include "textflag.h"

TEXT ·SumV2(SB),NOSPLIT,$0-24
    MOVQ  xs+0(FP),DI
    MOVQ  xs+8(FP),SI
    MOVQ  $0,CX
    MOVQ  $0,AX

L1: CMPQ  AX,SI           // i < len(xs)
    JGE   Z1
    LEAQ  (DI)(AX*8),BX   // BX = &xs[i]
    MOVQ  (BX),BX         // BX = *BX
    ADDQ  BX,CX           // CX += BX
    INCQ  AX              // i++
    JMP   L1

Z1:  MOVQ  CX,ret+24(FP)
    RET
