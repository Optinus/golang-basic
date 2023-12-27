package main

import (
	"github.com/Optinus/golang-basic/array"
	"github.com/Optinus/golang-basic/condition"
	"github.com/Optinus/golang-basic/constan"
	"github.com/Optinus/golang-basic/function"
	"github.com/Optinus/golang-basic/loop"
	"github.com/Optinus/golang-basic/maps"
	"github.com/Optinus/golang-basic/operator"
	"github.com/Optinus/golang-basic/slice"
	"github.com/Optinus/golang-basic/struc"
	"github.com/Optinus/golang-basic/switc"
	"github.com/Optinus/golang-basic/typedata"
	"github.com/Optinus/golang-basic/variable"
)

func main() {
	// Memanggil fungsi-fungsi dari paket-paket yang diimpor
	variable.Variable()   // Contoh penggunaan variabel
	constan.Constan()     // Contoh penggunaan konstanta
	typedata.TypeData()   // Contoh penggunaan tipe data
	array.Array()         // Contoh penggunaan array
	slice.Slice()         // Contoh penggunaan slice
	operator.Operator()   // Contoh penggunaan operator
	condition.Condition() // Contoh penggunaan kondisi
	switc.Switch()        // Contoh penggunaan switch
	loop.Loop()           // Contoh penggunaan loop
	function.Function()   // Contoh penggunaan fungsi
	struc.Struct()        // Contoh penggunaan struktur data
	maps.Maps()           // Contoh penggunaan map
}
