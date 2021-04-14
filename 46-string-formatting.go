package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p) // {1, 2}

	fmt.Printf("%+v\n", p) // include struct names: {x:1, y:2}

	fmt.Printf("%#v\n", p) // Go syntax representation: main.point{x:1, y:2}

	fmt.Printf("%T\n", p) // type of a value: main.point

	fmt.Printf("%t\n", true) // boolean: true

	fmt.Printf("%d\n", 123) // decimal, standard base-10: 123

	fmt.Printf("%b\n", 14) // binary representation: 1110

	fmt.Printf("%c\n", 33) // character corresponding to the integer: !

	fmt.Printf("%x\n", 456) // hex encoding: 1c8

	fmt.Printf("%f\n", 78.9) // basic decimal formatting for float: 78.900000

	// %e and %E formats float in (slightly different versions of) scientific notation
	fmt.Printf("%e\n", 123400000.0) // 1.234000e+08
	fmt.Printf("%E\n", 123400000.0) // 1.234000E+08

	fmt.Printf("%s\n", "\"string\"") // basic string print: "string"

	fmt.Printf("%q\n", "\"string\"") // double quote strings: "\"string\""

	fmt.Printf("%x\n", "hex this") // renders string in base-16, with two output characters per byte of input: 6865782074686973

	fmt.Printf("%p\n", &p) // pointer: 0x42135100

	// Width precision of numbers
	// Default right-justified, padded with space
	fmt.Printf("|%6d|%6d|\n", 12, 345) // |    12|   345|

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // width precision with decimal precision: |  1.20|  3.45|

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // Use - flat to left justify: |1.20  |3.45  |

	fmt.Printf("|%6s|%6s|\n", "foo", "b") // width precision for strings: |   foo|     b|

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // Use - flat to left justify: |foo   |b     |

	// Sprintf formats ands returns a string (without printing it anywhere)
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Use Fprintf to format and print to io.Writers other than os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
