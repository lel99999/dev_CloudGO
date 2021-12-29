# dev_CloudGO
Rethinking the building of cloud native ecosystems
_ _ _

To view GO documentation locally in a web browser:

>Type the following:  $godoc -http=: 8999
>
>Open up a browser and go to the URL: http://localhost:8999

### Install Git, Mercurial and Bazaar
$brew install git
$brew install mercurial
$brew install bazaar

### Check on installed
$brew info git
$brew info mercurial
$brew info bazaar

### Install GO on Mac OSX with Homebrew
$brew install go

#### Set up GO workspace - set up GOPATH
export GOPATH="${HOME}/.go"

#### Set up GOROOT
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"

#### Get golint a static analysis tool
$go get -u github.com/golang/lint/golint
