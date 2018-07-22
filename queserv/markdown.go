package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type NodeType uint32

const (
	NodePara = iota
	NodeImag
	NodeOl
	NodeLi
	NodeRoot
)

type Node struct {
	FirstChild  *Node
	NextSibling *Node

	Type NodeType
	Data string
	Attr string   // e.g "width=100" for NodeImag, "type=A." for list
	List []string // just for NodeList
}

func AppendSibling(node *Node, sibling *Node) *Node {
	if node == nil {
		node = sibling
		return node.NextSibling
	}
	for node.NextSibling != nil {
		node = node.NextSibling
	}
	node.NextSibling = sibling
	return node.NextSibling
}

func AppendChild(p, child *Node) *Node {
	if p == nil {
		fmt.Println("nil parent node")
		return nil
	}
	if p.FirstChild == nil {
		p.FirstChild = child
		return p.FirstChild
	}
	p = p.FirstChild
	for p.NextSibling != nil {
		p = p.NextSibling
	}
	p.NextSibling = child
	return p.NextSibling
}

func parseImag(parent *Node, line string) *Node {
	if parent == nil {
		parent = &Node{Type: NodeImag}
	}
	j0 := strings.Index(line, "(")
	j1 := strings.Index(line, ")")
	k0 := strings.Index(line, "[")
	k1 := strings.Index(line, "]")
	parent.Data = line[j0+1 : j1]
	tmp := line[k0+1 : k1]
	parent.Attr = tmp[strings.Index(tmp, "=")+1:]
	return parent
}

func parseOl(parent *Node, par string) *Node {
	if parent == nil {
		parent = &Node{Type: NodeOl}
	}
	rd := bufio.NewReader(strings.NewReader(par))
	AppendChild(parent, parseLi(nil, rd))
	return parent
}

func parseLi(p *Node, rd *bufio.Reader) (firstLi *Node) {
	var isinpar bool = false
	var par *Node
	line, err := rd.ReadString('\n')
	for {
		if err != nil && !(err == io.EOF && len(line) > 0) {
			break
		}

		islist, c := isList(line)
		if islist { // new li
			if isinpar {
				isinpar = false
			}

			node := &Node{Type: NodeLi}
			if c == 'A' {
				node.Attr = "type=\"A\""
			} else if c == '1' {
				node.Attr = "type=\"1\""
			} else if c == '*' {
				node.Attr = "type=\"disc\""
			}
			if p == nil {
				firstLi, p = node, node
			} else {
				p = AppendSibling(p, node)
			}
			line = line[2:]
			continue

		} else { // in old li
			if isImage(line) { // li>imag
				if isinpar {
					isinpar = false
				}
				AppendChild(p, parseImag(nil, line))
			} else { // li > para
				if isinpar {
					par.Data += line
				} else {
					isinpar = true
					par = AppendChild(p, &Node{Type: NodePara, Data: line})
				}
			}
			goto NextIter
		}

	NextIter:
		line, err = rd.ReadString('\n')
	}
	return firstLi
}

func parsePara(node *Node, par string) *Node {
	node = &Node{Type: NodePara, Data: par}
	return node
}

func isList(par string) (bool, rune) {
	if len(par) < 3 {
		return false, '&'
	}
	if par[1] == '.' && par[2] == ' ' {
		if '0' <= par[0] && par[0] <= '9' {
			return true, '1'
		} else if 'A' <= par[0] && par[0] <= 'Z' {
			return true, 'A'
		} else if par[0] == '*' {
			return true, '*'
		}
	}
	return false, '&'
}

func isImage(par string) bool {
	if len(par) < 3 {
		return false
	}
	if par[0] == '!' && par[1] == '[' {
		return true
	}
	return false
}

// Parse parse s which is markdown string, and
// return it's structe governd by Node.
func Parse(s string) *Node {
	root := &Node{Type: NodeRoot}

	s = strings.Replace(s, "\r", "", -1)
	pars := strings.Split(s, "\n\n")
	for _, par := range pars {
		if isImage(par) {
			AppendChild(root, parseImag(nil, par))
			continue
		}
		if islist, _ := isList(par); islist == true {
			AppendChild(root, parseOl(nil, par))
			continue
		}
		AppendChild(root, parsePara(nil, par))
	}

	return root
}

// markdownToHTML translate markdown string s to html string
func markdownToHTML(s string) string {
	s = strings.Replace(s, "_blank", "________", -1)
	s = strings.Replace(s, "\\kh", "(   )", -1)
	s = strings.Replace(s, "~", " ", -1)
	result := ""
	root := Parse(s)
	for node := root.FirstChild; node != nil; node = node.NextSibling {
		if node.Type == NodePara {
			result += fmt.Sprintf("<p> %s </p>", node.Data)
		} else if node.Type == NodeImag {
			result += fmt.Sprintf("<p> <img src=\"%s\" width=\"%s\" align=\"top\" /> </p>", node.Data, node.Attr)
		} else if node.Type == NodeOl {
			result += "<ol>"
			for linode := node.FirstChild; linode != nil; linode = linode.NextSibling {
				result += fmt.Sprintf("<li %s>\n", linode.Attr)
				for p := linode.FirstChild; p != nil; p = p.NextSibling {
					if p.Type == NodePara {
						result += fmt.Sprintf("<p>%s</p>", p.Data)
					} else if p.Type == NodeImag {
						result += fmt.Sprintf("<p> <img src=\"%s\" width=\"%s\" align=\"top\" /> </p>", p.Data, p.Attr)
					}
				}
				result += fmt.Sprintf("</li>\n")
			}
			result += "</ol>"
		}
	}
	return result
}
