# Rest API for E3

### Set environment variable for GO
    export GOPATH=`pwd`
    export PATH="$PATH:$GOPATH/bin"
    git clone http://github.com/easydanie/e3RestAPI $GOPATH/src/

### Get Revel
    go get github.com/revel/cmd/revel
  - Your project will be in $GOPATH/src/e3plus

### Runserver
    PORT=<PORT> revel run e3plus
