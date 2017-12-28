package main


import (

	//"fmt"
	"net/http"
	"html/template"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)





func main()  {

	//todo html page to upload and display image

	server :=http.Server{
		Addr:":3000",

	}

	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	http.HandleFunc("/logout",logout)
	server.ListenAndServe()
	//todo function to upload image

	//todo function to resize image one big one small


	//todo function watermark image

}

func login(w http.ResponseWriter, r *http.Request)  {
	t, _:= template.ParseFiles("login.html")
	_, err := r.Cookie("Profile")
	if err !=nil{
		http.SetCookie(w,&http.Cookie{
			Name:"Profile",
			Value:"asfsdjhfjhejhjkfhdjkhfjkdshfjkahej",
			MaxAge:7000,
		})
	}else {
		http.Redirect(w, r, "/upload", 303)
	}
	t.Execute(w,"")
}

func upload(w http.ResponseWriter, r *http.Request)  {
	t, _:= template.ParseFiles("upload.html")
	var s string
	if r.Method == http.MethodPost{
		f,h,err := r.FormFile("fileToUpload")
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		defer f.Close()
		fmt.Println("\nfile",f,"\nheader:",h.Filename,"\n",h.Size,"\nerr",err)
		//read
		bs, err := ioutil.ReadAll(f)

		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		s = string(bs)

		dst, _ := os.Create(filepath.Join("./img/",h.Filename))

		defer dst.Close()

		dst.Write(bs)
	}
	t.Execute(w,s)
}

func logout(w http.ResponseWriter, r *http.Request)  {
	t, _:= template.ParseFiles("logout.html")


	c, _ := r.Cookie("Profile")

	if c.Name=="Profile"{

		c.Value="Talk"
	}

	t.Execute(w,"")
}
