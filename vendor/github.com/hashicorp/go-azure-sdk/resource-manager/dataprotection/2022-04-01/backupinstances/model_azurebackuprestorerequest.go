package backupinstances

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBackupRestoreRequest interface {
}

func unmarshalAzureBackupRestoreRequestImplementation(input []byte) (AzureBackupRestoreRequest, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureBackupRestoreRequest into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureBackupRecoveryPointBasedRestoreRequest") {
		var out AzureBackupRecoveryPointBasedRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupRecoveryPointBasedRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureBackupRecoveryTimeBasedRestoreRequest") {
		var out AzureBackupRecoveryTimeBasedRestoreRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupRecoveryTimeBasedRestoreRequest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureBackupRestoreWithRehydrationRequest") {
		var out AzureBackupRestoreWithRehydrationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureBackupRestoreWithRehydrationRequest: %+v", err)
		}
		return out, nil
	}

	type RawAzureBackupRestoreRequestImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawAzureBackupRestoreRequestImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
