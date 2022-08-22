package media

import (
	"context"
	"github.com/minio/minio/cmd"
	"github.com/minio/pkg/bucket/policy"
	"net/http"
)

type BucketBase interface {
	Shutdown(ctx context.Context) error
	BucketInfo() cmd.BucketInfo
	// GetPolicy 获取bucket的权限
	GetPolicy(context.Context, string) (*policy.Policy, error)
}

type BucketReadable interface {
	BucketBase
	// ListObjectsV2 startAfter is supported comparing with s3
	ListObjectsV2(ctx context.Context, prefix, continuationToken, delimiter string, maxKeys int, fetchOwner bool) (result cmd.ListObjectsV2Info, err error)
	// Walk lists all objects including versions, delete markers.
	Walk(ctx context.Context, prefix string, results chan<- cmd.ObjectInfo, opts cmd.ObjectOptions) error

	// Object operations.

	// GetObjectNInfo returns a GetObjectReader that satisfies the
	// ReadCloser interface. The Close method unlocks the object
	// after reading, so it must always be called after usage.
	//
	// IMPORTANTLY, when implementations return err != nil, this
	// function MUST NOT return a non-nil ReadCloser.
	GetObjectNInfo(ctx context.Context, object string, rs *cmd.HTTPRangeSpec, h http.Header, lockType cmd.LockType, opts cmd.ObjectOptions) (reader *cmd.GetObjectReader, err error)
	GetObjectInfo(ctx context.Context, object string, opts cmd.ObjectOptions) (objInfo cmd.ObjectInfo, err error)
}

// BucketWriteable 允许写操作的 layer
type BucketWriteable interface {
	BucketBase
	PutObject(ctx context.Context, object string, data *cmd.PutObjReader, opts cmd.ObjectOptions) (objInfo cmd.ObjectInfo, err error)
	DeleteObject(ctx context.Context, object string, opts cmd.ObjectOptions) (cmd.ObjectInfo, error)
	DeleteObjects(ctx context.Context, objects []cmd.ObjectToDelete, opts cmd.ObjectOptions) ([]cmd.DeletedObject, []error)
}

// BucketWithMultipart 支持分片上传的 bucket，是写操作的提升
//	更多细节可以查看 cmd.ObjectLayer
type BucketWithMultipart interface {
	BucketWriteable
	ListMultipartUploads(ctx context.Context, prefix, keyMarker, uploadIDMarker, delimiter string, maxUploads int) (result cmd.ListMultipartsInfo, err error)
	NewMultipartUpload(ctx context.Context, object string, opts cmd.ObjectOptions) (uploadID string, err error)
	PutObjectPart(ctx context.Context, object, uploadID string, partID int, data *cmd.PutObjReader, opts cmd.ObjectOptions) (info cmd.PartInfo, err error)
	GetMultipartInfo(ctx context.Context, object, uploadID string, opts cmd.ObjectOptions) (info cmd.MultipartInfo, err error)
	ListObjectParts(ctx context.Context, object, uploadID string, partNumberMarker int, maxParts int, opts cmd.ObjectOptions) (result cmd.ListPartsInfo, err error)
	AbortMultipartUpload(ctx context.Context, object, uploadID string, opts cmd.ObjectOptions) error
	CompleteMultipartUpload(ctx context.Context, object, uploadID string, uploadedParts []cmd.CompletePart, opts cmd.ObjectOptions) (objInfo cmd.ObjectInfo, err error)
}
