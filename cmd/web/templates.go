package main

import (
	"html/template"
	"path/filepath"
	"snippetbox.codegeek.net/internal/models"
)

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}
	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.gohtml". This will essentially give
	// us a slice of all the filepaths for our application 'page' templates
	// like: [ui/html/pages/home.gohtml ui/html/pages/view.gohtml]
	pages, err := filepath.Glob("./ui/html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}
	// Loop through the page filepaths one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.gohtml') from the full filepath
		// and assign it to the name variable.
		name := filepath.Base(page)
		// Create a slice containing the filepaths for our base template, any
		// partials and the page.
		files := []string{
			"./ui/html/base.gohtml",
			"./ui/html/partials/nav.gohtml",
			page,
		}
		// Parse the files into a template set.
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}
		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}
	// Return the map.
	return cache, nil
}

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
