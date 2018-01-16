package main

import (
	"fmt"
)

func main() {
	fmt.Println("red-black-tree")
	var data = []int{6, 5, 3, 1, 8, 7, 2, 4, 9, 0, 3}
	//var data = []int{6, 5, 3, 1}

	tree := NewTree()
	for _, v := range data {
		tree.insertB(v)
	}
	//fmt.Println(tree.size)
	tree.tdelete(tree.root, 2)
	printT(tree.root)
}

type node struct {
	l, r, p *node
	v       int
	c       bool //0 = red 1=black
}

type Tree struct {
	root *node
	last *node
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

func (n *node) getGp() *node {
	if n.p == nil {
		return nil
	}
	if n.p.p == nil {
		return nil
	}
	return n.p.p
}

func (n *node) getUn() *node {
	if n.getGp() == nil {
		return nil
	}
	if n.p == n.getGp().l {
		return n.getGp().r
	} else {
		return n.getGp().l
	}
}

// rotate cur node
func (tree *Tree) rotateR(n *node) {
	gp := n.getGp()
	p := n.p
	r := n.r

	n.r = p
	n.p = gp

	if p != nil {
		p.p = n
		p.l = r
		if r != nil {
			r.p = p
		}
	}
	if tree.root == p {
		tree.root = n
	}

	if gp != nil {
		if gp.l == p {
			gp.l = n
		} else {
			gp.r = n
		}
	}
}

func (tree *Tree) rotateL(n *node) {
	gp := n.getGp()
	p := n.p
	l := n.l

	n.l = p
	n.p = gp

	if p != nil {
		p.p = n
		p.r = l
		if l != nil {
			l.p = p
		}
	}

	if tree.root == p {
		tree.root = n
	}

	if gp != nil {
		if gp.l == p {
			gp.l = n
		} else {
			gp.r = n
		}
	}

}

func getMin() {

}

func getMax() {

}

func (tree *Tree) tdelete(n *node, v int) {
	if tree == nil {
		return
	}
	if v < n.v {
		tree.tdelete(n.l, v)
	}
	if v > n.v {
		tree.tdelete(n.r, v)
	}

	if n.v == v {
		if tree.deleteR(n) {
			tree.size--
		}
	}
}

func (tree *Tree) deleteR(n *node) bool {
	if n.l == nil && n.r == nil {
		if n.p != nil {
			if n.p.l == n {
				n.p.l = nil
			} else {
				n.p.r = nil
			}
		}
		n.p = nil
		n.v = 0
		n.l = nil
		n.r = nil
		n.c = false
		return false
	} else if n.l != nil && n.r == nil {
		if n.p != nil {
			if n.p.l == n {
				n.p.l = n.l
			} else {
				n.p.r = n.l
			}
		}
		n.l.p = n.p
	} else if n.r != nil && n.l == nil {
		if n.p != nil {
			if n.p.l == n {
				n.p.l = n.r
			} else {
				n.p.r = n.r
			}
		}
		n.r.p = n.p
	} else if n.r != nil && n.l != nil {
		//选择n节点的左子树中最大节点
		maxn := n.l
		for maxn.r != nil {
			maxn = maxn.r
		}
		n.v = maxn.v
		if maxn == n.l {
			maxn.p.l = n.l
		} else {
			maxn.p.r = maxn.l
		}
		fmt.Println(maxn)
		maxn.p = nil
		maxn.v = 0
		maxn.l = nil
		maxn.r = nil
		maxn.c = false
		return true
	}
	n.p = nil
	n.v = 0
	n.l = nil
	n.r = nil
	n.c = false
	return true
}

//修复红黑树平衡
func fix() {

}

func (tree *Tree) insertB(v int) {
	if tree.root == nil {
		tree.root = new(node)
		tree.root.v = v
		tree.root.c = true
		tree.size++
	}

	if v < tree.root.v {
		if tree.Insert(&tree.root.l, v, tree.root) {
			printT(tree.root)
			tree.size++
			tree.inserCase(tree.last)
		}
	}

	if v > tree.root.v {
		if tree.Insert(&tree.root.r, v, tree.root) {
			printT(tree.root)
			tree.size++
			tree.inserCase(tree.last)
		}
	}
	printT(tree.root)
}

func (tree *Tree) Insert(n **node, v int, fa *node) bool {

	pn := (*n)
	if (*n) == nil {
		(*n) = new(node)
		(*n).v = v
		(*n).p = fa
		tree.last = (*n)
		return true
	}

	if v > pn.v {
		tree.Insert(&(pn.r), v, *n)
	}

	if v < pn.v {
		tree.Insert(&(pn.l), v, *n)
	}

	if v == pn.v {
		return false
	}
	return true
}

func (tree *Tree) inserCase(n *node) {
	if n.p == nil {
		n.c = true
		tree.root = n
		return
	}

	if n.p.c == false {
		if n.getUn() != nil && n.getUn().c == false {
			fmt.Println("aaaaa", n)
			n.p.c = true
			n.getUn().c = true
			n.getGp().c = false
			tree.inserCase(n.getGp()) //if root node
		} else {
			//nil or black
			if n == n.p.r && n.getGp() != nil && n.p == n.getGp().l {
				tree.rotateL(n)
				tree.rotateR(n)
				n.c = true
				n.l.c = false
				n.r.c = false
			}
			if n == n.p.l && n.getGp() != nil && n.p == n.getGp().r {
				tree.rotateR(n)
				tree.rotateL(n)
				n.c = true
				n.l.c = false
				n.r.c = false
			}
			if n == n.p.l && n.getGp() != nil && n.p == n.getGp().l {

				n.p.c = true
				if n.getGp() != nil {
					n.getGp().c = false
				}
				tree.rotateR(n.p)
			}
			if n == n.p.r && n.getGp() != nil && n.p == n.getGp().r {
				n.p.c = true
				if n.getGp() != nil {
					n.getGp().c = false
				}
				tree.rotateL(n.p)
			}
		}

	}
}

//test print ----------------------------------------------
var fstr = make([]string, 20, 20)

func printT(tree *node) {
	fmt.Println(tree)
	fstr = make([]string, 8, 8)
	printTree(tree, 0, "")
	for i, str := range fstr {
		fmt.Println("L", i, str)
	}
}

func printTree(tree *node, i int, n string) {
	fmt.Println("-----", tree)
	i++
	str := " "
	for n := i; n < 9; n++ {
		str += "-"
	}
	var tmp string
	if tree.c == true {
		tmp = fmt.Sprintf("%s\033[40;37m [%d] \033[0m%s%s", str, tree.v, n, str)
	} else {
		tmp = fmt.Sprintf("%s\033[41;37m [%d] \033[0m%s%s", str, tree.v, n, str)
	}

	fstr[i] += tmp
	if tree.l != nil {
		printTree(tree.l, i, "L"+fmt.Sprintf("%d", tree.v))
	}
	if tree.r != nil {
		printTree(tree.r, i, "R"+fmt.Sprintf("%d", tree.v))
	}
}
