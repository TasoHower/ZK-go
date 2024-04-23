package prime

import (
	"fmt"
	"math"
)

/*
	质数
	对应 WTF-ZK 02_prime 部分
	教程参见：https://github.com/WTFAcademy/WTF-zk/blob/main/01_Integer/readme.md
*/

// 最常用的方法是埃拉托斯特尼筛法（Sieve of Eratosthenes）。
// 它的逻辑非常简单：先确定一个要搜寻的范围，然后把从到之间的所有质数的倍数剔除掉，剩下的就是范围内的所有素数。
func SieveOfEratosthenes(n int64) {
	var list []bool

	for i := int64(0);i<n;i++ {
		list = append(list,true)
	}

	list[0] = false
	list[1] = false

	for i := int64(2); i < int64(math.Sqrt(float64(n))); i++ {
		for j := int64(2); i*j <n;j++ {
			list[i*j] = false
		}
	}

	for n,b:= range list{
		if b {
			fmt.Println(n)
		}
	}
}