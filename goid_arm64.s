// +build arm64
// +build go1.9

#include "textflag.h"
// func goid() int64
TEXT ·goid(SB),NOSPLIT,$0-8
	MOVD g, R14
	MOVD 152(R14), R13
	MOVD R13, ret+0(FP)
	RET
