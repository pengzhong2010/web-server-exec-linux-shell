package controllers

import (
	"bytes"
	_ "encoding/hex"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pengzhong2010/web-server-exec-linux-shell/models"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {
	// fmt.Println("dd")
	cmd_model := u.GetString("model")
	if cmd_model == "" {
		u.Ctx.WriteString("model is empty")
		return
	}
	cmd_args := u.GetString("args")
	if cmd_args == "" {
		u.Ctx.WriteString("args is empty")
		return
	}

	cmd := exec.Command("sh", beego.AppConfig.String("execShellPath")+cmd_model+".sh", cmd_args)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
	info := models.Info{GetRandomString(), "b", out.String()}

	fmt.Println(cmd_model)
	fmt.Println(info.Id)
	fmt.Println(info.Status)
	fmt.Println(info.Mes)
	u.Data["json"] = &info
	u.ServeJSON()

}

func GetRandomString() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 24; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
