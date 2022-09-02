package mock

import (
	"context"
	"go_8_mage/week14/project_arch/ddd/tag"
)

type M struct {

}

func (m *M) Query(ctx context.Context,req *tag.QueryRequest)(*tag.TagSet,error){
	return &tag.TagSet{
		Items:[]*tag.Tag{
			{Key:"key1",Value: "value1"},
		},
	},nil
}
