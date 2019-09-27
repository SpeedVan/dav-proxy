package dav

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/webdav"
)

func Test(t *testing.T) {
	fs := webdav.Dir("/userfunc")

	f, err := fs.OpenFile(nil, "user", os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("OpenFile:%v\n", err)
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		fmt.Printf("Stat:%v\n", err)
	}
	fmt.Printf("IsDir:%v\n", fi.IsDir())

	// if sth, err := fs.Stat(nil, "user"); err != nil {
	// 	if os.IsNotExist(err) {
	// 		fmt.Printf("IsNotExist:%v\n", err)
	// 	}
	// 	fmt.Printf("Error:%v\n", err)
	// } else {
	// 	fmt.Printf("IsDir:%v\n", sth.IsDir())

	// }
}
