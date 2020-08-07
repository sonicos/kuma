// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: system/v1alpha1/zone_insight.proto

package v1alpha1

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
var _zone_insight_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ZoneInsight with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ZoneInsight) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSubscriptions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ZoneInsightValidationError{
					field:  fmt.Sprintf("Subscriptions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ZoneInsightValidationError is the validation error returned by
// ZoneInsight.Validate if the designated constraints aren't met.
type ZoneInsightValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZoneInsightValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZoneInsightValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZoneInsightValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZoneInsightValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZoneInsightValidationError) ErrorName() string { return "ZoneInsightValidationError" }

// Error satisfies the builtin error interface
func (e ZoneInsightValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZoneInsight.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZoneInsightValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZoneInsightValidationError{}

// Validate checks the field values on KDSSubscription with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *KDSSubscription) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 1 {
		return KDSSubscriptionValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetGlobalInstanceId()) < 1 {
		return KDSSubscriptionValidationError{
			field:  "GlobalInstanceId",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetConnectTime() == nil {
		return KDSSubscriptionValidationError{
			field:  "ConnectTime",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDisconnectTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KDSSubscriptionValidationError{
				field:  "DisconnectTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetStatus() == nil {
		return KDSSubscriptionValidationError{
			field:  "Status",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KDSSubscriptionValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// KDSSubscriptionValidationError is the validation error returned by
// KDSSubscription.Validate if the designated constraints aren't met.
type KDSSubscriptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KDSSubscriptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KDSSubscriptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KDSSubscriptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KDSSubscriptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KDSSubscriptionValidationError) ErrorName() string { return "KDSSubscriptionValidationError" }

// Error satisfies the builtin error interface
func (e KDSSubscriptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKDSSubscription.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KDSSubscriptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KDSSubscriptionValidationError{}

// Validate checks the field values on KDSSubscriptionStatus with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *KDSSubscriptionStatus) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLastUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KDSSubscriptionStatusValidationError{
				field:  "LastUpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTotal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KDSSubscriptionStatusValidationError{
				field:  "Total",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for key, val := range m.GetStat() {
		_ = val

		// no validation rules for Stat[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return KDSSubscriptionStatusValidationError{
					field:  fmt.Sprintf("Stat[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// KDSSubscriptionStatusValidationError is the validation error returned by
// KDSSubscriptionStatus.Validate if the designated constraints aren't met.
type KDSSubscriptionStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KDSSubscriptionStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KDSSubscriptionStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KDSSubscriptionStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KDSSubscriptionStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KDSSubscriptionStatusValidationError) ErrorName() string {
	return "KDSSubscriptionStatusValidationError"
}

// Error satisfies the builtin error interface
func (e KDSSubscriptionStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKDSSubscriptionStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KDSSubscriptionStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KDSSubscriptionStatusValidationError{}

// Validate checks the field values on KDSServiceStats with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *KDSServiceStats) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResponsesSent

	// no validation rules for ResponsesAcknowledged

	// no validation rules for ResponsesRejected

	return nil
}

// KDSServiceStatsValidationError is the validation error returned by
// KDSServiceStats.Validate if the designated constraints aren't met.
type KDSServiceStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KDSServiceStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KDSServiceStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KDSServiceStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KDSServiceStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KDSServiceStatsValidationError) ErrorName() string { return "KDSServiceStatsValidationError" }

// Error satisfies the builtin error interface
func (e KDSServiceStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKDSServiceStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KDSServiceStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KDSServiceStatsValidationError{}