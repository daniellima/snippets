package main

import (
	"fmt"

	"hello.go/enumDef"
)

func printEnum(e enumDef.Enum) {
	fmt.Println(e)
}

func EnumTestMain() {
	var e enumDef.Enum
	a := enumDef.VALUE2

	b := enumDef.VALUE1
	e = enumDef.VALUE2

	fmt.Println(b)
	printEnum(e)
	fmt.Println(enumDef.VALUE1 == a)
	fmt.Println(enumDef.VALUE2 == a)

	enumDef.VALUE1 = enumDef.VALUE2
	fmt.Println(enumDef.VALUE1)
}
