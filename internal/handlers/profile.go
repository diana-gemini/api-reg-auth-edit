package handler

import (
	"fmt"
	"api/internal/render"
	"api/internal/types"
	"net/http"
	"strings"
)

func (h *Handler) editProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/editprofile" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	user := h.getUserFromContext(r)

	render.Render(w, "ui/html/editprofile.html", render.WebPage{
		IsLoggedin: true,
		User:       user,
	})
}

func (h *Handler) profileEditSave(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profileeditsave" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}
	if r.Method == http.MethodGet {
	} else if r.Method == http.MethodPost {

		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}

		author := h.getUserFromContext(r)

		profile := &types.EditProfile{
			AuthorId:  author.Id,
			Username:  strings.TrimSpace(r.Form.Get("username")),
			Mobile:    strings.TrimSpace(r.Form.Get("mobile")),
			BirthDate: strings.TrimSpace(r.Form.Get("birthdate")),
		}

		err = h.service.PostService.UpdateProfile(profile)
		if err != nil {
			fmt.Println(err)
			ErrorPage(w, r, http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
}
