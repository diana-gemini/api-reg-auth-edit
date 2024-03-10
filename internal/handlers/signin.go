package handler

import (
	"fmt"
	"api/internal/cookies"
	"api/internal/render"
	"api/internal/types"
	"net/http"
	
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
	} else if r.Method == http.MethodPost { //
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		user := &types.GetUserData{
			Email: r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		userid, err := h.service.UserService.CheckLogin(user)
		if err == nil {
			cookieToken := cookies.SetCookie(w) 
			h.service.UserService.AddToken(userid, cookieToken)
	
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			render.Render(w, "ui/html/signin.html", render.WebPage{
				Errtext: "Username or password is incorrect",
			})

			return
		}

	} else {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	render.Render(w, "ui/html/signin.html", render.WebPage{})
}
