	.file	"shl.s"
	.text

	.p2align 4,,15
	.globl	ShlByConst
	.type	ShlByConst, @function
ShlByConst:
	.cfi_startproc
	// reg8 <<= const
	shr	$1, %al
	shl	$1, %al
	shl	$1, %cl
	shl	$1, %dl
	shl	$1, %bl
	shl	$1, %spl
	shl	$1, %bpl
	shl	$1, %sil
	shl	$1, %dil
	shl	$1, %r8b
	shl	$1, %r9b
	shl	$1, %r10b
	shl	$1, %r11b
	shl	$1, %r12b
	shl	$1, %r13b
	shl	$1, %r14b
	shl	$1, %r15b
        nop
	shl	$2, %al
	shl	$2, %cl
	shl	$2, %dl
	shl	$2, %bl
	shl	$2, %spl
	shl	$2, %bpl
	shl	$2, %sil
	shl	$2, %dil
	shl	$2, %r8b
	shl	$2, %r9b
	shl	$2, %r10b
	shl	$2, %r11b
	shl	$2, %r12b
	shl	$2, %r13b
	shl	$2, %r14b
	shl	$2, %r15b
	nop
	shl	$7, %al
	shl	$7, %cl
	shl	$7, %dl
	shl	$7, %bl
	shl	$7, %spl
	shl	$7, %bpl
	shl	$7, %sil
	shl	$7, %dil
	shl	$7, %r8b
	shl	$7, %r9b
	shl	$7, %r10b
	shl	$7, %r11b
	shl	$7, %r12b
	shl	$7, %r13b
	shl	$7, %r14b
	shl	$7, %r15b
        nop
        // mem8[0] <<= const
	shlb	$1,(%rax)
	shlb	$1,(%rcx)
	shlb	$1,(%rdx)
	shlb	$1,(%rbx)
	shlb	$1,(%rsp)
	shlb	$1,(%rbp)
	shlb	$1,(%rsi)
	shlb	$1,(%rdi)
	shlb	$1,(%r8)
	shlb	$1,(%r9)
	shlb	$1,(%r10)
	shlb	$1,(%r11)
	shlb	$1,(%r12)
	shlb	$1,(%r13)
	shlb	$1,(%r14)
	shlb	$1,(%r15)
        nop
	shlb	$2,(%rax)
	shlb	$2,(%rcx)
	shlb	$2,(%rdx)
	shlb	$2,(%rbx)
	shlb	$2,(%rsp)
	shlb	$2,(%rbp)
	shlb	$2,(%rsi)
	shlb	$2,(%rdi)
	shlb	$2,(%r8)
	shlb	$2,(%r9)
	shlb	$2,(%r10)
	shlb	$2,(%r11)
	shlb	$2,(%r12)
	shlb	$2,(%r13)
	shlb	$2,(%r14)
	shlb	$2,(%r15)
        nop
	shlb	$7,(%rax)
	shlb	$7,(%rcx)
	shlb	$7,(%rdx)
	shlb	$7,(%rbx)
	shlb	$7,(%rsp)
	shlb	$7,(%rbp)
	shlb	$7,(%rsi)
	shlb	$7,(%rdi)
	shlb	$7,(%r8)
	shlb	$7,(%r9)
	shlb	$7,(%r10)
	shlb	$7,(%r11)
	shlb	$7,(%r12)
	shlb	$7,(%r13)
	shlb	$7,(%r14)
	shlb	$7,(%r15)
        nop
        // mem8[off8] <<= const
	shlb	$1, 0x7F(%rax)
	shlb	$1, 0x7F(%rdx)
	shlb	$1, 0x7F(%rbx)
	shlb	$1, 0x7F(%rsp)
	shlb	$1, 0x7F(%rbp)
	shlb	$1, 0x7F(%rsi)
	shlb	$1, 0x7F(%rdi)
	shlb	$1, 0x7F(%r8)
	shlb	$1, 0x7F(%r9)
	shlb	$1, 0x7F(%r10)
	shlb	$1, 0x7F(%r11)
	shlb	$1, 0x7F(%r12)
	shlb	$1, 0x7F(%r13)
	shlb	$1, 0x7F(%r14)
	shlb	$1, 0x7F(%r15)
        nop
	shlb	$2, 0x7F(%rax)
	shlb	$2, 0x7F(%rdx)
	shlb	$2, 0x7F(%rbx)
	shlb	$2, 0x7F(%rsp)
	shlb	$2, 0x7F(%rbp)
	shlb	$2, 0x7F(%rsi)
	shlb	$2, 0x7F(%rdi)
	shlb	$2, 0x7F(%r8)
	shlb	$2, 0x7F(%r9)
	shlb	$2, 0x7F(%r10)
	shlb	$2, 0x7F(%r11)
	shlb	$2, 0x7F(%r12)
	shlb	$2, 0x7F(%r13)
	shlb	$2, 0x7F(%r14)
	shlb	$2, 0x7F(%r15)
        nop
	shlb	$7, 0x7F(%rax)
	shlb	$7, 0x7F(%rdx)
	shlb	$7, 0x7F(%rbx)
	shlb	$7, 0x7F(%rsp)
	shlb	$7, 0x7F(%rbp)
	shlb	$7, 0x7F(%rsi)
	shlb	$7, 0x7F(%rdi)
	shlb	$7, 0x7F(%r8)
	shlb	$7, 0x7F(%r9)
	shlb	$7, 0x7F(%r10)
	shlb	$7, 0x7F(%r11)
	shlb	$7, 0x7F(%r12)
	shlb	$7, 0x7F(%r13)
	shlb	$7, 0x7F(%r14)
	shlb	$7, 0x7F(%r15)
        nop
        // mem8[off32] <<= const
	shlb	$1, 0x12345678(%rax)
	shlb	$1, 0x12345678(%rdx)
	shlb	$1, 0x12345678(%rbx)
	shlb	$1, 0x12345678(%rsp)
	shlb	$1, 0x12345678(%rbp)
	shlb	$1, 0x12345678(%rsi)
	shlb	$1, 0x12345678(%rdi)
	shlb	$1, 0x12345678(%r8)
	shlb	$1, 0x12345678(%r9)
	shlb	$1, 0x12345678(%r10)
	shlb	$1, 0x12345678(%r11)
	shlb	$1, 0x12345678(%r12)
	shlb	$1, 0x12345678(%r13)
	shlb	$1, 0x12345678(%r14)
	shlb	$1, 0x12345678(%r15)
        nop
	shlb	$2, 0x12345678(%rax)
	shlb	$2, 0x12345678(%rdx)
	shlb	$2, 0x12345678(%rbx)
	shlb	$2, 0x12345678(%rsp)
	shlb	$2, 0x12345678(%rbp)
	shlb	$2, 0x12345678(%rsi)
	shlb	$2, 0x12345678(%rdi)
	shlb	$2, 0x12345678(%r8)
	shlb	$2, 0x12345678(%r9)
	shlb	$2, 0x12345678(%r10)
	shlb	$2, 0x12345678(%r11)
	shlb	$2, 0x12345678(%r12)
	shlb	$2, 0x12345678(%r13)
	shlb	$2, 0x12345678(%r14)
	shlb	$2, 0x12345678(%r15)
        nop
	shlb	$7, 0x12345678(%rax)
	shlb	$7, 0x12345678(%rdx)
	shlb	$7, 0x12345678(%rbx)
	shlb	$7, 0x12345678(%rsp)
	shlb	$7, 0x12345678(%rbp)
	shlb	$7, 0x12345678(%rsi)
	shlb	$7, 0x12345678(%rdi)
	shlb	$7, 0x12345678(%r8)
	shlb	$7, 0x12345678(%r9)
	shlb	$7, 0x12345678(%r10)
	shlb	$7, 0x12345678(%r11)
	shlb	$7, 0x12345678(%r12)
	shlb	$7, 0x12345678(%r13)
	shlb	$7, 0x12345678(%r14)
	shlb	$7, 0x12345678(%r15)
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
	// reg8 <<= %cl
	shlb	%cl,%al
	shlb	%cl,%cl
	shlb	%cl,%dl
	shlb	%cl,%bl
	shlb	%cl,%spl
	shlb	%cl,%bpl
	shlb	%cl,%sil
	shlb	%cl,%dil
	shlb	%cl,%r8b
	shlb	%cl,%r9b
	shlb	%cl,%r10b
	shlb	%cl,%r11b
	shlb	%cl,%r12b
	shlb	%cl,%r13b
	shlb	%cl,%r14b
	shlb	%cl,%r15b
        nop
	// mem8[0] <<= %cl
	shlb	%cl,(%rax)
	shlb	%cl,(%rdx)
	shlb	%cl,(%rbx)
	shlb	%cl,(%rsp)
	shlb	%cl,(%rbp)
	shlb	%cl,(%rsi)
	shlb	%cl,(%rdi)
	shlb	%cl,(%r8)
	shlb	%cl,(%r9)
	shlb	%cl,(%r10)
	shlb	%cl,(%r11)
	shlb	%cl,(%r12)
	shlb	%cl,(%r13)
	shlb	%cl,(%r14)
	shlb	%cl,(%r15)
        nop
	// mem8[off8] <<= %cl
	shlb	%cl, 0x7F(%rax)
	shlb	%cl, 0x7F(%rdx)
	shlb	%cl, 0x7F(%rbx)
	shlb	%cl, 0x7F(%rsp)
	shlb	%cl, 0x7F(%rbp)
	shlb	%cl, 0x7F(%rsi)
	shlb	%cl, 0x7F(%rdi)
	shlb	%cl, 0x7F(%r8)
	shlb	%cl, 0x7F(%r9)
	shlb	%cl, 0x7F(%r10)
	shlb	%cl, 0x7F(%r11)
	shlb	%cl, 0x7F(%r12)
	shlb	%cl, 0x7F(%r13)
	shlb	%cl, 0x7F(%r14)
	shlb	%cl, 0x7F(%r15)
        nop
	// mem8[off32] <<= %cl
	shlb	%cl, 0x12345678(%rax)
	shlb	%cl, 0x12345678(%rdx)
	shlb	%cl, 0x12345678(%rbx)
	shlb	%cl, 0x12345678(%rsp)
	shlb	%cl, 0x12345678(%rbp)
	shlb	%cl, 0x12345678(%rsi)
	shlb	%cl, 0x12345678(%rdi)
	shlb	%cl, 0x12345678(%r8)
	shlb	%cl, 0x12345678(%r9)
	shlb	%cl, 0x12345678(%r10)
	shlb	%cl, 0x12345678(%r11)
	shlb	%cl, 0x12345678(%r12)
	shlb	%cl, 0x12345678(%r13)
	shlb	%cl, 0x12345678(%r14)
	shlb	%cl, 0x12345678(%r15)
        nop
        ret
	.cfi_endproc
