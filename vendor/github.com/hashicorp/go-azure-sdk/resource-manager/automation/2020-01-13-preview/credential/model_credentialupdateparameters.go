package credential

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUpdateParameters struct {
	Name       *string                     `json:"name,omitempty"`
	Properties *CredentialUpdateProperties `json:"properties,omitempty"`
}
