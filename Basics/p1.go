// Nombre del paquete inicial siempre main
package main

// Adición de paquetes
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

/**** - VARIABLES Y DECLARACIONES - ****/

// var se puede usar en cualquier parte del fichero
// Se debe de especificar siempre el tipo al final
// Se puede iniciar más de una a la a vez
var globalVariable, boolean bool
var num1, num2 int = 1, 2
var num3, boolean2, num4 = 3, true, 4

var (
	boolean3   bool       = false
	maxInt     uint64     = 1<<64 - 1
	complexNum complex128 = cmplx.Sqrt(-5 + 12i)
	floatNum   float64    = math.Sqrt(float64(num3*num3 + num4*num4))
	uintNum    uint       = uint(floatNum)
)

// const nunca con :=
const num5 = 5

func variables() {
	// := Solo dentro de funciones,
	// establece el tipo del valor asignado

	// Si no se usan las variables, salta error
	//num5, num6, string1 := 5, 6, "string1"
}

/*
	TYPES

	bool
	string
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte // alias for uint8
	rune // alias for int32
     	 // represents a Unicode code point
	float32 float64
	complex64 complex128

	DEFAULT VALUES
	0, false, ""

	null = nil
*/

/*
	%T	Mostrar tipo variable
	%v	Valor
		%t	Boolean
		%d	int uint
		%g	float complex
		%s	string
		%p 	chan pointer address
*/

/*
	CÁLCULOS
		BINARIO
	multiplicador << cantidad zeros = posición binaria*multiplicador = resultado
	1 << 3	=  1000*1  =  8
	2 << 3  =  1000*2  =  16
	2 << 2  =  100*2  = 8

		POTENCIAS
	2**2 = 4
	2**3 = 8
*/

/**** - FUNCIONES - ****/

// Se tienen que especificar que tipo de variables se pasan
// Y que tipo se retornan (en caso de return)
func suma(x int, y int) int {
	return x + y
}

// Si son del mismo tipo, con solo ponerlo el último ya sirve
func resta(x, y int) int {
	return x - y
}

// Se pueden retornar múltiples valores
// Pero se debe especificar cada uno
func swap(x, y string) (string, string) {
	return y, x
}

// Se pueden retornar variables especificando cuales
func splitNum(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}

// Parentesis opcionales (los del if, for . . .)

func loops() int {
	numLoops := 0

	// FOR loop
	for i := 0; i < 10; i++ {
		numLoops++
	}
	for numLoops < 12 {
		numLoops++
	}

	if j := 12; j == numLoops {
		return j
	}
	return numLoops

}

func switchFuns() {
	fmt.Print("Estás usando ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
		break
	case "linux":
		fmt.Println("Linux")
		break
	default:
		fmt.Printf("%s\n", os)
	}
	fmt.Println()

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Función inicial siempre main
func main() {
	// Mostrar por pantalla
	var helloWorld = "Hello World!"
	defer fmt.Println(helloWorld, "defer se evalua en seguida pero la función se ejecuta al final")
	fmt.Println(helloWorld, "Date Now:", time.Now())
	fmt.Println(helloWorld, "Random Number:", rand.Intn(10))
	fmt.Println(helloWorld, "4 + 5 =", suma(4, 5))
	fmt.Println(helloWorld, "5 - 4 =", resta(5, 4))
	splitString := strings.Split(helloWorld, " ")
	a, b := swap(splitString[0], splitString[1])
	fmt.Println(helloWorld, "swapped =", a, b)

	// Si se retorna más de un valor,
	// no se puede poner junto al Print
	fmt.Print(helloWorld, " Number Split 14 = ")
	fmt.Println(splitNum(14))
	fmt.Println(helloWorld, "Boolean 1, 2 =", boolean, boolean2)
	fmt.Println(helloWorld, "num1 num2 num3 num4 =", num1, num2, num3, num4)
	fmt.Println(helloWorld, "UINT =", uintNum)
	switchFuns()
	fmt.Println()

	fmt.Println("Defer example")
	fmt.Println("counting")
	i := 0
	for ; i < 2; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Resultado", i)
}
