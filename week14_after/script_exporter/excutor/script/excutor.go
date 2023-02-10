package script

import "io"

func NewExcutor(modules string) *Excutor {
	return &Excutor{
		ModuleDir: modules,
	}
}

//负责执行脚本(模块）
//执行： 脚本的名称（模块名称），脚本的参数
//Script Excutor 需要在一个目录厦门 搜索脚本，搜索的目录需要提前定义
type Excutor struct {
	ModuleDir string
}

//moduleName 模块名称
//moduleParams 参数 Json
//继续的结果，IO 数据流
func (e *Excutor) Exec(moduleName string, moduleParams string, w io.WriteCloser) error {
	//	1.寻找模块


	//	2.执行模块，并实时返回执行结果

	return nil
}
