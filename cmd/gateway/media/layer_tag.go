package media

import (
	"context"
	"github.com/minio/minio-go/v7/pkg/tags"
	minio "github.com/minio/minio/cmd"
)

func (l *Layer) PutObjectTags(ctx context.Context, s string, s2 string, s3 string, options minio.ObjectOptions) (minio.ObjectInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Layer) GetObjectTags(ctx context.Context, s string, s2 string, options minio.ObjectOptions) (*tags.Tags, error) {
	//TODO implement me
	panic("implement me")
}
func (l *Layer) DeleteObjectTags(ctx context.Context, s string, s2 string, options minio.ObjectOptions) (minio.ObjectInfo, error) {
	//TODO implement me
	panic("implement me")
}
