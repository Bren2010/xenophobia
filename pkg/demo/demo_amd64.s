TEXT ·demo(SB),$0-8
	MOVQ $1, R15
	MOVQ ·num+0(SB), R8
	MOVQ R15, ret+0(FP)
	RET
