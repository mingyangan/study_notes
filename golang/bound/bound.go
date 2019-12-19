package main

func f1(is []int, bs []byte) {
	if len(is) >= 256 {
		for _, n := range bs {
			_ = is[n]
		}
	}
}

func f2(is []int, bs []byte) {
	if len(is) >= 256 {
		is = is[:len(is)]
		for _, n := range bs {
			_ = is[n]
		}
	}
}
