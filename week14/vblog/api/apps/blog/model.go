package blog


type Blog struct {
	//文章ID
	Id int
	//文章摘要信息.通过提取content内容获取
	Sumary string
	// 创建时间
	CreateAt int64
	//更新时间
	UpdateAt int64
	//发布时间
	PublishAt int64
	//状态 草稿/发布
	Status Status
	//用户提交数据
	*CreateBlogRequest
}
//列表数据
type BlogSet struct {
	Items []*Blog
}

func NewCreateBlogRequest() *CreateBlogRequest{
	return &CreateBlogRequest{}
}

type CreateBlogRequest struct {
	//	文章图片
	TitleImg string
	//文章标题
	TitleName string
	//文章副标题
	SubTitle string
	//文章内容
	Content string
	//文章作者
	Author string

}

func NewPutUpdateBlogRequest() *UpdateBlogRequest{
	return &UpdateBlogRequest{
		UpdateMode: UPDATE_MODE_PUT,
	}
}

func NewPatchUpdateBlogRequest() *UpdateBlogRequest{
	return &UpdateBlogRequest{
		UpdateMode: UPDATE_MODE_PATCH,
	}
}

type UpdateBlogRequest struct {
	UpdateMode UpdateMode
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	Id int
}

func NewQueryBlogRequest() *QueryBlogRequest{
	return &QueryBlogRequest{
		PageSize: 20,
		PageNumber: 1,
	}
}

type QueryBlogRequest struct {
	PageSize int
	PageNumber int
	Keywords string
}


type DescribeBlogRequest struct {
	Id int
}

type UpdateBlogStatusRequest struct {
	//文章id
	Id int
	//状态 草稿/发布
	Status Status

}
