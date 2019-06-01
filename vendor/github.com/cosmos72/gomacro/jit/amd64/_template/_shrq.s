	.file	"shr.s"
	.text

	.p2align 4,,15
	.globl	ShrByConst
	.type	ShrByConst, @function
ShrByConst:
	.cfi_startproc
	// reg64 <<= const
	shrq	$1, %rax
	shrq	$1, %rcx
	shrq	$1, %rdx
	shrq	$1, %rbx
	shrq	$1, %rsp
	shrq	$1, %rbp
	shrq	$1, %rsi
	shrq	$1, %rdi
	shrq	$1, %r8
	shrq	$1, %r9
	shrq	$1, %r10
	shrq	$1, %r11
	shrq	$1, %r12
	shrq	$1, %r13
	shrq	$1, %r14
	shrq	$1, %r15
        nop
	shrq	$2, %rax
	shrq	$2, %rcx
	shrq	$2, %rdx
	shrq	$2, %rbx
	shrq	$2, %rsp
	shrq	$2, %rbp
	shrq	$2, %rsi
	shrq	$2, %rdi
	shrq	$2, %r8
	shrq	$2, %r9
	shrq	$2, %r10
	shrq	$2, %r11
	shrq	$2, %r12
	shrq	$2, %r13
	shrq	$2, %r14
	shrq	$2, %r15
	nop
	shrq	$7, %rax
	shrq	$7, %rcx
	shrq	$7, %rdx
	shrq	$7, %rbx
	shrq	$7, %rsp
	shrq	$7, %rbp
	shrq	$7, %rsi
	shrq	$7, %rdi
	shrq	$7, %r8
	shrq	$7, %r9
	shrq	$7, %r10
	shrq	$7, %r11
	shrq	$7, %r12
	shrq	$7, %r13
	shrq	$7, %r14
	shrq	$7, %r15
        nop
        // mem64[0] <<= const
	shrq	$1,(%rax)
	shrq	$1,(%rcx)
	shrq	$1,(%rdx)
	shrq	$1,(%rbx)
	shrq	$1,(%rsp)
	shrq	$1,(%rbp)
	shrq	$1,(%rsi)
	shrq	$1,(%rdi)
	shrq	$1,(%r8)
	shrq	$1,(%r9)
	shrq	$1,(%r10)
	shrq	$1,(%r11)
	shrq	$1,(%r12)
	shrq	$1,(%r13)
	shrq	$1,(%r14)
	shrq	$1,(%r15)
        nop
	shrq	$2,(%rax)
	shrq	$2,(%rcx)
	shrq	$2,(%rdx)
	shrq	$2,(%rbx)
	shrq	$2,(%rsp)
	shrq	$2,(%rbp)
	shrq	$2,(%rsi)
	shrq	$2,(%rdi)
	shrq	$2,(%r8)
	shrq	$2,(%r9)
	shrq	$2,(%r10)
	shrq	$2,(%r11)
	shrq	$2,(%r12)
	shrq	$2,(%r13)
	shrq	$2,(%r14)
	shrq	$2,(%r15)
        nop
	shrq	$7,(%rax)
	shrq	$7,(%rcx)
	shrq	$7,(%rdx)
	shrq	$7,(%rbx)
	shrq	$7,(%rsp)
	shrq	$7,(%rbp)
	shrq	$7,(%rsi)
	shrq	$7,(%rdi)
	shrq	$7,(%r8)
	shrq	$7,(%r9)
	shrq	$7,(%r10)
	shrq	$7,(%r11)
	shrq	$7,(%r12)
	shrq	$7,(%r13)
	shrq	$7,(%r14)
	shrq	$7,(%r15)
        nop
        // mem64[off8] <<= const
	shrq	$1, 0x7F(%rax)
	shrq	$1, 0x7F(%rdx)
	shrq	$1, 0x7F(%rbx)
	shrq	$1, 0x7F(%rsp)
	shrq	$1, 0x7F(%rbp)
	shrq	$1, 0x7F(%rsi)
	shrq	$1, 0x7F(%rdi)
	shrq	$1, 0x7F(%r8)
	shrq	$1, 0x7F(%r9)
	shrq	$1, 0x7F(%r10)
	shrq	$1, 0x7F(%r11)
	shrq	$1, 0x7F(%r12)
	shrq	$1, 0x7F(%r13)
	shrq	$1, 0x7F(%r14)
	shrq	$1, 0x7F(%r15)
        nop
	shrq	$2, 0x7F(%rax)
	shrq	$2, 0x7F(%rdx)
	shrq	$2, 0x7F(%rbx)
	shrq	$2, 0x7F(%rsp)
	shrq	$2, 0x7F(%rbp)
	shrq	$2, 0x7F(%rsi)
	shrq	$2, 0x7F(%rdi)
	shrq	$2, 0x7F(%r8)
	shrq	$2, 0x7F(%r9)
	shrq	$2, 0x7F(%r10)
	shrq	$2, 0x7F(%r11)
	shrq	$2, 0x7F(%r12)
	shrq	$2, 0x7F(%r13)
	shrq	$2, 0x7F(%r14)
	shrq	$2, 0x7F(%r15)
        nop
	shrq	$7, 0x7F(%rax)
	shrq	$7, 0x7F(%rdx)
	shrq	$7, 0x7F(%rbx)
	shrq	$7, 0x7F(%rsp)
	shrq	$7, 0x7F(%rbp)
	shrq	$7, 0x7F(%rsi)
	shrq	$7, 0x7F(%rdi)
	shrq	$7, 0x7F(%r8)
	shrq	$7, 0x7F(%r9)
	shrq	$7, 0x7F(%r10)
	shrq	$7, 0x7F(%r11)
	shrq	$7, 0x7F(%r12)
	shrq	$7, 0x7F(%r13)
	shrq	$7, 0x7F(%r14)
	shrq	$7, 0x7F(%r15)
        nop
        // mem64[off32] <<= const
	shrq	$1, 0x12345678(%rax)
	shrq	$1, 0x12345678(%rdx)
	shrq	$1, 0x12345678(%rbx)
	shrq	$1, 0x12345678(%rsp)
	shrq	$1, 0x12345678(%rbp)
	shrq	$1, 0x12345678(%rsi)
	shrq	$1, 0x12345678(%rdi)
	shrq	$1, 0x12345678(%r8)
	shrq	$1, 0x12345678(%r9)
	shrq	$1, 0x12345678(%r10)
	shrq	$1, 0x12345678(%r11)
	shrq	$1, 0x12345678(%r12)
	shrq	$1, 0x12345678(%r13)
	shrq	$1, 0x12345678(%r14)
	shrq	$1, 0x12345678(%r15)
        nop
	shrq	$2, 0x12345678(%rax)
	shrq	$2, 0x12345678(%rdx)
	shrq	$2, 0x12345678(%rbx)
	shrq	$2, 0x12345678(%rsp)
	shrq	$2, 0x12345678(%rbp)
	shrq	$2, 0x12345678(%rsi)
	shrq	$2, 0x12345678(%rdi)
	shrq	$2, 0x12345678(%r8)
	shrq	$2, 0x12345678(%r9)
	shrq	$2, 0x12345678(%r10)
	shrq	$2, 0x12345678(%r11)
	shrq	$2, 0x12345678(%r12)
	shrq	$2, 0x12345678(%r13)
	shrq	$2, 0x12345678(%r14)
	shrq	$2, 0x12345678(%r15)
        nop
	shrq	$7, 0x12345678(%rax)
	shrq	$7, 0x12345678(%rdx)
	shrq	$7, 0x12345678(%rbx)
	shrq	$7, 0x12345678(%rsp)
	shrq	$7, 0x12345678(%rbp)
	shrq	$7, 0x12345678(%rsi)
	shrq	$7, 0x12345678(%rdi)
	shrq	$7, 0x12345678(%r8)
	shrq	$7, 0x12345678(%r9)
	shrq	$7, 0x12345678(%r10)
	shrq	$7, 0x12345678(%r11)
	shrq	$7, 0x12345678(%r12)
	shrq	$7, 0x12345678(%r13)
	shrq	$7, 0x12345678(%r14)
	shrq	$7, 0x12345678(%r15)
        nop
        nop
        ret
        ret
	.cfi_endproc



	.p2align 4,,15
	.globl	ShrByCl
	.type	ShrByCl, @function
ShrByCl:
	.cfi_startproc
	// reg64 <<= %cl
	shrq	%cl,%rax
	shrq	%cl,%rcx
	shrq	%cl,%rdx
	shrq	%cl,%rbx
	shrq	%cl,%rsp
	shrq	%cl,%rbp
	shrq	%cl,%rsi
	shrq	%cl,%rdi
	shrq	%cl,%r8
	shrq	%cl,%r9
	shrq	%cl,%r10
	shrq	%cl,%r11
	shrq	%cl,%r12
	shrq	%cl,%r13
	shrq	%cl,%r14
	shrq	%cl,%r15
        nop
	// mem64[0] <<= %cl
	shrq	%cl,(%rax)
	shrq	%cl,(%rdx)
	shrq	%cl,(%rbx)
	shrq	%cl,(%rsp)
	shrq	%cl,(%rbp)
	shrq	%cl,(%rsi)
	shrq	%cl,(%rdi)
	shrq	%cl,(%r8)
	shrq	%cl,(%r9)
	shrq	%cl,(%r10)
	shrq	%cl,(%r11)
	shrq	%cl,(%r12)
	shrq	%cl,(%r13)
	shrq	%cl,(%r14)
	shrq	%cl,(%r15)
        nop
	// mem64[off8] <<= %cl
	shrq	%cl, 0x7F(%rax)
	shrq	%cl, 0x7F(%rdx)
	shrq	%cl, 0x7F(%rbx)
	shrq	%cl, 0x7F(%rsp)
	shrq	%cl, 0x7F(%rbp)
	shrq	%cl, 0x7F(%rsi)
	shrq	%cl, 0x7F(%rdi)
	shrq	%cl, 0x7F(%r8)
	shrq	%cl, 0x7F(%r9)
	shrq	%cl, 0x7F(%r10)
	shrq	%cl, 0x7F(%r11)
	shrq	%cl, 0x7F(%r12)
	shrq	%cl, 0x7F(%r13)
	shrq	%cl, 0x7F(%r14)
	shrq	%cl, 0x7F(%r15)
        nop
	// mem64[off32] <<= %cl
	shrq	%cl, 0x12345678(%rax)
	shrq	%cl, 0x12345678(%rdx)
	shrq	%cl, 0x12345678(%rbx)
	shrq	%cl, 0x12345678(%rsp)
	shrq	%cl, 0x12345678(%rbp)
	shrq	%cl, 0x12345678(%rsi)
	shrq	%cl, 0x12345678(%rdi)
	shrq	%cl, 0x12345678(%r8)
	shrq	%cl, 0x12345678(%r9)
	shrq	%cl, 0x12345678(%r10)
	shrq	%cl, 0x12345678(%r11)
	shrq	%cl, 0x12345678(%r12)
	shrq	%cl, 0x12345678(%r13)
	shrq	%cl, 0x12345678(%r14)
	shrq	%cl, 0x12345678(%r15)
        nop
        ret
	.cfi_endproc
