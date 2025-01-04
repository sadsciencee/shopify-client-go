package types

// MetafieldAdminAccess Possible admin access settings for metafields.
type MetafieldAdminAccess string

const (
	// MetafieldAdminAccessPrivate Owner gets full access. No one else has access rights.
	MetafieldAdminAccessPrivate MetafieldAdminAccess = "PRIVATE"
	// MetafieldAdminAccessPublicRead Owner gets full access. All applications and the merchant have read-only access.
	MetafieldAdminAccessPublicRead MetafieldAdminAccess = "PUBLIC_READ"
	// MetafieldAdminAccessPublicReadWrite Owner gets full access. All applications and the merchant have read and write access.
	MetafieldAdminAccessPublicReadWrite MetafieldAdminAccess = "PUBLIC_READ_WRITE"
	// MetafieldAdminAccessMerchantRead Owner gets full access. The merchant has read-only access. No one else has access rights.
	MetafieldAdminAccessMerchantRead MetafieldAdminAccess = "MERCHANT_READ"
	// MetafieldAdminAccessMerchantReadWrite Owner gets full access. The merchant has read and write access. No one else has access rights.
	MetafieldAdminAccessMerchantReadWrite MetafieldAdminAccess = "MERCHANT_READ_WRITE"
)

func (e MetafieldAdminAccess) ToInput() *MetafieldAdminAccessInput {
	var result MetafieldAdminAccessInput
	switch e {
	case MetafieldAdminAccessPrivate:
		result = MetafieldAdminAccessInputMerchantRead
	case MetafieldAdminAccessPublicRead:
		result = MetafieldAdminAccessInputMerchantReadWrite
	case MetafieldAdminAccessPublicReadWrite:
		result = MetafieldAdminAccessInputMerchantReadWrite
	case MetafieldAdminAccessMerchantRead:
		result = MetafieldAdminAccessInputMerchantRead
	case MetafieldAdminAccessMerchantReadWrite:
		result = MetafieldAdminAccessInputMerchantReadWrite
	}
	return &result
}

// MetafieldAdminAccessInput The possible values for setting metafield Admin API access.
type MetafieldAdminAccessInput string

const (
	// MetafieldAdminAccessInputMerchantRead Owner gets full access. The merchant has read-only access. No one else has access rights.
	MetafieldAdminAccessInputMerchantRead MetafieldAdminAccessInput = "MERCHANT_READ"
	// MetafieldAdminAccessInputMerchantReadWrite Owner gets full access. The merchant has read and write access. No one else has access rights.
	MetafieldAdminAccessInputMerchantReadWrite MetafieldAdminAccessInput = "MERCHANT_READ_WRITE"
)

func (e MetafieldAdminAccessInput) ToValue() *MetafieldAdminAccess {
	var result MetafieldAdminAccess
	switch e {
	case MetafieldAdminAccessInputMerchantRead:
		result = MetafieldAdminAccessMerchantRead
	case MetafieldAdminAccessInputMerchantReadWrite:
		result = MetafieldAdminAccessMerchantReadWrite
	}
	return &result
}

// MetafieldCustomerAccountAccess Defines how the metafields of a definition can be accessed in the Customer Account API.
type MetafieldCustomerAccountAccess string

const (
	// MetafieldCustomerAccountAccessReadWrite The Customer Account API can read and write metafields.
	MetafieldCustomerAccountAccessReadWrite MetafieldCustomerAccountAccess = "READ_WRITE"
	// MetafieldCustomerAccountAccessRead The Customer Account API can read metafields.
	MetafieldCustomerAccountAccessRead MetafieldCustomerAccountAccess = "READ"
	// MetafieldCustomerAccountAccessNone The Customer Account API cannot access metafields.
	MetafieldCustomerAccountAccessNone MetafieldCustomerAccountAccess = "NONE"
)

func (e MetafieldCustomerAccountAccess) ToInput() *MetafieldCustomerAccountAccessInput {
	var result MetafieldCustomerAccountAccessInput
	switch e {
	case MetafieldCustomerAccountAccessReadWrite:
		result = MetafieldCustomerAccountAccessInputReadWrite
	case MetafieldCustomerAccountAccessRead:
		result = MetafieldCustomerAccountAccessInputRead
	case MetafieldCustomerAccountAccessNone:
		result = MetafieldCustomerAccountAccessInputNone
	}
	return &result
}

// MetafieldCustomerAccountAccessInput The possible values for setting metafield Customer Account API access.
type MetafieldCustomerAccountAccessInput string

const (
	// MetafieldCustomerAccountAccessInputReadWrite The Customer Account API can read and write metafields.
	MetafieldCustomerAccountAccessInputReadWrite MetafieldCustomerAccountAccessInput = "READ_WRITE"
	// MetafieldCustomerAccountAccessInputRead The Customer Account API can read metafields.
	MetafieldCustomerAccountAccessInputRead MetafieldCustomerAccountAccessInput = "READ"
	// MetafieldCustomerAccountAccessInputNone The Customer Account API cannot access metafields.
	MetafieldCustomerAccountAccessInputNone MetafieldCustomerAccountAccessInput = "NONE"
)

func (e MetafieldCustomerAccountAccessInput) ToValue() *MetafieldCustomerAccountAccess {
	var result MetafieldCustomerAccountAccess
	switch e {
	case MetafieldCustomerAccountAccessInputReadWrite:
		result = MetafieldCustomerAccountAccessReadWrite
	case MetafieldCustomerAccountAccessInputRead:
		result = MetafieldCustomerAccountAccessRead
	case MetafieldCustomerAccountAccessInputNone:
		result = MetafieldCustomerAccountAccessNone

	}
	return &result
}

// MetafieldStorefrontAccess Defines how the metafields of a definition can be accessed in Storefront API surface areas, including Liquid and the GraphQL Storefront API.
type MetafieldStorefrontAccess string

const (
	// MetafieldStorefrontAccessPublicRead Metafields are accessible in the GraphQL Storefront API and online store Liquid templates.
	MetafieldStorefrontAccessPublicRead MetafieldStorefrontAccess = "PUBLIC_READ"
	// MetafieldStorefrontAccessNone Metafields are not accessible in any Storefront API surface area.
	MetafieldStorefrontAccessNone MetafieldStorefrontAccess = "NONE"
)

func (e MetafieldStorefrontAccess) ToInput() *MetafieldStorefrontAccessInput {
	var result MetafieldStorefrontAccessInput
	switch e {
	case MetafieldStorefrontAccessPublicRead:
		result = MetafieldStorefrontAccessInputPublicRead
	case MetafieldStorefrontAccessNone:
		result = MetafieldStorefrontAccessInputNone
	}
	return &result
}

// MetafieldStorefrontAccessInput The possible values for setting metafield storefront access. Storefront accesss governs both Liquid and the GraphQL Storefront API.
type MetafieldStorefrontAccessInput string

const (
	// MetafieldStorefrontAccessInputPublicRead Metafields are accessible in the GraphQL Storefront API and online store Liquid templates.
	MetafieldStorefrontAccessInputPublicRead MetafieldStorefrontAccessInput = "PUBLIC_READ"
	// MetafieldStorefrontAccessInputNone Metafields are not accessible in any Storefront API surface area.
	MetafieldStorefrontAccessInputNone MetafieldStorefrontAccessInput = "NONE"
)

func (e MetafieldStorefrontAccessInput) ToValue() *MetafieldStorefrontAccess {
	var result MetafieldStorefrontAccess
	switch e {
	case MetafieldStorefrontAccessInputPublicRead:
		result = MetafieldStorefrontAccessPublicRead
	case MetafieldStorefrontAccessInputNone:
		result = MetafieldStorefrontAccessNone
	}
	return &result
}

// MetafieldDefinitionValidationStatus Possible metafield definition validation statuses.
type MetafieldDefinitionValidationStatus string

const (
	// MetafieldDefinitionValidationStatusAllValid All of this definition's metafields are valid.
	MetafieldDefinitionValidationStatusAllValid MetafieldDefinitionValidationStatus = "ALL_VALID"
	// MetafieldDefinitionValidationStatusInProgress Asynchronous validation of this definition's metafields is in progress.
	MetafieldDefinitionValidationStatusInProgress MetafieldDefinitionValidationStatus = "IN_PROGRESS"
	// MetafieldDefinitionValidationStatusSomeInvalid Some of this definition's metafields are invalid.
	MetafieldDefinitionValidationStatusSomeInvalid MetafieldDefinitionValidationStatus = "SOME_INVALID"
)

// MetafieldValueType Legacy type information for the stored value. Replaced by `type`.
type MetafieldValueType string

const (
	// MetafieldValueTypeString A text field.
	MetafieldValueTypeString MetafieldValueType = "STRING"
	// MetafieldValueTypeInteger A whole number.
	MetafieldValueTypeInteger MetafieldValueType = "INTEGER"
	// MetafieldValueTypeJsonString A JSON string.
	MetafieldValueTypeJsonString MetafieldValueType = "JSON_STRING"
	// MetafieldValueTypeBoolean A `true` or `false` value.
	MetafieldValueTypeBoolean MetafieldValueType = "BOOLEAN"
)

func (e MetafieldValueType) ToInput() *string {
	result := string(e)
	return &result
}
