# Reduce Size of CurrentTime App

### The app
Start a server and display the current time to the user when they visit the source page.
```golang
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	curTime := time.Now().Format("02.01.2006 15:04:05")

	fmt.Fprintf(w, "%s", curTime)
}
```

#### Results
Method | Size (MB) | Directory | Dockerfile
:--- | :---: | :---: | :---: |
standard | **704** | [link](./1_standard) | [link](./1_standard/Dockerfile) |
alpine | **263** | [link](./2_alpine) | [link](./2_alpine/Dockerfile) |
multi-stage | **9.85** | [link](./3_multiStage) | [link](./3_multiStage/Dockerfile) |