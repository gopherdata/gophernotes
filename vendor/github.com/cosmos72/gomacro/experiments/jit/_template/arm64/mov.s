	.p2align 4,,15
	.globl	mov
	.type	mov, @function
mov:
	.cfi_startproc
	mov	x0, x0
	mov	x0, x1
	mov	x0, x2
	mov	x0, x3
	mov	x0, x15
	mov	x0, x30
	nop
	mov	x1, x0
	mov	x1, x1
	mov	x1, x2
	mov	x1, x3
	mov	x1, x15
	mov	x1, x30
	nop
	mov	x2, x0
	mov	x2, x1
	mov	x2, x2
	mov	x2, x3
	mov	x2, x15
	mov	x2, x30
	nop
	mov	x3, x0
	mov	x3, x1
	mov	x3, x2
	mov	x3, x3
	mov	x3, x15
	mov	x3, x30
	nop
	mov	x15, x0
	mov	x15, x1
	mov	x15, x2
	mov	x15, x3
	mov	x15, x15
	mov	x15, x30
	nop
	mov	x30, x0
	mov	x30, x1
	mov	x30, x2
	mov	x30, x3
	mov	x30, x15
	mov	x30, x30
	ret
	.cfi_endproc
	
