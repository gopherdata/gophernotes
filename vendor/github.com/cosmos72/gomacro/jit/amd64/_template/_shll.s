	.file	"shl.s"
	.text

	.p2align 4,,15
	.globl	ShlByConst
	.type	ShlByConst, @function
ShlByConst:
	.cfi_startproc
	// reg16 <<= const
	shll	$1, %eax
	shll	$1, %ecx
	shll	$1, %edx
	shll	$1, %ebx
	shll	$1, %esp
	shll	$1, %ebp
	shll	$1, %esi
	shll	$1, %edi
	shll	$1, %r8d
	shll	$1, %r9d
	shll	$1, %r10d
	shll	$1, %r11d
	shll	$1, %r12d
	shll	$1, %r13d
	shll	$1, %r14d
	shll	$1, %r15d
        nop
	shll	$2, %eax
	shll	$2, %ecx
	shll	$2, %edx
	shll	$2, %ebx
	shll	$2, %esp
	shll	$2, %ebp
	shll	$2, %esi
	shll	$2, %edi
	shll	$2, %r8d
	shll	$2, %r9d
	shll	$2, %r10d
	shll	$2, %r11d
	shll	$2, %r12d
	shll	$2, %r13d
	shll	$2, %r14d
	shll	$2, %r15d
	nop
	shll	$7, %eax
	shll	$7, %ecx
	shll	$7, %edx
	shll	$7, %ebx
	shll	$7, %esp
	shll	$7, %ebp
	shll	$7, %esi
	shll	$7, %edi
	shll	$7, %r8d
	shll	$7, %r9d
	shll	$7, %r10d
	shll	$7, %r11d
	shll	$7, %r12d
	shll	$7, %r13d
	shll	$7, %r14d
	shll	$7, %r15d
        nop
        // mem16[0] <<= const
	shll	$1,(%rax)
	shll	$1,(%rcx)
	shll	$1,(%rdx)
	shll	$1,(%rbx)
	shll	$1,(%rsp)
	shll	$1,(%rbp)
	shll	$1,(%rsi)
	shll	$1,(%rdi)
	shll	$1,(%r8)
	shll	$1,(%r9)
	shll	$1,(%r10)
	shll	$1,(%r11)
	shll	$1,(%r12)
	shll	$1,(%r13)
	shll	$1,(%r14)
	shll	$1,(%r15)
        nop
	shll	$2,(%rax)
	shll	$2,(%rcx)
	shll	$2,(%rdx)
	shll	$2,(%rbx)
	shll	$2,(%rsp)
	shll	$2,(%rbp)
	shll	$2,(%rsi)
	shll	$2,(%rdi)
	shll	$2,(%r8)
	shll	$2,(%r9)
	shll	$2,(%r10)
	shll	$2,(%r11)
	shll	$2,(%r12)
	shll	$2,(%r13)
	shll	$2,(%r14)
	shll	$2,(%r15)
        nop
	shll	$7,(%rax)
	shll	$7,(%rcx)
	shll	$7,(%rdx)
	shll	$7,(%rbx)
	shll	$7,(%rsp)
	shll	$7,(%rbp)
	shll	$7,(%rsi)
	shll	$7,(%rdi)
	shll	$7,(%r8)
	shll	$7,(%r9)
	shll	$7,(%r10)
	shll	$7,(%r11)
	shll	$7,(%r12)
	shll	$7,(%r13)
	shll	$7,(%r14)
	shll	$7,(%r15)
        nop
        // mem16[off8] <<= const
	shll	$1, 0x7F(%rax)
	shll	$1, 0x7F(%rdx)
	shll	$1, 0x7F(%rbx)
	shll	$1, 0x7F(%rsp)
	shll	$1, 0x7F(%rbp)
	shll	$1, 0x7F(%rsi)
	shll	$1, 0x7F(%rdi)
	shll	$1, 0x7F(%r8)
	shll	$1, 0x7F(%r9)
	shll	$1, 0x7F(%r10)
	shll	$1, 0x7F(%r11)
	shll	$1, 0x7F(%r12)
	shll	$1, 0x7F(%r13)
	shll	$1, 0x7F(%r14)
	shll	$1, 0x7F(%r15)
        nop
	shll	$2, 0x7F(%rax)
	shll	$2, 0x7F(%rdx)
	shll	$2, 0x7F(%rbx)
	shll	$2, 0x7F(%rsp)
	shll	$2, 0x7F(%rbp)
	shll	$2, 0x7F(%rsi)
	shll	$2, 0x7F(%rdi)
	shll	$2, 0x7F(%r8)
	shll	$2, 0x7F(%r9)
	shll	$2, 0x7F(%r10)
	shll	$2, 0x7F(%r11)
	shll	$2, 0x7F(%r12)
	shll	$2, 0x7F(%r13)
	shll	$2, 0x7F(%r14)
	shll	$2, 0x7F(%r15)
        nop
	shll	$7, 0x7F(%rax)
	shll	$7, 0x7F(%rdx)
	shll	$7, 0x7F(%rbx)
	shll	$7, 0x7F(%rsp)
	shll	$7, 0x7F(%rbp)
	shll	$7, 0x7F(%rsi)
	shll	$7, 0x7F(%rdi)
	shll	$7, 0x7F(%r8)
	shll	$7, 0x7F(%r9)
	shll	$7, 0x7F(%r10)
	shll	$7, 0x7F(%r11)
	shll	$7, 0x7F(%r12)
	shll	$7, 0x7F(%r13)
	shll	$7, 0x7F(%r14)
	shll	$7, 0x7F(%r15)
        nop
        // mem16[off32] <<= const
	shll	$1, 0x12345678(%rax)
	shll	$1, 0x12345678(%rdx)
	shll	$1, 0x12345678(%rbx)
	shll	$1, 0x12345678(%rsp)
	shll	$1, 0x12345678(%rbp)
	shll	$1, 0x12345678(%rsi)
	shll	$1, 0x12345678(%rdi)
	shll	$1, 0x12345678(%r8)
	shll	$1, 0x12345678(%r9)
	shll	$1, 0x12345678(%r10)
	shll	$1, 0x12345678(%r11)
	shll	$1, 0x12345678(%r12)
	shll	$1, 0x12345678(%r13)
	shll	$1, 0x12345678(%r14)
	shll	$1, 0x12345678(%r15)
        nop
	shll	$2, 0x12345678(%rax)
	shll	$2, 0x12345678(%rdx)
	shll	$2, 0x12345678(%rbx)
	shll	$2, 0x12345678(%rsp)
	shll	$2, 0x12345678(%rbp)
	shll	$2, 0x12345678(%rsi)
	shll	$2, 0x12345678(%rdi)
	shll	$2, 0x12345678(%r8)
	shll	$2, 0x12345678(%r9)
	shll	$2, 0x12345678(%r10)
	shll	$2, 0x12345678(%r11)
	shll	$2, 0x12345678(%r12)
	shll	$2, 0x12345678(%r13)
	shll	$2, 0x12345678(%r14)
	shll	$2, 0x12345678(%r15)
        nop
	shll	$7, 0x12345678(%rax)
	shll	$7, 0x12345678(%rdx)
	shll	$7, 0x12345678(%rbx)
	shll	$7, 0x12345678(%rsp)
	shll	$7, 0x12345678(%rbp)
	shll	$7, 0x12345678(%rsi)
	shll	$7, 0x12345678(%rdi)
	shll	$7, 0x12345678(%r8)
	shll	$7, 0x12345678(%r9)
	shll	$7, 0x12345678(%r10)
	shll	$7, 0x12345678(%r11)
	shll	$7, 0x12345678(%r12)
	shll	$7, 0x12345678(%r13)
	shll	$7, 0x12345678(%r14)
	shll	$7, 0x12345678(%r15)
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
	// reg16 <<= %cl
	shll	%cl,%eax
	shll	%cl,%ecx
	shll	%cl,%edx
	shll	%cl,%ebx
	shll	%cl,%esp
	shll	%cl,%ebp
	shll	%cl,%esi
	shll	%cl,%edi
	shll	%cl,%r8d
	shll	%cl,%r9d
	shll	%cl,%r10d
	shll	%cl,%r11d
	shll	%cl,%r12d
	shll	%cl,%r13d
	shll	%cl,%r14d
	shll	%cl,%r15d
        nop
	// mem16[0] <<= %cl
	shll	%cl,(%rax)
	shll	%cl,(%rdx)
	shll	%cl,(%rbx)
	shll	%cl,(%rsp)
	shll	%cl,(%rbp)
	shll	%cl,(%rsi)
	shll	%cl,(%rdi)
	shll	%cl,(%r8)
	shll	%cl,(%r9)
	shll	%cl,(%r10)
	shll	%cl,(%r11)
	shll	%cl,(%r12)
	shll	%cl,(%r13)
	shll	%cl,(%r14)
	shll	%cl,(%r15)
        nop
	// mem16[off8] <<= %cl
	shll	%cl, 0x7F(%rax)
	shll	%cl, 0x7F(%rdx)
	shll	%cl, 0x7F(%rbx)
	shll	%cl, 0x7F(%rsp)
	shll	%cl, 0x7F(%rbp)
	shll	%cl, 0x7F(%rsi)
	shll	%cl, 0x7F(%rdi)
	shll	%cl, 0x7F(%r8)
	shll	%cl, 0x7F(%r9)
	shll	%cl, 0x7F(%r10)
	shll	%cl, 0x7F(%r11)
	shll	%cl, 0x7F(%r12)
	shll	%cl, 0x7F(%r13)
	shll	%cl, 0x7F(%r14)
	shll	%cl, 0x7F(%r15)
        nop
	// mem16[off32] <<= %cl
	shll	%cl, 0x12345678(%rax)
	shll	%cl, 0x12345678(%rdx)
	shll	%cl, 0x12345678(%rbx)
	shll	%cl, 0x12345678(%rsp)
	shll	%cl, 0x12345678(%rbp)
	shll	%cl, 0x12345678(%rsi)
	shll	%cl, 0x12345678(%rdi)
	shll	%cl, 0x12345678(%r8)
	shll	%cl, 0x12345678(%r9)
	shll	%cl, 0x12345678(%r10)
	shll	%cl, 0x12345678(%r11)
	shll	%cl, 0x12345678(%r12)
	shll	%cl, 0x12345678(%r13)
	shll	%cl, 0x12345678(%r14)
	shll	%cl, 0x12345678(%r15)
        nop
        ret
	.cfi_endproc
