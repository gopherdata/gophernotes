	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Zero
	.type	Zero, @function
Zero:
	.cfi_startproc
	xor	%eax,%eax
	xor	%ecx,%ecx
	xor	%edx,%edx
	xor	%ebx,%ebx
	xor	%esp,%esp
	xor	%ebp,%ebp
	xor	%esi,%esi
	xor	%edi,%edi
	xor	%r8d,%r8d
	xor	%r9d,%r9d
	xor	%r10d,%r10d
	xor	%r11d,%r11d
	xor	%r12d,%r12d
	xor	%r13d,%r13d
	xor	%r14d,%r14d
	xor	%r15d,%r15d
	nop
	xor	%rax,%rax
	xor	%rcx,%rcx
	xor	%rdx,%rdx
	xor	%rbx,%rbx
	xor	%rsp,%rsp
	xor	%rbp,%rbp
	xor	%rsi,%rsi
	xor	%rdi,%rdi
	xor	%r8,%r8
	xor	%r9,%r9
	xor	%r10,%r10
	xor	%r11,%r11
	xor	%r12,%r12
	xor	%r13,%r13
	xor	%r14,%r14
	xor	%r15,%r15
	ret
	.cfi_endproc
        
        
