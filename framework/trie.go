package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

type node struct {
	isLast  bool              // 该节点是否能成为一个独立的uri；是否自身就是一个终极节点
	segment string            // uri 中的字符串
	handler ControllerHandler // 节点包含控制器
	childs  []*node           // 这个节点下的子节点
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
}

func newNode() *node {
	return &node{
		isLast:  false,
		segment: "",
		childs:  []*node{},
	}
}

// isWildSegment 判断segment是否以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childs) == 0 {
		return nil
	}

	// 如果是通配符，则下层节点都满足需求
	if isWildSegment(segment) {
		return n.childs
	}

	nodes := make([]*node, 0, len(n.childs))
	// 过滤所有的下一层子节点
	for _, cnode := range n.childs {
		if isWildSegment(cnode.segment) {
			// 如果下一层子节点有通配符，则满足需求
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			// 如果下一层没有通配符，但是文本完全匹配，则满足需求
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	// 使用分隔符将uri切割为两个部分
	segments := strings.SplitN(uri, "/", 2)
	// 第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}
	// 匹配符合的下一层子节点
	cnodes := n.filterChildNodes(segment)
	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在的，直接返回nil
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		// 如果segmment已经是最后一个节点，判断这些cnode是否有isLast标志
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		// 都不是最后一个节点
		return nil
	}

	// 如果有两个segment，递归每个子结点继续进行查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// AddRouter 增加路由节点
func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root
	// 确认路由是否冲突
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")
	// 对每个segment
	for index, segment := range segments {
		// 最终进入 Node segment 的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1
		// 标记是否有合适的子节点
		var objNode *node

		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的子节点
		if len(childNodes) > 0 {
			// 如果有segment相同的segment，则选择这个子节点
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前node的节点
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.childs = append(n.childs, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	// 直接复用 matchNode 函数，uri是不带通配符的地址
	matchNode := tree.root.matchNode(uri)
	if matchNode != nil {
		return matchNode.handler
	}
	return nil
}
