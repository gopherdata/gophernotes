	.file	"movabs.s"
	.text
	
	.p2align 4,,15
	.globl	Movabs
	.type	Movabs, @function
Movabs:
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
