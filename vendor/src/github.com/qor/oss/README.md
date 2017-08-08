# OSS

QOR OSS provide common interface to operate files in cloud storage, ftp, filesystem...

# Usage

```go
import (
	"github.com/oss/filesystem"
	"github.com/oss/s3"
	awss3 "github.com/aws/aws-sdk-go/s3"
)

func main() {
	storage := s3.New(s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: awss3.BucketCannedACLPublicRead})
	// storage := filesystem.New("/tmp")

	// Save a reader interface into storage
	storage.Put("/sample.txt", reader)

	// Get file with path
	storage.Get("/sample.txt")

	// Delete file with path
	storage.Delete("/sample.txt")

	// List all objects under path
	storage.List("/")
}
```
