// Filename: main.go
package main

import (
	"log"
	"net/http"
	"time"
	"html/template"
)


// Middleware function to log requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Printf("Method: %s | URI: %s | Time: %s",
				r.Method, r.URL.Path, start.Format(time.RFC3339))
			next.ServeHTTP(w, r)
		})
}

// Handlers for blog pages
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/home.html")
}

func week1Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/week1.tmpl")
	if err != nil { 
     log.Print(err.Error())
     http.Error(w, "Internal Server Error",http.StatusInternalServerError)   
     return 
	}
	data := map[string]string {
    "Title": "Life in Adavance Database",
	"Header": "Week 1",
	}
	
	err = tmpl.Execute(w, data)
	if err != nil { 
		log.Print(err.Error())
		http.Error(w, "Internal Server Error",http.StatusInternalServerError)   
		return 
   }
}

func week2Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/week2.tmpl")
	if err != nil { 
     log.Print(err.Error())
     http.Error(w, "Internal Server Error",http.StatusInternalServerError)   
     return 
	}
	data := map[string]string {
    "Title": "Life in Adavance Database",
	"Header": "Week 1",
	}
	
	err = tmpl.Execute(w, data)
	if err != nil { 
		log.Print(err.Error())
		http.Error(w, "Internal Server Error",http.StatusInternalServerError)   
		return 
   }
}

func week3Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/week3.html")
}

func week4Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/week4.html")
}

func week5Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/week5.html")
}

func main() {
	mux := http.NewServeMux()

	// Serve static files (CSS)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Serve image files from the "resource" folder
	resourceServer := http.FileServer(http.Dir("./resource"))
	mux.Handle("/resource/", http.StripPrefix("/resource/", resourceServer))

	// Register handlers for each blog post
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/week1", week1Handler)
	mux.HandleFunc("/week2", week2Handler)
	mux.HandleFunc("/week3", week3Handler)
	mux.HandleFunc("/week4", week4Handler)
	mux.HandleFunc("/week5", week5Handler)

	// Wrap the router with logging middleware
	loggedMux := loggingMiddleware(mux)

	// Start the server
	port := ":4000"
	log.Println("Starting server on", port)
	err := http.ListenAndServe(port, loggedMux)
	if err != nil {
		log.Fatal(err)
	}
}
