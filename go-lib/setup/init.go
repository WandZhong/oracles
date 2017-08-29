/* Package providings setup / initialization of common structures, drivers, ...
 * for other tools and applications
 */

package setup

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/robert-zaremba/log15/rollbar"
)

// GitVersion should be substituted during build time by the git version. This is done
// using go linker flags:
// -ldflags "-X bitbucket.org/sweetbridge/oracles/go-lib/setup.GitVersion=$(git describe)
var GitVersion string

var logger log.Logger

// Init initializes packages
func Init() {
	var err error
	logger, err = log.New(log.RootName,
		log.Config{Color: true, TimeFmt: "sec", Level: "DEBUG"}, rollbar.Config{})
	if err != nil {
		panic(err)
	}
}
