package media

import (
	"context"
	minio "github.com/minio/minio/cmd"
)

// TransitionObject TODO: 可以用这接口来控制File和Dash的转换？
func (l *Layer) TransitionObject(ctx context.Context, bucket, object string, opts minio.ObjectOptions) error {
	return minio.NotImplemented{}
}

func (l *Layer) RestoreTransitionedObject(ctx context.Context, bucket, object string, opts minio.ObjectOptions) error {
	return minio.NotImplemented{}
}
