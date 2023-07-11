package binarytree

import (
	"encoding/json"
	"fmt"
	"os"
)

type TreeNodeInput struct {
  Id string `json:"id"`
  Left string `json:"left"`
  Right string `json:"right"`
  Value string `json:"value"`
}

type TreeNode struct {
  id string
  leftId string
  rightId string
  Value string // value
  Left *TreeNode // pointer to left treenode
  Right *TreeNode // pointer to right treenode
}

func GetBytesFromFile(path string) ([]byte, error) {
  dat, err := os.ReadFile(path)
  if err != nil {
    fmt.Println("Parsing error: Unable to parse raw file into bytes")
    return []byte{}, err 
  }

  return dat, nil
}

// I like the idea of explicitly returing pointers
// so that you are inheratly able to return a nil
// pointer.
// I don't think this is very idomatic go however
// from what I have seen before.
// In this case it also makes sense to have concepts
// of pointers, as the type is recursive in nature
func generateTree(data []TreeNodeInput) *TreeNode {
  nodes := make([]TreeNode, len(data)) // create memory for pointers
  for i, itr := range data {
    currNode := TreeNode{
      id: itr.Id,
      leftId: itr.Left,
      rightId: itr.Right,
      Value: itr.Value,
    }
    nodes[i] = currNode
  }
  for i := range nodes {
    for j := range nodes {
      if nodes[j].id == nodes[i].leftId {
        nodes[i].Left = &nodes[j]
      }
      if nodes[j].id == nodes[i].rightId {
        nodes[i].Right = &nodes[j]
      }
    }
  }
  if len(nodes) == 0 {
    return nil
  }
  return &nodes[0]
}

func CreateFromJson(data []byte) (*TreeNode, error)  {
  var rawData []TreeNodeInput
  err := json.Unmarshal(data, &rawData)
  if err != nil {
    fmt.Println("Parsing error: Unable to parse json")
    return nil, err
  }
  node := generateTree(rawData)
  return node, nil 
}


func InOrderTraversal(rootNode *TreeNode) {
  if rootNode == nil {
    return
  }
  InOrderTraversal(rootNode.Left)
  fmt.Println(rootNode.Value)
  InOrderTraversal(rootNode.Right)
}
