package main

// 编译器自动进行逃逸分析
func foo()*int{
	var x int
	return &x
}

func bar()int{
	x:=new(int)
	return *x
}

func main(){

}
