// qrcodewriter
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/qrwriter/helpers.go
// Original time: 2023/06/21 07:56

package qrwriter

import (
	"fmt"
	"github.com/jwalton/gchalk"
)

func Changelog() {
	//fmt.Printf("\x1b[2J")
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION         DATE                    COMMENT
-------         ----                    -------
0.100           2023.06.21              initial version
\n`)
}

func Red(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightRed().Bold(sentence))
}

func Green(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightGreen().Bold(sentence))
}

func White(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightWhite().Bold(sentence))
}

func Yellow(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightYellow().Bold(sentence))
}
