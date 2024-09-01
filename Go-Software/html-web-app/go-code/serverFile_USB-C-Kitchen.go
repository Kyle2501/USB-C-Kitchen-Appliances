 # ~ Learn-CSS-Layout-online ~
// ,  ° . +
package main

import (


)

// ,  ° . +
type htmlPageData struct {
    pageTitle string
    pagePath string
    pageList []pageNav
    
}

type pageNav struct {
    pageTitle string
    pageLink string
}


// ,  ° . +
func app_welcome_center_page() {


}



// . appHandler, ~ for Data Pages °
func appHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,
  pageTitle := "~ Learn.CSS-Layout.online - // - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
// ,  ° . +
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  if pagePath == "/settings" {
      pageTitle = "Settings Page"
      pageList = pageList
  }
  
  
// ,  ° . + 
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler




// . pageHandler, ~ for User Pages °
func pageHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +

  pageTitle := "~ Learn.CSS-Layout.online - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  if pagePath == "/account" {
      pageTitle = "Account Page"
      pageList = pageList
  }
  
  if pagePath == "/profile" {
      pageTitle = "Profile Page"
      pageList = pageList
  }
  
  
// ,  ° . +
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }


// ,  ° . +
pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  pageHandler



// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ Learn.CSS-Layout.online - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  if pagePath == "/" {
      pageTitle = "Index Page"
      pageList = pageList
  }
  
    if pagePath == "/front" {
      pageTitle = "Front Page"
      pageList = pageList
  }
  
    if pagePath == "/main" {
      pageTitle = "Main Page"
      pageList = pageList
  }
  
    if pagePath == "/home" {
      pageTitle = "Home Page"
      pageList = pageList
  }
  
    if pagePath == "/start" {
      pageTitle = "Start Page"
      pageList = pageList
  }

// ,  ° . +
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler




//  .  html url routes 
//  .  as well as terminal cli logs

func main() {
// ,  ° . +
  appName := "~ Learn.CSS-Layout.online - // - Website App"

// ,  ° . +
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/front", indexHandler)
  http.HandleFunc("/main", indexHandler)
  http.HandleFunc("/home", indexHandler)
  http.HandleFunc("/start", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/user", pageHandler)
  http.HandleFunc("/account", pageHandler)
  http.HandleFunc("/profile", pageHandler)
  
  http.HandleFunc("/portfolio", pageHandler)
  http.HandleFunc("/resume", pageHandler)
  
  
  // ,  ° . +
  http.HandleFunc("/settings", appHandler)
  




// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
// ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }