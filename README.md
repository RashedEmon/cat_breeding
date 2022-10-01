# Cat Breeding

# Technologies used in this application.

1. Golang
2. Beego
3. Javascript
4. HTML
5. Bootstrap

# How to install and use this application

1. Install go https://go.dev/doc/install.

2. Create those directories inside your Users directory:
   /go
   /go/src
   /go/bin

3. Go to parent directory of go and Create mod file.
   <code>cd ..</code>
   <code>go mod init go</code>

4. Set GOPATH (User your own path):
   <code>go env -w GOPATH=C:\Users\<your_user>\go</code>

5. Set GOBIN (User your own path):
   <code>go env -w GOBIN=C:\Users\<your_user>\go\bin</code>

6. Install dependencies:
   <code>go get -u github.com/beego/beego/v2@latest</code>
   <code>go get -u github.com/beego/bee/v2@latest</code>
   <code>go get github.com/beego/beego/v2/server/web@v2.0.4</code>

7. Build Bee tools (User your own path):
   <code>cd C:\Users\<your_user>\go\pkg\mod\github.com\beego\bee\v2@v2.0.4</code>
   <code>go build</code>
   <code>copy bee.exe %GOPATH%\bin</code>

8. Set bee tools to environment variable.

9. Change directory to src and clone git repository
   <code>cd src</code>
   <code>git clone https://github.com/RashedEmon/cat_breeding.git</code>

10. Go to project root
    <code>cd cat_breeding</code>

11. Run the application.
    <code>bee run</code>
