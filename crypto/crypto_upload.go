package osscrypto

import (
	"fmt"

	"github.com/liuyongshuai/aliyun-oss-go-sdk"
)

// UploadFile with multi part mode
func (bucket CryptoBucket) UploadFile(objectKey, filePath string, partSize int64, options ...oss.Option) error {
	return fmt.Errorf("CryptoBucket doesn't support UploadFile")
}
