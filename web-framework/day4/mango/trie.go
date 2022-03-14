package mango

import (
	"log"
	"strings"
)

// 1. 明确需求。观察main函数里路由注册的规则，一个 Method，一个路由 pattern，一个处理函数，需要处理这个 pattern
// - 搜索是否有匹配这个 pattern 的注册路由
// - 把新增的 pattern 记录下来，插入前缀树中
// - 请求到来的时候查询时候有对应的路由并处理结果
// - 匹配规则有两种 /:abc, /*
// - 解析路由的时候，需要把变量赋值保存
// 2. 设计数据结构
// 3. 设计函数处理
// 4. 写单元测试，覆盖全面
// 5. 代码简洁化：多余的分支，提前变量赋值 DRY

type node struct {
	// 完整路径, 最后的节点才有值 eg: /p/:lang/doc
	pattern string
	// 当前部分路径, eg: doc, :lang
	part string
	// 子节点, eg [doc, :lang, other]
	children []*node
	// 是否是模糊匹配, part 包含 : 或 * 时为 true
	isWild bool
}

// 查找子节点里第一个匹配成功的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// 查找子节点里所有匹配成功的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// 插入节点
func (n *node) insert(pattern string, parts []string, heigth int) {
	if len(parts) == heigth {
		n.pattern = pattern
		return
	}

	part := parts[heigth]
	child := n.matchChild(parts[heigth])
	if child == nil {
		child = &node{part: parts[heigth], isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, heigth+1)
}

// 查找匹配模式的节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	log.Println("current part:", part, height)

	for _, child := range children {
		log.Println("child:", child)
		matchChild := child.search(parts, height+1)
		if matchChild != nil {
			return matchChild
		}
	}

	return nil
}
