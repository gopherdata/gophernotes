	.file	"_jmp.s"
	.text

	.p2align 4,,15
	.globl	Jmp
	.type	Jmp, @function
Jmp:
.LFB1:
	.cfi_startproc
	jmp	.L4
.L3:
    jmp .L3
.L4:
	ret
	.cfi_endproc
