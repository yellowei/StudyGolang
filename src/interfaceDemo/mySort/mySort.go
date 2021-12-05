package mySort

// 定义Hero结构体
type Hero struct {
	Name string
	Age int
}

// 定义Hero切片
type HeroSlice []Hero

// 实现sort里的Interface接口
func (heroSlice HeroSlice) Len() int  {
	return len(heroSlice)
}

func (heroSlice HeroSlice) Less(i, j int) bool {
	return heroSlice[i].Age < heroSlice[j].Age
}

func (heroSlice HeroSlice) Swap(i, j int)  {
	//temp := heroSlice[i]
	//heroSlice[i] = heroSlice[j]
	//heroSlice[j] = temp
	heroSlice[i], heroSlice[j] = heroSlice[j], heroSlice[i]
}
