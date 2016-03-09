# Rest API for E3

### Set environment variable for GO and get revel
    mkdir <projectname>
    cd <projectname>
    export GOPATH=`pwd`
    export PATH="$PATH:$GOPATH/bin"
    go get github.com/revel/cmd/revel

### Get project
    git clone http://github.com/easydanie/e3RestAPI $GOPATH/src/e3RestAPI
  - Your project will be in $GOPATH/src/e3RestAPI

### Runserver
    PORT=<PORT> revel run e3RestAPI
