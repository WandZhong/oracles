/* Package providings setup / initialization of common structures, drivers, ...
 * for other tools and applications
 */

package setup

import (
	"os"
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
)

// GitVersion should be substituted during build time by the git version. This is done
// using go linker flags:
// -ldflags "-X bitbucket.org/sweetbridge/oracles/go-lib/setup.GitVersion=$(git describe)
var GitVersion string

// envName is the application stage environment name (eg production, dev, backstage, ...).
// It is set using `SB_ENV` environment variable.
var envName string

var logger = log.Root()

// init initializes packages
func init() {
	if envName = os.Getenv("SB_ENV"); envName == "" {
		envName = "dev"
	} else {
		envName = strings.ToLower(envName)
		assert(checkAppName(envName))
	}
	if envName == "prod" {
		envName = "production"
	}
}
