package blog

type Blog struct {
	Id int
	*CreateBlogRequest
}

type CreateBlogRequest struct {
	Title string
	Content string
}