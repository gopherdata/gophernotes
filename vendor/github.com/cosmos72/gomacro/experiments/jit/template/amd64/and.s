	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	and
	.type	and, @function
and:
	.cfi_startproc
	and	%rax,%rax
	and	%rax,%rcx
	and	%rax,%rdx
	and	%rax,%rbx
	and	%rax,%rsp
	and	%rax,%rbp
	and	%rax,%rsi
	and	%rax,%rdi
	and	%rax,%r8
	and	%rax,%r9
	and	%rax,%r10
	and	%rax,%r11
	and	%rax,%r12
	and	%rax,%r13
	and	%rax,%r14
	and	%rax,%r15
	nop
	and	%rcx,%rax
	and	%rcx,%rcx
	and	%rcx,%rdx
	and	%rcx,%rbx
	and	%rcx,%rsp
	and	%rcx,%rbp
	and	%rcx,%rsi
	and	%rcx,%rdi
	and	%rcx,%r8
	and	%rcx,%r9
	and	%rcx,%r10
	and	%rcx,%r11
	and	%rcx,%r12
	and	%rcx,%r13
	and	%rcx,%r14
	and	%rcx,%r15
	nop
	and	%rdx,%rax
	and	%rdx,%rcx
	and	%rdx,%rdx
	and	%rdx,%rbx
	and	%rdx,%rsp
	and	%rdx,%rbp
	and	%rdx,%rsi
	and	%rdx,%rdi
	and	%rdx,%r8
	and	%rdx,%r9
	and	%rdx,%r10
	and	%rdx,%r11
	and	%rdx,%r12
	and	%rdx,%r13
	and	%rdx,%r14
	and	%rdx,%r15
	nop
	and	%rbx,%rax
	and	%rbx,%rcx
	and	%rbx,%rdx
	and	%rbx,%rbx
	and	%rbx,%rsp
	and	%rbx,%rbp
	and	%rbx,%rsi
	and	%rbx,%rdi
	and	%rbx,%r8
	and	%rbx,%r9
	and	%rbx,%r10
	and	%rbx,%r11
	and	%rbx,%r12
	and	%rbx,%r13
	and	%rbx,%r14
	and	%rbx,%r15
	nop
	and	%rsp,%rax
	and	%rsp,%rcx
	and	%rsp,%rdx
	and	%rsp,%rbx
	and	%rsp,%rsp
	and	%rsp,%rbp
	and	%rsp,%rsi
	and	%rsp,%rdi
	and	%rsp,%r8
	and	%rsp,%r9
	and	%rsp,%r10
	and	%rsp,%r11
	and	%rsp,%r12
	and	%rsp,%r13
	and	%rsp,%r14
	and	%rsp,%r15
	nop
	and	%rbp,%rax
	and	%rbp,%rcx
	and	%rbp,%rdx
	and	%rbp,%rbx
	and	%rbp,%rsp
	and	%rbp,%rbp
	and	%rbp,%rsi
	and	%rbp,%rdi
	and	%rbp,%r8
	and	%rbp,%r9
	and	%rbp,%r10
	and	%rbp,%r11
	and	%rbp,%r12
	and	%rbp,%r13
	and	%rbp,%r14
	and	%rbp,%r15
	nop
	and	%rsi,%rax
	and	%rsi,%rcx
	and	%rsi,%rdx
	and	%rsi,%rbx
	and	%rsi,%rsp
	and	%rsi,%rbp
	and	%rsi,%rsi
	and	%rsi,%rdi
	and	%rsi,%r8
	and	%rsi,%r9
	and	%rsi,%r10
	and	%rsi,%r11
	and	%rsi,%r12
	and	%rsi,%r13
	and	%rsi,%r14
	and	%rsi,%r15
	nop
	and	%rdi,%rax
	and	%rdi,%rcx
	and	%rdi,%rdx
	and	%rdi,%rbx
	and	%rdi,%rsp
	and	%rdi,%rbp
	and	%rdi,%rsi
	and	%rdi,%rdi
	and	%rdi,%r8
	and	%rdi,%r9
	and	%rdi,%r10
	and	%rdi,%r11
	and	%rdi,%r12
	and	%rdi,%r13
	and	%rdi,%r14
	and	%rdi,%r15
	nop
	and	%r8, %rax
	and	%r8, %rcx
	and	%r8, %rdx
	and	%r8, %rbx
	and	%r8, %rsp
	and	%r8, %rbp
	and	%r8, %rsi
	and	%r8, %rdi
	and	%r8, %r8
	and	%r8, %r9
	and	%r8, %r10
	and	%r8, %r11
	and	%r8, %r12
	and	%r8, %r13
	and	%r8, %r14
	and	%r8, %r15
	nop
	and	%r12,%rax
	and	%r12,%rcx
	and	%r12,%rdx
	and	%r12,%rbx
	and	%r12,%rsp
	and	%r12,%rbp
	and	%r12,%rsi
	and	%r12,%rdi
	and	%r12,%r8
	and	%r12,%r9
	and	%r12,%r10
	and	%r12,%r11
	and	%r12,%r12
	and	%r12,%r13
	and	%r12,%r14
	and	%r12,%r15
	nop
	and	%r15,%rax
	and	%r15,%rcx
	and	%r15,%rdx
	and	%r15,%rbx
	and	%r15,%rsp
	and	%r15,%rbp
	and	%r15,%rsi
	and	%r15,%rdi
	and	%r15,%r8
	and	%r15,%r9
	and	%r15,%r10
	and	%r15,%r11
	and	%r15,%r12
	and	%r15,%r13
	and	%r15,%r14
	and	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	and_u32
	.type	and_u32, @function
and_u32:
	.cfi_startproc
	and	$0x55667788,%eax
	and	$0x55667788,%ecx
	and	$0x55667788,%edx
	and	$0x55667788,%ebx
	and	$0x55667788,%esp
	and	$0x55667788,%ebp
	and	$0x55667788,%esi
	and	$0x55667788,%edi
	and	$0x55667788,%r8d
	and	$0x55667788,%r9d
	and	$0x55667788,%r10d
	and	$0x55667788,%r11d
	and	$0x55667788,%r12d
	and	$0x55667788,%r13d
	and	$0x55667788,%r14d
	and	$0x55667788,%r15d
	.cfi_endproc

	.p2align 4,,15
	.globl	and_s32
	.type	and_s32, @function
and_s32:
	.cfi_startproc
	and	$-0x55667788,%rax
	and	$-0x55667788,%rcx
	and	$-0x55667788,%rdx
	and	$-0x55667788,%rbx
	and	$-0x55667788,%rsp
	and	$-0x55667788,%rbp
	and	$-0x55667788,%rsi
	and	$-0x55667788,%rdi
	and	$-0x55667788,%r8
	and	$-0x55667788,%r9
	and	$-0x55667788,%r10
	and	$-0x55667788,%r11
	and	$-0x55667788,%r12
	and	$-0x55667788,%r13
	and	$-0x55667788,%r14
	and	$-0x55667788,%r15
	.cfi_endproc

