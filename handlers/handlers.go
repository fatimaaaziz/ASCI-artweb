package ourcode

import (
	"net/http"
	"text/template"
)

// HomeHandler handles GET requests to the root URL ("/").
// It renders the main HTML template for the form page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error  405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" && r.URL.Path != "/templates/index.html" {
		http.Error(w, " ERRORE 404 Page not found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "ERRORE 404 Template not found", http.StatusNotFound)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "ERRORE 500 Internal server error", http.StatusInternalServerError)
		return
	}
}

// AsciiArtHandler handles POST requests to "/ascii-art".
// It processes the submitted text and banner, validates input,
// generates ASCII art, and re-renders the template with the result.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if len(text) > 1000 {
		http.Error(w, "Error bad request 400", http.StatusBadRequest)
		return
	}
	if text == "" || banner == "" {
		http.Error(w, "Error bad request 400", http.StatusBadRequest)
		return
	}

	validBanners := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}

	if !validBanners[banner] {
		http.Error(w, "banner not found", http.StatusNotFound)
		return
	}

	result, err := GenerateASCIIArt(text, banner)
	if err != nil {
		renderWithError(w, text, banner, "/ Error generating ASCII art: "+err.Error())
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}

	data := PageData{
		Input:  text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
