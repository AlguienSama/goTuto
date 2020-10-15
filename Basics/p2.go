package main

import (
	"fmt"
	"strings"
)

/*
	POINTERS
	Funcionan obteniendo la dirección de memoria de la variable
*/
func pointers() {

	var printResults = func(x int, p *int) {
		fmt.Printf("Type of x = %T\nType of p = %T\nValue of x = %v\nValue of p = %v\n",
			x, p, x, *p)
	}

	i := 42
	p := &i

	printResults(i, p)

	*p /= 2

	printResults(i, p)

}

/*
	ESTRUCTURAS
	Parecido a un objeto, estructuras de datos
*/
func structs() {

	// Se define la estructura
	type Structure struct {
		X int
		Y string
	}

	// Se definen los valores
	s := Structure{2, "Hello"}
	fmt.Println("First structure =>", s)
	s.Y = "World!"
	fmt.Println("Y modify =>", s)

	// Se pueden usar pointers
	p := &s
	p.X = 4
	fmt.Println("Use with Pointer =>", p)

	s2 := Structure{Y: "Hello World!"}
	fmt.Println("Si no pones algún valor, lo deja por defecto => ", s2)
}

/*
	ARRAYS
	Pues un array
*/
func arrays() {

	// Declarar un array vacío
	var a [2]string
	a[0] = "Hello"
	a[1] = "World !"
	fmt.Println(a)

	// Declarar array con valores
	a2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a2)

	// Posiciones de los array
	var a3 []int = a2[1:4]
	fmt.Println(a3)

	a4 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}
	fmt.Println(a4)

	/*
		len(array) -> Largada
		cap(array) -> Capacidad
		make([]arrayType, length, capacity?) -> Nuevo array
		strings.Join(array, valor) -> String con los valores del array
		append(array, values...) -> Añadir valores a un array
		delete(array, index) -> Elimina un elemento del array
	*/

	fmt.Printf("A2 Tamaño actual = %d - Capacidad máxima = %d - Array = %v\n", len(a2), cap(a2), a2)

	a5 := make([]int, 2, 5)
	fmt.Printf("A5 Tamaño actual = %d - Capacidad máxima = %d - Array = %v\n", len(a5), cap(a5), a5)

	// Multidimensional array
	tabla := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	tabla[1][1] = "X"

	fmt.Println("Tabla")
	for i := 0; i < len(tabla); i++ {
		fmt.Printf("%s\n", strings.Join(tabla[i], " "))
	}

	// La capacidad máxima debe ser variable []
	fmt.Printf("A3 Tamaño actual = %d - Capacidad máxima = %d - Array = %v\n", len(a3), cap(a3), a3)
	a3 = append(a3, 7, 8, 9)
	fmt.Printf("A3 Tamaño actual = %d - Capacidad máxima = %d - Array = %v\n", len(a3), cap(a3), a3)

	// FOREACH
	// for _, v := range a { only take value }
	// for i := range a { only take index }
	for index, value := range a3 {
		fmt.Printf("i -> %d v -> %d | ", index, value)
	}
	fmt.Println()

	// ARRAY ASOCIATIVO / MAP
	type Coordenadas struct {
		Lat, Long float64
	}

	var a7 map[string]Coordenadas

	a7 = make(map[string]Coordenadas)
	a7["Ubicación1"] = Coordenadas{
		0.4234, 1.2432,
	}

	fmt.Println(a7)

	var a8 = map[string]Coordenadas{
		"Ubicación1": {0.4234, 1.2432},
		"Ubicación2": {0.4234, 1.2432},
	}
	fmt.Println(a8)

	// Retorna el valor (en caso que no haya, por defecto) y boolean de si existe o no
	ubicacion, existe := a8["UbicaciónX"]
	fmt.Println(ubicacion, existe)
}

/*
	FUNCIONES
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func runAdder() {
	num1, num2 := adder(), adder()

	for i := 0; i < 4; i++ {
		fmt.Println(i, " -> ", num1(i), num2(i+i))
	}
}

type Numbers struct {
	N1, N2 float64
}

func (n Numbers) Suma() float64 {
	return n.N1 + n.N2
}

func runMethods() {
	n := Numbers{3, 4}
	fmt.Printf("%v + %v = %v\n", n.N1, n.N2, n.Suma())
	n.increment(2.0)
	fmt.Println(n)
}

func (n *Numbers) increment(i float64) {
	n.N1 += i
	n.N2 += i
}

func main() {
	fmt.Println("Hello World")
	fmt.Println("\n --- POINTERS --- ")
	pointers()
	fmt.Println("\n --- STRUCTS --- ")
	structs()
	fmt.Println("\n --- ARRAYS --- ")
	arrays()
	fmt.Println("\n --- FUNCIONES --- ")
	runAdder()
	runMethods()
}
