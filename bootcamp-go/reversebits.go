package bootcamp

func ReverseBits(n byte) byte {
	var reversed byte = 0
	for i := 0; i < 8; i++ {
		reversed <<= 1
		reversed |= (n & 1)
		n >>= 1
	}
	return reversed
}
