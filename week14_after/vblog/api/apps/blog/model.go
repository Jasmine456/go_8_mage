package blog

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go_8_mage/week14/vblog/api/apps/tag"
	"net/http"
	"strconv"
	"time"
)

var (
	validate = validator.New()
)

func NewCreateBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreateAt:          time.Now().Unix(),
		CreateBlogRequest: req,
		Status:            STATUS_DRAFT,
		Tags:              []*tag.Tag{},
	}
}

type Blog struct {
	//文章ID
	Id int `json:"id"`
	//文章摘要信息.通过提取content内容获取
	Sumary string `json:"summary" gorm:"-"`
	// 创建时间
	CreateAt int64 `json:"create_at"`
	//更新时间
	UpdateAt int64 `json:"update_at"`
	//发布时间
	PublishAt int64 `json:"publish_at"`
	//状态 草稿/发布
	Status Status `json:"status"`
	//用户提交数据
	*CreateBlogRequest
	//	博客标签
	Tags []*tag.Tag `json:"tags"`
}

func (b *Blog) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}
func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

//列表数据
type BlogSet struct {
	//总条目个数,用于前端分页
	Total int64
	//列表数据
	Items []*Blog
}

func (b *BlogSet) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{}
}

type CreateBlogRequest struct {
	//	文章图片
	TitleImg string `json:"title_img"`
	//文章标题
	TitleName string `json:"title_name" validate:"required"`
	//文章副标题
	SubTitle string `json:"sub_title"`
	//文章内容
	Content string `json:"content" validate:"required"`
	//文章作者
	Author string `json:"author"`
}

func (req *CreateBlogRequest) Validate() error {
	return validate.Struct(req)
}

func NewPutUpdateBlogRequest(id int) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

func NewPatchUpdateBlogRequest(id int) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PATCH,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	BlogId     int
	UpdateMode UpdateMode
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	Id int
}

func NewDeleteBlogRequest(id int) *DeleteBlogRequest {
	return &DeleteBlogRequest{Id: id}
}
func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

//使用http标准库的原始方法来获取
// http query string: ?keywords=b&page_size=20&page_number=1
func NewQueryBlogRequestFromHTTP(r *http.Request) *QueryBlogRequest {
	qs := r.URL.Query()

	req := NewQueryBlogRequest()
	req.Keywords = qs.Get("keywords")

	psStr := qs.Get("page_size")
	if psStr != "" {
		req.PageSize, _ = strconv.Atoi(psStr)
	}

	pnStr := qs.Get("page_number")
	if pnStr != "" {
		req.PageNumber, _ = strconv.Atoi(pnStr)
	}
	return req
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
}

func (req *QueryBlogRequest) Offset() int {
	return (req.PageNumber - 1) * req.PageSize
}

func NewDescribeBlogRequest(id int) *DescribeBlogRequest {
	return &DescribeBlogRequest{Id: id}
}

type DescribeBlogRequest struct {
	Id int
}

func NewUpdateBlogStatusRequest(id int, status Status) *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{
		Id:     id,
		Status: status,
	}
}

func NewDefaultUpdateBlogStatusRequest() *UpdateBlogStatusRequest {
	return &UpdateBlogStatusRequest{}
}

type UpdateBlogStatusRequest struct {
	//文章id
	Id int `json:"id"`
	//文章状态 草稿/发布
	Status Status `json:"status"`
}
