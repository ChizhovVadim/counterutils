// +build avx
//+build !noasm !appengine
// AUTO-GENERATED BY C2GOASM -- DO NOT EDIT

TEXT ·_update_hidden(SB), $56-56

	MOVQ previous_outputs+0(FP), DI
	MOVQ update_indices+8(FP), SI
	MOVQ update_coeffs+16(FP), DX
	MOVQ update_size+24(FP), CX
	MOVQ weights+32(FP), R8
	MOVQ outputs+40(FP), R9
	MOVQ outputs_len+48(FP), R10
	ADDQ $8, SP
	MOVQ R10, 40(SP)

	LONG $0x2404894c             // mov    qword [rsp], r8
	LONG $0x24548948; BYTE $0x20 // mov    qword [rsp + 32], rdx
	LONG $0x2824448b             // mov    eax, dword 40[rsp] /* [rbp + 16] */
	WORD $0xc085                 // test    eax, eax
	JLE  LBB0_33
	WORD $0x8941; BYTE $0xcb     // mov    r11d, ecx
	WORD $0x8941; BYTE $0xc4     // mov    r12d, eax
	WORD $0xf883; BYTE $0x20     // cmp    eax, 32
	JB   LBB0_2
	LONG $0xa7048d4a             // lea    rax, [rdi + 4*r12]
	WORD $0x394c; BYTE $0xc8     // cmp    rax, r9
	JBE  LBB0_6
	LONG $0xa1048d4b             // lea    rax, [r9 + 4*r12]
	WORD $0x3948; BYTE $0xf8     // cmp    rax, rdi
	JBE  LBB0_6

LBB0_2:
	WORD $0xc031 // xor    eax, eax

LBB0_12:
	WORD $0x8948; BYTE $0xc3 // mov    rbx, rax
	WORD $0xf748; BYTE $0xd3 // not    rbx
	WORD $0x014c; BYTE $0xe3 // add    rbx, r12
	WORD $0x894c; BYTE $0xe1 // mov    rcx, r12
	LONG $0x03e18348         // and    rcx, 3
	JE   LBB0_14

LBB0_13:
	LONG $0x0410fac5; BYTE $0x87   // vmovss    xmm0, dword [rdi + 4*rax]
	LONG $0x117ac1c4; WORD $0x8104 // vmovss    dword [r9 + 4*rax], xmm0
	LONG $0x01c08348               // add    rax, 1
	LONG $0xffc18348               // add    rcx, -1
	JNE  LBB0_13

LBB0_14:
	LONG $0x03fb8348 // cmp    rbx, 3
	LONG $0x28244c8b // mov    ecx, dword 40[rsp] /* [rbp + 16] */
	JB   LBB0_16

LBB0_15:
	LONG $0x0410fac5; BYTE $0x87               // vmovss    xmm0, dword [rdi + 4*rax]
	LONG $0x117ac1c4; WORD $0x8104             // vmovss    dword [r9 + 4*rax], xmm0
	LONG $0x4410fac5; WORD $0x0487             // vmovss    xmm0, dword [rdi + 4*rax + 4]
	LONG $0x117ac1c4; WORD $0x8144; BYTE $0x04 // vmovss    dword [r9 + 4*rax + 4], xmm0
	LONG $0x4410fac5; WORD $0x0887             // vmovss    xmm0, dword [rdi + 4*rax + 8]
	LONG $0x117ac1c4; WORD $0x8144; BYTE $0x08 // vmovss    dword [r9 + 4*rax + 8], xmm0
	LONG $0x4410fac5; WORD $0x0c87             // vmovss    xmm0, dword [rdi + 4*rax + 12]
	LONG $0x117ac1c4; WORD $0x8144; BYTE $0x0c // vmovss    dword [r9 + 4*rax + 12], xmm0
	LONG $0x04c08348                           // add    rax, 4
	WORD $0x3949; BYTE $0xc4                   // cmp    r12, rax
	JNE  LBB0_15
	JMP  LBB0_16

LBB0_6:
	WORD $0x8944; BYTE $0xe0 // mov    eax, r12d
	WORD $0xe083; BYTE $0xe0 // and    eax, -32
	LONG $0xe0488d48         // lea    rcx, [rax - 32]
	WORD $0x8949; BYTE $0xca // mov    r10, rcx
	LONG $0x05eac149         // shr    r10, 5
	LONG $0x01c28349         // add    r10, 1
	WORD $0x8548; BYTE $0xc9 // test    rcx, rcx
	JE   LBB0_34
	WORD $0x894d; BYTE $0xd6 // mov    r14, r10
	LONG $0xfee68349         // and    r14, -2
	WORD $0xf749; BYTE $0xde // neg    r14
	WORD $0xdb31             // xor    ebx, ebx
	LONG $0x28244c8b         // mov    ecx, dword 40[rsp] /* [rbp + 16] */

LBB0_8:
	LONG $0x0410fcc5; BYTE $0x9f               // vmovups    ymm0, yword [rdi + 4*rbx]
	LONG $0x4c10fcc5; WORD $0x209f             // vmovups    ymm1, yword [rdi + 4*rbx + 32]
	LONG $0x5410fcc5; WORD $0x409f             // vmovups    ymm2, yword [rdi + 4*rbx + 64]
	LONG $0x5c10fcc5; WORD $0x609f             // vmovups    ymm3, yword [rdi + 4*rbx + 96]
	LONG $0x117cc1c4; WORD $0x9904             // vmovups    yword [r9 + 4*rbx], ymm0
	LONG $0x117cc1c4; WORD $0x994c; BYTE $0x20 // vmovups    yword [r9 + 4*rbx + 32], ymm1
	LONG $0x117cc1c4; WORD $0x9954; BYTE $0x40 // vmovups    yword [r9 + 4*rbx + 64], ymm2
	LONG $0x117cc1c4; WORD $0x995c; BYTE $0x60 // vmovups    yword [r9 + 4*rbx + 96], ymm3
	QUAD $0x0000809f8410fcc5; BYTE $0x00       // vmovups    ymm0, yword [rdi + 4*rbx + 128]
	QUAD $0x0000a09f8c10fcc5; BYTE $0x00       // vmovups    ymm1, yword [rdi + 4*rbx + 160]
	QUAD $0x0000c09f9410fcc5; BYTE $0x00       // vmovups    ymm2, yword [rdi + 4*rbx + 192]
	QUAD $0x0000e09f9c10fcc5; BYTE $0x00       // vmovups    ymm3, yword [rdi + 4*rbx + 224]
	QUAD $0x00809984117cc1c4; WORD $0x0000     // vmovups    yword [r9 + 4*rbx + 128], ymm0
	QUAD $0x00a0998c117cc1c4; WORD $0x0000     // vmovups    yword [r9 + 4*rbx + 160], ymm1
	QUAD $0x00c09994117cc1c4; WORD $0x0000     // vmovups    yword [r9 + 4*rbx + 192], ymm2
	QUAD $0x00e0999c117cc1c4; WORD $0x0000     // vmovups    yword [r9 + 4*rbx + 224], ymm3
	LONG $0x40c38348                           // add    rbx, 64
	LONG $0x02c68349                           // add    r14, 2
	JNE  LBB0_8
	LONG $0x01c2f641                           // test    r10b, 1
	JE   LBB0_11

LBB0_10:
	LONG $0x0410fcc5; BYTE $0x9f               // vmovups    ymm0, yword [rdi + 4*rbx]
	LONG $0x4c10fcc5; WORD $0x209f             // vmovups    ymm1, yword [rdi + 4*rbx + 32]
	LONG $0x5410fcc5; WORD $0x409f             // vmovups    ymm2, yword [rdi + 4*rbx + 64]
	LONG $0x5c10fcc5; WORD $0x609f             // vmovups    ymm3, yword [rdi + 4*rbx + 96]
	LONG $0x117cc1c4; WORD $0x9904             // vmovups    yword [r9 + 4*rbx], ymm0
	LONG $0x117cc1c4; WORD $0x994c; BYTE $0x20 // vmovups    yword [r9 + 4*rbx + 32], ymm1
	LONG $0x117cc1c4; WORD $0x9954; BYTE $0x40 // vmovups    yword [r9 + 4*rbx + 64], ymm2
	LONG $0x117cc1c4; WORD $0x995c; BYTE $0x60 // vmovups    yword [r9 + 4*rbx + 96], ymm3

LBB0_11:
	WORD $0x394c; BYTE $0xe0 // cmp    rax, r12
	JNE  LBB0_12

LBB0_16:
	WORD $0xc985                 // test    ecx, ecx
	JLE  LBB0_33
	WORD $0x8545; BYTE $0xdb     // test    r11d, r11d
	JLE  LBB0_33
	WORD $0x8945; BYTE $0xd8     // mov    r8d, r11d
	LONG $0xa1048d4b             // lea    rax, [r9 + 4*r12]
	LONG $0x24448948; BYTE $0x10 // mov    qword [rsp + 16], rax
	LONG $0x24048b48             // mov    rax, qword [rsp]
	LONG $0xa0148d4a             // lea    rdx, [rax + 4*r12]
	WORD $0x8945; BYTE $0xe3     // mov    r11d, r12d
	LONG $0xe0e38341             // and    r11d, -32
	WORD $0x894d; BYTE $0xe5     // mov    r13, r12
	WORD $0xf749; BYTE $0xdd     // neg    r13
	LONG $0x60788d48             // lea    rdi, [rax + 96]
	LONG $0x247c8948; BYTE $0x08 // mov    qword [rsp + 8], rdi
	LONG $0x04c08348             // add    rax, 4
	LONG $0x24448948; BYTE $0x18 // mov    qword [rsp + 24], rax
	WORD $0xff31                 // xor    edi, edi
	WORD $0x634c; BYTE $0xf9     // movsxd    r15, ecx
	JMP  LBB0_20

LBB0_19:
	LONG $0x01c78348         // add    rdi, 1
	WORD $0x394c; BYTE $0xc7 // cmp    rdi, r8
	JE   LBB0_33

LBB0_20:
	LONG $0x14bf0f4c; BYTE $0x7e // movsx    r10, word [rsi + 2*rdi]
	LONG $0xd7af0f4d             // imul    r10, r15
	LONG $0x24448b48; BYTE $0x20 // mov    rax, qword [rsp + 32]
	LONG $0x3804be0f             // movsx    eax, byte [rax + rdi]
	LONG $0xc02acac5             // vcvtsi2ss    xmm0, xmm6, eax
	WORD $0xf983; BYTE $0x20     // cmp    ecx, 32
	JB   LBB0_21
	LONG $0x92048d4a             // lea    rax, [rdx + 4*r10]
	WORD $0x394c; BYTE $0xc8     // cmp    rax, r9
	JBE  LBB0_25
	LONG $0x24048b48             // mov    rax, qword [rsp]
	LONG $0x90048d4a             // lea    rax, [rax + 4*r10]
	LONG $0x24443b48; BYTE $0x10 // cmp    rax, qword [rsp + 16]
	JAE  LBB0_25

LBB0_21:
	WORD $0xdb31 // xor    ebx, ebx

LBB0_28:
	WORD $0x8949; BYTE $0xde       // mov    r14, rbx
	LONG $0x01c4f641               // test    r12b, 1
	JE   LBB0_30
	LONG $0x13048d4a               // lea    rax, [rbx + r10]
	LONG $0x240c8b48               // mov    rcx, qword [rsp]
	LONG $0x0c59fac5; BYTE $0x81   // vmulss    xmm1, xmm0, dword [rcx + 4*rax]
	LONG $0x28244c8b               // mov    ecx, dword 40[rsp] /* [rbp + 16] */
	LONG $0x5872c1c4; WORD $0x990c // vaddss    xmm1, xmm1, dword [r9 + 4*rbx]
	LONG $0x117ac1c4; WORD $0x990c // vmovss    dword [r9 + 4*rbx], xmm1
	WORD $0x8949; BYTE $0xde       // mov    r14, rbx
	LONG $0x01ce8349               // or    r14, 1

LBB0_30:
	WORD $0xf748; BYTE $0xd3     // not    rbx
	WORD $0x394c; BYTE $0xeb     // cmp    rbx, r13
	JE   LBB0_19
	LONG $0x24448b48; BYTE $0x18 // mov    rax, qword [rsp + 24]
	LONG $0x90048d4a             // lea    rax, [rax + 4*r10]

LBB0_32:
	LONG $0x597aa1c4; WORD $0xb04c; BYTE $0xfc // vmulss    xmm1, xmm0, dword [rax + 4*r14 - 4]
	LONG $0x587281c4; WORD $0xb10c             // vaddss    xmm1, xmm1, dword [r9 + 4*r14]
	LONG $0x117a81c4; WORD $0xb10c             // vmovss    dword [r9 + 4*r14], xmm1
	LONG $0x597aa1c4; WORD $0xb00c             // vmulss    xmm1, xmm0, dword [rax + 4*r14]
	LONG $0x587281c4; WORD $0xb14c; BYTE $0x04 // vaddss    xmm1, xmm1, dword [r9 + 4*r14 + 4]
	LONG $0x117a81c4; WORD $0xb14c; BYTE $0x04 // vmovss    dword [r9 + 4*r14 + 4], xmm1
	LONG $0x02c68349                           // add    r14, 2
	WORD $0x394d; BYTE $0xf4                   // cmp    r12, r14
	JNE  LBB0_32
	JMP  LBB0_19

LBB0_25:
	LONG $0x0479e3c4; WORD $0x00c8 // vpermilps    xmm1, xmm0, 0
	LONG $0x1875e3c4; WORD $0x01c9 // vinsertf128    ymm1, ymm1, xmm1, 1
	LONG $0x24448b48; BYTE $0x08   // mov    rax, qword [rsp + 8]
	LONG $0x901c8d4a               // lea    rbx, [rax + 4*r10]
	WORD $0xc031                   // xor    eax, eax

LBB0_26:
	LONG $0x5459f4c5; WORD $0xa083             // vmulps    ymm2, ymm1, yword [rbx + 4*rax - 96]
	LONG $0x5c59f4c5; WORD $0xc083             // vmulps    ymm3, ymm1, yword [rbx + 4*rax - 64]
	LONG $0x6459f4c5; WORD $0xe083             // vmulps    ymm4, ymm1, yword [rbx + 4*rax - 32]
	LONG $0x2c59f4c5; BYTE $0x83               // vmulps    ymm5, ymm1, yword [rbx + 4*rax]
	LONG $0x586cc1c4; WORD $0x8114             // vaddps    ymm2, ymm2, yword [r9 + 4*rax]
	LONG $0x5864c1c4; WORD $0x815c; BYTE $0x20 // vaddps    ymm3, ymm3, yword [r9 + 4*rax + 32]
	LONG $0x585cc1c4; WORD $0x8164; BYTE $0x40 // vaddps    ymm4, ymm4, yword [r9 + 4*rax + 64]
	LONG $0x5854c1c4; WORD $0x816c; BYTE $0x60 // vaddps    ymm5, ymm5, yword [r9 + 4*rax + 96]
	LONG $0x117cc1c4; WORD $0x8114             // vmovups    yword [r9 + 4*rax], ymm2
	LONG $0x117cc1c4; WORD $0x815c; BYTE $0x20 // vmovups    yword [r9 + 4*rax + 32], ymm3
	LONG $0x117cc1c4; WORD $0x8164; BYTE $0x40 // vmovups    yword [r9 + 4*rax + 64], ymm4
	LONG $0x117cc1c4; WORD $0x816c; BYTE $0x60 // vmovups    yword [r9 + 4*rax + 96], ymm5
	LONG $0x20c08348                           // add    rax, 32
	WORD $0x3949; BYTE $0xc3                   // cmp    r11, rax
	JNE  LBB0_26
	WORD $0x894c; BYTE $0xdb                   // mov    rbx, r11
	WORD $0x394d; BYTE $0xe3                   // cmp    r11, r12
	JE   LBB0_19
	JMP  LBB0_28

LBB0_33:
	SUBQ $8, SP
	VZEROUPPER
	RET

LBB0_34:
	WORD $0xdb31     // xor    ebx, ebx
	LONG $0x28244c8b // mov    ecx, dword 40[rsp] /* [rbp + 16] */
	LONG $0x01c2f641 // test    r10b, 1
	JNE  LBB0_10
	JMP  LBB0_11

TEXT ·_quick_feed(SB), $0-40

	MOVQ hidden_outputs+0(FP), DI
	MOVQ hidden_outputs_len+8(FP), SI
	MOVQ weights+16(FP), DX
	MOVQ weights_len+24(FP), CX
	MOVQ res+32(FP), R8

	WORD $0xc985             // test    ecx, ecx
	JLE  LBB1_1
	WORD $0xc889             // mov    eax, ecx
	LONG $0xc057f8c5         // vxorps    xmm0, xmm0, xmm0
	WORD $0xf983; BYTE $0x20 // cmp    ecx, 32
	JAE  LBB1_4
	WORD $0xc931             // xor    ecx, ecx
	LONG $0xc957f0c5         // vxorps    xmm1, xmm1, xmm1
	JMP  LBB1_7

LBB1_1:
	LONG $0xc957f0c5 // vxorps    xmm1, xmm1, xmm1
	JMP  LBB1_8

LBB1_4:
	WORD $0xc189             // mov    ecx, eax
	WORD $0xe183; BYTE $0xe0 // and    ecx, -32
	LONG $0xc957f0c5         // vxorps    xmm1, xmm1, xmm1
	WORD $0xf631             // xor    esi, esi
	LONG $0xd257e8c5         // vxorps    xmm2, xmm2, xmm2
	LONG $0xdb57e0c5         // vxorps    xmm3, xmm3, xmm3
	LONG $0xe457d8c5         // vxorps    xmm4, xmm4, xmm4
	LONG $0xed57d0c5         // vxorps    xmm5, xmm5, xmm5

LBB1_5:
	LONG $0x345ff4c5; BYTE $0xb7   // vmaxps    ymm6, ymm1, yword [rdi + 4*rsi]
	LONG $0x7c5ff4c5; WORD $0x20b7 // vmaxps    ymm7, ymm1, yword [rdi + 4*rsi + 32]
	LONG $0x445f74c5; WORD $0x40b7 // vmaxps    ymm8, ymm1, yword [rdi + 4*rsi + 64]
	LONG $0x4c5f74c5; WORD $0x60b7 // vmaxps    ymm9, ymm1, yword [rdi + 4*rsi + 96]
	LONG $0x3459ccc5; BYTE $0xb2   // vmulps    ymm6, ymm6, yword [rdx + 4*rsi]
	LONG $0xd258ccc5               // vaddps    ymm2, ymm6, ymm2
	LONG $0x7459c4c5; WORD $0x20b2 // vmulps    ymm6, ymm7, yword [rdx + 4*rsi + 32]
	LONG $0xdb58ccc5               // vaddps    ymm3, ymm6, ymm3
	LONG $0x7459bcc5; WORD $0x40b2 // vmulps    ymm6, ymm8, yword [rdx + 4*rsi + 64]
	LONG $0x7c59b4c5; WORD $0x60b2 // vmulps    ymm7, ymm9, yword [rdx + 4*rsi + 96]
	LONG $0xe458ccc5               // vaddps    ymm4, ymm6, ymm4
	LONG $0xed58c4c5               // vaddps    ymm5, ymm7, ymm5
	LONG $0x20c68348               // add    rsi, 32
	WORD $0x3948; BYTE $0xf1       // cmp    rcx, rsi
	JNE  LBB1_5
	LONG $0xca58e4c5               // vaddps    ymm1, ymm3, ymm2
	LONG $0xc958dcc5               // vaddps    ymm1, ymm4, ymm1
	LONG $0xc958d4c5               // vaddps    ymm1, ymm5, ymm1
	LONG $0x197de3c4; WORD $0x01ca // vextractf128    xmm2, ymm1, 1
	LONG $0xca58f0c5               // vaddps    xmm1, xmm1, xmm2
	LONG $0x0579e3c4; WORD $0x01d1 // vpermilpd    xmm2, xmm1, 1
	LONG $0xca58f0c5               // vaddps    xmm1, xmm1, xmm2
	LONG $0xd116fac5               // vmovshdup    xmm2, xmm1
	LONG $0xca58f2c5               // vaddss    xmm1, xmm1, xmm2
	WORD $0x3948; BYTE $0xc1       // cmp    rcx, rax
	JE   LBB1_8

LBB1_7:
	LONG $0x145ffac5; BYTE $0x8f // vmaxss    xmm2, xmm0, dword [rdi + 4*rcx]
	LONG $0x1459eac5; BYTE $0x8a // vmulss    xmm2, xmm2, dword [rdx + 4*rcx]
	LONG $0xc958eac5             // vaddss    xmm1, xmm2, xmm1
	LONG $0x01c18348             // add    rcx, 1
	WORD $0x3948; BYTE $0xc8     // cmp    rax, rcx
	JNE  LBB1_7

LBB1_8:
	LONG $0x117ac1c4; BYTE $0x08 // vmovss    dword [r8], xmm1
	VZEROUPPER
	RET
