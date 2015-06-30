.global _sum

.text
_sum:
	mov	$0,%eax
	mov %eax,-24(%rsp)
	ret

.data
