# Rest API for E3

### Set environment variable for GO and get revel
    export GOPATH=`pwd`
    export PATH="$PATH:$GOPATH/bin"
    go get github.com/revel/cmd/revel

### Get project
    git clone http://github.com/easydanie/e3RestAPI $GOPATH/src/
  - Your project will be in $GOPATH/src/e3plus

### Runserver
    PORT=<PORT> revel run e3plus
