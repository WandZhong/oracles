package roles

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
)

// WithRoles provides interface for contracts which implement Roles contract.
type WithRoles interface {
	Owner(opts *bind.CallOpts) (common.Address, error)
	HasRole(opts *bind.CallOpts, roleName string) (bool, error)
	SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error)
}

// SenderHasRole checks if contract sender has specified role
func SenderHasRole(roleName string, ctr WithRoles) (bool, errstack.E) {
	hasRole, txErr := ctr.SenderHasRole(nil, roleName)
	return hasRole, errstack.WrapAsInfF(txErr, "checking role in %T", ctr)
	// return
}

// SenderIsOwnerOrHasRole checks if `who` is owner or sender has role
func SenderIsOwnerOrHasRole(who common.Address, roleName string, ctr WithRoles) (bool, errstack.E) {
	owner, txErr := ctr.Owner(nil)
	if txErr != nil {
		return false, errstack.WrapAsInfF(txErr, "reading owner from %T", ctr)
	}
	if owner == who {
		return true, nil
	}
	return SenderHasRole(roleName, ctr)
}

// SenderIsOwnerOrHasRoleErrp wraps SenderIsOwnerOrHasRole and assigns false response as
// an error to the putter.
func SenderIsOwnerOrHasRoleErrp(who common.Address, roleName string, ctr WithRoles, errp errstack.Putter) error {
	if ok, err := SenderIsOwnerOrHasRole(who, roleName, ctr); err != nil {
		return err
	} else if !ok {
		errp.Put(fmt.Sprintf("User [%s] is not authorised", who.Hex()))
	}
	return nil
}
