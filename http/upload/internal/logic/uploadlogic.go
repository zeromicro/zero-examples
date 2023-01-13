package logic

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"upload/internal/svc"
	"upload/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 MB

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(r *http.Request) (resp *types.Response, err error) {
	_ = r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := os.Create(path.Join(l.svcCtx.Config.Path, handler.Filename))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)

	return &types.Response{
		OK: 0,
	}, nil
}
