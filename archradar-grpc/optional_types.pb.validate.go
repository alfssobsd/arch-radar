// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: optional_types.proto

package archradar_grpc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on DecimalOptional with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DecimalOptional) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Units

	// no validation rules for Nanos

	return nil
}

// DecimalOptionalValidationError is the validation error returned by
// DecimalOptional.Validate if the designated constraints aren't met.
type DecimalOptionalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecimalOptionalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecimalOptionalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecimalOptionalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecimalOptionalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecimalOptionalValidationError) ErrorName() string { return "DecimalOptionalValidationError" }

// Error satisfies the builtin error interface
func (e DecimalOptionalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecimalOptional.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecimalOptionalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecimalOptionalValidationError{}

// Validate checks the field values on UUIDOptional with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UUIDOptional) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Value

	return nil
}

// UUIDOptionalValidationError is the validation error returned by
// UUIDOptional.Validate if the designated constraints aren't met.
type UUIDOptionalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UUIDOptionalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UUIDOptionalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UUIDOptionalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UUIDOptionalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UUIDOptionalValidationError) ErrorName() string { return "UUIDOptionalValidationError" }

// Error satisfies the builtin error interface
func (e UUIDOptionalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUUIDOptional.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UUIDOptionalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UUIDOptionalValidationError{}