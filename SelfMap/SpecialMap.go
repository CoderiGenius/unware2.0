package SelfMap

var SpecialMap map[string]string
func InitMap()  {
	SpecialMap = make(map[string]string)
	SpecialMap["刘国柱"] = "院长"
	SpecialMap["周艳平"] = "老师"
	SpecialMap["刘妍君"] = "老师"
	SpecialMap["周思成"] = "会长"

}

func FindHashMap(name string)  string{
	//fmt.Println(11111,SpecialMap["周思成"])
	return SpecialMap[name]
	
}