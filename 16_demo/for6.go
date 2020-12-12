package main

import "fmt"

func main() {

	/*
		当 j==4 和 j==5 的时候，没有任何输出：标签的作用对象为外部循环，
		因此 i 会直接变成下一个循环的值，而此时 j 的值就被重设为 0，即它的初始值。
		如果将 continue 改为 break，则不会只退出内层循环，而是直接退出外层循环了。
		另外，还可以使用 goto 语句和标签配合使用来模拟循环。
	*/
LABLE1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABLE1
			}
			fmt.Printf("i is:%d,and j is:%d\n", i, j)
		}
	}



	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
