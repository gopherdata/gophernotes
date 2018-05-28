	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	neg
	.type	neg, @function
neg:
	.cfi_startproc
	neg	%rax
	neg	%rcx
	neg	%rdx
	neg	%rbx
	neg	%rsp
	neg	%rbp
	neg	%rsi
	neg	%rdi
	neg	%r8
	neg	%r9
	neg	%r10
	neg	%r11
	neg	%r12
	neg	%r13
	neg	%r14
	neg	%r15
	
	ret
	.cfi_endproc

