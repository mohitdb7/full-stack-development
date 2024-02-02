package domain

import (
	"file-upload-go/common"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type IUpload interface {
	SaveFile(multipart.File, *multipart.FileHeader) (string, error)
}

type LocalUpload struct {
}

func (lu LocalUpload) SaveFile(file multipart.File, handler *multipart.FileHeader) (string, error) {
	//2. Retrieve file from form-data
	//<Form-id> is the form key that we will read from. Client should use the same form key when uploading the file
	defer file.Close()

	fmt.Printf("File name: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("File header: %+v\n", handler.Header)

	//3. Create a temporary file to our directory
	tempFolderPath := fmt.Sprintf("%s%s", common.RootPath, "/tempFiles")
	tempFileName := fmt.Sprintf("upload-%s-*.%s", common.FileNameWithoutExtension(handler.Filename), filepath.Ext(handler.Filename))

	tempFile, err := os.CreateTemp(tempFolderPath, tempFileName)
	if err != nil {
		errStr := fmt.Sprintf("Error in creating the file %s\n", err)
		fmt.Println(errStr)
		return errStr, err
	}

	defer tempFile.Close()

	//4. Write upload file bytes to your new file
	filebytes, err := io.ReadAll(file)
	if err != nil {
		errStr := fmt.Sprintf("Error in reading the file buffer %s\n", err)
		fmt.Println(errStr)
		return errStr, err
	}

	tempFile.Write(filebytes)
	return "Successfully uploaded\n", nil
}

func NewLocalUpload() LocalUpload {
	return LocalUpload{}
}
