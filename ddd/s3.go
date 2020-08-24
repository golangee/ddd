package ddd

type BucketSpec struct {
}

func S3(buckets ...*BucketSpec) *LayerSpec {
	return nil
}

// Bucket creates a new Repository with the given type name.
func Bucket(name TypeName) *BucketSpec {
	return &BucketSpec{}
}

type FolderSpec struct {
}

func Folder(name TypeName) *FolderSpec {
	return &FolderSpec{}
}

func Filesystem(folders ...*FolderSpec) *LayerSpec {
	return nil
}
