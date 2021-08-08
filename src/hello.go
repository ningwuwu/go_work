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
func main() {
	//test1()
	test2()
	//a := '一'
	//fmt.Printf("%T, %v", a, a)
}
