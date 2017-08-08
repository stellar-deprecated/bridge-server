package s3

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/qor/oss"
)

// Client S3 storage
type Client struct {
	*s3.S3
	Config *Config
}

// Config S3 client config
type Config struct {
	AccessID     string
	AccessKey    string
	Region       string
	Bucket       string
	SessionToken string
	ACL          string
	Endpoint     string
}

func EC2RoleAwsConfig(config *Config) *aws.Config {
	ec2m := ec2metadata.New(session.New(), &aws.Config{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Endpoint:   aws.String("http://169.254.169.254/latest"),
	})

	cr := credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{
		Client: ec2m,
	})

	return &aws.Config{
		Region:      aws.String(config.Region),
		Credentials: cr,
	}
}

// New initialize S3 storage
func New(config *Config) *Client {
	if config.ACL == "" {
		config.ACL = s3.BucketCannedACLPublicRead
	}

	client := &Client{Config: config}

	if config.AccessID == "" && config.AccessKey == "" {
		client.S3 = s3.New(session.New(), EC2RoleAwsConfig(config))
	} else {
		creds := credentials.NewStaticCredentials(config.AccessID, config.AccessKey, config.SessionToken)
		if _, err := creds.Get(); err == nil {
			client.S3 = s3.New(session.New(), &aws.Config{
				Region:      &config.Region,
				Credentials: creds,
			})
		}
	}

	return client
}

// Get receive file with given path
func (client Client) Get(path string) (file *os.File, err error) {
	getResponse, err := client.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(client.ToRelativePath(path)),
	})

	if err == nil {
		if file, err = ioutil.TempFile("/tmp", "s3"); err == nil {
			_, err = io.Copy(file, getResponse.Body)
			file.Seek(0, 0)
		}
	}

	return file, err
}

// Put store a reader into given path
func (client Client) Put(urlPath string, reader io.Reader) (*oss.Object, error) {
	if seeker, ok := reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}

	urlPath = client.ToRelativePath(urlPath)
	buffer, err := ioutil.ReadAll(reader)

	fileType := mime.TypeByExtension(path.Ext(urlPath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	params := &s3.PutObjectInput{
		Bucket:        aws.String(client.Config.Bucket), // required
		Key:           aws.String(urlPath),              // required
		ACL:           aws.String(client.Config.ACL),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(int64(len(buffer))),
		ContentType:   aws.String(fileType),
	}

	_, err = client.S3.PutObject(params)

	now := time.Now()
	return &oss.Object{
		Path:             urlPath,
		Name:             filepath.Base(urlPath),
		LastModified:     &now,
		StorageInterface: client,
	}, err
}

// Delete delete file
func (client Client) Delete(path string) error {
	_, err := client.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(client.ToRelativePath(path)),
	})
	return err
}

// List list all objects under current path
func (client Client) List(path string) ([]*oss.Object, error) {
	var objects []*oss.Object
	var prefix string

	if path != "" {
		prefix = strings.Trim(path, "/") + "/"
	}

	listObjectsResponse, err := client.S3.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(client.Config.Bucket),
		Prefix: aws.String(prefix),
	})

	if err == nil {
		for _, content := range listObjectsResponse.Contents {
			objects = append(objects, &oss.Object{
				Path:             client.ToRelativePath(*content.Key),
				Name:             filepath.Base(*content.Key),
				LastModified:     content.LastModified,
				StorageInterface: client,
			})
		}
	}

	return objects, err
}

// GetEndpoint get endpoint, FileSystem's endpoint is /
func (client Client) GetEndpoint() string {
	if client.Config.Endpoint != "" {
		return client.Config.Endpoint
	}

	endpoint := client.S3.Endpoint
	for _, prefix := range []string{"https://", "http://"} {
		endpoint = strings.TrimPrefix(endpoint, prefix)
	}

	return client.Config.Bucket + "." + endpoint
}

var urlRegexp = regexp.MustCompile(`(https?:)?//((\w+).)+(\w+)/`)

func (client Client) ToRelativePath(urlPath string) string {
	if urlRegexp.MatchString(urlPath) {
		if u, err := url.Parse(urlPath); err == nil {
			return u.Path
		}
	}

	return "/" + strings.TrimPrefix(urlPath, "/")
}
