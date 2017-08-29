## 1. Install go
## https://github.com/golang/go/wiki/Ubuntu
sudo apt-get install make snapd
sudo snap install --classic go

if grep -q "GOPATH" ~/.profile; then
        echo "GOPATH already defined"
else
        echo "export GOPATH=\`go env GOPATH\`" >> ~/.profile
        echo "export PATH=\$GOPATH/bin:\$PATH" >> ~/.profile
fi

. ~/.profile

## 2. download
git clone git@bitbucket.org:sweetbridge/oracles.git $GOPATH/src/bitbucket.org/sweetbridge/oracles

## 3. Build jobs (shuld be called whenever jenkins tries to build the package)
## source profile if env is not defined:
# . ~/profile
cd $GOPATH/src/bitbucket.org/sweetbridge/oracles
git pull
make install-deps build

# the executables land in the project root directory (oracles).
