package app


import (
	"fmt"

	"github.com/otiai10/gosseract"
)



func ReadImageText(path string) string {

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(path)
	client.SetWhitelist(`!"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_abcdefghijklmnopqrstuvwxyz{|}~`+"`")

	text, _ := client.Text()
	fmt.Println(text)

	return text
}
