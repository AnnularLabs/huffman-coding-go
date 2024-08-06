package huffman

import (
	"fmt"
	"log"
	"strings"
)

func Coding(nodeList []*HNode, word string) string {
	var result strings.Builder

	for _, ch := range word {
		child, err := findIndex(ch)
		if err != nil {
			log.Fatal(err)
		}
		parent := nodeList[child].Parent

		// 模拟栈获取编码
		var stack []int

		// 一直往上找
		for parent != -1 {
			// 当前节点是在左边默认赋值0
			if nodeList[parent].LeftChild == child {
				stack = append(stack, 0)
			} else {
				stack = append(stack, 1)
			}
			child = parent
			parent = nodeList[child].Parent
		}

		// 获取栈中的数据
		for i := len(stack) - 1; i >= 0; i-- {
			result.WriteString(fmt.Sprint(stack[i]))
		}
	}

	return result.String()
}
