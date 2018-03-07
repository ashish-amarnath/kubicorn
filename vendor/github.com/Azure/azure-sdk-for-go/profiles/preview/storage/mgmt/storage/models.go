// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package storage

import original "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-06-01/storage"

type AccountsClient = original.AccountsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type AccessTier = original.AccessTier

const (
	Cool AccessTier = original.Cool
	Hot  AccessTier = original.Hot
)

type AccountStatus = original.AccountStatus

const (
	Available   AccountStatus = original.Available
	Unavailable AccountStatus = original.Unavailable
)

type Action = original.Action

const (
	Allow Action = original.Allow
)

type Bypass = original.Bypass

const (
	AzureServices Bypass = original.AzureServices
	Logging       Bypass = original.Logging
	Metrics       Bypass = original.Metrics
	None          Bypass = original.None
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type HTTPProtocol = original.HTTPProtocol

const (
	HTTPS     HTTPProtocol = original.HTTPS
	Httpshttp HTTPProtocol = original.Httpshttp
)

type KeyPermission = original.KeyPermission

const (
	Full KeyPermission = original.Full
	Read KeyPermission = original.Read
)

type KeySource = original.KeySource

const (
	MicrosoftKeyvault KeySource = original.MicrosoftKeyvault
	MicrosoftStorage  KeySource = original.MicrosoftStorage
)

type Kind = original.Kind

const (
	BlobStorage Kind = original.BlobStorage
	Storage     Kind = original.Storage
)

type Permissions = original.Permissions

const (
	A Permissions = original.A
	C Permissions = original.C
	D Permissions = original.D
	L Permissions = original.L
	P Permissions = original.P
	R Permissions = original.R
	U Permissions = original.U
	W Permissions = original.W
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	ResolvingDNS ProvisioningState = original.ResolvingDNS
	Succeeded    ProvisioningState = original.Succeeded
)

type Reason = original.Reason

const (
	AccountNameInvalid Reason = original.AccountNameInvalid
	AlreadyExists      Reason = original.AlreadyExists
)

type ReasonCode = original.ReasonCode

const (
	NotAvailableForSubscription ReasonCode = original.NotAvailableForSubscription
	QuotaID                     ReasonCode = original.QuotaID
)

type Services = original.Services

const (
	B Services = original.B
	F Services = original.F
	Q Services = original.Q
	T Services = original.T
)

type SignedResource = original.SignedResource

const (
	SignedResourceB SignedResource = original.SignedResourceB
	SignedResourceC SignedResource = original.SignedResourceC
	SignedResourceF SignedResource = original.SignedResourceF
	SignedResourceS SignedResource = original.SignedResourceS
)

type SignedResourceTypes = original.SignedResourceTypes

const (
	SignedResourceTypesC SignedResourceTypes = original.SignedResourceTypesC
	SignedResourceTypesO SignedResourceTypes = original.SignedResourceTypesO
	SignedResourceTypesS SignedResourceTypes = original.SignedResourceTypesS
)

type SkuName = original.SkuName

const (
	PremiumLRS    SkuName = original.PremiumLRS
	StandardGRS   SkuName = original.StandardGRS
	StandardLRS   SkuName = original.StandardLRS
	StandardRAGRS SkuName = original.StandardRAGRS
	StandardZRS   SkuName = original.StandardZRS
)

type SkuTier = original.SkuTier

const (
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type State = original.State

const (
	StateDeprovisioning       State = original.StateDeprovisioning
	StateFailed               State = original.StateFailed
	StateNetworkSourceDeleted State = original.StateNetworkSourceDeleted
	StateProvisioning         State = original.StateProvisioning
	StateSucceeded            State = original.StateSucceeded
)

type UsageUnit = original.UsageUnit

const (
	Bytes           UsageUnit = original.Bytes
	BytesPerSecond  UsageUnit = original.BytesPerSecond
	Count           UsageUnit = original.Count
	CountsPerSecond UsageUnit = original.CountsPerSecond
	Percent         UsageUnit = original.Percent
	Seconds         UsageUnit = original.Seconds
)

type Account = original.Account
type AccountCheckNameAvailabilityParameters = original.AccountCheckNameAvailabilityParameters
type AccountCreateParameters = original.AccountCreateParameters
type AccountKey = original.AccountKey
type AccountListKeysResult = original.AccountListKeysResult
type AccountListResult = original.AccountListResult
type AccountProperties = original.AccountProperties
type AccountPropertiesCreateParameters = original.AccountPropertiesCreateParameters
type AccountPropertiesUpdateParameters = original.AccountPropertiesUpdateParameters
type AccountRegenerateKeyParameters = original.AccountRegenerateKeyParameters
type AccountSasParameters = original.AccountSasParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type CustomDomain = original.CustomDomain
type Dimension = original.Dimension
type Encryption = original.Encryption
type EncryptionService = original.EncryptionService
type EncryptionServices = original.EncryptionServices
type Endpoints = original.Endpoints
type Identity = original.Identity
type IPRule = original.IPRule
type KeyVaultProperties = original.KeyVaultProperties
type ListAccountSasResponse = original.ListAccountSasResponse
type ListServiceSasResponse = original.ListServiceSasResponse
type MetricSpecification = original.MetricSpecification
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type Resource = original.Resource
type Restriction = original.Restriction
type ServiceSasParameters = original.ServiceSasParameters
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type SKUCapability = original.SKUCapability
type SkuListResult = original.SkuListResult
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type VirtualNetworkRule = original.VirtualNetworkRule
type OperationsClient = original.OperationsClient
type SkusClient = original.SkusClient
type UsageClient = original.UsageClient

func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsageClient(subscriptionID string) UsageClient {
	return original.NewUsageClient(subscriptionID)
}
func NewUsageClientWithBaseURI(baseURI string, subscriptionID string) UsageClient {
	return original.NewUsageClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
