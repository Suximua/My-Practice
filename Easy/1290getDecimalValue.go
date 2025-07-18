//二进制链表转整数

//题目描述：给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。
//请你返回该链表所表示数字的 十进制值 。

//示例：输入：head = [1,0,1]
//输出：5
//解释：二进制数 (101) 转化为十进制数 (5)

package main

import(
  "fmt"
  "math/big"
)

type ListNode struct{
  Val int
  Next *ListNode
}

func getDecimalValue(head *ListNode) int{
  res := 0   //初始化结果为0
  current := head   //当前节点指针初始化为头节点
  for current ！= nil{   //遍历链表直到末尾
    res = res * 2 + current.Val   // 核心计算【res*2 => 左移一位（等价于乘以2的幂次）+ current.Val => 当前位的二进制】
    current = current.Next
  }
  return res
}

//测试示例
func main(){   // 创建测试链表1->0->1
  head := &ListNode{Val:1}
  head.Next := &ListNode{Val:0}
  head.Next.Next := &ListNode{Val:1}

  decimal := getDecimalValue(head)

  fmt.Println("链表[1,0,1]转换为十进制是：", decimal)
}


-----------------------------------------------------------------------------------------------------
一、创建结构体递归结构
1.链表
线性数据结构，由一系列节点组成，每个节点包含两部分：数据部分和指向下一个节点的指针（双向链表会有指向前一个节点的指针）
链表不必连续存储，在插入和删除操作上具有天然优势

链表的结构体：
type ListNode struct{
  Val int //数据部分
  Next *ListNode //指针
}
创建一个链表，并打印节点内容：
head := &ListNode{Val: 1}
node2 := &ListNode{Val: 2}
node3 := &ListNode{Val: 3}
head.Next = node2
node2.Next = node3
node := head
for {
    if node == nil {
        break
    } else {
        fmt.Println(node.Val)
        node = node.Next
    }
}

链表在实际场景中的应用：
·实现栈和队列
·解决数据频繁插入和删除的场景

2.二叉树
二叉树的结构体：
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}
创建一个二叉树，并用前序遍历打印二叉树节点：
func preorderTraversal(root *TreeNode){
  if root == nil{
    return
  }
  fmt.Println(root.Val)
  if root.Left != nil{
    preorderTraversal(root.Left)
  }
  if root.Right != nil{
    preorderTraversal(root.Right)
  }
}
root := &TreeNode{Val:1}
root.Left = &TreeNode{Val:2}
root.Right = &TreeNode{Val:3}
root.Left.Left = &TreeNode{Val:4}
root.Left.Right = &TreeNode{Val:5}
preorderTraversal(root)
