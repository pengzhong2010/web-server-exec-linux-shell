package controllers

import (
	"bytes"
	"dockerimagesmake/models"
	_ "encoding/hex"
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	// fmt.Println("dd")
	cmd_model := u.GetString("model")
	if cmd_model == "" {
		u.Ctx.WriteString("model is empty")
		return
	}
	cmd := exec.Command("sh", "/data/githook/k8s-apply/apply.sh", ""+cmd_model+"")
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
