	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	not
	.type	not, @function
not:
	.cfi_startproc
	not	%rax
	not	%rcx
	not	%rdx
	not	%rbx
	not	%rsp
	not	%rbp
	not	%rsi
	not	%rdi
	not	%r8
	not	%r9
	not	%r10
	not	%r11
	not	%r12
	not	%r13
	not	%r14
	not	%r15
	
	ret
	.cfi_endproc

