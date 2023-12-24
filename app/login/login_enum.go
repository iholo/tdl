// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.8
// Revision: 3d844c8ecc59661ed7aa17bfd65727bc06a60ad8
// Build Date: 2023-09-18T14:55:21Z
// Built By: goreleaser

package login

import (
	"fmt"
	"strings"
)

const (
	// TypeDesktop is a Type of type Desktop.
	TypeDesktop Type = iota
	// TypeCode is a Type of type Code.
	TypeCode
	// TypeQr is a Type of type Qr.
	TypeQr
)

var ErrInvalidType = fmt.Errorf("not a valid Type, try [%s]", strings.Join(_TypeNames, ", "))

const _TypeName = "desktopcodeqr"

var _TypeNames = []string{
	_TypeName[0:7],
	_TypeName[7:11],
	_TypeName[11:13],
}

// TypeNames returns a list of possible string values of Type.
func TypeNames() []string {
	tmp := make([]string, len(_TypeNames))
	copy(tmp, _TypeNames)
	return tmp
}

// TypeValues returns a list of the values for Type
func TypeValues() []Type {
	return []Type{
		TypeDesktop,
		TypeCode,
		TypeQr,
	}
}

var _TypeMap = map[Type]string{
	TypeDesktop: _TypeName[0:7],
	TypeCode:    _TypeName[7:11],
	TypeQr:      _TypeName[11:13],
}

// String implements the Stringer interface.
func (x Type) String() string {
	if str, ok := _TypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Type(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x Type) IsValid() bool {
	_, ok := _TypeMap[x]
	return ok
}

var _TypeValue = map[string]Type{
	_TypeName[0:7]:                    TypeDesktop,
	strings.ToLower(_TypeName[0:7]):   TypeDesktop,
	_TypeName[7:11]:                   TypeCode,
	strings.ToLower(_TypeName[7:11]):  TypeCode,
	_TypeName[11:13]:                  TypeQr,
	strings.ToLower(_TypeName[11:13]): TypeQr,
}

// ParseType attempts to convert a string to a Type.
func ParseType(name string) (Type, error) {
	if x, ok := _TypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _TypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Type(0), fmt.Errorf("%s is %w", name, ErrInvalidType)
}

// Set implements the Golang flag.Value interface func.
func (x *Type) Set(val string) error {
	v, err := ParseType(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func.
func (x *Type) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface.
func (x *Type) Type() string {
	return "Type"
}