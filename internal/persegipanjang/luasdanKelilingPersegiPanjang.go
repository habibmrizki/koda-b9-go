package persegipanjang

func LuasPersegiPanjang(panjang, lebar uint8) uint8 {
	return panjang * lebar
}

func KelilingPersegiPanjang(panjang, lebar uint8) uint8 {
	return 2 * (panjang + lebar)
}

func LuasDanKelilingPersegiPanjang(panjang, lebar uint8) (luas uint8, keliling uint8) {
	return LuasPersegiPanjang(panjang, lebar), KelilingPersegiPanjang(panjang, lebar)
}
