package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/logger"
	"github.com/synchthia/remonpi/util"
)

func main() {
	logger.Init()
	logrus.Printf("[AEHA] Init...")
	signal := [][]int{
		{0x23, 0xCB, 0x26, 0x01, 0x00, 0x24, 0x03, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13F},
	}

	//signal := [][]int{
	//	{0x28, 0x61, 0x3D, 0x13, 0xEC, 0xB8, 0x47},
	//	{0x28, 0x61, 0x6D, 0xFF, 0x00, 0xFF, 0x00},
	//	{0x28, 0x61, 0xCD, 0xFF, 0x00, 0xFF, 0x00},
	//}
	//rsig := [][]int{}

	code := util.SignalToCode(430, signal, 13300)
	for _, v := range code {
		fmt.Printf("%d ", v*1000)
	}
	//fmt.Println(code)

	//for i := 0; i < len(code); i += 2 {
	//	//fmt.Printf("[%d, %d]", i, i+1)
	//	if i < len(code)-1 {
	//		rsig = append(rsig, []int{code[i], code[i+1]})
	//	}
	//}

	//util.CodeToAEHA(rsig)
}
