[//]: # (Image References)
[image_0]: ../misc/basic_parsing_01.PNG

# Golang+Docker - HTTP Handler
Act on HTTP header/body

## test_01
![form processing][image_0] 

### Use
- Server
    - run `main.go`
- POST request
    - [POSTMAN](https://www.getpostman.com/)
    - set `body` (see image example)
    - hit [send]


### basic form processing
```golang
// form performs basic form processing.
func form(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	if len(r.Form["name"]) > 0 {
		fmt.Fprintf(w, "Hi %v!", r.Form["name"][0])
	} else {
		fmt.Fprintln(w, "Why don't you have a name?")
	}
}
```

### Image
```
REPOSITORY                 TAG                 IMAGE ID            CREATED              SIZE
jackburdick/body           latest              473c2de2f365        About a minute ago   9.86MB
```

#### Resources
- [processing-form-request-data-in-golang](https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be)
