/* Generated by re2c */
//line "golang/002_fill_eof_rule.--lang(go).re":1
//line "golang/002_fill_eof_rule.--lang(go).re":3


package main

import "fmt"
import "os"

var SIZE int = 10

type YYCTYPE byte
type Input struct {
	file   *os.File
	data   []byte
	cursor int
	marker int
	token  int
	limit  int
}

func peek(input *Input) func() YYCTYPE {
	in := input
	return func() YYCTYPE {
		return YYCTYPE(in.data[in.cursor])
	}
}

func skip(input *Input) func() {
	in := input
	return func() {
		in.cursor++
	}
}

func backup(input *Input) func() {
	in := input
	return func() {
		in.marker = in.cursor
	}
}

func restore(input *Input) func() {
	in := input
	return func() {
		in.cursor = in.marker
	}
}

func fill(input *Input) func() int {
	in := input
	return func() int {
		if in.limit < SIZE {
			return 1
		}

		copy(in.data[0:], in.data[in.token:in.cursor])

		in.cursor = in.cursor - in.token
		in.marker = in.marker - in.token
		in.limit = in.limit - in.token
		in.token = 0

		n, _ := in.file.Read(in.data[in.limit:SIZE])
		in.limit = in.limit + n
		in.data[in.limit] = 0
		fmt.Printf("fill: %v '%s'\n", input.data[:input.limit+1], string(input.data[:input.limit]))

		if n == 0 {
			return 1
		}
		return 0
	}
}

func lessthan(input *Input) func(int) bool {
	in := input
	return func(n int) bool {
		return in.limit-in.cursor < n
	}
}

func Lex(input *Input) int {
	YYPEEK := peek(input)
	YYSKIP := skip(input)
	YYBACKUP := backup(input)
	YYRESTORE := restore(input)
	YYFILL := fill(input)
	YYLESSTHAN := lessthan(input)
	input.token = input.cursor

	
//line "golang/002_fill_eof_rule.--lang(go).go":95
{
	var yych YYCTYPE
yyFillLabel0:
	yych = YYPEEK ();
	switch (yych) {
	case '\t':	fallthrough
	case '\n':	fallthrough
	case '\r':	fallthrough
	case ' ':	goto yy4;
	case '0':	fallthrough
	case '1':	fallthrough
	case '2':	fallthrough
	case '3':	fallthrough
	case '4':	fallthrough
	case '5':	fallthrough
	case '6':	fallthrough
	case '7':	fallthrough
	case '8':	fallthrough
	case '9':	goto yy7;
	default:
		if (YYLESSTHAN (1)) {
			if (YYFILL () == 0) {
				goto yyFillLabel0;
			}
			goto yyeofrule2;
		}
		goto yy2;
	}
yy2:
	YYSKIP ();
//line "golang/002_fill_eof_rule.--lang(go).re":95
	{
		fmt.Println("error")
		return -1
	}
//line "golang/002_fill_eof_rule.--lang(go).go":131
yy4:
	YYSKIP ();
yyFillLabel1:
	yych = YYPEEK ();
	switch (yych) {
	case '\t':	fallthrough
	case '\n':	fallthrough
	case '\r':	fallthrough
	case ' ':	goto yy4;
	default:
		if (YYLESSTHAN (1)) {
			if (YYFILL () == 0) {
				goto yyFillLabel1;
			}
		}
		goto yy6;
	}
yy6:
//line "golang/002_fill_eof_rule.--lang(go).re":115
	{
		return 3
	}
//line "golang/002_fill_eof_rule.--lang(go).go":154
yy7:
	YYSKIP ();
	YYBACKUP ();
yyFillLabel2:
	yych = YYPEEK ();
	switch (yych) {
	case '.':	goto yy10;
	case '0':	fallthrough
	case '1':	fallthrough
	case '2':	fallthrough
	case '3':	fallthrough
	case '4':	fallthrough
	case '5':	fallthrough
	case '6':	fallthrough
	case '7':	fallthrough
	case '8':	fallthrough
	case '9':	goto yy7;
	default:
		if (YYLESSTHAN (1)) {
			if (YYFILL () == 0) {
				goto yyFillLabel2;
			}
		}
		goto yy9;
	}
yy9:
//line "golang/002_fill_eof_rule.--lang(go).re":105
	{
		fmt.Printf("number: %v\n", string(input.data[input.token:input.cursor]))
		return 1
	}
//line "golang/002_fill_eof_rule.--lang(go).go":186
yy10:
	YYSKIP ();
yyFillLabel3:
	yych = YYPEEK ();
	switch (yych) {
	case '0':	fallthrough
	case '1':	fallthrough
	case '2':	fallthrough
	case '3':	fallthrough
	case '4':	fallthrough
	case '5':	fallthrough
	case '6':	fallthrough
	case '7':	fallthrough
	case '8':	fallthrough
	case '9':	goto yy12;
	default:
		if (YYLESSTHAN (1)) {
			if (YYFILL () == 0) {
				goto yyFillLabel3;
			}
		}
		goto yy11;
	}
yy11:
	YYRESTORE ();
	goto yy9;
yy12:
	YYSKIP ();
yyFillLabel4:
	yych = YYPEEK ();
	switch (yych) {
	case '0':	fallthrough
	case '1':	fallthrough
	case '2':	fallthrough
	case '3':	fallthrough
	case '4':	fallthrough
	case '5':	fallthrough
	case '6':	fallthrough
	case '7':	fallthrough
	case '8':	fallthrough
	case '9':	goto yy12;
	default:
		if (YYLESSTHAN (1)) {
			if (YYFILL () == 0) {
				goto yyFillLabel4;
			}
		}
		goto yy14;
	}
yy14:
//line "golang/002_fill_eof_rule.--lang(go).re":110
	{
		fmt.Println("number-2")
		return 2
	}
//line "golang/002_fill_eof_rule.--lang(go).go":242
yyeofrule2:
//line "golang/002_fill_eof_rule.--lang(go).re":100
	{
		fmt.Println("end")
		return 0
	}
//line "golang/002_fill_eof_rule.--lang(go).go":249
}
//line "golang/002_fill_eof_rule.--lang(go).re":118

}

func main() {
	tmpfile := "input.txt"
	f, _ := os.Create(tmpfile)
	f.WriteString("     123456789     \n")
	f.Seek(0, 0)

	input := &Input{
		file:   f,
		data:   make([]byte, SIZE+1),
		cursor: SIZE,
		marker: SIZE,
		token:  SIZE,
		limit:  SIZE,
	}

	result := 9999
	for result > 0 {
		result = Lex(input)
	}

	f.Close()
	os.Remove(tmpfile)
}
