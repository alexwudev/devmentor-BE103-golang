package encoder

import (
	"unsafe"
)

var endianness int

func init() {
	var b [2]byte
	*(*uint16)(unsafe.Pointer(&b)) = uint16(0xABCD)

	switch b[0] {
	case 0xCD:
		endianness = 0 // LE
	case 0xAB:
		endianness = 1 // BE
	default:
		panic("could not determine endianness")
	}
}

// "00010203...96979899" cast to []uint16
var intLELookup = [100]uint16{
	0x3030, 0x3130, 0x3230, 0x3330, 0x3430, 0x3530, 0x3630, 0x3730, 0x3830, 0x3930,
	0x3031, 0x3131, 0x3231, 0x3331, 0x3431, 0x3531, 0x3631, 0x3731, 0x3831, 0x3931,
	0x3032, 0x3132, 0x3232, 0x3332, 0x3432, 0x3532, 0x3632, 0x3732, 0x3832, 0x3932,
	0x3033, 0x3133, 0x3233, 0x3333, 0x3433, 0x3533, 0x3633, 0x3733, 0x3833, 0x3933,
	0x3034, 0x3134, 0x3234, 0x3334, 0x3434, 0x3534, 0x3634, 0x3734, 0x3834, 0x3934,
	0x3035, 0x3135, 0x3235, 0x3335, 0x3435, 0x3535, 0x3635, 0x3735, 0x3835, 0x3935,
	0x3036, 0x3136, 0x3236, 0x3336, 0x3436, 0x3536, 0x3636, 0x3736, 0x3836, 0x3936,
	0x3037, 0x3137, 0x3237, 0x3337, 0x3437, 0x3537, 0x3637, 0x3737, 0x3837, 0x3937,
	0x3038, 0x3138, 0x3238, 0x3338, 0x3438, 0x3538, 0x3638, 0x3738, 0x3838, 0x3938,
	0x3039, 0x3139, 0x3239, 0x3339, 0x3439, 0x3539, 0x3639, 0x3739, 0x3839, 0x3939,
}

var intBELookup = [100]uint16{
	0x3030, 0x3031, 0x3032, 0x3033, 0x3034, 0x3035, 0x3036, 0x3037, 0x3038, 0x3039,
	0x3130, 0x3131, 0x3132, 0x3133, 0x3134, 0x3135, 0x3136, 0x3137, 0x3138, 0x3139,
	0x3230, 0x3231, 0x3232, 0x3233, 0x3234, 0x3235, 0x3236, 0x3237, 0x3238, 0x3239,
	0x3330, 0x3331, 0x3332, 0x3333, 0x3334, 0x3335, 0x3336, 0x3337, 0x3338, 0x3339,
	0x3430, 0x3431, 0x3432, 0x3433, 0x3434, 0x3435, 0x3436, 0x3437, 0x3438, 0x3439,
	0x3530, 0x3531, 0x3532, 0x3533, 0x3534, 0x3535, 0x3536, 0x3537, 0x3538, 0x3539,
	0x3630, 0x3631, 0x3632, 0x3633, 0x3634, 0x3635, 0x3636, 0x3637, 0x3638, 0x3639,
	0x3730, 0x3731, 0x3732, 0x3733, 0x3734, 0x3735, 0x3736, 0x3737, 0x3738, 0x3739,
	0x3830, 0x3831, 0x3832, 0x3833, 0x3834, 0x3835, 0x3836, 0x3837, 0x3838, 0x3839,
	0x3930, 0x3931, 0x3932, 0x3933, 0x3934, 0x3935, 0x3936, 0x3937, 0x3938, 0x3939,
}

var intLookup = [2]*[100]uint16{&intLELookup, &intBELookup}

func numMask(numBitSize uint8) uint64 {
	return 1<<numBitSize - 1
}

func AppendInt(_ *RuntimeContext, out []byte, p uintptr, code *Opcode) []byte {
	var u64 uint64
	switch code.NumBitSize {
	case 8:
		u64 = (uint64)(**(**uint8)(unsafe.Pointer(&p)))
	case 16:
		u64 = (uint64)(**(**uint16)(unsafe.Pointer(&p)))
	case 32:
		u64 = (uint64)(**(**uint32)(unsafe.Pointer(&p)))
	case 64:
		u64 = **(**uint64)(unsafe.Pointer(&p))
	}
	mask := numMask(code.NumBitSize)
	n := u64 & mask
	negative := (u64>>(code.NumBitSize-1))&1 == 1
	if !negative {
		if n < 10 {
			return append(out, byte(n+'0'))
		} else if n < 100 {
			u := intLELookup[n]
			return append(out, byte(u), byte(u>>8))
		}
	} else {
		n = -n & mask
	}

	lookup := intLookup[endianness]

	var b [22]byte
	u := (*[11]uint16)(unsafe.Pointer(&b))
	i := 11

	for n >= 100 {
		j := n % 100
		n /= 100
		i--
		u[i] = lookup[j]
	}

	i--
	u[i] = lookup[n]

	i *= 2 // convert to byte index
	if n < 10 {
		i++ // remove leading zero
	}
	if negative {
		i--
		b[i] = '-'
	}

	return append(out, b[i:]...)
}

func AppendUint(_ *RuntimeContext, out []byte, p uintptr, code *Opcode) []byte {
	var u64 uint64
	switch code.NumBitSize {
	case 8:
		u64 = (uint64)(**(**uint8)(unsafe.Pointer(&p)))
	case 16:
		u64 = (uint64)(**(**uint16)(unsafe.Pointer(&p)))
	case 32:
		u64 = (uint64)(**(**uint32)(unsafe.Pointer(&p)))
	case 64:
		u64 = **(**uint64)(unsafe.Pointer(&p))
	}
	mask := numMask(code.NumBitSize)
	n := u64 & mask
	if n < 10 {
		return append(out, byte(n+'0'))
	} else if n < 100 {
		u := intLELookup[n]
		return append(out, byte(u), byte(u>>8))
	}

	lookup := intLookup[endianness]

	var b [22]byte
	u := (*[11]uint16)(unsafe.Pointer(&b))
	i := 11

	for n >= 100 {
		j := n % 100
		n /= 100
		i--
		u[i] = lookup[j]
	}

	i--
	u[i] = lookup[n]

	i *= 2 // convert to byte index
	if n < 10 {
		i++ // remove leading zero
	}
	return append(out, b[i:]...)
}
