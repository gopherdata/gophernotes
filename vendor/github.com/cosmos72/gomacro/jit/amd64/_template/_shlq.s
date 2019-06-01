	.file	"shl.s"
	.text

	.p2align 4,,15
	.globl	ShlByConst
	.type	ShlByConst, @function
ShlByConst:
	.cfi_startproc
	// reg64 <<= const
	shlq	$1, %rax
	shlq	$1, %rcx
	shlq	$1, %rdx
	shlq	$1, %rbx
	shlq	$1, %rsp
	shlq	$1, %rbp
	shlq	$1, %rsi
	shlq	$1, %rdi
	shlq	$1, %r8
	shlq	$1, %r9
	shlq	$1, %r10
	shlq	$1, %r11
	shlq	$1, %r12
	shlq	$1, %r13
	shlq	$1, %r14
	shlq	$1, %r15
        nop
	shlq	$2, %rax
	shlq	$2, %rcx
	shlq	$2, %rdx
	shlq	$2, %rbx
	shlq	$2, %rsp
	shlq	$2, %rbp
	shlq	$2, %rsi
	shlq	$2, %rdi
	shlq	$2, %r8
	shlq	$2, %r9
	shlq	$2, %r10
	shlq	$2, %r11
	shlq	$2, %r12
	shlq	$2, %r13
	shlq	$2, %r14
	shlq	$2, %r15
	nop
	shlq	$7, %rax
	shlq	$7, %rcx
	shlq	$7, %rdx
	shlq	$7, %rbx
	shlq	$7, %rsp
	shlq	$7, %rbp
	shlq	$7, %rsi
	shlq	$7, %rdi
	shlq	$7, %r8
	shlq	$7, %r9
	shlq	$7, %r10
	shlq	$7, %r11
	shlq	$7, %r12
	shlq	$7, %r13
	shlq	$7, %r14
	shlq	$7, %r15
        nop
        // mem64[0] <<= const
	shlq	$1,(%rax)
	shlq	$1,(%rcx)
	shlq	$1,(%rdx)
	shlq	$1,(%rbx)
	shlq	$1,(%rsp)
	shlq	$1,(%rbp)
	shlq	$1,(%rsi)
	shlq	$1,(%rdi)
	shlq	$1,(%r8)
	shlq	$1,(%r9)
	shlq	$1,(%r10)
	shlq	$1,(%r11)
	shlq	$1,(%r12)
	shlq	$1,(%r13)
	shlq	$1,(%r14)
	shlq	$1,(%r15)
        nop
	shlq	$2,(%rax)
	shlq	$2,(%rcx)
	shlq	$2,(%rdx)
	shlq	$2,(%rbx)
	shlq	$2,(%rsp)
	shlq	$2,(%rbp)
	shlq	$2,(%rsi)
	shlq	$2,(%rdi)
	shlq	$2,(%r8)
	shlq	$2,(%r9)
	shlq	$2,(%r10)
	shlq	$2,(%r11)
	shlq	$2,(%r12)
	shlq	$2,(%r13)
	shlq	$2,(%r14)
	shlq	$2,(%r15)
        nop
	shlq	$7,(%rax)
	shlq	$7,(%rcx)
	shlq	$7,(%rdx)
	shlq	$7,(%rbx)
	shlq	$7,(%rsp)
	shlq	$7,(%rbp)
	shlq	$7,(%rsi)
	shlq	$7,(%rdi)
	shlq	$7,(%r8)
	shlq	$7,(%r9)
	shlq	$7,(%r10)
	shlq	$7,(%r11)
	shlq	$7,(%r12)
	shlq	$7,(%r13)
	shlq	$7,(%r14)
	shlq	$7,(%r15)
        nop
        // mem64[off8] <<= const
	shlq	$1, 0x7F(%rax)
	shlq	$1, 0x7F(%rdx)
	shlq	$1, 0x7F(%rbx)
	shlq	$1, 0x7F(%rsp)
	shlq	$1, 0x7F(%rbp)
	shlq	$1, 0x7F(%rsi)
	shlq	$1, 0x7F(%rdi)
	shlq	$1, 0x7F(%r8)
	shlq	$1, 0x7F(%r9)
	shlq	$1, 0x7F(%r10)
	shlq	$1, 0x7F(%r11)
	shlq	$1, 0x7F(%r12)
	shlq	$1, 0x7F(%r13)
	shlq	$1, 0x7F(%r14)
	shlq	$1, 0x7F(%r15)
        nop
	shlq	$2, 0x7F(%rax)
	shlq	$2, 0x7F(%rdx)
	shlq	$2, 0x7F(%rbx)
	shlq	$2, 0x7F(%rsp)
	shlq	$2, 0x7F(%rbp)
	shlq	$2, 0x7F(%rsi)
	shlq	$2, 0x7F(%rdi)
	shlq	$2, 0x7F(%r8)
	shlq	$2, 0x7F(%r9)
	shlq	$2, 0x7F(%r10)
	shlq	$2, 0x7F(%r11)
	shlq	$2, 0x7F(%r12)
	shlq	$2, 0x7F(%r13)
	shlq	$2, 0x7F(%r14)
	shlq	$2, 0x7F(%r15)
        nop
	shlq	$7, 0x7F(%rax)
	shlq	$7, 0x7F(%rdx)
	shlq	$7, 0x7F(%rbx)
	shlq	$7, 0x7F(%rsp)
	shlq	$7, 0x7F(%rbp)
	shlq	$7, 0x7F(%rsi)
	shlq	$7, 0x7F(%rdi)
	shlq	$7, 0x7F(%r8)
	shlq	$7, 0x7F(%r9)
	shlq	$7, 0x7F(%r10)
	shlq	$7, 0x7F(%r11)
	shlq	$7, 0x7F(%r12)
	shlq	$7, 0x7F(%r13)
	shlq	$7, 0x7F(%r14)
	shlq	$7, 0x7F(%r15)
        nop
        // mem64[off32] <<= const
	shlq	$1, 0x12345678(%rax)
	shlq	$1, 0x12345678(%rdx)
	shlq	$1, 0x12345678(%rbx)
	shlq	$1, 0x12345678(%rsp)
	shlq	$1, 0x12345678(%rbp)
	shlq	$1, 0x12345678(%rsi)
	shlq	$1, 0x12345678(%rdi)
	shlq	$1, 0x12345678(%r8)
	shlq	$1, 0x12345678(%r9)
	shlq	$1, 0x12345678(%r10)
	shlq	$1, 0x12345678(%r11)
	shlq	$1, 0x12345678(%r12)
	shlq	$1, 0x12345678(%r13)
	shlq	$1, 0x12345678(%r14)
	shlq	$1, 0x12345678(%r15)
        nop
	shlq	$2, 0x12345678(%rax)
	shlq	$2, 0x12345678(%rdx)
	shlq	$2, 0x12345678(%rbx)
	shlq	$2, 0x12345678(%rsp)
	shlq	$2, 0x12345678(%rbp)
	shlq	$2, 0x12345678(%rsi)
	shlq	$2, 0x12345678(%rdi)
	shlq	$2, 0x12345678(%r8)
	shlq	$2, 0x12345678(%r9)
	shlq	$2, 0x12345678(%r10)
	shlq	$2, 0x12345678(%r11)
	shlq	$2, 0x12345678(%r12)
	shlq	$2, 0x12345678(%r13)
	shlq	$2, 0x12345678(%r14)
	shlq	$2, 0x12345678(%r15)
        nop
	shlq	$7, 0x12345678(%rax)
	shlq	$7, 0x12345678(%rdx)
	shlq	$7, 0x12345678(%rbx)
	shlq	$7, 0x12345678(%rsp)
	shlq	$7, 0x12345678(%rbp)
	shlq	$7, 0x12345678(%rsi)
	shlq	$7, 0x12345678(%rdi)
	shlq	$7, 0x12345678(%r8)
	shlq	$7, 0x12345678(%r9)
	shlq	$7, 0x12345678(%r10)
	shlq	$7, 0x12345678(%r11)
	shlq	$7, 0x12345678(%r12)
	shlq	$7, 0x12345678(%r13)
	shlq	$7, 0x12345678(%r14)
	shlq	$7, 0x12345678(%r15)
        nop
        nop
        ret
        ret
	.cfi_endproc



	.p2align 4,,15
	.globl	ShlByCl
	.type	ShlByCl, @function
ShlByCl:
	.cfi_startproc
	// reg64 <<= %cl
	shlq	%cl,%rax
	shlq	%cl,%rcx
	shlq	%cl,%rdx
	shlq	%cl,%rbx
	shlq	%cl,%rsp
	shlq	%cl,%rbp
	shlq	%cl,%rsi
	shlq	%cl,%rdi
	shlq	%cl,%r8
	shlq	%cl,%r9
	shlq	%cl,%r10
	shlq	%cl,%r11
	shlq	%cl,%r12
	shlq	%cl,%r13
	shlq	%cl,%r14
	shlq	%cl,%r15
        nop
	// mem64[0] <<= %cl
	shlq	%cl,(%rax)
	shlq	%cl,(%rdx)
	shlq	%cl,(%rbx)
	shlq	%cl,(%rsp)
	shlq	%cl,(%rbp)
	shlq	%cl,(%rsi)
	shlq	%cl,(%rdi)
	shlq	%cl,(%r8)
	shlq	%cl,(%r9)
	shlq	%cl,(%r10)
	shlq	%cl,(%r11)
	shlq	%cl,(%r12)
	shlq	%cl,(%r13)
	shlq	%cl,(%r14)
	shlq	%cl,(%r15)
        nop
	// mem64[off8] <<= %cl
	shlq	%cl, 0x7F(%rax)
	shlq	%cl, 0x7F(%rdx)
	shlq	%cl, 0x7F(%rbx)
	shlq	%cl, 0x7F(%rsp)
	shlq	%cl, 0x7F(%rbp)
	shlq	%cl, 0x7F(%rsi)
	shlq	%cl, 0x7F(%rdi)
	shlq	%cl, 0x7F(%r8)
	shlq	%cl, 0x7F(%r9)
	shlq	%cl, 0x7F(%r10)
	shlq	%cl, 0x7F(%r11)
	shlq	%cl, 0x7F(%r12)
	shlq	%cl, 0x7F(%r13)
	shlq	%cl, 0x7F(%r14)
	shlq	%cl, 0x7F(%r15)
        nop
	// mem64[off32] <<= %cl
	shlq	%cl, 0x12345678(%rax)
	shlq	%cl, 0x12345678(%rdx)
	shlq	%cl, 0x12345678(%rbx)
	shlq	%cl, 0x12345678(%rsp)
	shlq	%cl, 0x12345678(%rbp)
	shlq	%cl, 0x12345678(%rsi)
	shlq	%cl, 0x12345678(%rdi)
	shlq	%cl, 0x12345678(%r8)
	shlq	%cl, 0x12345678(%r9)
	shlq	%cl, 0x12345678(%r10)
	shlq	%cl, 0x12345678(%r11)
	shlq	%cl, 0x12345678(%r12)
	shlq	%cl, 0x12345678(%r13)
	shlq	%cl, 0x12345678(%r14)
	shlq	%cl, 0x12345678(%r15)
        nop
        ret
	.cfi_endproc
