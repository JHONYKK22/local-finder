## **Local finder**   
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)

This is a small program that uses ***goroutines*** and ***<span style="color:#00ADD8">Go</span> channels*** to traverse directories concurrently and search for ***files/folders*** whose names match or are similar to the specified term. 

>- It works for Windows, Linux and Mac.

<br><br>


To see the options, run:
```sh
go run .\main.go -help
```
Or
```sh
./local-finder-build --help
```
<br><br>

---



If you don't specify the path, default path is current directory

```sh
./local-finder-build --find "name" 
```
Or
```sh
go run .\main.go -find "name" --path "initial_path"
```
Or
```sh
./local-finder-build --find "name" -path "initial_path"
```
<br><br>

---

If you want to keep the results in a file, run

```sh
go run .\main.go -find="name" -path "initial_path" > result_file_name.txt
```
or

``` sh
./local-finder-build --find "name" --path="initial_path" > result_file_name.txt
```
<br><br>

---

If you want create all the builds for Linux, mac and windows, run:

``` sh
./builder.sh
```