package entities

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type File struct {
	Id         int64  `json:"id" db:"id"`
	Name       string `json:"file_name" db:"name"`
	PathUrl    string `json:"file_data" db:"pathUrl"`
	Type       string `json:"file_type" db:"type"`
	EntityType string `json:"entity_type" db:"entityType"`
	EntityId   int64  `json:"entity_id" db:"entityId"` //- id ไว้ใช้อิงตาม model ที่เกี่ยวข้อง
	CreatedAt  string `json:"created_at" db:"createdAt"`
	UpdatedAt  string `json:"updated_at" db:"updatedAt"`
}

func (f *File) Base64toPng(c *fiber.Ctx) (*string, *string, error) {

	if len(f.PathUrl) == 0 || f.Type != "PNG" {
		return nil, nil, errors.New("Invalid file data or file type, expected PNG file type but got " + f.Type)
	}

	hasher := sha256.New()
	hasher.Write([]byte(f.PathUrl))
	hash := hex.EncodeToString(hasher.Sum(nil))

	path := "public/image/"
	pngFilename := path + hash + ".png"

	if _, err := os.Stat(pngFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(f.PathUrl))
		m, _, err := image.Decode(reader)
		if err != nil {
			return nil, nil, err
		}
		// bounds := m.Bounds()
		// fmt.Println(bounds, formatString)

		osFile, errOnOpenFIle := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if errOnOpenFIle != nil {
			return nil, nil, err
		}
		err = png.Encode(osFile, m)
		if err != nil {
			return nil, nil, err
		}
		buffer := new(bytes.Buffer)
		errWhileEncoding := png.Encode(buffer, m) // img is your image.Image
		if errWhileEncoding != nil {
			return nil, nil, err
		}
		base64url := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		filePathData := fmt.Sprintf("%s/%s", c.Hostname(), pngFilename)
		log.Info("Create new PNG file name: ", pngFilename, "as the output")

		return &base64url, &filePathData, nil
	}

	data, err := os.ReadFile(pngFilename)
	if err != nil {
		return nil, nil, err
	}

	base64url := "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
	filePathData := fmt.Sprintf("%s/%s", c.Hostname(), pngFilename)
	log.Info("Reusing exist PNG file name: ", pngFilename, "as the output")

	return &base64url, &filePathData, nil

}

func (f *File) Base64toJpg(c *fiber.Ctx) (*string, *string, error) {

	if len(f.PathUrl) == 0 || f.Type != "JPG" {
		return nil, nil, errors.New("Invalid file data or file type, expected PNG file type but got " + f.Type)
	}

	hasher := sha256.New()
	hasher.Write([]byte(f.PathUrl))
	hash := hex.EncodeToString(hasher.Sum(nil))

	jpgFilename := "public/image/" + hash + ".jpg"

	if _, err := os.Stat(jpgFilename); os.IsNotExist(err) {
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(f.PathUrl))
		m, formatString, err := image.Decode(reader)
		if err != nil {
			return nil, nil, err
		}
		bounds := m.Bounds()
		fmt.Println("base64toJpg", bounds, formatString)

		osFile, err := os.OpenFile(jpgFilename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return nil, nil, err
		}

		err = jpeg.Encode(osFile, m, &jpeg.Options{Quality: 75})
		if err != nil {
			return nil, nil, err
		}

		buffer := new(bytes.Buffer)
		errWhileEncoding := jpeg.Encode(buffer, m, nil) // img is your image.Image
		if errWhileEncoding != nil {
			log.Fatal(errWhileEncoding)
		}
		base64url := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
		filePathData := fmt.Sprintf("%s/%s", c.Hostname(), jpgFilename)
		log.Info("Create new JPG file name: ", jpgFilename, "as the output")

		return &base64url, &filePathData, nil
	}

	data, err := os.ReadFile(jpgFilename)

	if err != nil {
		return nil, nil, err
	}

	base64url := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(data)
	filePathData := fmt.Sprintf("%s/%s", c.Hostname(), jpgFilename)

	log.Info("Reusing exist JPG file name: ", jpgFilename, "as the output")

	return &base64url, &filePathData, nil
}

func (f *File) Base64toFile(c *fiber.Ctx, includeDomain bool) (*string, *string, error) {
	if len(f.PathUrl) == 0 || f.Type != "PDF" || f.Type != "MARKDOWN_FILE" {
		return nil, nil, errors.New("Invalid file data or file type, expected PDF file type but got " + f.Type)
	}

	// encode blob to string
	hasher := sha256.New()
	hasher.Write([]byte(f.PathUrl))
	hash := hex.EncodeToString(hasher.Sum(nil))

	var fileName string
	switch f.Type {
	case "PDF":
		fileName = "public/file/" + hash + ".pdf"
	case "MARKDOWN_FILE":
		fileName = "public/file/" + hash + ".md"
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {

		data, err := base64.StdEncoding.DecodeString(f.PathUrl)
		if err != nil {
			return nil, nil, err
		}

		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			return nil, nil, err
		}

		srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.Type), base64.StdEncoding.EncodeToString(data))

		log.Info("Reusing exist ", f.Type, " file name: ", fileName, "as the output")

		filePathData := fileName
		if includeDomain {
			filePathData = fmt.Sprintf("%s/%s", c.Hostname(), fileName)
		}
		return &srcFile, &filePathData, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, nil, err
	}

	srcFile := fmt.Sprintf("data:file/%s;base64,%s", strings.ToLower(f.Type), base64.StdEncoding.EncodeToString(data))
	filePathData := fileName
	if includeDomain {
		filePathData = fmt.Sprintf("%s/%s", c.Hostname(), fileName)
	}
	log.Info("Reusing exist ", f.Type, " file name: ", fileName, "as the output")

	return &srcFile, &filePathData, nil

}

func (file *File) EncodeBase64toFile(c *fiber.Ctx, domainIncludeOnFile bool) (*string, *string, int, error) {
	var (
		base64urlRes string
		fPathDatRes  string
	)
	switch file.Type {
	case "PNG":
		base64url, fPathDat, err := file.Base64toPng(c)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "JPG":
		base64url, fPathDat, err := file.Base64toJpg(c)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	case "PDF":
		base64url, fPathDat, err := file.Base64toFile(c, domainIncludeOnFile)
		if err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		base64urlRes = *base64url
		fPathDatRes = *fPathDat
	default:
		return nil, nil, http.StatusUnsupportedMediaType, errors.New("Only except for PNG, JPG, PDF AND MD for now")
	}

	return &base64urlRes, &fPathDatRes, http.StatusOK, nil
}

func (f *File) DeleteFile(c *fiber.Ctx) error {
	if len(f.PathUrl) == 0 || f.Type != "PDF" || f.Type != "MARKDOWN_FILE" || f.Type != "PNG" || f.Type != "JPG" {
		return errors.New("Invalid file data or file type, expected file data but got nil or file type is not supported ")
	}

	hasher := sha256.New()
	hasher.Write([]byte(f.PathUrl))
	hash := hex.EncodeToString(hasher.Sum(nil))

	var fileName string
	if f.Type == "PDF" {
		fileName = "public/file/" + hash + ".pdf"
	} else {
		fileName = "public/image/" + hash + fmt.Sprintf(".%s", strings.ToLower(f.Type))
	}
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	log.Info("Deleted file name: ", fileName)
	return nil

}
