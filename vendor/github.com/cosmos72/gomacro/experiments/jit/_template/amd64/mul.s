	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	mul
	.type	mul, @function
mul:
	.cfi_startproc
	imul	%rax,%rax
	imul	%rax,%rcx
	imul	%rax,%rdx
	imul	%rax,%rbx
	imul	%rax,%rsp
	imul	%rax,%rbp
	imul	%rax,%rsi
	imul	%rax,%rdi
	imul	%rax,%r8
	imul	%rax,%r9
	imul	%rax,%r10
	imul	%rax,%r11
	imul	%rax,%r12
	imul	%rax,%r13
	imul	%rax,%r14
	imul	%rax,%r15
	nop
	imul	%rcx,%rax
	imul	%rcx,%rcx
	imul	%rcx,%rdx
	imul	%rcx,%rbx
	imul	%rcx,%rsp
	imul	%rcx,%rbp
	imul	%rcx,%rsi
	imul	%rcx,%rdi
	imul	%rcx,%r8
	imul	%rcx,%r9
	imul	%rcx,%r10
	imul	%rcx,%r11
	imul	%rcx,%r12
	imul	%rcx,%r13
	imul	%rcx,%r14
	imul	%rcx,%r15
	nop
	imul	%rdx,%rax
	imul	%rdx,%rcx
	imul	%rdx,%rdx
	imul	%rdx,%rbx
	imul	%rdx,%rsp
	imul	%rdx,%rbp
	imul	%rdx,%rsi
	imul	%rdx,%rdi
	imul	%rdx,%r8
	imul	%rdx,%r9
	imul	%rdx,%r10
	imul	%rdx,%r11
	imul	%rdx,%r12
	imul	%rdx,%r13
	imul	%rdx,%r14
	imul	%rdx,%r15
	nop
	imul	%rbx,%rax
	imul	%rbx,%rcx
	imul	%rbx,%rdx
	imul	%rbx,%rbx
	imul	%rbx,%rsp
	imul	%rbx,%rbp
	imul	%rbx,%rsi
	imul	%rbx,%rdi
	imul	%rbx,%r8
	imul	%rbx,%r9
	imul	%rbx,%r10
	imul	%rbx,%r11
	imul	%rbx,%r12
	imul	%rbx,%r13
	imul	%rbx,%r14
	imul	%rbx,%r15
	nop
	imul	%rsp,%rax
	imul	%rsp,%rcx
	imul	%rsp,%rdx
	imul	%rsp,%rbx
	imul	%rsp,%rsp
	imul	%rsp,%rbp
	imul	%rsp,%rsi
	imul	%rsp,%rdi
	imul	%rsp,%r8
	imul	%rsp,%r9
	imul	%rsp,%r10
	imul	%rsp,%r11
	imul	%rsp,%r12
	imul	%rsp,%r13
	imul	%rsp,%r14
	imul	%rsp,%r15
	nop
	imul	%rbp,%rax
	imul	%rbp,%rcx
	imul	%rbp,%rdx
	imul	%rbp,%rbx
	imul	%rbp,%rsp
	imul	%rbp,%rbp
	imul	%rbp,%rsi
	imul	%rbp,%rdi
	imul	%rbp,%r8
	imul	%rbp,%r9
	imul	%rbp,%r10
	imul	%rbp,%r11
	imul	%rbp,%r12
	imul	%rbp,%r13
	imul	%rbp,%r14
	imul	%rbp,%r15
	nop
	imul	%rsi,%rax
	imul	%rsi,%rcx
	imul	%rsi,%rdx
	imul	%rsi,%rbx
	imul	%rsi,%rsp
	imul	%rsi,%rbp
	imul	%rsi,%rsi
	imul	%rsi,%rdi
	imul	%rsi,%r8
	imul	%rsi,%r9
	imul	%rsi,%r10
	imul	%rsi,%r11
	imul	%rsi,%r12
	imul	%rsi,%r13
	imul	%rsi,%r14
	imul	%rsi,%r15
	nop
	imul	%rdi,%rax
	imul	%rdi,%rcx
	imul	%rdi,%rdx
	imul	%rdi,%rbx
	imul	%rdi,%rsp
	imul	%rdi,%rbp
	imul	%rdi,%rsi
	imul	%rdi,%rdi
	imul	%rdi,%r8
	imul	%rdi,%r9
	imul	%rdi,%r10
	imul	%rdi,%r11
	imul	%rdi,%r12
	imul	%rdi,%r13
	imul	%rdi,%r14
	imul	%rdi,%r15
	nop
	imul	%r8, %rax
	imul	%r8, %rcx
	imul	%r8, %rdx
	imul	%r8, %rbx
	imul	%r8, %rsp
	imul	%r8, %rbp
	imul	%r8, %rsi
	imul	%r8, %rdi
	imul	%r8, %r8
	imul	%r8, %r9
	imul	%r8, %r10
	imul	%r8, %r11
	imul	%r8, %r12
	imul	%r8, %r13
	imul	%r8, %r14
	imul	%r8, %r15
	nop
	imul	%r12,%rax
	imul	%r12,%rcx
	imul	%r12,%rdx
	imul	%r12,%rbx
	imul	%r12,%rsp
	imul	%r12,%rbp
	imul	%r12,%rsi
	imul	%r12,%rdi
	imul	%r12,%r8
	imul	%r12,%r9
	imul	%r12,%r10
	imul	%r12,%r11
	imul	%r12,%r12
	imul	%r12,%r13
	imul	%r12,%r14
	imul	%r12,%r15
	nop
	imul	%r15,%rax
	imul	%r15,%rcx
	imul	%r15,%rdx
	imul	%r15,%rbx
	imul	%r15,%rsp
	imul	%r15,%rbp
	imul	%r15,%rsi
	imul	%r15,%rdi
	imul	%r15,%r8
	imul	%r15,%r9
	imul	%r15,%r10
	imul	%r15,%r11
	imul	%r15,%r12
	imul	%r15,%r13
	imul	%r15,%r14
	imul	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	mul_s32
	.type	mul_s32, @function
mul_s32:
	.cfi_startproc
	imul	$-0x55667788,%rax
	imul	$-0x55667788,%rcx
	imul	$-0x55667788,%rdx
	imul	$-0x55667788,%rbx
	imul	$-0x55667788,%rsp
	imul	$-0x55667788,%rbp
	imul	$-0x55667788,%rsi
	imul	$-0x55667788,%rdi
	imul	$-0x55667788,%r8
	imul	$-0x55667788,%r9
	imul	$-0x55667788,%r10
	imul	$-0x55667788,%r11
	imul	$-0x55667788,%r12
	imul	$-0x55667788,%r13
	imul	$-0x55667788,%r14
	imul	$-0x55667788,%r15
	.cfi_endproc
