package main

import (
	"fmt"
	"os"
	"sort"
)

const NOMBRE_ARCH_ASC string = "asecendente.txt"
const NOMBRE_ARCH_DESC string = "descendente.txt"

func generarArchivo(nombreArchivo string, slice ...string) {
	file, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for _, val := range slice {
		file.WriteString(val + "\n")
	}
	file.Close()
}

func main() {
	var numCadenas int
	var slice []string

	fmt.Print("Dame la cantidad de cadenas a escribir: ")
	fmt.Scanf("%d", &numCadenas)
	fmt.Scanln() //Un enter se queda atorado en el buffer y nos da una entrada extra.

	for i := 0; i < numCadenas; i++ {
		aux := ""
		fmt.Scanln(&aux)
		slice = append(slice, aux)
	}
	fmt.Println("Slice inicial", slice)
	sort.Strings(slice)
	fmt.Println("Ordenamiento ascendente: ", slice)
	generarArchivo(NOMBRE_ARCH_ASC, slice...)
	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	fmt.Println("Ordenamiento descendente: ", slice)
	generarArchivo(NOMBRE_ARCH_DESC, slice...)
}
