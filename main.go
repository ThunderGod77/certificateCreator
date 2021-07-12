package main

import (
	"bytes"
	"certificateCreator/certificate"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/signal"
	"strings"
)

func fileType(file *multipart.FileHeader) (string, error) {

	fileType := strings.Split(file.Header.Get("Content-Type"), "/")
	if fileType[0] != "image" {
		return "", errors.New("incorrect file type")
	}
	return fileType[1], nil
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Use(logger.New())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		f, err := os.OpenFile("hello.pdf", os.O_RDWR, 0644)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		defer f.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, f); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		c.Set("Content-Disposition", "attachment; filename=lol.pdf")
		c.Set("Content-Type", "application/pdf")
		c.Write(buf.Bytes())
		//c.SendFile("./hello.pdf")

		return nil
	})
	app.Post("/certificate", func(c *fiber.Ctx) error {

		file, err := c.FormFile("logo")
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		imgT, err := fileType(file)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		f, err := file.Open()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		defer f.Close()
		signFile, err := c.FormFile("sign")
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		if signFile.Header.Get("Content-Type") != "image/svg+xml" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": "Incorrect signature file type!"})
		}
		sF, err := signFile.Open()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": err.Error()})
		}
		defer sF.Close()

		name := c.FormValue("name")
		description := c.FormValue("description")
		style := c.FormValue("style")
		local := c.FormValue("local")

		p := certificate.CreateCertificate(f, imgT, sF, name, description, style)

		if local == "yes" {
			err = p.OutputFileAndClose("hello.pdf")
			if err != nil {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": err.Error()})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": "none"})
		}

		buf := bytes.NewBuffer(nil)

		err = p.Output(buf)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": err.Error()})
		}
		c.Set("Content-Disposition", "attachment; filename=lol.pdf")
		c.Set("Content-Type", "application/pdf")
		c.Write(buf.Bytes())

		return nil

	})

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
