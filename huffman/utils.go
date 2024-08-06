package huffman

import "errors"

func SelectMin(nodeList []*HNode, length int) (int, int) {

	newHuffmanTree := make([]*HNode, length+1)
	copy(newHuffmanTree, nodeList[:length+1])

	// 筛选出没有被选中过的节点，parent不等于1的
	for i := 1; i <= length; i++ {
		if nodeList[i].Parent != -1 {
			newHuffmanTree[i] = nil
		}
	}

	// 获取newHuffmanTree中权重最小的值
	minFirst := searchHuffmanTreeMinValue(newHuffmanTree)
	// 将minFirst排除
	newHuffmanTree[minFirst] = nil
	// 获取newHuffmanTree中权重最小的值
	minSecond := searchHuffmanTreeMinValue(newHuffmanTree)
	return minFirst, minSecond
}

func findIndex(letter rune) (int, error) {
	for i, ch := range LetterFrequencies {
		if ch == letter {
			return i, nil
		}
	}
	return -1, errors.New("not found letter")
}

func searchHuffmanTreeMinValue(nodeList []*HNode) int {
	i, j := 0, len(nodeList)-1
	for i < j {
		if nodeList[i] == nil || (nodeList[j] != nil && nodeList[i].Weight >= nodeList[j].Weight) {
			i++
		} else if nodeList[j] == nil || nodeList[j].Weight >= nodeList[i].Weight {
			j--
		}
	}
	return j
}
