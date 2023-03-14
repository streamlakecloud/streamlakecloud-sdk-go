# 服务端文件上传接口

## 流程

1. 初始化 VodClient，配置 AK 等参数
2. 打开本地图片文件或者视频、封面等文件，获取 file handler
3. 调用 VodClient 的对应接口上传图片或者视频

请注意提供的接口都是同步调用，有必要的话可以 wrapper 成异步接口。

## 代码示例

### 上传图片文件

```go
/// returns VOD MediaId and error.
/// MediaId maybe empty if error occurs.
func UploadImage(fullPath string) (string, error){
	client := NewVodClient(nil)
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: "PUT_ACCESS_KEY_HERE"}
    uploadToken, err1 :=client.UploadImageFileFromPath(fullPath)
    if err1 != nil {
		return "", err1
    }

	verifyReq := VerifyUploadTokenRequest{
		VodUploadToken: uploadToken,
	}
	verifyResp, err2 := client.VerifyUploadToken(verifyReq)
	if err2 != nil {
		return "", err2
	} else {
		mediaId := verifyResp.ResponseData.MediaId
		if len(mediaId) == 0 {
			return "", fmt.Errorf("no MediaId")
		}
        return mediaId, nil
	}
```

### 上传视频与封面（可选）
```go
/// returns VOD MediaId and error.
/// MediaId maybe empty if error occurs.
func UploadVideoWithCover(videoPath string, coverPath string) (string, error){
	client := NewVodClient(nil)
	client.ServiceInfo.Credentials = base.Credentials{AccessKey: "PUT_ACCESS_KEY_HERE"}
    uploadToken, err1 := client.UploadVideoFileWithOptionalCoverFromPath(videoPath, coverPath)
    if err1 != nil {
		return "", err1
    }

	verifyReq := VerifyUploadTokenRequest{
		VodUploadToken: uploadToken,
	}
	verifyResp, err2 := client.VerifyUploadToken(verifyReq)
	if err2 != nil {
		return "", err2
	} else {
		mediaId := verifyResp.ResponseData.MediaId
		if len(mediaId) == 0 {
			return "", fmt.Errorf("no MediaId")
		}
        return mediaId, nil
	}
```

同时 SDK 提供了直接操作 `*os.File` 和提供文件 basename 的版本供特定场景使用，请参阅 `UploadImageFile` 和 `UploadVideoFileWithOptionalCover`。