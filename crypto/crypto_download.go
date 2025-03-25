package osscrypto

import (
	"fmt"

	"github.com/liuyongshuai/aliyun-oss-go-sdk"
)

// DownloadFile with multi part mode, temporarily not supported
func (bucket CryptoBucket) DownloadFile(objectKey, filePath string, partSize int64, options ...oss.Option) error {
	return fmt.Errorf("CryptoBucket doesn't support DownloadFile")
}
