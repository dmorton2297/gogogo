package main

import (
  "fmt"

  "deepspace/binarytree"
)

func main() {
  fmt.Println("Running in main")
  data, err := binarytree.GetBytesFromFile("thedatar.json")
  if err != nil {
    fmt.Println("Error in main: Exiting")
    return
  }
  node, treeError := binarytree.CreateFromJson(data)

  if treeError != nil {
    fmt.Println(treeError)
  }
  binarytree.InOrderTraversal(node)
}

