package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeIn() []string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	words := strings.Fields(text)
	return words
}

func sharpWriteIn() []string {
	fmt.Print("#> ")
	return writeIn()
}

func BTC_UINT_FORMATER(b uint64) string {
	const unit = 100000000
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.8f %s", float64(b)/float64(div), "BTC")
}
