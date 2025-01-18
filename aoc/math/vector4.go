package math

type Vector4[T Numbers] struct {
	U0, U1, U2, U3 T
}

func NewVector4[T Numbers](u0, u1, u2, u3 T) Vector4[T] {
	return Vector4[T]{U0: u0, U1: u1, U2: u2, U3: u3}
}

func (v Vector4[T]) Add(av Vector4[T]) Vector4[T] {
	return NewVector4(v.U0+av.U0, v.U1+av.U1, v.U2+av.U2, v.U3+av.U3)
}

func (v Vector4[T]) ManhattanDst(v2 Vector4[T]) T {
	return Abs(v.U0-v2.U0) + Abs(v.U1-v2.U1) + Abs(v.U2-v2.U2) + Abs(v.U3-v2.U3)
}
