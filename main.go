package main

import (
	"fmt"
	"strings"
)

type HNode struct {
	Str        rune
	Weight     int
	Parent     int
	LeftChild  int
	RightChild int
}

// 初始化
func NewHNode(str rune, weight int) *HNode {
	return &HNode{
		Str:        str,
		Weight:     weight,
		Parent:     -1,
		LeftChild:  -1,
		RightChild: -1,
	}
}

// 默认的初始化HNode
func NewDefaultHNode() *HNode {
	return &HNode{
		Str:        ' ',
		Weight:     999,
		Parent:     -1,
		LeftChild:  -1,
		RightChild: -1,
	}
}

// LetterFrequencies 是最小字符的顺序配置
var LetterFrequencies = []rune{
	'q', 'j', '!', '?', 'z',
	'x', 'v', 'k', 'w', 'y',
	'f', 'b', 'g', 'h', 'm',
	'.', 'p', 'd', 'u', ' ',
	'c', 'l', 's', 'n', 't',
	'o', 'i', 'r', 'a', 'e',
}

func InitHuffmanTree() []*HNode {
	nodeList := make([]*HNode, 1)

	for i, char := range LetterFrequencies {
		nodeList = append(nodeList, NewHNode(char, i+1))
	}

	for i := 1; i <= len(LetterFrequencies)-1; i++ {
		nodeList = append(nodeList, NewDefaultHNode())
	}
	return nodeList
}

func CreateHuffman(nodeList []*HNode, nodeListLength int) {
	for i := nodeListLength + 1; i < 2*nodeListLength; i++ {
		minFirst, minSecond := SelectMin(nodeList, i-1)
		nodeList[i].Weight = nodeList[minFirst].Weight + nodeList[minSecond].Weight

		nodeList[minFirst].Parent = i
		nodeList[minSecond].Parent = i

		nodeList[i].LeftChild = minFirst
		nodeList[i].RightChild = minSecond
	}
}

func SelectMin(nodeList []*HNode, selectLength int) (int, int) {
	newHuffmanTree := make([]*HNode, selectLength+1)
	copy(newHuffmanTree, nodeList[:selectLength+1])

	// 筛选出没被选中过的节点
	for i := 1; i <= selectLength; i++ {
		if nodeList[i].Parent != -1 {
			newHuffmanTree[i] = nil
		}
	}

	minFirst := SearchHuffmanTreeMinValue(newHuffmanTree)
	newHuffmanTree[minFirst] = nil
	minSecond := SearchHuffmanTreeMinValue(newHuffmanTree)

	return minFirst, minSecond
}

// 寻找筛选完数组中的最小值
func SearchHuffmanTreeMinValue(newNodeList []*HNode) int {
	i, j := 0, len(newNodeList)-1

	for i < j {
		if newNodeList[i] == nil || (newNodeList[j] != nil && newNodeList[i].Weight >= newNodeList[j].Weight) {
			i++
		} else if newNodeList[j] == nil || newNodeList[j].Weight >= newNodeList[i].Weight {
			j--
		}
	}
	return j
}

func findIndex(word rune) int {
	for i, ch := range LetterFrequencies {
		if ch == word {
			return i
		}
	}
	return -1
}

func Coding(nodeList []*HNode, words string) string {
	var resultCode strings.Builder

	for _, ch := range words {
		child := findIndex(ch) + 1
		fmt.Println(child)
		parent := nodeList[child].Parent

		var codeStack []int

		for parent != -1 {
			if nodeList[parent].LeftChild == child {
				codeStack = append(codeStack, 0)
			} else {
				codeStack = append(codeStack, 1)
			}
			child = parent
			parent = nodeList[child].Parent
		}

		for i := len(codeStack) - 1; i >= 0; i-- {
			resultCode.WriteString(fmt.Sprint(codeStack[i]))
		}

	}

	return resultCode.String()
}

func main() {
	// 初始化 Huffman 树
	nodeList := InitHuffmanTree()

	CreateHuffman(nodeList, len(LetterFrequencies))

	str := "hello"
	result := Coding(nodeList, str)
	fmt.Println(result)
	for i, node := range nodeList {
		if node != nil {
			fmt.Printf("Node %d: Char='%c', Weight=%d, Parent=%d, LeftChild=%d, RightChild=%d\n",
				i, node.Str, node.Weight, node.Parent, node.LeftChild, node.RightChild)
		} else {
			fmt.Printf("Node %d: <nil>\n", i)
		}
	}
}
