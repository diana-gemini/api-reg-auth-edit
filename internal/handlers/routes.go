package handler

import (
	"net/http"
)

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/signup", h.signup)
	mux.HandleFunc("/signin", h.signIn)
	mux.HandleFunc("/logout", h.logout)
	mux.HandleFunc("/editprofile", h.requireAuth(h.editProfile))
	mux.HandleFunc("/profileeditsave", h.requireAuth(h.profileEditSave))
	return h.middleware(mux)
}
