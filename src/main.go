package main

import (
	"fmt"
	"log"

	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Skype\Phone`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	var v uint32 = 0
	k.SetDWordValue("DisableFileTransfer", v)
	s, _, err := k.GetIntegerValue("DisableFileTransfer")
	fmt.Printf("Done %q", s)
	if err := k.Close(); err != nil {
		log.Fatal(err)
	}
}
