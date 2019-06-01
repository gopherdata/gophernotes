	.file	"arith.s"
	.text
/*
	.p2align 4,,15
	.globl	Movzdq
	.type	Movzdq, @function
Movzdq:
	.cfi_startproc
	// reg -> reg
	movzdq	%eax,%rax
	movzdq	%eax,%rcx
	movzdq	%eax,%rdx
	movzdq	%eax,%rbx
	movzdq	%eax,%rsp
	movzdq	%eax,%rbp
	movzdq	%eax,%rsi
	movzdq	%eax,%rdi
	movzdq	%eax,%r8
	movzdq	%eax,%r9
	movzdq	%eax,%r10
	movzdq	%eax,%r11
	movzdq	%eax,%r12
	movzdq	%eax,%r13
	movzdq	%eax,%r14
	movzdq	%eax,%r15
	nop
	movzdq	%ecx,%rax
	movzdq	%ecx,%rcx
	movzdq	%ecx,%rdx
	movzdq	%ecx,%rbx
	movzdq	%ecx,%rsp
	movzdq	%ecx,%rbp
	movzdq	%ecx,%rsi
	movzdq	%ecx,%rdi
	movzdq	%ecx,%r8
	movzdq	%ecx,%r9
	movzdq	%ecx,%r10
	movzdq	%ecx,%r11
	movzdq	%ecx,%r12
	movzdq	%ecx,%r13
	movzdq	%ecx,%r14
	movzdq	%ecx,%r15
	nop
	movzdq	%edx,%rax
	movzdq	%edx,%rcx
	movzdq	%edx,%rdx
	movzdq	%edx,%rbx
	movzdq	%edx,%rsp
	movzdq	%edx,%rbp
	movzdq	%edx,%rsi
	movzdq	%edx,%rdi
	movzdq	%edx,%r8
	movzdq	%edx,%r9
	movzdq	%edx,%r10
	movzdq	%edx,%r11
	movzdq	%edx,%r12
	movzdq	%edx,%r13
	movzdq	%edx,%r14
	movzdq	%edx,%r15
	nop
	movzdq	%ebx,%rax
	movzdq	%ebx,%rcx
	movzdq	%ebx,%rdx
	movzdq	%ebx,%rbx
	movzdq	%ebx,%rsp
	movzdq	%ebx,%rbp
	movzdq	%ebx,%rsi
	movzdq	%ebx,%rdi
	movzdq	%ebx,%r8
	movzdq	%ebx,%r9
	movzdq	%ebx,%r10
	movzdq	%ebx,%r11
	movzdq	%ebx,%r12
	movzdq	%ebx,%r13
	movzdq	%ebx,%r14
	movzdq	%ebx,%r15
	nop
	movzdq	%esp,%rax
	movzdq	%esp,%rcx
	movzdq	%esp,%rdx
	movzdq	%esp,%rbx
	movzdq	%esp,%rsp
	movzdq	%esp,%rbp
	movzdq	%esp,%rsi
	movzdq	%esp,%rdi
	movzdq	%esp,%r8
	movzdq	%esp,%r9
	movzdq	%esp,%r10
	movzdq	%esp,%r11
	movzdq	%esp,%r12
	movzdq	%esp,%r13
	movzdq	%esp,%r14
	movzdq	%esp,%r15
	nop
	movzdq	%ebp,%rax
	movzdq	%ebp,%rcx
	movzdq	%ebp,%rdx
	movzdq	%ebp,%rbx
	movzdq	%ebp,%rsp
	movzdq	%ebp,%rbp
	movzdq	%ebp,%rsi
	movzdq	%ebp,%rdi
	movzdq	%ebp,%r8
	movzdq	%ebp,%r9
	movzdq	%ebp,%r10
	movzdq	%ebp,%r11
	movzdq	%ebp,%r12
	movzdq	%ebp,%r13
	movzdq	%ebp,%r14
	movzdq	%ebp,%r15
	nop
	movzdq	%esi,%rax
	movzdq	%esi,%rcx
	movzdq	%esi,%rdx
	movzdq	%esi,%rbx
	movzdq	%esi,%rsp
	movzdq	%esi,%rbp
	movzdq	%esi,%rsi
	movzdq	%esi,%rdi
	movzdq	%esi,%r8
	movzdq	%esi,%r9
	movzdq	%esi,%r10
	movzdq	%esi,%r11
	movzdq	%esi,%r12
	movzdq	%esi,%r13
	movzdq	%esi,%r14
	movzdq	%esi,%r15
	nop
	movzdq	%edi,%rax
	movzdq	%edi,%rcx
	movzdq	%edi,%rdx
	movzdq	%edi,%rbx
	movzdq	%edi,%rsp
	movzdq	%edi,%rbp
	movzdq	%edi,%rsi
	movzdq	%edi,%rdi
	movzdq	%edi,%r8
	movzdq	%edi,%r9
	movzdq	%edi,%r10
	movzdq	%edi,%r11
	movzdq	%edi,%r12
	movzdq	%edi,%r13
	movzdq	%edi,%r14
	movzdq	%edi,%r15
	nop
	movzdq	%r8d, %rax
	movzdq	%r8d, %rcx
	movzdq	%r8d, %rdx
	movzdq	%r8d, %rbx
	movzdq	%r8d, %rsp
	movzdq	%r8d, %rbp
	movzdq	%r8d, %rsi
	movzdq	%r8d, %rdi
	movzdq	%r8d, %r8
	movzdq	%r8d, %r9
	movzdq	%r8d, %r10
	movzdq	%r8d, %r11
	movzdq	%r8d, %r12
	movzdq	%r8d, %r13
	movzdq	%r8d, %r14
	movzdq	%r8d, %r15
	nop
	movzdq	%r9d, %rax
	movzdq	%r9d, %rcx
	movzdq	%r9d, %rdx
	movzdq	%r9d, %rbx
	movzdq	%r9d, %rsp
	movzdq	%r9d, %rbp
	movzdq	%r9d, %rsi
	movzdq	%r9d, %rdi
	movzdq	%r9d, %r8
	movzdq	%r9d, %r9
	movzdq	%r9d, %r10
	movzdq	%r9d, %r11
	movzdq	%r9d, %r12
	movzdq	%r9d, %r13
	movzdq	%r9d, %r14
	movzdq	%r9d, %r15
	nop
	movzdq	%r10d,%rax
	movzdq	%r10d,%rcx
	movzdq	%r10d,%rdx
	movzdq	%r10d,%rbx
	movzdq	%r10d,%rsp
	movzdq	%r10d,%rbp
	movzdq	%r10d,%rsi
	movzdq	%r10d,%rdi
	movzdq	%r10d,%r8
	movzdq	%r10d,%r9
	movzdq	%r10d,%r10
	movzdq	%r10d,%r11
	movzdq	%r10d,%r12
	movzdq	%r10d,%r13
	movzdq	%r10d,%r14
	movzdq	%r10d,%r15
	nop
	movzdq	%r11d,%rax
	movzdq	%r11d,%rcx
	movzdq	%r11d,%rdx
	movzdq	%r11d,%rbx
	movzdq	%r11d,%rsp
	movzdq	%r11d,%rbp
	movzdq	%r11d,%rsi
	movzdq	%r11d,%rdi
	movzdq	%r11d,%r8
	movzdq	%r11d,%r9
	movzdq	%r11d,%r10
	movzdq	%r11d,%r11
	movzdq	%r11d,%r12
	movzdq	%r11d,%r13
	movzdq	%r11d,%r14
	movzdq	%r11d,%r15
	nop
	movzdq	%r12d,%rax
	movzdq	%r12d,%rcx
	movzdq	%r12d,%rdx
	movzdq	%r12d,%rbx
	movzdq	%r12d,%rsp
	movzdq	%r12d,%rbp
	movzdq	%r12d,%rsi
	movzdq	%r12d,%rdi
	movzdq	%r12d,%r8
	movzdq	%r12d,%r9
	movzdq	%r12d,%r10
	movzdq	%r12d,%r11
	movzdq	%r12d,%r12
	movzdq	%r12d,%r13
	movzdq	%r12d,%r14
	movzdq	%r12d,%r15
	nop
	movzdq	%r13d,%rax
	movzdq	%r13d,%rcx
	movzdq	%r13d,%rdx
	movzdq	%r13d,%rbx
	movzdq	%r13d,%rsp
	movzdq	%r13d,%rbp
	movzdq	%r13d,%rsi
	movzdq	%r13d,%rdi
	movzdq	%r13d,%r8
	movzdq	%r13d,%r9
	movzdq	%r13d,%r10
	movzdq	%r13d,%r11
	movzdq	%r13d,%r12
	movzdq	%r13d,%r13
	movzdq	%r13d,%r14
	movzdq	%r13d,%r15
	nop
	movzdq	%r14d,%rax
	movzdq	%r14d,%rcx
	movzdq	%r14d,%rdx
	movzdq	%r14d,%rbx
	movzdq	%r14d,%rsp
	movzdq	%r14d,%rbp
	movzdq	%r14d,%rsi
	movzdq	%r14d,%rdi
	movzdq	%r14d,%r8
	movzdq	%r14d,%r9
	movzdq	%r14d,%r10
	movzdq	%r14d,%r11
	movzdq	%r14d,%r12
	movzdq	%r14d,%r13
	movzdq	%r14d,%r14
	movzdq	%r14d,%r15
	nop
	movzdq	%r15d,%rax
	movzdq	%r15d,%rcx
	movzdq	%r15d,%rdx
	movzdq	%r15d,%rbx
	movzdq	%r15d,%rsp
	movzdq	%r15d,%rbp
	movzdq	%r15d,%rsi
	movzdq	%r15d,%rdi
	movzdq	%r15d,%r8
	movzdq	%r15d,%r9
	movzdq	%r15d,%r10
	movzdq	%r15d,%r11
	movzdq	%r15d,%r12
	movzdq	%r15d,%r13
	movzdq	%r15d,%r14
	movzdq	%r15d,%r15
	nop
        nop
	// (reg) -> reg
	movzdq	(%rax),%rax
	movzdq	(%rax),%rcx
	movzdq	(%rax),%rdx
	movzdq	(%rax),%rbx
	movzdq	(%rax),%rsp
	movzdq	(%rax),%rbp
	movzdq	(%rax),%rsi
	movzdq	(%rax),%rdi
	movzdq	(%rax),%r8
	movzdq	(%rax),%r9
	movzdq	(%rax),%r10
	movzdq	(%rax),%r11
	movzdq	(%rax),%r12
	movzdq	(%rax),%r13
	movzdq	(%rax),%r14
	movzdq	(%rax),%r15
	nop
	movzdq	(%rcx),%rax
	movzdq	(%rcx),%rcx
	movzdq	(%rcx),%rdx
	movzdq	(%rcx),%rbx
	movzdq	(%rcx),%rsp
	movzdq	(%rcx),%rbp
	movzdq	(%rcx),%rsi
	movzdq	(%rcx),%rdi
	movzdq	(%rcx),%r8
	movzdq	(%rcx),%r9
	movzdq	(%rcx),%r10
	movzdq	(%rcx),%r11
	movzdq	(%rcx),%r12
	movzdq	(%rcx),%r13
	movzdq	(%rcx),%r14
	movzdq	(%rcx),%r15
	nop
	movzdq	(%rdx),%rax
	movzdq	(%rdx),%rcx
	movzdq	(%rdx),%rdx
	movzdq	(%rdx),%rbx
	movzdq	(%rdx),%rsp
	movzdq	(%rdx),%rbp
	movzdq	(%rdx),%rsi
	movzdq	(%rdx),%rdi
	movzdq	(%rdx),%r8
	movzdq	(%rdx),%r9
	movzdq	(%rdx),%r10
	movzdq	(%rdx),%r11
	movzdq	(%rdx),%r12
	movzdq	(%rdx),%r13
	movzdq	(%rdx),%r14
	movzdq	(%rdx),%r15
	nop
	movzdq	(%rbx),%rax
	movzdq	(%rbx),%rcx
	movzdq	(%rbx),%rdx
	movzdq	(%rbx),%rbx
	movzdq	(%rbx),%rsp
	movzdq	(%rbx),%rbp
	movzdq	(%rbx),%rsi
	movzdq	(%rbx),%rdi
	movzdq	(%rbx),%r8
	movzdq	(%rbx),%r9
	movzdq	(%rbx),%r10
	movzdq	(%rbx),%r11
	movzdq	(%rbx),%r12
	movzdq	(%rbx),%r13
	movzdq	(%rbx),%r14
	movzdq	(%rbx),%r15
	nop
	movzdq	(%rsp),%rax
	movzdq	(%rsp),%rcx
	movzdq	(%rsp),%rdx
	movzdq	(%rsp),%rbx
	movzdq	(%rsp),%rsp
	movzdq	(%rsp),%rbp
	movzdq	(%rsp),%rsi
	movzdq	(%rsp),%rdi
	movzdq	(%rsp),%r8
	movzdq	(%rsp),%r9
	movzdq	(%rsp),%r10
	movzdq	(%rsp),%r11
	movzdq	(%rsp),%r12
	movzdq	(%rsp),%r13
	movzdq	(%rsp),%r14
	movzdq	(%rsp),%r15
	nop
	movzdq	(%rbp),%rax
	movzdq	(%rbp),%rcx
	movzdq	(%rbp),%rdx
	movzdq	(%rbp),%rbx
	movzdq	(%rbp),%rsp
	movzdq	(%rbp),%rbp
	movzdq	(%rbp),%rsi
	movzdq	(%rbp),%rdi
	movzdq	(%rbp),%r8
	movzdq	(%rbp),%r9
	movzdq	(%rbp),%r10
	movzdq	(%rbp),%r11
	movzdq	(%rbp),%r12
	movzdq	(%rbp),%r13
	movzdq	(%rbp),%r14
	movzdq	(%rbp),%r15
	nop
	movzdq	(%rsi),%rax
	movzdq	(%rsi),%rcx
	movzdq	(%rsi),%rdx
	movzdq	(%rsi),%rbx
	movzdq	(%rsi),%rsp
	movzdq	(%rsi),%rbp
	movzdq	(%rsi),%rsi
	movzdq	(%rsi),%rdi
	movzdq	(%rsi),%r8
	movzdq	(%rsi),%r9
	movzdq	(%rsi),%r10
	movzdq	(%rsi),%r11
	movzdq	(%rsi),%r12
	movzdq	(%rsi),%r13
	movzdq	(%rsi),%r14
	movzdq	(%rsi),%r15
	nop
	movzdq	(%rdi),%rax
	movzdq	(%rdi),%rcx
	movzdq	(%rdi),%rdx
	movzdq	(%rdi),%rbx
	movzdq	(%rdi),%rsp
	movzdq	(%rdi),%rbp
	movzdq	(%rdi),%rsi
	movzdq	(%rdi),%rdi
	movzdq	(%rdi),%r8
	movzdq	(%rdi),%r9
	movzdq	(%rdi),%r10
	movzdq	(%rdi),%r11
	movzdq	(%rdi),%r12
	movzdq	(%rdi),%r13
	movzdq	(%rdi),%r14
	movzdq	(%rdi),%r15
	nop
	movzdq	(%r8), %rax
	movzdq	(%r8), %rcx
	movzdq	(%r8), %rdx
	movzdq	(%r8), %rbx
	movzdq	(%r8), %rsp
	movzdq	(%r8), %rbp
	movzdq	(%r8), %rsi
	movzdq	(%r8), %rdi
	movzdq	(%r8), %r8
	movzdq	(%r8), %r9
	movzdq	(%r8), %r10
	movzdq	(%r8), %r11
	movzdq	(%r8), %r12
	movzdq	(%r8), %r13
	movzdq	(%r8), %r14
	movzdq	(%r8), %r15
	nop
	movzdq	(%r9), %rax
	movzdq	(%r9), %rcx
	movzdq	(%r9), %rdx
	movzdq	(%r9), %rbx
	movzdq	(%r9), %rsp
	movzdq	(%r9), %rbp
	movzdq	(%r9), %rsi
	movzdq	(%r9), %rdi
	movzdq	(%r9), %r8
	movzdq	(%r9), %r9
	movzdq	(%r9), %r10
	movzdq	(%r9), %r11
	movzdq	(%r9), %r12
	movzdq	(%r9), %r13
	movzdq	(%r9), %r14
	movzdq	(%r9), %r15
	nop
	movzdq	(%r10),%rax
	movzdq	(%r10),%rcx
	movzdq	(%r10),%rdx
	movzdq	(%r10),%rbx
	movzdq	(%r10),%rsp
	movzdq	(%r10),%rbp
	movzdq	(%r10),%rsi
	movzdq	(%r10),%rdi
	movzdq	(%r10),%r8
	movzdq	(%r10),%r9
	movzdq	(%r10),%r10
	movzdq	(%r10),%r11
	movzdq	(%r10),%r12
	movzdq	(%r10),%r13
	movzdq	(%r10),%r14
	movzdq	(%r10),%r15
	nop
	movzdq	(%r11),%rax
	movzdq	(%r11),%rcx
	movzdq	(%r11),%rdx
	movzdq	(%r11),%rbx
	movzdq	(%r11),%rsp
	movzdq	(%r11),%rbp
	movzdq	(%r11),%rsi
	movzdq	(%r11),%rdi
	movzdq	(%r11),%r8
	movzdq	(%r11),%r9
	movzdq	(%r11),%r10
	movzdq	(%r11),%r11
	movzdq	(%r11),%r12
	movzdq	(%r11),%r13
	movzdq	(%r11),%r14
	movzdq	(%r11),%r15
	nop
	movzdq	(%r12),%rax
	movzdq	(%r12),%rcx
	movzdq	(%r12),%rdx
	movzdq	(%r12),%rbx
	movzdq	(%r12),%rsp
	movzdq	(%r12),%rbp
	movzdq	(%r12),%rsi
	movzdq	(%r12),%rdi
	movzdq	(%r12),%r8
	movzdq	(%r12),%r9
	movzdq	(%r12),%r10
	movzdq	(%r12),%r11
	movzdq	(%r12),%r12
	movzdq	(%r12),%r13
	movzdq	(%r12),%r14
	movzdq	(%r12),%r15
	nop
	movzdq	(%r13),%rax
	movzdq	(%r13),%rcx
	movzdq	(%r13),%rdx
	movzdq	(%r13),%rbx
	movzdq	(%r13),%rsp
	movzdq	(%r13),%rbp
	movzdq	(%r13),%rsi
	movzdq	(%r13),%rdi
	movzdq	(%r13),%r8
	movzdq	(%r13),%r9
	movzdq	(%r13),%r10
	movzdq	(%r13),%r11
	movzdq	(%r13),%r12
	movzdq	(%r13),%r13
	movzdq	(%r13),%r14
	movzdq	(%r13),%r15
	nop
	movzdq	(%r14),%rax
	movzdq	(%r14),%rcx
	movzdq	(%r14),%rdx
	movzdq	(%r14),%rbx
	movzdq	(%r14),%rsp
	movzdq	(%r14),%rbp
	movzdq	(%r14),%rsi
	movzdq	(%r14),%rdi
	movzdq	(%r14),%r8
	movzdq	(%r14),%r9
	movzdq	(%r14),%r10
	movzdq	(%r14),%r11
	movzdq	(%r14),%r12
	movzdq	(%r14),%r13
	movzdq	(%r14),%r14
	movzdq	(%r14),%r15
	nop
	movzdq	(%r15),%rax
	movzdq	(%r15),%rcx
	movzdq	(%r15),%rdx
	movzdq	(%r15),%rbx
	movzdq	(%r15),%rsp
	movzdq	(%r15),%rbp
	movzdq	(%r15),%rsi
	movzdq	(%r15),%rdi
	movzdq	(%r15),%r8
	movzdq	(%r15),%r9
	movzdq	(%r15),%r10
	movzdq	(%r15),%r11
	movzdq	(%r15),%r12
	movzdq	(%r15),%r13
	movzdq	(%r15),%r14
	movzdq	(%r15),%r15
	nop
	nop
	// off8(reg) -> reg
	movzdq	0x7F(%rax),%rax
	movzdq	0x7F(%rax),%rcx
	movzdq	0x7F(%rax),%rdx
	movzdq	0x7F(%rax),%rbx
	movzdq	0x7F(%rax),%rsp
	movzdq	0x7F(%rax),%rbp
	movzdq	0x7F(%rax),%rsi
	movzdq	0x7F(%rax),%rdi
	movzdq	0x7F(%rax),%r8
	movzdq	0x7F(%rax),%r9
	movzdq	0x7F(%rax),%r10
	movzdq	0x7F(%rax),%r11
	movzdq	0x7F(%rax),%r12
	movzdq	0x7F(%rax),%r13
	movzdq	0x7F(%rax),%r14
	movzdq	0x7F(%rax),%r15
	nop
	movzdq	0x7F(%rcx),%rax
	movzdq	0x7F(%rcx),%rcx
	movzdq	0x7F(%rcx),%rdx
	movzdq	0x7F(%rcx),%rbx
	movzdq	0x7F(%rcx),%rsp
	movzdq	0x7F(%rcx),%rbp
	movzdq	0x7F(%rcx),%rsi
	movzdq	0x7F(%rcx),%rdi
	movzdq	0x7F(%rcx),%r8
	movzdq	0x7F(%rcx),%r9
	movzdq	0x7F(%rcx),%r10
	movzdq	0x7F(%rcx),%r11
	movzdq	0x7F(%rcx),%r12
	movzdq	0x7F(%rcx),%r13
	movzdq	0x7F(%rcx),%r14
	movzdq	0x7F(%rcx),%r15
	nop
	movzdq	0x7F(%rdx),%rax
	movzdq	0x7F(%rdx),%rcx
	movzdq	0x7F(%rdx),%rdx
	movzdq	0x7F(%rdx),%rbx
	movzdq	0x7F(%rdx),%rsp
	movzdq	0x7F(%rdx),%rbp
	movzdq	0x7F(%rdx),%rsi
	movzdq	0x7F(%rdx),%rdi
	movzdq	0x7F(%rdx),%r8
	movzdq	0x7F(%rdx),%r9
	movzdq	0x7F(%rdx),%r10
	movzdq	0x7F(%rdx),%r11
	movzdq	0x7F(%rdx),%r12
	movzdq	0x7F(%rdx),%r13
	movzdq	0x7F(%rdx),%r14
	movzdq	0x7F(%rdx),%r15
	nop
	movzdq	0x7F(%rbx),%rax
	movzdq	0x7F(%rbx),%rcx
	movzdq	0x7F(%rbx),%rdx
	movzdq	0x7F(%rbx),%rbx
	movzdq	0x7F(%rbx),%rsp
	movzdq	0x7F(%rbx),%rbp
	movzdq	0x7F(%rbx),%rsi
	movzdq	0x7F(%rbx),%rdi
	movzdq	0x7F(%rbx),%r8
	movzdq	0x7F(%rbx),%r9
	movzdq	0x7F(%rbx),%r10
	movzdq	0x7F(%rbx),%r11
	movzdq	0x7F(%rbx),%r12
	movzdq	0x7F(%rbx),%r13
	movzdq	0x7F(%rbx),%r14
	movzdq	0x7F(%rbx),%r15
	nop
	movzdq	0x7F(%rsp),%rax
	movzdq	0x7F(%rsp),%rcx
	movzdq	0x7F(%rsp),%rdx
	movzdq	0x7F(%rsp),%rbx
	movzdq	0x7F(%rsp),%rsp
	movzdq	0x7F(%rsp),%rbp
	movzdq	0x7F(%rsp),%rsi
	movzdq	0x7F(%rsp),%rdi
	movzdq	0x7F(%rsp),%r8
	movzdq	0x7F(%rsp),%r9
	movzdq	0x7F(%rsp),%r10
	movzdq	0x7F(%rsp),%r11
	movzdq	0x7F(%rsp),%r12
	movzdq	0x7F(%rsp),%r13
	movzdq	0x7F(%rsp),%r14
	movzdq	0x7F(%rsp),%r15
	nop
	movzdq	0x7F(%rbp),%rax
	movzdq	0x7F(%rbp),%rcx
	movzdq	0x7F(%rbp),%rdx
	movzdq	0x7F(%rbp),%rbx
	movzdq	0x7F(%rbp),%rsp
	movzdq	0x7F(%rbp),%rbp
	movzdq	0x7F(%rbp),%rsi
	movzdq	0x7F(%rbp),%rdi
	movzdq	0x7F(%rbp),%r8
	movzdq	0x7F(%rbp),%r9
	movzdq	0x7F(%rbp),%r10
	movzdq	0x7F(%rbp),%r11
	movzdq	0x7F(%rbp),%r12
	movzdq	0x7F(%rbp),%r13
	movzdq	0x7F(%rbp),%r14
	movzdq	0x7F(%rbp),%r15
	nop
	movzdq	0x7F(%rsi),%rax
	movzdq	0x7F(%rsi),%rcx
	movzdq	0x7F(%rsi),%rdx
	movzdq	0x7F(%rsi),%rbx
	movzdq	0x7F(%rsi),%rsp
	movzdq	0x7F(%rsi),%rbp
	movzdq	0x7F(%rsi),%rsi
	movzdq	0x7F(%rsi),%rdi
	movzdq	0x7F(%rsi),%r8
	movzdq	0x7F(%rsi),%r9
	movzdq	0x7F(%rsi),%r10
	movzdq	0x7F(%rsi),%r11
	movzdq	0x7F(%rsi),%r12
	movzdq	0x7F(%rsi),%r13
	movzdq	0x7F(%rsi),%r14
	movzdq	0x7F(%rsi),%r15
	nop
	movzdq	0x7F(%rdi),%rax
	movzdq	0x7F(%rdi),%rcx
	movzdq	0x7F(%rdi),%rdx
	movzdq	0x7F(%rdi),%rbx
	movzdq	0x7F(%rdi),%rsp
	movzdq	0x7F(%rdi),%rbp
	movzdq	0x7F(%rdi),%rsi
	movzdq	0x7F(%rdi),%rdi
	movzdq	0x7F(%rdi),%r8
	movzdq	0x7F(%rdi),%r9
	movzdq	0x7F(%rdi),%r10
	movzdq	0x7F(%rdi),%r11
	movzdq	0x7F(%rdi),%r12
	movzdq	0x7F(%rdi),%r13
	movzdq	0x7F(%rdi),%r14
	movzdq	0x7F(%rdi),%r15
	nop
	movzdq	0x7F(%r8), %rax
	movzdq	0x7F(%r8), %rcx
	movzdq	0x7F(%r8), %rdx
	movzdq	0x7F(%r8), %rbx
	movzdq	0x7F(%r8), %rsp
	movzdq	0x7F(%r8), %rbp
	movzdq	0x7F(%r8), %rsi
	movzdq	0x7F(%r8), %rdi
	movzdq	0x7F(%r8), %r8
	movzdq	0x7F(%r8), %r9
	movzdq	0x7F(%r8), %r10
	movzdq	0x7F(%r8), %r11
	movzdq	0x7F(%r8), %r12
	movzdq	0x7F(%r8), %r13
	movzdq	0x7F(%r8), %r14
	movzdq	0x7F(%r8), %r15
	nop
	movzdq	0x7F(%r9), %rax
	movzdq	0x7F(%r9), %rcx
	movzdq	0x7F(%r9), %rdx
	movzdq	0x7F(%r9), %rbx
	movzdq	0x7F(%r9), %rsp
	movzdq	0x7F(%r9), %rbp
	movzdq	0x7F(%r9), %rsi
	movzdq	0x7F(%r9), %rdi
	movzdq	0x7F(%r9), %r8
	movzdq	0x7F(%r9), %r9
	movzdq	0x7F(%r9), %r10
	movzdq	0x7F(%r9), %r11
	movzdq	0x7F(%r9), %r12
	movzdq	0x7F(%r9), %r13
	movzdq	0x7F(%r9), %r14
	movzdq	0x7F(%r9), %r15
	nop
	movzdq	0x7F(%r10),%rax
	movzdq	0x7F(%r10),%rcx
	movzdq	0x7F(%r10),%rdx
	movzdq	0x7F(%r10),%rbx
	movzdq	0x7F(%r10),%rsp
	movzdq	0x7F(%r10),%rbp
	movzdq	0x7F(%r10),%rsi
	movzdq	0x7F(%r10),%rdi
	movzdq	0x7F(%r10),%r8
	movzdq	0x7F(%r10),%r9
	movzdq	0x7F(%r10),%r10
	movzdq	0x7F(%r10),%r11
	movzdq	0x7F(%r10),%r12
	movzdq	0x7F(%r10),%r13
	movzdq	0x7F(%r10),%r14
	movzdq	0x7F(%r10),%r15
	nop
	movzdq	0x7F(%r11),%rax
	movzdq	0x7F(%r11),%rcx
	movzdq	0x7F(%r11),%rdx
	movzdq	0x7F(%r11),%rbx
	movzdq	0x7F(%r11),%rsp
	movzdq	0x7F(%r11),%rbp
	movzdq	0x7F(%r11),%rsi
	movzdq	0x7F(%r11),%rdi
	movzdq	0x7F(%r11),%r8
	movzdq	0x7F(%r11),%r9
	movzdq	0x7F(%r11),%r10
	movzdq	0x7F(%r11),%r11
	movzdq	0x7F(%r11),%r12
	movzdq	0x7F(%r11),%r13
	movzdq	0x7F(%r11),%r14
	movzdq	0x7F(%r11),%r15
	nop
	movzdq	0x7F(%r12),%rax
	movzdq	0x7F(%r12),%rcx
	movzdq	0x7F(%r12),%rdx
	movzdq	0x7F(%r12),%rbx
	movzdq	0x7F(%r12),%rsp
	movzdq	0x7F(%r12),%rbp
	movzdq	0x7F(%r12),%rsi
	movzdq	0x7F(%r12),%rdi
	movzdq	0x7F(%r12),%r8
	movzdq	0x7F(%r12),%r9
	movzdq	0x7F(%r12),%r10
	movzdq	0x7F(%r12),%r11
	movzdq	0x7F(%r12),%r12
	movzdq	0x7F(%r12),%r13
	movzdq	0x7F(%r12),%r14
	movzdq	0x7F(%r12),%r15
	nop
	movzdq	0x7F(%r13),%rax
	movzdq	0x7F(%r13),%rcx
	movzdq	0x7F(%r13),%rdx
	movzdq	0x7F(%r13),%rbx
	movzdq	0x7F(%r13),%rsp
	movzdq	0x7F(%r13),%rbp
	movzdq	0x7F(%r13),%rsi
	movzdq	0x7F(%r13),%rdi
	movzdq	0x7F(%r13),%r8
	movzdq	0x7F(%r13),%r9
	movzdq	0x7F(%r13),%r10
	movzdq	0x7F(%r13),%r11
	movzdq	0x7F(%r13),%r12
	movzdq	0x7F(%r13),%r13
	movzdq	0x7F(%r13),%r14
	movzdq	0x7F(%r13),%r15
	nop
	movzdq	0x7F(%r14),%rax
	movzdq	0x7F(%r14),%rcx
	movzdq	0x7F(%r14),%rdx
	movzdq	0x7F(%r14),%rbx
	movzdq	0x7F(%r14),%rsp
	movzdq	0x7F(%r14),%rbp
	movzdq	0x7F(%r14),%rsi
	movzdq	0x7F(%r14),%rdi
	movzdq	0x7F(%r14),%r8
	movzdq	0x7F(%r14),%r9
	movzdq	0x7F(%r14),%r10
	movzdq	0x7F(%r14),%r11
	movzdq	0x7F(%r14),%r12
	movzdq	0x7F(%r14),%r13
	movzdq	0x7F(%r14),%r14
	movzdq	0x7F(%r14),%r15
	nop
	movzdq	0x7F(%r15),%rax
	movzdq	0x7F(%r15),%rcx
	movzdq	0x7F(%r15),%rdx
	movzdq	0x7F(%r15),%rbx
	movzdq	0x7F(%r15),%rsp
	movzdq	0x7F(%r15),%rbp
	movzdq	0x7F(%r15),%rsi
	movzdq	0x7F(%r15),%rdi
	movzdq	0x7F(%r15),%r8
	movzdq	0x7F(%r15),%r9
	movzdq	0x7F(%r15),%r10
	movzdq	0x7F(%r15),%r11
	movzdq	0x7F(%r15),%r12
	movzdq	0x7F(%r15),%r13
	movzdq	0x7F(%r15),%r14
	movzdq	0x7F(%r15),%r15
	nop
	nop
	// off32(reg) -> reg
	movzdq	0x12345678(%rax),%rax
	movzdq	0x12345678(%rax),%rcx
	movzdq	0x12345678(%rax),%rdx
	movzdq	0x12345678(%rax),%rbx
	movzdq	0x12345678(%rax),%rsp
	movzdq	0x12345678(%rax),%rbp
	movzdq	0x12345678(%rax),%rsi
	movzdq	0x12345678(%rax),%rdi
	movzdq	0x12345678(%rax),%r8
	movzdq	0x12345678(%rax),%r9
	movzdq	0x12345678(%rax),%r10
	movzdq	0x12345678(%rax),%r11
	movzdq	0x12345678(%rax),%r12
	movzdq	0x12345678(%rax),%r13
	movzdq	0x12345678(%rax),%r14
	movzdq	0x12345678(%rax),%r15
	nop
	movzdq	0x12345678(%rcx),%rax
	movzdq	0x12345678(%rcx),%rcx
	movzdq	0x12345678(%rcx),%rdx
	movzdq	0x12345678(%rcx),%rbx
	movzdq	0x12345678(%rcx),%rsp
	movzdq	0x12345678(%rcx),%rbp
	movzdq	0x12345678(%rcx),%rsi
	movzdq	0x12345678(%rcx),%rdi
	movzdq	0x12345678(%rcx),%r8
	movzdq	0x12345678(%rcx),%r9
	movzdq	0x12345678(%rcx),%r10
	movzdq	0x12345678(%rcx),%r11
	movzdq	0x12345678(%rcx),%r12
	movzdq	0x12345678(%rcx),%r13
	movzdq	0x12345678(%rcx),%r14
	movzdq	0x12345678(%rcx),%r15
	nop
	movzdq	0x12345678(%rdx),%rax
	movzdq	0x12345678(%rdx),%rcx
	movzdq	0x12345678(%rdx),%rdx
	movzdq	0x12345678(%rdx),%rbx
	movzdq	0x12345678(%rdx),%rsp
	movzdq	0x12345678(%rdx),%rbp
	movzdq	0x12345678(%rdx),%rsi
	movzdq	0x12345678(%rdx),%rdi
	movzdq	0x12345678(%rdx),%r8
	movzdq	0x12345678(%rdx),%r9
	movzdq	0x12345678(%rdx),%r10
	movzdq	0x12345678(%rdx),%r11
	movzdq	0x12345678(%rdx),%r12
	movzdq	0x12345678(%rdx),%r13
	movzdq	0x12345678(%rdx),%r14
	movzdq	0x12345678(%rdx),%r15
	nop
	movzdq	0x12345678(%rbx),%rax
	movzdq	0x12345678(%rbx),%rcx
	movzdq	0x12345678(%rbx),%rdx
	movzdq	0x12345678(%rbx),%rbx
	movzdq	0x12345678(%rbx),%rsp
	movzdq	0x12345678(%rbx),%rbp
	movzdq	0x12345678(%rbx),%rsi
	movzdq	0x12345678(%rbx),%rdi
	movzdq	0x12345678(%rbx),%r8
	movzdq	0x12345678(%rbx),%r9
	movzdq	0x12345678(%rbx),%r10
	movzdq	0x12345678(%rbx),%r11
	movzdq	0x12345678(%rbx),%r12
	movzdq	0x12345678(%rbx),%r13
	movzdq	0x12345678(%rbx),%r14
	movzdq	0x12345678(%rbx),%r15
	nop
	movzdq	0x12345678(%rsp),%rax
	movzdq	0x12345678(%rsp),%rcx
	movzdq	0x12345678(%rsp),%rdx
	movzdq	0x12345678(%rsp),%rbx
	movzdq	0x12345678(%rsp),%rsp
	movzdq	0x12345678(%rsp),%rbp
	movzdq	0x12345678(%rsp),%rsi
	movzdq	0x12345678(%rsp),%rdi
	movzdq	0x12345678(%rsp),%r8
	movzdq	0x12345678(%rsp),%r9
	movzdq	0x12345678(%rsp),%r10
	movzdq	0x12345678(%rsp),%r11
	movzdq	0x12345678(%rsp),%r12
	movzdq	0x12345678(%rsp),%r13
	movzdq	0x12345678(%rsp),%r14
	movzdq	0x12345678(%rsp),%r15
	nop
	movzdq	0x12345678(%rbp),%rax
	movzdq	0x12345678(%rbp),%rcx
	movzdq	0x12345678(%rbp),%rdx
	movzdq	0x12345678(%rbp),%rbx
	movzdq	0x12345678(%rbp),%rsp
	movzdq	0x12345678(%rbp),%rbp
	movzdq	0x12345678(%rbp),%rsi
	movzdq	0x12345678(%rbp),%rdi
	movzdq	0x12345678(%rbp),%r8
	movzdq	0x12345678(%rbp),%r9
	movzdq	0x12345678(%rbp),%r10
	movzdq	0x12345678(%rbp),%r11
	movzdq	0x12345678(%rbp),%r12
	movzdq	0x12345678(%rbp),%r13
	movzdq	0x12345678(%rbp),%r14
	movzdq	0x12345678(%rbp),%r15
	nop
	movzdq	0x12345678(%rsi),%rax
	movzdq	0x12345678(%rsi),%rcx
	movzdq	0x12345678(%rsi),%rdx
	movzdq	0x12345678(%rsi),%rbx
	movzdq	0x12345678(%rsi),%rsp
	movzdq	0x12345678(%rsi),%rbp
	movzdq	0x12345678(%rsi),%rsi
	movzdq	0x12345678(%rsi),%rdi
	movzdq	0x12345678(%rsi),%r8
	movzdq	0x12345678(%rsi),%r9
	movzdq	0x12345678(%rsi),%r10
	movzdq	0x12345678(%rsi),%r11
	movzdq	0x12345678(%rsi),%r12
	movzdq	0x12345678(%rsi),%r13
	movzdq	0x12345678(%rsi),%r14
	movzdq	0x12345678(%rsi),%r15
	nop
	movzdq	0x12345678(%rdi),%rax
	movzdq	0x12345678(%rdi),%rcx
	movzdq	0x12345678(%rdi),%rdx
	movzdq	0x12345678(%rdi),%rbx
	movzdq	0x12345678(%rdi),%rsp
	movzdq	0x12345678(%rdi),%rbp
	movzdq	0x12345678(%rdi),%rsi
	movzdq	0x12345678(%rdi),%rdi
	movzdq	0x12345678(%rdi),%r8
	movzdq	0x12345678(%rdi),%r9
	movzdq	0x12345678(%rdi),%r10
	movzdq	0x12345678(%rdi),%r11
	movzdq	0x12345678(%rdi),%r12
	movzdq	0x12345678(%rdi),%r13
	movzdq	0x12345678(%rdi),%r14
	movzdq	0x12345678(%rdi),%r15
	nop
	movzdq	0x12345678(%r8), %rax
	movzdq	0x12345678(%r8), %rcx
	movzdq	0x12345678(%r8), %rdx
	movzdq	0x12345678(%r8), %rbx
	movzdq	0x12345678(%r8), %rsp
	movzdq	0x12345678(%r8), %rbp
	movzdq	0x12345678(%r8), %rsi
	movzdq	0x12345678(%r8), %rdi
	movzdq	0x12345678(%r8), %r8
	movzdq	0x12345678(%r8), %r9
	movzdq	0x12345678(%r8), %r10
	movzdq	0x12345678(%r8), %r11
	movzdq	0x12345678(%r8), %r12
	movzdq	0x12345678(%r8), %r13
	movzdq	0x12345678(%r8), %r14
	movzdq	0x12345678(%r8), %r15
	nop
	movzdq	0x12345678(%r9), %rax
	movzdq	0x12345678(%r9), %rcx
	movzdq	0x12345678(%r9), %rdx
	movzdq	0x12345678(%r9), %rbx
	movzdq	0x12345678(%r9), %rsp
	movzdq	0x12345678(%r9), %rbp
	movzdq	0x12345678(%r9), %rsi
	movzdq	0x12345678(%r9), %rdi
	movzdq	0x12345678(%r9), %r8
	movzdq	0x12345678(%r9), %r9
	movzdq	0x12345678(%r9), %r10
	movzdq	0x12345678(%r9), %r11
	movzdq	0x12345678(%r9), %r12
	movzdq	0x12345678(%r9), %r13
	movzdq	0x12345678(%r9), %r14
	movzdq	0x12345678(%r9), %r15
	nop
	movzdq	0x12345678(%r10),%rax
	movzdq	0x12345678(%r10),%rcx
	movzdq	0x12345678(%r10),%rdx
	movzdq	0x12345678(%r10),%rbx
	movzdq	0x12345678(%r10),%rsp
	movzdq	0x12345678(%r10),%rbp
	movzdq	0x12345678(%r10),%rsi
	movzdq	0x12345678(%r10),%rdi
	movzdq	0x12345678(%r10),%r8
	movzdq	0x12345678(%r10),%r9
	movzdq	0x12345678(%r10),%r10
	movzdq	0x12345678(%r10),%r11
	movzdq	0x12345678(%r10),%r12
	movzdq	0x12345678(%r10),%r13
	movzdq	0x12345678(%r10),%r14
	movzdq	0x12345678(%r10),%r15
	nop
	movzdq	0x12345678(%r11),%rax
	movzdq	0x12345678(%r11),%rcx
	movzdq	0x12345678(%r11),%rdx
	movzdq	0x12345678(%r11),%rbx
	movzdq	0x12345678(%r11),%rsp
	movzdq	0x12345678(%r11),%rbp
	movzdq	0x12345678(%r11),%rsi
	movzdq	0x12345678(%r11),%rdi
	movzdq	0x12345678(%r11),%r8
	movzdq	0x12345678(%r11),%r9
	movzdq	0x12345678(%r11),%r10
	movzdq	0x12345678(%r11),%r11
	movzdq	0x12345678(%r11),%r12
	movzdq	0x12345678(%r11),%r13
	movzdq	0x12345678(%r11),%r14
	movzdq	0x12345678(%r11),%r15
	nop
	movzdq	0x12345678(%r12),%rax
	movzdq	0x12345678(%r12),%rcx
	movzdq	0x12345678(%r12),%rdx
	movzdq	0x12345678(%r12),%rbx
	movzdq	0x12345678(%r12),%rsp
	movzdq	0x12345678(%r12),%rbp
	movzdq	0x12345678(%r12),%rsi
	movzdq	0x12345678(%r12),%rdi
	movzdq	0x12345678(%r12),%r8
	movzdq	0x12345678(%r12),%r9
	movzdq	0x12345678(%r12),%r10
	movzdq	0x12345678(%r12),%r11
	movzdq	0x12345678(%r12),%r12
	movzdq	0x12345678(%r12),%r13
	movzdq	0x12345678(%r12),%r14
	movzdq	0x12345678(%r12),%r15
	nop
	movzdq	0x12345678(%r13),%rax
	movzdq	0x12345678(%r13),%rcx
	movzdq	0x12345678(%r13),%rdx
	movzdq	0x12345678(%r13),%rbx
	movzdq	0x12345678(%r13),%rsp
	movzdq	0x12345678(%r13),%rbp
	movzdq	0x12345678(%r13),%rsi
	movzdq	0x12345678(%r13),%rdi
	movzdq	0x12345678(%r13),%r8
	movzdq	0x12345678(%r13),%r9
	movzdq	0x12345678(%r13),%r10
	movzdq	0x12345678(%r13),%r11
	movzdq	0x12345678(%r13),%r12
	movzdq	0x12345678(%r13),%r13
	movzdq	0x12345678(%r13),%r14
	movzdq	0x12345678(%r13),%r15
	nop
	movzdq	0x12345678(%r14),%rax
	movzdq	0x12345678(%r14),%rcx
	movzdq	0x12345678(%r14),%rdx
	movzdq	0x12345678(%r14),%rbx
	movzdq	0x12345678(%r14),%rsp
	movzdq	0x12345678(%r14),%rbp
	movzdq	0x12345678(%r14),%rsi
	movzdq	0x12345678(%r14),%rdi
	movzdq	0x12345678(%r14),%r8
	movzdq	0x12345678(%r14),%r9
	movzdq	0x12345678(%r14),%r10
	movzdq	0x12345678(%r14),%r11
	movzdq	0x12345678(%r14),%r12
	movzdq	0x12345678(%r14),%r13
	movzdq	0x12345678(%r14),%r14
	movzdq	0x12345678(%r14),%r15
	nop
	movzdq	0x12345678(%r15),%rax
	movzdq	0x12345678(%r15),%rcx
	movzdq	0x12345678(%r15),%rdx
	movzdq	0x12345678(%r15),%rbx
	movzdq	0x12345678(%r15),%rsp
	movzdq	0x12345678(%r15),%rbp
	movzdq	0x12345678(%r15),%rsi
	movzdq	0x12345678(%r15),%rdi
	movzdq	0x12345678(%r15),%r8
	movzdq	0x12345678(%r15),%r9
	movzdq	0x12345678(%r15),%r10
	movzdq	0x12345678(%r15),%r11
	movzdq	0x12345678(%r15),%r12
	movzdq	0x12345678(%r15),%r13
	movzdq	0x12345678(%r15),%r14
	movzdq	0x12345678(%r15),%r15
	nop
	nop
	ret
	.cfi_endproc
*/
        
	.p2align 4,,15
	.globl	Movslq
	.type	Movslq, @function
Movslq:
	.cfi_startproc
	// reg -> reg
	movslq	%eax,%rax
	movslq	%eax,%rcx
	movslq	%eax,%rdx
	movslq	%eax,%rbx
	movslq	%eax,%rsp
	movslq	%eax,%rbp
	movslq	%eax,%rsi
	movslq	%eax,%rdi
	movslq	%eax,%r8
	movslq	%eax,%r9
	movslq	%eax,%r10
	movslq	%eax,%r11
	movslq	%eax,%r12
	movslq	%eax,%r13
	movslq	%eax,%r14
	movslq	%eax,%r15
	nop
	movslq	%ecx,%rax
	movslq	%ecx,%rcx
	movslq	%ecx,%rdx
	movslq	%ecx,%rbx
	movslq	%ecx,%rsp
	movslq	%ecx,%rbp
	movslq	%ecx,%rsi
	movslq	%ecx,%rdi
	movslq	%ecx,%r8
	movslq	%ecx,%r9
	movslq	%ecx,%r10
	movslq	%ecx,%r11
	movslq	%ecx,%r12
	movslq	%ecx,%r13
	movslq	%ecx,%r14
	movslq	%ecx,%r15
	nop
	movslq	%edx,%rax
	movslq	%edx,%rcx
	movslq	%edx,%rdx
	movslq	%edx,%rbx
	movslq	%edx,%rsp
	movslq	%edx,%rbp
	movslq	%edx,%rsi
	movslq	%edx,%rdi
	movslq	%edx,%r8
	movslq	%edx,%r9
	movslq	%edx,%r10
	movslq	%edx,%r11
	movslq	%edx,%r12
	movslq	%edx,%r13
	movslq	%edx,%r14
	movslq	%edx,%r15
	nop
	movslq	%ebx,%rax
	movslq	%ebx,%rcx
	movslq	%ebx,%rdx
	movslq	%ebx,%rbx
	movslq	%ebx,%rsp
	movslq	%ebx,%rbp
	movslq	%ebx,%rsi
	movslq	%ebx,%rdi
	movslq	%ebx,%r8
	movslq	%ebx,%r9
	movslq	%ebx,%r10
	movslq	%ebx,%r11
	movslq	%ebx,%r12
	movslq	%ebx,%r13
	movslq	%ebx,%r14
	movslq	%ebx,%r15
	nop
	movslq	%esp,%rax
	movslq	%esp,%rcx
	movslq	%esp,%rdx
	movslq	%esp,%rbx
	movslq	%esp,%rsp
	movslq	%esp,%rbp
	movslq	%esp,%rsi
	movslq	%esp,%rdi
	movslq	%esp,%r8
	movslq	%esp,%r9
	movslq	%esp,%r10
	movslq	%esp,%r11
	movslq	%esp,%r12
	movslq	%esp,%r13
	movslq	%esp,%r14
	movslq	%esp,%r15
	nop
	movslq	%ebp,%rax
	movslq	%ebp,%rcx
	movslq	%ebp,%rdx
	movslq	%ebp,%rbx
	movslq	%ebp,%rsp
	movslq	%ebp,%rbp
	movslq	%ebp,%rsi
	movslq	%ebp,%rdi
	movslq	%ebp,%r8
	movslq	%ebp,%r9
	movslq	%ebp,%r10
	movslq	%ebp,%r11
	movslq	%ebp,%r12
	movslq	%ebp,%r13
	movslq	%ebp,%r14
	movslq	%ebp,%r15
	nop
	movslq	%esi,%rax
	movslq	%esi,%rcx
	movslq	%esi,%rdx
	movslq	%esi,%rbx
	movslq	%esi,%rsp
	movslq	%esi,%rbp
	movslq	%esi,%rsi
	movslq	%esi,%rdi
	movslq	%esi,%r8
	movslq	%esi,%r9
	movslq	%esi,%r10
	movslq	%esi,%r11
	movslq	%esi,%r12
	movslq	%esi,%r13
	movslq	%esi,%r14
	movslq	%esi,%r15
	nop
	movslq	%edi,%rax
	movslq	%edi,%rcx
	movslq	%edi,%rdx
	movslq	%edi,%rbx
	movslq	%edi,%rsp
	movslq	%edi,%rbp
	movslq	%edi,%rsi
	movslq	%edi,%rdi
	movslq	%edi,%r8
	movslq	%edi,%r9
	movslq	%edi,%r10
	movslq	%edi,%r11
	movslq	%edi,%r12
	movslq	%edi,%r13
	movslq	%edi,%r14
	movslq	%edi,%r15
	nop
	movslq	%r8d, %rax
	movslq	%r8d, %rcx
	movslq	%r8d, %rdx
	movslq	%r8d, %rbx
	movslq	%r8d, %rsp
	movslq	%r8d, %rbp
	movslq	%r8d, %rsi
	movslq	%r8d, %rdi
	movslq	%r8d, %r8
	movslq	%r8d, %r9
	movslq	%r8d, %r10
	movslq	%r8d, %r11
	movslq	%r8d, %r12
	movslq	%r8d, %r13
	movslq	%r8d, %r14
	movslq	%r8d, %r15
	nop
	movslq	%r9d, %rax
	movslq	%r9d, %rcx
	movslq	%r9d, %rdx
	movslq	%r9d, %rbx
	movslq	%r9d, %rsp
	movslq	%r9d, %rbp
	movslq	%r9d, %rsi
	movslq	%r9d, %rdi
	movslq	%r9d, %r8
	movslq	%r9d, %r9
	movslq	%r9d, %r10
	movslq	%r9d, %r11
	movslq	%r9d, %r12
	movslq	%r9d, %r13
	movslq	%r9d, %r14
	movslq	%r9d, %r15
	nop
	movslq	%r10d,%rax
	movslq	%r10d,%rcx
	movslq	%r10d,%rdx
	movslq	%r10d,%rbx
	movslq	%r10d,%rsp
	movslq	%r10d,%rbp
	movslq	%r10d,%rsi
	movslq	%r10d,%rdi
	movslq	%r10d,%r8
	movslq	%r10d,%r9
	movslq	%r10d,%r10
	movslq	%r10d,%r11
	movslq	%r10d,%r12
	movslq	%r10d,%r13
	movslq	%r10d,%r14
	movslq	%r10d,%r15
	nop
	movslq	%r11d,%rax
	movslq	%r11d,%rcx
	movslq	%r11d,%rdx
	movslq	%r11d,%rbx
	movslq	%r11d,%rsp
	movslq	%r11d,%rbp
	movslq	%r11d,%rsi
	movslq	%r11d,%rdi
	movslq	%r11d,%r8
	movslq	%r11d,%r9
	movslq	%r11d,%r10
	movslq	%r11d,%r11
	movslq	%r11d,%r12
	movslq	%r11d,%r13
	movslq	%r11d,%r14
	movslq	%r11d,%r15
	nop
	movslq	%r12d,%rax
	movslq	%r12d,%rcx
	movslq	%r12d,%rdx
	movslq	%r12d,%rbx
	movslq	%r12d,%rsp
	movslq	%r12d,%rbp
	movslq	%r12d,%rsi
	movslq	%r12d,%rdi
	movslq	%r12d,%r8
	movslq	%r12d,%r9
	movslq	%r12d,%r10
	movslq	%r12d,%r11
	movslq	%r12d,%r12
	movslq	%r12d,%r13
	movslq	%r12d,%r14
	movslq	%r12d,%r15
	nop
	movslq	%r13d,%rax
	movslq	%r13d,%rcx
	movslq	%r13d,%rdx
	movslq	%r13d,%rbx
	movslq	%r13d,%rsp
	movslq	%r13d,%rbp
	movslq	%r13d,%rsi
	movslq	%r13d,%rdi
	movslq	%r13d,%r8
	movslq	%r13d,%r9
	movslq	%r13d,%r10
	movslq	%r13d,%r11
	movslq	%r13d,%r12
	movslq	%r13d,%r13
	movslq	%r13d,%r14
	movslq	%r13d,%r15
	nop
	movslq	%r14d,%rax
	movslq	%r14d,%rcx
	movslq	%r14d,%rdx
	movslq	%r14d,%rbx
	movslq	%r14d,%rsp
	movslq	%r14d,%rbp
	movslq	%r14d,%rsi
	movslq	%r14d,%rdi
	movslq	%r14d,%r8
	movslq	%r14d,%r9
	movslq	%r14d,%r10
	movslq	%r14d,%r11
	movslq	%r14d,%r12
	movslq	%r14d,%r13
	movslq	%r14d,%r14
	movslq	%r14d,%r15
	nop
	movslq	%r15d,%rax
	movslq	%r15d,%rcx
	movslq	%r15d,%rdx
	movslq	%r15d,%rbx
	movslq	%r15d,%rsp
	movslq	%r15d,%rbp
	movslq	%r15d,%rsi
	movslq	%r15d,%rdi
	movslq	%r15d,%r8
	movslq	%r15d,%r9
	movslq	%r15d,%r10
	movslq	%r15d,%r11
	movslq	%r15d,%r12
	movslq	%r15d,%r13
	movslq	%r15d,%r14
	movslq	%r15d,%r15
	nop
        nop
	// (reg) -> reg
	movslq	(%rax),%rax
	movslq	(%rax),%rcx
	movslq	(%rax),%rdx
	movslq	(%rax),%rbx
	movslq	(%rax),%rsp
	movslq	(%rax),%rbp
	movslq	(%rax),%rsi
	movslq	(%rax),%rdi
	movslq	(%rax),%r8
	movslq	(%rax),%r9
	movslq	(%rax),%r10
	movslq	(%rax),%r11
	movslq	(%rax),%r12
	movslq	(%rax),%r13
	movslq	(%rax),%r14
	movslq	(%rax),%r15
	nop
	movslq	(%rcx),%rax
	movslq	(%rcx),%rcx
	movslq	(%rcx),%rdx
	movslq	(%rcx),%rbx
	movslq	(%rcx),%rsp
	movslq	(%rcx),%rbp
	movslq	(%rcx),%rsi
	movslq	(%rcx),%rdi
	movslq	(%rcx),%r8
	movslq	(%rcx),%r9
	movslq	(%rcx),%r10
	movslq	(%rcx),%r11
	movslq	(%rcx),%r12
	movslq	(%rcx),%r13
	movslq	(%rcx),%r14
	movslq	(%rcx),%r15
	nop
	movslq	(%rdx),%rax
	movslq	(%rdx),%rcx
	movslq	(%rdx),%rdx
	movslq	(%rdx),%rbx
	movslq	(%rdx),%rsp
	movslq	(%rdx),%rbp
	movslq	(%rdx),%rsi
	movslq	(%rdx),%rdi
	movslq	(%rdx),%r8
	movslq	(%rdx),%r9
	movslq	(%rdx),%r10
	movslq	(%rdx),%r11
	movslq	(%rdx),%r12
	movslq	(%rdx),%r13
	movslq	(%rdx),%r14
	movslq	(%rdx),%r15
	nop
	movslq	(%rbx),%rax
	movslq	(%rbx),%rcx
	movslq	(%rbx),%rdx
	movslq	(%rbx),%rbx
	movslq	(%rbx),%rsp
	movslq	(%rbx),%rbp
	movslq	(%rbx),%rsi
	movslq	(%rbx),%rdi
	movslq	(%rbx),%r8
	movslq	(%rbx),%r9
	movslq	(%rbx),%r10
	movslq	(%rbx),%r11
	movslq	(%rbx),%r12
	movslq	(%rbx),%r13
	movslq	(%rbx),%r14
	movslq	(%rbx),%r15
	nop
	movslq	(%rsp),%rax
	movslq	(%rsp),%rcx
	movslq	(%rsp),%rdx
	movslq	(%rsp),%rbx
	movslq	(%rsp),%rsp
	movslq	(%rsp),%rbp
	movslq	(%rsp),%rsi
	movslq	(%rsp),%rdi
	movslq	(%rsp),%r8
	movslq	(%rsp),%r9
	movslq	(%rsp),%r10
	movslq	(%rsp),%r11
	movslq	(%rsp),%r12
	movslq	(%rsp),%r13
	movslq	(%rsp),%r14
	movslq	(%rsp),%r15
	nop
	movslq	(%rbp),%rax
	movslq	(%rbp),%rcx
	movslq	(%rbp),%rdx
	movslq	(%rbp),%rbx
	movslq	(%rbp),%rsp
	movslq	(%rbp),%rbp
	movslq	(%rbp),%rsi
	movslq	(%rbp),%rdi
	movslq	(%rbp),%r8
	movslq	(%rbp),%r9
	movslq	(%rbp),%r10
	movslq	(%rbp),%r11
	movslq	(%rbp),%r12
	movslq	(%rbp),%r13
	movslq	(%rbp),%r14
	movslq	(%rbp),%r15
	nop
	movslq	(%rsi),%rax
	movslq	(%rsi),%rcx
	movslq	(%rsi),%rdx
	movslq	(%rsi),%rbx
	movslq	(%rsi),%rsp
	movslq	(%rsi),%rbp
	movslq	(%rsi),%rsi
	movslq	(%rsi),%rdi
	movslq	(%rsi),%r8
	movslq	(%rsi),%r9
	movslq	(%rsi),%r10
	movslq	(%rsi),%r11
	movslq	(%rsi),%r12
	movslq	(%rsi),%r13
	movslq	(%rsi),%r14
	movslq	(%rsi),%r15
	nop
	movslq	(%rdi),%rax
	movslq	(%rdi),%rcx
	movslq	(%rdi),%rdx
	movslq	(%rdi),%rbx
	movslq	(%rdi),%rsp
	movslq	(%rdi),%rbp
	movslq	(%rdi),%rsi
	movslq	(%rdi),%rdi
	movslq	(%rdi),%r8
	movslq	(%rdi),%r9
	movslq	(%rdi),%r10
	movslq	(%rdi),%r11
	movslq	(%rdi),%r12
	movslq	(%rdi),%r13
	movslq	(%rdi),%r14
	movslq	(%rdi),%r15
	nop
	movslq	(%r8), %rax
	movslq	(%r8), %rcx
	movslq	(%r8), %rdx
	movslq	(%r8), %rbx
	movslq	(%r8), %rsp
	movslq	(%r8), %rbp
	movslq	(%r8), %rsi
	movslq	(%r8), %rdi
	movslq	(%r8), %r8
	movslq	(%r8), %r9
	movslq	(%r8), %r10
	movslq	(%r8), %r11
	movslq	(%r8), %r12
	movslq	(%r8), %r13
	movslq	(%r8), %r14
	movslq	(%r8), %r15
	nop
	movslq	(%r9), %rax
	movslq	(%r9), %rcx
	movslq	(%r9), %rdx
	movslq	(%r9), %rbx
	movslq	(%r9), %rsp
	movslq	(%r9), %rbp
	movslq	(%r9), %rsi
	movslq	(%r9), %rdi
	movslq	(%r9), %r8
	movslq	(%r9), %r9
	movslq	(%r9), %r10
	movslq	(%r9), %r11
	movslq	(%r9), %r12
	movslq	(%r9), %r13
	movslq	(%r9), %r14
	movslq	(%r9), %r15
	nop
	movslq	(%r10),%rax
	movslq	(%r10),%rcx
	movslq	(%r10),%rdx
	movslq	(%r10),%rbx
	movslq	(%r10),%rsp
	movslq	(%r10),%rbp
	movslq	(%r10),%rsi
	movslq	(%r10),%rdi
	movslq	(%r10),%r8
	movslq	(%r10),%r9
	movslq	(%r10),%r10
	movslq	(%r10),%r11
	movslq	(%r10),%r12
	movslq	(%r10),%r13
	movslq	(%r10),%r14
	movslq	(%r10),%r15
	nop
	movslq	(%r11),%rax
	movslq	(%r11),%rcx
	movslq	(%r11),%rdx
	movslq	(%r11),%rbx
	movslq	(%r11),%rsp
	movslq	(%r11),%rbp
	movslq	(%r11),%rsi
	movslq	(%r11),%rdi
	movslq	(%r11),%r8
	movslq	(%r11),%r9
	movslq	(%r11),%r10
	movslq	(%r11),%r11
	movslq	(%r11),%r12
	movslq	(%r11),%r13
	movslq	(%r11),%r14
	movslq	(%r11),%r15
	nop
	movslq	(%r12),%rax
	movslq	(%r12),%rcx
	movslq	(%r12),%rdx
	movslq	(%r12),%rbx
	movslq	(%r12),%rsp
	movslq	(%r12),%rbp
	movslq	(%r12),%rsi
	movslq	(%r12),%rdi
	movslq	(%r12),%r8
	movslq	(%r12),%r9
	movslq	(%r12),%r10
	movslq	(%r12),%r11
	movslq	(%r12),%r12
	movslq	(%r12),%r13
	movslq	(%r12),%r14
	movslq	(%r12),%r15
	nop
	movslq	(%r13),%rax
	movslq	(%r13),%rcx
	movslq	(%r13),%rdx
	movslq	(%r13),%rbx
	movslq	(%r13),%rsp
	movslq	(%r13),%rbp
	movslq	(%r13),%rsi
	movslq	(%r13),%rdi
	movslq	(%r13),%r8
	movslq	(%r13),%r9
	movslq	(%r13),%r10
	movslq	(%r13),%r11
	movslq	(%r13),%r12
	movslq	(%r13),%r13
	movslq	(%r13),%r14
	movslq	(%r13),%r15
	nop
	movslq	(%r14),%rax
	movslq	(%r14),%rcx
	movslq	(%r14),%rdx
	movslq	(%r14),%rbx
	movslq	(%r14),%rsp
	movslq	(%r14),%rbp
	movslq	(%r14),%rsi
	movslq	(%r14),%rdi
	movslq	(%r14),%r8
	movslq	(%r14),%r9
	movslq	(%r14),%r10
	movslq	(%r14),%r11
	movslq	(%r14),%r12
	movslq	(%r14),%r13
	movslq	(%r14),%r14
	movslq	(%r14),%r15
	nop
	movslq	(%r15),%rax
	movslq	(%r15),%rcx
	movslq	(%r15),%rdx
	movslq	(%r15),%rbx
	movslq	(%r15),%rsp
	movslq	(%r15),%rbp
	movslq	(%r15),%rsi
	movslq	(%r15),%rdi
	movslq	(%r15),%r8
	movslq	(%r15),%r9
	movslq	(%r15),%r10
	movslq	(%r15),%r11
	movslq	(%r15),%r12
	movslq	(%r15),%r13
	movslq	(%r15),%r14
	movslq	(%r15),%r15
	nop
	nop
	// off8(reg) -> reg
	movslq	0x7F(%rax),%rax
	movslq	0x7F(%rax),%rcx
	movslq	0x7F(%rax),%rdx
	movslq	0x7F(%rax),%rbx
	movslq	0x7F(%rax),%rsp
	movslq	0x7F(%rax),%rbp
	movslq	0x7F(%rax),%rsi
	movslq	0x7F(%rax),%rdi
	movslq	0x7F(%rax),%r8
	movslq	0x7F(%rax),%r9
	movslq	0x7F(%rax),%r10
	movslq	0x7F(%rax),%r11
	movslq	0x7F(%rax),%r12
	movslq	0x7F(%rax),%r13
	movslq	0x7F(%rax),%r14
	movslq	0x7F(%rax),%r15
	nop
	movslq	0x7F(%rcx),%rax
	movslq	0x7F(%rcx),%rcx
	movslq	0x7F(%rcx),%rdx
	movslq	0x7F(%rcx),%rbx
	movslq	0x7F(%rcx),%rsp
	movslq	0x7F(%rcx),%rbp
	movslq	0x7F(%rcx),%rsi
	movslq	0x7F(%rcx),%rdi
	movslq	0x7F(%rcx),%r8
	movslq	0x7F(%rcx),%r9
	movslq	0x7F(%rcx),%r10
	movslq	0x7F(%rcx),%r11
	movslq	0x7F(%rcx),%r12
	movslq	0x7F(%rcx),%r13
	movslq	0x7F(%rcx),%r14
	movslq	0x7F(%rcx),%r15
	nop
	movslq	0x7F(%rdx),%rax
	movslq	0x7F(%rdx),%rcx
	movslq	0x7F(%rdx),%rdx
	movslq	0x7F(%rdx),%rbx
	movslq	0x7F(%rdx),%rsp
	movslq	0x7F(%rdx),%rbp
	movslq	0x7F(%rdx),%rsi
	movslq	0x7F(%rdx),%rdi
	movslq	0x7F(%rdx),%r8
	movslq	0x7F(%rdx),%r9
	movslq	0x7F(%rdx),%r10
	movslq	0x7F(%rdx),%r11
	movslq	0x7F(%rdx),%r12
	movslq	0x7F(%rdx),%r13
	movslq	0x7F(%rdx),%r14
	movslq	0x7F(%rdx),%r15
	nop
	movslq	0x7F(%rbx),%rax
	movslq	0x7F(%rbx),%rcx
	movslq	0x7F(%rbx),%rdx
	movslq	0x7F(%rbx),%rbx
	movslq	0x7F(%rbx),%rsp
	movslq	0x7F(%rbx),%rbp
	movslq	0x7F(%rbx),%rsi
	movslq	0x7F(%rbx),%rdi
	movslq	0x7F(%rbx),%r8
	movslq	0x7F(%rbx),%r9
	movslq	0x7F(%rbx),%r10
	movslq	0x7F(%rbx),%r11
	movslq	0x7F(%rbx),%r12
	movslq	0x7F(%rbx),%r13
	movslq	0x7F(%rbx),%r14
	movslq	0x7F(%rbx),%r15
	nop
	movslq	0x7F(%rsp),%rax
	movslq	0x7F(%rsp),%rcx
	movslq	0x7F(%rsp),%rdx
	movslq	0x7F(%rsp),%rbx
	movslq	0x7F(%rsp),%rsp
	movslq	0x7F(%rsp),%rbp
	movslq	0x7F(%rsp),%rsi
	movslq	0x7F(%rsp),%rdi
	movslq	0x7F(%rsp),%r8
	movslq	0x7F(%rsp),%r9
	movslq	0x7F(%rsp),%r10
	movslq	0x7F(%rsp),%r11
	movslq	0x7F(%rsp),%r12
	movslq	0x7F(%rsp),%r13
	movslq	0x7F(%rsp),%r14
	movslq	0x7F(%rsp),%r15
	nop
	movslq	0x7F(%rbp),%rax
	movslq	0x7F(%rbp),%rcx
	movslq	0x7F(%rbp),%rdx
	movslq	0x7F(%rbp),%rbx
	movslq	0x7F(%rbp),%rsp
	movslq	0x7F(%rbp),%rbp
	movslq	0x7F(%rbp),%rsi
	movslq	0x7F(%rbp),%rdi
	movslq	0x7F(%rbp),%r8
	movslq	0x7F(%rbp),%r9
	movslq	0x7F(%rbp),%r10
	movslq	0x7F(%rbp),%r11
	movslq	0x7F(%rbp),%r12
	movslq	0x7F(%rbp),%r13
	movslq	0x7F(%rbp),%r14
	movslq	0x7F(%rbp),%r15
	nop
	movslq	0x7F(%rsi),%rax
	movslq	0x7F(%rsi),%rcx
	movslq	0x7F(%rsi),%rdx
	movslq	0x7F(%rsi),%rbx
	movslq	0x7F(%rsi),%rsp
	movslq	0x7F(%rsi),%rbp
	movslq	0x7F(%rsi),%rsi
	movslq	0x7F(%rsi),%rdi
	movslq	0x7F(%rsi),%r8
	movslq	0x7F(%rsi),%r9
	movslq	0x7F(%rsi),%r10
	movslq	0x7F(%rsi),%r11
	movslq	0x7F(%rsi),%r12
	movslq	0x7F(%rsi),%r13
	movslq	0x7F(%rsi),%r14
	movslq	0x7F(%rsi),%r15
	nop
	movslq	0x7F(%rdi),%rax
	movslq	0x7F(%rdi),%rcx
	movslq	0x7F(%rdi),%rdx
	movslq	0x7F(%rdi),%rbx
	movslq	0x7F(%rdi),%rsp
	movslq	0x7F(%rdi),%rbp
	movslq	0x7F(%rdi),%rsi
	movslq	0x7F(%rdi),%rdi
	movslq	0x7F(%rdi),%r8
	movslq	0x7F(%rdi),%r9
	movslq	0x7F(%rdi),%r10
	movslq	0x7F(%rdi),%r11
	movslq	0x7F(%rdi),%r12
	movslq	0x7F(%rdi),%r13
	movslq	0x7F(%rdi),%r14
	movslq	0x7F(%rdi),%r15
	nop
	movslq	0x7F(%r8), %rax
	movslq	0x7F(%r8), %rcx
	movslq	0x7F(%r8), %rdx
	movslq	0x7F(%r8), %rbx
	movslq	0x7F(%r8), %rsp
	movslq	0x7F(%r8), %rbp
	movslq	0x7F(%r8), %rsi
	movslq	0x7F(%r8), %rdi
	movslq	0x7F(%r8), %r8
	movslq	0x7F(%r8), %r9
	movslq	0x7F(%r8), %r10
	movslq	0x7F(%r8), %r11
	movslq	0x7F(%r8), %r12
	movslq	0x7F(%r8), %r13
	movslq	0x7F(%r8), %r14
	movslq	0x7F(%r8), %r15
	nop
	movslq	0x7F(%r9), %rax
	movslq	0x7F(%r9), %rcx
	movslq	0x7F(%r9), %rdx
	movslq	0x7F(%r9), %rbx
	movslq	0x7F(%r9), %rsp
	movslq	0x7F(%r9), %rbp
	movslq	0x7F(%r9), %rsi
	movslq	0x7F(%r9), %rdi
	movslq	0x7F(%r9), %r8
	movslq	0x7F(%r9), %r9
	movslq	0x7F(%r9), %r10
	movslq	0x7F(%r9), %r11
	movslq	0x7F(%r9), %r12
	movslq	0x7F(%r9), %r13
	movslq	0x7F(%r9), %r14
	movslq	0x7F(%r9), %r15
	nop
	movslq	0x7F(%r10),%rax
	movslq	0x7F(%r10),%rcx
	movslq	0x7F(%r10),%rdx
	movslq	0x7F(%r10),%rbx
	movslq	0x7F(%r10),%rsp
	movslq	0x7F(%r10),%rbp
	movslq	0x7F(%r10),%rsi
	movslq	0x7F(%r10),%rdi
	movslq	0x7F(%r10),%r8
	movslq	0x7F(%r10),%r9
	movslq	0x7F(%r10),%r10
	movslq	0x7F(%r10),%r11
	movslq	0x7F(%r10),%r12
	movslq	0x7F(%r10),%r13
	movslq	0x7F(%r10),%r14
	movslq	0x7F(%r10),%r15
	nop
	movslq	0x7F(%r11),%rax
	movslq	0x7F(%r11),%rcx
	movslq	0x7F(%r11),%rdx
	movslq	0x7F(%r11),%rbx
	movslq	0x7F(%r11),%rsp
	movslq	0x7F(%r11),%rbp
	movslq	0x7F(%r11),%rsi
	movslq	0x7F(%r11),%rdi
	movslq	0x7F(%r11),%r8
	movslq	0x7F(%r11),%r9
	movslq	0x7F(%r11),%r10
	movslq	0x7F(%r11),%r11
	movslq	0x7F(%r11),%r12
	movslq	0x7F(%r11),%r13
	movslq	0x7F(%r11),%r14
	movslq	0x7F(%r11),%r15
	nop
	movslq	0x7F(%r12),%rax
	movslq	0x7F(%r12),%rcx
	movslq	0x7F(%r12),%rdx
	movslq	0x7F(%r12),%rbx
	movslq	0x7F(%r12),%rsp
	movslq	0x7F(%r12),%rbp
	movslq	0x7F(%r12),%rsi
	movslq	0x7F(%r12),%rdi
	movslq	0x7F(%r12),%r8
	movslq	0x7F(%r12),%r9
	movslq	0x7F(%r12),%r10
	movslq	0x7F(%r12),%r11
	movslq	0x7F(%r12),%r12
	movslq	0x7F(%r12),%r13
	movslq	0x7F(%r12),%r14
	movslq	0x7F(%r12),%r15
	nop
	movslq	0x7F(%r13),%rax
	movslq	0x7F(%r13),%rcx
	movslq	0x7F(%r13),%rdx
	movslq	0x7F(%r13),%rbx
	movslq	0x7F(%r13),%rsp
	movslq	0x7F(%r13),%rbp
	movslq	0x7F(%r13),%rsi
	movslq	0x7F(%r13),%rdi
	movslq	0x7F(%r13),%r8
	movslq	0x7F(%r13),%r9
	movslq	0x7F(%r13),%r10
	movslq	0x7F(%r13),%r11
	movslq	0x7F(%r13),%r12
	movslq	0x7F(%r13),%r13
	movslq	0x7F(%r13),%r14
	movslq	0x7F(%r13),%r15
	nop
	movslq	0x7F(%r14),%rax
	movslq	0x7F(%r14),%rcx
	movslq	0x7F(%r14),%rdx
	movslq	0x7F(%r14),%rbx
	movslq	0x7F(%r14),%rsp
	movslq	0x7F(%r14),%rbp
	movslq	0x7F(%r14),%rsi
	movslq	0x7F(%r14),%rdi
	movslq	0x7F(%r14),%r8
	movslq	0x7F(%r14),%r9
	movslq	0x7F(%r14),%r10
	movslq	0x7F(%r14),%r11
	movslq	0x7F(%r14),%r12
	movslq	0x7F(%r14),%r13
	movslq	0x7F(%r14),%r14
	movslq	0x7F(%r14),%r15
	nop
	movslq	0x7F(%r15),%rax
	movslq	0x7F(%r15),%rcx
	movslq	0x7F(%r15),%rdx
	movslq	0x7F(%r15),%rbx
	movslq	0x7F(%r15),%rsp
	movslq	0x7F(%r15),%rbp
	movslq	0x7F(%r15),%rsi
	movslq	0x7F(%r15),%rdi
	movslq	0x7F(%r15),%r8
	movslq	0x7F(%r15),%r9
	movslq	0x7F(%r15),%r10
	movslq	0x7F(%r15),%r11
	movslq	0x7F(%r15),%r12
	movslq	0x7F(%r15),%r13
	movslq	0x7F(%r15),%r14
	movslq	0x7F(%r15),%r15
	nop
	nop
	// off32(reg) -> reg
	movslq	0x12345678(%rax),%rax
	movslq	0x12345678(%rax),%rcx
	movslq	0x12345678(%rax),%rdx
	movslq	0x12345678(%rax),%rbx
	movslq	0x12345678(%rax),%rsp
	movslq	0x12345678(%rax),%rbp
	movslq	0x12345678(%rax),%rsi
	movslq	0x12345678(%rax),%rdi
	movslq	0x12345678(%rax),%r8
	movslq	0x12345678(%rax),%r9
	movslq	0x12345678(%rax),%r10
	movslq	0x12345678(%rax),%r11
	movslq	0x12345678(%rax),%r12
	movslq	0x12345678(%rax),%r13
	movslq	0x12345678(%rax),%r14
	movslq	0x12345678(%rax),%r15
	nop
	movslq	0x12345678(%rcx),%rax
	movslq	0x12345678(%rcx),%rcx
	movslq	0x12345678(%rcx),%rdx
	movslq	0x12345678(%rcx),%rbx
	movslq	0x12345678(%rcx),%rsp
	movslq	0x12345678(%rcx),%rbp
	movslq	0x12345678(%rcx),%rsi
	movslq	0x12345678(%rcx),%rdi
	movslq	0x12345678(%rcx),%r8
	movslq	0x12345678(%rcx),%r9
	movslq	0x12345678(%rcx),%r10
	movslq	0x12345678(%rcx),%r11
	movslq	0x12345678(%rcx),%r12
	movslq	0x12345678(%rcx),%r13
	movslq	0x12345678(%rcx),%r14
	movslq	0x12345678(%rcx),%r15
	nop
	movslq	0x12345678(%rdx),%rax
	movslq	0x12345678(%rdx),%rcx
	movslq	0x12345678(%rdx),%rdx
	movslq	0x12345678(%rdx),%rbx
	movslq	0x12345678(%rdx),%rsp
	movslq	0x12345678(%rdx),%rbp
	movslq	0x12345678(%rdx),%rsi
	movslq	0x12345678(%rdx),%rdi
	movslq	0x12345678(%rdx),%r8
	movslq	0x12345678(%rdx),%r9
	movslq	0x12345678(%rdx),%r10
	movslq	0x12345678(%rdx),%r11
	movslq	0x12345678(%rdx),%r12
	movslq	0x12345678(%rdx),%r13
	movslq	0x12345678(%rdx),%r14
	movslq	0x12345678(%rdx),%r15
	nop
	movslq	0x12345678(%rbx),%rax
	movslq	0x12345678(%rbx),%rcx
	movslq	0x12345678(%rbx),%rdx
	movslq	0x12345678(%rbx),%rbx
	movslq	0x12345678(%rbx),%rsp
	movslq	0x12345678(%rbx),%rbp
	movslq	0x12345678(%rbx),%rsi
	movslq	0x12345678(%rbx),%rdi
	movslq	0x12345678(%rbx),%r8
	movslq	0x12345678(%rbx),%r9
	movslq	0x12345678(%rbx),%r10
	movslq	0x12345678(%rbx),%r11
	movslq	0x12345678(%rbx),%r12
	movslq	0x12345678(%rbx),%r13
	movslq	0x12345678(%rbx),%r14
	movslq	0x12345678(%rbx),%r15
	nop
	movslq	0x12345678(%rsp),%rax
	movslq	0x12345678(%rsp),%rcx
	movslq	0x12345678(%rsp),%rdx
	movslq	0x12345678(%rsp),%rbx
	movslq	0x12345678(%rsp),%rsp
	movslq	0x12345678(%rsp),%rbp
	movslq	0x12345678(%rsp),%rsi
	movslq	0x12345678(%rsp),%rdi
	movslq	0x12345678(%rsp),%r8
	movslq	0x12345678(%rsp),%r9
	movslq	0x12345678(%rsp),%r10
	movslq	0x12345678(%rsp),%r11
	movslq	0x12345678(%rsp),%r12
	movslq	0x12345678(%rsp),%r13
	movslq	0x12345678(%rsp),%r14
	movslq	0x12345678(%rsp),%r15
	nop
	movslq	0x12345678(%rbp),%rax
	movslq	0x12345678(%rbp),%rcx
	movslq	0x12345678(%rbp),%rdx
	movslq	0x12345678(%rbp),%rbx
	movslq	0x12345678(%rbp),%rsp
	movslq	0x12345678(%rbp),%rbp
	movslq	0x12345678(%rbp),%rsi
	movslq	0x12345678(%rbp),%rdi
	movslq	0x12345678(%rbp),%r8
	movslq	0x12345678(%rbp),%r9
	movslq	0x12345678(%rbp),%r10
	movslq	0x12345678(%rbp),%r11
	movslq	0x12345678(%rbp),%r12
	movslq	0x12345678(%rbp),%r13
	movslq	0x12345678(%rbp),%r14
	movslq	0x12345678(%rbp),%r15
	nop
	movslq	0x12345678(%rsi),%rax
	movslq	0x12345678(%rsi),%rcx
	movslq	0x12345678(%rsi),%rdx
	movslq	0x12345678(%rsi),%rbx
	movslq	0x12345678(%rsi),%rsp
	movslq	0x12345678(%rsi),%rbp
	movslq	0x12345678(%rsi),%rsi
	movslq	0x12345678(%rsi),%rdi
	movslq	0x12345678(%rsi),%r8
	movslq	0x12345678(%rsi),%r9
	movslq	0x12345678(%rsi),%r10
	movslq	0x12345678(%rsi),%r11
	movslq	0x12345678(%rsi),%r12
	movslq	0x12345678(%rsi),%r13
	movslq	0x12345678(%rsi),%r14
	movslq	0x12345678(%rsi),%r15
	nop
	movslq	0x12345678(%rdi),%rax
	movslq	0x12345678(%rdi),%rcx
	movslq	0x12345678(%rdi),%rdx
	movslq	0x12345678(%rdi),%rbx
	movslq	0x12345678(%rdi),%rsp
	movslq	0x12345678(%rdi),%rbp
	movslq	0x12345678(%rdi),%rsi
	movslq	0x12345678(%rdi),%rdi
	movslq	0x12345678(%rdi),%r8
	movslq	0x12345678(%rdi),%r9
	movslq	0x12345678(%rdi),%r10
	movslq	0x12345678(%rdi),%r11
	movslq	0x12345678(%rdi),%r12
	movslq	0x12345678(%rdi),%r13
	movslq	0x12345678(%rdi),%r14
	movslq	0x12345678(%rdi),%r15
	nop
	movslq	0x12345678(%r8), %rax
	movslq	0x12345678(%r8), %rcx
	movslq	0x12345678(%r8), %rdx
	movslq	0x12345678(%r8), %rbx
	movslq	0x12345678(%r8), %rsp
	movslq	0x12345678(%r8), %rbp
	movslq	0x12345678(%r8), %rsi
	movslq	0x12345678(%r8), %rdi
	movslq	0x12345678(%r8), %r8
	movslq	0x12345678(%r8), %r9
	movslq	0x12345678(%r8), %r10
	movslq	0x12345678(%r8), %r11
	movslq	0x12345678(%r8), %r12
	movslq	0x12345678(%r8), %r13
	movslq	0x12345678(%r8), %r14
	movslq	0x12345678(%r8), %r15
	nop
	movslq	0x12345678(%r9), %rax
	movslq	0x12345678(%r9), %rcx
	movslq	0x12345678(%r9), %rdx
	movslq	0x12345678(%r9), %rbx
	movslq	0x12345678(%r9), %rsp
	movslq	0x12345678(%r9), %rbp
	movslq	0x12345678(%r9), %rsi
	movslq	0x12345678(%r9), %rdi
	movslq	0x12345678(%r9), %r8
	movslq	0x12345678(%r9), %r9
	movslq	0x12345678(%r9), %r10
	movslq	0x12345678(%r9), %r11
	movslq	0x12345678(%r9), %r12
	movslq	0x12345678(%r9), %r13
	movslq	0x12345678(%r9), %r14
	movslq	0x12345678(%r9), %r15
	nop
	movslq	0x12345678(%r10),%rax
	movslq	0x12345678(%r10),%rcx
	movslq	0x12345678(%r10),%rdx
	movslq	0x12345678(%r10),%rbx
	movslq	0x12345678(%r10),%rsp
	movslq	0x12345678(%r10),%rbp
	movslq	0x12345678(%r10),%rsi
	movslq	0x12345678(%r10),%rdi
	movslq	0x12345678(%r10),%r8
	movslq	0x12345678(%r10),%r9
	movslq	0x12345678(%r10),%r10
	movslq	0x12345678(%r10),%r11
	movslq	0x12345678(%r10),%r12
	movslq	0x12345678(%r10),%r13
	movslq	0x12345678(%r10),%r14
	movslq	0x12345678(%r10),%r15
	nop
	movslq	0x12345678(%r11),%rax
	movslq	0x12345678(%r11),%rcx
	movslq	0x12345678(%r11),%rdx
	movslq	0x12345678(%r11),%rbx
	movslq	0x12345678(%r11),%rsp
	movslq	0x12345678(%r11),%rbp
	movslq	0x12345678(%r11),%rsi
	movslq	0x12345678(%r11),%rdi
	movslq	0x12345678(%r11),%r8
	movslq	0x12345678(%r11),%r9
	movslq	0x12345678(%r11),%r10
	movslq	0x12345678(%r11),%r11
	movslq	0x12345678(%r11),%r12
	movslq	0x12345678(%r11),%r13
	movslq	0x12345678(%r11),%r14
	movslq	0x12345678(%r11),%r15
	nop
	movslq	0x12345678(%r12),%rax
	movslq	0x12345678(%r12),%rcx
	movslq	0x12345678(%r12),%rdx
	movslq	0x12345678(%r12),%rbx
	movslq	0x12345678(%r12),%rsp
	movslq	0x12345678(%r12),%rbp
	movslq	0x12345678(%r12),%rsi
	movslq	0x12345678(%r12),%rdi
	movslq	0x12345678(%r12),%r8
	movslq	0x12345678(%r12),%r9
	movslq	0x12345678(%r12),%r10
	movslq	0x12345678(%r12),%r11
	movslq	0x12345678(%r12),%r12
	movslq	0x12345678(%r12),%r13
	movslq	0x12345678(%r12),%r14
	movslq	0x12345678(%r12),%r15
	nop
	movslq	0x12345678(%r13),%rax
	movslq	0x12345678(%r13),%rcx
	movslq	0x12345678(%r13),%rdx
	movslq	0x12345678(%r13),%rbx
	movslq	0x12345678(%r13),%rsp
	movslq	0x12345678(%r13),%rbp
	movslq	0x12345678(%r13),%rsi
	movslq	0x12345678(%r13),%rdi
	movslq	0x12345678(%r13),%r8
	movslq	0x12345678(%r13),%r9
	movslq	0x12345678(%r13),%r10
	movslq	0x12345678(%r13),%r11
	movslq	0x12345678(%r13),%r12
	movslq	0x12345678(%r13),%r13
	movslq	0x12345678(%r13),%r14
	movslq	0x12345678(%r13),%r15
	nop
	movslq	0x12345678(%r14),%rax
	movslq	0x12345678(%r14),%rcx
	movslq	0x12345678(%r14),%rdx
	movslq	0x12345678(%r14),%rbx
	movslq	0x12345678(%r14),%rsp
	movslq	0x12345678(%r14),%rbp
	movslq	0x12345678(%r14),%rsi
	movslq	0x12345678(%r14),%rdi
	movslq	0x12345678(%r14),%r8
	movslq	0x12345678(%r14),%r9
	movslq	0x12345678(%r14),%r10
	movslq	0x12345678(%r14),%r11
	movslq	0x12345678(%r14),%r12
	movslq	0x12345678(%r14),%r13
	movslq	0x12345678(%r14),%r14
	movslq	0x12345678(%r14),%r15
	nop
	movslq	0x12345678(%r15),%rax
	movslq	0x12345678(%r15),%rcx
	movslq	0x12345678(%r15),%rdx
	movslq	0x12345678(%r15),%rbx
	movslq	0x12345678(%r15),%rsp
	movslq	0x12345678(%r15),%rbp
	movslq	0x12345678(%r15),%rsi
	movslq	0x12345678(%r15),%rdi
	movslq	0x12345678(%r15),%r8
	movslq	0x12345678(%r15),%r9
	movslq	0x12345678(%r15),%r10
	movslq	0x12345678(%r15),%r11
	movslq	0x12345678(%r15),%r12
	movslq	0x12345678(%r15),%r13
	movslq	0x12345678(%r15),%r14
	movslq	0x12345678(%r15),%r15
	nop
	nop
	ret
	.cfi_endproc

