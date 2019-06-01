	.file	"lea.s"
	.text


	.p2align 4,,15
	.globl	Lea
	.type	Lea, @function
Lea:
	.cfi_startproc
        // reg64 = &mem64[0]
	lea	(%rax),%rax
	lea	(%rax),%rcx
	lea	(%rax),%rdx
	lea	(%rax),%rbx
	lea	(%rax),%rsp
	lea	(%rax),%rbp
	lea	(%rax),%rsi
	lea	(%rax),%rdi
	lea	(%rax),%r8
	lea	(%rax),%r9
	lea	(%rax),%r10
	lea	(%rax),%r11
	lea	(%rax),%r12
	lea	(%rax),%r13
	lea	(%rax),%r14
	lea	(%rax),%r15
	nop
	lea	(%rcx),%rax
	lea	(%rcx),%rcx
	lea	(%rcx),%rdx
	lea	(%rcx),%rbx
	lea	(%rcx),%rsp
	lea	(%rcx),%rbp
	lea	(%rcx),%rsi
	lea	(%rcx),%rdi
	lea	(%rcx),%r8
	lea	(%rcx),%r9
	lea	(%rcx),%r10
	lea	(%rcx),%r11
	lea	(%rcx),%r12
	lea	(%rcx),%r13
	lea	(%rcx),%r14
	lea	(%rcx),%r15
	nop
	lea	(%rdx),%rax
	lea	(%rdx),%rcx
	lea	(%rdx),%rdx
	lea	(%rdx),%rbx
	lea	(%rdx),%rsp
	lea	(%rdx),%rbp
	lea	(%rdx),%rsi
	lea	(%rdx),%rdi
	lea	(%rdx),%r8
	lea	(%rdx),%r9
	lea	(%rdx),%r10
	lea	(%rdx),%r11
	lea	(%rdx),%r12
	lea	(%rdx),%r13
	lea	(%rdx),%r14
	lea	(%rdx),%r15
	nop
	lea	(%rbx),%rax
	lea	(%rbx),%rcx
	lea	(%rbx),%rdx
	lea	(%rbx),%rbx
	lea	(%rbx),%rsp
	lea	(%rbx),%rbp
	lea	(%rbx),%rsi
	lea	(%rbx),%rdi
	lea	(%rbx),%r8
	lea	(%rbx),%r9
	lea	(%rbx),%r10
	lea	(%rbx),%r11
	lea	(%rbx),%r12
	lea	(%rbx),%r13
	lea	(%rbx),%r14
	lea	(%rbx),%r15
	nop
	lea	(%rsp),%rax
	lea	(%rsp),%rcx
	lea	(%rsp),%rdx
	lea	(%rsp),%rbx
	lea	(%rsp),%rsp
	lea	(%rsp),%rbp
	lea	(%rsp),%rsi
	lea	(%rsp),%rdi
	lea	(%rsp),%r8
	lea	(%rsp),%r9
	lea	(%rsp),%r10
	lea	(%rsp),%r11
	lea	(%rsp),%r12
	lea	(%rsp),%r13
	lea	(%rsp),%r14
	lea	(%rsp),%r15
	nop
	lea	(%rbp),%rax
	lea	(%rbp),%rcx
	lea	(%rbp),%rdx
	lea	(%rbp),%rbx
	lea	(%rbp),%rsp
	lea	(%rbp),%rbp
	lea	(%rbp),%rsi
	lea	(%rbp),%rdi
	lea	(%rbp),%r8
	lea	(%rbp),%r9
	lea	(%rbp),%r10
	lea	(%rbp),%r11
	lea	(%rbp),%r12
	lea	(%rbp),%r13
	lea	(%rbp),%r14
	lea	(%rbp),%r15
	nop
	lea	(%rsi),%rax
	lea	(%rsi),%rcx
	lea	(%rsi),%rdx
	lea	(%rsi),%rbx
	lea	(%rsi),%rsp
	lea	(%rsi),%rbp
	lea	(%rsi),%rsi
	lea	(%rsi),%rdi
	lea	(%rsi),%r8
	lea	(%rsi),%r9
	lea	(%rsi),%r10
	lea	(%rsi),%r11
	lea	(%rsi),%r12
	lea	(%rsi),%r13
	lea	(%rsi),%r14
	lea	(%rsi),%r15
	nop
	lea	(%rdi),%rax
	lea	(%rdi),%rcx
	lea	(%rdi),%rdx
	lea	(%rdi),%rbx
	lea	(%rdi),%rsp
	lea	(%rdi),%rbp
	lea	(%rdi),%rsi
	lea	(%rdi),%rdi
	lea	(%rdi),%r8
	lea	(%rdi),%r9
	lea	(%rdi),%r10
	lea	(%rdi),%r11
	lea	(%rdi),%r12
	lea	(%rdi),%r13
	lea	(%rdi),%r14
	lea	(%rdi),%r15
	nop
	lea	(%r8), %rax
	lea	(%r8), %rcx
	lea	(%r8), %rdx
	lea	(%r8), %rbx
	lea	(%r8), %rsp
	lea	(%r8), %rbp
	lea	(%r8), %rsi
	lea	(%r8), %rdi
	lea	(%r8), %r8
	lea	(%r8), %r9
	lea	(%r8), %r10
	lea	(%r8), %r11
	lea	(%r8), %r12
	lea	(%r8), %r13
	lea	(%r8), %r14
	lea	(%r8), %r15
	nop
	lea	(%r9), %rax
	lea	(%r9), %rcx
	lea	(%r9), %rdx
	lea	(%r9), %rbx
	lea	(%r9), %rsp
	lea	(%r9), %rbp
	lea	(%r9), %rsi
	lea	(%r9), %rdi
	lea	(%r9), %r8
	lea	(%r9), %r9
	lea	(%r9), %r10
	lea	(%r9), %r11
	lea	(%r9), %r12
	lea	(%r9), %r13
	lea	(%r9), %r14
	lea	(%r9), %r15
	nop
	lea	(%r10),%rax
	lea	(%r10),%rcx
	lea	(%r10),%rdx
	lea	(%r10),%rbx
	lea	(%r10),%rsp
	lea	(%r10),%rbp
	lea	(%r10),%rsi
	lea	(%r10),%rdi
	lea	(%r10),%r8
	lea	(%r10),%r9
	lea	(%r10),%r10
	lea	(%r10),%r11
	lea	(%r10),%r12
	lea	(%r10),%r13
	lea	(%r10),%r14
	lea	(%r10),%r15
	nop
	lea	(%r11),%rax
	lea	(%r11),%rcx
	lea	(%r11),%rdx
	lea	(%r11),%rbx
	lea	(%r11),%rsp
	lea	(%r11),%rbp
	lea	(%r11),%rsi
	lea	(%r11),%rdi
	lea	(%r11),%r8
	lea	(%r11),%r9
	lea	(%r11),%r10
	lea	(%r11),%r11
	lea	(%r11),%r12
	lea	(%r11),%r13
	lea	(%r11),%r14
	lea	(%r11),%r15
	nop
	lea	(%r12),%rax
	lea	(%r12),%rcx
	lea	(%r12),%rdx
	lea	(%r12),%rbx
	lea	(%r12),%rsp
	lea	(%r12),%rbp
	lea	(%r12),%rsi
	lea	(%r12),%rdi
	lea	(%r12),%r8
	lea	(%r12),%r9
	lea	(%r12),%r10
	lea	(%r12),%r11
	lea	(%r12),%r12
	lea	(%r12),%r13
	lea	(%r12),%r14
	lea	(%r12),%r15
	nop
	lea	(%r13),%rax
	lea	(%r13),%rcx
	lea	(%r13),%rdx
	lea	(%r13),%rbx
	lea	(%r13),%rsp
	lea	(%r13),%rbp
	lea	(%r13),%rsi
	lea	(%r13),%rdi
	lea	(%r13),%r8
	lea	(%r13),%r9
	lea	(%r13),%r10
	lea	(%r13),%r11
	lea	(%r13),%r12
	lea	(%r13),%r13
	lea	(%r13),%r14
	lea	(%r13),%r15
	nop
	lea	(%r14),%rax
	lea	(%r14),%rcx
	lea	(%r14),%rdx
	lea	(%r14),%rbx
	lea	(%r14),%rsp
	lea	(%r14),%rbp
	lea	(%r14),%rsi
	lea	(%r14),%rdi
	lea	(%r14),%r8
	lea	(%r14),%r9
	lea	(%r14),%r10
	lea	(%r14),%r11
	lea	(%r14),%r12
	lea	(%r14),%r13
	lea	(%r14),%r14
	lea	(%r14),%r15
	nop
	lea	(%r15),%rax
	lea	(%r15),%rcx
	lea	(%r15),%rdx
	lea	(%r15),%rbx
	lea	(%r15),%rsp
	lea	(%r15),%rbp
	lea	(%r15),%rsi
	lea	(%r15),%rdi
	lea	(%r15),%r8
	lea	(%r15),%r9
	lea	(%r15),%r10
	lea	(%r15),%r11
	lea	(%r15),%r12
	lea	(%r15),%r13
	lea	(%r15),%r14
	lea	(%r15),%r15
        nop
        nop
        // reg64 = &mem64[off8]
	lea	0x7F(%rax),%rax
	lea	0x7F(%rax),%rcx
	lea	0x7F(%rax),%rdx
	lea	0x7F(%rax),%rbx
	lea	0x7F(%rax),%rsp
	lea	0x7F(%rax),%rbp
	lea	0x7F(%rax),%rsi
	lea	0x7F(%rax),%rdi
	lea	0x7F(%rax),%r8
	lea	0x7F(%rax),%r9
	lea	0x7F(%rax),%r10
	lea	0x7F(%rax),%r11
	lea	0x7F(%rax),%r12
	lea	0x7F(%rax),%r13
	lea	0x7F(%rax),%r14
	lea	0x7F(%rax),%r15
	nop
	lea	0x7F(%rcx),%rax
	lea	0x7F(%rcx),%rcx
	lea	0x7F(%rcx),%rdx
	lea	0x7F(%rcx),%rbx
	lea	0x7F(%rcx),%rsp
	lea	0x7F(%rcx),%rbp
	lea	0x7F(%rcx),%rsi
	lea	0x7F(%rcx),%rdi
	lea	0x7F(%rcx),%r8
	lea	0x7F(%rcx),%r9
	lea	0x7F(%rcx),%r10
	lea	0x7F(%rcx),%r11
	lea	0x7F(%rcx),%r12
	lea	0x7F(%rcx),%r13
	lea	0x7F(%rcx),%r14
	lea	0x7F(%rcx),%r15
	nop
	lea	0x7F(%rdx),%rax
	lea	0x7F(%rdx),%rcx
	lea	0x7F(%rdx),%rdx
	lea	0x7F(%rdx),%rbx
	lea	0x7F(%rdx),%rsp
	lea	0x7F(%rdx),%rbp
	lea	0x7F(%rdx),%rsi
	lea	0x7F(%rdx),%rdi
	lea	0x7F(%rdx),%r8
	lea	0x7F(%rdx),%r9
	lea	0x7F(%rdx),%r10
	lea	0x7F(%rdx),%r11
	lea	0x7F(%rdx),%r12
	lea	0x7F(%rdx),%r13
	lea	0x7F(%rdx),%r14
	lea	0x7F(%rdx),%r15
	nop
	lea	0x7F(%rbx),%rax
	lea	0x7F(%rbx),%rcx
	lea	0x7F(%rbx),%rdx
	lea	0x7F(%rbx),%rbx
	lea	0x7F(%rbx),%rsp
	lea	0x7F(%rbx),%rbp
	lea	0x7F(%rbx),%rsi
	lea	0x7F(%rbx),%rdi
	lea	0x7F(%rbx),%r8
	lea	0x7F(%rbx),%r9
	lea	0x7F(%rbx),%r10
	lea	0x7F(%rbx),%r11
	lea	0x7F(%rbx),%r12
	lea	0x7F(%rbx),%r13
	lea	0x7F(%rbx),%r14
	lea	0x7F(%rbx),%r15
	nop
	lea	0x7F(%rsp),%rax
	lea	0x7F(%rsp),%rcx
	lea	0x7F(%rsp),%rdx
	lea	0x7F(%rsp),%rbx
	lea	0x7F(%rsp),%rsp
	lea	0x7F(%rsp),%rbp
	lea	0x7F(%rsp),%rsi
	lea	0x7F(%rsp),%rdi
	lea	0x7F(%rsp),%r8
	lea	0x7F(%rsp),%r9
	lea	0x7F(%rsp),%r10
	lea	0x7F(%rsp),%r11
	lea	0x7F(%rsp),%r12
	lea	0x7F(%rsp),%r13
	lea	0x7F(%rsp),%r14
	lea	0x7F(%rsp),%r15
	nop
	lea	0x7F(%rbp),%rax
	lea	0x7F(%rbp),%rcx
	lea	0x7F(%rbp),%rdx
	lea	0x7F(%rbp),%rbx
	lea	0x7F(%rbp),%rsp
	lea	0x7F(%rbp),%rbp
	lea	0x7F(%rbp),%rsi
	lea	0x7F(%rbp),%rdi
	lea	0x7F(%rbp),%r8
	lea	0x7F(%rbp),%r9
	lea	0x7F(%rbp),%r10
	lea	0x7F(%rbp),%r11
	lea	0x7F(%rbp),%r12
	lea	0x7F(%rbp),%r13
	lea	0x7F(%rbp),%r14
	lea	0x7F(%rbp),%r15
	nop
	lea	0x7F(%rsi),%rax
	lea	0x7F(%rsi),%rcx
	lea	0x7F(%rsi),%rdx
	lea	0x7F(%rsi),%rbx
	lea	0x7F(%rsi),%rsp
	lea	0x7F(%rsi),%rbp
	lea	0x7F(%rsi),%rsi
	lea	0x7F(%rsi),%rdi
	lea	0x7F(%rsi),%r8
	lea	0x7F(%rsi),%r9
	lea	0x7F(%rsi),%r10
	lea	0x7F(%rsi),%r11
	lea	0x7F(%rsi),%r12
	lea	0x7F(%rsi),%r13
	lea	0x7F(%rsi),%r14
	lea	0x7F(%rsi),%r15
	nop
	lea	0x7F(%rdi),%rax
	lea	0x7F(%rdi),%rcx
	lea	0x7F(%rdi),%rdx
	lea	0x7F(%rdi),%rbx
	lea	0x7F(%rdi),%rsp
	lea	0x7F(%rdi),%rbp
	lea	0x7F(%rdi),%rsi
	lea	0x7F(%rdi),%rdi
	lea	0x7F(%rdi),%r8
	lea	0x7F(%rdi),%r9
	lea	0x7F(%rdi),%r10
	lea	0x7F(%rdi),%r11
	lea	0x7F(%rdi),%r12
	lea	0x7F(%rdi),%r13
	lea	0x7F(%rdi),%r14
	lea	0x7F(%rdi),%r15
	nop
	lea	0x7F(%r8), %rax
	lea	0x7F(%r8), %rcx
	lea	0x7F(%r8), %rdx
	lea	0x7F(%r8), %rbx
	lea	0x7F(%r8), %rsp
	lea	0x7F(%r8), %rbp
	lea	0x7F(%r8), %rsi
	lea	0x7F(%r8), %rdi
	lea	0x7F(%r8), %r8
	lea	0x7F(%r8), %r9
	lea	0x7F(%r8), %r10
	lea	0x7F(%r8), %r11
	lea	0x7F(%r8), %r12
	lea	0x7F(%r8), %r13
	lea	0x7F(%r8), %r14
	lea	0x7F(%r8), %r15
	nop
	lea	0x7F(%r9), %rax
	lea	0x7F(%r9), %rcx
	lea	0x7F(%r9), %rdx
	lea	0x7F(%r9), %rbx
	lea	0x7F(%r9), %rsp
	lea	0x7F(%r9), %rbp
	lea	0x7F(%r9), %rsi
	lea	0x7F(%r9), %rdi
	lea	0x7F(%r9), %r8
	lea	0x7F(%r9), %r9
	lea	0x7F(%r9), %r10
	lea	0x7F(%r9), %r11
	lea	0x7F(%r9), %r12
	lea	0x7F(%r9), %r13
	lea	0x7F(%r9), %r14
	lea	0x7F(%r9), %r15
	nop
	lea	0x7F(%r10),%rax
	lea	0x7F(%r10),%rcx
	lea	0x7F(%r10),%rdx
	lea	0x7F(%r10),%rbx
	lea	0x7F(%r10),%rsp
	lea	0x7F(%r10),%rbp
	lea	0x7F(%r10),%rsi
	lea	0x7F(%r10),%rdi
	lea	0x7F(%r10),%r8
	lea	0x7F(%r10),%r9
	lea	0x7F(%r10),%r10
	lea	0x7F(%r10),%r11
	lea	0x7F(%r10),%r12
	lea	0x7F(%r10),%r13
	lea	0x7F(%r10),%r14
	lea	0x7F(%r10),%r15
	nop
	lea	0x7F(%r11),%rax
	lea	0x7F(%r11),%rcx
	lea	0x7F(%r11),%rdx
	lea	0x7F(%r11),%rbx
	lea	0x7F(%r11),%rsp
	lea	0x7F(%r11),%rbp
	lea	0x7F(%r11),%rsi
	lea	0x7F(%r11),%rdi
	lea	0x7F(%r11),%r8
	lea	0x7F(%r11),%r9
	lea	0x7F(%r11),%r10
	lea	0x7F(%r11),%r11
	lea	0x7F(%r11),%r12
	lea	0x7F(%r11),%r13
	lea	0x7F(%r11),%r14
	lea	0x7F(%r11),%r15
	nop
	lea	0x7F(%r12),%rax
	lea	0x7F(%r12),%rcx
	lea	0x7F(%r12),%rdx
	lea	0x7F(%r12),%rbx
	lea	0x7F(%r12),%rsp
	lea	0x7F(%r12),%rbp
	lea	0x7F(%r12),%rsi
	lea	0x7F(%r12),%rdi
	lea	0x7F(%r12),%r8
	lea	0x7F(%r12),%r9
	lea	0x7F(%r12),%r10
	lea	0x7F(%r12),%r11
	lea	0x7F(%r12),%r12
	lea	0x7F(%r12),%r13
	lea	0x7F(%r12),%r14
	lea	0x7F(%r12),%r15
	nop
	lea	0x7F(%r13),%rax
	lea	0x7F(%r13),%rcx
	lea	0x7F(%r13),%rdx
	lea	0x7F(%r13),%rbx
	lea	0x7F(%r13),%rsp
	lea	0x7F(%r13),%rbp
	lea	0x7F(%r13),%rsi
	lea	0x7F(%r13),%rdi
	lea	0x7F(%r13),%r8
	lea	0x7F(%r13),%r9
	lea	0x7F(%r13),%r10
	lea	0x7F(%r13),%r11
	lea	0x7F(%r13),%r12
	lea	0x7F(%r13),%r13
	lea	0x7F(%r13),%r14
	lea	0x7F(%r13),%r15
	nop
	lea	0x7F(%r14),%rax
	lea	0x7F(%r14),%rcx
	lea	0x7F(%r14),%rdx
	lea	0x7F(%r14),%rbx
	lea	0x7F(%r14),%rsp
	lea	0x7F(%r14),%rbp
	lea	0x7F(%r14),%rsi
	lea	0x7F(%r14),%rdi
	lea	0x7F(%r14),%r8
	lea	0x7F(%r14),%r9
	lea	0x7F(%r14),%r10
	lea	0x7F(%r14),%r11
	lea	0x7F(%r14),%r12
	lea	0x7F(%r14),%r13
	lea	0x7F(%r14),%r14
	lea	0x7F(%r14),%r15
	nop
	lea	0x7F(%r15),%rax
	lea	0x7F(%r15),%rcx
	lea	0x7F(%r15),%rdx
	lea	0x7F(%r15),%rbx
	lea	0x7F(%r15),%rsp
	lea	0x7F(%r15),%rbp
	lea	0x7F(%r15),%rsi
	lea	0x7F(%r15),%rdi
	lea	0x7F(%r15),%r8
	lea	0x7F(%r15),%r9
	lea	0x7F(%r15),%r10
	lea	0x7F(%r15),%r11
	lea	0x7F(%r15),%r12
	lea	0x7F(%r15),%r13
	lea	0x7F(%r15),%r14
	lea	0x7F(%r15),%r15
        nop
        nop
        // reg64 = &mem64[off32]
	lea	0x12345678(%rax),%rax
	lea	0x12345678(%rax),%rcx
	lea	0x12345678(%rax),%rdx
	lea	0x12345678(%rax),%rbx
	lea	0x12345678(%rax),%rsp
	lea	0x12345678(%rax),%rbp
	lea	0x12345678(%rax),%rsi
	lea	0x12345678(%rax),%rdi
	lea	0x12345678(%rax),%r8
	lea	0x12345678(%rax),%r9
	lea	0x12345678(%rax),%r10
	lea	0x12345678(%rax),%r11
	lea	0x12345678(%rax),%r12
	lea	0x12345678(%rax),%r13
	lea	0x12345678(%rax),%r14
	lea	0x12345678(%rax),%r15
	nop
	lea	0x12345678(%rcx),%rax
	lea	0x12345678(%rcx),%rcx
	lea	0x12345678(%rcx),%rdx
	lea	0x12345678(%rcx),%rbx
	lea	0x12345678(%rcx),%rsp
	lea	0x12345678(%rcx),%rbp
	lea	0x12345678(%rcx),%rsi
	lea	0x12345678(%rcx),%rdi
	lea	0x12345678(%rcx),%r8
	lea	0x12345678(%rcx),%r9
	lea	0x12345678(%rcx),%r10
	lea	0x12345678(%rcx),%r11
	lea	0x12345678(%rcx),%r12
	lea	0x12345678(%rcx),%r13
	lea	0x12345678(%rcx),%r14
	lea	0x12345678(%rcx),%r15
	nop
	lea	0x12345678(%rdx),%rax
	lea	0x12345678(%rdx),%rcx
	lea	0x12345678(%rdx),%rdx
	lea	0x12345678(%rdx),%rbx
	lea	0x12345678(%rdx),%rsp
	lea	0x12345678(%rdx),%rbp
	lea	0x12345678(%rdx),%rsi
	lea	0x12345678(%rdx),%rdi
	lea	0x12345678(%rdx),%r8
	lea	0x12345678(%rdx),%r9
	lea	0x12345678(%rdx),%r10
	lea	0x12345678(%rdx),%r11
	lea	0x12345678(%rdx),%r12
	lea	0x12345678(%rdx),%r13
	lea	0x12345678(%rdx),%r14
	lea	0x12345678(%rdx),%r15
	nop
	lea	0x12345678(%rbx),%rax
	lea	0x12345678(%rbx),%rcx
	lea	0x12345678(%rbx),%rdx
	lea	0x12345678(%rbx),%rbx
	lea	0x12345678(%rbx),%rsp
	lea	0x12345678(%rbx),%rbp
	lea	0x12345678(%rbx),%rsi
	lea	0x12345678(%rbx),%rdi
	lea	0x12345678(%rbx),%r8
	lea	0x12345678(%rbx),%r9
	lea	0x12345678(%rbx),%r10
	lea	0x12345678(%rbx),%r11
	lea	0x12345678(%rbx),%r12
	lea	0x12345678(%rbx),%r13
	lea	0x12345678(%rbx),%r14
	lea	0x12345678(%rbx),%r15
	nop
	lea	0x12345678(%rsp),%rax
	lea	0x12345678(%rsp),%rcx
	lea	0x12345678(%rsp),%rdx
	lea	0x12345678(%rsp),%rbx
	lea	0x12345678(%rsp),%rsp
	lea	0x12345678(%rsp),%rbp
	lea	0x12345678(%rsp),%rsi
	lea	0x12345678(%rsp),%rdi
	lea	0x12345678(%rsp),%r8
	lea	0x12345678(%rsp),%r9
	lea	0x12345678(%rsp),%r10
	lea	0x12345678(%rsp),%r11
	lea	0x12345678(%rsp),%r12
	lea	0x12345678(%rsp),%r13
	lea	0x12345678(%rsp),%r14
	lea	0x12345678(%rsp),%r15
	nop
	lea	0x12345678(%rbp),%rax
	lea	0x12345678(%rbp),%rcx
	lea	0x12345678(%rbp),%rdx
	lea	0x12345678(%rbp),%rbx
	lea	0x12345678(%rbp),%rsp
	lea	0x12345678(%rbp),%rbp
	lea	0x12345678(%rbp),%rsi
	lea	0x12345678(%rbp),%rdi
	lea	0x12345678(%rbp),%r8
	lea	0x12345678(%rbp),%r9
	lea	0x12345678(%rbp),%r10
	lea	0x12345678(%rbp),%r11
	lea	0x12345678(%rbp),%r12
	lea	0x12345678(%rbp),%r13
	lea	0x12345678(%rbp),%r14
	lea	0x12345678(%rbp),%r15
	nop
	lea	0x12345678(%rsi),%rax
	lea	0x12345678(%rsi),%rcx
	lea	0x12345678(%rsi),%rdx
	lea	0x12345678(%rsi),%rbx
	lea	0x12345678(%rsi),%rsp
	lea	0x12345678(%rsi),%rbp
	lea	0x12345678(%rsi),%rsi
	lea	0x12345678(%rsi),%rdi
	lea	0x12345678(%rsi),%r8
	lea	0x12345678(%rsi),%r9
	lea	0x12345678(%rsi),%r10
	lea	0x12345678(%rsi),%r11
	lea	0x12345678(%rsi),%r12
	lea	0x12345678(%rsi),%r13
	lea	0x12345678(%rsi),%r14
	lea	0x12345678(%rsi),%r15
	nop
	lea	0x12345678(%rdi),%rax
	lea	0x12345678(%rdi),%rcx
	lea	0x12345678(%rdi),%rdx
	lea	0x12345678(%rdi),%rbx
	lea	0x12345678(%rdi),%rsp
	lea	0x12345678(%rdi),%rbp
	lea	0x12345678(%rdi),%rsi
	lea	0x12345678(%rdi),%rdi
	lea	0x12345678(%rdi),%r8
	lea	0x12345678(%rdi),%r9
	lea	0x12345678(%rdi),%r10
	lea	0x12345678(%rdi),%r11
	lea	0x12345678(%rdi),%r12
	lea	0x12345678(%rdi),%r13
	lea	0x12345678(%rdi),%r14
	lea	0x12345678(%rdi),%r15
	nop
	lea	0x12345678(%r8), %rax
	lea	0x12345678(%r8), %rcx
	lea	0x12345678(%r8), %rdx
	lea	0x12345678(%r8), %rbx
	lea	0x12345678(%r8), %rsp
	lea	0x12345678(%r8), %rbp
	lea	0x12345678(%r8), %rsi
	lea	0x12345678(%r8), %rdi
	lea	0x12345678(%r8), %r8
	lea	0x12345678(%r8), %r9
	lea	0x12345678(%r8), %r10
	lea	0x12345678(%r8), %r11
	lea	0x12345678(%r8), %r12
	lea	0x12345678(%r8), %r13
	lea	0x12345678(%r8), %r14
	lea	0x12345678(%r8), %r15
	nop
	lea	0x12345678(%r9), %rax
	lea	0x12345678(%r9), %rcx
	lea	0x12345678(%r9), %rdx
	lea	0x12345678(%r9), %rbx
	lea	0x12345678(%r9), %rsp
	lea	0x12345678(%r9), %rbp
	lea	0x12345678(%r9), %rsi
	lea	0x12345678(%r9), %rdi
	lea	0x12345678(%r9), %r8
	lea	0x12345678(%r9), %r9
	lea	0x12345678(%r9), %r10
	lea	0x12345678(%r9), %r11
	lea	0x12345678(%r9), %r12
	lea	0x12345678(%r9), %r13
	lea	0x12345678(%r9), %r14
	lea	0x12345678(%r9), %r15
	nop
	lea	0x12345678(%r10),%rax
	lea	0x12345678(%r10),%rcx
	lea	0x12345678(%r10),%rdx
	lea	0x12345678(%r10),%rbx
	lea	0x12345678(%r10),%rsp
	lea	0x12345678(%r10),%rbp
	lea	0x12345678(%r10),%rsi
	lea	0x12345678(%r10),%rdi
	lea	0x12345678(%r10),%r8
	lea	0x12345678(%r10),%r9
	lea	0x12345678(%r10),%r10
	lea	0x12345678(%r10),%r11
	lea	0x12345678(%r10),%r12
	lea	0x12345678(%r10),%r13
	lea	0x12345678(%r10),%r14
	lea	0x12345678(%r10),%r15
	nop
	lea	0x12345678(%r11),%rax
	lea	0x12345678(%r11),%rcx
	lea	0x12345678(%r11),%rdx
	lea	0x12345678(%r11),%rbx
	lea	0x12345678(%r11),%rsp
	lea	0x12345678(%r11),%rbp
	lea	0x12345678(%r11),%rsi
	lea	0x12345678(%r11),%rdi
	lea	0x12345678(%r11),%r8
	lea	0x12345678(%r11),%r9
	lea	0x12345678(%r11),%r10
	lea	0x12345678(%r11),%r11
	lea	0x12345678(%r11),%r12
	lea	0x12345678(%r11),%r13
	lea	0x12345678(%r11),%r14
	lea	0x12345678(%r11),%r15
	nop
	lea	0x12345678(%r12),%rax
	lea	0x12345678(%r12),%rcx
	lea	0x12345678(%r12),%rdx
	lea	0x12345678(%r12),%rbx
	lea	0x12345678(%r12),%rsp
	lea	0x12345678(%r12),%rbp
	lea	0x12345678(%r12),%rsi
	lea	0x12345678(%r12),%rdi
	lea	0x12345678(%r12),%r8
	lea	0x12345678(%r12),%r9
	lea	0x12345678(%r12),%r10
	lea	0x12345678(%r12),%r11
	lea	0x12345678(%r12),%r12
	lea	0x12345678(%r12),%r13
	lea	0x12345678(%r12),%r14
	lea	0x12345678(%r12),%r15
	nop
	lea	0x12345678(%r13),%rax
	lea	0x12345678(%r13),%rcx
	lea	0x12345678(%r13),%rdx
	lea	0x12345678(%r13),%rbx
	lea	0x12345678(%r13),%rsp
	lea	0x12345678(%r13),%rbp
	lea	0x12345678(%r13),%rsi
	lea	0x12345678(%r13),%rdi
	lea	0x12345678(%r13),%r8
	lea	0x12345678(%r13),%r9
	lea	0x12345678(%r13),%r10
	lea	0x12345678(%r13),%r11
	lea	0x12345678(%r13),%r12
	lea	0x12345678(%r13),%r13
	lea	0x12345678(%r13),%r14
	lea	0x12345678(%r13),%r15
	nop
	lea	0x12345678(%r14),%rax
	lea	0x12345678(%r14),%rcx
	lea	0x12345678(%r14),%rdx
	lea	0x12345678(%r14),%rbx
	lea	0x12345678(%r14),%rsp
	lea	0x12345678(%r14),%rbp
	lea	0x12345678(%r14),%rsi
	lea	0x12345678(%r14),%rdi
	lea	0x12345678(%r14),%r8
	lea	0x12345678(%r14),%r9
	lea	0x12345678(%r14),%r10
	lea	0x12345678(%r14),%r11
	lea	0x12345678(%r14),%r12
	lea	0x12345678(%r14),%r13
	lea	0x12345678(%r14),%r14
	lea	0x12345678(%r14),%r15
	nop
	lea	0x12345678(%r15),%rax
	lea	0x12345678(%r15),%rcx
	lea	0x12345678(%r15),%rdx
	lea	0x12345678(%r15),%rbx
	lea	0x12345678(%r15),%rsp
	lea	0x12345678(%r15),%rbp
	lea	0x12345678(%r15),%rsi
	lea	0x12345678(%r15),%rdi
	lea	0x12345678(%r15),%r8
	lea	0x12345678(%r15),%r9
	lea	0x12345678(%r15),%r10
	lea	0x12345678(%r15),%r11
	lea	0x12345678(%r15),%r12
	lea	0x12345678(%r15),%r13
	lea	0x12345678(%r15),%r14
	lea	0x12345678(%r15),%r15
        nop
        nop
        // reg64 = &mem64[reg64]
	lea	(%rax,%rax,1),%rax
	lea	(%rax,%rax,1),%rcx
	lea	(%rax,%rax,1),%rdx
	lea	(%rax,%rax,1),%rbx
	lea	(%rax,%rax,1),%rsp
	lea	(%rax,%rax,1),%rbp
	lea	(%rax,%rax,1),%rsi
	lea	(%rax,%rax,1),%rdi
	lea	(%rax,%rax,1),%r8
	lea	(%rax,%rax,1),%r9
	lea	(%rax,%rax,1),%r10
	lea	(%rax,%rax,1),%r11
	lea	(%rax,%rax,1),%r12
	lea	(%rax,%rax,1),%r13
	lea	(%rax,%rax,1),%r14
	lea	(%rax,%rax,1),%r15
	nop
	lea	(%rax,%rax,1),%rax
	lea	(%rcx,%rax,1),%rax
	lea	(%rdx,%rax,1),%rax
	lea	(%rbx,%rax,1),%rax
	lea	(%rsp,%rax,1),%rax
	lea	(%rbp,%rax,1),%rax
	lea	(%rsi,%rax,1),%rax
	lea	(%rdi,%rax,1),%rax
	lea	(%r8 ,%rax,1),%rax
	lea	(%r9 ,%rax,1),%rax
	lea	(%r10,%rax,1),%rax
	lea	(%r11,%rax,1),%rax
	lea	(%r12,%rax,1),%rax
	lea	(%r13,%rax,1),%rax
	lea	(%r14,%rax,1),%rax
	lea	(%r15,%rax,1),%rax
	nop
	lea	(%rax,%rax,1),%rax
	lea	(%rax,%rcx,1),%rax
	lea	(%rax,%rdx,1),%rax
	lea	(%rax,%rbx,1),%rax
        // lea  (%rax,%rsp,1),%rax 
	xor     %rsp, %rsp
	lea	(%rax,%rbp,1),%rax
	lea	(%rax,%rsi,1),%rax
	lea	(%rax,%rdi,1),%rax
	lea	(%rax,%r8 ,1),%rax
	lea	(%rax,%r9 ,1),%rax
	lea	(%rax,%r10,1),%rax
	lea	(%rax,%r11,1),%rax
	lea	(%rax,%r12,1),%rax
	lea	(%rax,%r13,1),%rax
	lea	(%rax,%r14,1),%rax
	lea	(%rax,%r15,1),%rax
        nop
        nop
        // reg64 = &mem64[reg64*2]
	lea	(%rax,%rax,2),%rax
	lea	(%rax,%rax,2),%rcx
	lea	(%rax,%rax,2),%rdx
	lea	(%rax,%rax,2),%rbx
	lea	(%rax,%rax,2),%rsp
	lea	(%rax,%rax,2),%rbp
	lea	(%rax,%rax,2),%rsi
	lea	(%rax,%rax,2),%rdi
	lea	(%rax,%rax,2),%r8
	lea	(%rax,%rax,2),%r9
	lea	(%rax,%rax,2),%r10
	lea	(%rax,%rax,2),%r11
	lea	(%rax,%rax,2),%r12
	lea	(%rax,%rax,2),%r13
	lea	(%rax,%rax,2),%r14
	lea	(%rax,%rax,2),%r15
	nop
	lea	(%rax,%rax,2),%rax
	lea	(%rcx,%rax,2),%rax
	lea	(%rdx,%rax,2),%rax
	lea	(%rbx,%rax,2),%rax
	lea	(%rsp,%rax,2),%rax
	lea	(%rbp,%rax,2),%rax
	lea	(%rsi,%rax,2),%rax
	lea	(%rdi,%rax,2),%rax
	lea	(%r8 ,%rax,2),%rax
	lea	(%r9 ,%rax,2),%rax
	lea	(%r10,%rax,2),%rax
	lea	(%r11,%rax,2),%rax
	lea	(%r12,%rax,2),%rax
	lea	(%r13,%rax,2),%rax
	lea	(%r14,%rax,2),%rax
	lea	(%r15,%rax,2),%rax
	nop
	lea	(%rax,%rax,2),%rax
	lea	(%rax,%rcx,2),%rax
	lea	(%rax,%rdx,2),%rax
	lea	(%rax,%rbx,2),%rax
	// lea  (%rax,%rsp,2),%rax
        xor     %rsp, %rsp
	lea	(%rax,%rbp,2),%rax
	lea	(%rax,%rsi,2),%rax
	lea	(%rax,%rdi,2),%rax
	lea	(%rax,%r8 ,2),%rax
	lea	(%rax,%r9 ,2),%rax
	lea	(%rax,%r10,2),%rax
	lea	(%rax,%r11,2),%rax
	lea	(%rax,%r12,2),%rax
	lea	(%rax,%r13,2),%rax
	lea	(%rax,%r14,2),%rax
	lea	(%rax,%r15,2),%rax
        nop
        nop
        // reg64 = &mem64[reg64*4]
	lea	(%rax,%rax,4),%rax
	lea	(%rax,%rax,4),%rcx
	lea	(%rax,%rax,4),%rdx
	lea	(%rax,%rax,4),%rbx
	lea	(%rax,%rax,4),%rsp
	lea	(%rax,%rax,4),%rbp
	lea	(%rax,%rax,4),%rsi
	lea	(%rax,%rax,4),%rdi
	lea	(%rax,%rax,4),%r8
	lea	(%rax,%rax,4),%r9
	lea	(%rax,%rax,4),%r10
	lea	(%rax,%rax,4),%r11
	lea	(%rax,%rax,4),%r12
	lea	(%rax,%rax,4),%r13
	lea	(%rax,%rax,4),%r14
	lea	(%rax,%rax,4),%r15
	nop
	lea	(%rax,%rax,4),%rax
	lea	(%rcx,%rax,4),%rax
	lea	(%rdx,%rax,4),%rax
	lea	(%rbx,%rax,4),%rax
	lea	(%rsp,%rax,4),%rax
	lea	(%rbp,%rax,4),%rax
	lea	(%rsi,%rax,4),%rax
	lea	(%rdi,%rax,4),%rax
	lea	(%r8 ,%rax,4),%rax
	lea	(%r9 ,%rax,4),%rax
	lea	(%r10,%rax,4),%rax
	lea	(%r11,%rax,4),%rax
	lea	(%r12,%rax,4),%rax
	lea	(%r13,%rax,4),%rax
	lea	(%r14,%rax,4),%rax
	lea	(%r15,%rax,4),%rax
	nop
	lea	(%rax,%rax,4),%rax
	lea	(%rax,%rcx,4),%rax
	lea	(%rax,%rdx,4),%rax
	lea	(%rax,%rbx,4),%rax
	// lea  (%rax,%rsp,4),%rax
        xor     %rsp, %rsp
	lea	(%rax,%rbp,4),%rax
	lea	(%rax,%rsi,4),%rax
	lea	(%rax,%rdi,4),%rax
	lea	(%rax,%r8 ,4),%rax
	lea	(%rax,%r9 ,4),%rax
	lea	(%rax,%r10,4),%rax
	lea	(%rax,%r11,4),%rax
	lea	(%rax,%r12,4),%rax
	lea	(%rax,%r13,4),%rax
	lea	(%rax,%r14,4),%rax
	lea	(%rax,%r15,4),%rax
        nop
        nop
        // reg64 = &mem64[reg64*8]
	lea	(%rax,%rax,8),%rax
	lea	(%rax,%rax,8),%rcx
	lea	(%rax,%rax,8),%rdx
	lea	(%rax,%rax,8),%rbx
	lea	(%rax,%rax,8),%rsp
	lea	(%rax,%rax,8),%rbp
	lea	(%rax,%rax,8),%rsi
	lea	(%rax,%rax,8),%rdi
	lea	(%rax,%rax,8),%r8
	lea	(%rax,%rax,8),%r9
	lea	(%rax,%rax,8),%r10
	lea	(%rax,%rax,8),%r11
	lea	(%rax,%rax,8),%r12
	lea	(%rax,%rax,8),%r13
	lea	(%rax,%rax,8),%r14
	lea	(%rax,%rax,8),%r15
	nop
	lea	(%rax,%rax,8),%rax
	lea	(%rcx,%rax,8),%rax
	lea	(%rdx,%rax,8),%rax
	lea	(%rbx,%rax,8),%rax
	lea	(%rsp,%rax,8),%rax
	lea	(%rbp,%rax,8),%rax
	lea	(%rsi,%rax,8),%rax
	lea	(%rdi,%rax,8),%rax
	lea	(%r8 ,%rax,8),%rax
	lea	(%r9 ,%rax,8),%rax
	lea	(%r10,%rax,8),%rax
	lea	(%r11,%rax,8),%rax
	lea	(%r12,%rax,8),%rax
	lea	(%r13,%rax,8),%rax
	lea	(%r14,%rax,8),%rax
	lea	(%r15,%rax,8),%rax
	nop
	lea	(%rax,%rax,8),%rax
	lea	(%rax,%rcx,8),%rax
	lea	(%rax,%rdx,8),%rax
	lea	(%rax,%rbx,8),%rax
	// lea  (%rax,%rsp,8),%rax
        xor     %rsp, %rsp
	lea	(%rax,%rbp,8),%rax
	lea	(%rax,%rsi,8),%rax
	lea	(%rax,%rdi,8),%rax
	lea	(%rax,%r8 ,8),%rax
	lea	(%rax,%r9 ,8),%rax
	lea	(%rax,%r10,8),%rax
	lea	(%rax,%r11,8),%rax
	lea	(%rax,%r12,8),%rax
	lea	(%rax,%r13,8),%rax
	lea	(%rax,%r14,8),%rax
	lea	(%rax,%r15,8),%rax
        nop
        nop
        // reg64 = &[reg64*1]
	lea	(,%rax,1),%rax
	lea	(,%rax,1),%rcx
	lea	(,%rax,1),%rdx
	lea	(,%rax,1),%rbx
	lea	(,%rax,1),%rsp
	lea	(,%rax,1),%rbp
	lea	(,%rax,1),%rsi
	lea	(,%rax,1),%rdi
	lea	(,%rax,1),%r8
	lea	(,%rax,1),%r9
	lea	(,%rax,1),%r10
	lea	(,%rax,1),%r11
	lea	(,%rax,1),%r12
	lea	(,%rax,1),%r13
	lea	(,%rax,1),%r14
	lea	(,%rax,1),%r15
	nop
	lea	(,%rax,1),%rax
	lea	(,%rcx,1),%rax
	lea	(,%rdx,1),%rax
	lea	(,%rbx,1),%rax
	// lea  (,%rsp,1),%rax
        xor     %rsp, %rsp
	lea	(,%rbp,1),%rax
	lea	(,%rsi,1),%rax
	lea	(,%rdi,1),%rax
	lea	(,%r8 ,1),%rax
	lea	(,%r9 ,1),%rax
	lea	(,%r10,1),%rax
	lea	(,%r11,1),%rax
	lea	(,%r12,1),%rax
	lea	(,%r13,1),%rax
	lea	(,%r14,1),%rax
	lea	(,%r15,1),%rax
        nop
        nop
        // reg64 = &[reg64*2]
	lea	(,%rax,2),%rax
	lea	(,%rax,2),%rcx
	lea	(,%rax,2),%rdx
	lea	(,%rax,2),%rbx
	lea	(,%rax,2),%rsp
	lea	(,%rax,2),%rbp
	lea	(,%rax,2),%rsi
	lea	(,%rax,2),%rdi
	lea	(,%rax,2),%r8
	lea	(,%rax,2),%r9
	lea	(,%rax,2),%r10
	lea	(,%rax,2),%r11
	lea	(,%rax,2),%r12
	lea	(,%rax,2),%r13
	lea	(,%rax,2),%r14
	lea	(,%rax,2),%r15
	nop
	lea	(,%rax,2),%rax
	lea	(,%rcx,2),%rax
	lea	(,%rdx,2),%rax
	lea	(,%rbx,2),%rax
	// lea  (,%rsp,2),%rax
        xor     %rsp, %rsp
	lea	(,%rbp,2),%rax
	lea	(,%rsi,2),%rax
	lea	(,%rdi,2),%rax
	lea	(,%r8 ,2),%rax
	lea	(,%r9 ,2),%rax
	lea	(,%r10,2),%rax
	lea	(,%r11,2),%rax
	lea	(,%r12,2),%rax
	lea	(,%r13,2),%rax
	lea	(,%r14,2),%rax
	lea	(,%r15,2),%rax
        nop
        nop
        // reg64 = &[reg64*4]
	lea	(,%rax,4),%rax
	lea	(,%rax,4),%rcx
	lea	(,%rax,4),%rdx
	lea	(,%rax,4),%rbx
	lea	(,%rax,4),%rsp
	lea	(,%rax,4),%rbp
	lea	(,%rax,4),%rsi
	lea	(,%rax,4),%rdi
	lea	(,%rax,4),%r8
	lea	(,%rax,4),%r9
	lea	(,%rax,4),%r10
	lea	(,%rax,4),%r11
	lea	(,%rax,4),%r12
	lea	(,%rax,4),%r13
	lea	(,%rax,4),%r14
	lea	(,%rax,4),%r15
	nop
	lea	(,%rax,4),%rax
	lea	(,%rcx,4),%rax
	lea	(,%rdx,4),%rax
	lea	(,%rbx,4),%rax
	// lea  (,%rsp,4),%rax
        xor     %rsp, %rsp
	lea	(,%rbp,4),%rax
	lea	(,%rsi,4),%rax
	lea	(,%rdi,4),%rax
	lea	(,%r8 ,4),%rax
	lea	(,%r9 ,4),%rax
	lea	(,%r10,4),%rax
	lea	(,%r11,4),%rax
	lea	(,%r12,4),%rax
	lea	(,%r13,4),%rax
	lea	(,%r14,4),%rax
	lea	(,%r15,4),%rax
        nop
        nop
        // reg64 = &[reg64*8]
	lea	(,%rax,8),%rax
	lea	(,%rax,8),%rcx
	lea	(,%rax,8),%rdx
	lea	(,%rax,8),%rbx
	lea	(,%rax,8),%rsp
	lea	(,%rax,8),%rbp
	lea	(,%rax,8),%rsi
	lea	(,%rax,8),%rdi
	lea	(,%rax,8),%r8
	lea	(,%rax,8),%r9
	lea	(,%rax,8),%r10
	lea	(,%rax,8),%r11
	lea	(,%rax,8),%r12
	lea	(,%rax,8),%r13
	lea	(,%rax,8),%r14
	lea	(,%rax,8),%r15
	nop
	lea	(,%rax,8),%rax
	lea	(,%rcx,8),%rax
	lea	(,%rdx,8),%rax
	lea	(,%rbx,8),%rax
	// lea  (,%rsp,8),%rax
        xor     %rsp, %rsp
	lea	(,%rbp,8),%rax
	lea	(,%rsi,8),%rax
	lea	(,%rdi,8),%rax
	lea	(,%r8 ,8),%rax
	lea	(,%r9 ,8),%rax
	lea	(,%r10,8),%rax
	lea	(,%r11,8),%rax
	lea	(,%r12,8),%rax
	lea	(,%r13,8),%rax
	lea	(,%r14,8),%rax
	lea	(,%r15,8),%rax
	ret
	.cfi_endproc


