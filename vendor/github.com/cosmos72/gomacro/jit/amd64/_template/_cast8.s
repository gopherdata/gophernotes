	.file	"arith.s"
	.text

	.p2align 4,,15
	.globl	Movzbq
	.type	Movzbq, @function
Movzbq:
	.cfi_startproc
	// reg -> reg
	movzbq	%al,%rax
	movzbq	%al,%rcx
	movzbq	%al,%rdx
	movzbq	%al,%rbx
	movzbq	%al,%rsp
	movzbq	%al,%rbp
	movzbq	%al,%rsi
	movzbq	%al,%rdi
	movzbq	%al,%r8
	movzbq	%al,%r9
	movzbq	%al,%r10
	movzbq	%al,%r11
	movzbq	%al,%r12
	movzbq	%al,%r13
	movzbq	%al,%r14
	movzbq	%al,%r15
	nop
	movzbq	%cl,%rax
	movzbq	%cl,%rcx
	movzbq	%cl,%rdx
	movzbq	%cl,%rbx
	movzbq	%cl,%rsp
	movzbq	%cl,%rbp
	movzbq	%cl,%rsi
	movzbq	%cl,%rdi
	movzbq	%cl,%r8
	movzbq	%cl,%r9
	movzbq	%cl,%r10
	movzbq	%cl,%r11
	movzbq	%cl,%r12
	movzbq	%cl,%r13
	movzbq	%cl,%r14
	movzbq	%cl,%r15
	nop
	movzbq	%dl,%rax
	movzbq	%dl,%rcx
	movzbq	%dl,%rdx
	movzbq	%dl,%rbx
	movzbq	%dl,%rsp
	movzbq	%dl,%rbp
	movzbq	%dl,%rsi
	movzbq	%dl,%rdi
	movzbq	%dl,%r8
	movzbq	%dl,%r9
	movzbq	%dl,%r10
	movzbq	%dl,%r11
	movzbq	%dl,%r12
	movzbq	%dl,%r13
	movzbq	%dl,%r14
	movzbq	%dl,%r15
	nop
	movzbq	%bl,%rax
	movzbq	%bl,%rcx
	movzbq	%bl,%rdx
	movzbq	%bl,%rbx
	movzbq	%bl,%rsp
	movzbq	%bl,%rbp
	movzbq	%bl,%rsi
	movzbq	%bl,%rdi
	movzbq	%bl,%r8
	movzbq	%bl,%r9
	movzbq	%bl,%r10
	movzbq	%bl,%r11
	movzbq	%bl,%r12
	movzbq	%bl,%r13
	movzbq	%bl,%r14
	movzbq	%bl,%r15
	nop
	movzbq	%spl,%rax
	movzbq	%spl,%rcx
	movzbq	%spl,%rdx
	movzbq	%spl,%rbx
	movzbq	%spl,%rsp
	movzbq	%spl,%rbp
	movzbq	%spl,%rsi
	movzbq	%spl,%rdi
	movzbq	%spl,%r8
	movzbq	%spl,%r9
	movzbq	%spl,%r10
	movzbq	%spl,%r11
	movzbq	%spl,%r12
	movzbq	%spl,%r13
	movzbq	%spl,%r14
	movzbq	%spl,%r15
	nop
	movzbq	%bpl,%rax
	movzbq	%bpl,%rcx
	movzbq	%bpl,%rdx
	movzbq	%bpl,%rbx
	movzbq	%bpl,%rsp
	movzbq	%bpl,%rbp
	movzbq	%bpl,%rsi
	movzbq	%bpl,%rdi
	movzbq	%bpl,%r8
	movzbq	%bpl,%r9
	movzbq	%bpl,%r10
	movzbq	%bpl,%r11
	movzbq	%bpl,%r12
	movzbq	%bpl,%r13
	movzbq	%bpl,%r14
	movzbq	%bpl,%r15
	nop
	movzbq	%sil,%rax
	movzbq	%sil,%rcx
	movzbq	%sil,%rdx
	movzbq	%sil,%rbx
	movzbq	%sil,%rsp
	movzbq	%sil,%rbp
	movzbq	%sil,%rsi
	movzbq	%sil,%rdi
	movzbq	%sil,%r8
	movzbq	%sil,%r9
	movzbq	%sil,%r10
	movzbq	%sil,%r11
	movzbq	%sil,%r12
	movzbq	%sil,%r13
	movzbq	%sil,%r14
	movzbq	%sil,%r15
	nop
	movzbq	%dil,%rax
	movzbq	%dil,%rcx
	movzbq	%dil,%rdx
	movzbq	%dil,%rbx
	movzbq	%dil,%rsp
	movzbq	%dil,%rbp
	movzbq	%dil,%rsi
	movzbq	%dil,%rdi
	movzbq	%dil,%r8
	movzbq	%dil,%r9
	movzbq	%dil,%r10
	movzbq	%dil,%r11
	movzbq	%dil,%r12
	movzbq	%dil,%r13
	movzbq	%dil,%r14
	movzbq	%dil,%r15
	nop
	movzbq	%r8b, %rax
	movzbq	%r8b, %rcx
	movzbq	%r8b, %rdx
	movzbq	%r8b, %rbx
	movzbq	%r8b, %rsp
	movzbq	%r8b, %rbp
	movzbq	%r8b, %rsi
	movzbq	%r8b, %rdi
	movzbq	%r8b, %r8
	movzbq	%r8b, %r9
	movzbq	%r8b, %r10
	movzbq	%r8b, %r11
	movzbq	%r8b, %r12
	movzbq	%r8b, %r13
	movzbq	%r8b, %r14
	movzbq	%r8b, %r15
	nop
	movzbq	%r9b, %rax
	movzbq	%r9b, %rcx
	movzbq	%r9b, %rdx
	movzbq	%r9b, %rbx
	movzbq	%r9b, %rsp
	movzbq	%r9b, %rbp
	movzbq	%r9b, %rsi
	movzbq	%r9b, %rdi
	movzbq	%r9b, %r8
	movzbq	%r9b, %r9
	movzbq	%r9b, %r10
	movzbq	%r9b, %r11
	movzbq	%r9b, %r12
	movzbq	%r9b, %r13
	movzbq	%r9b, %r14
	movzbq	%r9b, %r15
	nop
	movzbq	%r10b,%rax
	movzbq	%r10b,%rcx
	movzbq	%r10b,%rdx
	movzbq	%r10b,%rbx
	movzbq	%r10b,%rsp
	movzbq	%r10b,%rbp
	movzbq	%r10b,%rsi
	movzbq	%r10b,%rdi
	movzbq	%r10b,%r8
	movzbq	%r10b,%r9
	movzbq	%r10b,%r10
	movzbq	%r10b,%r11
	movzbq	%r10b,%r12
	movzbq	%r10b,%r13
	movzbq	%r10b,%r14
	movzbq	%r10b,%r15
	nop
	movzbq	%r11b,%rax
	movzbq	%r11b,%rcx
	movzbq	%r11b,%rdx
	movzbq	%r11b,%rbx
	movzbq	%r11b,%rsp
	movzbq	%r11b,%rbp
	movzbq	%r11b,%rsi
	movzbq	%r11b,%rdi
	movzbq	%r11b,%r8
	movzbq	%r11b,%r9
	movzbq	%r11b,%r10
	movzbq	%r11b,%r11
	movzbq	%r11b,%r12
	movzbq	%r11b,%r13
	movzbq	%r11b,%r14
	movzbq	%r11b,%r15
	nop
	movzbq	%r12b,%rax
	movzbq	%r12b,%rcx
	movzbq	%r12b,%rdx
	movzbq	%r12b,%rbx
	movzbq	%r12b,%rsp
	movzbq	%r12b,%rbp
	movzbq	%r12b,%rsi
	movzbq	%r12b,%rdi
	movzbq	%r12b,%r8
	movzbq	%r12b,%r9
	movzbq	%r12b,%r10
	movzbq	%r12b,%r11
	movzbq	%r12b,%r12
	movzbq	%r12b,%r13
	movzbq	%r12b,%r14
	movzbq	%r12b,%r15
	nop
	movzbq	%r13b,%rax
	movzbq	%r13b,%rcx
	movzbq	%r13b,%rdx
	movzbq	%r13b,%rbx
	movzbq	%r13b,%rsp
	movzbq	%r13b,%rbp
	movzbq	%r13b,%rsi
	movzbq	%r13b,%rdi
	movzbq	%r13b,%r8
	movzbq	%r13b,%r9
	movzbq	%r13b,%r10
	movzbq	%r13b,%r11
	movzbq	%r13b,%r12
	movzbq	%r13b,%r13
	movzbq	%r13b,%r14
	movzbq	%r13b,%r15
	nop
	movzbq	%r14b,%rax
	movzbq	%r14b,%rcx
	movzbq	%r14b,%rdx
	movzbq	%r14b,%rbx
	movzbq	%r14b,%rsp
	movzbq	%r14b,%rbp
	movzbq	%r14b,%rsi
	movzbq	%r14b,%rdi
	movzbq	%r14b,%r8
	movzbq	%r14b,%r9
	movzbq	%r14b,%r10
	movzbq	%r14b,%r11
	movzbq	%r14b,%r12
	movzbq	%r14b,%r13
	movzbq	%r14b,%r14
	movzbq	%r14b,%r15
	nop
	movzbq	%r15b,%rax
	movzbq	%r15b,%rcx
	movzbq	%r15b,%rdx
	movzbq	%r15b,%rbx
	movzbq	%r15b,%rsp
	movzbq	%r15b,%rbp
	movzbq	%r15b,%rsi
	movzbq	%r15b,%rdi
	movzbq	%r15b,%r8
	movzbq	%r15b,%r9
	movzbq	%r15b,%r10
	movzbq	%r15b,%r11
	movzbq	%r15b,%r12
	movzbq	%r15b,%r13
	movzbq	%r15b,%r14
	movzbq	%r15b,%r15
	nop
        nop
	// (reg) -> reg
	movzbq	(%rax),%rax
	movzbq	(%rax),%rcx
	movzbq	(%rax),%rdx
	movzbq	(%rax),%rbx
	movzbq	(%rax),%rsp
	movzbq	(%rax),%rbp
	movzbq	(%rax),%rsi
	movzbq	(%rax),%rdi
	movzbq	(%rax),%r8
	movzbq	(%rax),%r9
	movzbq	(%rax),%r10
	movzbq	(%rax),%r11
	movzbq	(%rax),%r12
	movzbq	(%rax),%r13
	movzbq	(%rax),%r14
	movzbq	(%rax),%r15
	nop
	movzbq	(%rcx),%rax
	movzbq	(%rcx),%rcx
	movzbq	(%rcx),%rdx
	movzbq	(%rcx),%rbx
	movzbq	(%rcx),%rsp
	movzbq	(%rcx),%rbp
	movzbq	(%rcx),%rsi
	movzbq	(%rcx),%rdi
	movzbq	(%rcx),%r8
	movzbq	(%rcx),%r9
	movzbq	(%rcx),%r10
	movzbq	(%rcx),%r11
	movzbq	(%rcx),%r12
	movzbq	(%rcx),%r13
	movzbq	(%rcx),%r14
	movzbq	(%rcx),%r15
	nop
	movzbq	(%rdx),%rax
	movzbq	(%rdx),%rcx
	movzbq	(%rdx),%rdx
	movzbq	(%rdx),%rbx
	movzbq	(%rdx),%rsp
	movzbq	(%rdx),%rbp
	movzbq	(%rdx),%rsi
	movzbq	(%rdx),%rdi
	movzbq	(%rdx),%r8
	movzbq	(%rdx),%r9
	movzbq	(%rdx),%r10
	movzbq	(%rdx),%r11
	movzbq	(%rdx),%r12
	movzbq	(%rdx),%r13
	movzbq	(%rdx),%r14
	movzbq	(%rdx),%r15
	nop
	movzbq	(%rbx),%rax
	movzbq	(%rbx),%rcx
	movzbq	(%rbx),%rdx
	movzbq	(%rbx),%rbx
	movzbq	(%rbx),%rsp
	movzbq	(%rbx),%rbp
	movzbq	(%rbx),%rsi
	movzbq	(%rbx),%rdi
	movzbq	(%rbx),%r8
	movzbq	(%rbx),%r9
	movzbq	(%rbx),%r10
	movzbq	(%rbx),%r11
	movzbq	(%rbx),%r12
	movzbq	(%rbx),%r13
	movzbq	(%rbx),%r14
	movzbq	(%rbx),%r15
	nop
	movzbq	(%rsp),%rax
	movzbq	(%rsp),%rcx
	movzbq	(%rsp),%rdx
	movzbq	(%rsp),%rbx
	movzbq	(%rsp),%rsp
	movzbq	(%rsp),%rbp
	movzbq	(%rsp),%rsi
	movzbq	(%rsp),%rdi
	movzbq	(%rsp),%r8
	movzbq	(%rsp),%r9
	movzbq	(%rsp),%r10
	movzbq	(%rsp),%r11
	movzbq	(%rsp),%r12
	movzbq	(%rsp),%r13
	movzbq	(%rsp),%r14
	movzbq	(%rsp),%r15
	nop
	movzbq	(%rbp),%rax
	movzbq	(%rbp),%rcx
	movzbq	(%rbp),%rdx
	movzbq	(%rbp),%rbx
	movzbq	(%rbp),%rsp
	movzbq	(%rbp),%rbp
	movzbq	(%rbp),%rsi
	movzbq	(%rbp),%rdi
	movzbq	(%rbp),%r8
	movzbq	(%rbp),%r9
	movzbq	(%rbp),%r10
	movzbq	(%rbp),%r11
	movzbq	(%rbp),%r12
	movzbq	(%rbp),%r13
	movzbq	(%rbp),%r14
	movzbq	(%rbp),%r15
	nop
	movzbq	(%rsi),%rax
	movzbq	(%rsi),%rcx
	movzbq	(%rsi),%rdx
	movzbq	(%rsi),%rbx
	movzbq	(%rsi),%rsp
	movzbq	(%rsi),%rbp
	movzbq	(%rsi),%rsi
	movzbq	(%rsi),%rdi
	movzbq	(%rsi),%r8
	movzbq	(%rsi),%r9
	movzbq	(%rsi),%r10
	movzbq	(%rsi),%r11
	movzbq	(%rsi),%r12
	movzbq	(%rsi),%r13
	movzbq	(%rsi),%r14
	movzbq	(%rsi),%r15
	nop
	movzbq	(%rdi),%rax
	movzbq	(%rdi),%rcx
	movzbq	(%rdi),%rdx
	movzbq	(%rdi),%rbx
	movzbq	(%rdi),%rsp
	movzbq	(%rdi),%rbp
	movzbq	(%rdi),%rsi
	movzbq	(%rdi),%rdi
	movzbq	(%rdi),%r8
	movzbq	(%rdi),%r9
	movzbq	(%rdi),%r10
	movzbq	(%rdi),%r11
	movzbq	(%rdi),%r12
	movzbq	(%rdi),%r13
	movzbq	(%rdi),%r14
	movzbq	(%rdi),%r15
	nop
	movzbq	(%r8), %rax
	movzbq	(%r8), %rcx
	movzbq	(%r8), %rdx
	movzbq	(%r8), %rbx
	movzbq	(%r8), %rsp
	movzbq	(%r8), %rbp
	movzbq	(%r8), %rsi
	movzbq	(%r8), %rdi
	movzbq	(%r8), %r8
	movzbq	(%r8), %r9
	movzbq	(%r8), %r10
	movzbq	(%r8), %r11
	movzbq	(%r8), %r12
	movzbq	(%r8), %r13
	movzbq	(%r8), %r14
	movzbq	(%r8), %r15
	nop
	movzbq	(%r9), %rax
	movzbq	(%r9), %rcx
	movzbq	(%r9), %rdx
	movzbq	(%r9), %rbx
	movzbq	(%r9), %rsp
	movzbq	(%r9), %rbp
	movzbq	(%r9), %rsi
	movzbq	(%r9), %rdi
	movzbq	(%r9), %r8
	movzbq	(%r9), %r9
	movzbq	(%r9), %r10
	movzbq	(%r9), %r11
	movzbq	(%r9), %r12
	movzbq	(%r9), %r13
	movzbq	(%r9), %r14
	movzbq	(%r9), %r15
	nop
	movzbq	(%r10),%rax
	movzbq	(%r10),%rcx
	movzbq	(%r10),%rdx
	movzbq	(%r10),%rbx
	movzbq	(%r10),%rsp
	movzbq	(%r10),%rbp
	movzbq	(%r10),%rsi
	movzbq	(%r10),%rdi
	movzbq	(%r10),%r8
	movzbq	(%r10),%r9
	movzbq	(%r10),%r10
	movzbq	(%r10),%r11
	movzbq	(%r10),%r12
	movzbq	(%r10),%r13
	movzbq	(%r10),%r14
	movzbq	(%r10),%r15
	nop
	movzbq	(%r11),%rax
	movzbq	(%r11),%rcx
	movzbq	(%r11),%rdx
	movzbq	(%r11),%rbx
	movzbq	(%r11),%rsp
	movzbq	(%r11),%rbp
	movzbq	(%r11),%rsi
	movzbq	(%r11),%rdi
	movzbq	(%r11),%r8
	movzbq	(%r11),%r9
	movzbq	(%r11),%r10
	movzbq	(%r11),%r11
	movzbq	(%r11),%r12
	movzbq	(%r11),%r13
	movzbq	(%r11),%r14
	movzbq	(%r11),%r15
	nop
	movzbq	(%r12),%rax
	movzbq	(%r12),%rcx
	movzbq	(%r12),%rdx
	movzbq	(%r12),%rbx
	movzbq	(%r12),%rsp
	movzbq	(%r12),%rbp
	movzbq	(%r12),%rsi
	movzbq	(%r12),%rdi
	movzbq	(%r12),%r8
	movzbq	(%r12),%r9
	movzbq	(%r12),%r10
	movzbq	(%r12),%r11
	movzbq	(%r12),%r12
	movzbq	(%r12),%r13
	movzbq	(%r12),%r14
	movzbq	(%r12),%r15
	nop
	movzbq	(%r13),%rax
	movzbq	(%r13),%rcx
	movzbq	(%r13),%rdx
	movzbq	(%r13),%rbx
	movzbq	(%r13),%rsp
	movzbq	(%r13),%rbp
	movzbq	(%r13),%rsi
	movzbq	(%r13),%rdi
	movzbq	(%r13),%r8
	movzbq	(%r13),%r9
	movzbq	(%r13),%r10
	movzbq	(%r13),%r11
	movzbq	(%r13),%r12
	movzbq	(%r13),%r13
	movzbq	(%r13),%r14
	movzbq	(%r13),%r15
	nop
	movzbq	(%r14),%rax
	movzbq	(%r14),%rcx
	movzbq	(%r14),%rdx
	movzbq	(%r14),%rbx
	movzbq	(%r14),%rsp
	movzbq	(%r14),%rbp
	movzbq	(%r14),%rsi
	movzbq	(%r14),%rdi
	movzbq	(%r14),%r8
	movzbq	(%r14),%r9
	movzbq	(%r14),%r10
	movzbq	(%r14),%r11
	movzbq	(%r14),%r12
	movzbq	(%r14),%r13
	movzbq	(%r14),%r14
	movzbq	(%r14),%r15
	nop
	movzbq	(%r15),%rax
	movzbq	(%r15),%rcx
	movzbq	(%r15),%rdx
	movzbq	(%r15),%rbx
	movzbq	(%r15),%rsp
	movzbq	(%r15),%rbp
	movzbq	(%r15),%rsi
	movzbq	(%r15),%rdi
	movzbq	(%r15),%r8
	movzbq	(%r15),%r9
	movzbq	(%r15),%r10
	movzbq	(%r15),%r11
	movzbq	(%r15),%r12
	movzbq	(%r15),%r13
	movzbq	(%r15),%r14
	movzbq	(%r15),%r15
	nop
	nop
	// off8(reg) -> reg
	movzbq	0x7F(%rax),%rax
	movzbq	0x7F(%rax),%rcx
	movzbq	0x7F(%rax),%rdx
	movzbq	0x7F(%rax),%rbx
	movzbq	0x7F(%rax),%rsp
	movzbq	0x7F(%rax),%rbp
	movzbq	0x7F(%rax),%rsi
	movzbq	0x7F(%rax),%rdi
	movzbq	0x7F(%rax),%r8
	movzbq	0x7F(%rax),%r9
	movzbq	0x7F(%rax),%r10
	movzbq	0x7F(%rax),%r11
	movzbq	0x7F(%rax),%r12
	movzbq	0x7F(%rax),%r13
	movzbq	0x7F(%rax),%r14
	movzbq	0x7F(%rax),%r15
	nop
	movzbq	0x7F(%rcx),%rax
	movzbq	0x7F(%rcx),%rcx
	movzbq	0x7F(%rcx),%rdx
	movzbq	0x7F(%rcx),%rbx
	movzbq	0x7F(%rcx),%rsp
	movzbq	0x7F(%rcx),%rbp
	movzbq	0x7F(%rcx),%rsi
	movzbq	0x7F(%rcx),%rdi
	movzbq	0x7F(%rcx),%r8
	movzbq	0x7F(%rcx),%r9
	movzbq	0x7F(%rcx),%r10
	movzbq	0x7F(%rcx),%r11
	movzbq	0x7F(%rcx),%r12
	movzbq	0x7F(%rcx),%r13
	movzbq	0x7F(%rcx),%r14
	movzbq	0x7F(%rcx),%r15
	nop
	movzbq	0x7F(%rdx),%rax
	movzbq	0x7F(%rdx),%rcx
	movzbq	0x7F(%rdx),%rdx
	movzbq	0x7F(%rdx),%rbx
	movzbq	0x7F(%rdx),%rsp
	movzbq	0x7F(%rdx),%rbp
	movzbq	0x7F(%rdx),%rsi
	movzbq	0x7F(%rdx),%rdi
	movzbq	0x7F(%rdx),%r8
	movzbq	0x7F(%rdx),%r9
	movzbq	0x7F(%rdx),%r10
	movzbq	0x7F(%rdx),%r11
	movzbq	0x7F(%rdx),%r12
	movzbq	0x7F(%rdx),%r13
	movzbq	0x7F(%rdx),%r14
	movzbq	0x7F(%rdx),%r15
	nop
	movzbq	0x7F(%rbx),%rax
	movzbq	0x7F(%rbx),%rcx
	movzbq	0x7F(%rbx),%rdx
	movzbq	0x7F(%rbx),%rbx
	movzbq	0x7F(%rbx),%rsp
	movzbq	0x7F(%rbx),%rbp
	movzbq	0x7F(%rbx),%rsi
	movzbq	0x7F(%rbx),%rdi
	movzbq	0x7F(%rbx),%r8
	movzbq	0x7F(%rbx),%r9
	movzbq	0x7F(%rbx),%r10
	movzbq	0x7F(%rbx),%r11
	movzbq	0x7F(%rbx),%r12
	movzbq	0x7F(%rbx),%r13
	movzbq	0x7F(%rbx),%r14
	movzbq	0x7F(%rbx),%r15
	nop
	movzbq	0x7F(%rsp),%rax
	movzbq	0x7F(%rsp),%rcx
	movzbq	0x7F(%rsp),%rdx
	movzbq	0x7F(%rsp),%rbx
	movzbq	0x7F(%rsp),%rsp
	movzbq	0x7F(%rsp),%rbp
	movzbq	0x7F(%rsp),%rsi
	movzbq	0x7F(%rsp),%rdi
	movzbq	0x7F(%rsp),%r8
	movzbq	0x7F(%rsp),%r9
	movzbq	0x7F(%rsp),%r10
	movzbq	0x7F(%rsp),%r11
	movzbq	0x7F(%rsp),%r12
	movzbq	0x7F(%rsp),%r13
	movzbq	0x7F(%rsp),%r14
	movzbq	0x7F(%rsp),%r15
	nop
	movzbq	0x7F(%rbp),%rax
	movzbq	0x7F(%rbp),%rcx
	movzbq	0x7F(%rbp),%rdx
	movzbq	0x7F(%rbp),%rbx
	movzbq	0x7F(%rbp),%rsp
	movzbq	0x7F(%rbp),%rbp
	movzbq	0x7F(%rbp),%rsi
	movzbq	0x7F(%rbp),%rdi
	movzbq	0x7F(%rbp),%r8
	movzbq	0x7F(%rbp),%r9
	movzbq	0x7F(%rbp),%r10
	movzbq	0x7F(%rbp),%r11
	movzbq	0x7F(%rbp),%r12
	movzbq	0x7F(%rbp),%r13
	movzbq	0x7F(%rbp),%r14
	movzbq	0x7F(%rbp),%r15
	nop
	movzbq	0x7F(%rsi),%rax
	movzbq	0x7F(%rsi),%rcx
	movzbq	0x7F(%rsi),%rdx
	movzbq	0x7F(%rsi),%rbx
	movzbq	0x7F(%rsi),%rsp
	movzbq	0x7F(%rsi),%rbp
	movzbq	0x7F(%rsi),%rsi
	movzbq	0x7F(%rsi),%rdi
	movzbq	0x7F(%rsi),%r8
	movzbq	0x7F(%rsi),%r9
	movzbq	0x7F(%rsi),%r10
	movzbq	0x7F(%rsi),%r11
	movzbq	0x7F(%rsi),%r12
	movzbq	0x7F(%rsi),%r13
	movzbq	0x7F(%rsi),%r14
	movzbq	0x7F(%rsi),%r15
	nop
	movzbq	0x7F(%rdi),%rax
	movzbq	0x7F(%rdi),%rcx
	movzbq	0x7F(%rdi),%rdx
	movzbq	0x7F(%rdi),%rbx
	movzbq	0x7F(%rdi),%rsp
	movzbq	0x7F(%rdi),%rbp
	movzbq	0x7F(%rdi),%rsi
	movzbq	0x7F(%rdi),%rdi
	movzbq	0x7F(%rdi),%r8
	movzbq	0x7F(%rdi),%r9
	movzbq	0x7F(%rdi),%r10
	movzbq	0x7F(%rdi),%r11
	movzbq	0x7F(%rdi),%r12
	movzbq	0x7F(%rdi),%r13
	movzbq	0x7F(%rdi),%r14
	movzbq	0x7F(%rdi),%r15
	nop
	movzbq	0x7F(%r8), %rax
	movzbq	0x7F(%r8), %rcx
	movzbq	0x7F(%r8), %rdx
	movzbq	0x7F(%r8), %rbx
	movzbq	0x7F(%r8), %rsp
	movzbq	0x7F(%r8), %rbp
	movzbq	0x7F(%r8), %rsi
	movzbq	0x7F(%r8), %rdi
	movzbq	0x7F(%r8), %r8
	movzbq	0x7F(%r8), %r9
	movzbq	0x7F(%r8), %r10
	movzbq	0x7F(%r8), %r11
	movzbq	0x7F(%r8), %r12
	movzbq	0x7F(%r8), %r13
	movzbq	0x7F(%r8), %r14
	movzbq	0x7F(%r8), %r15
	nop
	movzbq	0x7F(%r9), %rax
	movzbq	0x7F(%r9), %rcx
	movzbq	0x7F(%r9), %rdx
	movzbq	0x7F(%r9), %rbx
	movzbq	0x7F(%r9), %rsp
	movzbq	0x7F(%r9), %rbp
	movzbq	0x7F(%r9), %rsi
	movzbq	0x7F(%r9), %rdi
	movzbq	0x7F(%r9), %r8
	movzbq	0x7F(%r9), %r9
	movzbq	0x7F(%r9), %r10
	movzbq	0x7F(%r9), %r11
	movzbq	0x7F(%r9), %r12
	movzbq	0x7F(%r9), %r13
	movzbq	0x7F(%r9), %r14
	movzbq	0x7F(%r9), %r15
	nop
	movzbq	0x7F(%r10),%rax
	movzbq	0x7F(%r10),%rcx
	movzbq	0x7F(%r10),%rdx
	movzbq	0x7F(%r10),%rbx
	movzbq	0x7F(%r10),%rsp
	movzbq	0x7F(%r10),%rbp
	movzbq	0x7F(%r10),%rsi
	movzbq	0x7F(%r10),%rdi
	movzbq	0x7F(%r10),%r8
	movzbq	0x7F(%r10),%r9
	movzbq	0x7F(%r10),%r10
	movzbq	0x7F(%r10),%r11
	movzbq	0x7F(%r10),%r12
	movzbq	0x7F(%r10),%r13
	movzbq	0x7F(%r10),%r14
	movzbq	0x7F(%r10),%r15
	nop
	movzbq	0x7F(%r11),%rax
	movzbq	0x7F(%r11),%rcx
	movzbq	0x7F(%r11),%rdx
	movzbq	0x7F(%r11),%rbx
	movzbq	0x7F(%r11),%rsp
	movzbq	0x7F(%r11),%rbp
	movzbq	0x7F(%r11),%rsi
	movzbq	0x7F(%r11),%rdi
	movzbq	0x7F(%r11),%r8
	movzbq	0x7F(%r11),%r9
	movzbq	0x7F(%r11),%r10
	movzbq	0x7F(%r11),%r11
	movzbq	0x7F(%r11),%r12
	movzbq	0x7F(%r11),%r13
	movzbq	0x7F(%r11),%r14
	movzbq	0x7F(%r11),%r15
	nop
	movzbq	0x7F(%r12),%rax
	movzbq	0x7F(%r12),%rcx
	movzbq	0x7F(%r12),%rdx
	movzbq	0x7F(%r12),%rbx
	movzbq	0x7F(%r12),%rsp
	movzbq	0x7F(%r12),%rbp
	movzbq	0x7F(%r12),%rsi
	movzbq	0x7F(%r12),%rdi
	movzbq	0x7F(%r12),%r8
	movzbq	0x7F(%r12),%r9
	movzbq	0x7F(%r12),%r10
	movzbq	0x7F(%r12),%r11
	movzbq	0x7F(%r12),%r12
	movzbq	0x7F(%r12),%r13
	movzbq	0x7F(%r12),%r14
	movzbq	0x7F(%r12),%r15
	nop
	movzbq	0x7F(%r13),%rax
	movzbq	0x7F(%r13),%rcx
	movzbq	0x7F(%r13),%rdx
	movzbq	0x7F(%r13),%rbx
	movzbq	0x7F(%r13),%rsp
	movzbq	0x7F(%r13),%rbp
	movzbq	0x7F(%r13),%rsi
	movzbq	0x7F(%r13),%rdi
	movzbq	0x7F(%r13),%r8
	movzbq	0x7F(%r13),%r9
	movzbq	0x7F(%r13),%r10
	movzbq	0x7F(%r13),%r11
	movzbq	0x7F(%r13),%r12
	movzbq	0x7F(%r13),%r13
	movzbq	0x7F(%r13),%r14
	movzbq	0x7F(%r13),%r15
	nop
	movzbq	0x7F(%r14),%rax
	movzbq	0x7F(%r14),%rcx
	movzbq	0x7F(%r14),%rdx
	movzbq	0x7F(%r14),%rbx
	movzbq	0x7F(%r14),%rsp
	movzbq	0x7F(%r14),%rbp
	movzbq	0x7F(%r14),%rsi
	movzbq	0x7F(%r14),%rdi
	movzbq	0x7F(%r14),%r8
	movzbq	0x7F(%r14),%r9
	movzbq	0x7F(%r14),%r10
	movzbq	0x7F(%r14),%r11
	movzbq	0x7F(%r14),%r12
	movzbq	0x7F(%r14),%r13
	movzbq	0x7F(%r14),%r14
	movzbq	0x7F(%r14),%r15
	nop
	movzbq	0x7F(%r15),%rax
	movzbq	0x7F(%r15),%rcx
	movzbq	0x7F(%r15),%rdx
	movzbq	0x7F(%r15),%rbx
	movzbq	0x7F(%r15),%rsp
	movzbq	0x7F(%r15),%rbp
	movzbq	0x7F(%r15),%rsi
	movzbq	0x7F(%r15),%rdi
	movzbq	0x7F(%r15),%r8
	movzbq	0x7F(%r15),%r9
	movzbq	0x7F(%r15),%r10
	movzbq	0x7F(%r15),%r11
	movzbq	0x7F(%r15),%r12
	movzbq	0x7F(%r15),%r13
	movzbq	0x7F(%r15),%r14
	movzbq	0x7F(%r15),%r15
	nop
	nop
	// off32(reg) -> reg
	movzbq	0x12345678(%rax),%rax
	movzbq	0x12345678(%rax),%rcx
	movzbq	0x12345678(%rax),%rdx
	movzbq	0x12345678(%rax),%rbx
	movzbq	0x12345678(%rax),%rsp
	movzbq	0x12345678(%rax),%rbp
	movzbq	0x12345678(%rax),%rsi
	movzbq	0x12345678(%rax),%rdi
	movzbq	0x12345678(%rax),%r8
	movzbq	0x12345678(%rax),%r9
	movzbq	0x12345678(%rax),%r10
	movzbq	0x12345678(%rax),%r11
	movzbq	0x12345678(%rax),%r12
	movzbq	0x12345678(%rax),%r13
	movzbq	0x12345678(%rax),%r14
	movzbq	0x12345678(%rax),%r15
	nop
	movzbq	0x12345678(%rcx),%rax
	movzbq	0x12345678(%rcx),%rcx
	movzbq	0x12345678(%rcx),%rdx
	movzbq	0x12345678(%rcx),%rbx
	movzbq	0x12345678(%rcx),%rsp
	movzbq	0x12345678(%rcx),%rbp
	movzbq	0x12345678(%rcx),%rsi
	movzbq	0x12345678(%rcx),%rdi
	movzbq	0x12345678(%rcx),%r8
	movzbq	0x12345678(%rcx),%r9
	movzbq	0x12345678(%rcx),%r10
	movzbq	0x12345678(%rcx),%r11
	movzbq	0x12345678(%rcx),%r12
	movzbq	0x12345678(%rcx),%r13
	movzbq	0x12345678(%rcx),%r14
	movzbq	0x12345678(%rcx),%r15
	nop
	movzbq	0x12345678(%rdx),%rax
	movzbq	0x12345678(%rdx),%rcx
	movzbq	0x12345678(%rdx),%rdx
	movzbq	0x12345678(%rdx),%rbx
	movzbq	0x12345678(%rdx),%rsp
	movzbq	0x12345678(%rdx),%rbp
	movzbq	0x12345678(%rdx),%rsi
	movzbq	0x12345678(%rdx),%rdi
	movzbq	0x12345678(%rdx),%r8
	movzbq	0x12345678(%rdx),%r9
	movzbq	0x12345678(%rdx),%r10
	movzbq	0x12345678(%rdx),%r11
	movzbq	0x12345678(%rdx),%r12
	movzbq	0x12345678(%rdx),%r13
	movzbq	0x12345678(%rdx),%r14
	movzbq	0x12345678(%rdx),%r15
	nop
	movzbq	0x12345678(%rbx),%rax
	movzbq	0x12345678(%rbx),%rcx
	movzbq	0x12345678(%rbx),%rdx
	movzbq	0x12345678(%rbx),%rbx
	movzbq	0x12345678(%rbx),%rsp
	movzbq	0x12345678(%rbx),%rbp
	movzbq	0x12345678(%rbx),%rsi
	movzbq	0x12345678(%rbx),%rdi
	movzbq	0x12345678(%rbx),%r8
	movzbq	0x12345678(%rbx),%r9
	movzbq	0x12345678(%rbx),%r10
	movzbq	0x12345678(%rbx),%r11
	movzbq	0x12345678(%rbx),%r12
	movzbq	0x12345678(%rbx),%r13
	movzbq	0x12345678(%rbx),%r14
	movzbq	0x12345678(%rbx),%r15
	nop
	movzbq	0x12345678(%rsp),%rax
	movzbq	0x12345678(%rsp),%rcx
	movzbq	0x12345678(%rsp),%rdx
	movzbq	0x12345678(%rsp),%rbx
	movzbq	0x12345678(%rsp),%rsp
	movzbq	0x12345678(%rsp),%rbp
	movzbq	0x12345678(%rsp),%rsi
	movzbq	0x12345678(%rsp),%rdi
	movzbq	0x12345678(%rsp),%r8
	movzbq	0x12345678(%rsp),%r9
	movzbq	0x12345678(%rsp),%r10
	movzbq	0x12345678(%rsp),%r11
	movzbq	0x12345678(%rsp),%r12
	movzbq	0x12345678(%rsp),%r13
	movzbq	0x12345678(%rsp),%r14
	movzbq	0x12345678(%rsp),%r15
	nop
	movzbq	0x12345678(%rbp),%rax
	movzbq	0x12345678(%rbp),%rcx
	movzbq	0x12345678(%rbp),%rdx
	movzbq	0x12345678(%rbp),%rbx
	movzbq	0x12345678(%rbp),%rsp
	movzbq	0x12345678(%rbp),%rbp
	movzbq	0x12345678(%rbp),%rsi
	movzbq	0x12345678(%rbp),%rdi
	movzbq	0x12345678(%rbp),%r8
	movzbq	0x12345678(%rbp),%r9
	movzbq	0x12345678(%rbp),%r10
	movzbq	0x12345678(%rbp),%r11
	movzbq	0x12345678(%rbp),%r12
	movzbq	0x12345678(%rbp),%r13
	movzbq	0x12345678(%rbp),%r14
	movzbq	0x12345678(%rbp),%r15
	nop
	movzbq	0x12345678(%rsi),%rax
	movzbq	0x12345678(%rsi),%rcx
	movzbq	0x12345678(%rsi),%rdx
	movzbq	0x12345678(%rsi),%rbx
	movzbq	0x12345678(%rsi),%rsp
	movzbq	0x12345678(%rsi),%rbp
	movzbq	0x12345678(%rsi),%rsi
	movzbq	0x12345678(%rsi),%rdi
	movzbq	0x12345678(%rsi),%r8
	movzbq	0x12345678(%rsi),%r9
	movzbq	0x12345678(%rsi),%r10
	movzbq	0x12345678(%rsi),%r11
	movzbq	0x12345678(%rsi),%r12
	movzbq	0x12345678(%rsi),%r13
	movzbq	0x12345678(%rsi),%r14
	movzbq	0x12345678(%rsi),%r15
	nop
	movzbq	0x12345678(%rdi),%rax
	movzbq	0x12345678(%rdi),%rcx
	movzbq	0x12345678(%rdi),%rdx
	movzbq	0x12345678(%rdi),%rbx
	movzbq	0x12345678(%rdi),%rsp
	movzbq	0x12345678(%rdi),%rbp
	movzbq	0x12345678(%rdi),%rsi
	movzbq	0x12345678(%rdi),%rdi
	movzbq	0x12345678(%rdi),%r8
	movzbq	0x12345678(%rdi),%r9
	movzbq	0x12345678(%rdi),%r10
	movzbq	0x12345678(%rdi),%r11
	movzbq	0x12345678(%rdi),%r12
	movzbq	0x12345678(%rdi),%r13
	movzbq	0x12345678(%rdi),%r14
	movzbq	0x12345678(%rdi),%r15
	nop
	movzbq	0x12345678(%r8), %rax
	movzbq	0x12345678(%r8), %rcx
	movzbq	0x12345678(%r8), %rdx
	movzbq	0x12345678(%r8), %rbx
	movzbq	0x12345678(%r8), %rsp
	movzbq	0x12345678(%r8), %rbp
	movzbq	0x12345678(%r8), %rsi
	movzbq	0x12345678(%r8), %rdi
	movzbq	0x12345678(%r8), %r8
	movzbq	0x12345678(%r8), %r9
	movzbq	0x12345678(%r8), %r10
	movzbq	0x12345678(%r8), %r11
	movzbq	0x12345678(%r8), %r12
	movzbq	0x12345678(%r8), %r13
	movzbq	0x12345678(%r8), %r14
	movzbq	0x12345678(%r8), %r15
	nop
	movzbq	0x12345678(%r9), %rax
	movzbq	0x12345678(%r9), %rcx
	movzbq	0x12345678(%r9), %rdx
	movzbq	0x12345678(%r9), %rbx
	movzbq	0x12345678(%r9), %rsp
	movzbq	0x12345678(%r9), %rbp
	movzbq	0x12345678(%r9), %rsi
	movzbq	0x12345678(%r9), %rdi
	movzbq	0x12345678(%r9), %r8
	movzbq	0x12345678(%r9), %r9
	movzbq	0x12345678(%r9), %r10
	movzbq	0x12345678(%r9), %r11
	movzbq	0x12345678(%r9), %r12
	movzbq	0x12345678(%r9), %r13
	movzbq	0x12345678(%r9), %r14
	movzbq	0x12345678(%r9), %r15
	nop
	movzbq	0x12345678(%r10),%rax
	movzbq	0x12345678(%r10),%rcx
	movzbq	0x12345678(%r10),%rdx
	movzbq	0x12345678(%r10),%rbx
	movzbq	0x12345678(%r10),%rsp
	movzbq	0x12345678(%r10),%rbp
	movzbq	0x12345678(%r10),%rsi
	movzbq	0x12345678(%r10),%rdi
	movzbq	0x12345678(%r10),%r8
	movzbq	0x12345678(%r10),%r9
	movzbq	0x12345678(%r10),%r10
	movzbq	0x12345678(%r10),%r11
	movzbq	0x12345678(%r10),%r12
	movzbq	0x12345678(%r10),%r13
	movzbq	0x12345678(%r10),%r14
	movzbq	0x12345678(%r10),%r15
	nop
	movzbq	0x12345678(%r11),%rax
	movzbq	0x12345678(%r11),%rcx
	movzbq	0x12345678(%r11),%rdx
	movzbq	0x12345678(%r11),%rbx
	movzbq	0x12345678(%r11),%rsp
	movzbq	0x12345678(%r11),%rbp
	movzbq	0x12345678(%r11),%rsi
	movzbq	0x12345678(%r11),%rdi
	movzbq	0x12345678(%r11),%r8
	movzbq	0x12345678(%r11),%r9
	movzbq	0x12345678(%r11),%r10
	movzbq	0x12345678(%r11),%r11
	movzbq	0x12345678(%r11),%r12
	movzbq	0x12345678(%r11),%r13
	movzbq	0x12345678(%r11),%r14
	movzbq	0x12345678(%r11),%r15
	nop
	movzbq	0x12345678(%r12),%rax
	movzbq	0x12345678(%r12),%rcx
	movzbq	0x12345678(%r12),%rdx
	movzbq	0x12345678(%r12),%rbx
	movzbq	0x12345678(%r12),%rsp
	movzbq	0x12345678(%r12),%rbp
	movzbq	0x12345678(%r12),%rsi
	movzbq	0x12345678(%r12),%rdi
	movzbq	0x12345678(%r12),%r8
	movzbq	0x12345678(%r12),%r9
	movzbq	0x12345678(%r12),%r10
	movzbq	0x12345678(%r12),%r11
	movzbq	0x12345678(%r12),%r12
	movzbq	0x12345678(%r12),%r13
	movzbq	0x12345678(%r12),%r14
	movzbq	0x12345678(%r12),%r15
	nop
	movzbq	0x12345678(%r13),%rax
	movzbq	0x12345678(%r13),%rcx
	movzbq	0x12345678(%r13),%rdx
	movzbq	0x12345678(%r13),%rbx
	movzbq	0x12345678(%r13),%rsp
	movzbq	0x12345678(%r13),%rbp
	movzbq	0x12345678(%r13),%rsi
	movzbq	0x12345678(%r13),%rdi
	movzbq	0x12345678(%r13),%r8
	movzbq	0x12345678(%r13),%r9
	movzbq	0x12345678(%r13),%r10
	movzbq	0x12345678(%r13),%r11
	movzbq	0x12345678(%r13),%r12
	movzbq	0x12345678(%r13),%r13
	movzbq	0x12345678(%r13),%r14
	movzbq	0x12345678(%r13),%r15
	nop
	movzbq	0x12345678(%r14),%rax
	movzbq	0x12345678(%r14),%rcx
	movzbq	0x12345678(%r14),%rdx
	movzbq	0x12345678(%r14),%rbx
	movzbq	0x12345678(%r14),%rsp
	movzbq	0x12345678(%r14),%rbp
	movzbq	0x12345678(%r14),%rsi
	movzbq	0x12345678(%r14),%rdi
	movzbq	0x12345678(%r14),%r8
	movzbq	0x12345678(%r14),%r9
	movzbq	0x12345678(%r14),%r10
	movzbq	0x12345678(%r14),%r11
	movzbq	0x12345678(%r14),%r12
	movzbq	0x12345678(%r14),%r13
	movzbq	0x12345678(%r14),%r14
	movzbq	0x12345678(%r14),%r15
	nop
	movzbq	0x12345678(%r15),%rax
	movzbq	0x12345678(%r15),%rcx
	movzbq	0x12345678(%r15),%rdx
	movzbq	0x12345678(%r15),%rbx
	movzbq	0x12345678(%r15),%rsp
	movzbq	0x12345678(%r15),%rbp
	movzbq	0x12345678(%r15),%rsi
	movzbq	0x12345678(%r15),%rdi
	movzbq	0x12345678(%r15),%r8
	movzbq	0x12345678(%r15),%r9
	movzbq	0x12345678(%r15),%r10
	movzbq	0x12345678(%r15),%r11
	movzbq	0x12345678(%r15),%r12
	movzbq	0x12345678(%r15),%r13
	movzbq	0x12345678(%r15),%r14
	movzbq	0x12345678(%r15),%r15
	nop
	nop
	ret
	.cfi_endproc

	.p2align 4,,15
	.globl	Movsbq
	.type	Movsbq, @function
Movsbq:
	.cfi_startproc
	// reg -> reg
	movsbq	%al,%rax
	movsbq	%al,%rcx
	movsbq	%al,%rdx
	movsbq	%al,%rbx
	movsbq	%al,%rsp
	movsbq	%al,%rbp
	movsbq	%al,%rsi
	movsbq	%al,%rdi
	movsbq	%al,%r8
	movsbq	%al,%r9
	movsbq	%al,%r10
	movsbq	%al,%r11
	movsbq	%al,%r12
	movsbq	%al,%r13
	movsbq	%al,%r14
	movsbq	%al,%r15
	nop
	movsbq	%cl,%rax
	movsbq	%cl,%rcx
	movsbq	%cl,%rdx
	movsbq	%cl,%rbx
	movsbq	%cl,%rsp
	movsbq	%cl,%rbp
	movsbq	%cl,%rsi
	movsbq	%cl,%rdi
	movsbq	%cl,%r8
	movsbq	%cl,%r9
	movsbq	%cl,%r10
	movsbq	%cl,%r11
	movsbq	%cl,%r12
	movsbq	%cl,%r13
	movsbq	%cl,%r14
	movsbq	%cl,%r15
	nop
	movsbq	%dl,%rax
	movsbq	%dl,%rcx
	movsbq	%dl,%rdx
	movsbq	%dl,%rbx
	movsbq	%dl,%rsp
	movsbq	%dl,%rbp
	movsbq	%dl,%rsi
	movsbq	%dl,%rdi
	movsbq	%dl,%r8
	movsbq	%dl,%r9
	movsbq	%dl,%r10
	movsbq	%dl,%r11
	movsbq	%dl,%r12
	movsbq	%dl,%r13
	movsbq	%dl,%r14
	movsbq	%dl,%r15
	nop
	movsbq	%bl,%rax
	movsbq	%bl,%rcx
	movsbq	%bl,%rdx
	movsbq	%bl,%rbx
	movsbq	%bl,%rsp
	movsbq	%bl,%rbp
	movsbq	%bl,%rsi
	movsbq	%bl,%rdi
	movsbq	%bl,%r8
	movsbq	%bl,%r9
	movsbq	%bl,%r10
	movsbq	%bl,%r11
	movsbq	%bl,%r12
	movsbq	%bl,%r13
	movsbq	%bl,%r14
	movsbq	%bl,%r15
	nop
	movsbq	%spl,%rax
	movsbq	%spl,%rcx
	movsbq	%spl,%rdx
	movsbq	%spl,%rbx
	movsbq	%spl,%rsp
	movsbq	%spl,%rbp
	movsbq	%spl,%rsi
	movsbq	%spl,%rdi
	movsbq	%spl,%r8
	movsbq	%spl,%r9
	movsbq	%spl,%r10
	movsbq	%spl,%r11
	movsbq	%spl,%r12
	movsbq	%spl,%r13
	movsbq	%spl,%r14
	movsbq	%spl,%r15
	nop
	movsbq	%bpl,%rax
	movsbq	%bpl,%rcx
	movsbq	%bpl,%rdx
	movsbq	%bpl,%rbx
	movsbq	%bpl,%rsp
	movsbq	%bpl,%rbp
	movsbq	%bpl,%rsi
	movsbq	%bpl,%rdi
	movsbq	%bpl,%r8
	movsbq	%bpl,%r9
	movsbq	%bpl,%r10
	movsbq	%bpl,%r11
	movsbq	%bpl,%r12
	movsbq	%bpl,%r13
	movsbq	%bpl,%r14
	movsbq	%bpl,%r15
	nop
	movsbq	%sil,%rax
	movsbq	%sil,%rcx
	movsbq	%sil,%rdx
	movsbq	%sil,%rbx
	movsbq	%sil,%rsp
	movsbq	%sil,%rbp
	movsbq	%sil,%rsi
	movsbq	%sil,%rdi
	movsbq	%sil,%r8
	movsbq	%sil,%r9
	movsbq	%sil,%r10
	movsbq	%sil,%r11
	movsbq	%sil,%r12
	movsbq	%sil,%r13
	movsbq	%sil,%r14
	movsbq	%sil,%r15
	nop
	movsbq	%dil,%rax
	movsbq	%dil,%rcx
	movsbq	%dil,%rdx
	movsbq	%dil,%rbx
	movsbq	%dil,%rsp
	movsbq	%dil,%rbp
	movsbq	%dil,%rsi
	movsbq	%dil,%rdi
	movsbq	%dil,%r8
	movsbq	%dil,%r9
	movsbq	%dil,%r10
	movsbq	%dil,%r11
	movsbq	%dil,%r12
	movsbq	%dil,%r13
	movsbq	%dil,%r14
	movsbq	%dil,%r15
	nop
	movsbq	%r8b, %rax
	movsbq	%r8b, %rcx
	movsbq	%r8b, %rdx
	movsbq	%r8b, %rbx
	movsbq	%r8b, %rsp
	movsbq	%r8b, %rbp
	movsbq	%r8b, %rsi
	movsbq	%r8b, %rdi
	movsbq	%r8b, %r8
	movsbq	%r8b, %r9
	movsbq	%r8b, %r10
	movsbq	%r8b, %r11
	movsbq	%r8b, %r12
	movsbq	%r8b, %r13
	movsbq	%r8b, %r14
	movsbq	%r8b, %r15
	nop
	movsbq	%r9b, %rax
	movsbq	%r9b, %rcx
	movsbq	%r9b, %rdx
	movsbq	%r9b, %rbx
	movsbq	%r9b, %rsp
	movsbq	%r9b, %rbp
	movsbq	%r9b, %rsi
	movsbq	%r9b, %rdi
	movsbq	%r9b, %r8
	movsbq	%r9b, %r9
	movsbq	%r9b, %r10
	movsbq	%r9b, %r11
	movsbq	%r9b, %r12
	movsbq	%r9b, %r13
	movsbq	%r9b, %r14
	movsbq	%r9b, %r15
	nop
	movsbq	%r10b,%rax
	movsbq	%r10b,%rcx
	movsbq	%r10b,%rdx
	movsbq	%r10b,%rbx
	movsbq	%r10b,%rsp
	movsbq	%r10b,%rbp
	movsbq	%r10b,%rsi
	movsbq	%r10b,%rdi
	movsbq	%r10b,%r8
	movsbq	%r10b,%r9
	movsbq	%r10b,%r10
	movsbq	%r10b,%r11
	movsbq	%r10b,%r12
	movsbq	%r10b,%r13
	movsbq	%r10b,%r14
	movsbq	%r10b,%r15
	nop
	movsbq	%r11b,%rax
	movsbq	%r11b,%rcx
	movsbq	%r11b,%rdx
	movsbq	%r11b,%rbx
	movsbq	%r11b,%rsp
	movsbq	%r11b,%rbp
	movsbq	%r11b,%rsi
	movsbq	%r11b,%rdi
	movsbq	%r11b,%r8
	movsbq	%r11b,%r9
	movsbq	%r11b,%r10
	movsbq	%r11b,%r11
	movsbq	%r11b,%r12
	movsbq	%r11b,%r13
	movsbq	%r11b,%r14
	movsbq	%r11b,%r15
	nop
	movsbq	%r12b,%rax
	movsbq	%r12b,%rcx
	movsbq	%r12b,%rdx
	movsbq	%r12b,%rbx
	movsbq	%r12b,%rsp
	movsbq	%r12b,%rbp
	movsbq	%r12b,%rsi
	movsbq	%r12b,%rdi
	movsbq	%r12b,%r8
	movsbq	%r12b,%r9
	movsbq	%r12b,%r10
	movsbq	%r12b,%r11
	movsbq	%r12b,%r12
	movsbq	%r12b,%r13
	movsbq	%r12b,%r14
	movsbq	%r12b,%r15
	nop
	movsbq	%r13b,%rax
	movsbq	%r13b,%rcx
	movsbq	%r13b,%rdx
	movsbq	%r13b,%rbx
	movsbq	%r13b,%rsp
	movsbq	%r13b,%rbp
	movsbq	%r13b,%rsi
	movsbq	%r13b,%rdi
	movsbq	%r13b,%r8
	movsbq	%r13b,%r9
	movsbq	%r13b,%r10
	movsbq	%r13b,%r11
	movsbq	%r13b,%r12
	movsbq	%r13b,%r13
	movsbq	%r13b,%r14
	movsbq	%r13b,%r15
	nop
	movsbq	%r14b,%rax
	movsbq	%r14b,%rcx
	movsbq	%r14b,%rdx
	movsbq	%r14b,%rbx
	movsbq	%r14b,%rsp
	movsbq	%r14b,%rbp
	movsbq	%r14b,%rsi
	movsbq	%r14b,%rdi
	movsbq	%r14b,%r8
	movsbq	%r14b,%r9
	movsbq	%r14b,%r10
	movsbq	%r14b,%r11
	movsbq	%r14b,%r12
	movsbq	%r14b,%r13
	movsbq	%r14b,%r14
	movsbq	%r14b,%r15
	nop
	movsbq	%r15b,%rax
	movsbq	%r15b,%rcx
	movsbq	%r15b,%rdx
	movsbq	%r15b,%rbx
	movsbq	%r15b,%rsp
	movsbq	%r15b,%rbp
	movsbq	%r15b,%rsi
	movsbq	%r15b,%rdi
	movsbq	%r15b,%r8
	movsbq	%r15b,%r9
	movsbq	%r15b,%r10
	movsbq	%r15b,%r11
	movsbq	%r15b,%r12
	movsbq	%r15b,%r13
	movsbq	%r15b,%r14
	movsbq	%r15b,%r15
	nop
        nop
	// (reg) -> reg
	movsbq	(%rax),%rax
	movsbq	(%rax),%rcx
	movsbq	(%rax),%rdx
	movsbq	(%rax),%rbx
	movsbq	(%rax),%rsp
	movsbq	(%rax),%rbp
	movsbq	(%rax),%rsi
	movsbq	(%rax),%rdi
	movsbq	(%rax),%r8
	movsbq	(%rax),%r9
	movsbq	(%rax),%r10
	movsbq	(%rax),%r11
	movsbq	(%rax),%r12
	movsbq	(%rax),%r13
	movsbq	(%rax),%r14
	movsbq	(%rax),%r15
	nop
	movsbq	(%rcx),%rax
	movsbq	(%rcx),%rcx
	movsbq	(%rcx),%rdx
	movsbq	(%rcx),%rbx
	movsbq	(%rcx),%rsp
	movsbq	(%rcx),%rbp
	movsbq	(%rcx),%rsi
	movsbq	(%rcx),%rdi
	movsbq	(%rcx),%r8
	movsbq	(%rcx),%r9
	movsbq	(%rcx),%r10
	movsbq	(%rcx),%r11
	movsbq	(%rcx),%r12
	movsbq	(%rcx),%r13
	movsbq	(%rcx),%r14
	movsbq	(%rcx),%r15
	nop
	movsbq	(%rdx),%rax
	movsbq	(%rdx),%rcx
	movsbq	(%rdx),%rdx
	movsbq	(%rdx),%rbx
	movsbq	(%rdx),%rsp
	movsbq	(%rdx),%rbp
	movsbq	(%rdx),%rsi
	movsbq	(%rdx),%rdi
	movsbq	(%rdx),%r8
	movsbq	(%rdx),%r9
	movsbq	(%rdx),%r10
	movsbq	(%rdx),%r11
	movsbq	(%rdx),%r12
	movsbq	(%rdx),%r13
	movsbq	(%rdx),%r14
	movsbq	(%rdx),%r15
	nop
	movsbq	(%rbx),%rax
	movsbq	(%rbx),%rcx
	movsbq	(%rbx),%rdx
	movsbq	(%rbx),%rbx
	movsbq	(%rbx),%rsp
	movsbq	(%rbx),%rbp
	movsbq	(%rbx),%rsi
	movsbq	(%rbx),%rdi
	movsbq	(%rbx),%r8
	movsbq	(%rbx),%r9
	movsbq	(%rbx),%r10
	movsbq	(%rbx),%r11
	movsbq	(%rbx),%r12
	movsbq	(%rbx),%r13
	movsbq	(%rbx),%r14
	movsbq	(%rbx),%r15
	nop
	movsbq	(%rsp),%rax
	movsbq	(%rsp),%rcx
	movsbq	(%rsp),%rdx
	movsbq	(%rsp),%rbx
	movsbq	(%rsp),%rsp
	movsbq	(%rsp),%rbp
	movsbq	(%rsp),%rsi
	movsbq	(%rsp),%rdi
	movsbq	(%rsp),%r8
	movsbq	(%rsp),%r9
	movsbq	(%rsp),%r10
	movsbq	(%rsp),%r11
	movsbq	(%rsp),%r12
	movsbq	(%rsp),%r13
	movsbq	(%rsp),%r14
	movsbq	(%rsp),%r15
	nop
	movsbq	(%rbp),%rax
	movsbq	(%rbp),%rcx
	movsbq	(%rbp),%rdx
	movsbq	(%rbp),%rbx
	movsbq	(%rbp),%rsp
	movsbq	(%rbp),%rbp
	movsbq	(%rbp),%rsi
	movsbq	(%rbp),%rdi
	movsbq	(%rbp),%r8
	movsbq	(%rbp),%r9
	movsbq	(%rbp),%r10
	movsbq	(%rbp),%r11
	movsbq	(%rbp),%r12
	movsbq	(%rbp),%r13
	movsbq	(%rbp),%r14
	movsbq	(%rbp),%r15
	nop
	movsbq	(%rsi),%rax
	movsbq	(%rsi),%rcx
	movsbq	(%rsi),%rdx
	movsbq	(%rsi),%rbx
	movsbq	(%rsi),%rsp
	movsbq	(%rsi),%rbp
	movsbq	(%rsi),%rsi
	movsbq	(%rsi),%rdi
	movsbq	(%rsi),%r8
	movsbq	(%rsi),%r9
	movsbq	(%rsi),%r10
	movsbq	(%rsi),%r11
	movsbq	(%rsi),%r12
	movsbq	(%rsi),%r13
	movsbq	(%rsi),%r14
	movsbq	(%rsi),%r15
	nop
	movsbq	(%rdi),%rax
	movsbq	(%rdi),%rcx
	movsbq	(%rdi),%rdx
	movsbq	(%rdi),%rbx
	movsbq	(%rdi),%rsp
	movsbq	(%rdi),%rbp
	movsbq	(%rdi),%rsi
	movsbq	(%rdi),%rdi
	movsbq	(%rdi),%r8
	movsbq	(%rdi),%r9
	movsbq	(%rdi),%r10
	movsbq	(%rdi),%r11
	movsbq	(%rdi),%r12
	movsbq	(%rdi),%r13
	movsbq	(%rdi),%r14
	movsbq	(%rdi),%r15
	nop
	movsbq	(%r8), %rax
	movsbq	(%r8), %rcx
	movsbq	(%r8), %rdx
	movsbq	(%r8), %rbx
	movsbq	(%r8), %rsp
	movsbq	(%r8), %rbp
	movsbq	(%r8), %rsi
	movsbq	(%r8), %rdi
	movsbq	(%r8), %r8
	movsbq	(%r8), %r9
	movsbq	(%r8), %r10
	movsbq	(%r8), %r11
	movsbq	(%r8), %r12
	movsbq	(%r8), %r13
	movsbq	(%r8), %r14
	movsbq	(%r8), %r15
	nop
	movsbq	(%r9), %rax
	movsbq	(%r9), %rcx
	movsbq	(%r9), %rdx
	movsbq	(%r9), %rbx
	movsbq	(%r9), %rsp
	movsbq	(%r9), %rbp
	movsbq	(%r9), %rsi
	movsbq	(%r9), %rdi
	movsbq	(%r9), %r8
	movsbq	(%r9), %r9
	movsbq	(%r9), %r10
	movsbq	(%r9), %r11
	movsbq	(%r9), %r12
	movsbq	(%r9), %r13
	movsbq	(%r9), %r14
	movsbq	(%r9), %r15
	nop
	movsbq	(%r10),%rax
	movsbq	(%r10),%rcx
	movsbq	(%r10),%rdx
	movsbq	(%r10),%rbx
	movsbq	(%r10),%rsp
	movsbq	(%r10),%rbp
	movsbq	(%r10),%rsi
	movsbq	(%r10),%rdi
	movsbq	(%r10),%r8
	movsbq	(%r10),%r9
	movsbq	(%r10),%r10
	movsbq	(%r10),%r11
	movsbq	(%r10),%r12
	movsbq	(%r10),%r13
	movsbq	(%r10),%r14
	movsbq	(%r10),%r15
	nop
	movsbq	(%r11),%rax
	movsbq	(%r11),%rcx
	movsbq	(%r11),%rdx
	movsbq	(%r11),%rbx
	movsbq	(%r11),%rsp
	movsbq	(%r11),%rbp
	movsbq	(%r11),%rsi
	movsbq	(%r11),%rdi
	movsbq	(%r11),%r8
	movsbq	(%r11),%r9
	movsbq	(%r11),%r10
	movsbq	(%r11),%r11
	movsbq	(%r11),%r12
	movsbq	(%r11),%r13
	movsbq	(%r11),%r14
	movsbq	(%r11),%r15
	nop
	movsbq	(%r12),%rax
	movsbq	(%r12),%rcx
	movsbq	(%r12),%rdx
	movsbq	(%r12),%rbx
	movsbq	(%r12),%rsp
	movsbq	(%r12),%rbp
	movsbq	(%r12),%rsi
	movsbq	(%r12),%rdi
	movsbq	(%r12),%r8
	movsbq	(%r12),%r9
	movsbq	(%r12),%r10
	movsbq	(%r12),%r11
	movsbq	(%r12),%r12
	movsbq	(%r12),%r13
	movsbq	(%r12),%r14
	movsbq	(%r12),%r15
	nop
	movsbq	(%r13),%rax
	movsbq	(%r13),%rcx
	movsbq	(%r13),%rdx
	movsbq	(%r13),%rbx
	movsbq	(%r13),%rsp
	movsbq	(%r13),%rbp
	movsbq	(%r13),%rsi
	movsbq	(%r13),%rdi
	movsbq	(%r13),%r8
	movsbq	(%r13),%r9
	movsbq	(%r13),%r10
	movsbq	(%r13),%r11
	movsbq	(%r13),%r12
	movsbq	(%r13),%r13
	movsbq	(%r13),%r14
	movsbq	(%r13),%r15
	nop
	movsbq	(%r14),%rax
	movsbq	(%r14),%rcx
	movsbq	(%r14),%rdx
	movsbq	(%r14),%rbx
	movsbq	(%r14),%rsp
	movsbq	(%r14),%rbp
	movsbq	(%r14),%rsi
	movsbq	(%r14),%rdi
	movsbq	(%r14),%r8
	movsbq	(%r14),%r9
	movsbq	(%r14),%r10
	movsbq	(%r14),%r11
	movsbq	(%r14),%r12
	movsbq	(%r14),%r13
	movsbq	(%r14),%r14
	movsbq	(%r14),%r15
	nop
	movsbq	(%r15),%rax
	movsbq	(%r15),%rcx
	movsbq	(%r15),%rdx
	movsbq	(%r15),%rbx
	movsbq	(%r15),%rsp
	movsbq	(%r15),%rbp
	movsbq	(%r15),%rsi
	movsbq	(%r15),%rdi
	movsbq	(%r15),%r8
	movsbq	(%r15),%r9
	movsbq	(%r15),%r10
	movsbq	(%r15),%r11
	movsbq	(%r15),%r12
	movsbq	(%r15),%r13
	movsbq	(%r15),%r14
	movsbq	(%r15),%r15
	nop
	nop
	// off8(reg) -> reg
	movsbq	0x7F(%rax),%rax
	movsbq	0x7F(%rax),%rcx
	movsbq	0x7F(%rax),%rdx
	movsbq	0x7F(%rax),%rbx
	movsbq	0x7F(%rax),%rsp
	movsbq	0x7F(%rax),%rbp
	movsbq	0x7F(%rax),%rsi
	movsbq	0x7F(%rax),%rdi
	movsbq	0x7F(%rax),%r8
	movsbq	0x7F(%rax),%r9
	movsbq	0x7F(%rax),%r10
	movsbq	0x7F(%rax),%r11
	movsbq	0x7F(%rax),%r12
	movsbq	0x7F(%rax),%r13
	movsbq	0x7F(%rax),%r14
	movsbq	0x7F(%rax),%r15
	nop
	movsbq	0x7F(%rcx),%rax
	movsbq	0x7F(%rcx),%rcx
	movsbq	0x7F(%rcx),%rdx
	movsbq	0x7F(%rcx),%rbx
	movsbq	0x7F(%rcx),%rsp
	movsbq	0x7F(%rcx),%rbp
	movsbq	0x7F(%rcx),%rsi
	movsbq	0x7F(%rcx),%rdi
	movsbq	0x7F(%rcx),%r8
	movsbq	0x7F(%rcx),%r9
	movsbq	0x7F(%rcx),%r10
	movsbq	0x7F(%rcx),%r11
	movsbq	0x7F(%rcx),%r12
	movsbq	0x7F(%rcx),%r13
	movsbq	0x7F(%rcx),%r14
	movsbq	0x7F(%rcx),%r15
	nop
	movsbq	0x7F(%rdx),%rax
	movsbq	0x7F(%rdx),%rcx
	movsbq	0x7F(%rdx),%rdx
	movsbq	0x7F(%rdx),%rbx
	movsbq	0x7F(%rdx),%rsp
	movsbq	0x7F(%rdx),%rbp
	movsbq	0x7F(%rdx),%rsi
	movsbq	0x7F(%rdx),%rdi
	movsbq	0x7F(%rdx),%r8
	movsbq	0x7F(%rdx),%r9
	movsbq	0x7F(%rdx),%r10
	movsbq	0x7F(%rdx),%r11
	movsbq	0x7F(%rdx),%r12
	movsbq	0x7F(%rdx),%r13
	movsbq	0x7F(%rdx),%r14
	movsbq	0x7F(%rdx),%r15
	nop
	movsbq	0x7F(%rbx),%rax
	movsbq	0x7F(%rbx),%rcx
	movsbq	0x7F(%rbx),%rdx
	movsbq	0x7F(%rbx),%rbx
	movsbq	0x7F(%rbx),%rsp
	movsbq	0x7F(%rbx),%rbp
	movsbq	0x7F(%rbx),%rsi
	movsbq	0x7F(%rbx),%rdi
	movsbq	0x7F(%rbx),%r8
	movsbq	0x7F(%rbx),%r9
	movsbq	0x7F(%rbx),%r10
	movsbq	0x7F(%rbx),%r11
	movsbq	0x7F(%rbx),%r12
	movsbq	0x7F(%rbx),%r13
	movsbq	0x7F(%rbx),%r14
	movsbq	0x7F(%rbx),%r15
	nop
	movsbq	0x7F(%rsp),%rax
	movsbq	0x7F(%rsp),%rcx
	movsbq	0x7F(%rsp),%rdx
	movsbq	0x7F(%rsp),%rbx
	movsbq	0x7F(%rsp),%rsp
	movsbq	0x7F(%rsp),%rbp
	movsbq	0x7F(%rsp),%rsi
	movsbq	0x7F(%rsp),%rdi
	movsbq	0x7F(%rsp),%r8
	movsbq	0x7F(%rsp),%r9
	movsbq	0x7F(%rsp),%r10
	movsbq	0x7F(%rsp),%r11
	movsbq	0x7F(%rsp),%r12
	movsbq	0x7F(%rsp),%r13
	movsbq	0x7F(%rsp),%r14
	movsbq	0x7F(%rsp),%r15
	nop
	movsbq	0x7F(%rbp),%rax
	movsbq	0x7F(%rbp),%rcx
	movsbq	0x7F(%rbp),%rdx
	movsbq	0x7F(%rbp),%rbx
	movsbq	0x7F(%rbp),%rsp
	movsbq	0x7F(%rbp),%rbp
	movsbq	0x7F(%rbp),%rsi
	movsbq	0x7F(%rbp),%rdi
	movsbq	0x7F(%rbp),%r8
	movsbq	0x7F(%rbp),%r9
	movsbq	0x7F(%rbp),%r10
	movsbq	0x7F(%rbp),%r11
	movsbq	0x7F(%rbp),%r12
	movsbq	0x7F(%rbp),%r13
	movsbq	0x7F(%rbp),%r14
	movsbq	0x7F(%rbp),%r15
	nop
	movsbq	0x7F(%rsi),%rax
	movsbq	0x7F(%rsi),%rcx
	movsbq	0x7F(%rsi),%rdx
	movsbq	0x7F(%rsi),%rbx
	movsbq	0x7F(%rsi),%rsp
	movsbq	0x7F(%rsi),%rbp
	movsbq	0x7F(%rsi),%rsi
	movsbq	0x7F(%rsi),%rdi
	movsbq	0x7F(%rsi),%r8
	movsbq	0x7F(%rsi),%r9
	movsbq	0x7F(%rsi),%r10
	movsbq	0x7F(%rsi),%r11
	movsbq	0x7F(%rsi),%r12
	movsbq	0x7F(%rsi),%r13
	movsbq	0x7F(%rsi),%r14
	movsbq	0x7F(%rsi),%r15
	nop
	movsbq	0x7F(%rdi),%rax
	movsbq	0x7F(%rdi),%rcx
	movsbq	0x7F(%rdi),%rdx
	movsbq	0x7F(%rdi),%rbx
	movsbq	0x7F(%rdi),%rsp
	movsbq	0x7F(%rdi),%rbp
	movsbq	0x7F(%rdi),%rsi
	movsbq	0x7F(%rdi),%rdi
	movsbq	0x7F(%rdi),%r8
	movsbq	0x7F(%rdi),%r9
	movsbq	0x7F(%rdi),%r10
	movsbq	0x7F(%rdi),%r11
	movsbq	0x7F(%rdi),%r12
	movsbq	0x7F(%rdi),%r13
	movsbq	0x7F(%rdi),%r14
	movsbq	0x7F(%rdi),%r15
	nop
	movsbq	0x7F(%r8), %rax
	movsbq	0x7F(%r8), %rcx
	movsbq	0x7F(%r8), %rdx
	movsbq	0x7F(%r8), %rbx
	movsbq	0x7F(%r8), %rsp
	movsbq	0x7F(%r8), %rbp
	movsbq	0x7F(%r8), %rsi
	movsbq	0x7F(%r8), %rdi
	movsbq	0x7F(%r8), %r8
	movsbq	0x7F(%r8), %r9
	movsbq	0x7F(%r8), %r10
	movsbq	0x7F(%r8), %r11
	movsbq	0x7F(%r8), %r12
	movsbq	0x7F(%r8), %r13
	movsbq	0x7F(%r8), %r14
	movsbq	0x7F(%r8), %r15
	nop
	movsbq	0x7F(%r9), %rax
	movsbq	0x7F(%r9), %rcx
	movsbq	0x7F(%r9), %rdx
	movsbq	0x7F(%r9), %rbx
	movsbq	0x7F(%r9), %rsp
	movsbq	0x7F(%r9), %rbp
	movsbq	0x7F(%r9), %rsi
	movsbq	0x7F(%r9), %rdi
	movsbq	0x7F(%r9), %r8
	movsbq	0x7F(%r9), %r9
	movsbq	0x7F(%r9), %r10
	movsbq	0x7F(%r9), %r11
	movsbq	0x7F(%r9), %r12
	movsbq	0x7F(%r9), %r13
	movsbq	0x7F(%r9), %r14
	movsbq	0x7F(%r9), %r15
	nop
	movsbq	0x7F(%r10),%rax
	movsbq	0x7F(%r10),%rcx
	movsbq	0x7F(%r10),%rdx
	movsbq	0x7F(%r10),%rbx
	movsbq	0x7F(%r10),%rsp
	movsbq	0x7F(%r10),%rbp
	movsbq	0x7F(%r10),%rsi
	movsbq	0x7F(%r10),%rdi
	movsbq	0x7F(%r10),%r8
	movsbq	0x7F(%r10),%r9
	movsbq	0x7F(%r10),%r10
	movsbq	0x7F(%r10),%r11
	movsbq	0x7F(%r10),%r12
	movsbq	0x7F(%r10),%r13
	movsbq	0x7F(%r10),%r14
	movsbq	0x7F(%r10),%r15
	nop
	movsbq	0x7F(%r11),%rax
	movsbq	0x7F(%r11),%rcx
	movsbq	0x7F(%r11),%rdx
	movsbq	0x7F(%r11),%rbx
	movsbq	0x7F(%r11),%rsp
	movsbq	0x7F(%r11),%rbp
	movsbq	0x7F(%r11),%rsi
	movsbq	0x7F(%r11),%rdi
	movsbq	0x7F(%r11),%r8
	movsbq	0x7F(%r11),%r9
	movsbq	0x7F(%r11),%r10
	movsbq	0x7F(%r11),%r11
	movsbq	0x7F(%r11),%r12
	movsbq	0x7F(%r11),%r13
	movsbq	0x7F(%r11),%r14
	movsbq	0x7F(%r11),%r15
	nop
	movsbq	0x7F(%r12),%rax
	movsbq	0x7F(%r12),%rcx
	movsbq	0x7F(%r12),%rdx
	movsbq	0x7F(%r12),%rbx
	movsbq	0x7F(%r12),%rsp
	movsbq	0x7F(%r12),%rbp
	movsbq	0x7F(%r12),%rsi
	movsbq	0x7F(%r12),%rdi
	movsbq	0x7F(%r12),%r8
	movsbq	0x7F(%r12),%r9
	movsbq	0x7F(%r12),%r10
	movsbq	0x7F(%r12),%r11
	movsbq	0x7F(%r12),%r12
	movsbq	0x7F(%r12),%r13
	movsbq	0x7F(%r12),%r14
	movsbq	0x7F(%r12),%r15
	nop
	movsbq	0x7F(%r13),%rax
	movsbq	0x7F(%r13),%rcx
	movsbq	0x7F(%r13),%rdx
	movsbq	0x7F(%r13),%rbx
	movsbq	0x7F(%r13),%rsp
	movsbq	0x7F(%r13),%rbp
	movsbq	0x7F(%r13),%rsi
	movsbq	0x7F(%r13),%rdi
	movsbq	0x7F(%r13),%r8
	movsbq	0x7F(%r13),%r9
	movsbq	0x7F(%r13),%r10
	movsbq	0x7F(%r13),%r11
	movsbq	0x7F(%r13),%r12
	movsbq	0x7F(%r13),%r13
	movsbq	0x7F(%r13),%r14
	movsbq	0x7F(%r13),%r15
	nop
	movsbq	0x7F(%r14),%rax
	movsbq	0x7F(%r14),%rcx
	movsbq	0x7F(%r14),%rdx
	movsbq	0x7F(%r14),%rbx
	movsbq	0x7F(%r14),%rsp
	movsbq	0x7F(%r14),%rbp
	movsbq	0x7F(%r14),%rsi
	movsbq	0x7F(%r14),%rdi
	movsbq	0x7F(%r14),%r8
	movsbq	0x7F(%r14),%r9
	movsbq	0x7F(%r14),%r10
	movsbq	0x7F(%r14),%r11
	movsbq	0x7F(%r14),%r12
	movsbq	0x7F(%r14),%r13
	movsbq	0x7F(%r14),%r14
	movsbq	0x7F(%r14),%r15
	nop
	movsbq	0x7F(%r15),%rax
	movsbq	0x7F(%r15),%rcx
	movsbq	0x7F(%r15),%rdx
	movsbq	0x7F(%r15),%rbx
	movsbq	0x7F(%r15),%rsp
	movsbq	0x7F(%r15),%rbp
	movsbq	0x7F(%r15),%rsi
	movsbq	0x7F(%r15),%rdi
	movsbq	0x7F(%r15),%r8
	movsbq	0x7F(%r15),%r9
	movsbq	0x7F(%r15),%r10
	movsbq	0x7F(%r15),%r11
	movsbq	0x7F(%r15),%r12
	movsbq	0x7F(%r15),%r13
	movsbq	0x7F(%r15),%r14
	movsbq	0x7F(%r15),%r15
	nop
	nop
	// off32(reg) -> reg
	movsbq	0x12345678(%rax),%rax
	movsbq	0x12345678(%rax),%rcx
	movsbq	0x12345678(%rax),%rdx
	movsbq	0x12345678(%rax),%rbx
	movsbq	0x12345678(%rax),%rsp
	movsbq	0x12345678(%rax),%rbp
	movsbq	0x12345678(%rax),%rsi
	movsbq	0x12345678(%rax),%rdi
	movsbq	0x12345678(%rax),%r8
	movsbq	0x12345678(%rax),%r9
	movsbq	0x12345678(%rax),%r10
	movsbq	0x12345678(%rax),%r11
	movsbq	0x12345678(%rax),%r12
	movsbq	0x12345678(%rax),%r13
	movsbq	0x12345678(%rax),%r14
	movsbq	0x12345678(%rax),%r15
	nop
	movsbq	0x12345678(%rcx),%rax
	movsbq	0x12345678(%rcx),%rcx
	movsbq	0x12345678(%rcx),%rdx
	movsbq	0x12345678(%rcx),%rbx
	movsbq	0x12345678(%rcx),%rsp
	movsbq	0x12345678(%rcx),%rbp
	movsbq	0x12345678(%rcx),%rsi
	movsbq	0x12345678(%rcx),%rdi
	movsbq	0x12345678(%rcx),%r8
	movsbq	0x12345678(%rcx),%r9
	movsbq	0x12345678(%rcx),%r10
	movsbq	0x12345678(%rcx),%r11
	movsbq	0x12345678(%rcx),%r12
	movsbq	0x12345678(%rcx),%r13
	movsbq	0x12345678(%rcx),%r14
	movsbq	0x12345678(%rcx),%r15
	nop
	movsbq	0x12345678(%rdx),%rax
	movsbq	0x12345678(%rdx),%rcx
	movsbq	0x12345678(%rdx),%rdx
	movsbq	0x12345678(%rdx),%rbx
	movsbq	0x12345678(%rdx),%rsp
	movsbq	0x12345678(%rdx),%rbp
	movsbq	0x12345678(%rdx),%rsi
	movsbq	0x12345678(%rdx),%rdi
	movsbq	0x12345678(%rdx),%r8
	movsbq	0x12345678(%rdx),%r9
	movsbq	0x12345678(%rdx),%r10
	movsbq	0x12345678(%rdx),%r11
	movsbq	0x12345678(%rdx),%r12
	movsbq	0x12345678(%rdx),%r13
	movsbq	0x12345678(%rdx),%r14
	movsbq	0x12345678(%rdx),%r15
	nop
	movsbq	0x12345678(%rbx),%rax
	movsbq	0x12345678(%rbx),%rcx
	movsbq	0x12345678(%rbx),%rdx
	movsbq	0x12345678(%rbx),%rbx
	movsbq	0x12345678(%rbx),%rsp
	movsbq	0x12345678(%rbx),%rbp
	movsbq	0x12345678(%rbx),%rsi
	movsbq	0x12345678(%rbx),%rdi
	movsbq	0x12345678(%rbx),%r8
	movsbq	0x12345678(%rbx),%r9
	movsbq	0x12345678(%rbx),%r10
	movsbq	0x12345678(%rbx),%r11
	movsbq	0x12345678(%rbx),%r12
	movsbq	0x12345678(%rbx),%r13
	movsbq	0x12345678(%rbx),%r14
	movsbq	0x12345678(%rbx),%r15
	nop
	movsbq	0x12345678(%rsp),%rax
	movsbq	0x12345678(%rsp),%rcx
	movsbq	0x12345678(%rsp),%rdx
	movsbq	0x12345678(%rsp),%rbx
	movsbq	0x12345678(%rsp),%rsp
	movsbq	0x12345678(%rsp),%rbp
	movsbq	0x12345678(%rsp),%rsi
	movsbq	0x12345678(%rsp),%rdi
	movsbq	0x12345678(%rsp),%r8
	movsbq	0x12345678(%rsp),%r9
	movsbq	0x12345678(%rsp),%r10
	movsbq	0x12345678(%rsp),%r11
	movsbq	0x12345678(%rsp),%r12
	movsbq	0x12345678(%rsp),%r13
	movsbq	0x12345678(%rsp),%r14
	movsbq	0x12345678(%rsp),%r15
	nop
	movsbq	0x12345678(%rbp),%rax
	movsbq	0x12345678(%rbp),%rcx
	movsbq	0x12345678(%rbp),%rdx
	movsbq	0x12345678(%rbp),%rbx
	movsbq	0x12345678(%rbp),%rsp
	movsbq	0x12345678(%rbp),%rbp
	movsbq	0x12345678(%rbp),%rsi
	movsbq	0x12345678(%rbp),%rdi
	movsbq	0x12345678(%rbp),%r8
	movsbq	0x12345678(%rbp),%r9
	movsbq	0x12345678(%rbp),%r10
	movsbq	0x12345678(%rbp),%r11
	movsbq	0x12345678(%rbp),%r12
	movsbq	0x12345678(%rbp),%r13
	movsbq	0x12345678(%rbp),%r14
	movsbq	0x12345678(%rbp),%r15
	nop
	movsbq	0x12345678(%rsi),%rax
	movsbq	0x12345678(%rsi),%rcx
	movsbq	0x12345678(%rsi),%rdx
	movsbq	0x12345678(%rsi),%rbx
	movsbq	0x12345678(%rsi),%rsp
	movsbq	0x12345678(%rsi),%rbp
	movsbq	0x12345678(%rsi),%rsi
	movsbq	0x12345678(%rsi),%rdi
	movsbq	0x12345678(%rsi),%r8
	movsbq	0x12345678(%rsi),%r9
	movsbq	0x12345678(%rsi),%r10
	movsbq	0x12345678(%rsi),%r11
	movsbq	0x12345678(%rsi),%r12
	movsbq	0x12345678(%rsi),%r13
	movsbq	0x12345678(%rsi),%r14
	movsbq	0x12345678(%rsi),%r15
	nop
	movsbq	0x12345678(%rdi),%rax
	movsbq	0x12345678(%rdi),%rcx
	movsbq	0x12345678(%rdi),%rdx
	movsbq	0x12345678(%rdi),%rbx
	movsbq	0x12345678(%rdi),%rsp
	movsbq	0x12345678(%rdi),%rbp
	movsbq	0x12345678(%rdi),%rsi
	movsbq	0x12345678(%rdi),%rdi
	movsbq	0x12345678(%rdi),%r8
	movsbq	0x12345678(%rdi),%r9
	movsbq	0x12345678(%rdi),%r10
	movsbq	0x12345678(%rdi),%r11
	movsbq	0x12345678(%rdi),%r12
	movsbq	0x12345678(%rdi),%r13
	movsbq	0x12345678(%rdi),%r14
	movsbq	0x12345678(%rdi),%r15
	nop
	movsbq	0x12345678(%r8), %rax
	movsbq	0x12345678(%r8), %rcx
	movsbq	0x12345678(%r8), %rdx
	movsbq	0x12345678(%r8), %rbx
	movsbq	0x12345678(%r8), %rsp
	movsbq	0x12345678(%r8), %rbp
	movsbq	0x12345678(%r8), %rsi
	movsbq	0x12345678(%r8), %rdi
	movsbq	0x12345678(%r8), %r8
	movsbq	0x12345678(%r8), %r9
	movsbq	0x12345678(%r8), %r10
	movsbq	0x12345678(%r8), %r11
	movsbq	0x12345678(%r8), %r12
	movsbq	0x12345678(%r8), %r13
	movsbq	0x12345678(%r8), %r14
	movsbq	0x12345678(%r8), %r15
	nop
	movsbq	0x12345678(%r9), %rax
	movsbq	0x12345678(%r9), %rcx
	movsbq	0x12345678(%r9), %rdx
	movsbq	0x12345678(%r9), %rbx
	movsbq	0x12345678(%r9), %rsp
	movsbq	0x12345678(%r9), %rbp
	movsbq	0x12345678(%r9), %rsi
	movsbq	0x12345678(%r9), %rdi
	movsbq	0x12345678(%r9), %r8
	movsbq	0x12345678(%r9), %r9
	movsbq	0x12345678(%r9), %r10
	movsbq	0x12345678(%r9), %r11
	movsbq	0x12345678(%r9), %r12
	movsbq	0x12345678(%r9), %r13
	movsbq	0x12345678(%r9), %r14
	movsbq	0x12345678(%r9), %r15
	nop
	movsbq	0x12345678(%r10),%rax
	movsbq	0x12345678(%r10),%rcx
	movsbq	0x12345678(%r10),%rdx
	movsbq	0x12345678(%r10),%rbx
	movsbq	0x12345678(%r10),%rsp
	movsbq	0x12345678(%r10),%rbp
	movsbq	0x12345678(%r10),%rsi
	movsbq	0x12345678(%r10),%rdi
	movsbq	0x12345678(%r10),%r8
	movsbq	0x12345678(%r10),%r9
	movsbq	0x12345678(%r10),%r10
	movsbq	0x12345678(%r10),%r11
	movsbq	0x12345678(%r10),%r12
	movsbq	0x12345678(%r10),%r13
	movsbq	0x12345678(%r10),%r14
	movsbq	0x12345678(%r10),%r15
	nop
	movsbq	0x12345678(%r11),%rax
	movsbq	0x12345678(%r11),%rcx
	movsbq	0x12345678(%r11),%rdx
	movsbq	0x12345678(%r11),%rbx
	movsbq	0x12345678(%r11),%rsp
	movsbq	0x12345678(%r11),%rbp
	movsbq	0x12345678(%r11),%rsi
	movsbq	0x12345678(%r11),%rdi
	movsbq	0x12345678(%r11),%r8
	movsbq	0x12345678(%r11),%r9
	movsbq	0x12345678(%r11),%r10
	movsbq	0x12345678(%r11),%r11
	movsbq	0x12345678(%r11),%r12
	movsbq	0x12345678(%r11),%r13
	movsbq	0x12345678(%r11),%r14
	movsbq	0x12345678(%r11),%r15
	nop
	movsbq	0x12345678(%r12),%rax
	movsbq	0x12345678(%r12),%rcx
	movsbq	0x12345678(%r12),%rdx
	movsbq	0x12345678(%r12),%rbx
	movsbq	0x12345678(%r12),%rsp
	movsbq	0x12345678(%r12),%rbp
	movsbq	0x12345678(%r12),%rsi
	movsbq	0x12345678(%r12),%rdi
	movsbq	0x12345678(%r12),%r8
	movsbq	0x12345678(%r12),%r9
	movsbq	0x12345678(%r12),%r10
	movsbq	0x12345678(%r12),%r11
	movsbq	0x12345678(%r12),%r12
	movsbq	0x12345678(%r12),%r13
	movsbq	0x12345678(%r12),%r14
	movsbq	0x12345678(%r12),%r15
	nop
	movsbq	0x12345678(%r13),%rax
	movsbq	0x12345678(%r13),%rcx
	movsbq	0x12345678(%r13),%rdx
	movsbq	0x12345678(%r13),%rbx
	movsbq	0x12345678(%r13),%rsp
	movsbq	0x12345678(%r13),%rbp
	movsbq	0x12345678(%r13),%rsi
	movsbq	0x12345678(%r13),%rdi
	movsbq	0x12345678(%r13),%r8
	movsbq	0x12345678(%r13),%r9
	movsbq	0x12345678(%r13),%r10
	movsbq	0x12345678(%r13),%r11
	movsbq	0x12345678(%r13),%r12
	movsbq	0x12345678(%r13),%r13
	movsbq	0x12345678(%r13),%r14
	movsbq	0x12345678(%r13),%r15
	nop
	movsbq	0x12345678(%r14),%rax
	movsbq	0x12345678(%r14),%rcx
	movsbq	0x12345678(%r14),%rdx
	movsbq	0x12345678(%r14),%rbx
	movsbq	0x12345678(%r14),%rsp
	movsbq	0x12345678(%r14),%rbp
	movsbq	0x12345678(%r14),%rsi
	movsbq	0x12345678(%r14),%rdi
	movsbq	0x12345678(%r14),%r8
	movsbq	0x12345678(%r14),%r9
	movsbq	0x12345678(%r14),%r10
	movsbq	0x12345678(%r14),%r11
	movsbq	0x12345678(%r14),%r12
	movsbq	0x12345678(%r14),%r13
	movsbq	0x12345678(%r14),%r14
	movsbq	0x12345678(%r14),%r15
	nop
	movsbq	0x12345678(%r15),%rax
	movsbq	0x12345678(%r15),%rcx
	movsbq	0x12345678(%r15),%rdx
	movsbq	0x12345678(%r15),%rbx
	movsbq	0x12345678(%r15),%rsp
	movsbq	0x12345678(%r15),%rbp
	movsbq	0x12345678(%r15),%rsi
	movsbq	0x12345678(%r15),%rdi
	movsbq	0x12345678(%r15),%r8
	movsbq	0x12345678(%r15),%r9
	movsbq	0x12345678(%r15),%r10
	movsbq	0x12345678(%r15),%r11
	movsbq	0x12345678(%r15),%r12
	movsbq	0x12345678(%r15),%r13
	movsbq	0x12345678(%r15),%r14
	movsbq	0x12345678(%r15),%r15
	nop
	nop
	ret
	.cfi_endproc

