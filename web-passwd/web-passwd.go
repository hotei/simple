// web-passwd.go - borrowed from ?

// working

package main

// BUG(mdr): is a return needed after a redirect? apparently yes.

import (
	"fmt"
	"log"
	"net/http"
	//
	"github.com/hotei/simple/htpasswd"
)

const (
	passwdURL = "/passwd/"
	forumURL  = "/forum/"
)

var forumPassword string

func init() {
	// admin:{SHA}MCdMR5A70brHYzu/CXQxSeurgF8=  SHA1base64("passwd")
	forumPassword = "MCdMR5A70brHYzu/CXQxSeurgF8="
}

func handleMain(w http.ResponseWriter, rq *http.Request) {

	askPass := `
<html>
<head>
<title>Welcome to my web</title>
</head>
<body>
(splash)<br>
<hr>
<form action="." method="POST">
Enter your passwd: <input type="password" name="body" size=20 maxlength=20></input>
<input type="submit" name="pick" value="Go">
</form>
</body>
</html>
`

	// see if we already have good password in hand
	x, err := rq.Cookie("auth")
	fmt.Printf("handleMain: x[%v] err[%v]\n", x, err)
	if err == nil && x.Value == forumPassword {
		fmt.Printf("Redirecting to %s\n", forumURL)
		http.Redirect(w, rq, forumURL, 302)
		return
	}
	// if no cookie or wrong password in cookie send user back to password form

	err = rq.ParseForm()
	if err != nil {
		log.Printf("parseForm() err\n")
		return
	}
	if rq.Method == "POST" {
		buttonVal := rq.Form["pick"]
		log.Printf("Form value = %v\n", rq.Form)
		if len(buttonVal) > 0 && buttonVal[0] == "Go" { // anybutton == "SetPass"???
			userPasswd := rq.Form["body"][0]
			fmt.Printf("User provided password : %q\n", userPasswd)
			var cookie http.Cookie
			cookie.Name = "auth"
			cookie.Value, err = htpasswd.SHA1base64(userPasswd)
			if err != nil {
				fmt.Printf("%v\n", err)
				// BUG(mdr): advise user of error
				http.Redirect(w, rq, passwdURL, 302)
				return
			}
			if cookie.Value != forumPassword {
				w.Write([]byte(askPass))
				return
			}
			fmt.Printf("Storing cookie auth:%s\n", cookie.Value)
			http.SetCookie(w, &cookie)
			//w.Write([]byte("whatever")) // required to actually send cookie? but also generates multiple write warning
			// we have (maybe) a good password, so try it out by looping back to self
			http.Redirect(w, rq, forumURL, 302)
			return
		}
		// can't happen?
		log.Panic()
	}
	w.Write([]byte(askPass)) // sometimes doesn't cause expected text to appear... why?
}

func DumpCookie(c *http.Cookie) {
	fmt.Printf("Name = %q\n", c.Name)
	fmt.Printf("Value = %q\n", c.Value)
	fmt.Printf("Path = %q\n", c.Path)
	fmt.Printf("Raw = %q\n", c.Raw)
}

func handleForum(w http.ResponseWriter, rq *http.Request) {
	proceed := `<html>
<head>
</head>
<body>
Test forum data is here<br>
</body>
</html>`
	fmt.Printf("entered forum handler\n")

	// this prolog has to be at front of every handler
	x, err := rq.Cookie("auth")
	if err != nil || x.Value != forumPassword {
		fmt.Printf("Redirecting to %s\n", passwdURL)
		http.Redirect(w, rq, passwdURL, 302)
		return
	}
	// end of prolog

	fmt.Printf("authorized viewing now\n")
	w.Write([]byte(proceed))
}

func main() {
	log.SetFlags( /*log.LstdFlags | */ log.Lshortfile)
	http.HandleFunc(forumURL, handleForum)
	http.HandleFunc(passwdURL, handleMain)
	http.HandleFunc("/", handleMain)
	fmt.Printf("Starting http server at %d\n", 12345)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Printf("error running docs webserver: %v", err)
	}
}
