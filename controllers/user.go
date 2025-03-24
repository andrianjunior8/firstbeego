package controllers

import (
	"archive/zip"
	"encoding/json"
	"firstbeego/models"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/xuri/excelize/v2"
)

type UserController struct {
    beego.Controller
}

// Get all users
func (c *UserController) Get() {
    users := []models.User{
        {ID: 1, Name: "Alice", Age: 25},
        {ID: 2, Name: "Bob", Age: 30},
    }

    c.Data["json"] = users
    c.ServeJSON()
}

// Get user by ID
func (c *UserController) GetById() {
    id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
    if err != nil {
        c.Ctx.ResponseWriter.WriteHeader(400)
        c.Data["json"] = map[string]string{"error": "Invalid ID"}
        c.ServeJSON()
        return
    }

    user := models.User{ID: id, Name: "User " + strconv.Itoa(id), Age: 20 + id}
    c.Data["json"] = user
    c.ServeJSON()
}

func (c *UserController) Post() {
    o := orm.NewOrm()

    var user models.User
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
        c.Ctx.ResponseWriter.WriteHeader(400)
        c.Data["json"] = map[string]string{"error": "Invalid JSON input"}
        c.ServeJSON()
        return
    }

    // Insert the user into the database
    id, err := o.Insert(&user)
    if err != nil {
        c.Ctx.ResponseWriter.WriteHeader(500)
        c.Data["json"] = map[string]string{"error": "Failed to insert user"}
        c.ServeJSON()
        return
    }

    f:= excelize.NewFile()

    defer func(){
        if err:= f.Close(); err!= nil {
            fmt.Println(err)
        } 
    }()

    index, err := f.NewSheet("sheet2")
    if err!=nil{
        fmt.Println(err)
        return
    }

    f.SetCellValue("Sheet2", "A2", "Hello world.")
    f.SetCellValue("Sheet1", "B2", 100)
    // Set active sheet of the workbook.
    f.SetActiveSheet(index)
    // Save spreadsheet by the given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
    }

    fmt.Println("creating zip archive...")
    archive, err := os.Create("archive.zip")
    if err != nil {
        panic(err)
    }
    defer archive.Close()
    zipWriter := zip.NewWriter(archive)

    fmt.Println("opening first file...")
    f1, err := os.Open("Book1.xlsx")
    if err != nil {
        panic(err)
    }
    defer f1.Close()

    fmt.Println("writing first file to archive...")
    w1, err := zipWriter.Create("Book1.xlsx")
    if err != nil {
        panic(err)
    }
    if _, err := io.Copy(w1, f1); err != nil {
        panic(err)
    }

    fmt.Println("closing zip archive...")
    zipWriter.Close()

    user.ID = int(id)
    c.Data["json"] = user
    c.ServeJSON()
}
