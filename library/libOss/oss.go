package libOss

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

var OssClient CosClient

type CosClient struct {
	C *cos.Client
}

func InitOssClient(ctx context.Context) {
	u, err := url.Parse(g.Cfg().MustGet(ctx, "oss.bucket_url").String())
	if err != nil {
		g.Log().Error(ctx, err)
		panic(err)
	}
	//su, err := url.Parse(g.Cfg().MustGet(ctx, "oss.service_url").String())
	//if err != nil {
	//	g.Log().Error(ctx, err)
	//	panic(err)
	//}
	// 用于 Get Service 查询，默认全地域 service.cos.myqcloud.com
	b := &cos.BaseURL{BucketURL: u} //ServiceURL: su,

	// 1.永久密钥
	OssClient.C = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  g.Cfg().MustGet(ctx, "oss.secret_id").String(),
			SecretKey: g.Cfg().MustGet(ctx, "oss.secret_key").String(), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
}

func (c *CosClient) UploadFile(ctx context.Context, key string, fd multipart.File) error {
	_, err := OssClient.C.Object.Put(context.Background(), key, fd, nil)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (c *CosClient) GetFile(ctx context.Context, key string) ([]byte, error) {
	resp, err := OssClient.C.Object.Get(ctx, key, nil)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return all, nil
}
