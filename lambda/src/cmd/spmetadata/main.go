package main

import (
	"fmt"
	"spmetadata"
)

func main() {
	entityDescriptor, _ := spmetadata.GetEntityDescriptor()
	fmt.Println(entityDescriptor)
}
