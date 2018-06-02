	.file	"reg.s"
	.text
	.p2align 4,,15
	.globl	Store
	.type	Store, @function
Store:
	.cfi_startproc
	mov	%rax, 648(%rdi)
	mov	%rcx, 648(%rdi)
	mov	%rdx, 648(%rdi)
	mov	%rbx, 648(%rdi)
	mov	%rsp, 648(%rdi)
	mov	%rbp, 648(%rdi)
	mov	%rsi, 648(%rdi)
	mov	%rdi, 648(%rdi)
	mov	%r8,  648(%rdi)
	mov	%r9,  648(%rdi)
	mov	%r10, 648(%rdi)
	mov	%r11, 648(%rdi)
	mov	%r12, 648(%rdi)
	mov	%r13, 648(%rdi)
	mov	%r14, 648(%rdi)
	mov	%r15, 648(%rdi)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Load
	.type	Load, @function
Load:
	.cfi_startproc
	mov	648(%rdi), %rax
	mov	648(%rdi), %rcx
	mov	648(%rdi), %rdx
	mov	648(%rdi), %rbx
	mov	648(%rdi), %rsp
	mov	648(%rdi), %rbp
	mov	648(%rdi), %rsi
	mov	648(%rdi), %rdi
	mov	648(%rdi), %r8
	mov	648(%rdi), %r9
	mov	648(%rdi), %r10
	mov	648(%rdi), %r11
	mov	648(%rdi), %r12
	mov	648(%rdi), %r13
	mov	648(%rdi), %r14
	mov	648(%rdi), %r15
	ret
	.cfi_endproc

	.p2align 4,,15
	.globl	Load_s32
	.type	Load_s32, @function
Load_s32:
	.cfi_startproc
	mov	$-0x11223344, %rax
	mov	$-0x11223344, %rcx
	mov	$-0x11223344, %rdx
	mov	$-0x11223344, %rbx
	mov	$-0x11223344, %rsp
	mov	$-0x11223344, %rbp
	mov	$-0x11223344, %rsi
	mov	$-0x11223344, %rdi
	mov	$-0x11223344, %r8
	mov	$-0x11223344, %r9
	mov	$-0x11223344, %r10
	mov	$-0x11223344, %r11
	mov	$-0x11223344, %r12
	mov	$-0x11223344, %r13
	mov	$-0x11223344, %r14
	mov	$-0x11223344, %r15
	ret
	.cfi_endproc

	.p2align 4,,15
	.globl	Load_u32
	.type	Load_u32, @function
Load_u32:
	.cfi_startproc
	movl	$0xffaa9988, %eax
	movl	$0xffaa9988, %ecx
	movl	$0xffaa9988, %edx
	movl	$0xffaa9988, %ebx
	movl	$0xffaa9988, %esp
	movl	$0xffaa9988, %ebp
	movl	$0xffaa9988, %esi
	movl	$0xffaa9988, %edi
	movl	$0xffaa9988, %r8d
	movl	$0xffaa9988, %r9d
	movl	$0xffaa9988, %r10d
	movl	$0xffaa9988, %r11d
	movl	$0xffaa9988, %r12d
	movl	$0xffaa9988, %r13d
	movl	$0xffaa9988, %r14d
	movl	$0xffaa9988, %r15d
	ret
	.cfi_endproc



	.p2align 4,,15
	.globl	Load_64
	.type	Load_64, @function
Load_64:
	.cfi_startproc
	movabs	$0x5566778899aabbcc, %rax
	movabs	$0x5566778899aabbcc, %rcx
	movabs	$0x5566778899aabbcc, %rdx
	movabs	$0x5566778899aabbcc, %rbx
	movabs	$0x5566778899aabbcc, %rsp
	movabs	$0x5566778899aabbcc, %rbp
	movabs	$0x5566778899aabbcc, %rsi
	movabs	$0x5566778899aabbcc, %rdi
	movabs	$0x5566778899aabbcc, %r8
	movabs	$0x5566778899aabbcc, %r9
	movabs	$0x5566778899aabbcc, %r10
	movabs	$0x5566778899aabbcc, %r11
	movabs	$0x5566778899aabbcc, %r12
	movabs	$0x5566778899aabbcc, %r13
	movabs	$0x5566778899aabbcc, %r14
	movabs	$0x5566778899aabbcc, %r15
	ret
	.cfi_endproc
