package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

// QuestionRecord for recording question
type QuestionRecord struct {
	Number   int
	Question template.HTML
}

// TestBaseTemplateItem used in questionsPages page
type TestBaseTemplateItem struct {
	DB   string
	Type string
}

type QuestionsPageData struct {
	DB        string                     //题库数据库名称
	Types     []string                   //题库内试题类型
	Questions map[string][]template.HTML //[试题类型: 试题列表]
}

//type QuestType string
type QuestItem struct {
	QuestNote string
	Questions []int
}
type TestInfoForm struct {
	DatabaseName     string                // 存储试题库的数据库
	PaperFileName    string                //试卷名称
	CourseName       string                //课程名称
	TeacherForCourse string                // 任课教师
	TeacherForTest   string                // 出卷教师
	YearAndTerm      string                // 学年学期
	ClassName        string                // 班级名称
	OpenOrCloseTest  string                // 开卷闭卷？
	ABCSelected      string                // A,B,C卷？
	Questions        map[string]*QuestItem // [questType]*QuestItem
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/scripts", "./scripts")
	r.Static("/figures", "./figures")

	// "/" => index.html
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"DB": DBConfig().DataBaseItems})
	})

	// "/login" => login.html
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// "/questionsPages/*"  GET
	r.GET("/questionsPages/:dbname", func(c *gin.Context) {
		pageData := QuestionsPageData{}
		dbname := c.Param("dbname")
		if FillTestPageDataFromDB(dbname, &pageData) == false {
			return
		}
		c.HTML(http.StatusOK, "testpaperinfo.html", gin.H{
			"QuesTypes": pageData.Types,
			"Questions": pageData.Questions,
		})
		return
	})

	// "questionsPages/*"  POST
	r.POST("/questionsPages/:dbname", func(c *gin.Context) {
		// 从试题库中提取所有试题
		dbname := c.Param("dbname")
		pageData := QuestionsPageData{}
		if FillTestPageDataFromDB(dbname, &pageData) == false {
			return
		}

		// 从浏览器客户端提交的表单提取试卷信息
		testInfoForm := TestInfoForm{}
		testInfoForm.Questions = make(map[string]*QuestItem)
		testInfoForm.DatabaseName = dbname
		testInfoForm.PaperFileName = c.PostForm("paperFileName")
		testInfoForm.CourseName = c.PostForm("courseName")
		testInfoForm.TeacherForCourse = c.PostForm("teacherForCourse")
		testInfoForm.TeacherForTest = c.PostForm("teacherForTest")
		testInfoForm.YearAndTerm = c.PostForm("yearAndTerm")
		testInfoForm.ClassName = c.PostForm("className")
		testInfoForm.ABCSelected = c.PostForm("abcSelected")
		testInfoForm.OpenOrCloseTest = c.PostForm("openOrCloseTest")

		// 从浏览器客户端提交的表单提取试题类型、试题描述及试题题目列表
		for _, qtype := range pageData.Types {
			if v := c.PostForm(qtype + "Name"); v != "" { // 题目
				v = strings.TrimSpace(v)
				v = strings.Replace(v, ",", " ", -1)
				vs := strings.Fields(v)
				questions := make([]int, 0)
				for _, numstr := range vs {
					var num int
					fmt.Sscanf(numstr, "%d", &num)
					questions = append(questions, num)
				}
				if len(questions) == 0 { // 没有试题，略过该类型的试题
					continue
				}
				// 提取题目描述
				questNote := c.PostForm(qtype + "Note")
				questItem := QuestItem{
					QuestNote: questNote,
					Questions: make([]int, 0),
				}
				for _, v := range questions {
					questItem.Questions = append(questItem.Questions, v)
				}
				testInfoForm.Questions[qtype] = &questItem
			}
		}
		for qtype, quest := range testInfoForm.Questions {
			fmt.Println(qtype, ": ", quest)
		}

		// 提交testInfoForm 到 生产程序 生成latex, pdf 等文件
		questionFileName, answerFileName := ProduceTestPaper(&testInfoForm)
		fmt.Println(questionFileName, answerFileName)

		xelatexProgram, err := exec.LookPath("xelatex.exe")
		if err != nil {
			xelatexProgram, err = exec.LookPath("xelatex")
			if err != nil {
				fmt.Print("Can't find xelatex execute program!\n")
				return
			}
		}
		cmd1 := exec.Command(xelatexProgram, questionFileName)
		cmd2 := exec.Command(xelatexProgram, answerFileName)
		err = cmd1.Run()
		if err != nil {
			fmt.Printf("error when xelatex %s\n", questionFileName)
			return
		}
		err = cmd2.Run()
		if err != nil {
			fmt.Printf("error when xelatex %s\n", answerFileName)
			return
		}

		cmd3 := exec.Command(xelatexProgram, questionFileName)
		cmd4 := exec.Command(xelatexProgram, answerFileName)
		cmd3.Run()
		cmd4.Run()

		// locals := make(map[string]interface{})
		// images := []string{}
		// images = append(images, testInfoForm.PaperFileName+"Q.pdf")
		// images = append(images, testInfoForm.PaperFileName+"A.pdf")
		// locals["images"] = images
		// templ := path.Join(Templates_Dir, "listpdf.html")
		// template.Must(template.ParseFiles(templ)).Execute(w, locals)

		return
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

// FillTestPageDataFromDB 返回数据库内所有试题
func FillTestPageDataFromDB(dbname string, pageData *QuestionsPageData) bool {
	base := DBItemByName(dbname)

	// open db for dbname
	db, err := OpenDB(base.Path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer db.Close()
	pageData.DB = dbname

	// fill pageData.Types
	rows, err := db.Query("SELECT name FROM sqlite_master;")
	if err != nil {
		fmt.Println(err)
		return false
	}
	for rows.Next() {
		var t string
		if rows.Scan(&t) == nil {
			pageData.Types = append(pageData.Types, t)
		}
	}
	rows.Close()

	// fill pageData.Questions
	M := make(map[string][]template.HTML)
	for _, qtype := range pageData.Types {
		questions := make([]template.HTML, 0)
		rows, err := db.Query(fmt.Sprintf("SELECT question FROM \"%s\"", qtype))
		if err != nil {
			fmt.Println(err)
			return false
		}
		for rows.Next() {
			var t string
			rows.Scan(&t)
			questions = append(questions, template.HTML(markdownToHTML(t)))
		}
		rows.Close()
		M[qtype] = questions
	}
	pageData.Questions = M

	return true
}
