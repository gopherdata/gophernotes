	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	sub
	.type	sub, @function
sub:
	.cfi_startproc
	sub	%rax,%rax
	sub	%rax,%rcx
	sub	%rax,%rdx
	sub	%rax,%rbx
	sub	%rax,%rsp
	sub	%rax,%rbp
	sub	%rax,%rsi
	sub	%rax,%rdi
	sub	%rax,%r8
	sub	%rax,%r9
	sub	%rax,%r10
	sub	%rax,%r11
	sub	%rax,%r12
	sub	%rax,%r13
	sub	%rax,%r14
	sub	%rax,%r15
	nop
	sub	%rcx,%rax
	sub	%rcx,%rcx
	sub	%rcx,%rdx
	sub	%rcx,%rbx
	sub	%rcx,%rsp
	sub	%rcx,%rbp
	sub	%rcx,%rsi
	sub	%rcx,%rdi
	sub	%rcx,%r8
	sub	%rcx,%r9
	sub	%rcx,%r10
	sub	%rcx,%r11
	sub	%rcx,%r12
	sub	%rcx,%r13
	sub	%rcx,%r14
	sub	%rcx,%r15
	nop
	sub	%rdx,%rax
	sub	%rdx,%rcx
	sub	%rdx,%rdx
	sub	%rdx,%rbx
	sub	%rdx,%rsp
	sub	%rdx,%rbp
	sub	%rdx,%rsi
	sub	%rdx,%rdi
	sub	%rdx,%r8
	sub	%rdx,%r9
	sub	%rdx,%r10
	sub	%rdx,%r11
	sub	%rdx,%r12
	sub	%rdx,%r13
	sub	%rdx,%r14
	sub	%rdx,%r15
	nop
	sub	%rbx,%rax
	sub	%rbx,%rcx
	sub	%rbx,%rdx
	sub	%rbx,%rbx
	sub	%rbx,%rsp
	sub	%rbx,%rbp
	sub	%rbx,%rsi
	sub	%rbx,%rdi
	sub	%rbx,%r8
	sub	%rbx,%r9
	sub	%rbx,%r10
	sub	%rbx,%r11
	sub	%rbx,%r12
	sub	%rbx,%r13
	sub	%rbx,%r14
	sub	%rbx,%r15
	nop
	sub	%rsp,%rax
	sub	%rsp,%rcx
	sub	%rsp,%rdx
	sub	%rsp,%rbx
	sub	%rsp,%rsp
	sub	%rsp,%rbp
	sub	%rsp,%rsi
	sub	%rsp,%rdi
	sub	%rsp,%r8
	sub	%rsp,%r9
	sub	%rsp,%r10
	sub	%rsp,%r11
	sub	%rsp,%r12
	sub	%rsp,%r13
	sub	%rsp,%r14
	sub	%rsp,%r15
	nop
	sub	%rbp,%rax
	sub	%rbp,%rcx
	sub	%rbp,%rdx
	sub	%rbp,%rbx
	sub	%rbp,%rsp
	sub	%rbp,%rbp
	sub	%rbp,%rsi
	sub	%rbp,%rdi
	sub	%rbp,%r8
	sub	%rbp,%r9
	sub	%rbp,%r10
	sub	%rbp,%r11
	sub	%rbp,%r12
	sub	%rbp,%r13
	sub	%rbp,%r14
	sub	%rbp,%r15
	nop
	sub	%rsi,%rax
	sub	%rsi,%rcx
	sub	%rsi,%rdx
	sub	%rsi,%rbx
	sub	%rsi,%rsp
	sub	%rsi,%rbp
	sub	%rsi,%rsi
	sub	%rsi,%rdi
	sub	%rsi,%r8
	sub	%rsi,%r9
	sub	%rsi,%r10
	sub	%rsi,%r11
	sub	%rsi,%r12
	sub	%rsi,%r13
	sub	%rsi,%r14
	sub	%rsi,%r15
	nop
	sub	%rdi,%rax
	sub	%rdi,%rcx
	sub	%rdi,%rdx
	sub	%rdi,%rbx
	sub	%rdi,%rsp
	sub	%rdi,%rbp
	sub	%rdi,%rsi
	sub	%rdi,%rdi
	sub	%rdi,%r8
	sub	%rdi,%r9
	sub	%rdi,%r10
	sub	%rdi,%r11
	sub	%rdi,%r12
	sub	%rdi,%r13
	sub	%rdi,%r14
	sub	%rdi,%r15
	nop
	sub	%r8, %rax
	sub	%r8, %rcx
	sub	%r8, %rdx
	sub	%r8, %rbx
	sub	%r8, %rsp
	sub	%r8, %rbp
	sub	%r8, %rsi
	sub	%r8, %rdi
	sub	%r8, %r8
	sub	%r8, %r9
	sub	%r8, %r10
	sub	%r8, %r11
	sub	%r8, %r12
	sub	%r8, %r13
	sub	%r8, %r14
	sub	%r8, %r15
	nop
	sub	%r12,%rax
	sub	%r12,%rcx
	sub	%r12,%rdx
	sub	%r12,%rbx
	sub	%r12,%rsp
	sub	%r12,%rbp
	sub	%r12,%rsi
	sub	%r12,%rdi
	sub	%r12,%r8
	sub	%r12,%r9
	sub	%r12,%r10
	sub	%r12,%r11
	sub	%r12,%r12
	sub	%r12,%r13
	sub	%r12,%r14
	sub	%r12,%r15
	nop
	sub	%r15,%rax
	sub	%r15,%rcx
	sub	%r15,%rdx
	sub	%r15,%rbx
	sub	%r15,%rsp
	sub	%r15,%rbp
	sub	%r15,%rsi
	sub	%r15,%rdi
	sub	%r15,%r8
	sub	%r15,%r9
	sub	%r15,%r10
	sub	%r15,%r11
	sub	%r15,%r12
	sub	%r15,%r13
	sub	%r15,%r14
	sub	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	sub_s32
	.type	sub_s32, @function
sub_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xe8, 0x78, 0x88, 0x99, 0xaa
	// sub	$-0x55667788,%rax
	sub	$-0x55667788,%rcx
	sub	$-0x55667788,%rdx
	sub	$-0x55667788,%rbx
	sub	$-0x55667788,%rsp
	sub	$-0x55667788,%rbp
	sub	$-0x55667788,%rsi
	sub	$-0x55667788,%rdi
	sub	$-0x55667788,%r8
	sub	$-0x55667788,%r9
	sub	$-0x55667788,%r10
	sub	$-0x55667788,%r11
	sub	$-0x55667788,%r12
	sub	$-0x55667788,%r13
	sub	$-0x55667788,%r14
	sub	$-0x55667788,%r15
	.cfi_endproc
