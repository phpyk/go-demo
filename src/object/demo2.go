package main

import "fmt"

/**
面向对象--自定义类型
 */

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)
type Color byte

type Box struct {
	width,height,depth float64
	color Color
}
//给Box对象定义方法：计算体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
//给Box对象定义方法：设置颜色
func (b *Box) SetColor(c Color) {
	b.color = c
}
type BoxList []Box
//给BoxList对象定义方法：求体积最大的盒子的颜色
func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _,b := range bl {
		if bv := b.Volume();bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}
func (bl BoxList) Biggest() Box {
	bb := bl[0]
	v := bb.Volume()
	for _,b := range bl {
		if bv := b.Volume();bv > v {
			v = bv
			bb = b
		}
	}
	return bb
}

func (bl BoxList) PaintItBlack()  {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITH", "BLACK", "BLUE", "RED", "YELLOW",}
	return strings[c]
}

func main() {
	boxes := BoxList{
		{4, 4, 4, RED,},
		{10, 10, 10, YELLOW,},
		{1, 1, 20, BLACK,},
		{3, 6, 20, BLUE,},
		{8, 8, 8, WHITE,},
	}
	fmt.Printf("一共有%d个盒子\n",len(boxes))
	fmt.Println("第一个盒子的体积是：",boxes[0].Volume())
	fmt.Println("最后一个盒子的体积是：",boxes[len(boxes)-1].Volume())
	fmt.Println("体积最大的盒子的颜色是：",boxes.Biggest().color.String())
	fmt.Println("体积最大的盒子的体积是：",boxes.Biggest().Volume())
	fmt.Println("开始全部涂成黑色......")
	boxes.PaintItBlack()
	fmt.Println("第二个盒子的颜色是：",boxes[1].color.String())
	fmt.Println("体积最大的盒子的颜色是：",boxes.BiggestColor().String())
}
