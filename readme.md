## Gin framework
Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
### Installation 
First of all clone the project using either ssh or https
```shell 
git clone git@github.com:Ardouz11/web-service-gin.git
```    
after run 
```shell 
go get .
``` 
to download gin framework 
### Running 
Run 
``` 
go run .
``` 
to run the project. after this u can check in your browser this url `http://localhost:8081/albums` to get the albums or via your terminal using 
```shell 
curl -i -X GET http://localhost:8081/albums
```
To add new album use this command 
```shell 
curl http://localhost:8081/albums -i  -H "Content-Type: application/json" -d '{"id": "6","title": "Excelsior","artist": "Lmorphine","price": 36}'
```
To get an album by its ID you can either check on your browser using this url `http//localhost:8081/albums/id of the wanted album` or using the terminal like 
```shell 
curl http://localhost:8081/albums/3 -i  -H "Content-Type: application/json" -X GET
```
