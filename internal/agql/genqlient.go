// Sample generated code for test lib

package agql

import (
	"context"
	"github.com/sadsciencee/shopify-client-go/types"

	"github.com/Khan/genqlient/graphql"
)

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload includes the requested fields of the GraphQL type MetafieldDefinitionCreatePayload.
// The GraphQL type's documentation follows.
//
// Return type for `metafieldDefinitionCreate` mutation.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload struct {
	// The metafield definition that was created.
	CreatedDefinition *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition `json:"createdDefinition"`
	// The list of errors that occurred from executing the mutation.
	UserErrors []*CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError `json:"userErrors"`
}

// GetCreatedDefinition returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload.CreatedDefinition, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload) GetCreatedDefinition() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition {
	return v.CreatedDefinition
}

// GetUserErrors returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload.UserErrors, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload) GetUserErrors() []*CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError {
	return v.UserErrors
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition includes the requested fields of the GraphQL type MetafieldDefinition.
// The GraphQL type's documentation follows.
//
// Metafield definitions enable you to define additional validation constraints for metafields, and enable the merchant to edit metafield values in context.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition struct {
	// A globally-unique ID.
	Id string `json:"id"`
	// The container for a group of metafields that the metafield definition is associated with.
	Namespace string `json:"namespace"`
	// The unique identifier for the metafield definition within its namespace.
	Key string `json:"key"`
	// The human-readable name of the metafield definition.
	Name string `json:"name"`
	// The description of the metafield definition.
	Description *string `json:"description"`
	// The resource type that the metafield definition is attached to.
	OwnerType MetafieldOwnerType `json:"ownerType"`
	// The type of data that each of the metafields that belong to the metafield definition will store. Refer to the list of [supported types](https://shopify.dev/apps/metafields/types).
	Type *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType `json:"type"`
	// The validation status for the metafields that belong to the metafield definition.
	ValidationStatus types.MetafieldDefinitionValidationStatus `json:"validationStatus"`
	// The position of the metafield definition in the pinned list.
	PinnedPosition *int `json:"pinnedPosition"`
	// Whether the metafield definition can be used as a collection condition.
	UseAsCollectionCondition bool `json:"useAsCollectionCondition"`
	// The count of the metafields that belong to the metafield definition.
	MetafieldsCount int `json:"metafieldsCount"`
	// The access settings associated with the metafield definition.
	Access *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess `json:"access"`
	// The capabilities of the metafield definition.
	Capabilities *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities `json:"capabilities"`
	// A list of [validation options](https://shopify.dev/apps/metafields/definitions/validation) for the metafields that belong to the metafield definition. For example, for a metafield definition with the type `date`, you can set a minimum date validation so that each of the metafields that belong to it can only store dates after the specified minimum.
	Validations []*CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation `json:"validations"`
}

// GetId returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Id, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetId() string {
	return v.Id
}

// GetNamespace returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Namespace, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetNamespace() string {
	return v.Namespace
}

// GetKey returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Key, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetKey() string {
	return v.Key
}

// GetName returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Name, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetName() string {
	return v.Name
}

// GetDescription returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Description, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetDescription() *string {
	return v.Description
}

// GetOwnerType returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.OwnerType, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetOwnerType() MetafieldOwnerType {
	return v.OwnerType
}

// GetType returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Type, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetType() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType {
	return v.Type
}

// GetValidationStatus returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.ValidationStatus, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetValidationStatus() types.MetafieldDefinitionValidationStatus {
	return v.ValidationStatus
}

// GetPinnedPosition returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.PinnedPosition, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetPinnedPosition() *int {
	return v.PinnedPosition
}

// GetUseAsCollectionCondition returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.UseAsCollectionCondition, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetUseAsCollectionCondition() bool {
	return v.UseAsCollectionCondition
}

// GetMetafieldsCount returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.MetafieldsCount, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetMetafieldsCount() int {
	return v.MetafieldsCount
}

// GetAccess returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Access, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetAccess() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess {
	return v.Access
}

// GetCapabilities returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Capabilities, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetCapabilities() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities {
	return v.Capabilities
}

// GetValidations returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition.Validations, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinition) GetValidations() []*CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation {
	return v.Validations
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess includes the requested fields of the GraphQL type MetafieldAccess.
// The GraphQL type's documentation follows.
//
// The access settings for this metafield definition.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess struct {
	// The default admin access setting used for the metafields under this definition.
	Admin *types.MetafieldAdminAccess `json:"admin"`
	// The storefront access setting used for the metafields under this definition.
	Storefront *types.MetafieldStorefrontAccess `json:"storefront"`
}

// GetAdmin returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess.Admin, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess) GetAdmin() *types.MetafieldAdminAccess {
	return v.Admin
}

// GetStorefront returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess.Storefront, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionAccessMetafieldAccess) GetStorefront() *types.MetafieldStorefrontAccess {
	return v.Storefront
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities includes the requested fields of the GraphQL type MetafieldCapabilities.
// The GraphQL type's documentation follows.
//
// Provides the capabilities of a metafield definition.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities struct {
	// Indicate whether a metafield definition can be used as a smart collection condition.
	SmartCollectionCondition *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition `json:"smartCollectionCondition"`
	// Indicate whether a metafield definition is configured for filtering.
	AdminFilterable *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable `json:"adminFilterable"`
}

// GetSmartCollectionCondition returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities.SmartCollectionCondition, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities) GetSmartCollectionCondition() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition {
	return v.SmartCollectionCondition
}

// GetAdminFilterable returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities.AdminFilterable, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilities) GetAdminFilterable() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable {
	return v.AdminFilterable
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable includes the requested fields of the GraphQL type MetafieldCapabilityAdminFilterable.
// The GraphQL type's documentation follows.
//
// Information about the admin filterable capability on a metafield definition.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable struct {
	// Indicates if the capability is enabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable.Enabled, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable) GetEnabled() bool {
	return v.Enabled
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition includes the requested fields of the GraphQL type MetafieldCapabilitySmartCollectionCondition.
// The GraphQL type's documentation follows.
//
// Information about the smart collection condition capability on a metafield definition.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition struct {
	// Indicates if the capability is enabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition.Enabled, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition) GetEnabled() bool {
	return v.Enabled
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType includes the requested fields of the GraphQL type MetafieldDefinitionType.
// The GraphQL type's documentation follows.
//
// A metafield definition type provides basic foundation and validation for a metafield.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType struct {
	// The name of the type for the metafield definition. See the list of [supported types](https://shopify.dev/apps/metafields/types).
	Name string `json:"name"`
	// The category associated with the metafield definition type.
	Category string `json:"category"`
	// Whether metafields without a definition can be migrated to a definition of this type.
	SupportsDefinitionMigrations bool `json:"supportsDefinitionMigrations"`
	// The supported validations for a metafield definition type.
	SupportedValidations []*CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation `json:"supportedValidations"`
}

// GetName returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType.Name, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType) GetName() string {
	return v.Name
}

// GetCategory returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType.Category, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType) GetCategory() string {
	return v.Category
}

// GetSupportsDefinitionMigrations returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType.SupportsDefinitionMigrations, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType) GetSupportsDefinitionMigrations() bool {
	return v.SupportsDefinitionMigrations
}

// GetSupportedValidations returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType.SupportedValidations, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionType) GetSupportedValidations() []*CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation {
	return v.SupportedValidations
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation includes the requested fields of the GraphQL type MetafieldDefinitionSupportedValidation.
// The GraphQL type's documentation follows.
//
// The type and name for the optional validation configuration of a metafield.  For example, a supported validation might consist of a `max` name and a `number_integer` type. This validation can then be used to enforce a maximum character length for a `single_line_text_field` metafield.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation struct {
	// The name of the metafield definition validation.
	Name string `json:"name"`
	// The type of input for the validation.
	Type string `json:"type"`
}

// GetName returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation.Name, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation) GetName() string {
	return v.Name
}

// GetType returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation.Type, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation) GetType() string {
	return v.Type
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation includes the requested fields of the GraphQL type MetafieldDefinitionValidation.
// The GraphQL type's documentation follows.
//
// A configured metafield definition validation.  For example, for a metafield definition of `number_integer` type, you can set a validation with the name `max` and a value of `15`. This validation will ensure that the value of the metafield is a number less than or equal to 15.  Refer to the [list of supported validations](https://shopify.dev/api/admin/graphql/reference/common-objects/metafieldDefinitionTypes#examples-Fetch_all_metafield_definition_types).
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation struct {
	// The validation name.
	Name string `json:"name"`
	// The validation value.
	Value *string `json:"value"`
}

// GetName returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation.Name, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation) GetName() string {
	return v.Name
}

// GetValue returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation.Value, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadCreatedDefinitionMetafieldDefinitionValidationsMetafieldDefinitionValidation) GetValue() *string {
	return v.Value
}

// CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError includes the requested fields of the GraphQL type MetafieldDefinitionCreateUserError.
// The GraphQL type's documentation follows.
//
// An error that occurs during the execution of `MetafieldDefinitionCreate`.
type CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError struct {
	// The path to the input field that caused the error.
	Field []string `json:"field"`
	// The error message.
	Message string `json:"message"`
	// The error code.
	Code *MetafieldDefinitionCreateUserErrorCode `json:"code"`
}

// GetField returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError.Field, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError) GetField() []string {
	return v.Field
}

// GetMessage returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError.Message, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError) GetMessage() string {
	return v.Message
}

// GetCode returns CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError.Code, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayloadUserErrorsMetafieldDefinitionCreateUserError) GetCode() *MetafieldDefinitionCreateUserErrorCode {
	return v.Code
}

// CreateMetafieldDefinitionResponse is returned by CreateMetafieldDefinition on success.
type CreateMetafieldDefinitionResponse struct {
	// Creates a metafield definition. Any metafields existing under the same owner type, namespace, and key will be checked against this definition and will have their type updated accordingly. For metafields that are not valid, they will remain unchanged but any attempts to update them must align with this definition.
	MetafieldDefinitionCreate *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload `json:"metafieldDefinitionCreate"`
}

// GetMetafieldDefinitionCreate returns CreateMetafieldDefinitionResponse.MetafieldDefinitionCreate, and is useful for accessing the field via an interface.
func (v *CreateMetafieldDefinitionResponse) GetMetafieldDefinitionCreate() *CreateMetafieldDefinitionMetafieldDefinitionCreateMetafieldDefinitionCreatePayload {
	return v.MetafieldDefinitionCreate
}

// CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload includes the requested fields of the GraphQL type MetaobjectDefinitionCreatePayload.
// The GraphQL type's documentation follows.
//
// Return type for `metaobjectDefinitionCreate` mutation.
type CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload struct {
	// The created metaobject definition.
	MetaobjectDefinition *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition `json:"metaobjectDefinition"`
	// The list of errors that occurred from executing the mutation.
	UserErrors []*CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError `json:"userErrors"`
}

// GetMetaobjectDefinition returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload.MetaobjectDefinition, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload) GetMetaobjectDefinition() *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition {
	return v.MetaobjectDefinition
}

// GetUserErrors returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload.UserErrors, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload) GetUserErrors() []*CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError {
	return v.UserErrors
}

// CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition includes the requested fields of the GraphQL type MetaobjectDefinition.
// The GraphQL type's documentation follows.
//
// Provides the definition of a generic object structure composed of metafields.
type CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition struct {
	// A globally-unique ID.
	Id string `json:"id"`
	// The human-readable name.
	Name string `json:"name"`
	// The type of the object definition. Defines the namespace of associated metafields.
	Type string `json:"type"`
	// The fields defined for this object type.
	FieldDefinitions []*CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition `json:"fieldDefinitions"`
}

// GetId returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition.Id, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition) GetId() string {
	return v.Id
}

// GetName returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition.Name, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition) GetName() string {
	return v.Name
}

// GetType returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition.Type, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition) GetType() string {
	return v.Type
}

// GetFieldDefinitions returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition.FieldDefinitions, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinition) GetFieldDefinitions() []*CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition {
	return v.FieldDefinitions
}

// CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition includes the requested fields of the GraphQL type MetaobjectFieldDefinition.
// The GraphQL type's documentation follows.
//
// Defines a field for a MetaobjectDefinition with properties such as the field's data type and validations.
type CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition struct {
	// The human-readable name.
	Name string `json:"name"`
	// A key name used to identify the field within the metaobject composition.
	Key string `json:"key"`
}

// GetName returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition.Name, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition) GetName() string {
	return v.Name
}

// GetKey returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition.Key, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadMetaobjectDefinitionFieldDefinitionsMetaobjectFieldDefinition) GetKey() string {
	return v.Key
}

// CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError includes the requested fields of the GraphQL type MetaobjectUserError.
// The GraphQL type's documentation follows.
//
// Defines errors encountered while managing metaobject resources.
type CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError struct {
	// The path to the input field that caused the error.
	Field []string `json:"field"`
	// The error message.
	Message string `json:"message"`
	// The error code.
	Code *MetaobjectUserErrorCode `json:"code"`
}

// GetField returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError.Field, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError) GetField() []string {
	return v.Field
}

// GetMessage returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError.Message, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError) GetMessage() string {
	return v.Message
}

// GetCode returns CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError.Code, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayloadUserErrorsMetaobjectUserError) GetCode() *MetaobjectUserErrorCode {
	return v.Code
}

// CreateMetaobjectDefinitionResponse is returned by CreateMetaobjectDefinition on success.
type CreateMetaobjectDefinitionResponse struct {
	// Creates a new metaobject definition.
	MetaobjectDefinitionCreate *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload `json:"metaobjectDefinitionCreate"`
}

// GetMetaobjectDefinitionCreate returns CreateMetaobjectDefinitionResponse.MetaobjectDefinitionCreate, and is useful for accessing the field via an interface.
func (v *CreateMetaobjectDefinitionResponse) GetMetaobjectDefinitionCreate() *CreateMetaobjectDefinitionMetaobjectDefinitionCreateMetaobjectDefinitionCreatePayload {
	return v.MetaobjectDefinitionCreate
}

// DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload includes the requested fields of the GraphQL type MetafieldDefinitionDeletePayload.
// The GraphQL type's documentation follows.
//
// Return type for `metafieldDefinitionDelete` mutation.
type DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload struct {
	// The ID of the deleted metafield definition.
	DeletedDefinitionId *string `json:"deletedDefinitionId"`
	// The list of errors that occurred from executing the mutation.
	UserErrors []*DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError `json:"userErrors"`
}

// GetDeletedDefinitionId returns DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload.DeletedDefinitionId, and is useful for accessing the field via an interface.
func (v *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload) GetDeletedDefinitionId() *string {
	return v.DeletedDefinitionId
}

// GetUserErrors returns DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload.UserErrors, and is useful for accessing the field via an interface.
func (v *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload) GetUserErrors() []*DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError {
	return v.UserErrors
}

// DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError includes the requested fields of the GraphQL type MetafieldDefinitionDeleteUserError.
// The GraphQL type's documentation follows.
//
// An error that occurs during the execution of `MetafieldDefinitionDelete`.
type DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError struct {
	// The path to the input field that caused the error.
	Field []string `json:"field"`
	// The error message.
	Message string `json:"message"`
	// The error code.
	Code *MetafieldDefinitionDeleteUserErrorCode `json:"code"`
}

// GetField returns DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError.Field, and is useful for accessing the field via an interface.
func (v *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError) GetField() []string {
	return v.Field
}

// GetMessage returns DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError.Message, and is useful for accessing the field via an interface.
func (v *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError) GetMessage() string {
	return v.Message
}

// GetCode returns DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError.Code, and is useful for accessing the field via an interface.
func (v *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayloadUserErrorsMetafieldDefinitionDeleteUserError) GetCode() *MetafieldDefinitionDeleteUserErrorCode {
	return v.Code
}

// DeleteMetafieldDefinitionResponse is returned by DeleteMetafieldDefinition on success.
type DeleteMetafieldDefinitionResponse struct {
	// Delete a metafield definition. Optionally deletes all associated metafields asynchronously when specified.
	MetafieldDefinitionDelete *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload `json:"metafieldDefinitionDelete"`
}

// GetMetafieldDefinitionDelete returns DeleteMetafieldDefinitionResponse.MetafieldDefinitionDelete, and is useful for accessing the field via an interface.
func (v *DeleteMetafieldDefinitionResponse) GetMetafieldDefinitionDelete() *DeleteMetafieldDefinitionMetafieldDefinitionDeleteMetafieldDefinitionDeletePayload {
	return v.MetafieldDefinitionDelete
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection includes the requested fields of the GraphQL type MetafieldDefinitionConnection.
// The GraphQL type's documentation follows.
//
// An auto-generated type for paginating through multiple MetafieldDefinitions.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection struct {
	// The connection between the node and its parent. Each edge contains a minimum of the edge's cursor and the node.
	Edges []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdge `json:"edges"`
	// An object that’s used to retrieve [cursor information](https://shopify.dev/api/usage/pagination-graphql) about the current page.
	PageInfo *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo `json:"pageInfo"`
}

// GetEdges returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection.Edges, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection) GetEdges() []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdge {
	return v.Edges
}

// GetPageInfo returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection.PageInfo, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection) GetPageInfo() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo {
	return v.PageInfo
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdge includes the requested fields of the GraphQL type MetafieldDefinitionEdge.
// The GraphQL type's documentation follows.
//
// An auto-generated type which holds one MetafieldDefinition and a cursor during pagination.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdge struct {
	// The item at the end of MetafieldDefinitionEdge.
	Node *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition `json:"node"`
}

// GetNode returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdge.Node, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdge) GetNode() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition {
	return v.Node
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition includes the requested fields of the GraphQL type MetafieldDefinition.
// The GraphQL type's documentation follows.
//
// Metafield definitions enable you to define additional validation constraints for metafields, and enable the merchant to edit metafield values in context.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition struct {
	// A globally-unique ID.
	Id string `json:"id"`
	// The unique identifier for the metafield definition within its namespace.
	Key string `json:"key"`
	// The human-readable name of the metafield definition.
	Name string `json:"name"`
	// The container for a group of metafields that the metafield definition is associated with.
	Namespace string `json:"namespace"`
	// The description of the metafield definition.
	Description *string `json:"description"`
	// The resource type that the metafield definition is attached to.
	OwnerType MetafieldOwnerType `json:"ownerType"`
	// The position of the metafield definition in the pinned list.
	PinnedPosition *int `json:"pinnedPosition"`
	// The type of data that each of the metafields that belong to the metafield definition will store. Refer to the list of [supported types](https://shopify.dev/apps/metafields/types).
	Type *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType `json:"type"`
	// The access settings associated with the metafield definition.
	Access *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess `json:"access"`
	// The capabilities of the metafield definition.
	Capabilities *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities `json:"capabilities"`
	// The constraints that determine what subtypes of resources a metafield definition applies to.
	Constraints *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints `json:"constraints"`
	// The count of the metafields that belong to the metafield definition.
	MetafieldsCount int `json:"metafieldsCount"`
	// The standard metafield definition template associated with the metafield definition.
	StandardTemplate *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate `json:"standardTemplate"`
	// Whether the metafield definition can be used as a collection condition.
	UseAsCollectionCondition bool `json:"useAsCollectionCondition"`
	// The validation status for the metafields that belong to the metafield definition.
	ValidationStatus types.MetafieldDefinitionValidationStatus `json:"validationStatus"`
	// A list of [validation options](https://shopify.dev/apps/metafields/definitions/validation) for the metafields that belong to the metafield definition. For example, for a metafield definition with the type `date`, you can set a minimum date validation so that each of the metafields that belong to it can only store dates after the specified minimum.
	Validations []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation `json:"validations"`
}

// GetId returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Id, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetId() string {
	return v.Id
}

// GetKey returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Key, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetKey() string {
	return v.Key
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetName() string {
	return v.Name
}

// GetNamespace returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Namespace, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetNamespace() string {
	return v.Namespace
}

// GetDescription returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Description, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetDescription() *string {
	return v.Description
}

// GetOwnerType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.OwnerType, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetOwnerType() MetafieldOwnerType {
	return v.OwnerType
}

// GetPinnedPosition returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.PinnedPosition, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetPinnedPosition() *int {
	return v.PinnedPosition
}

// GetType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Type, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetType() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType {
	return v.Type
}

// GetAccess returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Access, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetAccess() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess {
	return v.Access
}

// GetCapabilities returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Capabilities, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetCapabilities() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities {
	return v.Capabilities
}

// GetConstraints returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Constraints, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetConstraints() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints {
	return v.Constraints
}

// GetMetafieldsCount returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.MetafieldsCount, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetMetafieldsCount() int {
	return v.MetafieldsCount
}

// GetStandardTemplate returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.StandardTemplate, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetStandardTemplate() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate {
	return v.StandardTemplate
}

// GetUseAsCollectionCondition returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.UseAsCollectionCondition, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetUseAsCollectionCondition() bool {
	return v.UseAsCollectionCondition
}

// GetValidationStatus returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.ValidationStatus, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetValidationStatus() types.MetafieldDefinitionValidationStatus {
	return v.ValidationStatus
}

// GetValidations returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition.Validations, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinition) GetValidations() []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation {
	return v.Validations
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess includes the requested fields of the GraphQL type MetafieldAccess.
// The GraphQL type's documentation follows.
//
// The access settings for this metafield definition.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess struct {
	// The default admin access setting used for the metafields under this definition.
	Admin *types.MetafieldAdminAccess `json:"admin"`
	// The customer account access setting used for the metafields under this definition.
	CustomerAccount types.MetafieldCustomerAccountAccess `json:"customerAccount"`
	// The storefront access setting used for the metafields under this definition.
	Storefront *types.MetafieldStorefrontAccess `json:"storefront"`
}

// GetAdmin returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess.Admin, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess) GetAdmin() *types.MetafieldAdminAccess {
	return v.Admin
}

// GetCustomerAccount returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess.CustomerAccount, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess) GetCustomerAccount() types.MetafieldCustomerAccountAccess {
	return v.CustomerAccount
}

// GetStorefront returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess.Storefront, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionAccessMetafieldAccess) GetStorefront() *types.MetafieldStorefrontAccess {
	return v.Storefront
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities includes the requested fields of the GraphQL type MetafieldCapabilities.
// The GraphQL type's documentation follows.
//
// Provides the capabilities of a metafield definition.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities struct {
	// Indicate whether a metafield definition is configured for filtering.
	AdminFilterable *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable `json:"adminFilterable"`
	// Indicate whether a metafield definition can be used as a smart collection condition.
	SmartCollectionCondition *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition `json:"smartCollectionCondition"`
	// Indicate whether the metafield values for a metafield definition are required to be unique.
	UniqueValues *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues `json:"uniqueValues"`
}

// GetAdminFilterable returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities.AdminFilterable, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities) GetAdminFilterable() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable {
	return v.AdminFilterable
}

// GetSmartCollectionCondition returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities.SmartCollectionCondition, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities) GetSmartCollectionCondition() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition {
	return v.SmartCollectionCondition
}

// GetUniqueValues returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities.UniqueValues, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilities) GetUniqueValues() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues {
	return v.UniqueValues
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable includes the requested fields of the GraphQL type MetafieldCapabilityAdminFilterable.
// The GraphQL type's documentation follows.
//
// Information about the admin filterable capability on a metafield definition.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable struct {
	// Indicates if the definition is eligible to have the capability.
	Eligible bool `json:"eligible"`
	// Indicates if the capability is enabled.
	Enabled bool `json:"enabled"`
	// Determines the metafield definition's filter status for use in admin filtering.
	Status MetafieldDefinitionAdminFilterStatus `json:"status"`
}

// GetEligible returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable.Eligible, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable) GetEligible() bool {
	return v.Eligible
}

// GetEnabled returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable.Enabled, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable) GetEnabled() bool {
	return v.Enabled
}

// GetStatus returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable.Status, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesAdminFilterableMetafieldCapabilityAdminFilterable) GetStatus() MetafieldDefinitionAdminFilterStatus {
	return v.Status
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition includes the requested fields of the GraphQL type MetafieldCapabilitySmartCollectionCondition.
// The GraphQL type's documentation follows.
//
// Information about the smart collection condition capability on a metafield definition.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition struct {
	// Indicates if the definition is eligible to have the capability.
	Eligible bool `json:"eligible"`
	// Indicates if the capability is enabled.
	Enabled bool `json:"enabled"`
}

// GetEligible returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition.Eligible, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition) GetEligible() bool {
	return v.Eligible
}

// GetEnabled returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition.Enabled, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesSmartCollectionConditionMetafieldCapabilitySmartCollectionCondition) GetEnabled() bool {
	return v.Enabled
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues includes the requested fields of the GraphQL type MetafieldCapabilityUniqueValues.
// The GraphQL type's documentation follows.
//
// Information about the unique values capability on a metafield definition.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues struct {
	// Indicates if the definition is eligible to have the capability.
	Eligible bool `json:"eligible"`
	// Indicates if the capability is enabled.
	Enabled bool `json:"enabled"`
}

// GetEligible returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues.Eligible, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues) GetEligible() bool {
	return v.Eligible
}

// GetEnabled returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues.Enabled, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionCapabilitiesMetafieldCapabilitiesUniqueValuesMetafieldCapabilityUniqueValues) GetEnabled() bool {
	return v.Enabled
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints includes the requested fields of the GraphQL type MetafieldDefinitionConstraints.
// The GraphQL type's documentation follows.
//
// The constraints that determine what subtypes of resources a metafield definition applies to.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints struct {
	// The category of resource subtypes that the definition applies to.
	Key *string `json:"key"`
	// The specific constraint subtype values that the definition applies to.
	Values *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnection `json:"values"`
}

// GetKey returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints.Key, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints) GetKey() *string {
	return v.Key
}

// GetValues returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints.Values, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraints) GetValues() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnection {
	return v.Values
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnection includes the requested fields of the GraphQL type MetafieldDefinitionConstraintValueConnection.
// The GraphQL type's documentation follows.
//
// An auto-generated type for paginating through multiple MetafieldDefinitionConstraintValues.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnection struct {
	// The connection between the node and its parent. Each edge contains a minimum of the edge's cursor and the node.
	Edges []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdge `json:"edges"`
}

// GetEdges returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnection.Edges, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnection) GetEdges() []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdge {
	return v.Edges
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdge includes the requested fields of the GraphQL type MetafieldDefinitionConstraintValueEdge.
// The GraphQL type's documentation follows.
//
// An auto-generated type which holds one MetafieldDefinitionConstraintValue and a cursor during pagination.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdge struct {
	// The item at the end of MetafieldDefinitionConstraintValueEdge.
	Node *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdgeNodeMetafieldDefinitionConstraintValue `json:"node"`
}

// GetNode returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdge.Node, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdge) GetNode() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdgeNodeMetafieldDefinitionConstraintValue {
	return v.Node
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdgeNodeMetafieldDefinitionConstraintValue includes the requested fields of the GraphQL type MetafieldDefinitionConstraintValue.
// The GraphQL type's documentation follows.
//
// A constraint subtype value that the metafield definition applies to.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdgeNodeMetafieldDefinitionConstraintValue struct {
	// The subtype value of the constraint.
	Value string `json:"value"`
}

// GetValue returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdgeNodeMetafieldDefinitionConstraintValue.Value, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionConstraintsValuesMetafieldDefinitionConstraintValueConnectionEdgesMetafieldDefinitionConstraintValueEdgeNodeMetafieldDefinitionConstraintValue) GetValue() string {
	return v.Value
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate includes the requested fields of the GraphQL type StandardMetafieldDefinitionTemplate.
// The GraphQL type's documentation follows.
//
// Standard metafield definition templates provide preset configurations to create metafield definitions. Each template has a specific namespace and key that we've reserved to have specific meanings for common use cases.  Refer to the [list of standard metafield definitions](https://shopify.dev/apps/metafields/definitions/standard-definitions).
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate struct {
	// The description of the standard metafield definition.
	Description *string `json:"description"`
	// A globally-unique ID.
	Id string `json:"id"`
	// The key owned by the definition after the definition has been activated.
	Key string `json:"key"`
	// The human-readable name for the standard metafield definition.
	Name string `json:"name"`
	// The namespace owned by the definition after the definition has been activated.
	Namespace string `json:"namespace"`
	// The list of resource types that the standard metafield definition can be applied to.
	OwnerTypes []MetafieldOwnerType `json:"ownerTypes"`
	// The associated [metafield definition type](https://shopify.dev/apps/metafields/definitions/types) that the metafield stores.
	Type *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType `json:"type"`
	// The configured validations for the standard metafield definition.
	Validations []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation `json:"validations"`
	// Whether metafields for the definition are by default visible using the Storefront API.
	VisibleToStorefrontApi bool `json:"visibleToStorefrontApi"`
}

// GetDescription returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Description, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetDescription() *string {
	return v.Description
}

// GetId returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Id, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetId() string {
	return v.Id
}

// GetKey returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Key, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetKey() string {
	return v.Key
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetName() string {
	return v.Name
}

// GetNamespace returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Namespace, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetNamespace() string {
	return v.Namespace
}

// GetOwnerTypes returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.OwnerTypes, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetOwnerTypes() []MetafieldOwnerType {
	return v.OwnerTypes
}

// GetType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Type, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetType() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType {
	return v.Type
}

// GetValidations returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.Validations, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetValidations() []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation {
	return v.Validations
}

// GetVisibleToStorefrontApi returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate.VisibleToStorefrontApi, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplate) GetVisibleToStorefrontApi() bool {
	return v.VisibleToStorefrontApi
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType includes the requested fields of the GraphQL type MetafieldDefinitionType.
// The GraphQL type's documentation follows.
//
// A metafield definition type provides basic foundation and validation for a metafield.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType struct {
	// The category associated with the metafield definition type.
	Category string `json:"category"`
	// The name of the type for the metafield definition. See the list of [supported types](https://shopify.dev/apps/metafields/types).
	Name string `json:"name"`
	// The supported validations for a metafield definition type.
	SupportedValidations []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation `json:"supportedValidations"`
	// Whether metafields without a definition can be migrated to a definition of this type.
	SupportsDefinitionMigrations bool `json:"supportsDefinitionMigrations"`
	// The value type for a metafield created with this definition type.
	ValueType types.MetafieldValueType `json:"valueType"`
}

// GetCategory returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType.Category, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType) GetCategory() string {
	return v.Category
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType) GetName() string {
	return v.Name
}

// GetSupportedValidations returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType.SupportedValidations, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType) GetSupportedValidations() []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation {
	return v.SupportedValidations
}

// GetSupportsDefinitionMigrations returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType.SupportsDefinitionMigrations, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType) GetSupportsDefinitionMigrations() bool {
	return v.SupportsDefinitionMigrations
}

// GetValueType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType.ValueType, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionType) GetValueType() types.MetafieldValueType {
	return v.ValueType
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation includes the requested fields of the GraphQL type MetafieldDefinitionSupportedValidation.
// The GraphQL type's documentation follows.
//
// The type and name for the optional validation configuration of a metafield.  For example, a supported validation might consist of a `max` name and a `number_integer` type. This validation can then be used to enforce a maximum character length for a `single_line_text_field` metafield.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation struct {
	// The name of the metafield definition validation.
	Name string `json:"name"`
	// The type of input for the validation.
	Type string `json:"type"`
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation) GetName() string {
	return v.Name
}

// GetType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation.Type, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateTypeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation) GetType() string {
	return v.Type
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation includes the requested fields of the GraphQL type MetafieldDefinitionValidation.
// The GraphQL type's documentation follows.
//
// A configured metafield definition validation.  For example, for a metafield definition of `number_integer` type, you can set a validation with the name `max` and a value of `15`. This validation will ensure that the value of the metafield is a number less than or equal to 15.  Refer to the [list of supported validations](https://shopify.dev/api/admin/graphql/reference/common-objects/metafieldDefinitionTypes#examples-Fetch_all_metafield_definition_types).
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation struct {
	// The validation name.
	Name string `json:"name"`
	// The name for the metafield type of this validation.
	Type string `json:"type"`
	// The validation value.
	Value *string `json:"value"`
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation) GetName() string {
	return v.Name
}

// GetType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation.Type, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation) GetType() string {
	return v.Type
}

// GetValue returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation.Value, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionStandardTemplateStandardMetafieldDefinitionTemplateValidationsMetafieldDefinitionValidation) GetValue() *string {
	return v.Value
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType includes the requested fields of the GraphQL type MetafieldDefinitionType.
// The GraphQL type's documentation follows.
//
// A metafield definition type provides basic foundation and validation for a metafield.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType struct {
	// The category associated with the metafield definition type.
	Category string `json:"category"`
	// The name of the type for the metafield definition. See the list of [supported types](https://shopify.dev/apps/metafields/types).
	Name string `json:"name"`
	// The supported validations for a metafield definition type.
	SupportedValidations []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation `json:"supportedValidations"`
	// Whether metafields without a definition can be migrated to a definition of this type.
	SupportsDefinitionMigrations bool `json:"supportsDefinitionMigrations"`
	// The value type for a metafield created with this definition type.
	ValueType types.MetafieldValueType `json:"valueType"`
}

// GetCategory returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType.Category, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType) GetCategory() string {
	return v.Category
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType) GetName() string {
	return v.Name
}

// GetSupportedValidations returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType.SupportedValidations, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType) GetSupportedValidations() []*GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation {
	return v.SupportedValidations
}

// GetSupportsDefinitionMigrations returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType.SupportsDefinitionMigrations, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType) GetSupportsDefinitionMigrations() bool {
	return v.SupportsDefinitionMigrations
}

// GetValueType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType.ValueType, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionType) GetValueType() types.MetafieldValueType {
	return v.ValueType
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation includes the requested fields of the GraphQL type MetafieldDefinitionSupportedValidation.
// The GraphQL type's documentation follows.
//
// The type and name for the optional validation configuration of a metafield.  For example, a supported validation might consist of a `max` name and a `number_integer` type. This validation can then be used to enforce a maximum character length for a `single_line_text_field` metafield.
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation struct {
	// The name of the metafield definition validation.
	Name string `json:"name"`
	// The type of input for the validation.
	Type string `json:"type"`
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation) GetName() string {
	return v.Name
}

// GetType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation.Type, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionTypeSupportedValidationsMetafieldDefinitionSupportedValidation) GetType() string {
	return v.Type
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation includes the requested fields of the GraphQL type MetafieldDefinitionValidation.
// The GraphQL type's documentation follows.
//
// A configured metafield definition validation.  For example, for a metafield definition of `number_integer` type, you can set a validation with the name `max` and a value of `15`. This validation will ensure that the value of the metafield is a number less than or equal to 15.  Refer to the [list of supported validations](https://shopify.dev/api/admin/graphql/reference/common-objects/metafieldDefinitionTypes#examples-Fetch_all_metafield_definition_types).
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation struct {
	// The validation name.
	Name string `json:"name"`
	// The name for the metafield type of this validation.
	Type string `json:"type"`
	// The validation value.
	Value *string `json:"value"`
}

// GetName returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation.Name, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation) GetName() string {
	return v.Name
}

// GetType returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation.Type, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation) GetType() string {
	return v.Type
}

// GetValue returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation.Value, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionEdgesMetafieldDefinitionEdgeNodeMetafieldDefinitionValidationsMetafieldDefinitionValidation) GetValue() *string {
	return v.Value
}

// GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo includes the requested fields of the GraphQL type PageInfo.
// The GraphQL type's documentation follows.
//
// Returns information about pagination in a connection, in accordance with the [Relay specification](https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo). For more information, please read our [GraphQL Pagination Usage Guide](https://shopify.dev/api/usage/pagination-graphql).
type GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo struct {
	// Whether there are more pages to fetch following the current page.
	HasNextPage bool `json:"hasNextPage"`
	// The cursor corresponding to the last node in edges.
	EndCursor *string `json:"endCursor"`
}

// GetHasNextPage returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo.HasNextPage, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo) GetHasNextPage() bool {
	return v.HasNextPage
}

// GetEndCursor returns GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo.EndCursor, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnectionPageInfo) GetEndCursor() *string {
	return v.EndCursor
}

// GetMetafieldDefinitionsResponse is returned by GetMetafieldDefinitions on success.
type GetMetafieldDefinitionsResponse struct {
	// Returns a list of metafield definitions.
	MetafieldDefinitions *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection `json:"metafieldDefinitions"`
}

// GetMetafieldDefinitions returns GetMetafieldDefinitionsResponse.MetafieldDefinitions, and is useful for accessing the field via an interface.
func (v *GetMetafieldDefinitionsResponse) GetMetafieldDefinitions() *GetMetafieldDefinitionsMetafieldDefinitionsMetafieldDefinitionConnection {
	return v.MetafieldDefinitions
}

// GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection includes the requested fields of the GraphQL type MetaobjectDefinitionConnection.
// The GraphQL type's documentation follows.
//
// An auto-generated type for paginating through multiple MetaobjectDefinitions.
type GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection struct {
	// The connection between the node and its parent. Each edge contains a minimum of the edge's cursor and the node.
	Edges []*GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge `json:"edges"`
	// An object that’s used to retrieve [cursor information](https://shopify.dev/api/usage/pagination-graphql) about the current page.
	PageInfo *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo `json:"pageInfo"`
}

// GetEdges returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection.Edges, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection) GetEdges() []*GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge {
	return v.Edges
}

// GetPageInfo returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection.PageInfo, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection) GetPageInfo() *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo {
	return v.PageInfo
}

// GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge includes the requested fields of the GraphQL type MetaobjectDefinitionEdge.
// The GraphQL type's documentation follows.
//
// An auto-generated type which holds one MetaobjectDefinition and a cursor during pagination.
type GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge struct {
	// The position of each node in an array, used in [pagination](https://shopify.dev/api/usage/pagination-graphql).
	Cursor string `json:"cursor"`
	// The item at the end of MetaobjectDefinitionEdge.
	Node *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition `json:"node"`
}

// GetCursor returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge.Cursor, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge) GetCursor() string {
	return v.Cursor
}

// GetNode returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge.Node, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdge) GetNode() *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition {
	return v.Node
}

// GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition includes the requested fields of the GraphQL type MetaobjectDefinition.
// The GraphQL type's documentation follows.
//
// Provides the definition of a generic object structure composed of metafields.
type GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition struct {
	// A globally-unique ID.
	Id string `json:"id"`
	// The human-readable name.
	Name string `json:"name"`
	// The type of the object definition. Defines the namespace of associated metafields.
	Type string `json:"type"`
	// The administrative description.
	Description *string `json:"description"`
	// The count of metaobjects created for the definition.
	MetaobjectsCount int `json:"metaobjectsCount"`
}

// GetId returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition.Id, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition) GetId() string {
	return v.Id
}

// GetName returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition.Name, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition) GetName() string {
	return v.Name
}

// GetType returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition.Type, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition) GetType() string {
	return v.Type
}

// GetDescription returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition.Description, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition) GetDescription() *string {
	return v.Description
}

// GetMetaobjectsCount returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition.MetaobjectsCount, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionEdgesMetaobjectDefinitionEdgeNodeMetaobjectDefinition) GetMetaobjectsCount() int {
	return v.MetaobjectsCount
}

// GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo includes the requested fields of the GraphQL type PageInfo.
// The GraphQL type's documentation follows.
//
// Returns information about pagination in a connection, in accordance with the [Relay specification](https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo). For more information, please read our [GraphQL Pagination Usage Guide](https://shopify.dev/api/usage/pagination-graphql).
type GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo struct {
	// Whether there are more pages to fetch following the current page.
	HasNextPage bool `json:"hasNextPage"`
	// The cursor corresponding to the last node in edges.
	EndCursor *string `json:"endCursor"`
}

// GetHasNextPage returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo.HasNextPage, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo) GetHasNextPage() bool {
	return v.HasNextPage
}

// GetEndCursor returns GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo.EndCursor, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnectionPageInfo) GetEndCursor() *string {
	return v.EndCursor
}

// GetMetaobjectDefinitionsResponse is returned by GetMetaobjectDefinitions on success.
type GetMetaobjectDefinitionsResponse struct {
	// All metaobject definitions.
	MetaobjectDefinitions *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection `json:"metaobjectDefinitions"`
}

// GetMetaobjectDefinitions returns GetMetaobjectDefinitionsResponse.MetaobjectDefinitions, and is useful for accessing the field via an interface.
func (v *GetMetaobjectDefinitionsResponse) GetMetaobjectDefinitions() *GetMetaobjectDefinitionsMetaobjectDefinitionsMetaobjectDefinitionConnection {
	return v.MetaobjectDefinitions
}

// The input fields for the access settings for the metafields under the definition.
type MetafieldAccessInput struct {
	// The admin access setting to use for the metafields under this definition.
	Admin *types.MetafieldAdminAccessInput `json:"admin"`
	// The storefront access setting to use for the metafields under this definition.
	Storefront *types.MetafieldStorefrontAccessInput `json:"storefront"`
	// The Customer Account API access setting to use for the metafields under this definition.
	CustomerAccount *types.MetafieldCustomerAccountAccessInput `json:"customerAccount"`
}

// GetAdmin returns MetafieldAccessInput.Admin, and is useful for accessing the field via an interface.
func (v *MetafieldAccessInput) GetAdmin() *types.MetafieldAdminAccessInput { return v.Admin }

// GetStorefront returns MetafieldAccessInput.Storefront, and is useful for accessing the field via an interface.
func (v *MetafieldAccessInput) GetStorefront() *types.MetafieldStorefrontAccessInput {
	return v.Storefront
}

// GetCustomerAccount returns MetafieldAccessInput.CustomerAccount, and is useful for accessing the field via an interface.
func (v *MetafieldAccessInput) GetCustomerAccount() *types.MetafieldCustomerAccountAccessInput {
	return v.CustomerAccount
}

// The input fields for enabling and disabling the admin filterable capability.
type MetafieldCapabilityAdminFilterableInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns MetafieldCapabilityAdminFilterableInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetafieldCapabilityAdminFilterableInput) GetEnabled() bool { return v.Enabled }

// The input fields for creating a metafield capability.
type MetafieldCapabilityCreateInput struct {
	// The input for updating the smart collection condition capability.
	SmartCollectionCondition *MetafieldCapabilitySmartCollectionConditionInput `json:"smartCollectionCondition,omitempty"`
	// The input for updating the admin filterable capability.
	AdminFilterable *MetafieldCapabilityAdminFilterableInput `json:"adminFilterable,omitempty"`
	// The input for updating the unique values capability.
	UniqueValues *MetafieldCapabilityUniqueValuesInput `json:"uniqueValues,omitempty"`
}

// GetSmartCollectionCondition returns MetafieldCapabilityCreateInput.SmartCollectionCondition, and is useful for accessing the field via an interface.
func (v *MetafieldCapabilityCreateInput) GetSmartCollectionCondition() *MetafieldCapabilitySmartCollectionConditionInput {
	return v.SmartCollectionCondition
}

// GetAdminFilterable returns MetafieldCapabilityCreateInput.AdminFilterable, and is useful for accessing the field via an interface.
func (v *MetafieldCapabilityCreateInput) GetAdminFilterable() *MetafieldCapabilityAdminFilterableInput {
	return v.AdminFilterable
}

// GetUniqueValues returns MetafieldCapabilityCreateInput.UniqueValues, and is useful for accessing the field via an interface.
func (v *MetafieldCapabilityCreateInput) GetUniqueValues() *MetafieldCapabilityUniqueValuesInput {
	return v.UniqueValues
}

// The input fields for enabling and disabling the smart collection condition capability.
type MetafieldCapabilitySmartCollectionConditionInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns MetafieldCapabilitySmartCollectionConditionInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetafieldCapabilitySmartCollectionConditionInput) GetEnabled() bool { return v.Enabled }

// The input fields for enabling and disabling the unique values capability.
type MetafieldCapabilityUniqueValuesInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns MetafieldCapabilityUniqueValuesInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetafieldCapabilityUniqueValuesInput) GetEnabled() bool { return v.Enabled }

// Possible filter statuses associated with a metafield definition for use in admin filtering.
type MetafieldDefinitionAdminFilterStatus string

const (
	// The metafield definition cannot be used for admin filtering.
	MetafieldDefinitionAdminFilterStatusNotFilterable MetafieldDefinitionAdminFilterStatus = "NOT_FILTERABLE"
	// The metafield definition's metafields are currently being processed for admin filtering.
	MetafieldDefinitionAdminFilterStatusInProgress MetafieldDefinitionAdminFilterStatus = "IN_PROGRESS"
	// The metafield definition allows admin filtering by matching metafield values.
	MetafieldDefinitionAdminFilterStatusFilterable MetafieldDefinitionAdminFilterStatus = "FILTERABLE"
	// The metafield definition has failed to be enabled for admin filtering.
	MetafieldDefinitionAdminFilterStatusFailed MetafieldDefinitionAdminFilterStatus = "FAILED"
)

// The input fields required to create metafield definition constraints. Each constraint applies a metafield definition to a subtype of a resource.
type MetafieldDefinitionConstraintsInput struct {
	// The category of resource subtypes that the definition applies to.
	Key string `json:"key"`
	// The specific constraint subtype values that the definition applies to.
	Values []string `json:"values"`
}

// GetKey returns MetafieldDefinitionConstraintsInput.Key, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionConstraintsInput) GetKey() string { return v.Key }

// GetValues returns MetafieldDefinitionConstraintsInput.Values, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionConstraintsInput) GetValues() []string { return v.Values }

// Possible error codes that can be returned by `MetafieldDefinitionCreateUserError`.
type MetafieldDefinitionCreateUserErrorCode string

const (
	// The input value is invalid.
	MetafieldDefinitionCreateUserErrorCodeInvalid MetafieldDefinitionCreateUserErrorCode = "INVALID"
	// The input value isn't included in the list.
	MetafieldDefinitionCreateUserErrorCodeInclusion MetafieldDefinitionCreateUserErrorCode = "INCLUSION"
	// The input value needs to be blank.
	MetafieldDefinitionCreateUserErrorCodePresent MetafieldDefinitionCreateUserErrorCode = "PRESENT"
	// The input value is already taken.
	MetafieldDefinitionCreateUserErrorCodeTaken MetafieldDefinitionCreateUserErrorCode = "TAKEN"
	// The input value is too long.
	MetafieldDefinitionCreateUserErrorCodeTooLong MetafieldDefinitionCreateUserErrorCode = "TOO_LONG"
	// The input value is too short.
	MetafieldDefinitionCreateUserErrorCodeTooShort MetafieldDefinitionCreateUserErrorCode = "TOO_SHORT"
	// A capability is required for the definition type but is disabled.
	MetafieldDefinitionCreateUserErrorCodeCapabilityRequiredButDisabled MetafieldDefinitionCreateUserErrorCode = "CAPABILITY_REQUIRED_BUT_DISABLED"
	// The definition limit per owner type has exceeded.
	MetafieldDefinitionCreateUserErrorCodeResourceTypeLimitExceeded MetafieldDefinitionCreateUserErrorCode = "RESOURCE_TYPE_LIMIT_EXCEEDED"
	// The maximum limit of definitions per owner type has exceeded.
	MetafieldDefinitionCreateUserErrorCodeLimitExceeded MetafieldDefinitionCreateUserErrorCode = "LIMIT_EXCEEDED"
	// An invalid option.
	MetafieldDefinitionCreateUserErrorCodeInvalidOption MetafieldDefinitionCreateUserErrorCode = "INVALID_OPTION"
	// A duplicate option.
	MetafieldDefinitionCreateUserErrorCodeDuplicateOption MetafieldDefinitionCreateUserErrorCode = "DUPLICATE_OPTION"
	// This namespace and key combination is reserved for standard definitions.
	MetafieldDefinitionCreateUserErrorCodeReservedNamespaceKey MetafieldDefinitionCreateUserErrorCode = "RESERVED_NAMESPACE_KEY"
	// The pinned limit has been reached for the owner type.
	MetafieldDefinitionCreateUserErrorCodePinnedLimitReached MetafieldDefinitionCreateUserErrorCode = "PINNED_LIMIT_REACHED"
	// This namespace and key combination is already in use for a set of your metafields.
	MetafieldDefinitionCreateUserErrorCodeUnstructuredAlreadyExists MetafieldDefinitionCreateUserErrorCode = "UNSTRUCTURED_ALREADY_EXISTS"
	// The metafield definition does not support pinning.
	MetafieldDefinitionCreateUserErrorCodeUnsupportedPinning MetafieldDefinitionCreateUserErrorCode = "UNSUPPORTED_PINNING"
	// A field contains an invalid character.
	MetafieldDefinitionCreateUserErrorCodeInvalidCharacter MetafieldDefinitionCreateUserErrorCode = "INVALID_CHARACTER"
	// The definition type is not eligible to be used as collection condition.
	MetafieldDefinitionCreateUserErrorCodeTypeNotAllowedForConditions MetafieldDefinitionCreateUserErrorCode = "TYPE_NOT_ALLOWED_FOR_CONDITIONS"
	// You have reached the maximum allowed definitions for automated collections.
	MetafieldDefinitionCreateUserErrorCodeOwnerTypeLimitExceededForAutomatedCollections MetafieldDefinitionCreateUserErrorCode = "OWNER_TYPE_LIMIT_EXCEEDED_FOR_AUTOMATED_COLLECTIONS"
	// The metafield definition constraints are invalid.
	MetafieldDefinitionCreateUserErrorCodeInvalidConstraints MetafieldDefinitionCreateUserErrorCode = "INVALID_CONSTRAINTS"
	// The maximum limit of grants per definition type has been exceeded.
	MetafieldDefinitionCreateUserErrorCodeGrantLimitExceeded MetafieldDefinitionCreateUserErrorCode = "GRANT_LIMIT_EXCEEDED"
	// The input combination is invalid.
	MetafieldDefinitionCreateUserErrorCodeInvalidInputCombination MetafieldDefinitionCreateUserErrorCode = "INVALID_INPUT_COMBINATION"
	// The metafield definition capability is invalid.
	MetafieldDefinitionCreateUserErrorCodeInvalidCapability MetafieldDefinitionCreateUserErrorCode = "INVALID_CAPABILITY"
	// Admin access can only be specified for app-owned metafield definitions.
	MetafieldDefinitionCreateUserErrorCodeAdminAccessInputNotAllowed MetafieldDefinitionCreateUserErrorCode = "ADMIN_ACCESS_INPUT_NOT_ALLOWED"
)

// Possible error codes that can be returned by `MetafieldDefinitionDeleteUserError`.
type MetafieldDefinitionDeleteUserErrorCode string

const (
	// The input value needs to be blank.
	MetafieldDefinitionDeleteUserErrorCodePresent MetafieldDefinitionDeleteUserErrorCode = "PRESENT"
	// Definition not found.
	MetafieldDefinitionDeleteUserErrorCodeNotFound MetafieldDefinitionDeleteUserErrorCode = "NOT_FOUND"
	// An internal error occurred.
	MetafieldDefinitionDeleteUserErrorCodeInternalError MetafieldDefinitionDeleteUserErrorCode = "INTERNAL_ERROR"
	// Deleting an id type metafield definition requires deletion of its associated metafields.
	MetafieldDefinitionDeleteUserErrorCodeIdTypeDeletionError MetafieldDefinitionDeleteUserErrorCode = "ID_TYPE_DELETION_ERROR"
	// Deleting a reference type metafield definition requires deletion of its associated metafields.
	MetafieldDefinitionDeleteUserErrorCodeReferenceTypeDeletionError MetafieldDefinitionDeleteUserErrorCode = "REFERENCE_TYPE_DELETION_ERROR"
	// Deleting a definition in a reserved namespace requires deletion of its associated metafields.
	MetafieldDefinitionDeleteUserErrorCodeReservedNamespaceOrphanedMetafields MetafieldDefinitionDeleteUserErrorCode = "RESERVED_NAMESPACE_ORPHANED_METAFIELDS"
	// Action cannot proceed. Definition is currently in use.
	MetafieldDefinitionDeleteUserErrorCodeMetafieldDefinitionInUse MetafieldDefinitionDeleteUserErrorCode = "METAFIELD_DEFINITION_IN_USE"
	// Owner type can't be used in this mutation.
	MetafieldDefinitionDeleteUserErrorCodeDisallowedOwnerType MetafieldDefinitionDeleteUserErrorCode = "DISALLOWED_OWNER_TYPE"
)

// The input fields required to create a metafield definition.
type MetafieldDefinitionInput struct {
	// The container for a group of metafields that the metafield definition will be associated with. If omitted, the app-reserved namespace will be used.  Must be 3-255 characters long and only contain alphanumeric, hyphen, and underscore characters.
	Namespace *string `json:"namespace"`
	// The unique identifier for the metafield definition within its namespace.  Must be 2-64 characters long and only contain alphanumeric, hyphen, and underscore characters.
	Key string `json:"key"`
	// The human-readable name for the metafield definition.
	Name string `json:"name"`
	// The description for the metafield definition.
	Description *string `json:"description"`
	// The resource type that the metafield definition is attached to.
	OwnerType MetafieldOwnerType `json:"ownerType"`
	// The type of data that each of the metafields that belong to the metafield definition will store. Refer to the list of [supported types](https://shopify.dev/apps/metafields/types).
	Type string `json:"type"`
	// A list of [validation options](https://shopify.dev/apps/metafields/definitions/validation) for the metafields that belong to the metafield definition. For example, for a metafield definition with the type `date`, you can set a minimum date validation so that each of the metafields that belong to it can only store dates after the specified minimum.
	Validations []*MetafieldDefinitionValidationInput `json:"validations,omitempty"`
	// Whether the metafield definition can be used as a collection condition.
	UseAsCollectionCondition *bool `json:"useAsCollectionCondition"`
	// Whether to [pin](https://help.shopify.com/manual/custom-data/metafields/pinning-metafield-definitions) the metafield definition.
	Pin *bool `json:"pin"`
	// The access settings that apply to each of the metafields that belong to the metafield definition.
	Access *MetafieldAccessInput `json:"access,omitempty"`
	// The constraints that determine what resources a metafield definition applies to.
	Constraints *MetafieldDefinitionConstraintsInput `json:"constraints,omitempty"`
	// The capabilities of the metafield definition.
	Capabilities *MetafieldCapabilityCreateInput `json:"capabilities,omitempty"`
}

// GetNamespace returns MetafieldDefinitionInput.Namespace, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetNamespace() *string { return v.Namespace }

// GetKey returns MetafieldDefinitionInput.Key, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetKey() string { return v.Key }

// GetName returns MetafieldDefinitionInput.Name, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetName() string { return v.Name }

// GetDescription returns MetafieldDefinitionInput.Description, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetDescription() *string { return v.Description }

// GetOwnerType returns MetafieldDefinitionInput.OwnerType, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetOwnerType() MetafieldOwnerType { return v.OwnerType }

// GetType returns MetafieldDefinitionInput.Type, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetType() string { return v.Type }

// GetValidations returns MetafieldDefinitionInput.Validations, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetValidations() []*MetafieldDefinitionValidationInput {
	return v.Validations
}

// GetUseAsCollectionCondition returns MetafieldDefinitionInput.UseAsCollectionCondition, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetUseAsCollectionCondition() *bool {
	return v.UseAsCollectionCondition
}

// GetPin returns MetafieldDefinitionInput.Pin, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetPin() *bool { return v.Pin }

// GetAccess returns MetafieldDefinitionInput.Access, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetAccess() *MetafieldAccessInput { return v.Access }

// GetConstraints returns MetafieldDefinitionInput.Constraints, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetConstraints() *MetafieldDefinitionConstraintsInput {
	return v.Constraints
}

// GetCapabilities returns MetafieldDefinitionInput.Capabilities, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionInput) GetCapabilities() *MetafieldCapabilityCreateInput {
	return v.Capabilities
}

// The name and value for a metafield definition validation.  For example, for a metafield definition of `single_line_text_field` type, you can set a validation with the name `min` and a value of `10`. This validation will ensure that the value of the metafield is at least 10 characters.  Refer to the [list of supported validations](https://shopify.dev/api/admin/graphql/reference/common-objects/metafieldDefinitionTypes#examples-Fetch_all_metafield_definition_types).
type MetafieldDefinitionValidationInput struct {
	// The name for the metafield definition validation.
	Name string `json:"name"`
	// The value for the metafield definition validation.
	Value string `json:"value"`
}

// GetName returns MetafieldDefinitionValidationInput.Name, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionValidationInput) GetName() string { return v.Name }

// GetValue returns MetafieldDefinitionValidationInput.Value, and is useful for accessing the field via an interface.
func (v *MetafieldDefinitionValidationInput) GetValue() string { return v.Value }

// Possible types of a metafield's owner resource.
type MetafieldOwnerType string

const (
	// The Api Permission metafield owner type.
	MetafieldOwnerTypeApiPermission MetafieldOwnerType = "API_PERMISSION"
	// The Company metafield owner type.
	MetafieldOwnerTypeCompany MetafieldOwnerType = "COMPANY"
	// The Company Location metafield owner type.
	MetafieldOwnerTypeCompanyLocation MetafieldOwnerType = "COMPANY_LOCATION"
	// The Payment Customization metafield owner type.
	MetafieldOwnerTypePaymentCustomization MetafieldOwnerType = "PAYMENT_CUSTOMIZATION"
	// The Validation metafield owner type.
	MetafieldOwnerTypeValidation MetafieldOwnerType = "VALIDATION"
	// The Customer metafield owner type.
	MetafieldOwnerTypeCustomer MetafieldOwnerType = "CUSTOMER"
	// The Delivery Customization metafield owner type.
	MetafieldOwnerTypeDeliveryCustomization MetafieldOwnerType = "DELIVERY_CUSTOMIZATION"
	// The draft order metafield owner type.
	MetafieldOwnerTypeDraftorder MetafieldOwnerType = "DRAFTORDER"
	// The GiftCardTransaction metafield owner type.
	MetafieldOwnerTypeGiftCardTransaction MetafieldOwnerType = "GIFT_CARD_TRANSACTION"
	// The Market metafield owner type.
	MetafieldOwnerTypeMarket MetafieldOwnerType = "MARKET"
	// The Cart Transform metafield owner type.
	MetafieldOwnerTypeCarttransform MetafieldOwnerType = "CARTTRANSFORM"
	// The Collection metafield owner type.
	MetafieldOwnerTypeCollection MetafieldOwnerType = "COLLECTION"
	// The Media Image metafield owner type.
	MetafieldOwnerTypeMediaImage MetafieldOwnerType = "MEDIA_IMAGE"
	// The Product metafield owner type.
	MetafieldOwnerTypeProduct MetafieldOwnerType = "PRODUCT"
	// The Product Variant metafield owner type.
	MetafieldOwnerTypeProductvariant MetafieldOwnerType = "PRODUCTVARIANT"
	// The Selling Plan metafield owner type.
	MetafieldOwnerTypeSellingPlan MetafieldOwnerType = "SELLING_PLAN"
	// The Article metafield owner type.
	MetafieldOwnerTypeArticle MetafieldOwnerType = "ARTICLE"
	// The Blog metafield owner type.
	MetafieldOwnerTypeBlog MetafieldOwnerType = "BLOG"
	// The Page metafield owner type.
	MetafieldOwnerTypePage MetafieldOwnerType = "PAGE"
	// The Fulfillment Constraint Rule metafield owner type.
	MetafieldOwnerTypeFulfillmentConstraintRule MetafieldOwnerType = "FULFILLMENT_CONSTRAINT_RULE"
	// The Order Routing Location Rule metafield owner type.
	MetafieldOwnerTypeOrderRoutingLocationRule MetafieldOwnerType = "ORDER_ROUTING_LOCATION_RULE"
	// The Discount metafield owner type.
	MetafieldOwnerTypeDiscount MetafieldOwnerType = "DISCOUNT"
	// The Order metafield owner type.
	MetafieldOwnerTypeOrder MetafieldOwnerType = "ORDER"
	// The Location metafield owner type.
	MetafieldOwnerTypeLocation MetafieldOwnerType = "LOCATION"
	// The Shop metafield owner type.
	MetafieldOwnerTypeShop MetafieldOwnerType = "SHOP"
)

// The input fields for configuring metaobject access controls.
type MetaobjectAccessInput struct {
	// Access configuration for Admin API surface areas, including the GraphQL Admin API.
	Admin *MetaobjectAdminAccessInput `json:"admin"`
	// Access configuration for Storefront API surface areas, including the GraphQL Storefront API and Liquid.
	Storefront *MetaobjectStorefrontAccess `json:"storefront"`
}

// GetAdmin returns MetaobjectAccessInput.Admin, and is useful for accessing the field via an interface.
func (v *MetaobjectAccessInput) GetAdmin() *MetaobjectAdminAccessInput { return v.Admin }

// GetStorefront returns MetaobjectAccessInput.Storefront, and is useful for accessing the field via an interface.
func (v *MetaobjectAccessInput) GetStorefront() *MetaobjectStorefrontAccess { return v.Storefront }

// Defines how the metaobjects of a definition can be accessed in admin API surface areas.
type MetaobjectAdminAccessInput string

const (
	// Applications that act on behalf of merchants can read metaobjects. Only the owning application can write metaobjects.
	MetaobjectAdminAccessInputMerchantRead MetaobjectAdminAccessInput = "MERCHANT_READ"
	// The owning application, as well as applications that act on behalf of merchants can read and write metaobjects. No other applications can read or write metaobjects.
	MetaobjectAdminAccessInputMerchantReadWrite MetaobjectAdminAccessInput = "MERCHANT_READ_WRITE"
)

// The input fields for creating a metaobject capability.
type MetaobjectCapabilityCreateInput struct {
	// The input for enabling the publishable capability.
	Publishable *MetaobjectCapabilityPublishableInput `json:"publishable,omitempty"`
	// The input for enabling the translatable capability.
	Translatable *MetaobjectCapabilityTranslatableInput `json:"translatable,omitempty"`
	// The input for enabling the renderable capability.
	Renderable *MetaobjectCapabilityRenderableInput `json:"renderable,omitempty"`
	// The input for enabling the Online Store capability.
	OnlineStore *MetaobjectCapabilityOnlineStoreInput `json:"onlineStore,omitempty"`
}

// GetPublishable returns MetaobjectCapabilityCreateInput.Publishable, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityCreateInput) GetPublishable() *MetaobjectCapabilityPublishableInput {
	return v.Publishable
}

// GetTranslatable returns MetaobjectCapabilityCreateInput.Translatable, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityCreateInput) GetTranslatable() *MetaobjectCapabilityTranslatableInput {
	return v.Translatable
}

// GetRenderable returns MetaobjectCapabilityCreateInput.Renderable, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityCreateInput) GetRenderable() *MetaobjectCapabilityRenderableInput {
	return v.Renderable
}

// GetOnlineStore returns MetaobjectCapabilityCreateInput.OnlineStore, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityCreateInput) GetOnlineStore() *MetaobjectCapabilityOnlineStoreInput {
	return v.OnlineStore
}

// The input fields of the Online Store capability.
type MetaobjectCapabilityDefinitionDataOnlineStoreInput struct {
	// The URL handle for accessing pages of this metaobject type in the Online Store.
	UrlHandle string `json:"urlHandle"`
	// Whether to redirect published metaobjects automatically when the URL handle changes.
	CreateRedirects *bool `json:"createRedirects"`
}

// GetUrlHandle returns MetaobjectCapabilityDefinitionDataOnlineStoreInput.UrlHandle, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityDefinitionDataOnlineStoreInput) GetUrlHandle() string {
	return v.UrlHandle
}

// GetCreateRedirects returns MetaobjectCapabilityDefinitionDataOnlineStoreInput.CreateRedirects, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityDefinitionDataOnlineStoreInput) GetCreateRedirects() *bool {
	return v.CreateRedirects
}

// The input fields of the renderable capability for SEO aliases.
type MetaobjectCapabilityDefinitionDataRenderableInput struct {
	// The metaobject field used as an alias for the SEO page title.
	MetaTitleKey *string `json:"metaTitleKey"`
	// The metaobject field used as an alias for the SEO page description.
	MetaDescriptionKey *string `json:"metaDescriptionKey"`
}

// GetMetaTitleKey returns MetaobjectCapabilityDefinitionDataRenderableInput.MetaTitleKey, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityDefinitionDataRenderableInput) GetMetaTitleKey() *string {
	return v.MetaTitleKey
}

// GetMetaDescriptionKey returns MetaobjectCapabilityDefinitionDataRenderableInput.MetaDescriptionKey, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityDefinitionDataRenderableInput) GetMetaDescriptionKey() *string {
	return v.MetaDescriptionKey
}

// The input fields for enabling and disabling the Online Store capability.
type MetaobjectCapabilityOnlineStoreInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
	// The data associated with the Online Store capability.
	Data *MetaobjectCapabilityDefinitionDataOnlineStoreInput `json:"data,omitempty"`
}

// GetEnabled returns MetaobjectCapabilityOnlineStoreInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityOnlineStoreInput) GetEnabled() bool { return v.Enabled }

// GetData returns MetaobjectCapabilityOnlineStoreInput.Data, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityOnlineStoreInput) GetData() *MetaobjectCapabilityDefinitionDataOnlineStoreInput {
	return v.Data
}

// The input fields for enabling and disabling the publishable capability.
type MetaobjectCapabilityPublishableInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns MetaobjectCapabilityPublishableInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityPublishableInput) GetEnabled() bool { return v.Enabled }

// The input fields for enabling and disabling the renderable capability.
type MetaobjectCapabilityRenderableInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
	// The data associated with the renderable capability.
	Data *MetaobjectCapabilityDefinitionDataRenderableInput `json:"data,omitempty"`
}

// GetEnabled returns MetaobjectCapabilityRenderableInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityRenderableInput) GetEnabled() bool { return v.Enabled }

// GetData returns MetaobjectCapabilityRenderableInput.Data, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityRenderableInput) GetData() *MetaobjectCapabilityDefinitionDataRenderableInput {
	return v.Data
}

// The input fields for enabling and disabling the translatable capability.
type MetaobjectCapabilityTranslatableInput struct {
	// Indicates whether the capability should be enabled or disabled.
	Enabled bool `json:"enabled"`
}

// GetEnabled returns MetaobjectCapabilityTranslatableInput.Enabled, and is useful for accessing the field via an interface.
func (v *MetaobjectCapabilityTranslatableInput) GetEnabled() bool { return v.Enabled }

// The input fields for creating a metaobject definition.
type MetaobjectDefinitionCreateInput struct {
	// A human-readable name for the definition. This can be changed at any time.
	Name *string `json:"name"`
	// An administrative description of the definition.
	Description *string `json:"description"`
	// The type of the metaobject definition. This can't be changed.  Must be 3-255 characters long and only contain alphanumeric, hyphen, and underscore characters.
	Type string `json:"type"`
	// A set of field definitions to create on this metaobject definition.
	FieldDefinitions []*MetaobjectFieldDefinitionCreateInput `json:"fieldDefinitions,omitempty"`
	// Access configuration for the metaobjects created with this definition.
	Access *MetaobjectAccessInput `json:"access,omitempty"`
	// The key of a field to reference as the display name for metaobjects of this type.
	DisplayNameKey *string `json:"displayNameKey"`
	// The capabilities of the metaobject definition.
	Capabilities *MetaobjectCapabilityCreateInput `json:"capabilities,omitempty"`
}

// GetName returns MetaobjectDefinitionCreateInput.Name, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetName() *string { return v.Name }

// GetDescription returns MetaobjectDefinitionCreateInput.Description, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetDescription() *string { return v.Description }

// GetType returns MetaobjectDefinitionCreateInput.Type, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetType() string { return v.Type }

// GetFieldDefinitions returns MetaobjectDefinitionCreateInput.FieldDefinitions, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetFieldDefinitions() []*MetaobjectFieldDefinitionCreateInput {
	return v.FieldDefinitions
}

// GetAccess returns MetaobjectDefinitionCreateInput.Access, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetAccess() *MetaobjectAccessInput { return v.Access }

// GetDisplayNameKey returns MetaobjectDefinitionCreateInput.DisplayNameKey, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetDisplayNameKey() *string { return v.DisplayNameKey }

// GetCapabilities returns MetaobjectDefinitionCreateInput.Capabilities, and is useful for accessing the field via an interface.
func (v *MetaobjectDefinitionCreateInput) GetCapabilities() *MetaobjectCapabilityCreateInput {
	return v.Capabilities
}

// The input fields for creating a metaobject field definition.
type MetaobjectFieldDefinitionCreateInput struct {
	// The key of the new field definition. This can't be changed.  Must be 2-64 characters long and only contain alphanumeric, hyphen, and underscore characters.
	Key string `json:"key"`
	// The metafield type applied to values of the field.
	Type string `json:"type"`
	// A human-readable name for the field. This can be changed at any time.
	Name *string `json:"name"`
	// An administrative description of the field.
	Description *string `json:"description"`
	// Whether metaobjects require a saved value for the field.
	Required *bool `json:"required"`
	// Custom validations that apply to values assigned to the field.
	Validations []*MetafieldDefinitionValidationInput `json:"validations,omitempty"`
}

// GetKey returns MetaobjectFieldDefinitionCreateInput.Key, and is useful for accessing the field via an interface.
func (v *MetaobjectFieldDefinitionCreateInput) GetKey() string { return v.Key }

// GetType returns MetaobjectFieldDefinitionCreateInput.Type, and is useful for accessing the field via an interface.
func (v *MetaobjectFieldDefinitionCreateInput) GetType() string { return v.Type }

// GetName returns MetaobjectFieldDefinitionCreateInput.Name, and is useful for accessing the field via an interface.
func (v *MetaobjectFieldDefinitionCreateInput) GetName() *string { return v.Name }

// GetDescription returns MetaobjectFieldDefinitionCreateInput.Description, and is useful for accessing the field via an interface.
func (v *MetaobjectFieldDefinitionCreateInput) GetDescription() *string { return v.Description }

// GetRequired returns MetaobjectFieldDefinitionCreateInput.Required, and is useful for accessing the field via an interface.
func (v *MetaobjectFieldDefinitionCreateInput) GetRequired() *bool { return v.Required }

// GetValidations returns MetaobjectFieldDefinitionCreateInput.Validations, and is useful for accessing the field via an interface.
func (v *MetaobjectFieldDefinitionCreateInput) GetValidations() []*MetafieldDefinitionValidationInput {
	return v.Validations
}

// Defines how the metaobjects of a definition can be accessed in Storefront API surface areas, including Liquid and the GraphQL Storefront API.
type MetaobjectStorefrontAccess string

const (
	// Metaobjects are not accessible in any Storefront API surface area.
	MetaobjectStorefrontAccessNone MetaobjectStorefrontAccess = "NONE"
	// Metaobjects are accessible in the GraphQL Storefront API by any application with the `unauthenticated_read_metaobjects` access scope. Metaobjects are accessible in online store Liquid templates.
	MetaobjectStorefrontAccessPublicRead MetaobjectStorefrontAccess = "PUBLIC_READ"
)

// Possible error codes that can be returned by `MetaobjectUserError`.
type MetaobjectUserErrorCode string

const (
	// The input value is invalid.
	MetaobjectUserErrorCodeInvalid MetaobjectUserErrorCode = "INVALID"
	// The input value isn't included in the list.
	MetaobjectUserErrorCodeInclusion MetaobjectUserErrorCode = "INCLUSION"
	// The input value is already taken.
	MetaobjectUserErrorCodeTaken MetaobjectUserErrorCode = "TAKEN"
	// The input value is too long.
	MetaobjectUserErrorCodeTooLong MetaobjectUserErrorCode = "TOO_LONG"
	// The input value is too short.
	MetaobjectUserErrorCodeTooShort MetaobjectUserErrorCode = "TOO_SHORT"
	// The input value needs to be blank.
	MetaobjectUserErrorCodePresent MetaobjectUserErrorCode = "PRESENT"
	// The input value is blank.
	MetaobjectUserErrorCodeBlank MetaobjectUserErrorCode = "BLANK"
	// The metafield type is invalid.
	MetaobjectUserErrorCodeInvalidType MetaobjectUserErrorCode = "INVALID_TYPE"
	// The value is invalid for the metafield type or the definition options.
	MetaobjectUserErrorCodeInvalidValue MetaobjectUserErrorCode = "INVALID_VALUE"
	// The value for the metafield definition option was invalid.
	MetaobjectUserErrorCodeInvalidOption MetaobjectUserErrorCode = "INVALID_OPTION"
	// Duplicate inputs were provided for this field key.
	MetaobjectUserErrorCodeDuplicateFieldInput MetaobjectUserErrorCode = "DUPLICATE_FIELD_INPUT"
	// No metaobject definition found for this type.
	MetaobjectUserErrorCodeUndefinedObjectType MetaobjectUserErrorCode = "UNDEFINED_OBJECT_TYPE"
	// No field definition found for this key.
	MetaobjectUserErrorCodeUndefinedObjectField MetaobjectUserErrorCode = "UNDEFINED_OBJECT_FIELD"
	// The specified field key is already in use.
	MetaobjectUserErrorCodeObjectFieldTaken MetaobjectUserErrorCode = "OBJECT_FIELD_TAKEN"
	// Missing required fields were found for this object.
	MetaobjectUserErrorCodeObjectFieldRequired MetaobjectUserErrorCode = "OBJECT_FIELD_REQUIRED"
	// The requested record couldn't be found.
	MetaobjectUserErrorCodeRecordNotFound MetaobjectUserErrorCode = "RECORD_NOT_FOUND"
	// An unexpected error occurred.
	MetaobjectUserErrorCodeInternalError MetaobjectUserErrorCode = "INTERNAL_ERROR"
	// The maximum number of metaobjects definitions has been exceeded.
	MetaobjectUserErrorCodeMaxDefinitionsExceeded MetaobjectUserErrorCode = "MAX_DEFINITIONS_EXCEEDED"
	// The maximum number of metaobjects per shop has been exceeded.
	MetaobjectUserErrorCodeMaxObjectsExceeded MetaobjectUserErrorCode = "MAX_OBJECTS_EXCEEDED"
	// The targeted object cannot be modified.
	MetaobjectUserErrorCodeImmutable MetaobjectUserErrorCode = "IMMUTABLE"
	// Not authorized.
	MetaobjectUserErrorCodeNotAuthorized MetaobjectUserErrorCode = "NOT_AUTHORIZED"
	// The provided name is reserved for system use.
	MetaobjectUserErrorCodeReservedName MetaobjectUserErrorCode = "RESERVED_NAME"
	// The display name cannot be the same when using the metaobject as a product option.
	MetaobjectUserErrorCodeDisplayNameConflict MetaobjectUserErrorCode = "DISPLAY_NAME_CONFLICT"
	// Admin access can only be specified on metaobject definitions that have an app-reserved type.
	MetaobjectUserErrorCodeAdminAccessInputNotAllowed MetaobjectUserErrorCode = "ADMIN_ACCESS_INPUT_NOT_ALLOWED"
	// The capability you are using is not enabled.
	MetaobjectUserErrorCodeCapabilityNotEnabled MetaobjectUserErrorCode = "CAPABILITY_NOT_ENABLED"
	// The Online Store URL handle is already taken.
	MetaobjectUserErrorCodeUrlHandleTaken MetaobjectUserErrorCode = "URL_HANDLE_TAKEN"
	// The Online Store URL handle is invalid.
	MetaobjectUserErrorCodeUrlHandleInvalid MetaobjectUserErrorCode = "URL_HANDLE_INVALID"
	// The Online Store URL handle cannot be blank.
	MetaobjectUserErrorCodeUrlHandleBlank MetaobjectUserErrorCode = "URL_HANDLE_BLANK"
	// Renderable data input is referencing an invalid field.
	MetaobjectUserErrorCodeFieldTypeInvalid MetaobjectUserErrorCode = "FIELD_TYPE_INVALID"
	// The input is missing required keys.
	MetaobjectUserErrorCodeMissingRequiredKeys MetaobjectUserErrorCode = "MISSING_REQUIRED_KEYS"
)

// __CreateMetafieldDefinitionInput is used internally by genqlient
type __CreateMetafieldDefinitionInput struct {
	Definition *MetafieldDefinitionInput `json:"definition,omitempty"`
}

// GetDefinition returns __CreateMetafieldDefinitionInput.Definition, and is useful for accessing the field via an interface.
func (v *__CreateMetafieldDefinitionInput) GetDefinition() *MetafieldDefinitionInput {
	return v.Definition
}

// __CreateMetaobjectDefinitionInput is used internally by genqlient
type __CreateMetaobjectDefinitionInput struct {
	Definition *MetaobjectDefinitionCreateInput `json:"definition,omitempty"`
}

// GetDefinition returns __CreateMetaobjectDefinitionInput.Definition, and is useful for accessing the field via an interface.
func (v *__CreateMetaobjectDefinitionInput) GetDefinition() *MetaobjectDefinitionCreateInput {
	return v.Definition
}

// __DeleteMetafieldDefinitionInput is used internally by genqlient
type __DeleteMetafieldDefinitionInput struct {
	Id string `json:"id"`
}

// GetId returns __DeleteMetafieldDefinitionInput.Id, and is useful for accessing the field via an interface.
func (v *__DeleteMetafieldDefinitionInput) GetId() string { return v.Id }

// __GetMetafieldDefinitionsInput is used internally by genqlient
type __GetMetafieldDefinitionsInput struct {
	OwnerType MetafieldOwnerType `json:"ownerType"`
	First     *int               `json:"first"`
}

// GetOwnerType returns __GetMetafieldDefinitionsInput.OwnerType, and is useful for accessing the field via an interface.
func (v *__GetMetafieldDefinitionsInput) GetOwnerType() MetafieldOwnerType { return v.OwnerType }

// GetFirst returns __GetMetafieldDefinitionsInput.First, and is useful for accessing the field via an interface.
func (v *__GetMetafieldDefinitionsInput) GetFirst() *int { return v.First }

// __GetMetaobjectDefinitionsInput is used internally by genqlient
type __GetMetaobjectDefinitionsInput struct {
	First int     `json:"first"`
	After *string `json:"after"`
}

// GetFirst returns __GetMetaobjectDefinitionsInput.First, and is useful for accessing the field via an interface.
func (v *__GetMetaobjectDefinitionsInput) GetFirst() int { return v.First }

// GetAfter returns __GetMetaobjectDefinitionsInput.After, and is useful for accessing the field via an interface.
func (v *__GetMetaobjectDefinitionsInput) GetAfter() *string { return v.After }

// The query or mutation executed by CreateMetafieldDefinition.
const CreateMetafieldDefinition_Operation = `
mutation CreateMetafieldDefinition ($definition: MetafieldDefinitionInput!) {
	metafieldDefinitionCreate(definition: $definition) {
		createdDefinition {
			id
			namespace
			key
			name
			description
			ownerType
			type {
				name
				category
				supportsDefinitionMigrations
				supportedValidations {
					name
					type
				}
			}
			validationStatus
			pinnedPosition
			useAsCollectionCondition
			metafieldsCount
			access {
				admin
				storefront
			}
			capabilities {
				smartCollectionCondition {
					enabled
				}
				adminFilterable {
					enabled
				}
			}
			validations {
				name
				value
			}
		}
		userErrors {
			field
			message
			code
		}
	}
}
`

func CreateMetafieldDefinition(
	ctx_ context.Context,
	client_ graphql.Client,
	definition *MetafieldDefinitionInput,
) (*CreateMetafieldDefinitionResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateMetafieldDefinition",
		Query:  CreateMetafieldDefinition_Operation,
		Variables: &__CreateMetafieldDefinitionInput{
			Definition: definition,
		},
	}
	var err_ error

	var data_ CreateMetafieldDefinitionResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by CreateMetaobjectDefinition.
const CreateMetaobjectDefinition_Operation = `
mutation CreateMetaobjectDefinition ($definition: MetaobjectDefinitionCreateInput!) {
	metaobjectDefinitionCreate(definition: $definition) {
		metaobjectDefinition {
			id
			name
			type
			fieldDefinitions {
				name
				key
			}
		}
		userErrors {
			field
			message
			code
		}
	}
}
`

func CreateMetaobjectDefinition(
	ctx_ context.Context,
	client_ graphql.Client,
	definition *MetaobjectDefinitionCreateInput,
) (*CreateMetaobjectDefinitionResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateMetaobjectDefinition",
		Query:  CreateMetaobjectDefinition_Operation,
		Variables: &__CreateMetaobjectDefinitionInput{
			Definition: definition,
		},
	}
	var err_ error

	var data_ CreateMetaobjectDefinitionResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by DeleteMetafieldDefinition.
const DeleteMetafieldDefinition_Operation = `
mutation DeleteMetafieldDefinition ($id: ID!) {
	metafieldDefinitionDelete(id: $id, deleteAllAssociatedMetafields: true) {
		deletedDefinitionId
		userErrors {
			field
			message
			code
		}
	}
}
`

func DeleteMetafieldDefinition(
	ctx_ context.Context,
	client_ graphql.Client,
	id string,
) (*DeleteMetafieldDefinitionResponse, error) {
	req_ := &graphql.Request{
		OpName: "DeleteMetafieldDefinition",
		Query:  DeleteMetafieldDefinition_Operation,
		Variables: &__DeleteMetafieldDefinitionInput{
			Id: id,
		},
	}
	var err_ error

	var data_ DeleteMetafieldDefinitionResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetMetafieldDefinitions.
const GetMetafieldDefinitions_Operation = `
query GetMetafieldDefinitions ($ownerType: MetafieldOwnerType!, $first: Int = 250) {
	metafieldDefinitions(first: $first, ownerType: $ownerType) {
		edges {
			node {
				id
				key
				name
				namespace
				description
				ownerType
				pinnedPosition
				type {
					category
					name
					supportedValidations {
						name
						type
					}
					supportsDefinitionMigrations
					valueType
				}
				access {
					admin
					customerAccount
					storefront
				}
				capabilities {
					adminFilterable {
						eligible
						enabled
						status
					}
					smartCollectionCondition {
						eligible
						enabled
					}
					uniqueValues {
						eligible
						enabled
					}
				}
				constraints {
					key
					values(first: 10) {
						edges {
							node {
								value
							}
						}
					}
				}
				metafieldsCount
				standardTemplate {
					description
					id
					key
					name
					namespace
					ownerTypes
					type {
						category
						name
						supportedValidations {
							name
							type
						}
						supportsDefinitionMigrations
						valueType
					}
					validations {
						name
						type
						value
					}
					visibleToStorefrontApi
				}
				useAsCollectionCondition
				validationStatus
				validations {
					name
					type
					value
				}
			}
		}
		pageInfo {
			hasNextPage
			endCursor
		}
	}
}
`

func GetMetafieldDefinitions(
	ctx_ context.Context,
	client_ graphql.Client,
	ownerType MetafieldOwnerType,
	first *int,
) (*GetMetafieldDefinitionsResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetMetafieldDefinitions",
		Query:  GetMetafieldDefinitions_Operation,
		Variables: &__GetMetafieldDefinitionsInput{
			OwnerType: ownerType,
			First:     first,
		},
	}
	var err_ error

	var data_ GetMetafieldDefinitionsResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by GetMetaobjectDefinitions.
const GetMetaobjectDefinitions_Operation = `
query GetMetaobjectDefinitions ($first: Int!, $after: String) {
	metaobjectDefinitions(first: $first, after: $after) {
		edges {
			cursor
			node {
				id
				name
				type
				description
				metaobjectsCount
			}
		}
		pageInfo {
			hasNextPage
			endCursor
		}
	}
}
`

func GetMetaobjectDefinitions(
	ctx_ context.Context,
	client_ graphql.Client,
	first int,
	after *string,
) (*GetMetaobjectDefinitionsResponse, error) {
	req_ := &graphql.Request{
		OpName: "GetMetaobjectDefinitions",
		Query:  GetMetaobjectDefinitions_Operation,
		Variables: &__GetMetaobjectDefinitionsInput{
			First: first,
			After: after,
		},
	}
	var err_ error

	var data_ GetMetaobjectDefinitionsResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
