	.file	"arith.s"
	.text
	.p2align 4,,15
	.globl	Add
	.type	Add, @function
Add:
	.cfi_startproc
	add	%rax,%rax
	add	%rax,%rcx
	add	%rax,%rdx
	add	%rax,%rbx
	add	%rax,%rsp
	add	%rax,%rbp
	add	%rax,%rsi
	add	%rax,%rdi
	add	%rax,%r8
	add	%rax,%r9
	add	%rax,%r10
	add	%rax,%r11
	add	%rax,%r12
	add	%rax,%r13
	add	%rax,%r14
	add	%rax,%r15
	nop
	add	%rcx,%rax
	add	%rcx,%rcx
	add	%rcx,%rdx
	add	%rcx,%rbx
	add	%rcx,%rsp
	add	%rcx,%rbp
	add	%rcx,%rsi
	add	%rcx,%rdi
	add	%rcx,%r8
	add	%rcx,%r9
	add	%rcx,%r10
	add	%rcx,%r11
	add	%rcx,%r12
	add	%rcx,%r13
	add	%rcx,%r14
	add	%rcx,%r15
	nop
	add	%rdx,%rax
	add	%rdx,%rcx
	add	%rdx,%rdx
	add	%rdx,%rbx
	add	%rdx,%rsp
	add	%rdx,%rbp
	add	%rdx,%rsi
	add	%rdx,%rdi
	add	%rdx,%r8
	add	%rdx,%r9
	add	%rdx,%r10
	add	%rdx,%r11
	add	%rdx,%r12
	add	%rdx,%r13
	add	%rdx,%r14
	add	%rdx,%r15
	nop
	add	%rbx,%rax
	add	%rbx,%rcx
	add	%rbx,%rdx
	add	%rbx,%rbx
	add	%rbx,%rsp
	add	%rbx,%rbp
	add	%rbx,%rsi
	add	%rbx,%rdi
	add	%rbx,%r8
	add	%rbx,%r9
	add	%rbx,%r10
	add	%rbx,%r11
	add	%rbx,%r12
	add	%rbx,%r13
	add	%rbx,%r14
	add	%rbx,%r15
	nop
	add	%rsp,%rax
	add	%rsp,%rcx
	add	%rsp,%rdx
	add	%rsp,%rbx
	add	%rsp,%rsp
	add	%rsp,%rbp
	add	%rsp,%rsi
	add	%rsp,%rdi
	add	%rsp,%r8
	add	%rsp,%r9
	add	%rsp,%r10
	add	%rsp,%r11
	add	%rsp,%r12
	add	%rsp,%r13
	add	%rsp,%r14
	add	%rsp,%r15
	nop
	add	%rbp,%rax
	add	%rbp,%rcx
	add	%rbp,%rdx
	add	%rbp,%rbx
	add	%rbp,%rsp
	add	%rbp,%rbp
	add	%rbp,%rsi
	add	%rbp,%rdi
	add	%rbp,%r8
	add	%rbp,%r9
	add	%rbp,%r10
	add	%rbp,%r11
	add	%rbp,%r12
	add	%rbp,%r13
	add	%rbp,%r14
	add	%rbp,%r15
	nop
	add	%rsi,%rax
	add	%rsi,%rcx
	add	%rsi,%rdx
	add	%rsi,%rbx
	add	%rsi,%rsp
	add	%rsi,%rbp
	add	%rsi,%rsi
	add	%rsi,%rdi
	add	%rsi,%r8
	add	%rsi,%r9
	add	%rsi,%r10
	add	%rsi,%r11
	add	%rsi,%r12
	add	%rsi,%r13
	add	%rsi,%r14
	add	%rsi,%r15
	nop
	add	%rdi,%rax
	add	%rdi,%rcx
	add	%rdi,%rdx
	add	%rdi,%rbx
	add	%rdi,%rsp
	add	%rdi,%rbp
	add	%rdi,%rsi
	add	%rdi,%rdi
	add	%rdi,%r8
	add	%rdi,%r9
	add	%rdi,%r10
	add	%rdi,%r11
	add	%rdi,%r12
	add	%rdi,%r13
	add	%rdi,%r14
	add	%rdi,%r15
	nop
	add	%r8, %rax
	add	%r8, %rcx
	add	%r8, %rdx
	add	%r8, %rbx
	add	%r8, %rsp
	add	%r8, %rbp
	add	%r8, %rsi
	add	%r8, %rdi
	add	%r8, %r8
	add	%r8, %r9
	add	%r8, %r10
	add	%r8, %r11
	add	%r8, %r12
	add	%r8, %r13
	add	%r8, %r14
	add	%r8, %r15
	nop
	add	%r12,%rax
	add	%r12,%rcx
	add	%r12,%rdx
	add	%r12,%rbx
	add	%r12,%rsp
	add	%r12,%rbp
	add	%r12,%rsi
	add	%r12,%rdi
	add	%r12,%r8
	add	%r12,%r9
	add	%r12,%r10
	add	%r12,%r11
	add	%r12,%r12
	add	%r12,%r13
	add	%r12,%r14
	add	%r12,%r15
	nop
	add	%r15,%rax
	add	%r15,%rcx
	add	%r15,%rdx
	add	%r15,%rbx
	add	%r15,%rsp
	add	%r15,%rbp
	add	%r15,%rsi
	add	%r15,%rdi
	add	%r15,%r8
	add	%r15,%r9
	add	%r15,%r10
	add	%r15,%r11
	add	%r15,%r12
	add	%r15,%r13
	add	%r15,%r14
	add	%r15,%r15
	
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	Add_s32
	.type	Add_s32, @function
Add_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// add	$-0x55667788,%rax
	add	$-0x55667788,%rcx
	add	$-0x55667788,%rdx
	add	$-0x55667788,%rbx
	add	$-0x55667788,%rsp
	add	$-0x55667788,%rbp
	add	$-0x55667788,%rsi
	add	$-0x55667788,%rdi
	add	$-0x55667788,%r8
	add	$-0x55667788,%r9
	add	$-0x55667788,%r10
	add	$-0x55667788,%r11
	add	$-0x55667788,%r12
	add	$-0x55667788,%r13
	add	$-0x55667788,%r14
	add	$-0x55667788,%r15
	.cfi_endproc
