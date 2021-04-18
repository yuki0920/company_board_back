package controllers

import "github.com/gofiber/fiber/v2"

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	filename := ""

	// フォームに入ったファイルを1つずつ保存する
	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8080/uploads/" + filename,
	})
}