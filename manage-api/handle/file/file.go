package file

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/zxbit2011/ant/manage-api/global"
	"github.com/zxbit2011/ant/manage-api/handle"
	"github.com/zxbit2011/ant/model"
	"github.com/zxbit2011/ant/utils"
	"github.com/zxbit2011/ant/utils/convert"
	"github.com/zxbit2011/ant/utils/filetype"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

var b float64 = 1024

// 单个文件上传
func UploadFile(c echo.Context) error {
	auth := c.Param("auth")
	loginInfo := global.GetLoginInfo(c)
	//放置目录
	dir := c.FormValue("dir")
	if dir == "" {
		dir = "default"
	}
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		return utils.ErrorNull(c, "获取文件失败")
	}
	if file.Size <= 0 {
		return utils.ErrorNull(c, "空文件")
	}
	if convert.MustInt64(file.Size) > global.Conf.FileUpload.MaxFileSize {
		return utils.ErrorNull(c, fmt.Sprintf("文件过大超出限制%vmb", fmt.Sprintf("%.2f", convert.MustFloat64(global.Conf.FileUpload.MaxFileSize)/b/b)))
	}
	src, err := file.Open()
	if err != nil {
		return utils.ErrorNull(c, "打开文件失败")
	}
	defer src.Close()

	buf, err := ioutil.ReadAll(src)
	if err != nil {
		return utils.ErrorNull(c, "获取文件格式错误")
	}
	ext := path.Ext(file.Filename)

	head := make([]byte, 261)
	_, _ = src.Read(head)
	//格式限制判断
	isExt := strings.Contains(fmt.Sprintf(",%v,", global.Conf.FileUpload.ExtFilter), fmt.Sprintf(",%v%v,", buf[0], buf[1]))
	if !isExt && !filetype.IsImage(buf) && !filetype.IsVideo(buf) && !filetype.IsAudio(buf) && !filetype.IsArchive(buf) && !filetype.IsDocument(buf) {
		return utils.ErrorNull(c, fmt.Sprintf("%v文件格式错误，ext：%v%v", ext, buf[0], buf[1]))
	}
	fileName := utils.ID()
	path := GetPath(auth, dir, fileName, ext, loginInfo)
	dst, err := utils.CreateFile(path)
	if err != nil {
		return utils.ErrorNull(c, "创建文件失败")
	}
	defer dst.Close()
	if _, err = dst.Write(buf); err != nil {
		return utils.ErrorNull(c, "保存文件失败")
	}
	path = fmt.Sprintf("/%s", path)

	fileLogID := utils.ID()
	err = global.DB.Create(&model.FileLog{
		ID:        fileLogID,
		Name:      file.Filename,
		Path:      path,
		Size:      file.Size,
		Ext:       ext,
		IP:        c.RealIP(),
		CreatedBy: loginInfo.ID,
	}).Error
	if err != nil {
		global.Log.Error(fmt.Sprintf("保存上传文件日志失败，ERROR：%s", err.Error()))
	}
	return utils.SuccessNullMsg(c, map[string]interface{}{
		"id":   fileLogID,
		"size": file.Size,
		"path": path,
		"url":  getUrl(path),
		"name": file.Filename,
		"ext":  ext,
	})
}

//批量文件上传
func PluploadFile(c echo.Context) error {
	if err := c.Request().ParseMultipartForm(global.Conf.FileUpload.MaxFileSize); err != nil {
		return utils.ErrorNull(c, fmt.Sprintf("文件过大超出限制%vmb", fmt.Sprintf("%.2f", convert.MustFloat64(global.Conf.FileUpload.MaxFileSize)/b/b)))
	}

	c.Response().Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	c.Response().Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Response().Header().Add("Cache-Control", "post-check=0, pre-check=0")
	c.Response().Header().Set("Pragma", "no-cache")

	auth := c.Param("auth")
	loginInfo := global.GetLoginInfo(c)
	//放置目录
	dir := c.FormValue("dir")
	if dir == "" {
		dir = "default"
	}
	var chunkStr = c.FormValue("chunk")
	var chunksStr = c.FormValue("chunks")
	var chunk = 0
	var chunks = 1
	if chunkStr != "" {
		chunk, _ = strconv.Atoi(chunkStr)
	}
	if chunksStr != "" {
		chunks, _ = strconv.Atoi(chunksStr)
	}
	var result = make(map[string]interface{}, 0)
	for _, fileHeaders := range c.Request().MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			ext := path.Ext(fileHeader.Filename)

			var filename = ""
			if c.FormValue("name") != "" {
				filename = c.FormValue("name")
			} else if fileHeader.Filename != "" {
				filename = fileHeader.Filename
			} else {
				filename = utils.ID()
			}

			buf, err := ioutil.ReadAll(file)
			if err != nil {
				return utils.ErrorNull(c, "获取文件格式错误")
			}

			paths := GetPath(auth, dir, filename, "", loginInfo)

			log.Println(paths)

			tempPath := paths + ".part"
			exists, _ := utils.PathExists(tempPath)
			var f *os.File
			if !exists {
				f, err = utils.CreateFile(tempPath)
				if err != nil {
					global.Log.Error("OpenFile error:%s", err.Error())
					return utils.ErrorNull(c, "创建文件失败")
				}
			} else {
				f, err = os.OpenFile(tempPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					global.Log.Error("OpenFile error:%s", err.Error())
					return utils.ErrorNull(c, "创建文件失败")
				}
			}
			if _, err = f.Write(buf); err != nil {
				panic(err)
			}
			_ = f.Close()

			if chunks == 0 || chunk == chunks-1 {
				_ = os.Rename(paths+".part", paths)
				fileLogID := utils.ID()
				err = global.DB.Create(&model.FileLog{
					ID:        fileLogID,
					Name:      filename,
					Path:      paths,
					Size:      fileHeader.Size,
					Ext:       ext,
					IP:        c.RealIP(),
					CreatedBy: loginInfo.ID,
				}).Error
				if err != nil {
					global.Log.Error(fmt.Sprintf("保存上传文件日志失败，ERROR：%s", err.Error()))
				}
				paths = fmt.Sprintf("/%s", paths)
				result[filename] = map[string]interface{}{
					"id":    fileLogID,
					"size":  fileHeader.Size,
					"paths": paths,
					"url":   getUrl(paths),
					"name":  filename,
					"ext":   ext,
				}
			}
		}
	}
	return utils.SuccessNullMsg(c, result)
}

func GetPath(auth, dir, fileName, suffix string, loginInfo model.SysUserLoginInfo) string {
	if dir == "" {
		dir = "file"
	}
	if auth != "" {
		dir += auth
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s%s", global.Conf.FileUpload.BasePath+global.Conf.FileUpload.Path, loginInfo.ID, dir, strings.Replace(utils.CurrentDate(), "-", "/", -1), fileName, suffix)
}

func getUrl(path string) string {
	if strings.Contains(path, "http") {
		return global.Conf.FileUpload.DoMain + path
	} else {
		return path
	}
}

// 文件数据删除
func DelFileLog(c echo.Context) error {
	loginInfo := global.GetLoginInfo(c)
	id := c.FormValue("id")
	if id == "" {
		return utils.ErrorNull(c, utils.GetParsFailResult)
	}
	var pf model.FileLog
	if err := global.DB.Model(&model.FileLog{}).First(&pf, "id=?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.ErrorNull(c, "该文件记录不存在")
		}
		global.Log.Error("GetProjectType error：", err)
		return utils.ErrorNull(c, "获取文件记录失败")
	}
	if err := handle.PowerCheck(loginInfo, pf.CreatedBy); err != nil {
		return utils.ErrorNull(c, err.Error())
	}
	if err := global.DB.Where("id = ?", id).Delete(&model.FileLog{}).Error; err != nil {
		global.Log.Error("DelFileLog error：%v", err)
		return utils.ErrorNull(c, utils.DeleteFailResult)
	}
	return utils.SuccessNull(c, utils.DeleteSuccessResult)
}
