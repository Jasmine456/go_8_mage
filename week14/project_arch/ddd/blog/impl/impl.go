package impl

import (
	"context"
	"go_8_mage/week14/project_arch/ddd/tag"
	"go_8_mage/week14/project_arch/ddd/blog"
)

func NewImpl(tag tag.Service) *Impl {
	return &Impl{
		tag: tag,
	}
}

type Impl struct {
	tag tag.Service
}

func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	i.tag.Query(ctx, tag.NewQueryRequest(1))
	return nil, nil
}
