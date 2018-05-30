	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	or
	.type	or, @function
or:
	.cfi_startproc
	or	%rax,%rax
	or	%rax,%rcx
	or	%rax,%rdx
	or	%rax,%rbx
	or	%rax,%rsp
	or	%rax,%rbp
	or	%rax,%rsi
	or	%rax,%rdi
	or	%rax,%r8
	or	%rax,%r9
	or	%rax,%r10
	or	%rax,%r11
	or	%rax,%r12
	or	%rax,%r13
	or	%rax,%r14
	or	%rax,%r15
	nop
	or	%rcx,%rax
	or	%rcx,%rcx
	or	%rcx,%rdx
	or	%rcx,%rbx
	or	%rcx,%rsp
	or	%rcx,%rbp
	or	%rcx,%rsi
	or	%rcx,%rdi
	or	%rcx,%r8
	or	%rcx,%r9
	or	%rcx,%r10
	or	%rcx,%r11
	or	%rcx,%r12
	or	%rcx,%r13
	or	%rcx,%r14
	or	%rcx,%r15
	nop
	or	%rdx,%rax
	or	%rdx,%rcx
	or	%rdx,%rdx
	or	%rdx,%rbx
	or	%rdx,%rsp
	or	%rdx,%rbp
	or	%rdx,%rsi
	or	%rdx,%rdi
	or	%rdx,%r8
	or	%rdx,%r9
	or	%rdx,%r10
	or	%rdx,%r11
	or	%rdx,%r12
	or	%rdx,%r13
	or	%rdx,%r14
	or	%rdx,%r15
	nop
	or	%rbx,%rax
	or	%rbx,%rcx
	or	%rbx,%rdx
	or	%rbx,%rbx
	or	%rbx,%rsp
	or	%rbx,%rbp
	or	%rbx,%rsi
	or	%rbx,%rdi
	or	%rbx,%r8
	or	%rbx,%r9
	or	%rbx,%r10
	or	%rbx,%r11
	or	%rbx,%r12
	or	%rbx,%r13
	or	%rbx,%r14
	or	%rbx,%r15
	nop
	or	%rsp,%rax
	or	%rsp,%rcx
	or	%rsp,%rdx
	or	%rsp,%rbx
	or	%rsp,%rsp
	or	%rsp,%rbp
	or	%rsp,%rsi
	or	%rsp,%rdi
	or	%rsp,%r8
	or	%rsp,%r9
	or	%rsp,%r10
	or	%rsp,%r11
	or	%rsp,%r12
	or	%rsp,%r13
	or	%rsp,%r14
	or	%rsp,%r15
	nop
	or	%rbp,%rax
	or	%rbp,%rcx
	or	%rbp,%rdx
	or	%rbp,%rbx
	or	%rbp,%rsp
	or	%rbp,%rbp
	or	%rbp,%rsi
	or	%rbp,%rdi
	or	%rbp,%r8
	or	%rbp,%r9
	or	%rbp,%r10
	or	%rbp,%r11
	or	%rbp,%r12
	or	%rbp,%r13
	or	%rbp,%r14
	or	%rbp,%r15
	nop
	or	%rsi,%rax
	or	%rsi,%rcx
	or	%rsi,%rdx
	or	%rsi,%rbx
	or	%rsi,%rsp
	or	%rsi,%rbp
	or	%rsi,%rsi
	or	%rsi,%rdi
	or	%rsi,%r8
	or	%rsi,%r9
	or	%rsi,%r10
	or	%rsi,%r11
	or	%rsi,%r12
	or	%rsi,%r13
	or	%rsi,%r14
	or	%rsi,%r15
	nop
	or	%rdi,%rax
	or	%rdi,%rcx
	or	%rdi,%rdx
	or	%rdi,%rbx
	or	%rdi,%rsp
	or	%rdi,%rbp
	or	%rdi,%rsi
	or	%rdi,%rdi
	or	%rdi,%r8
	or	%rdi,%r9
	or	%rdi,%r10
	or	%rdi,%r11
	or	%rdi,%r12
	or	%rdi,%r13
	or	%rdi,%r14
	or	%rdi,%r15
	nop
	or	%r8, %rax
	or	%r8, %rcx
	or	%r8, %rdx
	or	%r8, %rbx
	or	%r8, %rsp
	or	%r8, %rbp
	or	%r8, %rsi
	or	%r8, %rdi
	or	%r8, %r8
	or	%r8, %r9
	or	%r8, %r10
	or	%r8, %r11
	or	%r8, %r12
	or	%r8, %r13
	or	%r8, %r14
	or	%r8, %r15
	nop
	or	%r12,%rax
	or	%r12,%rcx
	or	%r12,%rdx
	or	%r12,%rbx
	or	%r12,%rsp
	or	%r12,%rbp
	or	%r12,%rsi
	or	%r12,%rdi
	or	%r12,%r8
	or	%r12,%r9
	or	%r12,%r10
	or	%r12,%r11
	or	%r12,%r12
	or	%r12,%r13
	or	%r12,%r14
	or	%r12,%r15
	nop
	or	%r15,%rax
	or	%r15,%rcx
	or	%r15,%rdx
	or	%r15,%rbx
	or	%r15,%rsp
	or	%r15,%rbp
	or	%r15,%rsi
	or	%r15,%rdi
	or	%r15,%r8
	or	%r15,%r9
	or	%r15,%r10
	or	%r15,%r11
	or	%r15,%r12
	or	%r15,%r13
	or	%r15,%r14
	or	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	or_s32
	.type	or_s32, @function
or_s32:
	.cfi_startproc
	or	$-0x55667788,%rax
	or	$-0x55667788,%rcx
	or	$-0x55667788,%rdx
	or	$-0x55667788,%rbx
	or	$-0x55667788,%rsp
	or	$-0x55667788,%rbp
	or	$-0x55667788,%rsi
	or	$-0x55667788,%rdi
	or	$-0x55667788,%r8
	or	$-0x55667788,%r9
	or	$-0x55667788,%r10
	or	$-0x55667788,%r11
	or	$-0x55667788,%r12
	or	$-0x55667788,%r13
	or	$-0x55667788,%r14
	or	$-0x55667788,%r15
	.cfi_endproc

