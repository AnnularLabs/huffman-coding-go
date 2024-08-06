package main

import (
	"fmt"

	"github.com/AnnularLabs/huffman-coding-go/huffman"
)

func main() {
	nodeList := huffman.InitHuffmanTree()
	huffman.CreateHuffman(nodeList, len(huffman.LetterFrequencies))
	str := "hello"
	result := huffman.Coding(nodeList, str)
	fmt.Println(result)
}
