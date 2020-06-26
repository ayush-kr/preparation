package main

import "log"

type single struct {
	singleInt int
}

var singleIns *single

func (s *single) Increment(increase int) {
	s.singleInt += increase
}

func (s *single) Decrement(decrease int) {
	s.singleInt -= decrease
}

func (s *single) Value() int {
	return s.singleInt
}

//GetInstance return single instance
func GetInstance() *single {
	if singleIns == nil {
		singleIns = &single{
			singleInt: 0,
		}
	}
	return singleIns
}

func main() {
	s := GetInstance()
	log.Println(s.Value())
	s.Increment(4)
	log.Println(s.Value())
	s.Decrement(3)
	log.Println(s.Value())
}
