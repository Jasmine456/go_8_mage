package blog

type Status int

const (
	//草稿
	STATUS_DRAFT Status = 0
	//	已发布
	STATUS_PUBLISH Status = 1
)

//更新模式
type UpdateMode string

const(
	//全量更新
	UPDATE_MODE_PUT UpdateMode = "put"
	//部分更新
	UPDATE_MODE_PATCH UpdateMode = "patch"
)

