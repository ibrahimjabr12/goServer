package main 
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"strconv"
	"encoding/json"
	"io/ioutil"
)
type person struct {
    name string `json:"name,omitempty"`
	age  int64	`json:"age,omitempty"`
	favColor string `json:"favColor,omitempty"`
}
var people = make(map[string]person)
var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("temp/*"))
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	tpl.ExecuteTemplate(w,"add.gohtml",nil)
}
func add(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	tpl.ExecuteTemplate(w,"add.gohtml",nil)
	n:= r.Form["name"][0]
	a, err := strconv.ParseInt(r.Form["age"][0], 10, 64)
	if err!= nil {
		panic(err)
	}
	fC:= r.Form["favColor"][0]
	 people[n]=person{
		name:n,
		age:a,
		favColor:fC,
	}
}
func getOnePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name :=ps.ByName("name") 
	if name == ""{
		name = r.Form["Fname"][0]
	}
	jsonData, err := json.MarshalIndent(people[name],""," ")
	if err != nil{
		panic(err)
	}
	tpl.ExecuteTemplate(w,"getone.gohtml",string(jsonData))
}
func getAllPeople(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	jsonData, err := json.MarshalIndent(people,""," ")
	if err != nil{
		panic(err)
	}
	_ = ioutil.WriteFile("allPeople.json", jsonData, 0644)
	var de = make(map[string]person)
	if err := json.Unmarshal(jsonData, &de); err != nil {
        panic(err)
    }
	tpl.ExecuteTemplate(w,"getone.gohtml",string(jsonData))
}
  func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.POST("/people", add)
	router.GET("/people/:name", getOnePerson)
	router.GET("/people", getAllPeople)
	// http.HandleFunc("/people",d)
	// http.HandleFunc("/people/",c)
	http.ListenAndServe(":3000",router)
  }