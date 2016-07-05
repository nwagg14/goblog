build: blog.go webserver.go sql.go
	go build blog.go sql.go webserver.go

run: blog 
	sudo ./blog
