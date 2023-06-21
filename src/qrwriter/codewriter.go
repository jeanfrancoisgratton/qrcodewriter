// qrcodewriter
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/qrwriter/codewriter.go
// Original time: 2023/06/17 13:57

package qrwriter

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var Port uint16 = 1718

var templ = template.Must(template.New("qr").Parse(templateStr))

func QRwrite() {
	addr := fmt.Sprintf(":%d", Port)
	fmt.Println(fmt.Sprintf("Open your browser on http://localhost:%s", White(strconv.FormatUint(uint64(Port), 10))))
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
