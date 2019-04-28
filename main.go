package main

import "joseph/CircleLink"
//2019.4.28
func main() {
	first := CircleLink.CreateFirstPeople()
	CircleLink.CreateCircleDesk(10)
	//CircleLink.ShowPeoples(first)
	CircleLink.GameStart(first,5)
}

//1,2,3

//5,4,6,2,1,3
