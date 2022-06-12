package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

var myBucket = "simpledy"
var myAccessKeyId = "LTAI5t9Na736WxVs5yy4KDu8"
var myAccessKeySecret = "tR4PjMXmAX15aLrKy0KYff3y81sElU"

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func UploadFile(localFilePath string, objectPath string) {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := myAccessKeyId
	accessKeySecret := myAccessKeySecret
	bucketName := myBucket
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName := objectPath //oss://simpledy/
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := localFilePath
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}
