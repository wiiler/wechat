package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"wechat/utils"
)

func (c *Client) ImgUpload(fileName string, filePath string) (string, error) {

	type Meta struct {
		FileName string `json:"filename" binding:"required"` // 商户上传的媒体图片的名称，商户自定义，必须以JPG、BMP、PNG为后缀。
		Sha256   string `json:"sha256" binding:"required"`   // 图片文件的文件摘要，即对图片文件的二进制内容进行sha256计算得到的值。
	}

	// 读取文件
	pictureByes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// 计算文件序列化后的sha256
	h := sha256.New()
	if _, err = h.Write(pictureByes); err != nil {
		return "", err
	}
	meta := &Meta{}
	pictureSha256 := h.Sum(nil)
	meta.FileName = fileName
	meta.Sha256 = fmt.Sprintf("%x", string(pictureSha256))
	metaByte, err := json.Marshal(meta)
	if err != nil {
		return "", err
	}
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	// 设置reqbody中的meta部分
	if err = c.CreateFormField(writer, "meta", "application/json", metaByte); err != nil {
		return "", err
	}

	// 设置reqbody中的file部分
	if err = c.CreateFormFile(writer, fileName, "image/jpg", pictureByes); err != nil {
		return "", err
	}

	if err = writer.Close(); err != nil {
		return "", err
	}
	fmt.Println(string(metaByte), "string(metaByte)")
	// fmt.Println(reqBody.String(), "reqBody.String()")
	fmt.Println(writer.FormDataContentType(), "writer.FormDataContentType()")

	header := make(map[string]string)
	header[utils.ContentType] = writer.FormDataContentType()
	response, err := c.Upload(ImgUploadUrl, string(metaByte), reqBody.String(), header)
	if err != nil {

		return "", err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	// 获取回包中的信息
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("read response body err:%s", err.Error())
		return "", err
	}

	type Media struct {
		MediaID string `json:"media_id"`
	}
	media := new(Media)
	json.Unmarshal(body, media)
	return media.MediaID, nil
}

func (c *Client) VideoUpload(fileName string, filePath string) (string, error) {

	type Meta struct {
		FileName string `json:"filename" binding:"required"` // 商户上传的媒体图片的名称。
		Sha256   string `json:"sha256" binding:"required"`   // 文件的文件摘要，即对文件的二进制内容进行sha256计算得到的值。
	}

	// 读取文件
	pictureByes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// 计算文件序列化后的sha256
	h := sha256.New()
	if _, err = h.Write(pictureByes); err != nil {
		return "", err
	}
	meta := &Meta{}
	pictureSha256 := h.Sum(nil)
	meta.FileName = fileName
	meta.Sha256 = fmt.Sprintf("%x", string(pictureSha256))
	metaByte, err := json.Marshal(meta)
	if err != nil {
		return "", err
	}
	reqBody := &bytes.Buffer{}
	writer := multipart.NewWriter(reqBody)

	// 设置reqbody中的meta部分
	if err = c.CreateFormField(writer, "meta", "application/json", metaByte); err != nil {
		return "", err
	}

	// 设置reqbody中的file部分
	if err = c.CreateFormFile(writer, fileName, "image/jpg", pictureByes); err != nil {
		return "", err
	}

	if err = writer.Close(); err != nil {
		return "", err
	}
	fmt.Println(string(metaByte), "string(metaByte)")
	// fmt.Println(reqBody.String(), "reqBody.String()")
	fmt.Println(writer.FormDataContentType(), "writer.FormDataContentType()")

	header := make(map[string]string)
	header[utils.ContentType] = writer.FormDataContentType()
	response, err := c.Upload(VideoUploadUrl, string(metaByte), reqBody.String(), header)
	if err != nil {

		return "", err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	// 获取回包中的信息
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("read response body err:%s", err.Error())
		return "", err
	}

	type Media struct {
		MediaID string `json:"media_id"`
	}
	media := new(Media)
	json.Unmarshal(body, media)
	return media.MediaID, nil
}
