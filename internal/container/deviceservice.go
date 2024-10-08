// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

	"github.com/edgexfoundry/device-sdk-go/v3/pkg/interfaces"
)

// DeviceServiceName contains the name of device service struct in the DIC.
var DeviceServiceName = di.TypeInstanceToName(models.DeviceService{})

// ProtocolDriverName contains the name of protocol driver implementation in the DIC.
var ProtocolDriverName = di.TypeInstanceToName((*interfaces.ProtocolDriver)(nil))

// AutoEventManagerName contains the name of autoevent manager implementation in the DIC
var AutoEventManagerName = di.TypeInstanceToName((*interfaces.AutoEventManager)(nil))

// ProfileScanName contains the name of profile scan implementation in the DIC.
var ProfileScanName = di.TypeInstanceToName((*interfaces.ProfileScan)(nil))

// DeviceServiceFrom helper function queries the DIC and returns device service struct.
func DeviceServiceFrom(get di.Get) *models.DeviceService {
	return get(DeviceServiceName).(*models.DeviceService)
}

// ProtocolDriverFrom helper function queries the DIC and returns protocol driver implementation.
func ProtocolDriverFrom(get di.Get) interfaces.ProtocolDriver {
	return get(ProtocolDriverName).(interfaces.ProtocolDriver)
}

// AutoEventManagerFrom helper function queries the DIC and returns autoevent manager implementation
func AutoEventManagerFrom(get di.Get) interfaces.AutoEventManager {
	return get(AutoEventManagerName).(interfaces.AutoEventManager)
}

// ProfileScanFrom helper function queries the DIC and returns profile scan implementation.
func ProfileScanFrom(get di.Get) interfaces.ProfileScan {
	casted, ok := get(ProfileScanName).(interfaces.ProfileScan)
	if ok {
		return casted
	}
	return nil
}

// DiscoveryRequestIdName contains the name of discovery request id implementation in the DIC.
var DiscoveryRequestIdName = di.TypeInstanceToName(new(string))

// DiscoveryRequestIdFrom helper function queries the DIC and returns discovery request id.
func DiscoveryRequestIdFrom(get di.Get) string {
	return get(DiscoveryRequestIdName).(string)
}
