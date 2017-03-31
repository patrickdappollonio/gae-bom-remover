package main

import (
	"bytes"
	"fmt"
	"strings"
)

const main = `<!DOCTYPE html>
<html lang="en">
  <head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Byte-order mark remover</title>

	<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
	<!--[if lt IE 9]>
	  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
	  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
	<![endif]-->
  </head>
  <body>
#####
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  </body>
</html>
`

func getTemplate(data string) string {
	data = strings.TrimSpace(data)

	var buf bytes.Buffer
	for _, line := range strings.Split(data, "\n") {
		buf.WriteString(fmt.Sprintf("\t%s\n", line))
	}

	return strings.Replace(main, "#####", buf.String(), 1)
}
