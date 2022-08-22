package media

import (
	"context"
	"errors"
	"github.com/minio/madmin-go"
	minio "github.com/minio/minio/cmd"
	"github.com/minio/pkg/bucket/policy"
	"net/http"
)

type Layer struct {
	minio.GatewayUnsupported
	// FsObjectLayer store path
	basePath string
	// FsObjectLayer
	fs minio.FSObjects
	fb *FileBucket
}

func (l *Layer) getBucket(bucket string) BucketBase {
	switch bucket {
	case "file":
		return l.fb
	default:
		return nil
	}
}

func (l *Layer) listBucket() []BucketBase {
	var res = make([]BucketBase, 0, 3)
	res = append(res, l.fb)
	return res
}

func (l *Layer) Shutdown(ctx context.Context) error {
	_ = l.fs.Shutdown(ctx)
	return l.fb.Shutdown(ctx)
}

func (l *Layer) BackendInfo() madmin.BackendInfo {
	info := l.fs.BackendInfo()
	info.Type = madmin.Gateway
	return info
}

func (l *Layer) StorageInfo(ctx context.Context) (minio.StorageInfo, []error) {
	info, err := l.fs.StorageInfo(ctx)
	if err != nil {
		return info, err
	}
	info.Backend = l.BackendInfo()
	return info, nil
}

func (l *Layer) LocalStorageInfo(ctx context.Context) (minio.StorageInfo, []error) {
	return l.fs.LocalStorageInfo(ctx)
}

func (l *Layer) GetBucketInfo(ctx context.Context, bucket string, opts minio.BucketOptions) (bucketInfo minio.BucketInfo, err error) {
	var bkt BucketBase
	if bkt = l.getBucket(bucket); bkt == nil {
		return minio.BucketInfo{}, errors.New("bucket not found")
	}
	return bkt.BucketInfo(), nil
}

func (l *Layer) ListBuckets(ctx context.Context, opts minio.BucketOptions) (buckets []minio.BucketInfo, err error) {
	// ignore deleted bucket
	if opts.Deleted {
		return nil, nil
	}
	for _, v := range l.listBucket() {
		buckets = append(buckets, v.BucketInfo())
	}
	return
}

// DeleteBucket Delete is not allowed
func (l *Layer) DeleteBucket(ctx context.Context, bucket string, opts minio.DeleteBucketOptions) error {
	return minio.NotImplemented{}
}

// ListObjects ListObjectsV1 is not supported
func (l *Layer) ListObjects(ctx context.Context, bucket, prefix, marker, delimiter string, maxKeys int) (result minio.ListObjectsInfo, err error) {
	return minio.ListObjectsInfo{}, minio.NotImplemented{}
}

func (l *Layer) ListObjectsV2(ctx context.Context, bucket, prefix, continuationToken, delimiter string, maxKeys int, fetchOwner bool, startAfter string) (result minio.ListObjectsV2Info, err error) {
	if startAfter != "" {
		return minio.ListObjectsV2Info{}, errors.New("param 'startAfter' is not supported")
	}
	var (
		bkt BucketReadable
		ok  bool
	)
	if bkt, ok = l.getBucket(bucket).(BucketReadable); !ok {
		return minio.ListObjectsV2Info{}, errors.New("bucket not found or unreadable")
	}
	return bkt.ListObjectsV2(ctx, prefix, continuationToken, delimiter, maxKeys, fetchOwner)
}

func (l *Layer) Walk(ctx context.Context, bucket, prefix string, results chan<- minio.ObjectInfo, opts minio.ObjectOptions) error {
	if f, ok := l.getBucket(bucket).(FileBucket); ok {
		return f.Walk(ctx, prefix, results, opts)
	}
	if f, ok := l.getBucket(bucket).(DashLayer); ok {
		return f.Walk(ctx, prefix, results, opts)
	}
	return minio.NotImplemented{}
}

func (l *Layer) GetObjectNInfo(ctx context.Context, bucket, object string, rs *minio.HTTPRangeSpec, h http.Header, lockType minio.LockType, opts minio.ObjectOptions) (reader *minio.GetObjectReader, err error) {
	//TODO implement me
	panic("implement me")
}

func (l *Layer) GetObjectInfo(ctx context.Context, bucket, object string, opts minio.ObjectOptions) (objInfo minio.ObjectInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (l *Layer) PutObject(ctx context.Context, bucket, object string, data *minio.PutObjReader, opts minio.ObjectOptions) (objInfo minio.ObjectInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (l *Layer) DeleteObject(ctx context.Context, bucket, object string, opts minio.ObjectOptions) (minio.ObjectInfo, error) {
	var (
		bkt BucketWriteable
		ok  bool
	)
	if bkt, ok = l.getBucket(bucket).(BucketWriteable); !ok {
		return minio.ObjectInfo{}, minio.NotImplemented{}
	}
	return bkt.DeleteObject(ctx, object, opts)
}

func (l *Layer) DeleteObjects(ctx context.Context, bucket string, objects []minio.ObjectToDelete, opts minio.ObjectOptions) ([]minio.DeletedObject, []error) {
	result := make([]minio.DeletedObject, len(objects))
	errors := make([]error, len(objects), len(objects))
	for i, _ := range result {
		errors[i] = minio.NotImplemented{}
	}
	return result, errors
}

func (l *Layer) GetBucketPolicy(ctx context.Context, s string) (*policy.Policy, error) {
	//TODO implement me
	panic("implement me")
}

func (l *Layer) DeleteBucketPolicy(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

// IsCompressionSupported FIXME: can we support compression ?
func (l *Layer) IsCompressionSupported() bool {
	return false
}
