// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Package audit records auditable events
package audit

import (
	"net"
	"time"

	"github.com/juju/errors"
	"github.com/juju/utils"
	"github.com/juju/version"
)

// AuditEntry represents an auditted event.
type AuditEntry struct {
	// JujuServerVersion is the version of the jujud that recorded
	// this AuditEntry.
	JujuServerVersion version.Number
	// ModelUUID is the ID of the model the audit entry was written
	// on.
	ModelUUID string
	// Timestamp is when the audit entry was generated. It must be
	// stored with the UTC locale.
	Timestamp time.Time
	// OriginIP is the IP of the machine from which the audit-event
	// was triggered.
	OriginIP net.IP
	// OriginType is the type of entity (e.g. model, user, action)
	// which triggered the audit event.
	OriginType string
	// OriginName is the name of the origin which triggered the
	// audit-event.
	OriginName string
	// Operation is the operation that was performed that triggered
	// the audit event.
	Operation string
	// Data is a catch-all for storing random data.
	Data map[string]interface{}
}

// Validate ensures that the entry considers itself to be in a
// complete and valid state.
func (e AuditEntry) Validate() error {
	if e.JujuServerVersion == version.Zero {
		return errors.NewNotValid(errors.NotAssignedf("JujuServerVersion"), "")
	}
	if e.ModelUUID == "" {
		return errors.NewNotValid(errors.NotAssignedf("ModelUUID"), "")
	}
	if utils.IsValidUUIDString(e.ModelUUID) == false {
		return errors.NotValidf("ModelUUID")
	}
	if e.Timestamp.IsZero() {
		return errors.NewNotValid(errors.NotAssignedf("Timestamp"), "")
	}
	if e.Timestamp.Location() != time.UTC {
		return errors.NewNotValid(errors.NotValidf("Timestamp"), "must be set to UTC")
	}
	if e.OriginIP == nil || e.OriginIP.IsUnspecified() {
		return errors.NewNotValid(errors.NotAssignedf("OriginIP"), "")
	}
	if e.OriginType == "" {
		return errors.NewNotValid(errors.NotAssignedf("OriginType"), "")
	}
	if e.OriginName == "" {
		return errors.NewNotValid(errors.NotAssignedf("OriginName"), "")
	}
	if e.Operation == "" {
		return errors.NewNotValid(errors.NotAssignedf("Operation"), "")
	}

	// Data remains unchecked as it is always optional.

	return nil
}
