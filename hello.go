package main

import (
	"fmt"

	"github.com/kniren/gota/dataframe"
	"github.com/kniren/gota/series"
)

// go get github.com/kniren/gota/dataframe  (para importar)
// "go build"  en la consola
// git config --global user.email b3mol@hotmail.com
// https://golang.org/doc/install?download=go1.9.4.windows-amd64.msi
// https://github.com/derekparker/delve/tree/master/Documentation/installation
// http://mclibre.org/consultar/informatica/lecciones/vsc-git-repositorio.html
func main() {
	a := "Arturo is in the town -"
	fmt.Printf("hello, world\n" + a)
	df := dataframe.New(
		series.New([]string{"b", "a", "g", "g", "g"}, series.String, "COL.1"),
		series.New([]int{1, 2, 3, 4, 5}, series.Int, "COL.2"),
		series.New([]float64{3.0, 4.0, 4.2, 4.3, 4.5}, series.Float, "COL.3"),
	)
	fmt.Println(df)
}
