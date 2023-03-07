package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/choisangh/image_crud_api/internal/global"
	"github.com/choisangh/image_crud_api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type RequestCreate struct {
	FileNO      string `json:"no" binding:"required"`
	ImageBase64 string `json:"image" binding:"required"`
}
type Requestput struct {
	ImageBase64 string `json:"image" binding:"required"`
}

type RequestURI struct {
	ImageID string `uri:"id" binding:"required"`
}

type Response struct {
	Res string `json:"res"`
}

const (
	FILE_NOT_VALID_MSG        = "file is not valid"
	FILE_ALREADY_EXIST_MSG    = "file is already exist"
	FILE_NOT_EXIST_MSG        = "file is not exist"
	FILE_FORMAT_NOT_VALID_MSG = "file is not image format"
	FILE_CREATING_ERR_MSG     = "file creating error"
	FILE_DELETE_ERR_MSG       = "failed to delete file"
	SUCCESS_MSG               = "success"
)

func CreateImage(c *gin.Context) {
	// 바인딩
	req := RequestCreate{}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_VALID_MSG})
		return
	}

	// 파일명이 중복되는지 확인
	filename := req.FileNO
	if utils.IsFileExists(filename, global.IMAGE_FILEPATH) {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_ALREADY_EXIST_MSG})
		return
	}

	// 이미지 파일인지 확인
	if !utils.IsValidImageFormat(req.ImageBase64) {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_FORMAT_NOT_VALID_MSG})
		return
	}

	// 이미지 파일 생성
	if err := utils.CreateImageFile(filename, req.ImageBase64, global.IMAGE_FILEPATH); err != nil {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_CREATING_ERR_MSG})
		return
	}

	// 성공 응답
	c.JSON(http.StatusOK, Response{Res: SUCCESS_MSG})
}

func ReadImage(c *gin.Context) {
	// 바인딩
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_VALID_MSG})
		return
	}

	// 읽을 파일이 존재하는지 확인
	filename := reqURI.ImageID
	if !utils.IsFileExists(filename, global.IMAGE_FILEPATH) {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_EXIST_MSG})
		return
	}

	// 파일 내보내기
	c.File(filepath.Join(global.IMAGE_FILEPATH, filename))
}

func PutImage(c *gin.Context) {
	// 바인딩
	req := Requestput{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_VALID_MSG})

		return
	}

	// 수정할 파일 존재하는지 확인
	filename := c.Param("id")
	if !utils.IsFileExists(filename, global.IMAGE_FILEPATH) {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_EXIST_MSG})
		return
	}

	// 이미지 파일인지 확인
	if !utils.IsValidImageFormat(req.ImageBase64) {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_FORMAT_NOT_VALID_MSG})
		return
	}

	// 이미지 파일 생성
	if err := utils.CreateImageFile(filename, req.ImageBase64, global.IMAGE_FILEPATH); err != nil {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_CREATING_ERR_MSG})
		return
	}

	// 성공 응답
	c.JSON(http.StatusOK, Response{Res: SUCCESS_MSG})

}

func DeleteImage(c *gin.Context) {
	// 바인딩
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_VALID_MSG})
		return
	}

	// 삭제할 파일이 존재하는지 확인
	filename := reqURI.ImageID
	if !utils.IsFileExists(filename, global.IMAGE_FILEPATH) {
		c.JSON(http.StatusBadRequest, Response{Res: FILE_NOT_EXIST_MSG})
		return
	}

	// 파일 삭제
	if err := os.Remove(filepath.Join(global.IMAGE_FILEPATH, filename)); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Res: FILE_DELETE_ERR_MSG})
		return
	}

	// 성공 응답
	c.JSON(http.StatusOK, Response{Res: SUCCESS_MSG})

}
