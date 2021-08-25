package closure

// https://zhuanlan.zhihu.com/p/56750616

// 闭包里没有引用环境，变量生命周期很短，调用完即释放
func first() func() int {
	return func() int {
		a := 1
		a++
		return a
	}
}

// 闭包里引用全局变量，变量生命周期就是全局变量生命周期
var a = 1

func second() func() int {
	return func() int {
		a++
		return a
	}
}

// 闭包里引用局部变量，变量生命周期长，调用完不释放，下次调用会继续引用
func third() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}

// go tool compile -S closure.go  > closure.S


/**
# 闭包函数

定义
当一个函数的返回值是另外一个函数，而返回的那个函数如果调用了其父函数内部的其他变量，如果返回的这个函数在外部被执行，就产生了闭包。

表现形式
使函数外部能够调用函数内部定义的变量

特性
函数嵌套函数
可以读取函数内部的变量
变量不会被垃圾回收机制回收


 */