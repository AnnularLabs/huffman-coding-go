package huffman

func InitHuffmanTree() []*HNode {

	// 排除下标为0的位置
	nodeList := make([]*HNode, 1)

	// 将字符存入nodeList中，并且权重递增
	for i, letter := range LetterFrequencies {
		nodeList = append(nodeList, NewHNode(letter, i+1))
	}

	// 增加huffman树需要的节点个数
	for i := 1; i <= len(LetterFrequencies)-1; i++ {
		nodeList = append(nodeList, NewDefaultHNode())
	}

	return nodeList

}

func CreateHuffman(nodeList []*HNode, length int) {

	// 从空白节点生成
	for i := length + 1; i < 2*length; i++ {
		minFirst, minSecond := SelectMin(nodeList, i-1)

		// 将当前节点的权重赋值为两个最小值权重的和
		nodeList[i].Weight = nodeList[minFirst].Weight + nodeList[minSecond].Weight
		// 两个最小值的parent指向当前节点
		nodeList[minFirst].Parent = i
		nodeList[minSecond].Parent = i
		// 当前节点的left等于minF，right等于minS
		nodeList[i].LeftChild = minFirst
		nodeList[i].RightChild = minSecond
	}

}
