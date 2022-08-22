package media

import (
	"context"
	minio "github.com/minio/minio/cmd"
)

func (l *Layer) getMultipartBucket(bucket string) BucketWithMultipart {
	if bucket == "file" {
		return l.fb
	}
	return nil
}

func (l *Layer) ListMultipartUploads(ctx context.Context, bucket, prefix, keyMarker, uploadIDMarker, delimiter string, maxUploads int) (result minio.ListMultipartsInfo, err error) {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return minio.ListMultipartsInfo{}, minio.NotImplemented{}
	}
	return bkt.ListMultipartUploads(ctx, prefix, keyMarker, uploadIDMarker, delimiter, maxUploads)
}

func (l *Layer) NewMultipartUpload(ctx context.Context, bucket, object string, opts minio.ObjectOptions) (uploadID string, err error) {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return "", minio.NotImplemented{}
	}
	return bkt.NewMultipartUpload(ctx, object, opts)
}

func (l *Layer) CopyObjectPart(ctx context.Context, srcBucket, srcObject, destBucket, destObject string, uploadID string, partID int, startOffset int64, length int64, srcInfo minio.ObjectInfo, srcOpts, dstOpts minio.ObjectOptions) (info minio.PartInfo, err error) {
	return minio.PartInfo{}, minio.NotImplemented{}
}

func (l *Layer) PutObjectPart(ctx context.Context, bucket, object, uploadID string, partID int, data *minio.PutObjReader, opts minio.ObjectOptions) (info minio.PartInfo, err error) {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return minio.PartInfo{}, minio.NotImplemented{}
	}
	return bkt.PutObjectPart(ctx, object, uploadID, partID, data, opts)
}

func (l *Layer) GetMultipartInfo(ctx context.Context, bucket, object, uploadID string, opts minio.ObjectOptions) (info minio.MultipartInfo, err error) {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return minio.MultipartInfo{}, minio.NotImplemented{}
	}
	return bkt.GetMultipartInfo(ctx, object, uploadID, opts)
}

func (l *Layer) ListObjectParts(ctx context.Context, bucket, object, uploadID string, partNumberMarker int, maxParts int, opts minio.ObjectOptions) (result minio.ListPartsInfo, err error) {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return minio.ListPartsInfo{}, minio.NotImplemented{}
	}
	return bkt.ListObjectParts(ctx, object, uploadID, partNumberMarker, maxParts, opts)
}

func (l *Layer) AbortMultipartUpload(ctx context.Context, bucket, object, uploadID string, opts minio.ObjectOptions) error {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return minio.NotImplemented{}
	}
	return bkt.AbortMultipartUpload(ctx, object, uploadID, opts)
}

func (l *Layer) CompleteMultipartUpload(ctx context.Context, bucket, object, uploadID string, uploadedParts []minio.CompletePart, opts minio.ObjectOptions) (objInfo minio.ObjectInfo, err error) {
	var bkt BucketWithMultipart
	if bkt = l.getMultipartBucket(bucket); bkt == nil {
		return minio.ObjectInfo{}, minio.NotImplemented{}
	}
	return bkt.CompleteMultipartUpload(ctx, object, uploadID, uploadedParts, opts)
}
