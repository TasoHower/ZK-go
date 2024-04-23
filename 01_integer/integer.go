package integer

import "fmt"

/*
	对应 WTF-ZK 01_integer 部分
	教程参见：https://github.com/WTFAcademy/WTF-zk/blob/main/01_Integer/readme.md
*/

type Integer struct {
	Value int64
}

// 加法
func (c *Integer) Add(a Integer) *Integer {
	c.Value += a.Value
	return c
}

// 减法
func (c *Integer) Sub(a Integer) *Integer {
	c.Value -= a.Value
	return c
}

// 乘法
func (c *Integer) Mul(a Integer) *Integer {
	c.Value *= a.Value
	return c
}

// 欧几里得除法
func (c *Integer)EuclideanDivision(a Integer)(*Integer,*Integer){
	var quotient,remainder Integer// 商和余数

	quotient.Value = c.Value/a.Value
	remainder.Value = c.Sub(Integer{Value: quotient.Value*a.Value}).Value

	return &quotient,&remainder
}

// 加法交换率
func AdditiveCommutativeLaw() {
	// 2+3+4 = 9
	var a1, a2, b, c Integer
	a1.Value = 2
	a2.Value = 2
	b.Value = 3
	c.Value = 4

	fmt.Printf("%d = %d", a1.Add(b).Add(c).Value, a2.Add(c).Add(b).Value)
}

// 乘法交换律
func MultiplicationCommutativeLaw() {
	// 2*3*4 = 24
	var a1, a2, b, c Integer
	a1.Value = 2
	a2.Value = 2
	b.Value = 3
	c.Value = 4

	fmt.Printf("%d = %d", a1.Mul(b).Mul(c).Value, a2.Mul(c).Mul(b).Value)
}

// 加法结合律
func AdditiveLaw(){

	var a1, a2, b, c Integer
	a1.Value = 2
	a2.Value = 2
	b.Value = 3
	c.Value = 4

	fmt.Printf("%d = %d", a1. Add(b).Add(c).Value, a2.Add(*c.Add(b)).Value)
}

// 乘法结合律
func MultiplicationLaw(){

	var a1, a2, b, c Integer
	a1.Value = 2
	a2.Value = 2
	b.Value = 3
	c.Value = 4

	fmt.Printf("%d = %d", a1.Mul(b).Mul(c).Value, a2.Mul(*c.Mul(b)).Value)
}

