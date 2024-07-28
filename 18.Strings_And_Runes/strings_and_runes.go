package main

import (
	"fmt"
	"unicode/utf8"
)


// A GO string is a read-only slice of bytes. The language and the standard library treat strings specially -as containers of text encoded in UTF-8. In other languages, strings are made of "characters". In GO, the concept of a character is called a "rune" - it's an integer that represents a Unicode code point. 
func main(){
	const s = "สวัสดี" // s is a string assigned to a literal value representing the world "hello" in Thai language. GO string literals are UTF-8 encoded text.

	fmt.Println("Len:", len(s))// since strings are equivalent to []byte, this will show length of the raw bytes stored within.

	for i := 0; i< len(s);i++{
		fmt.Printf("%x ",s[i])// indexing into a string produces the raw byte values at each index. generates the hex values of all the bytes that consitute the code point in s. 
	}

	fmt.Println()

	fmt.Println("Rune Count: ", utf8.RuneCountInString(s)) // to count how many runes are in a string , the "utf8 // to count how many runes are in a string , the "utf8" package is used. NOTE the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially. Some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising.

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx) // a range loop handles strings specially and decodes each rune along with its offset in the string
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i,  w:=0,0; i< len(s); i+= w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])	// the same iteration can be achieved by using the utf8.DecodeRuneInString function explicitly

		fmt.Printf("%#U starts at %d\n", runeValue, i)

		w  = width

		examineRune(runeValue) // this is how a rune value is passed to a function
	}
}

func examineRune(r rune){
    if r == 't' {	// values enclosed in single quotes are "rune literals" , we can compare a "rune value" to a "rune literal" directly
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }	
}