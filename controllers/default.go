package controllers

import (
	"ExpSend/models"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "index.tpl"
	c.TplName = "index.tpl"
	a := get_str.ApiResponse{}
	if apiKey := c.Ctx.Request.FormValue("apiKey"); apiKey == "" {
		return
	} else {
		outFileSingle, err := os.Create("views/SingleOptIn.html")
		if err != nil {
			panic(err)
		}
		defer outFileSingle.Close()
		outFileDouble, err := os.Create("views/DoubleOptIn.html")
		if err != nil {
			panic(err)
		}
		defer outFileDouble.Close()
		resp, err := http.Get("https://api.esv2.com/Api/Lists?apiKey=" + apiKey)
		if err != nil {
			log.Fatal(err)
		}
		if code := resp.StatusCode; code != 200 {
			c.Data["Code"] = "Неверный apiKey. Повторите попытку."
			return
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = xml.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		tplSingle, _ := template.New("SingleOptIn").Parse("<tr><td>{{.Id}}</td><td>{{.Name}}</td><td>{{.FriendlyName}}</td><td>{{.Language}}</td></tr>\n")
		tplDouble, _ := template.New("DoubleOptIn").Parse("<tr><td>{{.Id}}</td><td>{{.Name}}</td><td>{{.FriendlyName}}</td><td>{{.Language}}</td></tr>\n")
		for _, v := range a.Data {
			tplVars := map[string]string{
				"Id":           v.Id,
				"Name":         v.Name,
				"FriendlyName": v.FriendlyName,
				"Language":     v.Language,
			}
			if v.OptInMode == "SingleOptIn" {
				c.Data["Table1"] = v.OptInMode
				tplSingle.Execute(outFileSingle, tplVars)
			} else {
				c.Data["Table2"] = v.OptInMode
				tplDouble.Execute(outFileDouble, tplVars)
			}
		}
	}
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SingleOptIn"] = "SingleOptIn.html"
	c.LayoutSections["DoubleOptIn"] = "DoubleOptIn.html"
}
func (c *MainController) Post() {
	c.TplName = "index.tpl"
	if c.Ctx.Request.Method == "POST" {
		xmlBody := `
<ApiRequest xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xs="http://www.w3.org/2001/XMLSchema">
	<ApiKey>.ApiKey</ApiKey>
   	<Data>
    	<Recipients>
			<SeedLists>
    			<SeedList>.SeedList</SeedList>
			</SeedLists>
    	</Recipients>
    	<Content>
			<FromName>.FromName</FromName>
    		<FromEmail>.FromEmail</FromEmail>
    		<Subject>.Subject</Subject>
			<Html><![CDATA[.Table1]]><![CDATA[.Table2]]>
			</Html>
    	</Content>
    </Data>
 	</ApiRequest>`
		ApiKey := c.Ctx.Request.FormValue("apiKey")
		SeedList := c.Ctx.Request.FormValue("SeedList")
		FromName := c.Ctx.Request.FormValue("FromName")
		FromEmail := c.Ctx.Request.FormValue("FromEmail")
		Subject := c.Ctx.Request.FormValue("Subject")
		Table1, err := ioutil.ReadFile("views/SingleOptIn.html")
		if err != nil {
			panic(err)
		}
		Table2, err := ioutil.ReadFile("views/DoubleOptIn.html")
		if err != nil {
			panic(err)
		}
		xmlBody = strings.Replace(xmlBody, ".ApiKey", ApiKey, -1)
		xmlBody = strings.Replace(xmlBody, ".SeedList", SeedList, -1)
		xmlBody = strings.Replace(xmlBody, ".FromName", FromName, -1)
		xmlBody = strings.Replace(xmlBody, ".FromEmail", FromEmail, -1)
		xmlBody = strings.Replace(xmlBody, ".Subject", Subject, -1)
		xmlBody = strings.Replace(xmlBody, ".Table1", string(Table1), -1)
		xmlBody = strings.Replace(xmlBody, ".Table2", string(Table2), -1)
		b := get_str.ApiResponse2{}
		resp, err := http.Post("https://api.esv2.com/Api/Newsletters", "text/xml", strings.NewReader(xmlBody))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err = xml.Unmarshal(data, &b)
		if err != nil {
			panic(err)
		}
		c.Data["Code2"] = b.Data
	}
}
