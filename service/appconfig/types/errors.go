// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The input fails to satisfy the constraints specified by an AWS service.
type BadRequestException struct {
	Message *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request could not be processed because of conflict in the current state of
// the resource.
type ConflictException struct {
	Message *string
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

// There was an internal failure in the AppConfig service.
type InternalServerException struct {
	Message *string
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string             { return "InternalServerException" }
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The configuration size is too large.
type PayloadTooLargeException struct {
	Message *string

	Size    *float32
	Measure BytesMeasure
	Limit   *float32
}

func (e *PayloadTooLargeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PayloadTooLargeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PayloadTooLargeException) ErrorCode() string             { return "PayloadTooLargeException" }
func (e *PayloadTooLargeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested resource could not be found.
type ResourceNotFoundException struct {
	Message *string

	ResourceName *string
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

// The number of hosted configuration versions exceeds the limit for the AppConfig
// configuration store. Delete one or more versions and try again.
type ServiceQuotaExceededException struct {
	Message *string
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string             { return "ServiceQuotaExceededException" }
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
