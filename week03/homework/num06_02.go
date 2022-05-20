package main

/*
6.创建一个student结构体，包含姓名和语数外三门课的成绩。用一个slice容纳一个班的同学，
求每位同学的平均分和整个班三门课的平均分，全班同学平均分低于60的有几位
*/
type Grade struct {
	Chinses,Math,English,avg_3 float32
}
type student struct {
	name string
	grade Grade

}
type Class struct {
	student []student
}

func arg_sore(grade Grade)float32{
	 grade.avg_3=(grade.Chinses+grade.Math+grade.English)/3
	 return grade.avg_3
}

func arg_class(class Class)(en_all,ch_all,ma_all float32){
	//var en_all,ch_all,ma_all float32
	for _,ele:=range class.student{
		en_all+=ele.grade.English
		ch_all+=ele.grade.Chinses
		ma_all+=ele.grade.Math
	}
	return en_all,ch_all,ma_all
}

func arg_stu(class Class)(count int){

	for _,ele:= range class.student{
		if ele.grade.avg_3 <60{
			count +=1
		}
	}
	return count
}
