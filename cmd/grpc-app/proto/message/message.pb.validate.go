// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: message.proto

package message

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

// define the regex for a UUID once up-front
var _message_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Message with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Message) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Content

	// no validation rules for UserId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MessageValidationError is the validation error returned by Message.Validate
// if the designated constraints aren't met.
type MessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageValidationError) ErrorName() string { return "MessageValidationError" }

// Error satisfies the builtin error interface
func (e MessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageValidationError{}

// Validate checks the field values on MessageList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MessageList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageListValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MessageListValidationError is the validation error returned by
// MessageList.Validate if the designated constraints aren't met.
type MessageListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageListValidationError) ErrorName() string { return "MessageListValidationError" }

// Error satisfies the builtin error interface
func (e MessageListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageListValidationError{}

// Validate checks the field values on PostMessageParams with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PostMessageParams) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetContent()) > 1000 {
		return PostMessageParamsValidationError{
			field:  "Content",
			reason: "value length must be at most 1000 runes",
		}
	}

	return nil
}

// PostMessageParamsValidationError is the validation error returned by
// PostMessageParams.Validate if the designated constraints aren't met.
type PostMessageParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostMessageParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostMessageParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostMessageParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostMessageParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostMessageParamsValidationError) ErrorName() string {
	return "PostMessageParamsValidationError"
}

// Error satisfies the builtin error interface
func (e PostMessageParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostMessageParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostMessageParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostMessageParamsValidationError{}

// Validate checks the field values on PutMessageParams with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PutMessageParams) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if utf8.RuneCountInString(m.GetContent()) > 1000 {
		return PutMessageParamsValidationError{
			field:  "Content",
			reason: "value length must be at most 1000 runes",
		}
	}

	// no validation rules for UserId

	return nil
}

// PutMessageParamsValidationError is the validation error returned by
// PutMessageParams.Validate if the designated constraints aren't met.
type PutMessageParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PutMessageParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PutMessageParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PutMessageParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PutMessageParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PutMessageParamsValidationError) ErrorName() string { return "PutMessageParamsValidationError" }

// Error satisfies the builtin error interface
func (e PutMessageParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPutMessageParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PutMessageParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PutMessageParamsValidationError{}

// Validate checks the field values on DeleteMessageParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteMessageParams) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteMessageParamsValidationError is the validation error returned by
// DeleteMessageParams.Validate if the designated constraints aren't met.
type DeleteMessageParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMessageParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMessageParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMessageParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMessageParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMessageParamsValidationError) ErrorName() string {
	return "DeleteMessageParamsValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMessageParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMessageParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMessageParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMessageParamsValidationError{}
