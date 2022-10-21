package go_pkg

func Build() []byte {
	var a = make([]byte, 2)
	a = append(a, 0x01)
	a = append(a, 0x02)
	return a
}
