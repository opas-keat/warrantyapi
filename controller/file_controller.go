package controller

import (
	"fmt"
	"path/filepath"
	"time"
	"warrantyapi/common"
	"warrantyapi/configuration"
	"warrantyapi/model"

	"github.com/gofiber/fiber/v2"
)

type FileController struct {
	configuration.Config
}

func NewFileController(config configuration.Config) *FileController {
	return &FileController{Config: config}
}

func (controller FileController) Route(app *fiber.App) {
	apiV1 := app.Group(controller.Config.Get("API_CONTEXT_PATH") + "/v1")
	api := apiV1.Group("/file_attach")
	api.Post("/", controller.create)
	// api.Get("/", controller.list)
}

func (controller FileController) create(c *fiber.Ctx) error {
	fileAttachInput := new(model.FileAttach)
	if err := c.BodyParser(fileAttachInput); err != nil {
		print("An error occurred when parsing the FileAttachInput: " + err.Error())
	}
	//WT-202309101822161
	println("Id: " + fileAttachInput.LinkId)
	println("link_type: " + fileAttachInput.LinkType)
	// userName := middleware.GetUserNameFromToken(c)
	userName := "admin"
	extension := filepath.Ext(fileAttachInput.FileName)
	println("extension: " + extension)
	//Save data to FileAttach
	fileAttach := model.FileAttach{
		CreatedBy: userName,
		FileName:  fileAttachInput.FileName,
		FileType:  extension,
		FileSize:  fileAttachInput.FileSize,
		LinkType:  fileAttachInput.LinkType,
		LinkId:    fileAttachInput.LinkId,
		Module:    fileAttachInput.Module,
	}
	// println(fileAttachInput.LinkId[3:11])
	path := common.CreatePathFileForUpload(time.Now().Format("20060102"))
	println("path: " + path)
	file, err := c.FormFile(fileAttachInput.LinkType)
	if err != nil {
		println(err)
		return err
	}
	c.SaveFile(file, fmt.Sprintf(path+"%s", fileAttachInput.LinkId+"_"+fileAttachInput.LinkType+extension))
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    "000",
		Message: "Success",
		Data:    fileAttach,
	})
}
