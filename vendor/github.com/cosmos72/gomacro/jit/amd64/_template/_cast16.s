	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Movzwq
	.type	Movzwq, @function
Movzwq:
	.cfi_startproc
	// reg -> reg
	movzwq	%ax,%rax
	movzwq	%ax,%rcx
	movzwq	%ax,%rdx
	movzwq	%ax,%rbx
	movzwq	%ax,%rsp
	movzwq	%ax,%rbp
	movzwq	%ax,%rsi
	movzwq	%ax,%rdi
	movzwq	%ax,%r8
	movzwq	%ax,%r9
	movzwq	%ax,%r10
	movzwq	%ax,%r11
	movzwq	%ax,%r12
	movzwq	%ax,%r13
	movzwq	%ax,%r14
	movzwq	%ax,%r15
	nop
	movzwq	%cx,%rax
	movzwq	%cx,%rcx
	movzwq	%cx,%rdx
	movzwq	%cx,%rbx
	movzwq	%cx,%rsp
	movzwq	%cx,%rbp
	movzwq	%cx,%rsi
	movzwq	%cx,%rdi
	movzwq	%cx,%r8
	movzwq	%cx,%r9
	movzwq	%cx,%r10
	movzwq	%cx,%r11
	movzwq	%cx,%r12
	movzwq	%cx,%r13
	movzwq	%cx,%r14
	movzwq	%cx,%r15
	nop
	movzwq	%dx,%rax
	movzwq	%dx,%rcx
	movzwq	%dx,%rdx
	movzwq	%dx,%rbx
	movzwq	%dx,%rsp
	movzwq	%dx,%rbp
	movzwq	%dx,%rsi
	movzwq	%dx,%rdi
	movzwq	%dx,%r8
	movzwq	%dx,%r9
	movzwq	%dx,%r10
	movzwq	%dx,%r11
	movzwq	%dx,%r12
	movzwq	%dx,%r13
	movzwq	%dx,%r14
	movzwq	%dx,%r15
	nop
	movzwq	%bx,%rax
	movzwq	%bx,%rcx
	movzwq	%bx,%rdx
	movzwq	%bx,%rbx
	movzwq	%bx,%rsp
	movzwq	%bx,%rbp
	movzwq	%bx,%rsi
	movzwq	%bx,%rdi
	movzwq	%bx,%r8
	movzwq	%bx,%r9
	movzwq	%bx,%r10
	movzwq	%bx,%r11
	movzwq	%bx,%r12
	movzwq	%bx,%r13
	movzwq	%bx,%r14
	movzwq	%bx,%r15
	nop
	movzwq	%sp,%rax
	movzwq	%sp,%rcx
	movzwq	%sp,%rdx
	movzwq	%sp,%rbx
	movzwq	%sp,%rsp
	movzwq	%sp,%rbp
	movzwq	%sp,%rsi
	movzwq	%sp,%rdi
	movzwq	%sp,%r8
	movzwq	%sp,%r9
	movzwq	%sp,%r10
	movzwq	%sp,%r11
	movzwq	%sp,%r12
	movzwq	%sp,%r13
	movzwq	%sp,%r14
	movzwq	%sp,%r15
	nop
	movzwq	%bp,%rax
	movzwq	%bp,%rcx
	movzwq	%bp,%rdx
	movzwq	%bp,%rbx
	movzwq	%bp,%rsp
	movzwq	%bp,%rbp
	movzwq	%bp,%rsi
	movzwq	%bp,%rdi
	movzwq	%bp,%r8
	movzwq	%bp,%r9
	movzwq	%bp,%r10
	movzwq	%bp,%r11
	movzwq	%bp,%r12
	movzwq	%bp,%r13
	movzwq	%bp,%r14
	movzwq	%bp,%r15
	nop
	movzwq	%si,%rax
	movzwq	%si,%rcx
	movzwq	%si,%rdx
	movzwq	%si,%rbx
	movzwq	%si,%rsp
	movzwq	%si,%rbp
	movzwq	%si,%rsi
	movzwq	%si,%rdi
	movzwq	%si,%r8
	movzwq	%si,%r9
	movzwq	%si,%r10
	movzwq	%si,%r11
	movzwq	%si,%r12
	movzwq	%si,%r13
	movzwq	%si,%r14
	movzwq	%si,%r15
	nop
	movzwq	%di,%rax
	movzwq	%di,%rcx
	movzwq	%di,%rdx
	movzwq	%di,%rbx
	movzwq	%di,%rsp
	movzwq	%di,%rbp
	movzwq	%di,%rsi
	movzwq	%di,%rdi
	movzwq	%di,%r8
	movzwq	%di,%r9
	movzwq	%di,%r10
	movzwq	%di,%r11
	movzwq	%di,%r12
	movzwq	%di,%r13
	movzwq	%di,%r14
	movzwq	%di,%r15
	nop
	movzwq	%r8w, %rax
	movzwq	%r8w, %rcx
	movzwq	%r8w, %rdx
	movzwq	%r8w, %rbx
	movzwq	%r8w, %rsp
	movzwq	%r8w, %rbp
	movzwq	%r8w, %rsi
	movzwq	%r8w, %rdi
	movzwq	%r8w, %r8
	movzwq	%r8w, %r9
	movzwq	%r8w, %r10
	movzwq	%r8w, %r11
	movzwq	%r8w, %r12
	movzwq	%r8w, %r13
	movzwq	%r8w, %r14
	movzwq	%r8w, %r15
	nop
	movzwq	%r9w, %rax
	movzwq	%r9w, %rcx
	movzwq	%r9w, %rdx
	movzwq	%r9w, %rbx
	movzwq	%r9w, %rsp
	movzwq	%r9w, %rbp
	movzwq	%r9w, %rsi
	movzwq	%r9w, %rdi
	movzwq	%r9w, %r8
	movzwq	%r9w, %r9
	movzwq	%r9w, %r10
	movzwq	%r9w, %r11
	movzwq	%r9w, %r12
	movzwq	%r9w, %r13
	movzwq	%r9w, %r14
	movzwq	%r9w, %r15
	nop
	movzwq	%r10w,%rax
	movzwq	%r10w,%rcx
	movzwq	%r10w,%rdx
	movzwq	%r10w,%rbx
	movzwq	%r10w,%rsp
	movzwq	%r10w,%rbp
	movzwq	%r10w,%rsi
	movzwq	%r10w,%rdi
	movzwq	%r10w,%r8
	movzwq	%r10w,%r9
	movzwq	%r10w,%r10
	movzwq	%r10w,%r11
	movzwq	%r10w,%r12
	movzwq	%r10w,%r13
	movzwq	%r10w,%r14
	movzwq	%r10w,%r15
	nop
	movzwq	%r11w,%rax
	movzwq	%r11w,%rcx
	movzwq	%r11w,%rdx
	movzwq	%r11w,%rbx
	movzwq	%r11w,%rsp
	movzwq	%r11w,%rbp
	movzwq	%r11w,%rsi
	movzwq	%r11w,%rdi
	movzwq	%r11w,%r8
	movzwq	%r11w,%r9
	movzwq	%r11w,%r10
	movzwq	%r11w,%r11
	movzwq	%r11w,%r12
	movzwq	%r11w,%r13
	movzwq	%r11w,%r14
	movzwq	%r11w,%r15
	nop
	movzwq	%r12w,%rax
	movzwq	%r12w,%rcx
	movzwq	%r12w,%rdx
	movzwq	%r12w,%rbx
	movzwq	%r12w,%rsp
	movzwq	%r12w,%rbp
	movzwq	%r12w,%rsi
	movzwq	%r12w,%rdi
	movzwq	%r12w,%r8
	movzwq	%r12w,%r9
	movzwq	%r12w,%r10
	movzwq	%r12w,%r11
	movzwq	%r12w,%r12
	movzwq	%r12w,%r13
	movzwq	%r12w,%r14
	movzwq	%r12w,%r15
	nop
	movzwq	%r13w,%rax
	movzwq	%r13w,%rcx
	movzwq	%r13w,%rdx
	movzwq	%r13w,%rbx
	movzwq	%r13w,%rsp
	movzwq	%r13w,%rbp
	movzwq	%r13w,%rsi
	movzwq	%r13w,%rdi
	movzwq	%r13w,%r8
	movzwq	%r13w,%r9
	movzwq	%r13w,%r10
	movzwq	%r13w,%r11
	movzwq	%r13w,%r12
	movzwq	%r13w,%r13
	movzwq	%r13w,%r14
	movzwq	%r13w,%r15
	nop
	movzwq	%r14w,%rax
	movzwq	%r14w,%rcx
	movzwq	%r14w,%rdx
	movzwq	%r14w,%rbx
	movzwq	%r14w,%rsp
	movzwq	%r14w,%rbp
	movzwq	%r14w,%rsi
	movzwq	%r14w,%rdi
	movzwq	%r14w,%r8
	movzwq	%r14w,%r9
	movzwq	%r14w,%r10
	movzwq	%r14w,%r11
	movzwq	%r14w,%r12
	movzwq	%r14w,%r13
	movzwq	%r14w,%r14
	movzwq	%r14w,%r15
	nop
	movzwq	%r15w,%rax
	movzwq	%r15w,%rcx
	movzwq	%r15w,%rdx
	movzwq	%r15w,%rbx
	movzwq	%r15w,%rsp
	movzwq	%r15w,%rbp
	movzwq	%r15w,%rsi
	movzwq	%r15w,%rdi
	movzwq	%r15w,%r8
	movzwq	%r15w,%r9
	movzwq	%r15w,%r10
	movzwq	%r15w,%r11
	movzwq	%r15w,%r12
	movzwq	%r15w,%r13
	movzwq	%r15w,%r14
	movzwq	%r15w,%r15
	nop
        nop
	// (reg) -> reg
	movzwq	(%rax),%rax
	movzwq	(%rax),%rcx
	movzwq	(%rax),%rdx
	movzwq	(%rax),%rbx
	movzwq	(%rax),%rsp
	movzwq	(%rax),%rbp
	movzwq	(%rax),%rsi
	movzwq	(%rax),%rdi
	movzwq	(%rax),%r8
	movzwq	(%rax),%r9
	movzwq	(%rax),%r10
	movzwq	(%rax),%r11
	movzwq	(%rax),%r12
	movzwq	(%rax),%r13
	movzwq	(%rax),%r14
	movzwq	(%rax),%r15
	nop
	movzwq	(%rcx),%rax
	movzwq	(%rcx),%rcx
	movzwq	(%rcx),%rdx
	movzwq	(%rcx),%rbx
	movzwq	(%rcx),%rsp
	movzwq	(%rcx),%rbp
	movzwq	(%rcx),%rsi
	movzwq	(%rcx),%rdi
	movzwq	(%rcx),%r8
	movzwq	(%rcx),%r9
	movzwq	(%rcx),%r10
	movzwq	(%rcx),%r11
	movzwq	(%rcx),%r12
	movzwq	(%rcx),%r13
	movzwq	(%rcx),%r14
	movzwq	(%rcx),%r15
	nop
	movzwq	(%rdx),%rax
	movzwq	(%rdx),%rcx
	movzwq	(%rdx),%rdx
	movzwq	(%rdx),%rbx
	movzwq	(%rdx),%rsp
	movzwq	(%rdx),%rbp
	movzwq	(%rdx),%rsi
	movzwq	(%rdx),%rdi
	movzwq	(%rdx),%r8
	movzwq	(%rdx),%r9
	movzwq	(%rdx),%r10
	movzwq	(%rdx),%r11
	movzwq	(%rdx),%r12
	movzwq	(%rdx),%r13
	movzwq	(%rdx),%r14
	movzwq	(%rdx),%r15
	nop
	movzwq	(%rbx),%rax
	movzwq	(%rbx),%rcx
	movzwq	(%rbx),%rdx
	movzwq	(%rbx),%rbx
	movzwq	(%rbx),%rsp
	movzwq	(%rbx),%rbp
	movzwq	(%rbx),%rsi
	movzwq	(%rbx),%rdi
	movzwq	(%rbx),%r8
	movzwq	(%rbx),%r9
	movzwq	(%rbx),%r10
	movzwq	(%rbx),%r11
	movzwq	(%rbx),%r12
	movzwq	(%rbx),%r13
	movzwq	(%rbx),%r14
	movzwq	(%rbx),%r15
	nop
	movzwq	(%rsp),%rax
	movzwq	(%rsp),%rcx
	movzwq	(%rsp),%rdx
	movzwq	(%rsp),%rbx
	movzwq	(%rsp),%rsp
	movzwq	(%rsp),%rbp
	movzwq	(%rsp),%rsi
	movzwq	(%rsp),%rdi
	movzwq	(%rsp),%r8
	movzwq	(%rsp),%r9
	movzwq	(%rsp),%r10
	movzwq	(%rsp),%r11
	movzwq	(%rsp),%r12
	movzwq	(%rsp),%r13
	movzwq	(%rsp),%r14
	movzwq	(%rsp),%r15
	nop
	movzwq	(%rbp),%rax
	movzwq	(%rbp),%rcx
	movzwq	(%rbp),%rdx
	movzwq	(%rbp),%rbx
	movzwq	(%rbp),%rsp
	movzwq	(%rbp),%rbp
	movzwq	(%rbp),%rsi
	movzwq	(%rbp),%rdi
	movzwq	(%rbp),%r8
	movzwq	(%rbp),%r9
	movzwq	(%rbp),%r10
	movzwq	(%rbp),%r11
	movzwq	(%rbp),%r12
	movzwq	(%rbp),%r13
	movzwq	(%rbp),%r14
	movzwq	(%rbp),%r15
	nop
	movzwq	(%rsi),%rax
	movzwq	(%rsi),%rcx
	movzwq	(%rsi),%rdx
	movzwq	(%rsi),%rbx
	movzwq	(%rsi),%rsp
	movzwq	(%rsi),%rbp
	movzwq	(%rsi),%rsi
	movzwq	(%rsi),%rdi
	movzwq	(%rsi),%r8
	movzwq	(%rsi),%r9
	movzwq	(%rsi),%r10
	movzwq	(%rsi),%r11
	movzwq	(%rsi),%r12
	movzwq	(%rsi),%r13
	movzwq	(%rsi),%r14
	movzwq	(%rsi),%r15
	nop
	movzwq	(%rdi),%rax
	movzwq	(%rdi),%rcx
	movzwq	(%rdi),%rdx
	movzwq	(%rdi),%rbx
	movzwq	(%rdi),%rsp
	movzwq	(%rdi),%rbp
	movzwq	(%rdi),%rsi
	movzwq	(%rdi),%rdi
	movzwq	(%rdi),%r8
	movzwq	(%rdi),%r9
	movzwq	(%rdi),%r10
	movzwq	(%rdi),%r11
	movzwq	(%rdi),%r12
	movzwq	(%rdi),%r13
	movzwq	(%rdi),%r14
	movzwq	(%rdi),%r15
	nop
	movzwq	(%r8), %rax
	movzwq	(%r8), %rcx
	movzwq	(%r8), %rdx
	movzwq	(%r8), %rbx
	movzwq	(%r8), %rsp
	movzwq	(%r8), %rbp
	movzwq	(%r8), %rsi
	movzwq	(%r8), %rdi
	movzwq	(%r8), %r8
	movzwq	(%r8), %r9
	movzwq	(%r8), %r10
	movzwq	(%r8), %r11
	movzwq	(%r8), %r12
	movzwq	(%r8), %r13
	movzwq	(%r8), %r14
	movzwq	(%r8), %r15
	nop
	movzwq	(%r9), %rax
	movzwq	(%r9), %rcx
	movzwq	(%r9), %rdx
	movzwq	(%r9), %rbx
	movzwq	(%r9), %rsp
	movzwq	(%r9), %rbp
	movzwq	(%r9), %rsi
	movzwq	(%r9), %rdi
	movzwq	(%r9), %r8
	movzwq	(%r9), %r9
	movzwq	(%r9), %r10
	movzwq	(%r9), %r11
	movzwq	(%r9), %r12
	movzwq	(%r9), %r13
	movzwq	(%r9), %r14
	movzwq	(%r9), %r15
	nop
	movzwq	(%r10),%rax
	movzwq	(%r10),%rcx
	movzwq	(%r10),%rdx
	movzwq	(%r10),%rbx
	movzwq	(%r10),%rsp
	movzwq	(%r10),%rbp
	movzwq	(%r10),%rsi
	movzwq	(%r10),%rdi
	movzwq	(%r10),%r8
	movzwq	(%r10),%r9
	movzwq	(%r10),%r10
	movzwq	(%r10),%r11
	movzwq	(%r10),%r12
	movzwq	(%r10),%r13
	movzwq	(%r10),%r14
	movzwq	(%r10),%r15
	nop
	movzwq	(%r11),%rax
	movzwq	(%r11),%rcx
	movzwq	(%r11),%rdx
	movzwq	(%r11),%rbx
	movzwq	(%r11),%rsp
	movzwq	(%r11),%rbp
	movzwq	(%r11),%rsi
	movzwq	(%r11),%rdi
	movzwq	(%r11),%r8
	movzwq	(%r11),%r9
	movzwq	(%r11),%r10
	movzwq	(%r11),%r11
	movzwq	(%r11),%r12
	movzwq	(%r11),%r13
	movzwq	(%r11),%r14
	movzwq	(%r11),%r15
	nop
	movzwq	(%r12),%rax
	movzwq	(%r12),%rcx
	movzwq	(%r12),%rdx
	movzwq	(%r12),%rbx
	movzwq	(%r12),%rsp
	movzwq	(%r12),%rbp
	movzwq	(%r12),%rsi
	movzwq	(%r12),%rdi
	movzwq	(%r12),%r8
	movzwq	(%r12),%r9
	movzwq	(%r12),%r10
	movzwq	(%r12),%r11
	movzwq	(%r12),%r12
	movzwq	(%r12),%r13
	movzwq	(%r12),%r14
	movzwq	(%r12),%r15
	nop
	movzwq	(%r13),%rax
	movzwq	(%r13),%rcx
	movzwq	(%r13),%rdx
	movzwq	(%r13),%rbx
	movzwq	(%r13),%rsp
	movzwq	(%r13),%rbp
	movzwq	(%r13),%rsi
	movzwq	(%r13),%rdi
	movzwq	(%r13),%r8
	movzwq	(%r13),%r9
	movzwq	(%r13),%r10
	movzwq	(%r13),%r11
	movzwq	(%r13),%r12
	movzwq	(%r13),%r13
	movzwq	(%r13),%r14
	movzwq	(%r13),%r15
	nop
	movzwq	(%r14),%rax
	movzwq	(%r14),%rcx
	movzwq	(%r14),%rdx
	movzwq	(%r14),%rbx
	movzwq	(%r14),%rsp
	movzwq	(%r14),%rbp
	movzwq	(%r14),%rsi
	movzwq	(%r14),%rdi
	movzwq	(%r14),%r8
	movzwq	(%r14),%r9
	movzwq	(%r14),%r10
	movzwq	(%r14),%r11
	movzwq	(%r14),%r12
	movzwq	(%r14),%r13
	movzwq	(%r14),%r14
	movzwq	(%r14),%r15
	nop
	movzwq	(%r15),%rax
	movzwq	(%r15),%rcx
	movzwq	(%r15),%rdx
	movzwq	(%r15),%rbx
	movzwq	(%r15),%rsp
	movzwq	(%r15),%rbp
	movzwq	(%r15),%rsi
	movzwq	(%r15),%rdi
	movzwq	(%r15),%r8
	movzwq	(%r15),%r9
	movzwq	(%r15),%r10
	movzwq	(%r15),%r11
	movzwq	(%r15),%r12
	movzwq	(%r15),%r13
	movzwq	(%r15),%r14
	movzwq	(%r15),%r15
	nop
	nop
	// off8(reg) -> reg
	movzwq	0x7F(%rax),%rax
	movzwq	0x7F(%rax),%rcx
	movzwq	0x7F(%rax),%rdx
	movzwq	0x7F(%rax),%rbx
	movzwq	0x7F(%rax),%rsp
	movzwq	0x7F(%rax),%rbp
	movzwq	0x7F(%rax),%rsi
	movzwq	0x7F(%rax),%rdi
	movzwq	0x7F(%rax),%r8
	movzwq	0x7F(%rax),%r9
	movzwq	0x7F(%rax),%r10
	movzwq	0x7F(%rax),%r11
	movzwq	0x7F(%rax),%r12
	movzwq	0x7F(%rax),%r13
	movzwq	0x7F(%rax),%r14
	movzwq	0x7F(%rax),%r15
	nop
	movzwq	0x7F(%rcx),%rax
	movzwq	0x7F(%rcx),%rcx
	movzwq	0x7F(%rcx),%rdx
	movzwq	0x7F(%rcx),%rbx
	movzwq	0x7F(%rcx),%rsp
	movzwq	0x7F(%rcx),%rbp
	movzwq	0x7F(%rcx),%rsi
	movzwq	0x7F(%rcx),%rdi
	movzwq	0x7F(%rcx),%r8
	movzwq	0x7F(%rcx),%r9
	movzwq	0x7F(%rcx),%r10
	movzwq	0x7F(%rcx),%r11
	movzwq	0x7F(%rcx),%r12
	movzwq	0x7F(%rcx),%r13
	movzwq	0x7F(%rcx),%r14
	movzwq	0x7F(%rcx),%r15
	nop
	movzwq	0x7F(%rdx),%rax
	movzwq	0x7F(%rdx),%rcx
	movzwq	0x7F(%rdx),%rdx
	movzwq	0x7F(%rdx),%rbx
	movzwq	0x7F(%rdx),%rsp
	movzwq	0x7F(%rdx),%rbp
	movzwq	0x7F(%rdx),%rsi
	movzwq	0x7F(%rdx),%rdi
	movzwq	0x7F(%rdx),%r8
	movzwq	0x7F(%rdx),%r9
	movzwq	0x7F(%rdx),%r10
	movzwq	0x7F(%rdx),%r11
	movzwq	0x7F(%rdx),%r12
	movzwq	0x7F(%rdx),%r13
	movzwq	0x7F(%rdx),%r14
	movzwq	0x7F(%rdx),%r15
	nop
	movzwq	0x7F(%rbx),%rax
	movzwq	0x7F(%rbx),%rcx
	movzwq	0x7F(%rbx),%rdx
	movzwq	0x7F(%rbx),%rbx
	movzwq	0x7F(%rbx),%rsp
	movzwq	0x7F(%rbx),%rbp
	movzwq	0x7F(%rbx),%rsi
	movzwq	0x7F(%rbx),%rdi
	movzwq	0x7F(%rbx),%r8
	movzwq	0x7F(%rbx),%r9
	movzwq	0x7F(%rbx),%r10
	movzwq	0x7F(%rbx),%r11
	movzwq	0x7F(%rbx),%r12
	movzwq	0x7F(%rbx),%r13
	movzwq	0x7F(%rbx),%r14
	movzwq	0x7F(%rbx),%r15
	nop
	movzwq	0x7F(%rsp),%rax
	movzwq	0x7F(%rsp),%rcx
	movzwq	0x7F(%rsp),%rdx
	movzwq	0x7F(%rsp),%rbx
	movzwq	0x7F(%rsp),%rsp
	movzwq	0x7F(%rsp),%rbp
	movzwq	0x7F(%rsp),%rsi
	movzwq	0x7F(%rsp),%rdi
	movzwq	0x7F(%rsp),%r8
	movzwq	0x7F(%rsp),%r9
	movzwq	0x7F(%rsp),%r10
	movzwq	0x7F(%rsp),%r11
	movzwq	0x7F(%rsp),%r12
	movzwq	0x7F(%rsp),%r13
	movzwq	0x7F(%rsp),%r14
	movzwq	0x7F(%rsp),%r15
	nop
	movzwq	0x7F(%rbp),%rax
	movzwq	0x7F(%rbp),%rcx
	movzwq	0x7F(%rbp),%rdx
	movzwq	0x7F(%rbp),%rbx
	movzwq	0x7F(%rbp),%rsp
	movzwq	0x7F(%rbp),%rbp
	movzwq	0x7F(%rbp),%rsi
	movzwq	0x7F(%rbp),%rdi
	movzwq	0x7F(%rbp),%r8
	movzwq	0x7F(%rbp),%r9
	movzwq	0x7F(%rbp),%r10
	movzwq	0x7F(%rbp),%r11
	movzwq	0x7F(%rbp),%r12
	movzwq	0x7F(%rbp),%r13
	movzwq	0x7F(%rbp),%r14
	movzwq	0x7F(%rbp),%r15
	nop
	movzwq	0x7F(%rsi),%rax
	movzwq	0x7F(%rsi),%rcx
	movzwq	0x7F(%rsi),%rdx
	movzwq	0x7F(%rsi),%rbx
	movzwq	0x7F(%rsi),%rsp
	movzwq	0x7F(%rsi),%rbp
	movzwq	0x7F(%rsi),%rsi
	movzwq	0x7F(%rsi),%rdi
	movzwq	0x7F(%rsi),%r8
	movzwq	0x7F(%rsi),%r9
	movzwq	0x7F(%rsi),%r10
	movzwq	0x7F(%rsi),%r11
	movzwq	0x7F(%rsi),%r12
	movzwq	0x7F(%rsi),%r13
	movzwq	0x7F(%rsi),%r14
	movzwq	0x7F(%rsi),%r15
	nop
	movzwq	0x7F(%rdi),%rax
	movzwq	0x7F(%rdi),%rcx
	movzwq	0x7F(%rdi),%rdx
	movzwq	0x7F(%rdi),%rbx
	movzwq	0x7F(%rdi),%rsp
	movzwq	0x7F(%rdi),%rbp
	movzwq	0x7F(%rdi),%rsi
	movzwq	0x7F(%rdi),%rdi
	movzwq	0x7F(%rdi),%r8
	movzwq	0x7F(%rdi),%r9
	movzwq	0x7F(%rdi),%r10
	movzwq	0x7F(%rdi),%r11
	movzwq	0x7F(%rdi),%r12
	movzwq	0x7F(%rdi),%r13
	movzwq	0x7F(%rdi),%r14
	movzwq	0x7F(%rdi),%r15
	nop
	movzwq	0x7F(%r8), %rax
	movzwq	0x7F(%r8), %rcx
	movzwq	0x7F(%r8), %rdx
	movzwq	0x7F(%r8), %rbx
	movzwq	0x7F(%r8), %rsp
	movzwq	0x7F(%r8), %rbp
	movzwq	0x7F(%r8), %rsi
	movzwq	0x7F(%r8), %rdi
	movzwq	0x7F(%r8), %r8
	movzwq	0x7F(%r8), %r9
	movzwq	0x7F(%r8), %r10
	movzwq	0x7F(%r8), %r11
	movzwq	0x7F(%r8), %r12
	movzwq	0x7F(%r8), %r13
	movzwq	0x7F(%r8), %r14
	movzwq	0x7F(%r8), %r15
	nop
	movzwq	0x7F(%r9), %rax
	movzwq	0x7F(%r9), %rcx
	movzwq	0x7F(%r9), %rdx
	movzwq	0x7F(%r9), %rbx
	movzwq	0x7F(%r9), %rsp
	movzwq	0x7F(%r9), %rbp
	movzwq	0x7F(%r9), %rsi
	movzwq	0x7F(%r9), %rdi
	movzwq	0x7F(%r9), %r8
	movzwq	0x7F(%r9), %r9
	movzwq	0x7F(%r9), %r10
	movzwq	0x7F(%r9), %r11
	movzwq	0x7F(%r9), %r12
	movzwq	0x7F(%r9), %r13
	movzwq	0x7F(%r9), %r14
	movzwq	0x7F(%r9), %r15
	nop
	movzwq	0x7F(%r10),%rax
	movzwq	0x7F(%r10),%rcx
	movzwq	0x7F(%r10),%rdx
	movzwq	0x7F(%r10),%rbx
	movzwq	0x7F(%r10),%rsp
	movzwq	0x7F(%r10),%rbp
	movzwq	0x7F(%r10),%rsi
	movzwq	0x7F(%r10),%rdi
	movzwq	0x7F(%r10),%r8
	movzwq	0x7F(%r10),%r9
	movzwq	0x7F(%r10),%r10
	movzwq	0x7F(%r10),%r11
	movzwq	0x7F(%r10),%r12
	movzwq	0x7F(%r10),%r13
	movzwq	0x7F(%r10),%r14
	movzwq	0x7F(%r10),%r15
	nop
	movzwq	0x7F(%r11),%rax
	movzwq	0x7F(%r11),%rcx
	movzwq	0x7F(%r11),%rdx
	movzwq	0x7F(%r11),%rbx
	movzwq	0x7F(%r11),%rsp
	movzwq	0x7F(%r11),%rbp
	movzwq	0x7F(%r11),%rsi
	movzwq	0x7F(%r11),%rdi
	movzwq	0x7F(%r11),%r8
	movzwq	0x7F(%r11),%r9
	movzwq	0x7F(%r11),%r10
	movzwq	0x7F(%r11),%r11
	movzwq	0x7F(%r11),%r12
	movzwq	0x7F(%r11),%r13
	movzwq	0x7F(%r11),%r14
	movzwq	0x7F(%r11),%r15
	nop
	movzwq	0x7F(%r12),%rax
	movzwq	0x7F(%r12),%rcx
	movzwq	0x7F(%r12),%rdx
	movzwq	0x7F(%r12),%rbx
	movzwq	0x7F(%r12),%rsp
	movzwq	0x7F(%r12),%rbp
	movzwq	0x7F(%r12),%rsi
	movzwq	0x7F(%r12),%rdi
	movzwq	0x7F(%r12),%r8
	movzwq	0x7F(%r12),%r9
	movzwq	0x7F(%r12),%r10
	movzwq	0x7F(%r12),%r11
	movzwq	0x7F(%r12),%r12
	movzwq	0x7F(%r12),%r13
	movzwq	0x7F(%r12),%r14
	movzwq	0x7F(%r12),%r15
	nop
	movzwq	0x7F(%r13),%rax
	movzwq	0x7F(%r13),%rcx
	movzwq	0x7F(%r13),%rdx
	movzwq	0x7F(%r13),%rbx
	movzwq	0x7F(%r13),%rsp
	movzwq	0x7F(%r13),%rbp
	movzwq	0x7F(%r13),%rsi
	movzwq	0x7F(%r13),%rdi
	movzwq	0x7F(%r13),%r8
	movzwq	0x7F(%r13),%r9
	movzwq	0x7F(%r13),%r10
	movzwq	0x7F(%r13),%r11
	movzwq	0x7F(%r13),%r12
	movzwq	0x7F(%r13),%r13
	movzwq	0x7F(%r13),%r14
	movzwq	0x7F(%r13),%r15
	nop
	movzwq	0x7F(%r14),%rax
	movzwq	0x7F(%r14),%rcx
	movzwq	0x7F(%r14),%rdx
	movzwq	0x7F(%r14),%rbx
	movzwq	0x7F(%r14),%rsp
	movzwq	0x7F(%r14),%rbp
	movzwq	0x7F(%r14),%rsi
	movzwq	0x7F(%r14),%rdi
	movzwq	0x7F(%r14),%r8
	movzwq	0x7F(%r14),%r9
	movzwq	0x7F(%r14),%r10
	movzwq	0x7F(%r14),%r11
	movzwq	0x7F(%r14),%r12
	movzwq	0x7F(%r14),%r13
	movzwq	0x7F(%r14),%r14
	movzwq	0x7F(%r14),%r15
	nop
	movzwq	0x7F(%r15),%rax
	movzwq	0x7F(%r15),%rcx
	movzwq	0x7F(%r15),%rdx
	movzwq	0x7F(%r15),%rbx
	movzwq	0x7F(%r15),%rsp
	movzwq	0x7F(%r15),%rbp
	movzwq	0x7F(%r15),%rsi
	movzwq	0x7F(%r15),%rdi
	movzwq	0x7F(%r15),%r8
	movzwq	0x7F(%r15),%r9
	movzwq	0x7F(%r15),%r10
	movzwq	0x7F(%r15),%r11
	movzwq	0x7F(%r15),%r12
	movzwq	0x7F(%r15),%r13
	movzwq	0x7F(%r15),%r14
	movzwq	0x7F(%r15),%r15
	nop
	nop
	// off32(reg) -> reg
	movzwq	0x12345678(%rax),%rax
	movzwq	0x12345678(%rax),%rcx
	movzwq	0x12345678(%rax),%rdx
	movzwq	0x12345678(%rax),%rbx
	movzwq	0x12345678(%rax),%rsp
	movzwq	0x12345678(%rax),%rbp
	movzwq	0x12345678(%rax),%rsi
	movzwq	0x12345678(%rax),%rdi
	movzwq	0x12345678(%rax),%r8
	movzwq	0x12345678(%rax),%r9
	movzwq	0x12345678(%rax),%r10
	movzwq	0x12345678(%rax),%r11
	movzwq	0x12345678(%rax),%r12
	movzwq	0x12345678(%rax),%r13
	movzwq	0x12345678(%rax),%r14
	movzwq	0x12345678(%rax),%r15
	nop
	movzwq	0x12345678(%rcx),%rax
	movzwq	0x12345678(%rcx),%rcx
	movzwq	0x12345678(%rcx),%rdx
	movzwq	0x12345678(%rcx),%rbx
	movzwq	0x12345678(%rcx),%rsp
	movzwq	0x12345678(%rcx),%rbp
	movzwq	0x12345678(%rcx),%rsi
	movzwq	0x12345678(%rcx),%rdi
	movzwq	0x12345678(%rcx),%r8
	movzwq	0x12345678(%rcx),%r9
	movzwq	0x12345678(%rcx),%r10
	movzwq	0x12345678(%rcx),%r11
	movzwq	0x12345678(%rcx),%r12
	movzwq	0x12345678(%rcx),%r13
	movzwq	0x12345678(%rcx),%r14
	movzwq	0x12345678(%rcx),%r15
	nop
	movzwq	0x12345678(%rdx),%rax
	movzwq	0x12345678(%rdx),%rcx
	movzwq	0x12345678(%rdx),%rdx
	movzwq	0x12345678(%rdx),%rbx
	movzwq	0x12345678(%rdx),%rsp
	movzwq	0x12345678(%rdx),%rbp
	movzwq	0x12345678(%rdx),%rsi
	movzwq	0x12345678(%rdx),%rdi
	movzwq	0x12345678(%rdx),%r8
	movzwq	0x12345678(%rdx),%r9
	movzwq	0x12345678(%rdx),%r10
	movzwq	0x12345678(%rdx),%r11
	movzwq	0x12345678(%rdx),%r12
	movzwq	0x12345678(%rdx),%r13
	movzwq	0x12345678(%rdx),%r14
	movzwq	0x12345678(%rdx),%r15
	nop
	movzwq	0x12345678(%rbx),%rax
	movzwq	0x12345678(%rbx),%rcx
	movzwq	0x12345678(%rbx),%rdx
	movzwq	0x12345678(%rbx),%rbx
	movzwq	0x12345678(%rbx),%rsp
	movzwq	0x12345678(%rbx),%rbp
	movzwq	0x12345678(%rbx),%rsi
	movzwq	0x12345678(%rbx),%rdi
	movzwq	0x12345678(%rbx),%r8
	movzwq	0x12345678(%rbx),%r9
	movzwq	0x12345678(%rbx),%r10
	movzwq	0x12345678(%rbx),%r11
	movzwq	0x12345678(%rbx),%r12
	movzwq	0x12345678(%rbx),%r13
	movzwq	0x12345678(%rbx),%r14
	movzwq	0x12345678(%rbx),%r15
	nop
	movzwq	0x12345678(%rsp),%rax
	movzwq	0x12345678(%rsp),%rcx
	movzwq	0x12345678(%rsp),%rdx
	movzwq	0x12345678(%rsp),%rbx
	movzwq	0x12345678(%rsp),%rsp
	movzwq	0x12345678(%rsp),%rbp
	movzwq	0x12345678(%rsp),%rsi
	movzwq	0x12345678(%rsp),%rdi
	movzwq	0x12345678(%rsp),%r8
	movzwq	0x12345678(%rsp),%r9
	movzwq	0x12345678(%rsp),%r10
	movzwq	0x12345678(%rsp),%r11
	movzwq	0x12345678(%rsp),%r12
	movzwq	0x12345678(%rsp),%r13
	movzwq	0x12345678(%rsp),%r14
	movzwq	0x12345678(%rsp),%r15
	nop
	movzwq	0x12345678(%rbp),%rax
	movzwq	0x12345678(%rbp),%rcx
	movzwq	0x12345678(%rbp),%rdx
	movzwq	0x12345678(%rbp),%rbx
	movzwq	0x12345678(%rbp),%rsp
	movzwq	0x12345678(%rbp),%rbp
	movzwq	0x12345678(%rbp),%rsi
	movzwq	0x12345678(%rbp),%rdi
	movzwq	0x12345678(%rbp),%r8
	movzwq	0x12345678(%rbp),%r9
	movzwq	0x12345678(%rbp),%r10
	movzwq	0x12345678(%rbp),%r11
	movzwq	0x12345678(%rbp),%r12
	movzwq	0x12345678(%rbp),%r13
	movzwq	0x12345678(%rbp),%r14
	movzwq	0x12345678(%rbp),%r15
	nop
	movzwq	0x12345678(%rsi),%rax
	movzwq	0x12345678(%rsi),%rcx
	movzwq	0x12345678(%rsi),%rdx
	movzwq	0x12345678(%rsi),%rbx
	movzwq	0x12345678(%rsi),%rsp
	movzwq	0x12345678(%rsi),%rbp
	movzwq	0x12345678(%rsi),%rsi
	movzwq	0x12345678(%rsi),%rdi
	movzwq	0x12345678(%rsi),%r8
	movzwq	0x12345678(%rsi),%r9
	movzwq	0x12345678(%rsi),%r10
	movzwq	0x12345678(%rsi),%r11
	movzwq	0x12345678(%rsi),%r12
	movzwq	0x12345678(%rsi),%r13
	movzwq	0x12345678(%rsi),%r14
	movzwq	0x12345678(%rsi),%r15
	nop
	movzwq	0x12345678(%rdi),%rax
	movzwq	0x12345678(%rdi),%rcx
	movzwq	0x12345678(%rdi),%rdx
	movzwq	0x12345678(%rdi),%rbx
	movzwq	0x12345678(%rdi),%rsp
	movzwq	0x12345678(%rdi),%rbp
	movzwq	0x12345678(%rdi),%rsi
	movzwq	0x12345678(%rdi),%rdi
	movzwq	0x12345678(%rdi),%r8
	movzwq	0x12345678(%rdi),%r9
	movzwq	0x12345678(%rdi),%r10
	movzwq	0x12345678(%rdi),%r11
	movzwq	0x12345678(%rdi),%r12
	movzwq	0x12345678(%rdi),%r13
	movzwq	0x12345678(%rdi),%r14
	movzwq	0x12345678(%rdi),%r15
	nop
	movzwq	0x12345678(%r8), %rax
	movzwq	0x12345678(%r8), %rcx
	movzwq	0x12345678(%r8), %rdx
	movzwq	0x12345678(%r8), %rbx
	movzwq	0x12345678(%r8), %rsp
	movzwq	0x12345678(%r8), %rbp
	movzwq	0x12345678(%r8), %rsi
	movzwq	0x12345678(%r8), %rdi
	movzwq	0x12345678(%r8), %r8
	movzwq	0x12345678(%r8), %r9
	movzwq	0x12345678(%r8), %r10
	movzwq	0x12345678(%r8), %r11
	movzwq	0x12345678(%r8), %r12
	movzwq	0x12345678(%r8), %r13
	movzwq	0x12345678(%r8), %r14
	movzwq	0x12345678(%r8), %r15
	nop
	movzwq	0x12345678(%r9), %rax
	movzwq	0x12345678(%r9), %rcx
	movzwq	0x12345678(%r9), %rdx
	movzwq	0x12345678(%r9), %rbx
	movzwq	0x12345678(%r9), %rsp
	movzwq	0x12345678(%r9), %rbp
	movzwq	0x12345678(%r9), %rsi
	movzwq	0x12345678(%r9), %rdi
	movzwq	0x12345678(%r9), %r8
	movzwq	0x12345678(%r9), %r9
	movzwq	0x12345678(%r9), %r10
	movzwq	0x12345678(%r9), %r11
	movzwq	0x12345678(%r9), %r12
	movzwq	0x12345678(%r9), %r13
	movzwq	0x12345678(%r9), %r14
	movzwq	0x12345678(%r9), %r15
	nop
	movzwq	0x12345678(%r10),%rax
	movzwq	0x12345678(%r10),%rcx
	movzwq	0x12345678(%r10),%rdx
	movzwq	0x12345678(%r10),%rbx
	movzwq	0x12345678(%r10),%rsp
	movzwq	0x12345678(%r10),%rbp
	movzwq	0x12345678(%r10),%rsi
	movzwq	0x12345678(%r10),%rdi
	movzwq	0x12345678(%r10),%r8
	movzwq	0x12345678(%r10),%r9
	movzwq	0x12345678(%r10),%r10
	movzwq	0x12345678(%r10),%r11
	movzwq	0x12345678(%r10),%r12
	movzwq	0x12345678(%r10),%r13
	movzwq	0x12345678(%r10),%r14
	movzwq	0x12345678(%r10),%r15
	nop
	movzwq	0x12345678(%r11),%rax
	movzwq	0x12345678(%r11),%rcx
	movzwq	0x12345678(%r11),%rdx
	movzwq	0x12345678(%r11),%rbx
	movzwq	0x12345678(%r11),%rsp
	movzwq	0x12345678(%r11),%rbp
	movzwq	0x12345678(%r11),%rsi
	movzwq	0x12345678(%r11),%rdi
	movzwq	0x12345678(%r11),%r8
	movzwq	0x12345678(%r11),%r9
	movzwq	0x12345678(%r11),%r10
	movzwq	0x12345678(%r11),%r11
	movzwq	0x12345678(%r11),%r12
	movzwq	0x12345678(%r11),%r13
	movzwq	0x12345678(%r11),%r14
	movzwq	0x12345678(%r11),%r15
	nop
	movzwq	0x12345678(%r12),%rax
	movzwq	0x12345678(%r12),%rcx
	movzwq	0x12345678(%r12),%rdx
	movzwq	0x12345678(%r12),%rbx
	movzwq	0x12345678(%r12),%rsp
	movzwq	0x12345678(%r12),%rbp
	movzwq	0x12345678(%r12),%rsi
	movzwq	0x12345678(%r12),%rdi
	movzwq	0x12345678(%r12),%r8
	movzwq	0x12345678(%r12),%r9
	movzwq	0x12345678(%r12),%r10
	movzwq	0x12345678(%r12),%r11
	movzwq	0x12345678(%r12),%r12
	movzwq	0x12345678(%r12),%r13
	movzwq	0x12345678(%r12),%r14
	movzwq	0x12345678(%r12),%r15
	nop
	movzwq	0x12345678(%r13),%rax
	movzwq	0x12345678(%r13),%rcx
	movzwq	0x12345678(%r13),%rdx
	movzwq	0x12345678(%r13),%rbx
	movzwq	0x12345678(%r13),%rsp
	movzwq	0x12345678(%r13),%rbp
	movzwq	0x12345678(%r13),%rsi
	movzwq	0x12345678(%r13),%rdi
	movzwq	0x12345678(%r13),%r8
	movzwq	0x12345678(%r13),%r9
	movzwq	0x12345678(%r13),%r10
	movzwq	0x12345678(%r13),%r11
	movzwq	0x12345678(%r13),%r12
	movzwq	0x12345678(%r13),%r13
	movzwq	0x12345678(%r13),%r14
	movzwq	0x12345678(%r13),%r15
	nop
	movzwq	0x12345678(%r14),%rax
	movzwq	0x12345678(%r14),%rcx
	movzwq	0x12345678(%r14),%rdx
	movzwq	0x12345678(%r14),%rbx
	movzwq	0x12345678(%r14),%rsp
	movzwq	0x12345678(%r14),%rbp
	movzwq	0x12345678(%r14),%rsi
	movzwq	0x12345678(%r14),%rdi
	movzwq	0x12345678(%r14),%r8
	movzwq	0x12345678(%r14),%r9
	movzwq	0x12345678(%r14),%r10
	movzwq	0x12345678(%r14),%r11
	movzwq	0x12345678(%r14),%r12
	movzwq	0x12345678(%r14),%r13
	movzwq	0x12345678(%r14),%r14
	movzwq	0x12345678(%r14),%r15
	nop
	movzwq	0x12345678(%r15),%rax
	movzwq	0x12345678(%r15),%rcx
	movzwq	0x12345678(%r15),%rdx
	movzwq	0x12345678(%r15),%rbx
	movzwq	0x12345678(%r15),%rsp
	movzwq	0x12345678(%r15),%rbp
	movzwq	0x12345678(%r15),%rsi
	movzwq	0x12345678(%r15),%rdi
	movzwq	0x12345678(%r15),%r8
	movzwq	0x12345678(%r15),%r9
	movzwq	0x12345678(%r15),%r10
	movzwq	0x12345678(%r15),%r11
	movzwq	0x12345678(%r15),%r12
	movzwq	0x12345678(%r15),%r13
	movzwq	0x12345678(%r15),%r14
	movzwq	0x12345678(%r15),%r15
	nop
	nop
	ret
	.cfi_endproc

	.p2align 4,,15
	.globl	Movswq
	.type	Movswq, @function
Movswq:
	.cfi_startproc
	// reg -> reg
	movswq	%ax,%rax
	movswq	%ax,%rcx
	movswq	%ax,%rdx
	movswq	%ax,%rbx
	movswq	%ax,%rsp
	movswq	%ax,%rbp
	movswq	%ax,%rsi
	movswq	%ax,%rdi
	movswq	%ax,%r8
	movswq	%ax,%r9
	movswq	%ax,%r10
	movswq	%ax,%r11
	movswq	%ax,%r12
	movswq	%ax,%r13
	movswq	%ax,%r14
	movswq	%ax,%r15
	nop
	movswq	%cx,%rax
	movswq	%cx,%rcx
	movswq	%cx,%rdx
	movswq	%cx,%rbx
	movswq	%cx,%rsp
	movswq	%cx,%rbp
	movswq	%cx,%rsi
	movswq	%cx,%rdi
	movswq	%cx,%r8
	movswq	%cx,%r9
	movswq	%cx,%r10
	movswq	%cx,%r11
	movswq	%cx,%r12
	movswq	%cx,%r13
	movswq	%cx,%r14
	movswq	%cx,%r15
	nop
	movswq	%dx,%rax
	movswq	%dx,%rcx
	movswq	%dx,%rdx
	movswq	%dx,%rbx
	movswq	%dx,%rsp
	movswq	%dx,%rbp
	movswq	%dx,%rsi
	movswq	%dx,%rdi
	movswq	%dx,%r8
	movswq	%dx,%r9
	movswq	%dx,%r10
	movswq	%dx,%r11
	movswq	%dx,%r12
	movswq	%dx,%r13
	movswq	%dx,%r14
	movswq	%dx,%r15
	nop
	movswq	%bx,%rax
	movswq	%bx,%rcx
	movswq	%bx,%rdx
	movswq	%bx,%rbx
	movswq	%bx,%rsp
	movswq	%bx,%rbp
	movswq	%bx,%rsi
	movswq	%bx,%rdi
	movswq	%bx,%r8
	movswq	%bx,%r9
	movswq	%bx,%r10
	movswq	%bx,%r11
	movswq	%bx,%r12
	movswq	%bx,%r13
	movswq	%bx,%r14
	movswq	%bx,%r15
	nop
	movswq	%sp,%rax
	movswq	%sp,%rcx
	movswq	%sp,%rdx
	movswq	%sp,%rbx
	movswq	%sp,%rsp
	movswq	%sp,%rbp
	movswq	%sp,%rsi
	movswq	%sp,%rdi
	movswq	%sp,%r8
	movswq	%sp,%r9
	movswq	%sp,%r10
	movswq	%sp,%r11
	movswq	%sp,%r12
	movswq	%sp,%r13
	movswq	%sp,%r14
	movswq	%sp,%r15
	nop
	movswq	%bp,%rax
	movswq	%bp,%rcx
	movswq	%bp,%rdx
	movswq	%bp,%rbx
	movswq	%bp,%rsp
	movswq	%bp,%rbp
	movswq	%bp,%rsi
	movswq	%bp,%rdi
	movswq	%bp,%r8
	movswq	%bp,%r9
	movswq	%bp,%r10
	movswq	%bp,%r11
	movswq	%bp,%r12
	movswq	%bp,%r13
	movswq	%bp,%r14
	movswq	%bp,%r15
	nop
	movswq	%si,%rax
	movswq	%si,%rcx
	movswq	%si,%rdx
	movswq	%si,%rbx
	movswq	%si,%rsp
	movswq	%si,%rbp
	movswq	%si,%rsi
	movswq	%si,%rdi
	movswq	%si,%r8
	movswq	%si,%r9
	movswq	%si,%r10
	movswq	%si,%r11
	movswq	%si,%r12
	movswq	%si,%r13
	movswq	%si,%r14
	movswq	%si,%r15
	nop
	movswq	%di,%rax
	movswq	%di,%rcx
	movswq	%di,%rdx
	movswq	%di,%rbx
	movswq	%di,%rsp
	movswq	%di,%rbp
	movswq	%di,%rsi
	movswq	%di,%rdi
	movswq	%di,%r8
	movswq	%di,%r9
	movswq	%di,%r10
	movswq	%di,%r11
	movswq	%di,%r12
	movswq	%di,%r13
	movswq	%di,%r14
	movswq	%di,%r15
	nop
	movswq	%r8w, %rax
	movswq	%r8w, %rcx
	movswq	%r8w, %rdx
	movswq	%r8w, %rbx
	movswq	%r8w, %rsp
	movswq	%r8w, %rbp
	movswq	%r8w, %rsi
	movswq	%r8w, %rdi
	movswq	%r8w, %r8
	movswq	%r8w, %r9
	movswq	%r8w, %r10
	movswq	%r8w, %r11
	movswq	%r8w, %r12
	movswq	%r8w, %r13
	movswq	%r8w, %r14
	movswq	%r8w, %r15
	nop
	movswq	%r9w, %rax
	movswq	%r9w, %rcx
	movswq	%r9w, %rdx
	movswq	%r9w, %rbx
	movswq	%r9w, %rsp
	movswq	%r9w, %rbp
	movswq	%r9w, %rsi
	movswq	%r9w, %rdi
	movswq	%r9w, %r8
	movswq	%r9w, %r9
	movswq	%r9w, %r10
	movswq	%r9w, %r11
	movswq	%r9w, %r12
	movswq	%r9w, %r13
	movswq	%r9w, %r14
	movswq	%r9w, %r15
	nop
	movswq	%r10w,%rax
	movswq	%r10w,%rcx
	movswq	%r10w,%rdx
	movswq	%r10w,%rbx
	movswq	%r10w,%rsp
	movswq	%r10w,%rbp
	movswq	%r10w,%rsi
	movswq	%r10w,%rdi
	movswq	%r10w,%r8
	movswq	%r10w,%r9
	movswq	%r10w,%r10
	movswq	%r10w,%r11
	movswq	%r10w,%r12
	movswq	%r10w,%r13
	movswq	%r10w,%r14
	movswq	%r10w,%r15
	nop
	movswq	%r11w,%rax
	movswq	%r11w,%rcx
	movswq	%r11w,%rdx
	movswq	%r11w,%rbx
	movswq	%r11w,%rsp
	movswq	%r11w,%rbp
	movswq	%r11w,%rsi
	movswq	%r11w,%rdi
	movswq	%r11w,%r8
	movswq	%r11w,%r9
	movswq	%r11w,%r10
	movswq	%r11w,%r11
	movswq	%r11w,%r12
	movswq	%r11w,%r13
	movswq	%r11w,%r14
	movswq	%r11w,%r15
	nop
	movswq	%r12w,%rax
	movswq	%r12w,%rcx
	movswq	%r12w,%rdx
	movswq	%r12w,%rbx
	movswq	%r12w,%rsp
	movswq	%r12w,%rbp
	movswq	%r12w,%rsi
	movswq	%r12w,%rdi
	movswq	%r12w,%r8
	movswq	%r12w,%r9
	movswq	%r12w,%r10
	movswq	%r12w,%r11
	movswq	%r12w,%r12
	movswq	%r12w,%r13
	movswq	%r12w,%r14
	movswq	%r12w,%r15
	nop
	movswq	%r13w,%rax
	movswq	%r13w,%rcx
	movswq	%r13w,%rdx
	movswq	%r13w,%rbx
	movswq	%r13w,%rsp
	movswq	%r13w,%rbp
	movswq	%r13w,%rsi
	movswq	%r13w,%rdi
	movswq	%r13w,%r8
	movswq	%r13w,%r9
	movswq	%r13w,%r10
	movswq	%r13w,%r11
	movswq	%r13w,%r12
	movswq	%r13w,%r13
	movswq	%r13w,%r14
	movswq	%r13w,%r15
	nop
	movswq	%r14w,%rax
	movswq	%r14w,%rcx
	movswq	%r14w,%rdx
	movswq	%r14w,%rbx
	movswq	%r14w,%rsp
	movswq	%r14w,%rbp
	movswq	%r14w,%rsi
	movswq	%r14w,%rdi
	movswq	%r14w,%r8
	movswq	%r14w,%r9
	movswq	%r14w,%r10
	movswq	%r14w,%r11
	movswq	%r14w,%r12
	movswq	%r14w,%r13
	movswq	%r14w,%r14
	movswq	%r14w,%r15
	nop
	movswq	%r15w,%rax
	movswq	%r15w,%rcx
	movswq	%r15w,%rdx
	movswq	%r15w,%rbx
	movswq	%r15w,%rsp
	movswq	%r15w,%rbp
	movswq	%r15w,%rsi
	movswq	%r15w,%rdi
	movswq	%r15w,%r8
	movswq	%r15w,%r9
	movswq	%r15w,%r10
	movswq	%r15w,%r11
	movswq	%r15w,%r12
	movswq	%r15w,%r13
	movswq	%r15w,%r14
	movswq	%r15w,%r15
	nop
        nop
	// (reg) -> reg
	movswq	(%rax),%rax
	movswq	(%rax),%rcx
	movswq	(%rax),%rdx
	movswq	(%rax),%rbx
	movswq	(%rax),%rsp
	movswq	(%rax),%rbp
	movswq	(%rax),%rsi
	movswq	(%rax),%rdi
	movswq	(%rax),%r8
	movswq	(%rax),%r9
	movswq	(%rax),%r10
	movswq	(%rax),%r11
	movswq	(%rax),%r12
	movswq	(%rax),%r13
	movswq	(%rax),%r14
	movswq	(%rax),%r15
	nop
	movswq	(%rcx),%rax
	movswq	(%rcx),%rcx
	movswq	(%rcx),%rdx
	movswq	(%rcx),%rbx
	movswq	(%rcx),%rsp
	movswq	(%rcx),%rbp
	movswq	(%rcx),%rsi
	movswq	(%rcx),%rdi
	movswq	(%rcx),%r8
	movswq	(%rcx),%r9
	movswq	(%rcx),%r10
	movswq	(%rcx),%r11
	movswq	(%rcx),%r12
	movswq	(%rcx),%r13
	movswq	(%rcx),%r14
	movswq	(%rcx),%r15
	nop
	movswq	(%rdx),%rax
	movswq	(%rdx),%rcx
	movswq	(%rdx),%rdx
	movswq	(%rdx),%rbx
	movswq	(%rdx),%rsp
	movswq	(%rdx),%rbp
	movswq	(%rdx),%rsi
	movswq	(%rdx),%rdi
	movswq	(%rdx),%r8
	movswq	(%rdx),%r9
	movswq	(%rdx),%r10
	movswq	(%rdx),%r11
	movswq	(%rdx),%r12
	movswq	(%rdx),%r13
	movswq	(%rdx),%r14
	movswq	(%rdx),%r15
	nop
	movswq	(%rbx),%rax
	movswq	(%rbx),%rcx
	movswq	(%rbx),%rdx
	movswq	(%rbx),%rbx
	movswq	(%rbx),%rsp
	movswq	(%rbx),%rbp
	movswq	(%rbx),%rsi
	movswq	(%rbx),%rdi
	movswq	(%rbx),%r8
	movswq	(%rbx),%r9
	movswq	(%rbx),%r10
	movswq	(%rbx),%r11
	movswq	(%rbx),%r12
	movswq	(%rbx),%r13
	movswq	(%rbx),%r14
	movswq	(%rbx),%r15
	nop
	movswq	(%rsp),%rax
	movswq	(%rsp),%rcx
	movswq	(%rsp),%rdx
	movswq	(%rsp),%rbx
	movswq	(%rsp),%rsp
	movswq	(%rsp),%rbp
	movswq	(%rsp),%rsi
	movswq	(%rsp),%rdi
	movswq	(%rsp),%r8
	movswq	(%rsp),%r9
	movswq	(%rsp),%r10
	movswq	(%rsp),%r11
	movswq	(%rsp),%r12
	movswq	(%rsp),%r13
	movswq	(%rsp),%r14
	movswq	(%rsp),%r15
	nop
	movswq	(%rbp),%rax
	movswq	(%rbp),%rcx
	movswq	(%rbp),%rdx
	movswq	(%rbp),%rbx
	movswq	(%rbp),%rsp
	movswq	(%rbp),%rbp
	movswq	(%rbp),%rsi
	movswq	(%rbp),%rdi
	movswq	(%rbp),%r8
	movswq	(%rbp),%r9
	movswq	(%rbp),%r10
	movswq	(%rbp),%r11
	movswq	(%rbp),%r12
	movswq	(%rbp),%r13
	movswq	(%rbp),%r14
	movswq	(%rbp),%r15
	nop
	movswq	(%rsi),%rax
	movswq	(%rsi),%rcx
	movswq	(%rsi),%rdx
	movswq	(%rsi),%rbx
	movswq	(%rsi),%rsp
	movswq	(%rsi),%rbp
	movswq	(%rsi),%rsi
	movswq	(%rsi),%rdi
	movswq	(%rsi),%r8
	movswq	(%rsi),%r9
	movswq	(%rsi),%r10
	movswq	(%rsi),%r11
	movswq	(%rsi),%r12
	movswq	(%rsi),%r13
	movswq	(%rsi),%r14
	movswq	(%rsi),%r15
	nop
	movswq	(%rdi),%rax
	movswq	(%rdi),%rcx
	movswq	(%rdi),%rdx
	movswq	(%rdi),%rbx
	movswq	(%rdi),%rsp
	movswq	(%rdi),%rbp
	movswq	(%rdi),%rsi
	movswq	(%rdi),%rdi
	movswq	(%rdi),%r8
	movswq	(%rdi),%r9
	movswq	(%rdi),%r10
	movswq	(%rdi),%r11
	movswq	(%rdi),%r12
	movswq	(%rdi),%r13
	movswq	(%rdi),%r14
	movswq	(%rdi),%r15
	nop
	movswq	(%r8), %rax
	movswq	(%r8), %rcx
	movswq	(%r8), %rdx
	movswq	(%r8), %rbx
	movswq	(%r8), %rsp
	movswq	(%r8), %rbp
	movswq	(%r8), %rsi
	movswq	(%r8), %rdi
	movswq	(%r8), %r8
	movswq	(%r8), %r9
	movswq	(%r8), %r10
	movswq	(%r8), %r11
	movswq	(%r8), %r12
	movswq	(%r8), %r13
	movswq	(%r8), %r14
	movswq	(%r8), %r15
	nop
	movswq	(%r9), %rax
	movswq	(%r9), %rcx
	movswq	(%r9), %rdx
	movswq	(%r9), %rbx
	movswq	(%r9), %rsp
	movswq	(%r9), %rbp
	movswq	(%r9), %rsi
	movswq	(%r9), %rdi
	movswq	(%r9), %r8
	movswq	(%r9), %r9
	movswq	(%r9), %r10
	movswq	(%r9), %r11
	movswq	(%r9), %r12
	movswq	(%r9), %r13
	movswq	(%r9), %r14
	movswq	(%r9), %r15
	nop
	movswq	(%r10),%rax
	movswq	(%r10),%rcx
	movswq	(%r10),%rdx
	movswq	(%r10),%rbx
	movswq	(%r10),%rsp
	movswq	(%r10),%rbp
	movswq	(%r10),%rsi
	movswq	(%r10),%rdi
	movswq	(%r10),%r8
	movswq	(%r10),%r9
	movswq	(%r10),%r10
	movswq	(%r10),%r11
	movswq	(%r10),%r12
	movswq	(%r10),%r13
	movswq	(%r10),%r14
	movswq	(%r10),%r15
	nop
	movswq	(%r11),%rax
	movswq	(%r11),%rcx
	movswq	(%r11),%rdx
	movswq	(%r11),%rbx
	movswq	(%r11),%rsp
	movswq	(%r11),%rbp
	movswq	(%r11),%rsi
	movswq	(%r11),%rdi
	movswq	(%r11),%r8
	movswq	(%r11),%r9
	movswq	(%r11),%r10
	movswq	(%r11),%r11
	movswq	(%r11),%r12
	movswq	(%r11),%r13
	movswq	(%r11),%r14
	movswq	(%r11),%r15
	nop
	movswq	(%r12),%rax
	movswq	(%r12),%rcx
	movswq	(%r12),%rdx
	movswq	(%r12),%rbx
	movswq	(%r12),%rsp
	movswq	(%r12),%rbp
	movswq	(%r12),%rsi
	movswq	(%r12),%rdi
	movswq	(%r12),%r8
	movswq	(%r12),%r9
	movswq	(%r12),%r10
	movswq	(%r12),%r11
	movswq	(%r12),%r12
	movswq	(%r12),%r13
	movswq	(%r12),%r14
	movswq	(%r12),%r15
	nop
	movswq	(%r13),%rax
	movswq	(%r13),%rcx
	movswq	(%r13),%rdx
	movswq	(%r13),%rbx
	movswq	(%r13),%rsp
	movswq	(%r13),%rbp
	movswq	(%r13),%rsi
	movswq	(%r13),%rdi
	movswq	(%r13),%r8
	movswq	(%r13),%r9
	movswq	(%r13),%r10
	movswq	(%r13),%r11
	movswq	(%r13),%r12
	movswq	(%r13),%r13
	movswq	(%r13),%r14
	movswq	(%r13),%r15
	nop
	movswq	(%r14),%rax
	movswq	(%r14),%rcx
	movswq	(%r14),%rdx
	movswq	(%r14),%rbx
	movswq	(%r14),%rsp
	movswq	(%r14),%rbp
	movswq	(%r14),%rsi
	movswq	(%r14),%rdi
	movswq	(%r14),%r8
	movswq	(%r14),%r9
	movswq	(%r14),%r10
	movswq	(%r14),%r11
	movswq	(%r14),%r12
	movswq	(%r14),%r13
	movswq	(%r14),%r14
	movswq	(%r14),%r15
	nop
	movswq	(%r15),%rax
	movswq	(%r15),%rcx
	movswq	(%r15),%rdx
	movswq	(%r15),%rbx
	movswq	(%r15),%rsp
	movswq	(%r15),%rbp
	movswq	(%r15),%rsi
	movswq	(%r15),%rdi
	movswq	(%r15),%r8
	movswq	(%r15),%r9
	movswq	(%r15),%r10
	movswq	(%r15),%r11
	movswq	(%r15),%r12
	movswq	(%r15),%r13
	movswq	(%r15),%r14
	movswq	(%r15),%r15
	nop
	nop
	// off8(reg) -> reg
	movswq	0x7F(%rax),%rax
	movswq	0x7F(%rax),%rcx
	movswq	0x7F(%rax),%rdx
	movswq	0x7F(%rax),%rbx
	movswq	0x7F(%rax),%rsp
	movswq	0x7F(%rax),%rbp
	movswq	0x7F(%rax),%rsi
	movswq	0x7F(%rax),%rdi
	movswq	0x7F(%rax),%r8
	movswq	0x7F(%rax),%r9
	movswq	0x7F(%rax),%r10
	movswq	0x7F(%rax),%r11
	movswq	0x7F(%rax),%r12
	movswq	0x7F(%rax),%r13
	movswq	0x7F(%rax),%r14
	movswq	0x7F(%rax),%r15
	nop
	movswq	0x7F(%rcx),%rax
	movswq	0x7F(%rcx),%rcx
	movswq	0x7F(%rcx),%rdx
	movswq	0x7F(%rcx),%rbx
	movswq	0x7F(%rcx),%rsp
	movswq	0x7F(%rcx),%rbp
	movswq	0x7F(%rcx),%rsi
	movswq	0x7F(%rcx),%rdi
	movswq	0x7F(%rcx),%r8
	movswq	0x7F(%rcx),%r9
	movswq	0x7F(%rcx),%r10
	movswq	0x7F(%rcx),%r11
	movswq	0x7F(%rcx),%r12
	movswq	0x7F(%rcx),%r13
	movswq	0x7F(%rcx),%r14
	movswq	0x7F(%rcx),%r15
	nop
	movswq	0x7F(%rdx),%rax
	movswq	0x7F(%rdx),%rcx
	movswq	0x7F(%rdx),%rdx
	movswq	0x7F(%rdx),%rbx
	movswq	0x7F(%rdx),%rsp
	movswq	0x7F(%rdx),%rbp
	movswq	0x7F(%rdx),%rsi
	movswq	0x7F(%rdx),%rdi
	movswq	0x7F(%rdx),%r8
	movswq	0x7F(%rdx),%r9
	movswq	0x7F(%rdx),%r10
	movswq	0x7F(%rdx),%r11
	movswq	0x7F(%rdx),%r12
	movswq	0x7F(%rdx),%r13
	movswq	0x7F(%rdx),%r14
	movswq	0x7F(%rdx),%r15
	nop
	movswq	0x7F(%rbx),%rax
	movswq	0x7F(%rbx),%rcx
	movswq	0x7F(%rbx),%rdx
	movswq	0x7F(%rbx),%rbx
	movswq	0x7F(%rbx),%rsp
	movswq	0x7F(%rbx),%rbp
	movswq	0x7F(%rbx),%rsi
	movswq	0x7F(%rbx),%rdi
	movswq	0x7F(%rbx),%r8
	movswq	0x7F(%rbx),%r9
	movswq	0x7F(%rbx),%r10
	movswq	0x7F(%rbx),%r11
	movswq	0x7F(%rbx),%r12
	movswq	0x7F(%rbx),%r13
	movswq	0x7F(%rbx),%r14
	movswq	0x7F(%rbx),%r15
	nop
	movswq	0x7F(%rsp),%rax
	movswq	0x7F(%rsp),%rcx
	movswq	0x7F(%rsp),%rdx
	movswq	0x7F(%rsp),%rbx
	movswq	0x7F(%rsp),%rsp
	movswq	0x7F(%rsp),%rbp
	movswq	0x7F(%rsp),%rsi
	movswq	0x7F(%rsp),%rdi
	movswq	0x7F(%rsp),%r8
	movswq	0x7F(%rsp),%r9
	movswq	0x7F(%rsp),%r10
	movswq	0x7F(%rsp),%r11
	movswq	0x7F(%rsp),%r12
	movswq	0x7F(%rsp),%r13
	movswq	0x7F(%rsp),%r14
	movswq	0x7F(%rsp),%r15
	nop
	movswq	0x7F(%rbp),%rax
	movswq	0x7F(%rbp),%rcx
	movswq	0x7F(%rbp),%rdx
	movswq	0x7F(%rbp),%rbx
	movswq	0x7F(%rbp),%rsp
	movswq	0x7F(%rbp),%rbp
	movswq	0x7F(%rbp),%rsi
	movswq	0x7F(%rbp),%rdi
	movswq	0x7F(%rbp),%r8
	movswq	0x7F(%rbp),%r9
	movswq	0x7F(%rbp),%r10
	movswq	0x7F(%rbp),%r11
	movswq	0x7F(%rbp),%r12
	movswq	0x7F(%rbp),%r13
	movswq	0x7F(%rbp),%r14
	movswq	0x7F(%rbp),%r15
	nop
	movswq	0x7F(%rsi),%rax
	movswq	0x7F(%rsi),%rcx
	movswq	0x7F(%rsi),%rdx
	movswq	0x7F(%rsi),%rbx
	movswq	0x7F(%rsi),%rsp
	movswq	0x7F(%rsi),%rbp
	movswq	0x7F(%rsi),%rsi
	movswq	0x7F(%rsi),%rdi
	movswq	0x7F(%rsi),%r8
	movswq	0x7F(%rsi),%r9
	movswq	0x7F(%rsi),%r10
	movswq	0x7F(%rsi),%r11
	movswq	0x7F(%rsi),%r12
	movswq	0x7F(%rsi),%r13
	movswq	0x7F(%rsi),%r14
	movswq	0x7F(%rsi),%r15
	nop
	movswq	0x7F(%rdi),%rax
	movswq	0x7F(%rdi),%rcx
	movswq	0x7F(%rdi),%rdx
	movswq	0x7F(%rdi),%rbx
	movswq	0x7F(%rdi),%rsp
	movswq	0x7F(%rdi),%rbp
	movswq	0x7F(%rdi),%rsi
	movswq	0x7F(%rdi),%rdi
	movswq	0x7F(%rdi),%r8
	movswq	0x7F(%rdi),%r9
	movswq	0x7F(%rdi),%r10
	movswq	0x7F(%rdi),%r11
	movswq	0x7F(%rdi),%r12
	movswq	0x7F(%rdi),%r13
	movswq	0x7F(%rdi),%r14
	movswq	0x7F(%rdi),%r15
	nop
	movswq	0x7F(%r8), %rax
	movswq	0x7F(%r8), %rcx
	movswq	0x7F(%r8), %rdx
	movswq	0x7F(%r8), %rbx
	movswq	0x7F(%r8), %rsp
	movswq	0x7F(%r8), %rbp
	movswq	0x7F(%r8), %rsi
	movswq	0x7F(%r8), %rdi
	movswq	0x7F(%r8), %r8
	movswq	0x7F(%r8), %r9
	movswq	0x7F(%r8), %r10
	movswq	0x7F(%r8), %r11
	movswq	0x7F(%r8), %r12
	movswq	0x7F(%r8), %r13
	movswq	0x7F(%r8), %r14
	movswq	0x7F(%r8), %r15
	nop
	movswq	0x7F(%r9), %rax
	movswq	0x7F(%r9), %rcx
	movswq	0x7F(%r9), %rdx
	movswq	0x7F(%r9), %rbx
	movswq	0x7F(%r9), %rsp
	movswq	0x7F(%r9), %rbp
	movswq	0x7F(%r9), %rsi
	movswq	0x7F(%r9), %rdi
	movswq	0x7F(%r9), %r8
	movswq	0x7F(%r9), %r9
	movswq	0x7F(%r9), %r10
	movswq	0x7F(%r9), %r11
	movswq	0x7F(%r9), %r12
	movswq	0x7F(%r9), %r13
	movswq	0x7F(%r9), %r14
	movswq	0x7F(%r9), %r15
	nop
	movswq	0x7F(%r10),%rax
	movswq	0x7F(%r10),%rcx
	movswq	0x7F(%r10),%rdx
	movswq	0x7F(%r10),%rbx
	movswq	0x7F(%r10),%rsp
	movswq	0x7F(%r10),%rbp
	movswq	0x7F(%r10),%rsi
	movswq	0x7F(%r10),%rdi
	movswq	0x7F(%r10),%r8
	movswq	0x7F(%r10),%r9
	movswq	0x7F(%r10),%r10
	movswq	0x7F(%r10),%r11
	movswq	0x7F(%r10),%r12
	movswq	0x7F(%r10),%r13
	movswq	0x7F(%r10),%r14
	movswq	0x7F(%r10),%r15
	nop
	movswq	0x7F(%r11),%rax
	movswq	0x7F(%r11),%rcx
	movswq	0x7F(%r11),%rdx
	movswq	0x7F(%r11),%rbx
	movswq	0x7F(%r11),%rsp
	movswq	0x7F(%r11),%rbp
	movswq	0x7F(%r11),%rsi
	movswq	0x7F(%r11),%rdi
	movswq	0x7F(%r11),%r8
	movswq	0x7F(%r11),%r9
	movswq	0x7F(%r11),%r10
	movswq	0x7F(%r11),%r11
	movswq	0x7F(%r11),%r12
	movswq	0x7F(%r11),%r13
	movswq	0x7F(%r11),%r14
	movswq	0x7F(%r11),%r15
	nop
	movswq	0x7F(%r12),%rax
	movswq	0x7F(%r12),%rcx
	movswq	0x7F(%r12),%rdx
	movswq	0x7F(%r12),%rbx
	movswq	0x7F(%r12),%rsp
	movswq	0x7F(%r12),%rbp
	movswq	0x7F(%r12),%rsi
	movswq	0x7F(%r12),%rdi
	movswq	0x7F(%r12),%r8
	movswq	0x7F(%r12),%r9
	movswq	0x7F(%r12),%r10
	movswq	0x7F(%r12),%r11
	movswq	0x7F(%r12),%r12
	movswq	0x7F(%r12),%r13
	movswq	0x7F(%r12),%r14
	movswq	0x7F(%r12),%r15
	nop
	movswq	0x7F(%r13),%rax
	movswq	0x7F(%r13),%rcx
	movswq	0x7F(%r13),%rdx
	movswq	0x7F(%r13),%rbx
	movswq	0x7F(%r13),%rsp
	movswq	0x7F(%r13),%rbp
	movswq	0x7F(%r13),%rsi
	movswq	0x7F(%r13),%rdi
	movswq	0x7F(%r13),%r8
	movswq	0x7F(%r13),%r9
	movswq	0x7F(%r13),%r10
	movswq	0x7F(%r13),%r11
	movswq	0x7F(%r13),%r12
	movswq	0x7F(%r13),%r13
	movswq	0x7F(%r13),%r14
	movswq	0x7F(%r13),%r15
	nop
	movswq	0x7F(%r14),%rax
	movswq	0x7F(%r14),%rcx
	movswq	0x7F(%r14),%rdx
	movswq	0x7F(%r14),%rbx
	movswq	0x7F(%r14),%rsp
	movswq	0x7F(%r14),%rbp
	movswq	0x7F(%r14),%rsi
	movswq	0x7F(%r14),%rdi
	movswq	0x7F(%r14),%r8
	movswq	0x7F(%r14),%r9
	movswq	0x7F(%r14),%r10
	movswq	0x7F(%r14),%r11
	movswq	0x7F(%r14),%r12
	movswq	0x7F(%r14),%r13
	movswq	0x7F(%r14),%r14
	movswq	0x7F(%r14),%r15
	nop
	movswq	0x7F(%r15),%rax
	movswq	0x7F(%r15),%rcx
	movswq	0x7F(%r15),%rdx
	movswq	0x7F(%r15),%rbx
	movswq	0x7F(%r15),%rsp
	movswq	0x7F(%r15),%rbp
	movswq	0x7F(%r15),%rsi
	movswq	0x7F(%r15),%rdi
	movswq	0x7F(%r15),%r8
	movswq	0x7F(%r15),%r9
	movswq	0x7F(%r15),%r10
	movswq	0x7F(%r15),%r11
	movswq	0x7F(%r15),%r12
	movswq	0x7F(%r15),%r13
	movswq	0x7F(%r15),%r14
	movswq	0x7F(%r15),%r15
	nop
	nop
	// off32(reg) -> reg
	movswq	0x12345678(%rax),%rax
	movswq	0x12345678(%rax),%rcx
	movswq	0x12345678(%rax),%rdx
	movswq	0x12345678(%rax),%rbx
	movswq	0x12345678(%rax),%rsp
	movswq	0x12345678(%rax),%rbp
	movswq	0x12345678(%rax),%rsi
	movswq	0x12345678(%rax),%rdi
	movswq	0x12345678(%rax),%r8
	movswq	0x12345678(%rax),%r9
	movswq	0x12345678(%rax),%r10
	movswq	0x12345678(%rax),%r11
	movswq	0x12345678(%rax),%r12
	movswq	0x12345678(%rax),%r13
	movswq	0x12345678(%rax),%r14
	movswq	0x12345678(%rax),%r15
	nop
	movswq	0x12345678(%rcx),%rax
	movswq	0x12345678(%rcx),%rcx
	movswq	0x12345678(%rcx),%rdx
	movswq	0x12345678(%rcx),%rbx
	movswq	0x12345678(%rcx),%rsp
	movswq	0x12345678(%rcx),%rbp
	movswq	0x12345678(%rcx),%rsi
	movswq	0x12345678(%rcx),%rdi
	movswq	0x12345678(%rcx),%r8
	movswq	0x12345678(%rcx),%r9
	movswq	0x12345678(%rcx),%r10
	movswq	0x12345678(%rcx),%r11
	movswq	0x12345678(%rcx),%r12
	movswq	0x12345678(%rcx),%r13
	movswq	0x12345678(%rcx),%r14
	movswq	0x12345678(%rcx),%r15
	nop
	movswq	0x12345678(%rdx),%rax
	movswq	0x12345678(%rdx),%rcx
	movswq	0x12345678(%rdx),%rdx
	movswq	0x12345678(%rdx),%rbx
	movswq	0x12345678(%rdx),%rsp
	movswq	0x12345678(%rdx),%rbp
	movswq	0x12345678(%rdx),%rsi
	movswq	0x12345678(%rdx),%rdi
	movswq	0x12345678(%rdx),%r8
	movswq	0x12345678(%rdx),%r9
	movswq	0x12345678(%rdx),%r10
	movswq	0x12345678(%rdx),%r11
	movswq	0x12345678(%rdx),%r12
	movswq	0x12345678(%rdx),%r13
	movswq	0x12345678(%rdx),%r14
	movswq	0x12345678(%rdx),%r15
	nop
	movswq	0x12345678(%rbx),%rax
	movswq	0x12345678(%rbx),%rcx
	movswq	0x12345678(%rbx),%rdx
	movswq	0x12345678(%rbx),%rbx
	movswq	0x12345678(%rbx),%rsp
	movswq	0x12345678(%rbx),%rbp
	movswq	0x12345678(%rbx),%rsi
	movswq	0x12345678(%rbx),%rdi
	movswq	0x12345678(%rbx),%r8
	movswq	0x12345678(%rbx),%r9
	movswq	0x12345678(%rbx),%r10
	movswq	0x12345678(%rbx),%r11
	movswq	0x12345678(%rbx),%r12
	movswq	0x12345678(%rbx),%r13
	movswq	0x12345678(%rbx),%r14
	movswq	0x12345678(%rbx),%r15
	nop
	movswq	0x12345678(%rsp),%rax
	movswq	0x12345678(%rsp),%rcx
	movswq	0x12345678(%rsp),%rdx
	movswq	0x12345678(%rsp),%rbx
	movswq	0x12345678(%rsp),%rsp
	movswq	0x12345678(%rsp),%rbp
	movswq	0x12345678(%rsp),%rsi
	movswq	0x12345678(%rsp),%rdi
	movswq	0x12345678(%rsp),%r8
	movswq	0x12345678(%rsp),%r9
	movswq	0x12345678(%rsp),%r10
	movswq	0x12345678(%rsp),%r11
	movswq	0x12345678(%rsp),%r12
	movswq	0x12345678(%rsp),%r13
	movswq	0x12345678(%rsp),%r14
	movswq	0x12345678(%rsp),%r15
	nop
	movswq	0x12345678(%rbp),%rax
	movswq	0x12345678(%rbp),%rcx
	movswq	0x12345678(%rbp),%rdx
	movswq	0x12345678(%rbp),%rbx
	movswq	0x12345678(%rbp),%rsp
	movswq	0x12345678(%rbp),%rbp
	movswq	0x12345678(%rbp),%rsi
	movswq	0x12345678(%rbp),%rdi
	movswq	0x12345678(%rbp),%r8
	movswq	0x12345678(%rbp),%r9
	movswq	0x12345678(%rbp),%r10
	movswq	0x12345678(%rbp),%r11
	movswq	0x12345678(%rbp),%r12
	movswq	0x12345678(%rbp),%r13
	movswq	0x12345678(%rbp),%r14
	movswq	0x12345678(%rbp),%r15
	nop
	movswq	0x12345678(%rsi),%rax
	movswq	0x12345678(%rsi),%rcx
	movswq	0x12345678(%rsi),%rdx
	movswq	0x12345678(%rsi),%rbx
	movswq	0x12345678(%rsi),%rsp
	movswq	0x12345678(%rsi),%rbp
	movswq	0x12345678(%rsi),%rsi
	movswq	0x12345678(%rsi),%rdi
	movswq	0x12345678(%rsi),%r8
	movswq	0x12345678(%rsi),%r9
	movswq	0x12345678(%rsi),%r10
	movswq	0x12345678(%rsi),%r11
	movswq	0x12345678(%rsi),%r12
	movswq	0x12345678(%rsi),%r13
	movswq	0x12345678(%rsi),%r14
	movswq	0x12345678(%rsi),%r15
	nop
	movswq	0x12345678(%rdi),%rax
	movswq	0x12345678(%rdi),%rcx
	movswq	0x12345678(%rdi),%rdx
	movswq	0x12345678(%rdi),%rbx
	movswq	0x12345678(%rdi),%rsp
	movswq	0x12345678(%rdi),%rbp
	movswq	0x12345678(%rdi),%rsi
	movswq	0x12345678(%rdi),%rdi
	movswq	0x12345678(%rdi),%r8
	movswq	0x12345678(%rdi),%r9
	movswq	0x12345678(%rdi),%r10
	movswq	0x12345678(%rdi),%r11
	movswq	0x12345678(%rdi),%r12
	movswq	0x12345678(%rdi),%r13
	movswq	0x12345678(%rdi),%r14
	movswq	0x12345678(%rdi),%r15
	nop
	movswq	0x12345678(%r8), %rax
	movswq	0x12345678(%r8), %rcx
	movswq	0x12345678(%r8), %rdx
	movswq	0x12345678(%r8), %rbx
	movswq	0x12345678(%r8), %rsp
	movswq	0x12345678(%r8), %rbp
	movswq	0x12345678(%r8), %rsi
	movswq	0x12345678(%r8), %rdi
	movswq	0x12345678(%r8), %r8
	movswq	0x12345678(%r8), %r9
	movswq	0x12345678(%r8), %r10
	movswq	0x12345678(%r8), %r11
	movswq	0x12345678(%r8), %r12
	movswq	0x12345678(%r8), %r13
	movswq	0x12345678(%r8), %r14
	movswq	0x12345678(%r8), %r15
	nop
	movswq	0x12345678(%r9), %rax
	movswq	0x12345678(%r9), %rcx
	movswq	0x12345678(%r9), %rdx
	movswq	0x12345678(%r9), %rbx
	movswq	0x12345678(%r9), %rsp
	movswq	0x12345678(%r9), %rbp
	movswq	0x12345678(%r9), %rsi
	movswq	0x12345678(%r9), %rdi
	movswq	0x12345678(%r9), %r8
	movswq	0x12345678(%r9), %r9
	movswq	0x12345678(%r9), %r10
	movswq	0x12345678(%r9), %r11
	movswq	0x12345678(%r9), %r12
	movswq	0x12345678(%r9), %r13
	movswq	0x12345678(%r9), %r14
	movswq	0x12345678(%r9), %r15
	nop
	movswq	0x12345678(%r10),%rax
	movswq	0x12345678(%r10),%rcx
	movswq	0x12345678(%r10),%rdx
	movswq	0x12345678(%r10),%rbx
	movswq	0x12345678(%r10),%rsp
	movswq	0x12345678(%r10),%rbp
	movswq	0x12345678(%r10),%rsi
	movswq	0x12345678(%r10),%rdi
	movswq	0x12345678(%r10),%r8
	movswq	0x12345678(%r10),%r9
	movswq	0x12345678(%r10),%r10
	movswq	0x12345678(%r10),%r11
	movswq	0x12345678(%r10),%r12
	movswq	0x12345678(%r10),%r13
	movswq	0x12345678(%r10),%r14
	movswq	0x12345678(%r10),%r15
	nop
	movswq	0x12345678(%r11),%rax
	movswq	0x12345678(%r11),%rcx
	movswq	0x12345678(%r11),%rdx
	movswq	0x12345678(%r11),%rbx
	movswq	0x12345678(%r11),%rsp
	movswq	0x12345678(%r11),%rbp
	movswq	0x12345678(%r11),%rsi
	movswq	0x12345678(%r11),%rdi
	movswq	0x12345678(%r11),%r8
	movswq	0x12345678(%r11),%r9
	movswq	0x12345678(%r11),%r10
	movswq	0x12345678(%r11),%r11
	movswq	0x12345678(%r11),%r12
	movswq	0x12345678(%r11),%r13
	movswq	0x12345678(%r11),%r14
	movswq	0x12345678(%r11),%r15
	nop
	movswq	0x12345678(%r12),%rax
	movswq	0x12345678(%r12),%rcx
	movswq	0x12345678(%r12),%rdx
	movswq	0x12345678(%r12),%rbx
	movswq	0x12345678(%r12),%rsp
	movswq	0x12345678(%r12),%rbp
	movswq	0x12345678(%r12),%rsi
	movswq	0x12345678(%r12),%rdi
	movswq	0x12345678(%r12),%r8
	movswq	0x12345678(%r12),%r9
	movswq	0x12345678(%r12),%r10
	movswq	0x12345678(%r12),%r11
	movswq	0x12345678(%r12),%r12
	movswq	0x12345678(%r12),%r13
	movswq	0x12345678(%r12),%r14
	movswq	0x12345678(%r12),%r15
	nop
	movswq	0x12345678(%r13),%rax
	movswq	0x12345678(%r13),%rcx
	movswq	0x12345678(%r13),%rdx
	movswq	0x12345678(%r13),%rbx
	movswq	0x12345678(%r13),%rsp
	movswq	0x12345678(%r13),%rbp
	movswq	0x12345678(%r13),%rsi
	movswq	0x12345678(%r13),%rdi
	movswq	0x12345678(%r13),%r8
	movswq	0x12345678(%r13),%r9
	movswq	0x12345678(%r13),%r10
	movswq	0x12345678(%r13),%r11
	movswq	0x12345678(%r13),%r12
	movswq	0x12345678(%r13),%r13
	movswq	0x12345678(%r13),%r14
	movswq	0x12345678(%r13),%r15
	nop
	movswq	0x12345678(%r14),%rax
	movswq	0x12345678(%r14),%rcx
	movswq	0x12345678(%r14),%rdx
	movswq	0x12345678(%r14),%rbx
	movswq	0x12345678(%r14),%rsp
	movswq	0x12345678(%r14),%rbp
	movswq	0x12345678(%r14),%rsi
	movswq	0x12345678(%r14),%rdi
	movswq	0x12345678(%r14),%r8
	movswq	0x12345678(%r14),%r9
	movswq	0x12345678(%r14),%r10
	movswq	0x12345678(%r14),%r11
	movswq	0x12345678(%r14),%r12
	movswq	0x12345678(%r14),%r13
	movswq	0x12345678(%r14),%r14
	movswq	0x12345678(%r14),%r15
	nop
	movswq	0x12345678(%r15),%rax
	movswq	0x12345678(%r15),%rcx
	movswq	0x12345678(%r15),%rdx
	movswq	0x12345678(%r15),%rbx
	movswq	0x12345678(%r15),%rsp
	movswq	0x12345678(%r15),%rbp
	movswq	0x12345678(%r15),%rsi
	movswq	0x12345678(%r15),%rdi
	movswq	0x12345678(%r15),%r8
	movswq	0x12345678(%r15),%r9
	movswq	0x12345678(%r15),%r10
	movswq	0x12345678(%r15),%r11
	movswq	0x12345678(%r15),%r12
	movswq	0x12345678(%r15),%r13
	movswq	0x12345678(%r15),%r14
	movswq	0x12345678(%r15),%r15
	nop
	nop
	ret
	.cfi_endproc

