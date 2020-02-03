package main 

import "fmt"

// claim

var (
	x int 
	y int 
	z int 
	python bool 
	java bool 
)



//claim
/*
var x, y, z int
var python, java bool
*/


//initial
/*
var x, y, z int = 1, 2, 3 //有起始值可以省略型別，會直接取用初始化的型別
var python, java =  false, "no!" 
*/

//short claim(:=)  It can replace var when it is used in the func.
/*
func main(){
	var x, y, z int = 1, 2, 3
	python, java := true, "no!"
	fmt.Println(x, y, z, python ,java)
}
*/



func main(){
	fmt.Println(x, y, z, python, java)
}

/*
Conclusion 
	var a  // 不定型別的變數
    var a int // 宣告成 int
    var a int = 10 // 初始化同時宣告
    var a, b int // a 跟 b 都是 intvar a, b = 0, ""
    var a int , b string
    a := 0
    a, b, c := 0, true, "tacolin" // 這樣就可以不同型別寫在同一行
    var(
        a bool = false // 記得要不同行，不然會錯
        b int
        c = "hello"
    )
*/