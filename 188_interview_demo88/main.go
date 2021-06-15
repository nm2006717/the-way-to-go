package main

type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	//	错误做法
	//	1、map 需要初始化才能使用；
	//	2、指针不支持索引。修复代码如下：
	//s := new(Show)
	//s.Param["day"] = 2

	//s:=new(Show)
	//p:=make(Param)
	//p["day"] = 2
	//s.Param = &p
	//tmp := *s.Param
	//fmt.Println(tmp["day"])
}
