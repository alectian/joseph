package CircleLink

import (
	"fmt"
	"time"
)

type People struct {
	Num int
	PrePeople *People
	NextPeople *People
}

var overAllFirstPeople *People

//创建第一个people
func CreateFirstPeople() *People{
	firstPeople := &People{1,nil,nil}
	overAllFirstPeople = firstPeople
	return firstPeople
}

//创建6个人,围成一个环
func CreateCircleDesk(totalNum int){
	var tmpPeople = overAllFirstPeople
	for i:=2 ;i<=totalNum ;i++  {
		 tmpPeople = AddPeople(i,tmpPeople)
	}
	//闭环
	tmpPeople.NextPeople = overAllFirstPeople
	overAllFirstPeople.PrePeople = tmpPeople
}

func AddPeople(num int,people *People) *People{
	newPeople := &People{num,people,nil}
	people.NextPeople = newPeople
	return newPeople
}

//1,2,3,4,5,6
func ShowPeoples(start *People){
	for true {
		fmt.Println(start)
		start = start.NextPeople
		time.Sleep(500*time.Millisecond)
	}
}

func GameStart(people *People,outNum int) {
	num := 1
	for true {
		time.Sleep(500*time.Millisecond)
		if people.NextPeople == people {
			fmt.Println("游戏结束,最后剩余",people.Num,"号")
			break
		}
		if num % outNum == 0 {
			num = 0
			fmt.Println(people.Num,"out","--下一个起始",people.NextPeople.Num)
			people.PrePeople.NextPeople = people.NextPeople
			people.NextPeople.PrePeople = people.PrePeople
		}
		people = people.NextPeople
		num++
	}
}
