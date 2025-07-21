//有效单词

//题目描述：
/*有效单词 需要满足以下几个条件：

·至少 包含 3 个字符。
·由数字 0-9 和英文大小写字母组成。（不必包含所有这类字符。）
·至少 包含一个 元音字母 。
·至少 包含一个 辅音字母 。
·给你一个字符串 word 。如果 word 是一个有效单词，则返回 true ，否则返回 false 。

注意：
'a'、'e'、'i'、'o'、'u' 及其大写形式都属于 元音字母 。
英文中的 辅音字母 是指那些除元音字母之外的字母。*/

/*示例 1：

输入：word = "234Adas"

输出：true*/


package main

import (
  "unicode"
  "fmt"
)

func isVaild(word string) bool{
  if len(word) < 3{   //快速过滤短字符串
    return false
    }

  var hasVowel, hasConsonant = false, false
  vowels := map[rune]bool{'a'= true, 'e'= true, 'i'= true, 'o'= true, 'u'= true}   //哈希表，时间复杂度为O(1)
  for _, c := range word{
    if !(unicode.IsLetter(c) || unicode.IsDigit(c)){
      return false
  }
    if unicode.IsLetter(c){
      lower := unicode.Tolower(c)   //统一转为小写
      if vowels[lower]{   //若存在键'Lower'，则说明是元音
        hasVowel = true
      } else {
        hasConsonant = true
      }
    }
}
return hasVowel && hasConsonant
}

func main(){
  var testcase = "Abbb233"
  var res = isVaild(testcase)
  fmt.Println(res)
}


-----------------------------------------------------------------------------------------------------
一、代码解析
1.vowels := map[rune]bool{'a'= true, 'e'= true, 'i'= true, 'o'= true, 'u'= true}
  这句代码定义了一个 ​​映射（map）​​，用于快速判断某个字符是否为元音字母。其核心作用是：
  ​​键（Key）​​：元音字母的小写形式（rune类型，即单个Unicode字符）。
​​  值（Value）​​：固定为true，仅作为标记存在（实际只关心键是否存在）。

  map[rune]bool{...}
  • map[rune]bool:声明一个键为rune（Unicode字符）、值为bool的映射。
    => rune是Go语言中用于表示Unicode码点的类型（等价于int32），适合处理字符（如‘a’、‘你’）。  
    => bool类型的值仅用于标记存在性（存在元音时为true，否则默认false）
  • := ：短变量声明语法，自动推断类型

  {'a':true, 'e':true, ...}
  • 语法：通过字面量直接初始化映射，键值对用：分隔，多个键值对用，分隔。
  • 键的类型：所有键必须为rune类型（如'a'表示ASCII字符a）
  • 值的类型：所有值必须为bool类型（此处固定为true）

  ！！！
  • 键的类型必须为可比较类型（如rune、string、int）
  • 默认值：访问不存在的键时返回bool的零值false，可直接用于条件判断：if vowels['x']{...}//自动检查键是否存在，无需显
2.
    

    
  
  
