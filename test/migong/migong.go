package main

import (
	"os"
	"bufio"
	"strings"
	"io"
	"strconv"
)
type point struct { //对应迷宫的位置
	z,h int
}
//上下左右四个点的计算方法
func (p *point)Up() point{
	var up point
	up.z ,up.h = p.z+1,p.h
	return up
}

func (p *point)Left() point{
	var left point
	left.z ,left.h = p.z,p.h-1
	return left
}
func (p *point)Down() point{
	var down point
	down.z ,down.h = p.z-1,p.h
	return down
}
func (p *point)Right() point{
	var right point
	right.z ,right.h = p.z,p.h+1
	return right
}

type mini struct {
//	sep int               //值
	walkLine []point	  //坐标
}
//判断这个点是不是在地图中 地图是2维坐标
func (p *point )isPoint(max_z,max_h int) bool{

	where := false
	if  0 <= p.z && p.z <= max_z && 0 <= p.h  && p.h <= max_h{
		where = true
	}
	return where
}
//计算出最优路径
func walk(ditu [][]int,startP,endP *point) mini {

	var A []point //已发现未探索
	var B []point    // 已发现已探索
	var C mini    //未发现
	h := len(ditu[0])
	z := len(ditu)
	// 从开始 ， 上下左右 探索
	for {
		p := startP.Up()
		if p.isPoint(z, h) {
			A = append(A, p)
		}
		startP.Left()
		if p.isPoint(z, h) {
			A = append(A, p)
		}
		startP.Down()
		if p.isPoint(z, h) {
			A = append(A, p)
		}
		startP.Right()
		if p.isPoint(z, h) {
			A = append(A, p)
		}
	}

	for {
		if len(A) == 0 {
			break
			}
		p := A[0]
		p.z,p.h :=
		}
	}

func readFile(filename string)[][]int{
	maze := [][]int{}
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
		}
	buf := bufio.NewReader(f)
	lineNum := -1
	for {
		lineNum ++
		//一次读一行
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		//每行开启一个新的row，加入到 maze中
		func(lineNum int,line string){
			numString := strings.Split(line," ")
			row := []int{}
			for _,v :=  range numString {
				num1,_ := strconv.Atoi(v)
				row= append(row,int(num1))
			}
			maze = append(maze, row)
		}(lineNum,line)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	 return maze
}

func main() {
	a := readFile("test/migong/maze1")
	start := point{0,0}
	end := point{4,5}
	walk(a,&start,&end)
}