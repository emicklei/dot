package dot

import (
	"bytes"
	"fmt"
	"testing"
)

func TestIndentWriter(t *testing.T) {
	b := new(bytes.Buffer)
	i := NewIndentWriter(b)
	i.WriteString("doc {")
	i.NewLineIndentWhile(func() {
		fmt.Fprint(i, "chapter {")
		i.NewLineIndentWhile(func() {
			fmt.Fprint(i, "chapter text")
		})
		i.WriteString("}")
	})
	i.WriteString("}")
	fmt.Println(b.String())
}
