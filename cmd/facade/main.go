package main

import (
	"design-patterns/structural"
	"fmt"
)

func main() {
	fmt.Println(structural.NewFacade().GetAggregatedData())
}
