"".first STEXT nosplit size=13 args=0x8 locals=0x0
	0x0000 00000 (closure.go:6)	TEXT	"".first(SB), NOSPLIT|ABIInternal, $0-8
	0x0000 00000 (closure.go:6)	FUNCDATA	$0, gclocals·2a5305abe05176240e61b8620e19a815(SB)
	0x0000 00000 (closure.go:6)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:7)	LEAQ	"".first.func1·f(SB), AX
	0x0007 00007 (closure.go:7)	MOVQ	AX, "".~r0+8(SP)
	0x000c 00012 (closure.go:7)	RET
	0x0000 48 8d 05 00 00 00 00 48 89 44 24 08 c3           H......H.D$..
	rel 3+4 t=16 "".first.func1·f+0
"".second STEXT nosplit size=13 args=0x8 locals=0x0
	0x0000 00000 (closure.go:17)	TEXT	"".second(SB), NOSPLIT|ABIInternal, $0-8
	0x0000 00000 (closure.go:17)	FUNCDATA	$0, gclocals·2a5305abe05176240e61b8620e19a815(SB)
	0x0000 00000 (closure.go:17)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:18)	LEAQ	"".second.func1·f(SB), AX
	0x0007 00007 (closure.go:18)	MOVQ	AX, "".~r0+8(SP)
	0x000c 00012 (closure.go:18)	RET
	0x0000 48 8d 05 00 00 00 00 48 89 44 24 08 c3           H......H.D$..
	rel 3+4 t=16 "".second.func1·f+0
"".third STEXT size=152 args=0x8 locals=0x20
	0x0000 00000 (closure.go:25)	TEXT	"".third(SB), ABIInternal, $32-8
	0x0000 00000 (closure.go:25)	MOVQ	(TLS), CX
	0x0009 00009 (closure.go:25)	CMPQ	SP, 16(CX)
	0x000d 00013 (closure.go:25)	PCDATA	$0, $-2
	0x000d 00013 (closure.go:25)	JLS	142
	0x000f 00015 (closure.go:25)	PCDATA	$0, $-1
	0x000f 00015 (closure.go:25)	SUBQ	$32, SP
	0x0013 00019 (closure.go:25)	MOVQ	BP, 24(SP)
	0x0018 00024 (closure.go:25)	LEAQ	24(SP), BP
	0x001d 00029 (closure.go:25)	FUNCDATA	$0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
	0x001d 00029 (closure.go:25)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (closure.go:26)	LEAQ	type.int(SB), AX
	0x0024 00036 (closure.go:26)	MOVQ	AX, (SP)
	0x0028 00040 (closure.go:26)	PCDATA	$1, $0
	0x0028 00040 (closure.go:26)	CALL	runtime.newobject(SB)
	0x002d 00045 (closure.go:26)	MOVQ	8(SP), AX
	0x0032 00050 (closure.go:26)	MOVQ	AX, "".&a+16(SP)
	0x0037 00055 (closure.go:26)	MOVQ	$1, (AX)
	0x003e 00062 (closure.go:27)	LEAQ	type.noalg.struct { F uintptr; "".a *int }(SB), CX
	0x0045 00069 (closure.go:27)	MOVQ	CX, (SP)
	0x0049 00073 (closure.go:27)	PCDATA	$1, $1
	0x0049 00073 (closure.go:27)	CALL	runtime.newobject(SB)
	0x004e 00078 (closure.go:27)	MOVQ	8(SP), AX
	0x0053 00083 (closure.go:27)	LEAQ	"".third.func1(SB), CX
	0x005a 00090 (closure.go:27)	MOVQ	CX, (AX)
	0x005d 00093 (closure.go:27)	PCDATA	$0, $-2
	0x005d 00093 (closure.go:27)	CMPL	runtime.writeBarrier(SB), $0
	0x0064 00100 (closure.go:27)	JNE	126
	0x0066 00102 (closure.go:27)	MOVQ	"".&a+16(SP), CX
	0x006b 00107 (closure.go:27)	MOVQ	CX, 8(AX)
	0x006f 00111 (closure.go:27)	PCDATA	$0, $-1
	0x006f 00111 (closure.go:27)	MOVQ	AX, "".~r0+40(SP)
	0x0074 00116 (closure.go:27)	MOVQ	24(SP), BP
	0x0079 00121 (closure.go:27)	ADDQ	$32, SP
	0x007d 00125 (closure.go:27)	RET
	0x007e 00126 (closure.go:27)	PCDATA	$0, $-2
	0x007e 00126 (closure.go:27)	LEAQ	8(AX), DI
	0x0082 00130 (closure.go:27)	MOVQ	"".&a+16(SP), CX
	0x0087 00135 (closure.go:27)	CALL	runtime.gcWriteBarrierCX(SB)
	0x008c 00140 (closure.go:27)	JMP	111
	0x008e 00142 (closure.go:27)	NOP
	0x008e 00142 (closure.go:25)	PCDATA	$1, $-1
	0x008e 00142 (closure.go:25)	PCDATA	$0, $-2
	0x008e 00142 (closure.go:25)	CALL	runtime.morestack_noctxt(SB)
	0x0093 00147 (closure.go:25)	PCDATA	$0, $-1
	0x0093 00147 (closure.go:25)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 7f 48  eH..%....H;a.v.H
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 8d 05  .. H.l$.H.l$.H..
	0x0020 00 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b 44  ....H..$.....H.D
	0x0030 24 08 48 89 44 24 10 48 c7 00 01 00 00 00 48 8d  $.H.D$.H......H.
	0x0040 0d 00 00 00 00 48 89 0c 24 e8 00 00 00 00 48 8b  .....H..$.....H.
	0x0050 44 24 08 48 8d 0d 00 00 00 00 48 89 08 83 3d 00  D$.H......H...=.
	0x0060 00 00 00 00 75 18 48 8b 4c 24 10 48 89 48 08 48  ....u.H.L$.H.H.H
	0x0070 89 44 24 28 48 8b 6c 24 18 48 83 c4 20 c3 48 8d  .D$(H.l$.H.. .H.
	0x0080 78 08 48 8b 4c 24 10 e8 00 00 00 00 eb e1 e8 00  x.H.L$..........
	0x0090 00 00 00 e9 68 ff ff ff                          ....h...
	rel 5+4 t=17 TLS+0
	rel 32+4 t=16 type.int+0
	rel 41+4 t=8 runtime.newobject+0
	rel 65+4 t=16 type.noalg.struct { F uintptr; "".a *int }+0
	rel 74+4 t=8 runtime.newobject+0
	rel 86+4 t=16 "".third.func1+0
	rel 95+4 t=16 runtime.writeBarrier+-1
	rel 136+4 t=8 runtime.gcWriteBarrierCX+0
	rel 143+4 t=8 runtime.morestack_noctxt+0
"".first.func1 STEXT nosplit size=10 args=0x8 locals=0x0
	0x0000 00000 (closure.go:7)	TEXT	"".first.func1(SB), NOSPLIT|ABIInternal, $0-8
	0x0000 00000 (closure.go:7)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:7)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:10)	MOVQ	$2, "".~r0+8(SP)
	0x0009 00009 (closure.go:10)	RET
	0x0000 48 c7 44 24 08 02 00 00 00 c3                    H.D$......
"".second.func1 STEXT nosplit size=23 args=0x8 locals=0x0
	0x0000 00000 (closure.go:18)	TEXT	"".second.func1(SB), NOSPLIT|ABIInternal, $0-8
	0x0000 00000 (closure.go:18)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:18)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:19)	MOVQ	"".a(SB), AX
	0x0007 00007 (closure.go:19)	INCQ	AX
	0x000a 00010 (closure.go:19)	MOVQ	AX, "".a(SB)
	0x0011 00017 (closure.go:20)	MOVQ	AX, "".~r0+8(SP)
	0x0016 00022 (closure.go:20)	RET
	0x0000 48 8b 05 00 00 00 00 48 ff c0 48 89 05 00 00 00  H......H..H.....
	0x0010 00 48 89 44 24 08 c3                             .H.D$..
	rel 3+4 t=16 "".a+0
	rel 13+4 t=16 "".a+0
"".third.func1 STEXT nosplit size=19 args=0x8 locals=0x0
	0x0000 00000 (closure.go:27)	TEXT	"".third.func1(SB), NOSPLIT|NEEDCTXT|ABIInternal, $0-8
	0x0000 00000 (closure.go:27)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:27)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (closure.go:27)	MOVQ	8(DX), AX
	0x0004 00004 (closure.go:28)	MOVQ	(AX), CX
	0x0007 00007 (closure.go:28)	INCQ	CX
	0x000a 00010 (closure.go:28)	MOVQ	CX, (AX)
	0x000d 00013 (closure.go:29)	MOVQ	CX, "".~r0+8(SP)
	0x0012 00018 (closure.go:29)	RET
	0x0000 48 8b 42 08 48 8b 08 48 ff c1 48 89 08 48 89 4c  H.B.H..H..H..H.L
	0x0010 24 08 c3                                         $..
go.cuinfo.packagename. SDWARFINFO dupok size=0
	0x0000 63 6c 6f 73 75 72 65                             closure
"".a SNOPTRDATA size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".first.func1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".first.func1+0
"".second.func1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".second.func1+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*struct { F uintptr; a *int }- SRODATA dupok size=32
	0x0000 00 00 1d 2a 73 74 72 75 63 74 20 7b 20 46 20 75  ...*struct { F u
	0x0010 69 6e 74 70 74 72 3b 20 61 20 2a 69 6e 74 20 7d  intptr; a *int }
type.*struct { F uintptr; "".a *int } SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a8 2a eb 0e 08 08 08 36 00 00 00 00 00 00 00 00  .*.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; a *int }-+0
	rel 48+8 t=1 type.noalg.struct { F uintptr; "".a *int }+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..namedata..F- SRODATA dupok size=5
	0x0000 00 00 02 2e 46                                   ....F
type..namedata.a- SRODATA dupok size=4
	0x0000 00 00 01 61                                      ...a
type.noalg.struct { F uintptr; "".a *int } SRODATA dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 bf fb c7 5a 02 08 08 19 00 00 00 00 00 00 00 00  ...Z............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; a *int }-+0
	rel 44+4 t=6 type.*struct { F uintptr; "".a *int }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.noalg.struct { F uintptr; "".a *int }+80
	rel 80+8 t=1 type..namedata..F-+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.a-+0
	rel 112+8 t=1 type.*int+0
gclocals·2a5305abe05176240e61b8620e19a815 SRODATA dupok size=9
	0x0000 01 00 00 00 01 00 00 00 00                       .........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·263043c8f03e3241528dfae4e2812ef4 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
