	.file	"imul.s"
	.text

	.p2align 4,,15
	.globl	Imul_s32
	.type	Imul_s32, @function
Imul_s32:
	.cfi_startproc
	.byte 0x48, 0x81, 0xc0, 0x78, 0x88, 0x99, 0xaa
	// imul	$-0x55667788,%rax
	imul	$-0x55667788,%rcx
	imul	$-0x55667788,%rdx
	imul	$-0x55667788,%rbx
	imul	$-0x55667788,%rsp
	imul	$-0x55667788,%rbp
	imul	$-0x55667788,%rsi
	imul	$-0x55667788,%rdi
	imul	$-0x55667788,%r8
	imul	$-0x55667788,%r9
	imul	$-0x55667788,%r10
	imul	$-0x55667788,%r11
	imul	$-0x55667788,%r12
	imul	$-0x55667788,%r13
	imul	$-0x55667788,%r14
	imul	$-0x55667788,%r15
        ret
	.cfi_endproc



        .p2align 4,,15
	.globl	Imul8
	.type	Imul8, @function
Imul8:
	.cfi_startproc

        // %al *= reg8
	imulb	%al
	imulb	%cl
	imulb	%dl
	imulb	%bl
	imulb	%spl
	imulb	%bpl
	imulb	%sil
	imulb	%dil
	imulb	%r8b
	imulb	%r9b
	imulb	%r10b
	imulb	%r11b
	imulb	%r12b
	imulb	%r13b
	imulb	%r14b
	imulb	%r15b
        nop
        // %al *= mem8
	imulb	(%rax)
	imulb	(%rcx)
	imulb	(%rdx)
	imulb	(%rbx)
	imulb	(%rsp)
	imulb	(%rbp)
	imulb	(%rsi)
	imulb	(%rdi)
	imulb	(%r8)
	imulb	(%r9)
	imulb	(%r10)
	imulb	(%r11)
	imulb	(%r12)
	imulb	(%r13)
	imulb	(%r14)
	imulb	(%r15)
        nop
        // %al *= mem8[off8]
	imulb	0x7F(%rax)
	imulb	0x7F(%rcx)
	imulb	0x7F(%rdx)
	imulb	0x7F(%rbx)
	imulb	0x7F(%rsp)
	imulb	0x7F(%rbp)
	imulb	0x7F(%rsi)
	imulb	0x7F(%rdi)
	imulb	0x7F(%r8)
	imulb	0x7F(%r9)
	imulb	0x7F(%r10)
	imulb	0x7F(%r11)
	imulb	0x7F(%r12)
	imulb	0x7F(%r13)
	imulb	0x7F(%r14)
	imulb	0x7F(%r15)
        nop
        // %al *= mem8[off32]
	imulb	0x12345678(%rax)
	imulb	0x12345678(%rcx)
	imulb	0x12345678(%rdx)
	imulb	0x12345678(%rbx)
	imulb	0x12345678(%rsp)
	imulb	0x12345678(%rbp)
	imulb	0x12345678(%rsi)
	imulb	0x12345678(%rdi)
	imulb	0x12345678(%r8)
	imulb	0x12345678(%r9)
	imulb	0x12345678(%r10)
	imulb	0x12345678(%r11)
	imulb	0x12345678(%r12)
	imulb	0x12345678(%r13)
	imulb	0x12345678(%r14)
	imulb	0x12345678(%r15)
        ret
	.cfi_endproc
        
        .p2align 4,,15
	.globl	Imul16
	.type	Imul16, @function
Imul16:
	.cfi_startproc
	nop
        nop
        nop
        // reg16 *= reg16
	imulw	%ax,%ax
	imulw	%ax,%cx
	imulw	%ax,%dx
	imulw	%ax,%bx
	imulw	%ax,%sp
	imulw	%ax,%bp
	imulw	%ax,%si
	imulw	%ax,%di
	imulw	%ax,%r8w
	imulw	%ax,%r9w
	imulw	%ax,%r10w
	imulw	%ax,%r11w
	imulw	%ax,%r12w
	imulw	%ax,%r13w
	imulw	%ax,%r14w
	imulw	%ax,%r15w
	nop
	imulw	%cx,%ax
	imulw	%cx,%cx
	imulw	%cx,%dx
	imulw	%cx,%bx
	imulw	%cx,%sp
	imulw	%cx,%bp
	imulw	%cx,%si
	imulw	%cx,%di
	imulw	%cx,%r8w
	imulw	%cx,%r9w
	imulw	%cx,%r10w
	imulw	%cx,%r11w
	imulw	%cx,%r12w
	imulw	%cx,%r13w
	imulw	%cx,%r14w
	imulw	%cx,%r15w
	nop
	imulw	%dx,%ax
	imulw	%dx,%cx
	imulw	%dx,%dx
	imulw	%dx,%bx
	imulw	%dx,%sp
	imulw	%dx,%bp
	imulw	%dx,%si
	imulw	%dx,%di
	imulw	%dx,%r8w
	imulw	%dx,%r9w
	imulw	%dx,%r10w
	imulw	%dx,%r11w
	imulw	%dx,%r12w
	imulw	%dx,%r13w
	imulw	%dx,%r14w
	imulw	%dx,%r15w
	nop
	imulw	%bx,%ax
	imulw	%bx,%cx
	imulw	%bx,%dx
	imulw	%bx,%bx
	imulw	%bx,%sp
	imulw	%bx,%bp
	imulw	%bx,%si
	imulw	%bx,%di
	imulw	%bx,%r8w
	imulw	%bx,%r9w
	imulw	%bx,%r10w
	imulw	%bx,%r11w
	imulw	%bx,%r12w
	imulw	%bx,%r13w
	imulw	%bx,%r14w
	imulw	%bx,%r15w
	nop
	imulw	%sp,%ax
	imulw	%sp,%cx
	imulw	%sp,%dx
	imulw	%sp,%bx
	imulw	%sp,%sp
	imulw	%sp,%bp
	imulw	%sp,%si
	imulw	%sp,%di
	imulw	%sp,%r8w
	imulw	%sp,%r9w
	imulw	%sp,%r10w
	imulw	%sp,%r11w
	imulw	%sp,%r12w
	imulw	%sp,%r13w
	imulw	%sp,%r14w
	imulw	%sp,%r15w
	nop
	imulw	%bp,%ax
	imulw	%bp,%cx
	imulw	%bp,%dx
	imulw	%bp,%bx
	imulw	%bp,%sp
	imulw	%bp,%bp
	imulw	%bp,%si
	imulw	%bp,%di
	imulw	%bp,%r8w
	imulw	%bp,%r9w
	imulw	%bp,%r10w
	imulw	%bp,%r11w
	imulw	%bp,%r12w
	imulw	%bp,%r13w
	imulw	%bp,%r14w
	imulw	%bp,%r15w
	nop
	imulw	%si,%ax
	imulw	%si,%cx
	imulw	%si,%dx
	imulw	%si,%bx
	imulw	%si,%sp
	imulw	%si,%bp
	imulw	%si,%si
	imulw	%si,%di
	imulw	%si,%r8w
	imulw	%si,%r9w
	imulw	%si,%r10w
	imulw	%si,%r11w
	imulw	%si,%r12w
	imulw	%si,%r13w
	imulw	%si,%r14w
	imulw	%si,%r15w
	nop
	imulw	%di,%ax
	imulw	%di,%cx
	imulw	%di,%dx
	imulw	%di,%bx
	imulw	%di,%sp
	imulw	%di,%bp
	imulw	%di,%si
	imulw	%di,%di
	imulw	%di,%r8w
	imulw	%di,%r9w
	imulw	%di,%r10w
	imulw	%di,%r11w
	imulw	%di,%r12w
	imulw	%di,%r13w
	imulw	%di,%r14w
	imulw	%di,%r15w
	nop
	imulw	%r8w, %ax
	imulw	%r8w, %cx
	imulw	%r8w, %dx
	imulw	%r8w, %bx
	imulw	%r8w, %sp
	imulw	%r8w, %bp
	imulw	%r8w, %si
	imulw	%r8w, %di
	imulw	%r8w, %r8w
	imulw	%r8w, %r9w
	imulw	%r8w, %r10w
	imulw	%r8w, %r11w
	imulw	%r8w, %r12w
	imulw	%r8w, %r13w
	imulw	%r8w, %r14w
	imulw	%r8w, %r15w
	nop
	imulw	%r9w, %ax
	imulw	%r9w, %cx
	imulw	%r9w, %dx
	imulw	%r9w, %bx
	imulw	%r9w, %sp
	imulw	%r9w, %bp
	imulw	%r9w, %si
	imulw	%r9w, %di
	imulw	%r9w, %r8w
	imulw	%r9w, %r9w
	imulw	%r9w, %r10w
	imulw	%r9w, %r11w
	imulw	%r9w, %r12w
	imulw	%r9w, %r13w
	imulw	%r9w, %r14w
	imulw	%r9w, %r15w
	nop
	imulw	%r10w,%ax
	imulw	%r10w,%cx
	imulw	%r10w,%dx
	imulw	%r10w,%bx
	imulw	%r10w,%sp
	imulw	%r10w,%bp
	imulw	%r10w,%si
	imulw	%r10w,%di
	imulw	%r10w,%r8w
	imulw	%r10w,%r9w
	imulw	%r10w,%r10w
	imulw	%r10w,%r11w
	imulw	%r10w,%r12w
	imulw	%r10w,%r13w
	imulw	%r10w,%r14w
	imulw	%r10w,%r15w
	nop
	imulw	%r11w,%ax
	imulw	%r11w,%cx
	imulw	%r11w,%dx
	imulw	%r11w,%bx
	imulw	%r11w,%sp
	imulw	%r11w,%bp
	imulw	%r11w,%si
	imulw	%r11w,%di
	imulw	%r11w,%r8w
	imulw	%r11w,%r9w
	imulw	%r11w,%r10w
	imulw	%r11w,%r11w
	imulw	%r11w,%r12w
	imulw	%r11w,%r13w
	imulw	%r11w,%r14w
	imulw	%r11w,%r15w
	nop
	imulw	%r12w,%ax
	imulw	%r12w,%cx
	imulw	%r12w,%dx
	imulw	%r12w,%bx
	imulw	%r12w,%sp
	imulw	%r12w,%bp
	imulw	%r12w,%si
	imulw	%r12w,%di
	imulw	%r12w,%r8w
	imulw	%r12w,%r9w
	imulw	%r12w,%r10w
	imulw	%r12w,%r11w
	imulw	%r12w,%r12w
	imulw	%r12w,%r13w
	imulw	%r12w,%r14w
	imulw	%r12w,%r15w
	nop
	imulw	%r13w,%ax
	imulw	%r13w,%cx
	imulw	%r13w,%dx
	imulw	%r13w,%bx
	imulw	%r13w,%sp
	imulw	%r13w,%bp
	imulw	%r13w,%si
	imulw	%r13w,%di
	imulw	%r13w,%r8w
	imulw	%r13w,%r9w
	imulw	%r13w,%r10w
	imulw	%r13w,%r11w
	imulw	%r13w,%r12w
	imulw	%r13w,%r13w
	imulw	%r13w,%r14w
	imulw	%r13w,%r15w
	nop
	imulw	%r14w,%ax
	imulw	%r14w,%cx
	imulw	%r14w,%dx
	imulw	%r14w,%bx
	imulw	%r14w,%sp
	imulw	%r14w,%bp
	imulw	%r14w,%si
	imulw	%r14w,%di
	imulw	%r14w,%r8w
	imulw	%r14w,%r9w
	imulw	%r14w,%r10w
	imulw	%r14w,%r11w
	imulw	%r14w,%r12w
	imulw	%r14w,%r13w
	imulw	%r14w,%r14w
	imulw	%r14w,%r15w
	nop
	imulw	%r15w,%ax
	imulw	%r15w,%cx
	imulw	%r15w,%dx
	imulw	%r15w,%bx
	imulw	%r15w,%sp
	imulw	%r15w,%bp
	imulw	%r15w,%si
	imulw	%r15w,%di
	imulw	%r15w,%r8w
	imulw	%r15w,%r9w
	imulw	%r15w,%r10w
	imulw	%r15w,%r11w
	imulw	%r15w,%r12w
	imulw	%r15w,%r13w
	imulw	%r15w,%r14w
	imulw	%r15w,%r15w
        nop
        nop
        // mem *= reg NOT SUPPORTED!
        // reg16 *= mem16
	imulw	(%rax),%ax
	imulw	(%rcx),%ax
	imulw	(%rdx),%ax
	imulw	(%rbx),%ax
	imulw	(%rsp),%ax
	imulw	(%rbp),%ax
	imulw	(%rsi),%ax
	imulw	(%rdi),%ax
	imulw	(%r8 ),%ax
	imulw	(%r9 ),%ax
	imulw	(%r10),%ax
	imulw	(%r11),%ax
	imulw	(%r12),%ax
	imulw	(%r13),%ax
	imulw	(%r14),%ax
	imulw	(%r15),%ax
	nop                
	imulw	(%rax),%cx
	imulw	(%rcx),%cx
	imulw	(%rdx),%cx
	imulw	(%rbx),%cx
	imulw	(%rsp),%cx
	imulw	(%rbp),%cx
	imulw	(%rsi),%cx
	imulw	(%rdi),%cx
	imulw	(%r8 ),%cx
	imulw	(%r9 ),%cx
	imulw	(%r10),%cx
	imulw	(%r11),%cx
	imulw	(%r12),%cx
	imulw	(%r13),%cx
	imulw	(%r14),%cx
	imulw	(%r15),%cx
	nop                
	imulw	(%rax),%dx
	imulw	(%rcx),%dx
	imulw	(%rdx),%dx
	imulw	(%rbx),%dx
	imulw	(%rsp),%dx
	imulw	(%rbp),%dx
	imulw	(%rsi),%dx
	imulw	(%rdi),%dx
	imulw	(%r8 ),%dx
	imulw	(%r9 ),%dx
	imulw	(%r10),%dx
	imulw	(%r11),%dx
	imulw	(%r12),%dx
	imulw	(%r13),%dx
	imulw	(%r14),%dx
	imulw	(%r15),%dx
	nop                
	imulw	(%rax),%bx
	imulw	(%rcx),%bx
	imulw	(%rdx),%bx
	imulw	(%rbx),%bx
	imulw	(%rsp),%bx
	imulw	(%rbp),%bx
	imulw	(%rsi),%bx
	imulw	(%rdi),%bx
	imulw	(%r8 ),%bx
	imulw	(%r9 ),%bx
	imulw	(%r10),%bx
	imulw	(%r11),%bx
	imulw	(%r12),%bx
	imulw	(%r13),%bx
	imulw	(%r14),%bx
	imulw	(%r15),%bx
	nop                
	imulw	(%rax),%sp
	imulw	(%rcx),%sp
	imulw	(%rdx),%sp
	imulw	(%rbx),%sp
	imulw	(%rsp),%sp
	imulw	(%rbp),%sp
	imulw	(%rsi),%sp
	imulw	(%rdi),%sp
	imulw	(%r8 ),%sp
	imulw	(%r9 ),%sp
	imulw	(%r10),%sp
	imulw	(%r11),%sp
	imulw	(%r12),%sp
	imulw	(%r13),%sp
	imulw	(%r14),%sp
	imulw	(%r15),%sp
	nop                
	imulw	(%rax),%bp
	imulw	(%rcx),%bp
	imulw	(%rdx),%bp
	imulw	(%rbx),%bp
	imulw	(%rsp),%bp
	imulw	(%rbp),%bp
	imulw	(%rsi),%bp
	imulw	(%rdi),%bp
	imulw	(%r8 ),%bp
	imulw	(%r9 ),%bp
	imulw	(%r10),%bp
	imulw	(%r11),%bp
	imulw	(%r12),%bp
	imulw	(%r13),%bp
	imulw	(%r14),%bp
	imulw	(%r15),%bp
	nop                
	imulw	(%rax),%si
	imulw	(%rcx),%si
	imulw	(%rdx),%si
	imulw	(%rbx),%si
	imulw	(%rsp),%si
	imulw	(%rbp),%si
	imulw	(%rsi),%si
	imulw	(%rdi),%si
	imulw	(%r8 ),%si
	imulw	(%r9 ),%si
	imulw	(%r10),%si
	imulw	(%r11),%si
	imulw	(%r12),%si
	imulw	(%r13),%si
	imulw	(%r14),%si
	imulw	(%r15),%si
	nop                
	imulw	(%rax),%di
	imulw	(%rcx),%di
	imulw	(%rdx),%di
	imulw	(%rbx),%di
	imulw	(%rsp),%di
	imulw	(%rbp),%di
	imulw	(%rsi),%di
	imulw	(%rdi),%di
	imulw	(%r8 ),%di
	imulw	(%r9 ),%di
	imulw	(%r10),%di
	imulw	(%r11),%di
	imulw	(%r12),%di
	imulw	(%r13),%di
	imulw	(%r14),%di
	imulw	(%r15),%di
	nop
	imulw	(%rax),%r8w
	imulw	(%rcx),%r8w
	imulw	(%rdx),%r8w
	imulw	(%rbx),%r8w
	imulw	(%rsp),%r8w
	imulw	(%rbp),%r8w
	imulw	(%rsi),%r8w
	imulw	(%rdi),%r8w
	imulw	(%r8), %r8w
	imulw	(%r9), %r8w
	imulw	(%r10),%r8w
	imulw	(%r11),%r8w
	imulw	(%r12),%r8w
	imulw	(%r13),%r8w
	imulw	(%r14),%r8w
	imulw	(%r15),%r8w
	nop                  
	imulw	(%rax),%r9w
	imulw	(%rcx),%r9w
	imulw	(%rdx),%r9w
	imulw	(%rbx),%r9w
	imulw	(%rsp),%r9w
	imulw	(%rbp),%r9w
	imulw	(%rsi),%r9w
	imulw	(%rdi),%r9w
	imulw	(%r8), %r9w
	imulw	(%r9), %r9w
	imulw	(%r10),%r9w
	imulw	(%r11),%r9w
	imulw	(%r12),%r9w
	imulw	(%r13),%r9w
	imulw	(%r14),%r9w
	imulw	(%r15),%r9w
	nop                  
	imulw	(%rax),%r10w
	imulw	(%rcx),%r10w
	imulw	(%rdx),%r10w
	imulw	(%rbx),%r10w
	imulw	(%rsp),%r10w
	imulw	(%rbp),%r10w
	imulw	(%rsi),%r10w
	imulw	(%rdi),%r10w
	imulw	(%r8), %r10w
	imulw	(%r9), %r10w
	imulw	(%r10),%r10w
	imulw	(%r11),%r10w
	imulw	(%r12),%r10w
	imulw	(%r13),%r10w
	imulw	(%r14),%r10w
	imulw	(%r15),%r10w
	nop                  
	imulw	(%rax),%r11w
	imulw	(%rcx),%r11w
	imulw	(%rdx),%r11w
	imulw	(%rbx),%r11w
	imulw	(%rsp),%r11w
	imulw	(%rbp),%r11w
	imulw	(%rsi),%r11w
	imulw	(%rdi),%r11w
	imulw	(%r8), %r11w
	imulw	(%r9), %r11w
	imulw	(%r10),%r11w
	imulw	(%r11),%r11w
	imulw	(%r12),%r11w
	imulw	(%r13),%r11w
	imulw	(%r14),%r11w
	imulw	(%r15),%r11w
	nop                  
	imulw	(%rax),%r12w
	imulw	(%rcx),%r12w
	imulw	(%rdx),%r12w
	imulw	(%rbx),%r12w
	imulw	(%rsp),%r12w
	imulw	(%rbp),%r12w
	imulw	(%rsi),%r12w
	imulw	(%rdi),%r12w
	imulw	(%r8), %r12w
	imulw	(%r9), %r12w
	imulw	(%r10),%r12w
	imulw	(%r11),%r12w
	imulw	(%r12),%r12w
	imulw	(%r13),%r12w
	imulw	(%r14),%r12w
	imulw	(%r15),%r12w
	nop                  
	imulw	(%rax),%r13w
	imulw	(%rcx),%r13w
	imulw	(%rdx),%r13w
	imulw	(%rbx),%r13w
	imulw	(%rsp),%r13w
	imulw	(%rbp),%r13w
	imulw	(%rsi),%r13w
	imulw	(%rdi),%r13w
	imulw	(%r8), %r13w
	imulw	(%r9), %r13w
	imulw	(%r10),%r13w
	imulw	(%r11),%r13w
	imulw	(%r12),%r13w
	imulw	(%r13),%r13w
	imulw	(%r14),%r13w
	imulw	(%r15),%r13w
	nop                  
	imulw	(%rax),%r14w
	imulw	(%rcx),%r14w
	imulw	(%rdx),%r14w
	imulw	(%rbx),%r14w
	imulw	(%rsp),%r14w
	imulw	(%rbp),%r14w
	imulw	(%rsi),%r14w
	imulw	(%rdi),%r14w
	imulw	(%r8), %r14w
	imulw	(%r9), %r14w
	imulw	(%r10),%r14w
	imulw	(%r11),%r14w
	imulw	(%r12),%r14w
	imulw	(%r13),%r14w
	imulw	(%r14),%r14w
	imulw	(%r15),%r14w
	nop                  
	imulw	(%rax),%r15w
	imulw	(%rcx),%r15w
	imulw	(%rdx),%r15w
	imulw	(%rbx),%r15w
	imulw	(%rsp),%r15w
	imulw	(%rbp),%r15w
	imulw	(%rsi),%r15w
	imulw	(%rdi),%r15w
	imulw	(%r8), %r15w
	imulw	(%r9), %r15w
	imulw	(%r10),%r15w
	imulw	(%r11),%r15w
	imulw	(%r12),%r15w
	imulw	(%r13),%r15w
	imulw	(%r14),%r15w
	imulw	(%r15),%r15w
        nop
        nop
        // mem *= reg NOT SUPPORTED!
        // reg16 *= mem16[off8]
	imulw	0x7F(%rax),%ax
	imulw	0x7F(%rcx),%ax
	imulw	0x7F(%rdx),%ax
	imulw	0x7F(%rbx),%ax
	imulw	0x7F(%rsp),%ax
	imulw	0x7F(%rbp),%ax
	imulw	0x7F(%rsi),%ax
	imulw	0x7F(%rdi),%ax
	imulw	0x7F(%r8 ),%ax
	imulw	0x7F(%r9 ),%ax
	imulw	0x7F(%r10),%ax
	imulw	0x7F(%r11),%ax
	imulw	0x7F(%r12),%ax
	imulw	0x7F(%r13),%ax
	imulw	0x7F(%r14),%ax
	imulw	0x7F(%r15),%ax
	nop                
	imulw	0x7F(%rax),%cx
	imulw	0x7F(%rcx),%cx
	imulw	0x7F(%rdx),%cx
	imulw	0x7F(%rbx),%cx
	imulw	0x7F(%rsp),%cx
	imulw	0x7F(%rbp),%cx
	imulw	0x7F(%rsi),%cx
	imulw	0x7F(%rdi),%cx
	imulw	0x7F(%r8 ),%cx
	imulw	0x7F(%r9 ),%cx
	imulw	0x7F(%r10),%cx
	imulw	0x7F(%r11),%cx
	imulw	0x7F(%r12),%cx
	imulw	0x7F(%r13),%cx
	imulw	0x7F(%r14),%cx
	imulw	0x7F(%r15),%cx
	nop                
	imulw	0x7F(%rax),%dx
	imulw	0x7F(%rcx),%dx
	imulw	0x7F(%rdx),%dx
	imulw	0x7F(%rbx),%dx
	imulw	0x7F(%rsp),%dx
	imulw	0x7F(%rbp),%dx
	imulw	0x7F(%rsi),%dx
	imulw	0x7F(%rdi),%dx
	imulw	0x7F(%r8 ),%dx
	imulw	0x7F(%r9 ),%dx
	imulw	0x7F(%r10),%dx
	imulw	0x7F(%r11),%dx
	imulw	0x7F(%r12),%dx
	imulw	0x7F(%r13),%dx
	imulw	0x7F(%r14),%dx
	imulw	0x7F(%r15),%dx
	nop                
	imulw	0x7F(%rax),%bx
	imulw	0x7F(%rcx),%bx
	imulw	0x7F(%rdx),%bx
	imulw	0x7F(%rbx),%bx
	imulw	0x7F(%rsp),%bx
	imulw	0x7F(%rbp),%bx
	imulw	0x7F(%rsi),%bx
	imulw	0x7F(%rdi),%bx
	imulw	0x7F(%r8 ),%bx
	imulw	0x7F(%r9 ),%bx
	imulw	0x7F(%r10),%bx
	imulw	0x7F(%r11),%bx
	imulw	0x7F(%r12),%bx
	imulw	0x7F(%r13),%bx
	imulw	0x7F(%r14),%bx
	imulw	0x7F(%r15),%bx
	nop                
	imulw	0x7F(%rax),%sp
	imulw	0x7F(%rcx),%sp
	imulw	0x7F(%rdx),%sp
	imulw	0x7F(%rbx),%sp
	imulw	0x7F(%rsp),%sp
	imulw	0x7F(%rbp),%sp
	imulw	0x7F(%rsi),%sp
	imulw	0x7F(%rdi),%sp
	imulw	0x7F(%r8 ),%sp
	imulw	0x7F(%r9 ),%sp
	imulw	0x7F(%r10),%sp
	imulw	0x7F(%r11),%sp
	imulw	0x7F(%r12),%sp
	imulw	0x7F(%r13),%sp
	imulw	0x7F(%r14),%sp
	imulw	0x7F(%r15),%sp
	nop                
	imulw	0x7F(%rax),%bp
	imulw	0x7F(%rcx),%bp
	imulw	0x7F(%rdx),%bp
	imulw	0x7F(%rbx),%bp
	imulw	0x7F(%rsp),%bp
	imulw	0x7F(%rbp),%bp
	imulw	0x7F(%rsi),%bp
	imulw	0x7F(%rdi),%bp
	imulw	0x7F(%r8 ),%bp
	imulw	0x7F(%r9 ),%bp
	imulw	0x7F(%r10),%bp
	imulw	0x7F(%r11),%bp
	imulw	0x7F(%r12),%bp
	imulw	0x7F(%r13),%bp
	imulw	0x7F(%r14),%bp
	imulw	0x7F(%r15),%bp
	nop                
	imulw	0x7F(%rax),%si
	imulw	0x7F(%rcx),%si
	imulw	0x7F(%rdx),%si
	imulw	0x7F(%rbx),%si
	imulw	0x7F(%rsp),%si
	imulw	0x7F(%rbp),%si
	imulw	0x7F(%rsi),%si
	imulw	0x7F(%rdi),%si
	imulw	0x7F(%r8 ),%si
	imulw	0x7F(%r9 ),%si
	imulw	0x7F(%r10),%si
	imulw	0x7F(%r11),%si
	imulw	0x7F(%r12),%si
	imulw	0x7F(%r13),%si
	imulw	0x7F(%r14),%si
	imulw	0x7F(%r15),%si
	nop                
	imulw	0x7F(%rax),%di
	imulw	0x7F(%rcx),%di
	imulw	0x7F(%rdx),%di
	imulw	0x7F(%rbx),%di
	imulw	0x7F(%rsp),%di
	imulw	0x7F(%rbp),%di
	imulw	0x7F(%rsi),%di
	imulw	0x7F(%rdi),%di
	imulw	0x7F(%r8 ),%di
	imulw	0x7F(%r9 ),%di
	imulw	0x7F(%r10),%di
	imulw	0x7F(%r11),%di
	imulw	0x7F(%r12),%di
	imulw	0x7F(%r13),%di
	imulw	0x7F(%r14),%di
	imulw	0x7F(%r15),%di
	nop
	imulw	0x7F(%rax),%r8w
	imulw	0x7F(%rcx),%r8w
	imulw	0x7F(%rdx),%r8w
	imulw	0x7F(%rbx),%r8w
	imulw	0x7F(%rsp),%r8w
	imulw	0x7F(%rbp),%r8w
	imulw	0x7F(%rsi),%r8w
	imulw	0x7F(%rdi),%r8w
	imulw	0x7F(%r8), %r8w
	imulw	0x7F(%r9), %r8w
	imulw	0x7F(%r10),%r8w
	imulw	0x7F(%r11),%r8w
	imulw	0x7F(%r12),%r8w
	imulw	0x7F(%r13),%r8w
	imulw	0x7F(%r14),%r8w
	imulw	0x7F(%r15),%r8w
	nop                  
	imulw	0x7F(%rax),%r9w
	imulw	0x7F(%rcx),%r9w
	imulw	0x7F(%rdx),%r9w
	imulw	0x7F(%rbx),%r9w
	imulw	0x7F(%rsp),%r9w
	imulw	0x7F(%rbp),%r9w
	imulw	0x7F(%rsi),%r9w
	imulw	0x7F(%rdi),%r9w
	imulw	0x7F(%r8), %r9w
	imulw	0x7F(%r9), %r9w
	imulw	0x7F(%r10),%r9w
	imulw	0x7F(%r11),%r9w
	imulw	0x7F(%r12),%r9w
	imulw	0x7F(%r13),%r9w
	imulw	0x7F(%r14),%r9w
	imulw	0x7F(%r15),%r9w
	nop                  
	imulw	0x7F(%rax),%r10w
	imulw	0x7F(%rcx),%r10w
	imulw	0x7F(%rdx),%r10w
	imulw	0x7F(%rbx),%r10w
	imulw	0x7F(%rsp),%r10w
	imulw	0x7F(%rbp),%r10w
	imulw	0x7F(%rsi),%r10w
	imulw	0x7F(%rdi),%r10w
	imulw	0x7F(%r8), %r10w
	imulw	0x7F(%r9), %r10w
	imulw	0x7F(%r10),%r10w
	imulw	0x7F(%r11),%r10w
	imulw	0x7F(%r12),%r10w
	imulw	0x7F(%r13),%r10w
	imulw	0x7F(%r14),%r10w
	imulw	0x7F(%r15),%r10w
	nop                  
	imulw	0x7F(%rax),%r11w
	imulw	0x7F(%rcx),%r11w
	imulw	0x7F(%rdx),%r11w
	imulw	0x7F(%rbx),%r11w
	imulw	0x7F(%rsp),%r11w
	imulw	0x7F(%rbp),%r11w
	imulw	0x7F(%rsi),%r11w
	imulw	0x7F(%rdi),%r11w
	imulw	0x7F(%r8), %r11w
	imulw	0x7F(%r9), %r11w
	imulw	0x7F(%r10),%r11w
	imulw	0x7F(%r11),%r11w
	imulw	0x7F(%r12),%r11w
	imulw	0x7F(%r13),%r11w
	imulw	0x7F(%r14),%r11w
	imulw	0x7F(%r15),%r11w
	nop                  
	imulw	0x7F(%rax),%r12w
	imulw	0x7F(%rcx),%r12w
	imulw	0x7F(%rdx),%r12w
	imulw	0x7F(%rbx),%r12w
	imulw	0x7F(%rsp),%r12w
	imulw	0x7F(%rbp),%r12w
	imulw	0x7F(%rsi),%r12w
	imulw	0x7F(%rdi),%r12w
	imulw	0x7F(%r8), %r12w
	imulw	0x7F(%r9), %r12w
	imulw	0x7F(%r10),%r12w
	imulw	0x7F(%r11),%r12w
	imulw	0x7F(%r12),%r12w
	imulw	0x7F(%r13),%r12w
	imulw	0x7F(%r14),%r12w
	imulw	0x7F(%r15),%r12w
	nop                  
	imulw	0x7F(%rax),%r13w
	imulw	0x7F(%rcx),%r13w
	imulw	0x7F(%rdx),%r13w
	imulw	0x7F(%rbx),%r13w
	imulw	0x7F(%rsp),%r13w
	imulw	0x7F(%rbp),%r13w
	imulw	0x7F(%rsi),%r13w
	imulw	0x7F(%rdi),%r13w
	imulw	0x7F(%r8), %r13w
	imulw	0x7F(%r9), %r13w
	imulw	0x7F(%r10),%r13w
	imulw	0x7F(%r11),%r13w
	imulw	0x7F(%r12),%r13w
	imulw	0x7F(%r13),%r13w
	imulw	0x7F(%r14),%r13w
	imulw	0x7F(%r15),%r13w
	nop                  
	imulw	0x7F(%rax),%r14w
	imulw	0x7F(%rcx),%r14w
	imulw	0x7F(%rdx),%r14w
	imulw	0x7F(%rbx),%r14w
	imulw	0x7F(%rsp),%r14w
	imulw	0x7F(%rbp),%r14w
	imulw	0x7F(%rsi),%r14w
	imulw	0x7F(%rdi),%r14w
	imulw	0x7F(%r8), %r14w
	imulw	0x7F(%r9), %r14w
	imulw	0x7F(%r10),%r14w
	imulw	0x7F(%r11),%r14w
	imulw	0x7F(%r12),%r14w
	imulw	0x7F(%r13),%r14w
	imulw	0x7F(%r14),%r14w
	imulw	0x7F(%r15),%r14w
	nop                  
	imulw	0x7F(%rax),%r15w
	imulw	0x7F(%rcx),%r15w
	imulw	0x7F(%rdx),%r15w
	imulw	0x7F(%rbx),%r15w
	imulw	0x7F(%rsp),%r15w
	imulw	0x7F(%rbp),%r15w
	imulw	0x7F(%rsi),%r15w
	imulw	0x7F(%rdi),%r15w
	imulw	0x7F(%r8), %r15w
	imulw	0x7F(%r9), %r15w
	imulw	0x7F(%r10),%r15w
	imulw	0x7F(%r11),%r15w
	imulw	0x7F(%r12),%r15w
	imulw	0x7F(%r13),%r15w
	imulw	0x7F(%r14),%r15w
	imulw	0x7F(%r15),%r15w
        nop
        nop
        // mem *= reg NOT SUPPORTED!
        // reg16 *= mem16[off32]
	imulw	0x12345678(%rax),%ax
	imulw	0x12345678(%rcx),%ax
	imulw	0x12345678(%rdx),%ax
	imulw	0x12345678(%rbx),%ax
	imulw	0x12345678(%rsp),%ax
	imulw	0x12345678(%rbp),%ax
	imulw	0x12345678(%rsi),%ax
	imulw	0x12345678(%rdi),%ax
	imulw	0x12345678(%r8 ),%ax
	imulw	0x12345678(%r9 ),%ax
	imulw	0x12345678(%r10),%ax
	imulw	0x12345678(%r11),%ax
	imulw	0x12345678(%r12),%ax
	imulw	0x12345678(%r13),%ax
	imulw	0x12345678(%r14),%ax
	imulw	0x12345678(%r15),%ax
	nop                
	imulw	0x12345678(%rax),%cx
	imulw	0x12345678(%rcx),%cx
	imulw	0x12345678(%rdx),%cx
	imulw	0x12345678(%rbx),%cx
	imulw	0x12345678(%rsp),%cx
	imulw	0x12345678(%rbp),%cx
	imulw	0x12345678(%rsi),%cx
	imulw	0x12345678(%rdi),%cx
	imulw	0x12345678(%r8 ),%cx
	imulw	0x12345678(%r9 ),%cx
	imulw	0x12345678(%r10),%cx
	imulw	0x12345678(%r11),%cx
	imulw	0x12345678(%r12),%cx
	imulw	0x12345678(%r13),%cx
	imulw	0x12345678(%r14),%cx
	imulw	0x12345678(%r15),%cx
	nop                
	imulw	0x12345678(%rax),%dx
	imulw	0x12345678(%rcx),%dx
	imulw	0x12345678(%rdx),%dx
	imulw	0x12345678(%rbx),%dx
	imulw	0x12345678(%rsp),%dx
	imulw	0x12345678(%rbp),%dx
	imulw	0x12345678(%rsi),%dx
	imulw	0x12345678(%rdi),%dx
	imulw	0x12345678(%r8 ),%dx
	imulw	0x12345678(%r9 ),%dx
	imulw	0x12345678(%r10),%dx
	imulw	0x12345678(%r11),%dx
	imulw	0x12345678(%r12),%dx
	imulw	0x12345678(%r13),%dx
	imulw	0x12345678(%r14),%dx
	imulw	0x12345678(%r15),%dx
	nop                
	imulw	0x12345678(%rax),%bx
	imulw	0x12345678(%rcx),%bx
	imulw	0x12345678(%rdx),%bx
	imulw	0x12345678(%rbx),%bx
	imulw	0x12345678(%rsp),%bx
	imulw	0x12345678(%rbp),%bx
	imulw	0x12345678(%rsi),%bx
	imulw	0x12345678(%rdi),%bx
	imulw	0x12345678(%r8 ),%bx
	imulw	0x12345678(%r9 ),%bx
	imulw	0x12345678(%r10),%bx
	imulw	0x12345678(%r11),%bx
	imulw	0x12345678(%r12),%bx
	imulw	0x12345678(%r13),%bx
	imulw	0x12345678(%r14),%bx
	imulw	0x12345678(%r15),%bx
	nop                
	imulw	0x12345678(%rax),%sp
	imulw	0x12345678(%rcx),%sp
	imulw	0x12345678(%rdx),%sp
	imulw	0x12345678(%rbx),%sp
	imulw	0x12345678(%rsp),%sp
	imulw	0x12345678(%rbp),%sp
	imulw	0x12345678(%rsi),%sp
	imulw	0x12345678(%rdi),%sp
	imulw	0x12345678(%r8 ),%sp
	imulw	0x12345678(%r9 ),%sp
	imulw	0x12345678(%r10),%sp
	imulw	0x12345678(%r11),%sp
	imulw	0x12345678(%r12),%sp
	imulw	0x12345678(%r13),%sp
	imulw	0x12345678(%r14),%sp
	imulw	0x12345678(%r15),%sp
	nop                
	imulw	0x12345678(%rax),%bp
	imulw	0x12345678(%rcx),%bp
	imulw	0x12345678(%rdx),%bp
	imulw	0x12345678(%rbx),%bp
	imulw	0x12345678(%rsp),%bp
	imulw	0x12345678(%rbp),%bp
	imulw	0x12345678(%rsi),%bp
	imulw	0x12345678(%rdi),%bp
	imulw	0x12345678(%r8 ),%bp
	imulw	0x12345678(%r9 ),%bp
	imulw	0x12345678(%r10),%bp
	imulw	0x12345678(%r11),%bp
	imulw	0x12345678(%r12),%bp
	imulw	0x12345678(%r13),%bp
	imulw	0x12345678(%r14),%bp
	imulw	0x12345678(%r15),%bp
	nop                
	imulw	0x12345678(%rax),%si
	imulw	0x12345678(%rcx),%si
	imulw	0x12345678(%rdx),%si
	imulw	0x12345678(%rbx),%si
	imulw	0x12345678(%rsp),%si
	imulw	0x12345678(%rbp),%si
	imulw	0x12345678(%rsi),%si
	imulw	0x12345678(%rdi),%si
	imulw	0x12345678(%r8 ),%si
	imulw	0x12345678(%r9 ),%si
	imulw	0x12345678(%r10),%si
	imulw	0x12345678(%r11),%si
	imulw	0x12345678(%r12),%si
	imulw	0x12345678(%r13),%si
	imulw	0x12345678(%r14),%si
	imulw	0x12345678(%r15),%si
	nop                
	imulw	0x12345678(%rax),%di
	imulw	0x12345678(%rcx),%di
	imulw	0x12345678(%rdx),%di
	imulw	0x12345678(%rbx),%di
	imulw	0x12345678(%rsp),%di
	imulw	0x12345678(%rbp),%di
	imulw	0x12345678(%rsi),%di
	imulw	0x12345678(%rdi),%di
	imulw	0x12345678(%r8 ),%di
	imulw	0x12345678(%r9 ),%di
	imulw	0x12345678(%r10),%di
	imulw	0x12345678(%r11),%di
	imulw	0x12345678(%r12),%di
	imulw	0x12345678(%r13),%di
	imulw	0x12345678(%r14),%di
	imulw	0x12345678(%r15),%di
	nop
	imulw	0x12345678(%rax),%r8w
	imulw	0x12345678(%rcx),%r8w
	imulw	0x12345678(%rdx),%r8w
	imulw	0x12345678(%rbx),%r8w
	imulw	0x12345678(%rsp),%r8w
	imulw	0x12345678(%rbp),%r8w
	imulw	0x12345678(%rsi),%r8w
	imulw	0x12345678(%rdi),%r8w
	imulw	0x12345678(%r8), %r8w
	imulw	0x12345678(%r9), %r8w
	imulw	0x12345678(%r10),%r8w
	imulw	0x12345678(%r11),%r8w
	imulw	0x12345678(%r12),%r8w
	imulw	0x12345678(%r13),%r8w
	imulw	0x12345678(%r14),%r8w
	imulw	0x12345678(%r15),%r8w
	nop                  
	imulw	0x12345678(%rax),%r9w
	imulw	0x12345678(%rcx),%r9w
	imulw	0x12345678(%rdx),%r9w
	imulw	0x12345678(%rbx),%r9w
	imulw	0x12345678(%rsp),%r9w
	imulw	0x12345678(%rbp),%r9w
	imulw	0x12345678(%rsi),%r9w
	imulw	0x12345678(%rdi),%r9w
	imulw	0x12345678(%r8), %r9w
	imulw	0x12345678(%r9), %r9w
	imulw	0x12345678(%r10),%r9w
	imulw	0x12345678(%r11),%r9w
	imulw	0x12345678(%r12),%r9w
	imulw	0x12345678(%r13),%r9w
	imulw	0x12345678(%r14),%r9w
	imulw	0x12345678(%r15),%r9w
	nop                  
	imulw	0x12345678(%rax),%r10w
	imulw	0x12345678(%rcx),%r10w
	imulw	0x12345678(%rdx),%r10w
	imulw	0x12345678(%rbx),%r10w
	imulw	0x12345678(%rsp),%r10w
	imulw	0x12345678(%rbp),%r10w
	imulw	0x12345678(%rsi),%r10w
	imulw	0x12345678(%rdi),%r10w
	imulw	0x12345678(%r8), %r10w
	imulw	0x12345678(%r9), %r10w
	imulw	0x12345678(%r10),%r10w
	imulw	0x12345678(%r11),%r10w
	imulw	0x12345678(%r12),%r10w
	imulw	0x12345678(%r13),%r10w
	imulw	0x12345678(%r14),%r10w
	imulw	0x12345678(%r15),%r10w
	nop                  
	imulw	0x12345678(%rax),%r11w
	imulw	0x12345678(%rcx),%r11w
	imulw	0x12345678(%rdx),%r11w
	imulw	0x12345678(%rbx),%r11w
	imulw	0x12345678(%rsp),%r11w
	imulw	0x12345678(%rbp),%r11w
	imulw	0x12345678(%rsi),%r11w
	imulw	0x12345678(%rdi),%r11w
	imulw	0x12345678(%r8), %r11w
	imulw	0x12345678(%r9), %r11w
	imulw	0x12345678(%r10),%r11w
	imulw	0x12345678(%r11),%r11w
	imulw	0x12345678(%r12),%r11w
	imulw	0x12345678(%r13),%r11w
	imulw	0x12345678(%r14),%r11w
	imulw	0x12345678(%r15),%r11w
	nop                  
	imulw	0x12345678(%rax),%r12w
	imulw	0x12345678(%rcx),%r12w
	imulw	0x12345678(%rdx),%r12w
	imulw	0x12345678(%rbx),%r12w
	imulw	0x12345678(%rsp),%r12w
	imulw	0x12345678(%rbp),%r12w
	imulw	0x12345678(%rsi),%r12w
	imulw	0x12345678(%rdi),%r12w
	imulw	0x12345678(%r8), %r12w
	imulw	0x12345678(%r9), %r12w
	imulw	0x12345678(%r10),%r12w
	imulw	0x12345678(%r11),%r12w
	imulw	0x12345678(%r12),%r12w
	imulw	0x12345678(%r13),%r12w
	imulw	0x12345678(%r14),%r12w
	imulw	0x12345678(%r15),%r12w
	nop                  
	imulw	0x12345678(%rax),%r13w
	imulw	0x12345678(%rcx),%r13w
	imulw	0x12345678(%rdx),%r13w
	imulw	0x12345678(%rbx),%r13w
	imulw	0x12345678(%rsp),%r13w
	imulw	0x12345678(%rbp),%r13w
	imulw	0x12345678(%rsi),%r13w
	imulw	0x12345678(%rdi),%r13w
	imulw	0x12345678(%r8), %r13w
	imulw	0x12345678(%r9), %r13w
	imulw	0x12345678(%r10),%r13w
	imulw	0x12345678(%r11),%r13w
	imulw	0x12345678(%r12),%r13w
	imulw	0x12345678(%r13),%r13w
	imulw	0x12345678(%r14),%r13w
	imulw	0x12345678(%r15),%r13w
	nop                  
	imulw	0x12345678(%rax),%r14w
	imulw	0x12345678(%rcx),%r14w
	imulw	0x12345678(%rdx),%r14w
	imulw	0x12345678(%rbx),%r14w
	imulw	0x12345678(%rsp),%r14w
	imulw	0x12345678(%rbp),%r14w
	imulw	0x12345678(%rsi),%r14w
	imulw	0x12345678(%rdi),%r14w
	imulw	0x12345678(%r8), %r14w
	imulw	0x12345678(%r9), %r14w
	imulw	0x12345678(%r10),%r14w
	imulw	0x12345678(%r11),%r14w
	imulw	0x12345678(%r12),%r14w
	imulw	0x12345678(%r13),%r14w
	imulw	0x12345678(%r14),%r14w
	imulw	0x12345678(%r15),%r14w
	nop                  
	imulw	0x12345678(%rax),%r15w
	imulw	0x12345678(%rcx),%r15w
	imulw	0x12345678(%rdx),%r15w
	imulw	0x12345678(%rbx),%r15w
	imulw	0x12345678(%rsp),%r15w
	imulw	0x12345678(%rbp),%r15w
	imulw	0x12345678(%rsi),%r15w
	imulw	0x12345678(%rdi),%r15w
	imulw	0x12345678(%r8), %r15w
	imulw	0x12345678(%r9), %r15w
	imulw	0x12345678(%r10),%r15w
	imulw	0x12345678(%r11),%r15w
	imulw	0x12345678(%r12),%r15w
	imulw	0x12345678(%r13),%r15w
	imulw	0x12345678(%r14),%r15w
	imulw	0x12345678(%r15),%r15w
        ret
	.cfi_endproc

/*
        .p2align 4,,15
	.globl	Imul32
	.type	Imul32, @function
Imul32:
	.cfi_startproc
        // reg32 += reg32
	imul	%eax,%eax
	imul	%eax,%ecx
	imul	%eax,%edx
	imul	%eax,%ebx
	imul	%eax,%esp
	imul	%eax,%ebp
	imul	%eax,%esi
	imul	%eax,%edi
	imul	%eax,%r8d
	imul	%eax,%r9d
	imul	%eax,%r10d
	imul	%eax,%r11d
	imul	%eax,%r12d
	imul	%eax,%r13d
	imul	%eax,%r14d
	imul	%eax,%r15d
	nop
	imul	%ecx,%eax
	imul	%ecx,%ecx
	imul	%ecx,%edx
	imul	%ecx,%ebx
	imul	%ecx,%esp
	imul	%ecx,%ebp
	imul	%ecx,%esi
	imul	%ecx,%edi
	imul	%ecx,%r8d
	imul	%ecx,%r9d
	imul	%ecx,%r10d
	imul	%ecx,%r11d
	imul	%ecx,%r12d
	imul	%ecx,%r13d
	imul	%ecx,%r14d
	imul	%ecx,%r15d
	nop
	imul	%edx,%eax
	imul	%edx,%ecx
	imul	%edx,%edx
	imul	%edx,%ebx
	imul	%edx,%esp
	imul	%edx,%ebp
	imul	%edx,%esi
	imul	%edx,%edi
	imul	%edx,%r8d
	imul	%edx,%r9d
	imul	%edx,%r10d
	imul	%edx,%r11d
	imul	%edx,%r12d
	imul	%edx,%r13d
	imul	%edx,%r14d
	imul	%edx,%r15d
	nop
	imul	%ebx,%eax
	imul	%ebx,%ecx
	imul	%ebx,%edx
	imul	%ebx,%ebx
	imul	%ebx,%esp
	imul	%ebx,%ebp
	imul	%ebx,%esi
	imul	%ebx,%edi
	imul	%ebx,%r8d
	imul	%ebx,%r9d
	imul	%ebx,%r10d
	imul	%ebx,%r11d
	imul	%ebx,%r12d
	imul	%ebx,%r13d
	imul	%ebx,%r14d
	imul	%ebx,%r15d
	nop
	imul	%esp,%eax
	imul	%esp,%ecx
	imul	%esp,%edx
	imul	%esp,%ebx
	imul	%esp,%esp
	imul	%esp,%ebp
	imul	%esp,%esi
	imul	%esp,%edi
	imul	%esp,%r8d
	imul	%esp,%r9d
	imul	%esp,%r10d
	imul	%esp,%r11d
	imul	%esp,%r12d
	imul	%esp,%r13d
	imul	%esp,%r14d
	imul	%esp,%r15d
	nop
	imul	%ebp,%eax
	imul	%ebp,%ecx
	imul	%ebp,%edx
	imul	%ebp,%ebx
	imul	%ebp,%esp
	imul	%ebp,%ebp
	imul	%ebp,%esi
	imul	%ebp,%edi
	imul	%ebp,%r8d
	imul	%ebp,%r9d
	imul	%ebp,%r10d
	imul	%ebp,%r11d
	imul	%ebp,%r12d
	imul	%ebp,%r13d
	imul	%ebp,%r14d
	imul	%ebp,%r15d
	nop
	imul	%esi,%eax
	imul	%esi,%ecx
	imul	%esi,%edx
	imul	%esi,%ebx
	imul	%esi,%esp
	imul	%esi,%ebp
	imul	%esi,%esi
	imul	%esi,%edi
	imul	%esi,%r8d
	imul	%esi,%r9d
	imul	%esi,%r10d
	imul	%esi,%r11d
	imul	%esi,%r12d
	imul	%esi,%r13d
	imul	%esi,%r14d
	imul	%esi,%r15d
	nop
	imul	%edi,%eax
	imul	%edi,%ecx
	imul	%edi,%edx
	imul	%edi,%ebx
	imul	%edi,%esp
	imul	%edi,%ebp
	imul	%edi,%esi
	imul	%edi,%edi
	imul	%edi,%r8d
	imul	%edi,%r9d
	imul	%edi,%r10d
	imul	%edi,%r11d
	imul	%edi,%r12d
	imul	%edi,%r13d
	imul	%edi,%r14d
	imul	%edi,%r15d
	nop
	imul	%r8d, %eax
	imul	%r8d, %ecx
	imul	%r8d, %edx
	imul	%r8d, %ebx
	imul	%r8d, %esp
	imul	%r8d, %ebp
	imul	%r8d, %esi
	imul	%r8d, %edi
	imul	%r8d, %r8d
	imul	%r8d, %r9d
	imul	%r8d, %r10d
	imul	%r8d, %r11d
	imul	%r8d, %r12d
	imul	%r8d, %r13d
	imul	%r8d, %r14d
	imul	%r8d, %r15d
	nop
	imul	%r9d, %eax
	imul	%r9d, %ecx
	imul	%r9d, %edx
	imul	%r9d, %ebx
	imul	%r9d, %esp
	imul	%r9d, %ebp
	imul	%r9d, %esi
	imul	%r9d, %edi
	imul	%r9d, %r8d
	imul	%r9d, %r9d
	imul	%r9d, %r10d
	imul	%r9d, %r11d
	imul	%r9d, %r12d
	imul	%r9d, %r13d
	imul	%r9d, %r14d
	imul	%r9d, %r15d
	nop
	imul	%r10d,%eax
	imul	%r10d,%ecx
	imul	%r10d,%edx
	imul	%r10d,%ebx
	imul	%r10d,%esp
	imul	%r10d,%ebp
	imul	%r10d,%esi
	imul	%r10d,%edi
	imul	%r10d,%r8d
	imul	%r10d,%r9d
	imul	%r10d,%r10d
	imul	%r10d,%r11d
	imul	%r10d,%r12d
	imul	%r10d,%r13d
	imul	%r10d,%r14d
	imul	%r10d,%r15d
	nop
	imul	%r11d,%eax
	imul	%r11d,%ecx
	imul	%r11d,%edx
	imul	%r11d,%ebx
	imul	%r11d,%esp
	imul	%r11d,%ebp
	imul	%r11d,%esi
	imul	%r11d,%edi
	imul	%r11d,%r8d
	imul	%r11d,%r9d
	imul	%r11d,%r10d
	imul	%r11d,%r11d
	imul	%r11d,%r12d
	imul	%r11d,%r13d
	imul	%r11d,%r14d
	imul	%r11d,%r15d
	nop
	imul	%r12d,%eax
	imul	%r12d,%ecx
	imul	%r12d,%edx
	imul	%r12d,%ebx
	imul	%r12d,%esp
	imul	%r12d,%ebp
	imul	%r12d,%esi
	imul	%r12d,%edi
	imul	%r12d,%r8d
	imul	%r12d,%r9d
	imul	%r12d,%r10d
	imul	%r12d,%r11d
	imul	%r12d,%r12d
	imul	%r12d,%r13d
	imul	%r12d,%r14d
	imul	%r12d,%r15d
	nop
	imul	%r13d,%eax
	imul	%r13d,%ecx
	imul	%r13d,%edx
	imul	%r13d,%ebx
	imul	%r13d,%esp
	imul	%r13d,%ebp
	imul	%r13d,%esi
	imul	%r13d,%edi
	imul	%r13d,%r8d
	imul	%r13d,%r9d
	imul	%r13d,%r10d
	imul	%r13d,%r11d
	imul	%r13d,%r12d
	imul	%r13d,%r13d
	imul	%r13d,%r14d
	imul	%r13d,%r15d
	nop
	imul	%r14d,%eax
	imul	%r14d,%ecx
	imul	%r14d,%edx
	imul	%r14d,%ebx
	imul	%r14d,%esp
	imul	%r14d,%ebp
	imul	%r14d,%esi
	imul	%r14d,%edi
	imul	%r14d,%r8d
	imul	%r14d,%r9d
	imul	%r14d,%r10d
	imul	%r14d,%r11d
	imul	%r14d,%r12d
	imul	%r14d,%r13d
	imul	%r14d,%r14d
	imul	%r14d,%r15d
	nop
	imul	%r15d,%eax
	imul	%r15d,%ecx
	imul	%r15d,%edx
	imul	%r15d,%ebx
	imul	%r15d,%esp
	imul	%r15d,%ebp
	imul	%r15d,%esi
	imul	%r15d,%edi
	imul	%r15d,%r8d
	imul	%r15d,%r9d
	imul	%r15d,%r10d
	imul	%r15d,%r11d
	imul	%r15d,%r12d
	imul	%r15d,%r13d
	imul	%r15d,%r14d
	imul	%r15d,%r15d
        nop
        nop
        // mem32 *= reg32
	imul	%eax,(%rax)
	imul	%eax,(%rcx)
	imul	%eax,(%rdx)
	imul	%eax,(%rbx)
	imul	%eax,(%rsp)
	imul	%eax,(%rbp)
	imul	%eax,(%rsi)
	imul	%eax,(%rdi)
	imul	%eax,(%r8)
	imul	%eax,(%r9)
	imul	%eax,(%r10)
	imul	%eax,(%r11)
	imul	%eax,(%r12)
	imul	%eax,(%r13)
	imul	%eax,(%r14)
	imul	%eax,(%r15)
	nop
	imul	%ecx,(%rax)
	imul	%ecx,(%rcx)
	imul	%ecx,(%rdx)
	imul	%ecx,(%rbx)
	imul	%ecx,(%rsp)
	imul	%ecx,(%rbp)
	imul	%ecx,(%rsi)
	imul	%ecx,(%rdi)
	imul	%ecx,(%r8)
	imul	%ecx,(%r9)
	imul	%ecx,(%r10)
	imul	%ecx,(%r11)
	imul	%ecx,(%r12)
	imul	%ecx,(%r13)
	imul	%ecx,(%r14)
	imul	%ecx,(%r15)
	nop
	imul	%edx,(%rax)
	imul	%edx,(%rcx)
	imul	%edx,(%rdx)
	imul	%edx,(%rbx)
	imul	%edx,(%rsp)
	imul	%edx,(%rbp)
	imul	%edx,(%rsi)
	imul	%edx,(%rdi)
	imul	%edx,(%r8)
	imul	%edx,(%r9)
	imul	%edx,(%r10)
	imul	%edx,(%r11)
	imul	%edx,(%r12)
	imul	%edx,(%r13)
	imul	%edx,(%r14)
	imul	%edx,(%r15)
	nop
	imul	%ebx,(%rax)
	imul	%ebx,(%rcx)
	imul	%ebx,(%rdx)
	imul	%ebx,(%rbx)
	imul	%ebx,(%rsp)
	imul	%ebx,(%rbp)
	imul	%ebx,(%rsi)
	imul	%ebx,(%rdi)
	imul	%ebx,(%r8)
	imul	%ebx,(%r9)
	imul	%ebx,(%r10)
	imul	%ebx,(%r11)
	imul	%ebx,(%r12)
	imul	%ebx,(%r13)
	imul	%ebx,(%r14)
	imul	%ebx,(%r15)
	nop
	imul	%esp,(%rax)
	imul	%esp,(%rcx)
	imul	%esp,(%rdx)
	imul	%esp,(%rbx)
	imul	%esp,(%rsp)
	imul	%esp,(%rbp)
	imul	%esp,(%rsi)
	imul	%esp,(%rdi)
	imul	%esp,(%r8)
	imul	%esp,(%r9)
	imul	%esp,(%r10)
	imul	%esp,(%r11)
	imul	%esp,(%r12)
	imul	%esp,(%r13)
	imul	%esp,(%r14)
	imul	%esp,(%r15)
	nop
	imul	%ebp,(%rax)
	imul	%ebp,(%rcx)
	imul	%ebp,(%rdx)
	imul	%ebp,(%rbx)
	imul	%ebp,(%rsp)
	imul	%ebp,(%rbp)
	imul	%ebp,(%rsi)
	imul	%ebp,(%rdi)
	imul	%ebp,(%r8)
	imul	%ebp,(%r9)
	imul	%ebp,(%r10)
	imul	%ebp,(%r11)
	imul	%ebp,(%r12)
	imul	%ebp,(%r13)
	imul	%ebp,(%r14)
	imul	%ebp,(%r15)
	nop
	imul	%esi,(%rax)
	imul	%esi,(%rcx)
	imul	%esi,(%rdx)
	imul	%esi,(%rbx)
	imul	%esi,(%rsp)
	imul	%esi,(%rbp)
	imul	%esi,(%rsi)
	imul	%esi,(%rdi)
	imul	%esi,(%r8)
	imul	%esi,(%r9)
	imul	%esi,(%r10)
	imul	%esi,(%r11)
	imul	%esi,(%r12)
	imul	%esi,(%r13)
	imul	%esi,(%r14)
	imul	%esi,(%r15)
	nop
	imul	%edi,(%rax)
	imul	%edi,(%rcx)
	imul	%edi,(%rdx)
	imul	%edi,(%rbx)
	imul	%edi,(%rsp)
	imul	%edi,(%rbp)
	imul	%edi,(%rsi)
	imul	%edi,(%rdi)
	imul	%edi,(%r8)
	imul	%edi,(%r9)
	imul	%edi,(%r10)
	imul	%edi,(%r11)
	imul	%edi,(%r12)
	imul	%edi,(%r13)
	imul	%edi,(%r14)
	imul	%edi,(%r15)
	nop
	imul	%r8d, (%rax)
	imul	%r8d, (%rcx)
	imul	%r8d, (%rdx)
	imul	%r8d, (%rbx)
	imul	%r8d, (%rsp)
	imul	%r8d, (%rbp)
	imul	%r8d, (%rsi)
	imul	%r8d, (%rdi)
	imul	%r8d, (%r8)
	imul	%r8d, (%r9)
	imul	%r8d, (%r10)
	imul	%r8d, (%r11)
	imul	%r8d, (%r12)
	imul	%r8d, (%r13)
	imul	%r8d, (%r14)
	imul	%r8d, (%r15)
	nop
	imul	%r9d, (%rax)
	imul	%r9d, (%rcx)
	imul	%r9d, (%rdx)
	imul	%r9d, (%rbx)
	imul	%r9d, (%rsp)
	imul	%r9d, (%rbp)
	imul	%r9d, (%rsi)
	imul	%r9d, (%rdi)
	imul	%r9d, (%r8)
	imul	%r9d, (%r9)
	imul	%r9d, (%r10)
	imul	%r9d, (%r11)
	imul	%r9d, (%r12)
	imul	%r9d, (%r13)
	imul	%r9d, (%r14)
	imul	%r9d, (%r15)
	nop
	imul	%r10d,(%rax)
	imul	%r10d,(%rcx)
	imul	%r10d,(%rdx)
	imul	%r10d,(%rbx)
	imul	%r10d,(%rsp)
	imul	%r10d,(%rbp)
	imul	%r10d,(%rsi)
	imul	%r10d,(%rdi)
	imul	%r10d,(%r8)
	imul	%r10d,(%r9)
	imul	%r10d,(%r10)
	imul	%r10d,(%r11)
	imul	%r10d,(%r12)
	imul	%r10d,(%r13)
	imul	%r10d,(%r14)
	imul	%r10d,(%r15)
	nop
	imul	%r11d,(%rax)
	imul	%r11d,(%rcx)
	imul	%r11d,(%rdx)
	imul	%r11d,(%rbx)
	imul	%r11d,(%rsp)
	imul	%r11d,(%rbp)
	imul	%r11d,(%rsi)
	imul	%r11d,(%rdi)
	imul	%r11d,(%r8)
	imul	%r11d,(%r9)
	imul	%r11d,(%r10)
	imul	%r11d,(%r11)
	imul	%r11d,(%r12)
	imul	%r11d,(%r13)
	imul	%r11d,(%r14)
	imul	%r11d,(%r15)
	nop
	imul	%r12d,(%rax)
	imul	%r12d,(%rcx)
	imul	%r12d,(%rdx)
	imul	%r12d,(%rbx)
	imul	%r12d,(%rsp)
	imul	%r12d,(%rbp)
	imul	%r12d,(%rsi)
	imul	%r12d,(%rdi)
	imul	%r12d,(%r8)
	imul	%r12d,(%r9)
	imul	%r12d,(%r10)
	imul	%r12d,(%r11)
	imul	%r12d,(%r12)
	imul	%r12d,(%r13)
	imul	%r12d,(%r14)
	imul	%r12d,(%r15)
	nop
	imul	%r13d,(%rax)
	imul	%r13d,(%rcx)
	imul	%r13d,(%rdx)
	imul	%r13d,(%rbx)
	imul	%r13d,(%rsp)
	imul	%r13d,(%rbp)
	imul	%r13d,(%rsi)
	imul	%r13d,(%rdi)
	imul	%r13d,(%r8)
	imul	%r13d,(%r9)
	imul	%r13d,(%r10)
	imul	%r13d,(%r11)
	imul	%r13d,(%r12)
	imul	%r13d,(%r13)
	imul	%r13d,(%r14)
	imul	%r13d,(%r15)
	nop
	imul	%r14d,(%rax)
	imul	%r14d,(%rcx)
	imul	%r14d,(%rdx)
	imul	%r14d,(%rbx)
	imul	%r14d,(%rsp)
	imul	%r14d,(%rbp)
	imul	%r14d,(%rsi)
	imul	%r14d,(%rdi)
	imul	%r14d,(%r8)
	imul	%r14d,(%r9)
	imul	%r14d,(%r10)
	imul	%r14d,(%r11)
	imul	%r14d,(%r12)
	imul	%r14d,(%r13)
	imul	%r14d,(%r14)
	imul	%r14d,(%r15)
	nop
	imul	%r15d,(%rax)
	imul	%r15d,(%rcx)
	imul	%r15d,(%rdx)
	imul	%r15d,(%rbx)
	imul	%r15d,(%rsp)
	imul	%r15d,(%rbp)
	imul	%r15d,(%rsi)
	imul	%r15d,(%rdi)
	imul	%r15d,(%r8)
	imul	%r15d,(%r9)
	imul	%r15d,(%r10)
	imul	%r15d,(%r11)
	imul	%r15d,(%r12)
	imul	%r15d,(%r13)
	imul	%r15d,(%r14)
	imul	%r15d,(%r15)
        ret
	.cfi_endproc

        
        .p2align 4,,15
	.globl	Imul64
	.type	Imul64, @function
Imul64:
	.cfi_startproc
        // mem64 += reg64
	imul	%rax,(%rax)
	imul	%rax,(%rcx)
	imul	%rax,(%rdx)
	imul	%rax,(%rbx)
	imul	%rax,(%rsp)
	imul	%rax,(%rbp)
	imul	%rax,(%rsi)
	imul	%rax,(%rdi)
	imul	%rax,(%r8)
	imul	%rax,(%r9)
	imul	%rax,(%r10)
	imul	%rax,(%r11)
	imul	%rax,(%r12)
	imul	%rax,(%r13)
	imul	%rax,(%r14)
	imul	%rax,(%r15)
	nop
	imul	%rcx,(%rax)
	imul	%rcx,(%rcx)
	imul	%rcx,(%rdx)
	imul	%rcx,(%rbx)
	imul	%rcx,(%rsp)
	imul	%rcx,(%rbp)
	imul	%rcx,(%rsi)
	imul	%rcx,(%rdi)
	imul	%rcx,(%r8)
	imul	%rcx,(%r9)
	imul	%rcx,(%r10)
	imul	%rcx,(%r11)
	imul	%rcx,(%r12)
	imul	%rcx,(%r13)
	imul	%rcx,(%r14)
	imul	%rcx,(%r15)
	nop
	imul	%rdx,(%rax)
	imul	%rdx,(%rcx)
	imul	%rdx,(%rdx)
	imul	%rdx,(%rbx)
	imul	%rdx,(%rsp)
	imul	%rdx,(%rbp)
	imul	%rdx,(%rsi)
	imul	%rdx,(%rdi)
	imul	%rdx,(%r8)
	imul	%rdx,(%r9)
	imul	%rdx,(%r10)
	imul	%rdx,(%r11)
	imul	%rdx,(%r12)
	imul	%rdx,(%r13)
	imul	%rdx,(%r14)
	imul	%rdx,(%r15)
	nop
	imul	%rbx,(%rax)
	imul	%rbx,(%rcx)
	imul	%rbx,(%rdx)
	imul	%rbx,(%rbx)
	imul	%rbx,(%rsp)
	imul	%rbx,(%rbp)
	imul	%rbx,(%rsi)
	imul	%rbx,(%rdi)
	imul	%rbx,(%r8)
	imul	%rbx,(%r9)
	imul	%rbx,(%r10)
	imul	%rbx,(%r11)
	imul	%rbx,(%r12)
	imul	%rbx,(%r13)
	imul	%rbx,(%r14)
	imul	%rbx,(%r15)
	nop
	imul	%rsp,(%rax)
	imul	%rsp,(%rcx)
	imul	%rsp,(%rdx)
	imul	%rsp,(%rbx)
	imul	%rsp,(%rsp)
	imul	%rsp,(%rbp)
	imul	%rsp,(%rsi)
	imul	%rsp,(%rdi)
	imul	%rsp,(%r8)
	imul	%rsp,(%r9)
	imul	%rsp,(%r10)
	imul	%rsp,(%r11)
	imul	%rsp,(%r12)
	imul	%rsp,(%r13)
	imul	%rsp,(%r14)
	imul	%rsp,(%r15)
	nop
	imul	%rbp,(%rax)
	imul	%rbp,(%rcx)
	imul	%rbp,(%rdx)
	imul	%rbp,(%rbx)
	imul	%rbp,(%rsp)
	imul	%rbp,(%rbp)
	imul	%rbp,(%rsi)
	imul	%rbp,(%rdi)
	imul	%rbp,(%r8)
	imul	%rbp,(%r9)
	imul	%rbp,(%r10)
	imul	%rbp,(%r11)
	imul	%rbp,(%r12)
	imul	%rbp,(%r13)
	imul	%rbp,(%r14)
	imul	%rbp,(%r15)
	nop
	imul	%rsi,(%rax)
	imul	%rsi,(%rcx)
	imul	%rsi,(%rdx)
	imul	%rsi,(%rbx)
	imul	%rsi,(%rsp)
	imul	%rsi,(%rbp)
	imul	%rsi,(%rsi)
	imul	%rsi,(%rdi)
	imul	%rsi,(%r8)
	imul	%rsi,(%r9)
	imul	%rsi,(%r10)
	imul	%rsi,(%r11)
	imul	%rsi,(%r12)
	imul	%rsi,(%r13)
	imul	%rsi,(%r14)
	imul	%rsi,(%r15)
	nop
	imul	%rdi,(%rax)
	imul	%rdi,(%rcx)
	imul	%rdi,(%rdx)
	imul	%rdi,(%rbx)
	imul	%rdi,(%rsp)
	imul	%rdi,(%rbp)
	imul	%rdi,(%rsi)
	imul	%rdi,(%rdi)
	imul	%rdi,(%r8)
	imul	%rdi,(%r9)
	imul	%rdi,(%r10)
	imul	%rdi,(%r11)
	imul	%rdi,(%r12)
	imul	%rdi,(%r13)
	imul	%rdi,(%r14)
	imul	%rdi,(%r15)
	nop
	imul	%r8 ,(%rax)
	imul	%r8 ,(%rcx)
	imul	%r8 ,(%rdx)
	imul	%r8 ,(%rbx)
	imul	%r8 ,(%rsp)
	imul	%r8 ,(%rbp)
	imul	%r8 ,(%rsi)
	imul	%r8 ,(%rdi)
	imul	%r8 ,(%r8)
	imul	%r8 ,(%r9)
	imul	%r8 ,(%r10)
	imul	%r8 ,(%r11)
	imul	%r8 ,(%r12)
	imul	%r8 ,(%r13)
	imul	%r8 ,(%r14)
	imul	%r8 ,(%r15)
	nop
	imul	%r9 ,(%rax)
	imul	%r9 ,(%rcx)
	imul	%r9 ,(%rdx)
	imul	%r9 ,(%rbx)
	imul	%r9 ,(%rsp)
	imul	%r9 ,(%rbp)
	imul	%r9 ,(%rsi)
	imul	%r9 ,(%rdi)
	imul	%r9 ,(%r8)
	imul	%r9 ,(%r9)
	imul	%r9 ,(%r10)
	imul	%r9 ,(%r11)
	imul	%r9 ,(%r12)
	imul	%r9 ,(%r13)
	imul	%r9 ,(%r14)
	imul	%r9 ,(%r15)
	nop
	imul	%r10,(%rax)
	imul	%r10,(%rcx)
	imul	%r10,(%rdx)
	imul	%r10,(%rbx)
	imul	%r10,(%rsp)
	imul	%r10,(%rbp)
	imul	%r10,(%rsi)
	imul	%r10,(%rdi)
	imul	%r10,(%r8)
	imul	%r10,(%r9)
	imul	%r10,(%r10)
	imul	%r10,(%r11)
	imul	%r10,(%r12)
	imul	%r10,(%r13)
	imul	%r10,(%r14)
	imul	%r10,(%r15)
	nop
	imul	%r11,(%rax)
	imul	%r11,(%rcx)
	imul	%r11,(%rdx)
	imul	%r11,(%rbx)
	imul	%r11,(%rsp)
	imul	%r11,(%rbp)
	imul	%r11,(%rsi)
	imul	%r11,(%rdi)
	imul	%r11,(%r8)
	imul	%r11,(%r9)
	imul	%r11,(%r10)
	imul	%r11,(%r11)
	imul	%r11,(%r12)
	imul	%r11,(%r13)
	imul	%r11,(%r14)
	imul	%r11,(%r15)
	nop
	imul	%r12,(%rax)
	imul	%r12,(%rcx)
	imul	%r12,(%rdx)
	imul	%r12,(%rbx)
	imul	%r12,(%rsp)
	imul	%r12,(%rbp)
	imul	%r12,(%rsi)
	imul	%r12,(%rdi)
	imul	%r12,(%r8)
	imul	%r12,(%r9)
	imul	%r12,(%r10)
	imul	%r12,(%r11)
	imul	%r12,(%r12)
	imul	%r12,(%r13)
	imul	%r12,(%r14)
	imul	%r12,(%r15)
	nop
	imul	%r13,(%rax)
	imul	%r13,(%rcx)
	imul	%r13,(%rdx)
	imul	%r13,(%rbx)
	imul	%r13,(%rsp)
	imul	%r13,(%rbp)
	imul	%r13,(%rsi)
	imul	%r13,(%rdi)
	imul	%r13,(%r8)
	imul	%r13,(%r9)
	imul	%r13,(%r10)
	imul	%r13,(%r11)
	imul	%r13,(%r12)
	imul	%r13,(%r13)
	imul	%r13,(%r14)
	imul	%r13,(%r15)
	nop
	imul	%r14,(%rax)
	imul	%r14,(%rcx)
	imul	%r14,(%rdx)
	imul	%r14,(%rbx)
	imul	%r14,(%rsp)
	imul	%r14,(%rbp)
	imul	%r14,(%rsi)
	imul	%r14,(%rdi)
	imul	%r14,(%r8)
	imul	%r14,(%r9)
	imul	%r14,(%r10)
	imul	%r14,(%r11)
	imul	%r14,(%r12)
	imul	%r14,(%r13)
	imul	%r14,(%r14)
	imul	%r14,(%r15)
	nop
	imul	%r15,(%rax)
	imul	%r15,(%rcx)
	imul	%r15,(%rdx)
	imul	%r15,(%rbx)
	imul	%r15,(%rsp)
	imul	%r15,(%rbp)
	imul	%r15,(%rsi)
	imul	%r15,(%rdi)
	imul	%r15,(%r8)
	imul	%r15,(%r9)
	imul	%r15,(%r10)
	imul	%r15,(%r11)
	imul	%r15,(%r12)
	imul	%r15,(%r13)
	imul	%r15,(%r14)
	imul	%r15,(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	ImulMem8Reg
	.type	ImulMem8Reg, @function
ImulMem8Reg:
	.cfi_startproc
	imul	%rax,0x7f(%rax)
	imul	%rax,0x7f(%rcx)
	imul	%rax,0x7f(%rdx)
	imul	%rax,0x7f(%rbx)
	imul	%rax,0x7f(%rsp)
	imul	%rax,0x7f(%rbp)
	imul	%rax,0x7f(%rsi)
	imul	%rax,0x7f(%rdi)
	imul	%rax,0x7f(%r8)
	imul	%rax,0x7f(%r9)
	imul	%rax,0x7f(%r10)
	imul	%rax,0x7f(%r11)
	imul	%rax,0x7f(%r12)
	imul	%rax,0x7f(%r13)
	imul	%rax,0x7f(%r14)
	imul	%rax,0x7f(%r15)
	nop
	imul	%rcx,0x7f(%rax)
	imul	%rcx,0x7f(%rcx)
	imul	%rcx,0x7f(%rdx)
	imul	%rcx,0x7f(%rbx)
	imul	%rcx,0x7f(%rsp)
	imul	%rcx,0x7f(%rbp)
	imul	%rcx,0x7f(%rsi)
	imul	%rcx,0x7f(%rdi)
	imul	%rcx,0x7f(%r8)
	imul	%rcx,0x7f(%r9)
	imul	%rcx,0x7f(%r10)
	imul	%rcx,0x7f(%r11)
	imul	%rcx,0x7f(%r12)
	imul	%rcx,0x7f(%r13)
	imul	%rcx,0x7f(%r14)
	imul	%rcx,0x7f(%r15)
	nop
	imul	%rdx,0x7f(%rax)
	imul	%rdx,0x7f(%rcx)
	imul	%rdx,0x7f(%rdx)
	imul	%rdx,0x7f(%rbx)
	imul	%rdx,0x7f(%rsp)
	imul	%rdx,0x7f(%rbp)
	imul	%rdx,0x7f(%rsi)
	imul	%rdx,0x7f(%rdi)
	imul	%rdx,0x7f(%r8)
	imul	%rdx,0x7f(%r9)
	imul	%rdx,0x7f(%r10)
	imul	%rdx,0x7f(%r11)
	imul	%rdx,0x7f(%r12)
	imul	%rdx,0x7f(%r13)
	imul	%rdx,0x7f(%r14)
	imul	%rdx,0x7f(%r15)
	nop
	imul	%rbx,0x7f(%rax)
	imul	%rbx,0x7f(%rcx)
	imul	%rbx,0x7f(%rdx)
	imul	%rbx,0x7f(%rbx)
	imul	%rbx,0x7f(%rsp)
	imul	%rbx,0x7f(%rbp)
	imul	%rbx,0x7f(%rsi)
	imul	%rbx,0x7f(%rdi)
	imul	%rbx,0x7f(%r8)
	imul	%rbx,0x7f(%r9)
	imul	%rbx,0x7f(%r10)
	imul	%rbx,0x7f(%r11)
	imul	%rbx,0x7f(%r12)
	imul	%rbx,0x7f(%r13)
	imul	%rbx,0x7f(%r14)
	imul	%rbx,0x7f(%r15)
	nop
	imul	%rsp,0x7f(%rax)
	imul	%rsp,0x7f(%rcx)
	imul	%rsp,0x7f(%rdx)
	imul	%rsp,0x7f(%rbx)
	imul	%rsp,0x7f(%rsp)
	imul	%rsp,0x7f(%rbp)
	imul	%rsp,0x7f(%rsi)
	imul	%rsp,0x7f(%rdi)
	imul	%rsp,0x7f(%r8)
	imul	%rsp,0x7f(%r9)
	imul	%rsp,0x7f(%r10)
	imul	%rsp,0x7f(%r11)
	imul	%rsp,0x7f(%r12)
	imul	%rsp,0x7f(%r13)
	imul	%rsp,0x7f(%r14)
	imul	%rsp,0x7f(%r15)
	nop
	imul	%rbp,0x7f(%rax)
	imul	%rbp,0x7f(%rcx)
	imul	%rbp,0x7f(%rdx)
	imul	%rbp,0x7f(%rbx)
	imul	%rbp,0x7f(%rsp)
	imul	%rbp,0x7f(%rbp)
	imul	%rbp,0x7f(%rsi)
	imul	%rbp,0x7f(%rdi)
	imul	%rbp,0x7f(%r8)
	imul	%rbp,0x7f(%r9)
	imul	%rbp,0x7f(%r10)
	imul	%rbp,0x7f(%r11)
	imul	%rbp,0x7f(%r12)
	imul	%rbp,0x7f(%r13)
	imul	%rbp,0x7f(%r14)
	imul	%rbp,0x7f(%r15)
	nop
	imul	%rsi,0x7f(%rax)
	imul	%rsi,0x7f(%rcx)
	imul	%rsi,0x7f(%rdx)
	imul	%rsi,0x7f(%rbx)
	imul	%rsi,0x7f(%rsp)
	imul	%rsi,0x7f(%rbp)
	imul	%rsi,0x7f(%rsi)
	imul	%rsi,0x7f(%rdi)
	imul	%rsi,0x7f(%r8)
	imul	%rsi,0x7f(%r9)
	imul	%rsi,0x7f(%r10)
	imul	%rsi,0x7f(%r11)
	imul	%rsi,0x7f(%r12)
	imul	%rsi,0x7f(%r13)
	imul	%rsi,0x7f(%r14)
	imul	%rsi,0x7f(%r15)
	nop
	imul	%rdi,0x7f(%rax)
	imul	%rdi,0x7f(%rcx)
	imul	%rdi,0x7f(%rdx)
	imul	%rdi,0x7f(%rbx)
	imul	%rdi,0x7f(%rsp)
	imul	%rdi,0x7f(%rbp)
	imul	%rdi,0x7f(%rsi)
	imul	%rdi,0x7f(%rdi)
	imul	%rdi,0x7f(%r8)
	imul	%rdi,0x7f(%r9)
	imul	%rdi,0x7f(%r10)
	imul	%rdi,0x7f(%r11)
	imul	%rdi,0x7f(%r12)
	imul	%rdi,0x7f(%r13)
	imul	%rdi,0x7f(%r14)
	imul	%rdi,0x7f(%r15)
	nop
	imul	%r8 ,0x7f(%rax)
	imul	%r8 ,0x7f(%rcx)
	imul	%r8 ,0x7f(%rdx)
	imul	%r8 ,0x7f(%rbx)
	imul	%r8 ,0x7f(%rsp)
	imul	%r8 ,0x7f(%rbp)
	imul	%r8 ,0x7f(%rsi)
	imul	%r8 ,0x7f(%rdi)
	imul	%r8 ,0x7f(%r8)
	imul	%r8 ,0x7f(%r9)
	imul	%r8 ,0x7f(%r10)
	imul	%r8 ,0x7f(%r11)
	imul	%r8 ,0x7f(%r12)
	imul	%r8 ,0x7f(%r13)
	imul	%r8 ,0x7f(%r14)
	imul	%r8 ,0x7f(%r15)
	nop
	imul	%r9 ,0x7f(%rax)
	imul	%r9 ,0x7f(%rcx)
	imul	%r9 ,0x7f(%rdx)
	imul	%r9 ,0x7f(%rbx)
	imul	%r9 ,0x7f(%rsp)
	imul	%r9 ,0x7f(%rbp)
	imul	%r9 ,0x7f(%rsi)
	imul	%r9 ,0x7f(%rdi)
	imul	%r9 ,0x7f(%r8)
	imul	%r9 ,0x7f(%r9)
	imul	%r9 ,0x7f(%r10)
	imul	%r9 ,0x7f(%r11)
	imul	%r9 ,0x7f(%r12)
	imul	%r9 ,0x7f(%r13)
	imul	%r9 ,0x7f(%r14)
	imul	%r9 ,0x7f(%r15)
	nop
	imul	%r10,0x7f(%rax)
	imul	%r10,0x7f(%rcx)
	imul	%r10,0x7f(%rdx)
	imul	%r10,0x7f(%rbx)
	imul	%r10,0x7f(%rsp)
	imul	%r10,0x7f(%rbp)
	imul	%r10,0x7f(%rsi)
	imul	%r10,0x7f(%rdi)
	imul	%r10,0x7f(%r8)
	imul	%r10,0x7f(%r9)
	imul	%r10,0x7f(%r10)
	imul	%r10,0x7f(%r11)
	imul	%r10,0x7f(%r12)
	imul	%r10,0x7f(%r13)
	imul	%r10,0x7f(%r14)
	imul	%r10,0x7f(%r15)
	nop
	imul	%r11,0x7f(%rax)
	imul	%r11,0x7f(%rcx)
	imul	%r11,0x7f(%rdx)
	imul	%r11,0x7f(%rbx)
	imul	%r11,0x7f(%rsp)
	imul	%r11,0x7f(%rbp)
	imul	%r11,0x7f(%rsi)
	imul	%r11,0x7f(%rdi)
	imul	%r11,0x7f(%r8)
	imul	%r11,0x7f(%r9)
	imul	%r11,0x7f(%r10)
	imul	%r11,0x7f(%r11)
	imul	%r11,0x7f(%r12)
	imul	%r11,0x7f(%r13)
	imul	%r11,0x7f(%r14)
	imul	%r11,0x7f(%r15)
	nop
	imul	%r12,0x7f(%rax)
	imul	%r12,0x7f(%rcx)
	imul	%r12,0x7f(%rdx)
	imul	%r12,0x7f(%rbx)
	imul	%r12,0x7f(%rsp)
	imul	%r12,0x7f(%rbp)
	imul	%r12,0x7f(%rsi)
	imul	%r12,0x7f(%rdi)
	imul	%r12,0x7f(%r8)
	imul	%r12,0x7f(%r9)
	imul	%r12,0x7f(%r10)
	imul	%r12,0x7f(%r11)
	imul	%r12,0x7f(%r12)
	imul	%r12,0x7f(%r13)
	imul	%r12,0x7f(%r14)
	imul	%r12,0x7f(%r15)
	nop
	imul	%r13,0x7f(%rax)
	imul	%r13,0x7f(%rcx)
	imul	%r13,0x7f(%rdx)
	imul	%r13,0x7f(%rbx)
	imul	%r13,0x7f(%rsp)
	imul	%r13,0x7f(%rbp)
	imul	%r13,0x7f(%rsi)
	imul	%r13,0x7f(%rdi)
	imul	%r13,0x7f(%r8)
	imul	%r13,0x7f(%r9)
	imul	%r13,0x7f(%r10)
	imul	%r13,0x7f(%r11)
	imul	%r13,0x7f(%r12)
	imul	%r13,0x7f(%r13)
	imul	%r13,0x7f(%r14)
	imul	%r13,0x7f(%r15)
	nop
	imul	%r14,0x7f(%rax)
	imul	%r14,0x7f(%rcx)
	imul	%r14,0x7f(%rdx)
	imul	%r14,0x7f(%rbx)
	imul	%r14,0x7f(%rsp)
	imul	%r14,0x7f(%rbp)
	imul	%r14,0x7f(%rsi)
	imul	%r14,0x7f(%rdi)
	imul	%r14,0x7f(%r8)
	imul	%r14,0x7f(%r9)
	imul	%r14,0x7f(%r10)
	imul	%r14,0x7f(%r11)
	imul	%r14,0x7f(%r12)
	imul	%r14,0x7f(%r13)
	imul	%r14,0x7f(%r14)
	imul	%r14,0x7f(%r15)
	nop
	imul	%r15,0x7f(%rax)
	imul	%r15,0x7f(%rcx)
	imul	%r15,0x7f(%rdx)
	imul	%r15,0x7f(%rbx)
	imul	%r15,0x7f(%rsp)
	imul	%r15,0x7f(%rbp)
	imul	%r15,0x7f(%rsi)
	imul	%r15,0x7f(%rdi)
	imul	%r15,0x7f(%r8)
	imul	%r15,0x7f(%r9)
	imul	%r15,0x7f(%r10)
	imul	%r15,0x7f(%r11)
	imul	%r15,0x7f(%r12)
	imul	%r15,0x7f(%r13)
	imul	%r15,0x7f(%r14)
	imul	%r15,0x7f(%r15)
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	ImulMem32Reg
	.type	ImulMem32Reg, @function
ImulMem32Reg:
	.cfi_startproc
	imul	%rax,0x7f563412(%rax)
	imul	%rax,0x7f563412(%rcx)
	imul	%rax,0x7f563412(%rdx)
	imul	%rax,0x7f563412(%rbx)
	imul	%rax,0x7f563412(%rsp)
	imul	%rax,0x7f563412(%rbp)
	imul	%rax,0x7f563412(%rsi)
	imul	%rax,0x7f563412(%rdi)
	imul	%rax,0x7f563412(%r8)
	imul	%rax,0x7f563412(%r9)
	imul	%rax,0x7f563412(%r10)
	imul	%rax,0x7f563412(%r11)
	imul	%rax,0x7f563412(%r12)
	imul	%rax,0x7f563412(%r13)
	imul	%rax,0x7f563412(%r14)
	imul	%rax,0x7f563412(%r15)
	nop
	imul	%rcx,0x7f563412(%rax)
	imul	%rcx,0x7f563412(%rcx)
	imul	%rcx,0x7f563412(%rdx)
	imul	%rcx,0x7f563412(%rbx)
	imul	%rcx,0x7f563412(%rsp)
	imul	%rcx,0x7f563412(%rbp)
	imul	%rcx,0x7f563412(%rsi)
	imul	%rcx,0x7f563412(%rdi)
	imul	%rcx,0x7f563412(%r8)
	imul	%rcx,0x7f563412(%r9)
	imul	%rcx,0x7f563412(%r10)
	imul	%rcx,0x7f563412(%r11)
	imul	%rcx,0x7f563412(%r12)
	imul	%rcx,0x7f563412(%r13)
	imul	%rcx,0x7f563412(%r14)
	imul	%rcx,0x7f563412(%r15)
	nop
	imul	%rdx,0x7f563412(%rax)
	imul	%rdx,0x7f563412(%rcx)
	imul	%rdx,0x7f563412(%rdx)
	imul	%rdx,0x7f563412(%rbx)
	imul	%rdx,0x7f563412(%rsp)
	imul	%rdx,0x7f563412(%rbp)
	imul	%rdx,0x7f563412(%rsi)
	imul	%rdx,0x7f563412(%rdi)
	imul	%rdx,0x7f563412(%r8)
	imul	%rdx,0x7f563412(%r9)
	imul	%rdx,0x7f563412(%r10)
	imul	%rdx,0x7f563412(%r11)
	imul	%rdx,0x7f563412(%r12)
	imul	%rdx,0x7f563412(%r13)
	imul	%rdx,0x7f563412(%r14)
	imul	%rdx,0x7f563412(%r15)
	nop
	imul	%rbx,0x7f563412(%rax)
	imul	%rbx,0x7f563412(%rcx)
	imul	%rbx,0x7f563412(%rdx)
	imul	%rbx,0x7f563412(%rbx)
	imul	%rbx,0x7f563412(%rsp)
	imul	%rbx,0x7f563412(%rbp)
	imul	%rbx,0x7f563412(%rsi)
	imul	%rbx,0x7f563412(%rdi)
	imul	%rbx,0x7f563412(%r8)
	imul	%rbx,0x7f563412(%r9)
	imul	%rbx,0x7f563412(%r10)
	imul	%rbx,0x7f563412(%r11)
	imul	%rbx,0x7f563412(%r12)
	imul	%rbx,0x7f563412(%r13)
	imul	%rbx,0x7f563412(%r14)
	imul	%rbx,0x7f563412(%r15)
	nop
	imul	%rsp,0x7f563412(%rax)
	imul	%rsp,0x7f563412(%rcx)
	imul	%rsp,0x7f563412(%rdx)
	imul	%rsp,0x7f563412(%rbx)
	imul	%rsp,0x7f563412(%rsp)
	imul	%rsp,0x7f563412(%rbp)
	imul	%rsp,0x7f563412(%rsi)
	imul	%rsp,0x7f563412(%rdi)
	imul	%rsp,0x7f563412(%r8)
	imul	%rsp,0x7f563412(%r9)
	imul	%rsp,0x7f563412(%r10)
	imul	%rsp,0x7f563412(%r11)
	imul	%rsp,0x7f563412(%r12)
	imul	%rsp,0x7f563412(%r13)
	imul	%rsp,0x7f563412(%r14)
	imul	%rsp,0x7f563412(%r15)
	nop
	imul	%rbp,0x7f563412(%rax)
	imul	%rbp,0x7f563412(%rcx)
	imul	%rbp,0x7f563412(%rdx)
	imul	%rbp,0x7f563412(%rbx)
	imul	%rbp,0x7f563412(%rsp)
	imul	%rbp,0x7f563412(%rbp)
	imul	%rbp,0x7f563412(%rsi)
	imul	%rbp,0x7f563412(%rdi)
	imul	%rbp,0x7f563412(%r8)
	imul	%rbp,0x7f563412(%r9)
	imul	%rbp,0x7f563412(%r10)
	imul	%rbp,0x7f563412(%r11)
	imul	%rbp,0x7f563412(%r12)
	imul	%rbp,0x7f563412(%r13)
	imul	%rbp,0x7f563412(%r14)
	imul	%rbp,0x7f563412(%r15)
	nop
	imul	%rsi,0x7f563412(%rax)
	imul	%rsi,0x7f563412(%rcx)
	imul	%rsi,0x7f563412(%rdx)
	imul	%rsi,0x7f563412(%rbx)
	imul	%rsi,0x7f563412(%rsp)
	imul	%rsi,0x7f563412(%rbp)
	imul	%rsi,0x7f563412(%rsi)
	imul	%rsi,0x7f563412(%rdi)
	imul	%rsi,0x7f563412(%r8)
	imul	%rsi,0x7f563412(%r9)
	imul	%rsi,0x7f563412(%r10)
	imul	%rsi,0x7f563412(%r11)
	imul	%rsi,0x7f563412(%r12)
	imul	%rsi,0x7f563412(%r13)
	imul	%rsi,0x7f563412(%r14)
	imul	%rsi,0x7f563412(%r15)
	nop
	imul	%rdi,0x7f563412(%rax)
	imul	%rdi,0x7f563412(%rcx)
	imul	%rdi,0x7f563412(%rdx)
	imul	%rdi,0x7f563412(%rbx)
	imul	%rdi,0x7f563412(%rsp)
	imul	%rdi,0x7f563412(%rbp)
	imul	%rdi,0x7f563412(%rsi)
	imul	%rdi,0x7f563412(%rdi)
	imul	%rdi,0x7f563412(%r8)
	imul	%rdi,0x7f563412(%r9)
	imul	%rdi,0x7f563412(%r10)
	imul	%rdi,0x7f563412(%r11)
	imul	%rdi,0x7f563412(%r12)
	imul	%rdi,0x7f563412(%r13)
	imul	%rdi,0x7f563412(%r14)
	imul	%rdi,0x7f563412(%r15)
	nop
	imul	%r8 ,0x7f563412(%rax)
	imul	%r8 ,0x7f563412(%rcx)
	imul	%r8 ,0x7f563412(%rdx)
	imul	%r8 ,0x7f563412(%rbx)
	imul	%r8 ,0x7f563412(%rsp)
	imul	%r8 ,0x7f563412(%rbp)
	imul	%r8 ,0x7f563412(%rsi)
	imul	%r8 ,0x7f563412(%rdi)
	imul	%r8 ,0x7f563412(%r8)
	imul	%r8 ,0x7f563412(%r9)
	imul	%r8 ,0x7f563412(%r10)
	imul	%r8 ,0x7f563412(%r11)
	imul	%r8 ,0x7f563412(%r12)
	imul	%r8 ,0x7f563412(%r13)
	imul	%r8 ,0x7f563412(%r14)
	imul	%r8 ,0x7f563412(%r15)
	nop
	imul	%r9 ,0x7f563412(%rax)
	imul	%r9 ,0x7f563412(%rcx)
	imul	%r9 ,0x7f563412(%rdx)
	imul	%r9 ,0x7f563412(%rbx)
	imul	%r9 ,0x7f563412(%rsp)
	imul	%r9 ,0x7f563412(%rbp)
	imul	%r9 ,0x7f563412(%rsi)
	imul	%r9 ,0x7f563412(%rdi)
	imul	%r9 ,0x7f563412(%r8)
	imul	%r9 ,0x7f563412(%r9)
	imul	%r9 ,0x7f563412(%r10)
	imul	%r9 ,0x7f563412(%r11)
	imul	%r9 ,0x7f563412(%r12)
	imul	%r9 ,0x7f563412(%r13)
	imul	%r9 ,0x7f563412(%r14)
	imul	%r9 ,0x7f563412(%r15)
	nop
	imul	%r10,0x7f563412(%rax)
	imul	%r10,0x7f563412(%rcx)
	imul	%r10,0x7f563412(%rdx)
	imul	%r10,0x7f563412(%rbx)
	imul	%r10,0x7f563412(%rsp)
	imul	%r10,0x7f563412(%rbp)
	imul	%r10,0x7f563412(%rsi)
	imul	%r10,0x7f563412(%rdi)
	imul	%r10,0x7f563412(%r8)
	imul	%r10,0x7f563412(%r9)
	imul	%r10,0x7f563412(%r10)
	imul	%r10,0x7f563412(%r11)
	imul	%r10,0x7f563412(%r12)
	imul	%r10,0x7f563412(%r13)
	imul	%r10,0x7f563412(%r14)
	imul	%r10,0x7f563412(%r15)
	nop
	imul	%r11,0x7f563412(%rax)
	imul	%r11,0x7f563412(%rcx)
	imul	%r11,0x7f563412(%rdx)
	imul	%r11,0x7f563412(%rbx)
	imul	%r11,0x7f563412(%rsp)
	imul	%r11,0x7f563412(%rbp)
	imul	%r11,0x7f563412(%rsi)
	imul	%r11,0x7f563412(%rdi)
	imul	%r11,0x7f563412(%r8)
	imul	%r11,0x7f563412(%r9)
	imul	%r11,0x7f563412(%r10)
	imul	%r11,0x7f563412(%r11)
	imul	%r11,0x7f563412(%r12)
	imul	%r11,0x7f563412(%r13)
	imul	%r11,0x7f563412(%r14)
	imul	%r11,0x7f563412(%r15)
	nop
	imul	%r12,0x7f563412(%rax)
	imul	%r12,0x7f563412(%rcx)
	imul	%r12,0x7f563412(%rdx)
	imul	%r12,0x7f563412(%rbx)
	imul	%r12,0x7f563412(%rsp)
	imul	%r12,0x7f563412(%rbp)
	imul	%r12,0x7f563412(%rsi)
	imul	%r12,0x7f563412(%rdi)
	imul	%r12,0x7f563412(%r8)
	imul	%r12,0x7f563412(%r9)
	imul	%r12,0x7f563412(%r10)
	imul	%r12,0x7f563412(%r11)
	imul	%r12,0x7f563412(%r12)
	imul	%r12,0x7f563412(%r13)
	imul	%r12,0x7f563412(%r14)
	imul	%r12,0x7f563412(%r15)
	nop
	imul	%r13,0x7f563412(%rax)
	imul	%r13,0x7f563412(%rcx)
	imul	%r13,0x7f563412(%rdx)
	imul	%r13,0x7f563412(%rbx)
	imul	%r13,0x7f563412(%rsp)
	imul	%r13,0x7f563412(%rbp)
	imul	%r13,0x7f563412(%rsi)
	imul	%r13,0x7f563412(%rdi)
	imul	%r13,0x7f563412(%r8)
	imul	%r13,0x7f563412(%r9)
	imul	%r13,0x7f563412(%r10)
	imul	%r13,0x7f563412(%r11)
	imul	%r13,0x7f563412(%r12)
	imul	%r13,0x7f563412(%r13)
	imul	%r13,0x7f563412(%r14)
	imul	%r13,0x7f563412(%r15)
	nop
	imul	%r14,0x7f563412(%rax)
	imul	%r14,0x7f563412(%rcx)
	imul	%r14,0x7f563412(%rdx)
	imul	%r14,0x7f563412(%rbx)
	imul	%r14,0x7f563412(%rsp)
	imul	%r14,0x7f563412(%rbp)
	imul	%r14,0x7f563412(%rsi)
	imul	%r14,0x7f563412(%rdi)
	imul	%r14,0x7f563412(%r8)
	imul	%r14,0x7f563412(%r9)
	imul	%r14,0x7f563412(%r10)
	imul	%r14,0x7f563412(%r11)
	imul	%r14,0x7f563412(%r12)
	imul	%r14,0x7f563412(%r13)
	imul	%r14,0x7f563412(%r14)
	imul	%r14,0x7f563412(%r15)
	nop
	imul	%r15,0x7f563412(%rax)
	imul	%r15,0x7f563412(%rcx)
	imul	%r15,0x7f563412(%rdx)
	imul	%r15,0x7f563412(%rbx)
	imul	%r15,0x7f563412(%rsp)
	imul	%r15,0x7f563412(%rbp)
	imul	%r15,0x7f563412(%rsi)
	imul	%r15,0x7f563412(%rdi)
	imul	%r15,0x7f563412(%r8)
	imul	%r15,0x7f563412(%r9)
	imul	%r15,0x7f563412(%r10)
	imul	%r15,0x7f563412(%r11)
	imul	%r15,0x7f563412(%r12)
	imul	%r15,0x7f563412(%r13)
	imul	%r15,0x7f563412(%r14)
	imul	%r15,0x7f563412(%r15)
	ret
	.cfi_endproc
        
        
	.p2align 4,,15
	.globl	ImulRegMem
	.type	ImulRegMem, @function
ImulRegMem:
	.cfi_startproc
	imul	(%rax),%rax
	imul	(%rax),%rcx
	imul	(%rax),%rdx
	imul	(%rax),%rbx
	imul	(%rax),%rsp
	imul	(%rax),%rbp
	imul	(%rax),%rsi
	imul	(%rax),%rdi
	imul	(%rax),%r8
	imul	(%rax),%r9
	imul	(%rax),%r10
	imul	(%rax),%r11
	imul	(%rax),%r12
	imul	(%rax),%r13
	imul	(%rax),%r14
	imul	(%rax),%r15
	nop
	imul	(%rcx),%rax
	imul	(%rcx),%rcx
	imul	(%rcx),%rdx
	imul	(%rcx),%rbx
	imul	(%rcx),%rsp
	imul	(%rcx),%rbp
	imul	(%rcx),%rsi
	imul	(%rcx),%rdi
	imul	(%rcx),%r8
	imul	(%rcx),%r9
	imul	(%rcx),%r10
	imul	(%rcx),%r11
	imul	(%rcx),%r12
	imul	(%rcx),%r13
	imul	(%rcx),%r14
	imul	(%rcx),%r15
	nop
	imul	(%rdx),%rax
	imul	(%rdx),%rcx
	imul	(%rdx),%rdx
	imul	(%rdx),%rbx
	imul	(%rdx),%rsp
	imul	(%rdx),%rbp
	imul	(%rdx),%rsi
	imul	(%rdx),%rdi
	imul	(%rdx),%r8
	imul	(%rdx),%r9
	imul	(%rdx),%r10
	imul	(%rdx),%r11
	imul	(%rdx),%r12
	imul	(%rdx),%r13
	imul	(%rdx),%r14
	imul	(%rdx),%r15
	nop
	imul	(%rbx),%rax
	imul	(%rbx),%rcx
	imul	(%rbx),%rdx
	imul	(%rbx),%rbx
	imul	(%rbx),%rsp
	imul	(%rbx),%rbp
	imul	(%rbx),%rsi
	imul	(%rbx),%rdi
	imul	(%rbx),%r8
	imul	(%rbx),%r9
	imul	(%rbx),%r10
	imul	(%rbx),%r11
	imul	(%rbx),%r12
	imul	(%rbx),%r13
	imul	(%rbx),%r14
	imul	(%rbx),%r15
	nop
	imul	(%rsp),%rax
	imul	(%rsp),%rcx
	imul	(%rsp),%rdx
	imul	(%rsp),%rbx
	imul	(%rsp),%rsp
	imul	(%rsp),%rbp
	imul	(%rsp),%rsi
	imul	(%rsp),%rdi
	imul	(%rsp),%r8
	imul	(%rsp),%r9
	imul	(%rsp),%r10
	imul	(%rsp),%r11
	imul	(%rsp),%r12
	imul	(%rsp),%r13
	imul	(%rsp),%r14
	imul	(%rsp),%r15
	nop
	imul	(%rbp),%rax
	imul	(%rbp),%rcx
	imul	(%rbp),%rdx
	imul	(%rbp),%rbx
	imul	(%rbp),%rsp
	imul	(%rbp),%rbp
	imul	(%rbp),%rsi
	imul	(%rbp),%rdi
	imul	(%rbp),%r8
	imul	(%rbp),%r9
	imul	(%rbp),%r10
	imul	(%rbp),%r11
	imul	(%rbp),%r12
	imul	(%rbp),%r13
	imul	(%rbp),%r14
	imul	(%rbp),%r15
	nop
	imul	(%rsi),%rax
	imul	(%rsi),%rcx
	imul	(%rsi),%rdx
	imul	(%rsi),%rbx
	imul	(%rsi),%rsp
	imul	(%rsi),%rbp
	imul	(%rsi),%rsi
	imul	(%rsi),%rdi
	imul	(%rsi),%r8
	imul	(%rsi),%r9
	imul	(%rsi),%r10
	imul	(%rsi),%r11
	imul	(%rsi),%r12
	imul	(%rsi),%r13
	imul	(%rsi),%r14
	imul	(%rsi),%r15
	nop
	imul	(%rdi),%rax
	imul	(%rdi),%rcx
	imul	(%rdi),%rdx
	imul	(%rdi),%rbx
	imul	(%rdi),%rsp
	imul	(%rdi),%rbp
	imul	(%rdi),%rsi
	imul	(%rdi),%rdi
	imul	(%rdi),%r8
	imul	(%rdi),%r9
	imul	(%rdi),%r10
	imul	(%rdi),%r11
	imul	(%rdi),%r12
	imul	(%rdi),%r13
	imul	(%rdi),%r14
	imul	(%rdi),%r15
	nop
	imul	(%r8 ),%rax
	imul	(%r8 ),%rcx
	imul	(%r8 ),%rdx
	imul	(%r8 ),%rbx
	imul	(%r8 ),%rsp
	imul	(%r8 ),%rbp
	imul	(%r8 ),%rsi
	imul	(%r8 ),%rdi
	imul	(%r8 ),%r8
	imul	(%r8 ),%r9
	imul	(%r8 ),%r10
	imul	(%r8 ),%r11
	imul	(%r8 ),%r12
	imul	(%r8 ),%r13
	imul	(%r8 ),%r14
	imul	(%r8 ),%r15
	nop
	imul	(%r9 ),%rax
	imul	(%r9 ),%rcx
	imul	(%r9 ),%rdx
	imul	(%r9 ),%rbx
	imul	(%r9 ),%rsp
	imul	(%r9 ),%rbp
	imul	(%r9 ),%rsi
	imul	(%r9 ),%rdi
	imul	(%r9 ),%r8
	imul	(%r9 ),%r9
	imul	(%r9 ),%r10
	imul	(%r9 ),%r11
	imul	(%r9 ),%r12
	imul	(%r9 ),%r13
	imul	(%r9 ),%r14
	imul	(%r9 ),%r15
	nop
	imul	(%r10),%rax
	imul	(%r10),%rcx
	imul	(%r10),%rdx
	imul	(%r10),%rbx
	imul	(%r10),%rsp
	imul	(%r10),%rbp
	imul	(%r10),%rsi
	imul	(%r10),%rdi
	imul	(%r10),%r8
	imul	(%r10),%r9
	imul	(%r10),%r10
	imul	(%r10),%r11
	imul	(%r10),%r12
	imul	(%r10),%r13
	imul	(%r10),%r14
	imul	(%r10),%r15
	nop
	imul	(%r11),%rax
	imul	(%r11),%rcx
	imul	(%r11),%rdx
	imul	(%r11),%rbx
	imul	(%r11),%rsp
	imul	(%r11),%rbp
	imul	(%r11),%rsi
	imul	(%r11),%rdi
	imul	(%r11),%r8
	imul	(%r11),%r9
	imul	(%r11),%r10
	imul	(%r11),%r11
	imul	(%r11),%r12
	imul	(%r11),%r13
	imul	(%r11),%r14
	imul	(%r11),%r15
	nop
	imul	(%r12),%rax
	imul	(%r12),%rcx
	imul	(%r12),%rdx
	imul	(%r12),%rbx
	imul	(%r12),%rsp
	imul	(%r12),%rbp
	imul	(%r12),%rsi
	imul	(%r12),%rdi
	imul	(%r12),%r8
	imul	(%r12),%r9
	imul	(%r12),%r10
	imul	(%r12),%r11
	imul	(%r12),%r12
	imul	(%r12),%r13
	imul	(%r12),%r14
	imul	(%r12),%r15
	nop
	imul	(%r13),%rax
	imul	(%r13),%rcx
	imul	(%r13),%rdx
	imul	(%r13),%rbx
	imul	(%r13),%rsp
	imul	(%r13),%rbp
	imul	(%r13),%rsi
	imul	(%r13),%rdi
	imul	(%r13),%r8
	imul	(%r13),%r9
	imul	(%r13),%r10
	imul	(%r13),%r11
	imul	(%r13),%r12
	imul	(%r13),%r13
	imul	(%r13),%r14
	imul	(%r13),%r15
	nop
	imul	(%r14),%rax
	imul	(%r14),%rcx
	imul	(%r14),%rdx
	imul	(%r14),%rbx
	imul	(%r14),%rsp
	imul	(%r14),%rbp
	imul	(%r14),%rsi
	imul	(%r14),%rdi
	imul	(%r14),%r8
	imul	(%r14),%r9
	imul	(%r14),%r10
	imul	(%r14),%r11
	imul	(%r14),%r12
	imul	(%r14),%r13
	imul	(%r14),%r14
	imul	(%r14),%r15
	nop
	imul	(%r15),%rax
	imul	(%r15),%rcx
	imul	(%r15),%rdx
	imul	(%r15),%rbx
	imul	(%r15),%rsp
	imul	(%r15),%rbp
	imul	(%r15),%rsi
	imul	(%r15),%rdi
	imul	(%r15),%r8
	imul	(%r15),%r9
	imul	(%r15),%r10
	imul	(%r15),%r11
	imul	(%r15),%r12
	imul	(%r15),%r13
	imul	(%r15),%r14
	imul	(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	ImulRegMem8
	.type	ImulRegMem8 ,@function
ImulRegMem8:
	.cfi_startproc
	imul	0x7f(%rax),%rax
	imul	0x7f(%rax),%rcx
	imul	0x7f(%rax),%rdx
	imul	0x7f(%rax),%rbx
	imul	0x7f(%rax),%rsp
	imul	0x7f(%rax),%rbp
	imul	0x7f(%rax),%rsi
	imul	0x7f(%rax),%rdi
	imul	0x7f(%rax),%r8
	imul	0x7f(%rax),%r9
	imul	0x7f(%rax),%r10
	imul	0x7f(%rax),%r11
	imul	0x7f(%rax),%r12
	imul	0x7f(%rax),%r13
	imul	0x7f(%rax),%r14
	imul	0x7f(%rax),%r15
	nop
	imul	0x7f(%rcx),%rax
	imul	0x7f(%rcx),%rcx
	imul	0x7f(%rcx),%rdx
	imul	0x7f(%rcx),%rbx
	imul	0x7f(%rcx),%rsp
	imul	0x7f(%rcx),%rbp
	imul	0x7f(%rcx),%rsi
	imul	0x7f(%rcx),%rdi
	imul	0x7f(%rcx),%r8
	imul	0x7f(%rcx),%r9
	imul	0x7f(%rcx),%r10
	imul	0x7f(%rcx),%r11
	imul	0x7f(%rcx),%r12
	imul	0x7f(%rcx),%r13
	imul	0x7f(%rcx),%r14
	imul	0x7f(%rcx),%r15
	nop
	imul	0x7f(%rdx),%rax
	imul	0x7f(%rdx),%rcx
	imul	0x7f(%rdx),%rdx
	imul	0x7f(%rdx),%rbx
	imul	0x7f(%rdx),%rsp
	imul	0x7f(%rdx),%rbp
	imul	0x7f(%rdx),%rsi
	imul	0x7f(%rdx),%rdi
	imul	0x7f(%rdx),%r8
	imul	0x7f(%rdx),%r9
	imul	0x7f(%rdx),%r10
	imul	0x7f(%rdx),%r11
	imul	0x7f(%rdx),%r12
	imul	0x7f(%rdx),%r13
	imul	0x7f(%rdx),%r14
	imul	0x7f(%rdx),%r15
	nop
	imul	0x7f(%rbx),%rax
	imul	0x7f(%rbx),%rcx
	imul	0x7f(%rbx),%rdx
	imul	0x7f(%rbx),%rbx
	imul	0x7f(%rbx),%rsp
	imul	0x7f(%rbx),%rbp
	imul	0x7f(%rbx),%rsi
	imul	0x7f(%rbx),%rdi
	imul	0x7f(%rbx),%r8
	imul	0x7f(%rbx),%r9
	imul	0x7f(%rbx),%r10
	imul	0x7f(%rbx),%r11
	imul	0x7f(%rbx),%r12
	imul	0x7f(%rbx),%r13
	imul	0x7f(%rbx),%r14
	imul	0x7f(%rbx),%r15
	nop
	imul	0x7f(%rsp),%rax
	imul	0x7f(%rsp),%rcx
	imul	0x7f(%rsp),%rdx
	imul	0x7f(%rsp),%rbx
	imul	0x7f(%rsp),%rsp
	imul	0x7f(%rsp),%rbp
	imul	0x7f(%rsp),%rsi
	imul	0x7f(%rsp),%rdi
	imul	0x7f(%rsp),%r8
	imul	0x7f(%rsp),%r9
	imul	0x7f(%rsp),%r10
	imul	0x7f(%rsp),%r11
	imul	0x7f(%rsp),%r12
	imul	0x7f(%rsp),%r13
	imul	0x7f(%rsp),%r14
	imul	0x7f(%rsp),%r15
	nop
	imul	0x7f(%rbp),%rax
	imul	0x7f(%rbp),%rcx
	imul	0x7f(%rbp),%rdx
	imul	0x7f(%rbp),%rbx
	imul	0x7f(%rbp),%rsp
	imul	0x7f(%rbp),%rbp
	imul	0x7f(%rbp),%rsi
	imul	0x7f(%rbp),%rdi
	imul	0x7f(%rbp),%r8
	imul	0x7f(%rbp),%r9
	imul	0x7f(%rbp),%r10
	imul	0x7f(%rbp),%r11
	imul	0x7f(%rbp),%r12
	imul	0x7f(%rbp),%r13
	imul	0x7f(%rbp),%r14
	imul	0x7f(%rbp),%r15
	nop
	imul	0x7f(%rsi),%rax
	imul	0x7f(%rsi),%rcx
	imul	0x7f(%rsi),%rdx
	imul	0x7f(%rsi),%rbx
	imul	0x7f(%rsi),%rsp
	imul	0x7f(%rsi),%rbp
	imul	0x7f(%rsi),%rsi
	imul	0x7f(%rsi),%rdi
	imul	0x7f(%rsi),%r8
	imul	0x7f(%rsi),%r9
	imul	0x7f(%rsi),%r10
	imul	0x7f(%rsi),%r11
	imul	0x7f(%rsi),%r12
	imul	0x7f(%rsi),%r13
	imul	0x7f(%rsi),%r14
	imul	0x7f(%rsi),%r15
	nop
	imul	0x7f(%rdi),%rax
	imul	0x7f(%rdi),%rcx
	imul	0x7f(%rdi),%rdx
	imul	0x7f(%rdi),%rbx
	imul	0x7f(%rdi),%rsp
	imul	0x7f(%rdi),%rbp
	imul	0x7f(%rdi),%rsi
	imul	0x7f(%rdi),%rdi
	imul	0x7f(%rdi),%r8
	imul	0x7f(%rdi),%r9
	imul	0x7f(%rdi),%r10
	imul	0x7f(%rdi),%r11
	imul	0x7f(%rdi),%r12
	imul	0x7f(%rdi),%r13
	imul	0x7f(%rdi),%r14
	imul	0x7f(%rdi),%r15
	nop
	imul	0x7f(%r8 ),%rax
	imul	0x7f(%r8 ),%rcx
	imul	0x7f(%r8 ),%rdx
	imul	0x7f(%r8 ),%rbx
	imul	0x7f(%r8 ),%rsp
	imul	0x7f(%r8 ),%rbp
	imul	0x7f(%r8 ),%rsi
	imul	0x7f(%r8 ),%rdi
	imul	0x7f(%r8 ),%r8
	imul	0x7f(%r8 ),%r9
	imul	0x7f(%r8 ),%r10
	imul	0x7f(%r8 ),%r11
	imul	0x7f(%r8 ),%r12
	imul	0x7f(%r8 ),%r13
	imul	0x7f(%r8 ),%r14
	imul	0x7f(%r8 ),%r15
	nop
	imul	0x7f(%r9 ),%rax
	imul	0x7f(%r9 ),%rcx
	imul	0x7f(%r9 ),%rdx
	imul	0x7f(%r9 ),%rbx
	imul	0x7f(%r9 ),%rsp
	imul	0x7f(%r9 ),%rbp
	imul	0x7f(%r9 ),%rsi
	imul	0x7f(%r9 ),%rdi
	imul	0x7f(%r9 ),%r8
	imul	0x7f(%r9 ),%r9
	imul	0x7f(%r9 ),%r10
	imul	0x7f(%r9 ),%r11
	imul	0x7f(%r9 ),%r12
	imul	0x7f(%r9 ),%r13
	imul	0x7f(%r9 ),%r14
	imul	0x7f(%r9 ),%r15
	nop
	imul	0x7f(%r10),%rax
	imul	0x7f(%r10),%rcx
	imul	0x7f(%r10),%rdx
	imul	0x7f(%r10),%rbx
	imul	0x7f(%r10),%rsp
	imul	0x7f(%r10),%rbp
	imul	0x7f(%r10),%rsi
	imul	0x7f(%r10),%rdi
	imul	0x7f(%r10),%r8
	imul	0x7f(%r10),%r9
	imul	0x7f(%r10),%r10
	imul	0x7f(%r10),%r11
	imul	0x7f(%r10),%r12
	imul	0x7f(%r10),%r13
	imul	0x7f(%r10),%r14
	imul	0x7f(%r10),%r15
	nop
	imul	0x7f(%r11),%rax
	imul	0x7f(%r11),%rcx
	imul	0x7f(%r11),%rdx
	imul	0x7f(%r11),%rbx
	imul	0x7f(%r11),%rsp
	imul	0x7f(%r11),%rbp
	imul	0x7f(%r11),%rsi
	imul	0x7f(%r11),%rdi
	imul	0x7f(%r11),%r8
	imul	0x7f(%r11),%r9
	imul	0x7f(%r11),%r10
	imul	0x7f(%r11),%r11
	imul	0x7f(%r11),%r12
	imul	0x7f(%r11),%r13
	imul	0x7f(%r11),%r14
	imul	0x7f(%r11),%r15
	nop
	imul	0x7f(%r12),%rax
	imul	0x7f(%r12),%rcx
	imul	0x7f(%r12),%rdx
	imul	0x7f(%r12),%rbx
	imul	0x7f(%r12),%rsp
	imul	0x7f(%r12),%rbp
	imul	0x7f(%r12),%rsi
	imul	0x7f(%r12),%rdi
	imul	0x7f(%r12),%r8
	imul	0x7f(%r12),%r9
	imul	0x7f(%r12),%r10
	imul	0x7f(%r12),%r11
	imul	0x7f(%r12),%r12
	imul	0x7f(%r12),%r13
	imul	0x7f(%r12),%r14
	imul	0x7f(%r12),%r15
	nop
	imul	0x7f(%r13),%rax
	imul	0x7f(%r13),%rcx
	imul	0x7f(%r13),%rdx
	imul	0x7f(%r13),%rbx
	imul	0x7f(%r13),%rsp
	imul	0x7f(%r13),%rbp
	imul	0x7f(%r13),%rsi
	imul	0x7f(%r13),%rdi
	imul	0x7f(%r13),%r8
	imul	0x7f(%r13),%r9
	imul	0x7f(%r13),%r10
	imul	0x7f(%r13),%r11
	imul	0x7f(%r13),%r12
	imul	0x7f(%r13),%r13
	imul	0x7f(%r13),%r14
	imul	0x7f(%r13),%r15
	nop
	imul	0x7f(%r14),%rax
	imul	0x7f(%r14),%rcx
	imul	0x7f(%r14),%rdx
	imul	0x7f(%r14),%rbx
	imul	0x7f(%r14),%rsp
	imul	0x7f(%r14),%rbp
	imul	0x7f(%r14),%rsi
	imul	0x7f(%r14),%rdi
	imul	0x7f(%r14),%r8
	imul	0x7f(%r14),%r9
	imul	0x7f(%r14),%r10
	imul	0x7f(%r14),%r11
	imul	0x7f(%r14),%r12
	imul	0x7f(%r14),%r13
	imul	0x7f(%r14),%r14
	imul	0x7f(%r14),%r15
	nop
	imul	0x7f(%r15),%rax
	imul	0x7f(%r15),%rcx
	imul	0x7f(%r15),%rdx
	imul	0x7f(%r15),%rbx
	imul	0x7f(%r15),%rsp
	imul	0x7f(%r15),%rbp
	imul	0x7f(%r15),%rsi
	imul	0x7f(%r15),%rdi
	imul	0x7f(%r15),%r8
	imul	0x7f(%r15),%r9
	imul	0x7f(%r15),%r10
	imul	0x7f(%r15),%r11
	imul	0x7f(%r15),%r12
	imul	0x7f(%r15),%r13
	imul	0x7f(%r15),%r14
	imul	0x7f(%r15),%r15
	ret
	.cfi_endproc


	.p2align 4,,15
	.globl	ImulRegMem32
	.type	ImulRegMem32, @function
ImulRegMem32:
	.cfi_startproc
	imul	0x7f563412(%rax),%rax
	imul	0x7f563412(%rax),%rcx
	imul	0x7f563412(%rax),%rdx
	imul	0x7f563412(%rax),%rbx
	imul	0x7f563412(%rax),%rsp
	imul	0x7f563412(%rax),%rbp
	imul	0x7f563412(%rax),%rsi
	imul	0x7f563412(%rax),%rdi
	imul	0x7f563412(%rax),%r8
	imul	0x7f563412(%rax),%r9
	imul	0x7f563412(%rax),%r10
	imul	0x7f563412(%rax),%r11
	imul	0x7f563412(%rax),%r12
	imul	0x7f563412(%rax),%r13
	imul	0x7f563412(%rax),%r14
	imul	0x7f563412(%rax),%r15
	nop
	imul	0x7f563412(%rcx),%rax
	imul	0x7f563412(%rcx),%rcx
	imul	0x7f563412(%rcx),%rdx
	imul	0x7f563412(%rcx),%rbx
	imul	0x7f563412(%rcx),%rsp
	imul	0x7f563412(%rcx),%rbp
	imul	0x7f563412(%rcx),%rsi
	imul	0x7f563412(%rcx),%rdi
	imul	0x7f563412(%rcx),%r8
	imul	0x7f563412(%rcx),%r9
	imul	0x7f563412(%rcx),%r10
	imul	0x7f563412(%rcx),%r11
	imul	0x7f563412(%rcx),%r12
	imul	0x7f563412(%rcx),%r13
	imul	0x7f563412(%rcx),%r14
	imul	0x7f563412(%rcx),%r15
	nop
	imul	0x7f563412(%rdx),%rax
	imul	0x7f563412(%rdx),%rcx
	imul	0x7f563412(%rdx),%rdx
	imul	0x7f563412(%rdx),%rbx
	imul	0x7f563412(%rdx),%rsp
	imul	0x7f563412(%rdx),%rbp
	imul	0x7f563412(%rdx),%rsi
	imul	0x7f563412(%rdx),%rdi
	imul	0x7f563412(%rdx),%r8
	imul	0x7f563412(%rdx),%r9
	imul	0x7f563412(%rdx),%r10
	imul	0x7f563412(%rdx),%r11
	imul	0x7f563412(%rdx),%r12
	imul	0x7f563412(%rdx),%r13
	imul	0x7f563412(%rdx),%r14
	imul	0x7f563412(%rdx),%r15
	nop
	imul	0x7f563412(%rbx),%rax
	imul	0x7f563412(%rbx),%rcx
	imul	0x7f563412(%rbx),%rdx
	imul	0x7f563412(%rbx),%rbx
	imul	0x7f563412(%rbx),%rsp
	imul	0x7f563412(%rbx),%rbp
	imul	0x7f563412(%rbx),%rsi
	imul	0x7f563412(%rbx),%rdi
	imul	0x7f563412(%rbx),%r8
	imul	0x7f563412(%rbx),%r9
	imul	0x7f563412(%rbx),%r10
	imul	0x7f563412(%rbx),%r11
	imul	0x7f563412(%rbx),%r12
	imul	0x7f563412(%rbx),%r13
	imul	0x7f563412(%rbx),%r14
	imul	0x7f563412(%rbx),%r15
	nop
	imul	0x7f563412(%rsp),%rax
	imul	0x7f563412(%rsp),%rcx
	imul	0x7f563412(%rsp),%rdx
	imul	0x7f563412(%rsp),%rbx
	imul	0x7f563412(%rsp),%rsp
	imul	0x7f563412(%rsp),%rbp
	imul	0x7f563412(%rsp),%rsi
	imul	0x7f563412(%rsp),%rdi
	imul	0x7f563412(%rsp),%r8
	imul	0x7f563412(%rsp),%r9
	imul	0x7f563412(%rsp),%r10
	imul	0x7f563412(%rsp),%r11
	imul	0x7f563412(%rsp),%r12
	imul	0x7f563412(%rsp),%r13
	imul	0x7f563412(%rsp),%r14
	imul	0x7f563412(%rsp),%r15
	nop
	imul	0x7f563412(%rbp),%rax
	imul	0x7f563412(%rbp),%rcx
	imul	0x7f563412(%rbp),%rdx
	imul	0x7f563412(%rbp),%rbx
	imul	0x7f563412(%rbp),%rsp
	imul	0x7f563412(%rbp),%rbp
	imul	0x7f563412(%rbp),%rsi
	imul	0x7f563412(%rbp),%rdi
	imul	0x7f563412(%rbp),%r8
	imul	0x7f563412(%rbp),%r9
	imul	0x7f563412(%rbp),%r10
	imul	0x7f563412(%rbp),%r11
	imul	0x7f563412(%rbp),%r12
	imul	0x7f563412(%rbp),%r13
	imul	0x7f563412(%rbp),%r14
	imul	0x7f563412(%rbp),%r15
	nop
	imul	0x7f563412(%rsi),%rax
	imul	0x7f563412(%rsi),%rcx
	imul	0x7f563412(%rsi),%rdx
	imul	0x7f563412(%rsi),%rbx
	imul	0x7f563412(%rsi),%rsp
	imul	0x7f563412(%rsi),%rbp
	imul	0x7f563412(%rsi),%rsi
	imul	0x7f563412(%rsi),%rdi
	imul	0x7f563412(%rsi),%r8
	imul	0x7f563412(%rsi),%r9
	imul	0x7f563412(%rsi),%r10
	imul	0x7f563412(%rsi),%r11
	imul	0x7f563412(%rsi),%r12
	imul	0x7f563412(%rsi),%r13
	imul	0x7f563412(%rsi),%r14
	imul	0x7f563412(%rsi),%r15
	nop
	imul	0x7f563412(%rdi),%rax
	imul	0x7f563412(%rdi),%rcx
	imul	0x7f563412(%rdi),%rdx
	imul	0x7f563412(%rdi),%rbx
	imul	0x7f563412(%rdi),%rsp
	imul	0x7f563412(%rdi),%rbp
	imul	0x7f563412(%rdi),%rsi
	imul	0x7f563412(%rdi),%rdi
	imul	0x7f563412(%rdi),%r8
	imul	0x7f563412(%rdi),%r9
	imul	0x7f563412(%rdi),%r10
	imul	0x7f563412(%rdi),%r11
	imul	0x7f563412(%rdi),%r12
	imul	0x7f563412(%rdi),%r13
	imul	0x7f563412(%rdi),%r14
	imul	0x7f563412(%rdi),%r15
	nop
	imul	0x7f563412(%r8 ),%rax
	imul	0x7f563412(%r8 ),%rcx
	imul	0x7f563412(%r8 ),%rdx
	imul	0x7f563412(%r8 ),%rbx
	imul	0x7f563412(%r8 ),%rsp
	imul	0x7f563412(%r8 ),%rbp
	imul	0x7f563412(%r8 ),%rsi
	imul	0x7f563412(%r8 ),%rdi
	imul	0x7f563412(%r8 ),%r8
	imul	0x7f563412(%r8 ),%r9
	imul	0x7f563412(%r8 ),%r10
	imul	0x7f563412(%r8 ),%r11
	imul	0x7f563412(%r8 ),%r12
	imul	0x7f563412(%r8 ),%r13
	imul	0x7f563412(%r8 ),%r14
	imul	0x7f563412(%r8 ),%r15
	nop
	imul	0x7f563412(%r9 ),%rax
	imul	0x7f563412(%r9 ),%rcx
	imul	0x7f563412(%r9 ),%rdx
	imul	0x7f563412(%r9 ),%rbx
	imul	0x7f563412(%r9 ),%rsp
	imul	0x7f563412(%r9 ),%rbp
	imul	0x7f563412(%r9 ),%rsi
	imul	0x7f563412(%r9 ),%rdi
	imul	0x7f563412(%r9 ),%r8
	imul	0x7f563412(%r9 ),%r9
	imul	0x7f563412(%r9 ),%r10
	imul	0x7f563412(%r9 ),%r11
	imul	0x7f563412(%r9 ),%r12
	imul	0x7f563412(%r9 ),%r13
	imul	0x7f563412(%r9 ),%r14
	imul	0x7f563412(%r9 ),%r15
	nop
	imul	0x7f563412(%r10),%rax
	imul	0x7f563412(%r10),%rcx
	imul	0x7f563412(%r10),%rdx
	imul	0x7f563412(%r10),%rbx
	imul	0x7f563412(%r10),%rsp
	imul	0x7f563412(%r10),%rbp
	imul	0x7f563412(%r10),%rsi
	imul	0x7f563412(%r10),%rdi
	imul	0x7f563412(%r10),%r8
	imul	0x7f563412(%r10),%r9
	imul	0x7f563412(%r10),%r10
	imul	0x7f563412(%r10),%r11
	imul	0x7f563412(%r10),%r12
	imul	0x7f563412(%r10),%r13
	imul	0x7f563412(%r10),%r14
	imul	0x7f563412(%r10),%r15
	nop
	imul	0x7f563412(%r11),%rax
	imul	0x7f563412(%r11),%rcx
	imul	0x7f563412(%r11),%rdx
	imul	0x7f563412(%r11),%rbx
	imul	0x7f563412(%r11),%rsp
	imul	0x7f563412(%r11),%rbp
	imul	0x7f563412(%r11),%rsi
	imul	0x7f563412(%r11),%rdi
	imul	0x7f563412(%r11),%r8
	imul	0x7f563412(%r11),%r9
	imul	0x7f563412(%r11),%r10
	imul	0x7f563412(%r11),%r11
	imul	0x7f563412(%r11),%r12
	imul	0x7f563412(%r11),%r13
	imul	0x7f563412(%r11),%r14
	imul	0x7f563412(%r11),%r15
	nop
	imul	0x7f563412(%r12),%rax
	imul	0x7f563412(%r12),%rcx
	imul	0x7f563412(%r12),%rdx
	imul	0x7f563412(%r12),%rbx
	imul	0x7f563412(%r12),%rsp
	imul	0x7f563412(%r12),%rbp
	imul	0x7f563412(%r12),%rsi
	imul	0x7f563412(%r12),%rdi
	imul	0x7f563412(%r12),%r8
	imul	0x7f563412(%r12),%r9
	imul	0x7f563412(%r12),%r10
	imul	0x7f563412(%r12),%r11
	imul	0x7f563412(%r12),%r12
	imul	0x7f563412(%r12),%r13
	imul	0x7f563412(%r12),%r14
	imul	0x7f563412(%r12),%r15
	nop
	imul	0x7f563412(%r13),%rax
	imul	0x7f563412(%r13),%rcx
	imul	0x7f563412(%r13),%rdx
	imul	0x7f563412(%r13),%rbx
	imul	0x7f563412(%r13),%rsp
	imul	0x7f563412(%r13),%rbp
	imul	0x7f563412(%r13),%rsi
	imul	0x7f563412(%r13),%rdi
	imul	0x7f563412(%r13),%r8
	imul	0x7f563412(%r13),%r9
	imul	0x7f563412(%r13),%r10
	imul	0x7f563412(%r13),%r11
	imul	0x7f563412(%r13),%r12
	imul	0x7f563412(%r13),%r13
	imul	0x7f563412(%r13),%r14
	imul	0x7f563412(%r13),%r15
	nop
	imul	0x7f563412(%r14),%rax
	imul	0x7f563412(%r14),%rcx
	imul	0x7f563412(%r14),%rdx
	imul	0x7f563412(%r14),%rbx
	imul	0x7f563412(%r14),%rsp
	imul	0x7f563412(%r14),%rbp
	imul	0x7f563412(%r14),%rsi
	imul	0x7f563412(%r14),%rdi
	imul	0x7f563412(%r14),%r8
	imul	0x7f563412(%r14),%r9
	imul	0x7f563412(%r14),%r10
	imul	0x7f563412(%r14),%r11
	imul	0x7f563412(%r14),%r12
	imul	0x7f563412(%r14),%r13
	imul	0x7f563412(%r14),%r14
	imul	0x7f563412(%r14),%r15
	nop
	imul	0x7f563412(%r15),%rax
	imul	0x7f563412(%r15),%rcx
	imul	0x7f563412(%r15),%rdx
	imul	0x7f563412(%r15),%rbx
	imul	0x7f563412(%r15),%rsp
	imul	0x7f563412(%r15),%rbp
	imul	0x7f563412(%r15),%rsi
	imul	0x7f563412(%r15),%rdi
	imul	0x7f563412(%r15),%r8
	imul	0x7f563412(%r15),%r9
	imul	0x7f563412(%r15),%r10
	imul	0x7f563412(%r15),%r11
	imul	0x7f563412(%r15),%r12
	imul	0x7f563412(%r15),%r13
	imul	0x7f563412(%r15),%r14
	imul	0x7f563412(%r15),%r15
	ret
	.cfi_endproc


        nop
        nop
        // reg64 += reg64
	imul	%rax,%rax
	imul	%rax,%rcx
	imul	%rax,%rdx
	imul	%rax,%rbx
	imul	%rax,%rsp
	imul	%rax,%rbp
	imul	%rax,%rsi
	imul	%rax,%rdi
	imul	%rax,%r8
	imul	%rax,%r9
	imul	%rax,%r10
	imul	%rax,%r11
	imul	%rax,%r12
	imul	%rax,%r13
	imul	%rax,%r14
	imul	%rax,%r15
	nop
	imul	%rcx,%rax
	imul	%rcx,%rcx
	imul	%rcx,%rdx
	imul	%rcx,%rbx
	imul	%rcx,%rsp
	imul	%rcx,%rbp
	imul	%rcx,%rsi
	imul	%rcx,%rdi
	imul	%rcx,%r8
	imul	%rcx,%r9
	imul	%rcx,%r10
	imul	%rcx,%r11
	imul	%rcx,%r12
	imul	%rcx,%r13
	imul	%rcx,%r14
	imul	%rcx,%r15
	nop
	imul	%rdx,%rax
	imul	%rdx,%rcx
	imul	%rdx,%rdx
	imul	%rdx,%rbx
	imul	%rdx,%rsp
	imul	%rdx,%rbp
	imul	%rdx,%rsi
	imul	%rdx,%rdi
	imul	%rdx,%r8
	imul	%rdx,%r9
	imul	%rdx,%r10
	imul	%rdx,%r11
	imul	%rdx,%r12
	imul	%rdx,%r13
	imul	%rdx,%r14
	imul	%rdx,%r15
	nop
	imul	%rbx,%rax
	imul	%rbx,%rcx
	imul	%rbx,%rdx
	imul	%rbx,%rbx
	imul	%rbx,%rsp
	imul	%rbx,%rbp
	imul	%rbx,%rsi
	imul	%rbx,%rdi
	imul	%rbx,%r8
	imul	%rbx,%r9
	imul	%rbx,%r10
	imul	%rbx,%r11
	imul	%rbx,%r12
	imul	%rbx,%r13
	imul	%rbx,%r14
	imul	%rbx,%r15
	nop
	imul	%rsp,%rax
	imul	%rsp,%rcx
	imul	%rsp,%rdx
	imul	%rsp,%rbx
	imul	%rsp,%rsp
	imul	%rsp,%rbp
	imul	%rsp,%rsi
	imul	%rsp,%rdi
	imul	%rsp,%r8
	imul	%rsp,%r9
	imul	%rsp,%r10
	imul	%rsp,%r11
	imul	%rsp,%r12
	imul	%rsp,%r13
	imul	%rsp,%r14
	imul	%rsp,%r15
	nop
	imul	%rbp,%rax
	imul	%rbp,%rcx
	imul	%rbp,%rdx
	imul	%rbp,%rbx
	imul	%rbp,%rsp
	imul	%rbp,%rbp
	imul	%rbp,%rsi
	imul	%rbp,%rdi
	imul	%rbp,%r8
	imul	%rbp,%r9
	imul	%rbp,%r10
	imul	%rbp,%r11
	imul	%rbp,%r12
	imul	%rbp,%r13
	imul	%rbp,%r14
	imul	%rbp,%r15
	nop
	imul	%rsi,%rax
	imul	%rsi,%rcx
	imul	%rsi,%rdx
	imul	%rsi,%rbx
	imul	%rsi,%rsp
	imul	%rsi,%rbp
	imul	%rsi,%rsi
	imul	%rsi,%rdi
	imul	%rsi,%r8
	imul	%rsi,%r9
	imul	%rsi,%r10
	imul	%rsi,%r11
	imul	%rsi,%r12
	imul	%rsi,%r13
	imul	%rsi,%r14
	imul	%rsi,%r15
	nop
	imul	%rdi,%rax
	imul	%rdi,%rcx
	imul	%rdi,%rdx
	imul	%rdi,%rbx
	imul	%rdi,%rsp
	imul	%rdi,%rbp
	imul	%rdi,%rsi
	imul	%rdi,%rdi
	imul	%rdi,%r8
	imul	%rdi,%r9
	imul	%rdi,%r10
	imul	%rdi,%r11
	imul	%rdi,%r12
	imul	%rdi,%r13
	imul	%rdi,%r14
	imul	%rdi,%r15
	nop
	imul	%r8 ,%rax
	imul	%r8 ,%rcx
	imul	%r8 ,%rdx
	imul	%r8 ,%rbx
	imul	%r8 ,%rsp
	imul	%r8 ,%rbp
	imul	%r8 ,%rsi
	imul	%r8 ,%rdi
	imul	%r8 ,%r8
	imul	%r8 ,%r9
	imul	%r8 ,%r10
	imul	%r8 ,%r11
	imul	%r8 ,%r12
	imul	%r8 ,%r13
	imul	%r8 ,%r14
	imul	%r8 ,%r15
	nop
	imul	%r9 ,%rax
	imul	%r9 ,%rcx
	imul	%r9 ,%rdx
	imul	%r9 ,%rbx
	imul	%r9 ,%rsp
	imul	%r9 ,%rbp
	imul	%r9 ,%rsi
	imul	%r9 ,%rdi
	imul	%r9 ,%r8
	imul	%r9 ,%r9
	imul	%r9 ,%r10
	imul	%r9 ,%r11
	imul	%r9 ,%r12
	imul	%r9 ,%r13
	imul	%r9 ,%r14
	imul	%r9 ,%r15
	nop
	imul	%r10,%rax
	imul	%r10,%rcx
	imul	%r10,%rdx
	imul	%r10,%rbx
	imul	%r10,%rsp
	imul	%r10,%rbp
	imul	%r10,%rsi
	imul	%r10,%rdi
	imul	%r10,%r8
	imul	%r10,%r9
	imul	%r10,%r10
	imul	%r10,%r11
	imul	%r10,%r12
	imul	%r10,%r13
	imul	%r10,%r14
	imul	%r10,%r15
	nop
	imul	%r11,%rax
	imul	%r11,%rcx
	imul	%r11,%rdx
	imul	%r11,%rbx
	imul	%r11,%rsp
	imul	%r11,%rbp
	imul	%r11,%rsi
	imul	%r11,%rdi
	imul	%r11,%r8
	imul	%r11,%r9
	imul	%r11,%r10
	imul	%r11,%r11
	imul	%r11,%r12
	imul	%r11,%r13
	imul	%r11,%r14
	imul	%r11,%r15
	nop
	imul	%r12,%rax
	imul	%r12,%rcx
	imul	%r12,%rdx
	imul	%r12,%rbx
	imul	%r12,%rsp
	imul	%r12,%rbp
	imul	%r12,%rsi
	imul	%r12,%rdi
	imul	%r12,%r8
	imul	%r12,%r9
	imul	%r12,%r10
	imul	%r12,%r11
	imul	%r12,%r12
	imul	%r12,%r13
	imul	%r12,%r14
	imul	%r12,%r15
	nop
	imul	%r13,%rax
	imul	%r13,%rcx
	imul	%r13,%rdx
	imul	%r13,%rbx
	imul	%r13,%rsp
	imul	%r13,%rbp
	imul	%r13,%rsi
	imul	%r13,%rdi
	imul	%r13,%r8
	imul	%r13,%r9
	imul	%r13,%r10
	imul	%r13,%r11
	imul	%r13,%r12
	imul	%r13,%r13
	imul	%r13,%r14
	imul	%r13,%r15
	nop
	imul	%r14,%rax
	imul	%r14,%rcx
	imul	%r14,%rdx
	imul	%r14,%rbx
	imul	%r14,%rsp
	imul	%r14,%rbp
	imul	%r14,%rsi
	imul	%r14,%rdi
	imul	%r14,%r8
	imul	%r14,%r9
	imul	%r14,%r10
	imul	%r14,%r11
	imul	%r14,%r12
	imul	%r14,%r13
	imul	%r14,%r14
	imul	%r14,%r15
	nop
	imul	%r15,%rax
	imul	%r15,%rcx
	imul	%r15,%rdx
	imul	%r15,%rbx
	imul	%r15,%rsp
	imul	%r15,%rbp
	imul	%r15,%rsi
	imul	%r15,%rdi
	imul	%r15,%r8
	imul	%r15,%r9
	imul	%r15,%r10
	imul	%r15,%r11
	imul	%r15,%r12
	imul	%r15,%r13
	imul	%r15,%r14
	imul	%r15,%r15
	ret
	.cfi_endproc


*/
