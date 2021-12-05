package mySort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestHero(t *testing.T)  {

	//var hero = Hero{"张三", 18}
	//var heroes HeroSlice = HeroSlice{ hero }
	//fmt.Println(heroes)

	var heroes HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d",rand.Int()),
			Age: rand.Intn(100),
		}

		heroes = append(heroes, hero)
	}

	//输出排序前
	fmt.Println("-------------------输出排序前-------------------")
	for _, v := range heroes {
		fmt.Println(v)
	}

	fmt.Println("-------------------输出排序后-------------------")
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}

}