package main

import (
	"fmt"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

// go get github.com/kniren/gota/dataframe  (para importar)
// "go build"  en la consola
// https://golang.org/doc/install?download=go1.9.4.windows-amd64.msi
// https://github.com/derekparker/delve/tree/master/Documentation/installation
// http://mclibre.org/consultar/informatica/lecciones/vsc-git-repositorio.html
func main() {
	a := "Arturo is in the town -"
	fmt.Printf("hello, world\n" + a)
	df := dataframe.New(
		series.New([]string{"b", "a"}, series.String, "COL.1"),
		series.New([]int{1, 2}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0}, series.Float, "COL.3"),
	)
	fmt.Println(df)
}
