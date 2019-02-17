package main

import (
  "fmt"
)

// type Power struct{
//         age int
//         high int
//         name string
// }
// func printSlice(x []int){
//    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
// }
// const MAX = 3
// const name, age = "Kim", 22

// func fibo(n int) int{
//  if n < 2{
//     return n
//  }
//  return fibo(n-2) + fibo(n-1)
// }
// type Man interface {
//     name() string;
//     age() int;
// }

// type Woman struct {
// }

// func (woman Woman) name() string {
//    return "Jin Yawei"
// }
// func (woman Woman) age() int {
//    return 23;
// }

// type Men struct {
// }

// func ( men Men) name() string {
//    return "liweibin";
// }
// func ( men Men) age() int {
//     return 27;
// }

// 定义一个 DivideError 结构
type DivideError struct {
  dividee int
  divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
  strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
  return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
  if varDivider == 0 {
    dData := DivideError{
      dividee: varDividee,
      divider: varDivider,
    }
    errorMsg = dData.Error()
    return
  } else {
    return varDividee / varDivider, ""
  }

}

func main() {
  var m map[int]map[int]string
  m = make(map[int]map[int]string)
  m[1] = make(map[int]string)
  m[1][1] = "OK"
  a, ok := m[1][1]

  fmt.Println(a, ok)
  // // 正常情况
  //     if result, errorMsg := Divide(100, 10); errorMsg == "" {
  //         fmt.Println("100/10 = ", result)
  //     }
  //     // 当被除数为零的时候会返回错误信息
  //     if _, errorMsg := Divide(100, 0); errorMsg != "" {
  //         fmt.Println("errorMsg is: ", errorMsg)
  //     }

  // num := 12
  // for i:=0; i<num; i++{
  //  fmt.Println(fibo(i))
  // }

  //   var sum int = 17
  //   var count int = 5
  //   var mean float32

  //   mean = float32(sum)/float32(count)
  //   fmt.Printf("mean 的值为: %.4f\n",mean)

  // Print inserts blanks between arguments when neither is a string.
  // It does not add a newline to the output, so we add one explicitly.
  // fmt.Print("The vector (", a, b, ") has length ", h, ".\n")

  // Println always inserts spaces between its arguments,
  // so it cannot be used to produce the same output as Print in this case;
  // its output has extra spaces.
  // Also, Println always adds a newline to the output.
  // fmt.Println("The vector (", a, b, ") has length", h, ".")

  // Printf provides complete control but is more complex to use.
  // It does not add a newline to the output, so we add one explicitly
  // at the end of the format specifier string.
  // fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)

  // fmt.Printf("a b %d %d %d cd\n", 1, 2, 3)
  // var i Power = Power{age: 10, high: 178, name: "NewMan"}
  // fmt.Printf("type:%T\n", i)
  //    fmt.Printf("value:%v\n", i)
  //    fmt.Printf("value+:%+v\n", i)
  //    fmt.Printf("value#:%#v\n", i)

  //        fmt.Println("========interface========")
  //        var interf interface{} = i
  //        fmt.Printf("%v\n", interf)
  //        fmt.Println(interf)
  //        fmt.Println(i)

  // var test uint64
  // test = 92233720368547758

  // fmt.Printf("%t\n", test
  // type person struct{
  //  name string
  //  age int
  // }
  // var i = person{"bruce",24}
  // fmt.Println(i)
  //greats := [5]string {"Katano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
  //    var greats = []string{"Katano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}

  // fmt.Printf("%v %#q\n", greats, greats)

  //     hash := sha256.New()
  //     hash.Write([]byte("Hello world\n"))
  //     fmt.Printf("The hash is %x",hash.Sum(nil))
  // const(
  //             a = 1   //1
  //             b          //1
  //             c          //1
  //             d = "ha"   //独立值，iota += 1
  //             e          //"ha"   iota += 1
  //             f = 100    //iota +=1
  //             g          //100  iota +=1
  //             h = iota   //7,恢复计数
  //             i          //8
  //     )
  //     fmt.Println(a,b,c,d,e,f,g,h,i)

  /* 定义局部变量 */
  // var i, j int

  // for i=2; i < 100; i++ {
  //    for j=2; j <= (i/j); j++ {
  //       if(i%j==0) {
  //          break; // 如果发现因子，则不是素数
  //       }
  //    }
  //    if(j > (i/j)) {
  //       fmt.Printf("%d  是素数\n", i);
  //    }
  // }

  // var a int
  // var ptr *int
  // var pptr **int

  // a = 3000

  // /* 指针 ptr 地址 */
  // ptr = &a

  // /* 指向指针 ptr 地址 */
  // pptr = &ptr

  //  获取 pptr 的值
  // fmt.Printf("变量 a = %d\n", a )
  // fmt.Printf("指针变量 *ptr = %x\n", ptr )
  // fmt.Printf("指向指针的指针变量 **pptr = %x\n", pptr)
  // var countryCapitalMap map[string]string /*创建集合 */
  // countryCapitalMap = make(map[string]string)

  // /* map插入key - value对,各个国家对应的首都 */
  // countryCapitalMap [ "France" ] = "Paris"
  // countryCapitalMap [ "Italy" ] = "罗马"
  // countryCapitalMap [ "Japan" ] = "东京"
  // countryCapitalMap [ "India " ] = "新德里"
  // countryCapitalMap ["美国"] = "Washington DC"

  // /*使用键输出地图值 */
  // for country := range countryCapitalMap {
  //     fmt.Println(country, "首都是", countryCapitalMap [country])
  // }

  // /*查看元素在集合中是否存在 */
  // captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
  // fmt.Println(captial)
  // fmt.Println(ok)
  // if (ok) {
  //     fmt.Println("美国的首都是", captial)
  // } else {
  //     fmt.Println("美国的首都不存在")
  // }

}
