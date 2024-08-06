package huffman

type HNode struct {
	Letter     rune
	Weight     int
	Parent     int
	LeftChild  int
	RightChild int
}

func NewHNode(l rune, w int) *HNode {
	return &HNode{
		Letter:     l,
		Weight:     w,
		Parent:     -1,
		LeftChild:  -1,
		RightChild: -1,
	}
}

func NewDefaultHNode() *HNode {
	return &HNode{
		Letter:     ' ',
		Weight:     999,
		Parent:     -1,
		LeftChild:  -1,
		RightChild: -1,
	}
}
