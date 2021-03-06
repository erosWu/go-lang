package service

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

type User struct {
	Username string
	Password string
	Phone    string
	Email    string
}

func tableHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {

		}
		user := new(User)
		decoder := schema.NewDecoder()
		err = decoder.Decode(user, req.PostForm)
		if err != nil {

		}
		//将从html中获取到的相关信息按照自定义格式（json）输出
		formatter.HTML(w, http.StatusOK, "table", struct {
			Username string
			Password string
			Phone    string
			Email    string
		}{Username: user.Username, Password: user.Password, Phone: user.Phone, Email: user.Email})
	}
}
