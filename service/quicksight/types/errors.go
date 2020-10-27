// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// You don't have access to this item. The provided credentials couldn't be
// validated. You might not be authorized to carry out the request. Make sure that
// your account is authorized to use the Amazon QuickSight service, that your
// policies have the correct permissions, and that you are using the correct access
// keys.
type AccessDeniedException struct {
	Message *string

	RequestId *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource is already in a state that indicates an operation is happening that
// must complete before a new update can be applied.
type ConcurrentUpdatingException struct {
	Message *string

	RequestId *string
}

func (e *ConcurrentUpdatingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentUpdatingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentUpdatingException) ErrorCode() string             { return "ConcurrentUpdatingException" }
func (e *ConcurrentUpdatingException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Updating or deleting a resource can cause an inconsistent state.
type ConflictException struct {
	Message *string

	RequestId *string
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The domain specified isn't on the allow list. All domains for embedded
// dashboards must be added to the approved list by an Amazon QuickSight admin.
type DomainNotWhitelistedException struct {
	Message *string

	RequestId *string
}

func (e *DomainNotWhitelistedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DomainNotWhitelistedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DomainNotWhitelistedException) ErrorCode() string             { return "DomainNotWhitelistedException" }
func (e *DomainNotWhitelistedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The identity type specified isn't supported. Supported identity types include
// IAM and QUICKSIGHT.
type IdentityTypeNotSupportedException struct {
	Message *string

	RequestId *string
}

func (e *IdentityTypeNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdentityTypeNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdentityTypeNotSupportedException) ErrorCode() string {
	return "IdentityTypeNotSupportedException"
}
func (e *IdentityTypeNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal failure occurred.
type InternalFailureException struct {
	Message *string

	RequestId *string
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string             { return "InternalFailureException" }
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The NextToken value isn't valid.
type InvalidNextTokenException struct {
	Message *string

	RequestId *string
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more parameters has a value that isn't valid.
type InvalidParameterValueException struct {
	Message *string

	RequestId *string
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string             { return "InvalidParameterValueException" }
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A limit is exceeded.
type LimitExceededException struct {
	Message *string

	ResourceType ExceptionResourceType
	RequestId    *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more preconditions aren't met.
type PreconditionNotMetException struct {
	Message *string

	RequestId *string
}

func (e *PreconditionNotMetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionNotMetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionNotMetException) ErrorCode() string             { return "PreconditionNotMetException" }
func (e *PreconditionNotMetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user with the provided name isn't found. This error can happen in any
// operation that requires finding a user based on a provided user name, such as
// DeleteUser, DescribeUser, and so on.
type QuickSightUserNotFoundException struct {
	Message *string

	RequestId *string
}

func (e *QuickSightUserNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QuickSightUserNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QuickSightUserNotFoundException) ErrorCode() string {
	return "QuickSightUserNotFoundException"
}
func (e *QuickSightUserNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified already exists.
type ResourceExistsException struct {
	Message *string

	ResourceType ExceptionResourceType
	RequestId    *string
}

func (e *ResourceExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceExistsException) ErrorCode() string             { return "ResourceExistsException" }
func (e *ResourceExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more resources can't be found.
type ResourceNotFoundException struct {
	Message *string

	ResourceType ExceptionResourceType
	RequestId    *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This resource is currently unavailable.
type ResourceUnavailableException struct {
	Message *string

	RequestId    *string
	ResourceType ExceptionResourceType
}

func (e *ResourceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceUnavailableException) ErrorCode() string             { return "ResourceUnavailableException" }
func (e *ResourceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The number of minutes specified for the lifetime of a session isn't valid. The
// session lifetime must be 15-600 minutes.
type SessionLifetimeInMinutesInvalidException struct {
	Message *string

	RequestId *string
}

func (e *SessionLifetimeInMinutesInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorCode() string {
	return "SessionLifetimeInMinutesInvalidException"
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Access is throttled.
type ThrottlingException struct {
	Message *string

	RequestId *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error indicates that you are calling an operation on an Amazon QuickSight
// subscription where the edition doesn't include support for that operation.
// Amazon QuickSight currently has Standard Edition and Enterprise Edition. Not
// every operation and capability is available in every edition.
type UnsupportedUserEditionException struct {
	Message *string

	RequestId *string
}

func (e *UnsupportedUserEditionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedUserEditionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedUserEditionException) ErrorCode() string {
	return "UnsupportedUserEditionException"
}
func (e *UnsupportedUserEditionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
