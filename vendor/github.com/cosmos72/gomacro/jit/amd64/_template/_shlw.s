	.file	"shl.s"
	.text

	.p2align 4,,15
	.globl	ShlByConst
	.type	ShlByConst, @function
ShlByConst:
	.cfi_startproc
	// reg16 <<= const
	shlw	$1, %ax
	shlw	$1, %cx
	shlw	$1, %dx
	shlw	$1, %bx
	shlw	$1, %sp
	shlw	$1, %bp
	shlw	$1, %si
	shlw	$1, %di
	shlw	$1, %r8w
	shlw	$1, %r9w
	shlw	$1, %r10w
	shlw	$1, %r11w
	shlw	$1, %r12w
	shlw	$1, %r13w
	shlw	$1, %r14w
	shlw	$1, %r15w
        nop
	shlw	$2, %ax
	shlw	$2, %cx
	shlw	$2, %dx
	shlw	$2, %bx
	shlw	$2, %sp
	shlw	$2, %bp
	shlw	$2, %si
	shlw	$2, %di
	shlw	$2, %r8w
	shlw	$2, %r9w
	shlw	$2, %r10w
	shlw	$2, %r11w
	shlw	$2, %r12w
	shlw	$2, %r13w
	shlw	$2, %r14w
	shlw	$2, %r15w
	nop
	shlw	$7, %ax
	shlw	$7, %cx
	shlw	$7, %dx
	shlw	$7, %bx
	shlw	$7, %sp
	shlw	$7, %bp
	shlw	$7, %si
	shlw	$7, %di
	shlw	$7, %r8w
	shlw	$7, %r9w
	shlw	$7, %r10w
	shlw	$7, %r11w
	shlw	$7, %r12w
	shlw	$7, %r13w
	shlw	$7, %r14w
	shlw	$7, %r15w
        nop
        // mem16[0] <<= const
	shlw	$1,(%rax)
	shlw	$1,(%rcx)
	shlw	$1,(%rdx)
	shlw	$1,(%rbx)
	shlw	$1,(%rsp)
	shlw	$1,(%rbp)
	shlw	$1,(%rsi)
	shlw	$1,(%rdi)
	shlw	$1,(%r8)
	shlw	$1,(%r9)
	shlw	$1,(%r10)
	shlw	$1,(%r11)
	shlw	$1,(%r12)
	shlw	$1,(%r13)
	shlw	$1,(%r14)
	shlw	$1,(%r15)
        nop
	shlw	$2,(%rax)
	shlw	$2,(%rcx)
	shlw	$2,(%rdx)
	shlw	$2,(%rbx)
	shlw	$2,(%rsp)
	shlw	$2,(%rbp)
	shlw	$2,(%rsi)
	shlw	$2,(%rdi)
	shlw	$2,(%r8)
	shlw	$2,(%r9)
	shlw	$2,(%r10)
	shlw	$2,(%r11)
	shlw	$2,(%r12)
	shlw	$2,(%r13)
	shlw	$2,(%r14)
	shlw	$2,(%r15)
        nop
	shlw	$7,(%rax)
	shlw	$7,(%rcx)
	shlw	$7,(%rdx)
	shlw	$7,(%rbx)
	shlw	$7,(%rsp)
	shlw	$7,(%rbp)
	shlw	$7,(%rsi)
	shlw	$7,(%rdi)
	shlw	$7,(%r8)
	shlw	$7,(%r9)
	shlw	$7,(%r10)
	shlw	$7,(%r11)
	shlw	$7,(%r12)
	shlw	$7,(%r13)
	shlw	$7,(%r14)
	shlw	$7,(%r15)
        nop
        // mem16[off8] <<= const
	shlw	$1, 0x7F(%rax)
	shlw	$1, 0x7F(%rdx)
	shlw	$1, 0x7F(%rbx)
	shlw	$1, 0x7F(%rsp)
	shlw	$1, 0x7F(%rbp)
	shlw	$1, 0x7F(%rsi)
	shlw	$1, 0x7F(%rdi)
	shlw	$1, 0x7F(%r8)
	shlw	$1, 0x7F(%r9)
	shlw	$1, 0x7F(%r10)
	shlw	$1, 0x7F(%r11)
	shlw	$1, 0x7F(%r12)
	shlw	$1, 0x7F(%r13)
	shlw	$1, 0x7F(%r14)
	shlw	$1, 0x7F(%r15)
        nop
	shlw	$2, 0x7F(%rax)
	shlw	$2, 0x7F(%rdx)
	shlw	$2, 0x7F(%rbx)
	shlw	$2, 0x7F(%rsp)
	shlw	$2, 0x7F(%rbp)
	shlw	$2, 0x7F(%rsi)
	shlw	$2, 0x7F(%rdi)
	shlw	$2, 0x7F(%r8)
	shlw	$2, 0x7F(%r9)
	shlw	$2, 0x7F(%r10)
	shlw	$2, 0x7F(%r11)
	shlw	$2, 0x7F(%r12)
	shlw	$2, 0x7F(%r13)
	shlw	$2, 0x7F(%r14)
	shlw	$2, 0x7F(%r15)
        nop
	shlw	$7, 0x7F(%rax)
	shlw	$7, 0x7F(%rdx)
	shlw	$7, 0x7F(%rbx)
	shlw	$7, 0x7F(%rsp)
	shlw	$7, 0x7F(%rbp)
	shlw	$7, 0x7F(%rsi)
	shlw	$7, 0x7F(%rdi)
	shlw	$7, 0x7F(%r8)
	shlw	$7, 0x7F(%r9)
	shlw	$7, 0x7F(%r10)
	shlw	$7, 0x7F(%r11)
	shlw	$7, 0x7F(%r12)
	shlw	$7, 0x7F(%r13)
	shlw	$7, 0x7F(%r14)
	shlw	$7, 0x7F(%r15)
        nop
        // mem16[off32] <<= const
	shlw	$1, 0x12345678(%rax)
	shlw	$1, 0x12345678(%rdx)
	shlw	$1, 0x12345678(%rbx)
	shlw	$1, 0x12345678(%rsp)
	shlw	$1, 0x12345678(%rbp)
	shlw	$1, 0x12345678(%rsi)
	shlw	$1, 0x12345678(%rdi)
	shlw	$1, 0x12345678(%r8)
	shlw	$1, 0x12345678(%r9)
	shlw	$1, 0x12345678(%r10)
	shlw	$1, 0x12345678(%r11)
	shlw	$1, 0x12345678(%r12)
	shlw	$1, 0x12345678(%r13)
	shlw	$1, 0x12345678(%r14)
	shlw	$1, 0x12345678(%r15)
        nop
	shlw	$2, 0x12345678(%rax)
	shlw	$2, 0x12345678(%rdx)
	shlw	$2, 0x12345678(%rbx)
	shlw	$2, 0x12345678(%rsp)
	shlw	$2, 0x12345678(%rbp)
	shlw	$2, 0x12345678(%rsi)
	shlw	$2, 0x12345678(%rdi)
	shlw	$2, 0x12345678(%r8)
	shlw	$2, 0x12345678(%r9)
	shlw	$2, 0x12345678(%r10)
	shlw	$2, 0x12345678(%r11)
	shlw	$2, 0x12345678(%r12)
	shlw	$2, 0x12345678(%r13)
	shlw	$2, 0x12345678(%r14)
	shlw	$2, 0x12345678(%r15)
        nop
	shlw	$7, 0x12345678(%rax)
	shlw	$7, 0x12345678(%rdx)
	shlw	$7, 0x12345678(%rbx)
	shlw	$7, 0x12345678(%rsp)
	shlw	$7, 0x12345678(%rbp)
	shlw	$7, 0x12345678(%rsi)
	shlw	$7, 0x12345678(%rdi)
	shlw	$7, 0x12345678(%r8)
	shlw	$7, 0x12345678(%r9)
	shlw	$7, 0x12345678(%r10)
	shlw	$7, 0x12345678(%r11)
	shlw	$7, 0x12345678(%r12)
	shlw	$7, 0x12345678(%r13)
	shlw	$7, 0x12345678(%r14)
	shlw	$7, 0x12345678(%r15)
        nop
        nop
        ret
        ret
	.cfi_endproc



	.p2align 4,,15
	.globl	ShlwyC
	.type	ShlwyCl, @function
ShlwyCl:
	.cfi_startproc
	// reg16 <<= %cl
	shlw	%cl,%ax
	shlw	%cl,%cx
	shlw	%cl,%dx
	shlw	%cl,%bx
	shlw	%cl,%sp
	shlw	%cl,%bp
	shlw	%cl,%si
	shlw	%cl,%di
	shlw	%cl,%r8w
	shlw	%cl,%r9w
	shlw	%cl,%r10w
	shlw	%cl,%r11w
	shlw	%cl,%r12w
	shlw	%cl,%r13w
	shlw	%cl,%r14w
	shlw	%cl,%r15w
        nop
	// mem16[0] <<= %cl
	shlw	%cl,(%rax)
	shlw	%cl,(%rdx)
	shlw	%cl,(%rbx)
	shlw	%cl,(%rsp)
	shlw	%cl,(%rbp)
	shlw	%cl,(%rsi)
	shlw	%cl,(%rdi)
	shlw	%cl,(%r8)
	shlw	%cl,(%r9)
	shlw	%cl,(%r10)
	shlw	%cl,(%r11)
	shlw	%cl,(%r12)
	shlw	%cl,(%r13)
	shlw	%cl,(%r14)
	shlw	%cl,(%r15)
        nop
	// mem16[off8] <<= %cl
	shlw	%cl, 0x7F(%rax)
	shlw	%cl, 0x7F(%rdx)
	shlw	%cl, 0x7F(%rbx)
	shlw	%cl, 0x7F(%rsp)
	shlw	%cl, 0x7F(%rbp)
	shlw	%cl, 0x7F(%rsi)
	shlw	%cl, 0x7F(%rdi)
	shlw	%cl, 0x7F(%r8)
	shlw	%cl, 0x7F(%r9)
	shlw	%cl, 0x7F(%r10)
	shlw	%cl, 0x7F(%r11)
	shlw	%cl, 0x7F(%r12)
	shlw	%cl, 0x7F(%r13)
	shlw	%cl, 0x7F(%r14)
	shlw	%cl, 0x7F(%r15)
        nop
	// mem16[off32] <<= %cl
	shlw	%cl, 0x12345678(%rax)
	shlw	%cl, 0x12345678(%rdx)
	shlw	%cl, 0x12345678(%rbx)
	shlw	%cl, 0x12345678(%rsp)
	shlw	%cl, 0x12345678(%rbp)
	shlw	%cl, 0x12345678(%rsi)
	shlw	%cl, 0x12345678(%rdi)
	shlw	%cl, 0x12345678(%r8)
	shlw	%cl, 0x12345678(%r9)
	shlw	%cl, 0x12345678(%r10)
	shlw	%cl, 0x12345678(%r11)
	shlw	%cl, 0x12345678(%r12)
	shlw	%cl, 0x12345678(%r13)
	shlw	%cl, 0x12345678(%r14)
	shlw	%cl, 0x12345678(%r15)
        nop
        ret
	.cfi_endproc
