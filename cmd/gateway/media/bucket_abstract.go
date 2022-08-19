package media

// BucketAbstract 是一个模版类型，用来将读写操作映射为函数操作，达到实现 unix 中一切皆文件的设计
// 后续可以通过这个实现来暴露一些可写文件给用户来暴露操作接口
type BucketAbstract struct {
	BucketReadable
	BucketWriteable
}
