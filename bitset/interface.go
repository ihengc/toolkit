package bitset

/*
* @author: Heng ChenChi
* @date: 2022/12/21 0021 10:07
* @version: 1.0
* @description:
**/

type BitSetInterface interface {
	ClearAll()
	Clear(bitIndex int)
	ClearRange(fromIndex, toIndex int)

	Flip(bitIndex int)
	FlipRange(fromIndex, toIndex int)

	Get(bitIndex int) bool
	GetRange(fromIndex, toIndex int) bool

	Set(bitIndex int, value bool)
	SetRange(fromIndex, toIndex int, value bool)

	IsEmpty() bool
	Len() int
	Size() int
	Cardinality() int

	And(bitSet BitSetInterface)
	AndNot(bitSet BitSetInterface)
	Or(bitSet BitSetInterface)
	XOR(bitSet BitSetInterface)
}
