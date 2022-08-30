# Grepl
Grepl is a Grep-like command line utility built as a personal exercise in Go development. Implementation is based on [this](https://willdemaine.ghost.io/grep-from-first-principles-in-golang/) artcile by Will Demaine. You can read more about the motivation behind this project and my learnings [here](https://kumarpit.github.io/grepl).

# Getting Started
- Install Golang by following [these steps](https://go.dev/doc/install)
- Clone this repository into an empty directory using `git clone https://github.com/kumarpit/grepl.git`
- Install all dependencies using `go get ./..`
- Run the program using `go run . "hello" .`

# Usage
```golang
go run . <regexp> <path>
```
Grepl searches for `.txt` files recursively beginning at the specified path, after which it concurrently looks for matches.
After cloning this repository and running `go run . "hello" .`, you should see the following output:
```bash
pfiles/another.txt hello
words2.txt hello
words2.txt helloed 
words2.txt helloes 
words2.txt helloing
words2.txt hellos  
words.txt hello    
words.txt helloed  
words.txt helloes  
words.txt helloing 
words.txt hellos   
```

