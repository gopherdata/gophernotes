	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	idiv
	.type	idiv, @function
idiv:
	.cfi_startproc
	idiv	%rax
	idiv	%rcx
	idiv	%rdx
	idiv	%rbx
	idiv	%rsp
	idiv	%rbp
	idiv	%rsi
	idiv	%rdi
	idiv	%r8
	idiv	%r9
	idiv	%r10
	idiv	%r11
	idiv	%r12
	idiv	%r13
	idiv	%r14
	idiv	%r15
	idivq	0x288(%rdi)
	ret
	.cfi_endproc

	.globl	div
	.type	div, @function
div:
	.cfi_startproc
	div	%rax
	div	%rcx
	div	%rdx
	div	%rbx
	div	%rsp
	div	%rbp
	div	%rsi
	div	%rdi
	div	%r8
	div	%r9
	div	%r10
	div	%r11
	div	%r12
	div	%r13
	div	%r14
	div	%r15
	divq	0x288(%rdi)
	ret
	.cfi_endproc

