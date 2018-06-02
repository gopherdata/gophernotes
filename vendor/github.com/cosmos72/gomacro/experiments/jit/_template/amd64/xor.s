	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	xor
	.type	xor, @function
xor:
	.cfi_startproc
	xor	%rax,%rax
	xor	%rax,%rcx
	xor	%rax,%rdx
	xor	%rax,%rbx
	xor	%rax,%rsp
	xor	%rax,%rbp
	xor	%rax,%rsi
	xor	%rax,%rdi
	xor	%rax,%r8
	xor	%rax,%r9
	xor	%rax,%r10
	xor	%rax,%r11
	xor	%rax,%r12
	xor	%rax,%r13
	xor	%rax,%r14
	xor	%rax,%r15
	nop
	xor	%rcx,%rax
	xor	%rcx,%rcx
	xor	%rcx,%rdx
	xor	%rcx,%rbx
	xor	%rcx,%rsp
	xor	%rcx,%rbp
	xor	%rcx,%rsi
	xor	%rcx,%rdi
	xor	%rcx,%r8
	xor	%rcx,%r9
	xor	%rcx,%r10
	xor	%rcx,%r11
	xor	%rcx,%r12
	xor	%rcx,%r13
	xor	%rcx,%r14
	xor	%rcx,%r15
	nop
	xor	%rdx,%rax
	xor	%rdx,%rcx
	xor	%rdx,%rdx
	xor	%rdx,%rbx
	xor	%rdx,%rsp
	xor	%rdx,%rbp
	xor	%rdx,%rsi
	xor	%rdx,%rdi
	xor	%rdx,%r8
	xor	%rdx,%r9
	xor	%rdx,%r10
	xor	%rdx,%r11
	xor	%rdx,%r12
	xor	%rdx,%r13
	xor	%rdx,%r14
	xor	%rdx,%r15
	nop
	xor	%rbx,%rax
	xor	%rbx,%rcx
	xor	%rbx,%rdx
	xor	%rbx,%rbx
	xor	%rbx,%rsp
	xor	%rbx,%rbp
	xor	%rbx,%rsi
	xor	%rbx,%rdi
	xor	%rbx,%r8
	xor	%rbx,%r9
	xor	%rbx,%r10
	xor	%rbx,%r11
	xor	%rbx,%r12
	xor	%rbx,%r13
	xor	%rbx,%r14
	xor	%rbx,%r15
	nop
	xor	%rsp,%rax
	xor	%rsp,%rcx
	xor	%rsp,%rdx
	xor	%rsp,%rbx
	xor	%rsp,%rsp
	xor	%rsp,%rbp
	xor	%rsp,%rsi
	xor	%rsp,%rdi
	xor	%rsp,%r8
	xor	%rsp,%r9
	xor	%rsp,%r10
	xor	%rsp,%r11
	xor	%rsp,%r12
	xor	%rsp,%r13
	xor	%rsp,%r14
	xor	%rsp,%r15
	nop
	xor	%rbp,%rax
	xor	%rbp,%rcx
	xor	%rbp,%rdx
	xor	%rbp,%rbx
	xor	%rbp,%rsp
	xor	%rbp,%rbp
	xor	%rbp,%rsi
	xor	%rbp,%rdi
	xor	%rbp,%r8
	xor	%rbp,%r9
	xor	%rbp,%r10
	xor	%rbp,%r11
	xor	%rbp,%r12
	xor	%rbp,%r13
	xor	%rbp,%r14
	xor	%rbp,%r15
	nop
	xor	%rsi,%rax
	xor	%rsi,%rcx
	xor	%rsi,%rdx
	xor	%rsi,%rbx
	xor	%rsi,%rsp
	xor	%rsi,%rbp
	xor	%rsi,%rsi
	xor	%rsi,%rdi
	xor	%rsi,%r8
	xor	%rsi,%r9
	xor	%rsi,%r10
	xor	%rsi,%r11
	xor	%rsi,%r12
	xor	%rsi,%r13
	xor	%rsi,%r14
	xor	%rsi,%r15
	nop
	xor	%rdi,%rax
	xor	%rdi,%rcx
	xor	%rdi,%rdx
	xor	%rdi,%rbx
	xor	%rdi,%rsp
	xor	%rdi,%rbp
	xor	%rdi,%rsi
	xor	%rdi,%rdi
	xor	%rdi,%r8
	xor	%rdi,%r9
	xor	%rdi,%r10
	xor	%rdi,%r11
	xor	%rdi,%r12
	xor	%rdi,%r13
	xor	%rdi,%r14
	xor	%rdi,%r15
	nop
	xor	%r8, %rax
	xor	%r8, %rcx
	xor	%r8, %rdx
	xor	%r8, %rbx
	xor	%r8, %rsp
	xor	%r8, %rbp
	xor	%r8, %rsi
	xor	%r8, %rdi
	xor	%r8, %r8
	xor	%r8, %r9
	xor	%r8, %r10
	xor	%r8, %r11
	xor	%r8, %r12
	xor	%r8, %r13
	xor	%r8, %r14
	xor	%r8, %r15
	nop
	xor	%r12,%rax
	xor	%r12,%rcx
	xor	%r12,%rdx
	xor	%r12,%rbx
	xor	%r12,%rsp
	xor	%r12,%rbp
	xor	%r12,%rsi
	xor	%r12,%rdi
	xor	%r12,%r8
	xor	%r12,%r9
	xor	%r12,%r10
	xor	%r12,%r11
	xor	%r12,%r12
	xor	%r12,%r13
	xor	%r12,%r14
	xor	%r12,%r15
	nop
	xor	%r15,%rax
	xor	%r15,%rcx
	xor	%r15,%rdx
	xor	%r15,%rbx
	xor	%r15,%rsp
	xor	%r15,%rbp
	xor	%r15,%rsi
	xor	%r15,%rdi
	xor	%r15,%r8
	xor	%r15,%r9
	xor	%r15,%r10
	xor	%r15,%r11
	xor	%r15,%r12
	xor	%r15,%r13
	xor	%r15,%r14
	xor	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	xor_s32
	.type	xor_s32, @function
xor_s32:
	.cfi_startproc
	xor	$-0x55667788,%rax
	xor	$-0x55667788,%rcx
	xor	$-0x55667788,%rdx
	xor	$-0x55667788,%rbx
	xor	$-0x55667788,%rsp
	xor	$-0x55667788,%rbp
	xor	$-0x55667788,%rsi
	xor	$-0x55667788,%rdi
	xor	$-0x55667788,%r8
	xor	$-0x55667788,%r9
	xor	$-0x55667788,%r10
	xor	$-0x55667788,%r11
	xor	$-0x55667788,%r12
	xor	$-0x55667788,%r13
	xor	$-0x55667788,%r14
	xor	$-0x55667788,%r15
	.cfi_endproc

