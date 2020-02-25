package mysort

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Println("---------welcome mysort init----------")
}

type
(
	ElectedNode struct {
		Rate float64 //单位时间上台率=本节课上台次数/本节课在教室内的总时长
		Uid  int64
	}
)

//type ElectedNode struct {
//	name     string
//	mass     float64
//	distance float64
//}

type By func(p1, p2 *ElectedNode) bool

func (by By) Sort(electedNodes []ElectedNode) {
	ps := &electedNodeSorter{
		electedNodes: electedNodes,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

type electedNodeSorter struct {
	electedNodes []ElectedNode
	by      func(p1, p2 *ElectedNode) bool // Closure used in the Less method.
}

func (s *electedNodeSorter) Len() int {
	return len(s.electedNodes)
}

// Swap is part of sort.Interface.
func (s *electedNodeSorter) Swap(i, j int) {
	s.electedNodes[i], s.electedNodes[j] = s.electedNodes[j], s.electedNodes[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *electedNodeSorter) Less(i, j int) bool {
	return s.by(&s.electedNodes[i], &s.electedNodes[j])
}


func MySort() {
	var electedNodes = []ElectedNode{
		{Rate:1,Uid:1},
		{Rate:2,Uid:2},
		{Rate:3,Uid:3},
		{Rate:4,Uid:4},
		{Rate:5,Uid:5},
	}
	fmt.Println("--------------sort before -------------")
	for index,node := range electedNodes{
		fmt.Printf("index:%d node:%+v\n",index,node)
	}


	//比较函数
	rate := func (p1, p2 *ElectedNode) bool {
		return p1.Rate < p2.Rate
	}
	fmt.Println("--------------sort before -------------")
	By(rate).Sort(electedNodes)
	for index,node := range electedNodes{
		fmt.Printf("index:%d node:%+v\n",index,node)
	}
	fmt.Println("--------------sort end--- -------------")
}






