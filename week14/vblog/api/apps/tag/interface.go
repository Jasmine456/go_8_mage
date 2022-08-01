package tag

import "context"

type Service interface {
	//文章添加Tag
	AddTag(context.Context,*AddTagRequest)(*TagSet,error)
	//文章移除Tag
	RemoveTag(context.Context,*RemoveTagRequest)(*TagSet,error)
}


type AddTagRequest struct {
	//一次可以添加多个Tag
	Tags []*CreateTagRequest
}

type RemoveTagRequest struct {
	TagIds []int
}