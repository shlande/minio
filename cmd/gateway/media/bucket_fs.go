package media

// FileBucket 具体的媒体文件存放类型, 将一个简单的 fs 包装成bucket
type FileBucket struct {
	BucketReadable
	BucketWithMultipart
}
