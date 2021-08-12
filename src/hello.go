package main

import "fmt"

type Man struct {
	age, height int
}

func test1() {
	const (
		x = ' '
		y
		a = 10
		b
		c
	)
	const (
		l = 11 << iota
		m = iota
		n = 13 << iota
		o
	)
	const d int = 10
	const NAME string = "abc"
	const p int = iota
	const q int = iota
	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", d, d)
	fmt.Printf("%T, %v\n", l, l)
	fmt.Printf("%T, %v\n", m, m)
	fmt.Printf("%T, %v\n", n, n)
	fmt.Printf("%T, %v\n", o, o)
	fmt.Printf("%T, %v\n", p, p)
	fmt.Printf("%T, %v\n", q, q)
}
func test2() {
	// if的表达式不用带括号，但是大括号要和if在同一行，尽量都写到同一行
	if true {
		fmt.Println("true")
	}
	// case中不用带break，自动break
	// switch的表达式可以缺省，默认为true
	// fallthrough强制执行下一个case
	switch 1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2, 3, 4:
		fmt.Println("2, 3, 4")
	default:
		fmt.Println("2")
	}
	switch {
	case true:
		fmt.Println("3")
	case false:
		fmt.Println("4")
	default:
		fmt.Println("5")
	}
	// for循环有多种形式
	// 1. 三个表达式全部都存在
	// 2. 变量初始化可以放在外面
	// 3. 判断模块可以省略
	// 4. 三个表达式都不要
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	i := 2
	for ; i < 3; i++ {
		fmt.Printf("%d", i)
	}
	for ; ; i++ {
		if i == 4 {
			fmt.Printf("%d", i)
		}
	}
	for ; ; {
		if i == 5 {
			fmt.Printf("%d", i)
			break
		}
		i++
	}
}
type casefunc func(val int) bool
func main() {
	//test1()
	//test2()
	//a := '一'
	//fmt.Printf("%T, %v", a, a)
	//arr := []int {0, 1, 2, 3, 4, 5, 6}
	//res := filter(arr, isEven)
	//fmt.Println(res)
	//test3()
	// 比如说这里，res1得到的是匿名函数的地址，
	// 和counter这个函数没有任何关系，
	// 他只是一个代理作用，帮你得到这个匿名函数
	// 所以就算再counter中修改变量的值也不会影响匿名函数中值的改变
	//res1 := counter()
	//fmt.Println(res1())
	//fmt.Println(res1())
	// 这里如果重新引用了这个匿名函数，
	// 那么他和之前的匿名函数使用完全没有关联
	// 所以res2虽然和res1保存都是这个匿名函数的地址
	// 但是对于匿名函数的使用是完全分开的
	// res1还是继承上一次res1++,res2从0重新开始
	// 注意：使用匿名函数的使用一定要加括号，
	// 只有添加括号才认为你是通过这个函数地址去执行这个函数
	// 变量中存的是返回值，否则存的就是函数的地址
	//res2 := counter()
	//fmt.Println(res2())
	//fmt.Println(res1())
	//fmt.Printf("%T\n", res1)
	//fmt.Printf("%T", res1())
	auto_exec()
	fmt.Println()
	ptr()
}
func isEven(val int) bool {
	if val % 2 == 0 {
		return true
	}
	return false
}
/*func isOdd(val int) bool {
	if val % 2 == 0
		return false
	return true
}*/
// 函数做变量，回调函数
func filter(arr []int, myfunc casefunc) []int {
	var res []int
	for _, val := range arr {
		if myfunc(val) {
			res = append(res, val)
		}
	}
	return res
}
// 闭包函数
// 匿名函数的使用实际上就是直接访问这个匿名函数指针
// 所以他带有记忆属性
func test3() {
	res := func() int{
		i := 3
		return i
	}()
	fmt.Println(res)
}
func counter() (func() int) {
	i := 0
	return func() int {
		i++
		return i
	}
}
// 可变参数
func auto_exec() {
	scores := []float64 {99, 85, 77.5}
	sums, avg, count := average_sums(scores ...)
	fmt.Println(count)
	fmt.Println(avg)
	fmt.Printf("%.2f", sums)

}
func average_sums(num ... float64) (sums, avg float64, count int64) {
	for _, val := range num {
		count++
		sums += val
	}
	avg = sums / float64(count)
	return sums, avg, count
}
// 指针
func ptr() {
	var ip *int64 = nil
	var a int64 = 3
	ip = &a
	fmt.Println(ip)
	fmt.Println(&ip)
	fmt.Println(*ip)
}