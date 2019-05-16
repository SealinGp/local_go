package main;

import (
	"fmt"
);

func init() {
  fmt.Println("Content-Type:text/plain;charset=utf-8\n\n");
}
func main() {	
	u1 := multiply_self(2,4,true);

	fmt.Println(u1);
}

var u uint64;

//计算u1的m1次幂 u1^m1,(注:golang不支持默认参数)
func multiply_self(u1 uint64,m1 uint8,first bool) (uint64) {
	if m1 <= 1 {
		return u1;
	}
	
	if first {
		u = u1;
	}

	return multiply_self(u1*u,m1-1,false);
}
/*

默认函数参数,其实并不是一个很好的行为.我认为代码,
应该是没有任何隐喻的,需要其他人去查看的,仅从名字
就能知道大概的行为.而默认函数参数,可能导致调用者
不知道默认参数是什么出现一些问题.我不知道具体Go
这么做的原因是什么,但是在我概念中,一个行为一定是
确定的,无二意的,而默认参数的函数会可能导致错误.
Google的C++规范里面也禁止使用函数默认参数.类似的,
在这个规范的,凡是有可能导致二意行为的也是被禁止的.
比如如果没有必要的情况下,禁止类的默认拷贝构造函数,
避免某些情况下莫名的调用了,等等的.

作者：codedump
链接：https://www.zhihu.com/question/24368980/answer/84708971
来源：知乎
著作权归作者所有.商业转载请联系作者获得授权,非商业转载请注明出处.
*/