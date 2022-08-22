package media

import (
	"context"
	minio "github.com/minio/minio/cmd"
	"net/http"
)

const (
	mappedPath    = "/mapped/"
	mappedPathLen = len(mappedPath)
)

// FileBucket 具体的媒体文件存放类型, 将一个简单的 fs 包装成bucket
type FileBucket struct {
	//*cmd.FSObjects
	//BucketReadable
	BucketWithMultipart
}

func (f FileBucket) ListObjectsV2(ctx context.Context, prefix, continuationToken, delimiter string, maxKeys int, fetchOwner bool, startAfter string) (result minio.ListObjectsV2Info, err error) {
	// 判断是否是读取mapped文件夹
	return
}

func (f FileBucket) Walk(ctx context.Context, prefix string, results chan<- minio.ObjectInfo, opts minio.ObjectOptions) error {
	//TODO implement me
	panic("implement me")
}

func (f FileBucket) GetObjectNInfo(ctx context.Context, object string, rs *minio.HTTPRangeSpec, h http.Header, lockType minio.LockType, opts minio.ObjectOptions) (reader *minio.GetObjectReader, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FileBucket) GetObjectInfo(ctx context.Context, object string, opts minio.ObjectOptions) (objInfo minio.ObjectInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (f FileBucket) isMapped(path string) string {
	if len(path) >= mappedPathLen {
		return path[mappedPathLen:]
	}
	return ""
}
