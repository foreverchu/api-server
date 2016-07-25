package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
	"github.com/slene/blackfriday"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	files, err := ListDir("../../../doc", "api")
	if err != nil {
		log.Fatal(err)
	}
	if len(files) <= 0 {
		log.Fatal("length of files < 0")
	}

	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		md := NewMarkdown()

		title := fileName
		md.SetTitle(title)
		md.SetTemplateFile("github_style_markdown_template.tpl")

		err = md.Generate(file)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(getDstFileName(fileName))
		err = md.WriteTo(getDstFileName(fileName))
		if err != nil {
			log.Fatal(err)
		}
	}

}

func getDstFileName(src string) (dst string) {
	dst = strings.TrimSuffix(strings.TrimPrefix(src, "../../../doc/"), ".md")
	return fmt.Sprintf("../../static/%s.html", dst)
}

func ListDir(dirPth string, prefix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	prefix = strings.ToUpper(prefix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasPrefix(strings.ToUpper(fi.Name()), prefix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

type Markdown struct {
	inputReader  io.Reader
	outputReader io.Reader
	title        string
	templateFile string
}

func NewMarkdown() *Markdown {
	return &Markdown{}
}

func (md *Markdown) SetTitle(title string) {
	md.title = title
}

func (md *Markdown) SetTemplateFile(templateFile string) {
	md.templateFile = templateFile
}

func (md *Markdown) Output() io.Reader {
	return md.outputReader
}

func (md *Markdown) Generate(inputReader io.Reader) error {
	md.inputReader = inputReader
	input, err := ioutil.ReadAll(md.inputReader)
	if err != nil {
		return err
	}
	unsafe := blackfriday.MarkdownCommon(input)
	output := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	md.outputReader = bytes.NewReader(output)
	return nil
}

func (md *Markdown) WriteTo(fileName string) (err error) {
	content, err := ioutil.ReadAll(md.Output())
	if err != nil {
		return
	}
	htmlFile, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer htmlFile.Close()

	t, err := template.ParseFiles(md.templateFile)
	if err != nil {
		return
	}
	data := struct {
		Title   string
		Content string
	}{
		Title:   md.title,
		Content: string(content),
	}
	err = t.Execute(htmlFile, data)
	if err != nil {
		return
	}
	return
}

func (md *Markdown) Reset() {
	md.inputReader = nil
	md.outputReader = nil
}
