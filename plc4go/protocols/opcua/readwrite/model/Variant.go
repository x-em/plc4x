/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Variant is the corresponding interface of Variant
type Variant interface {
	VariantContract
	VariantRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsVariant is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariant()
}

// VariantContract provides a set of functions which can be overwritten by a sub struct
type VariantContract interface {
	// GetArrayLengthSpecified returns ArrayLengthSpecified (property field)
	GetArrayLengthSpecified() bool
	// GetArrayDimensionsSpecified returns ArrayDimensionsSpecified (property field)
	GetArrayDimensionsSpecified() bool
	// GetNoOfArrayDimensions returns NoOfArrayDimensions (property field)
	GetNoOfArrayDimensions() *int32
	// GetArrayDimensions returns ArrayDimensions (property field)
	GetArrayDimensions() []bool
	// IsVariant is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariant()
}

// VariantRequirements provides a set of functions which need to be implemented by a sub struct
type VariantRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetVariantType returns VariantType (discriminator field)
	GetVariantType() uint8
	// GetArrayLengthSpecified returns ArrayLengthSpecified (discriminator field)
	GetArrayLengthSpecified() bool
}

// _Variant is the data-structure of this message
type _Variant struct {
	_SubType                 Variant
	ArrayLengthSpecified     bool
	ArrayDimensionsSpecified bool
	NoOfArrayDimensions      *int32
	ArrayDimensions          []bool
}

var _ VariantContract = (*_Variant)(nil)

// NewVariant factory function for _Variant
func NewVariant(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_Variant {
	return &_Variant{ArrayLengthSpecified: arrayLengthSpecified, ArrayDimensionsSpecified: arrayDimensionsSpecified, NoOfArrayDimensions: noOfArrayDimensions, ArrayDimensions: arrayDimensions}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Variant) GetArrayLengthSpecified() bool {
	return m.ArrayLengthSpecified
}

func (m *_Variant) GetArrayDimensionsSpecified() bool {
	return m.ArrayDimensionsSpecified
}

func (m *_Variant) GetNoOfArrayDimensions() *int32 {
	return m.NoOfArrayDimensions
}

func (m *_Variant) GetArrayDimensions() []bool {
	return m.ArrayDimensions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariant(structType any) Variant {
	if casted, ok := structType.(Variant); ok {
		return casted
	}
	if casted, ok := structType.(*Variant); ok {
		return *casted
	}
	return nil
}

func (m *_Variant) GetTypeName() string {
	return "Variant"
}

func (m *_Variant) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (arrayLengthSpecified)
	lengthInBits += 1

	// Simple field (arrayDimensionsSpecified)
	lengthInBits += 1
	// Discriminator Field (VariantType)
	lengthInBits += 6

	// Optional Field (noOfArrayDimensions)
	if m.NoOfArrayDimensions != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.ArrayDimensions) > 0 {
		lengthInBits += 1 * uint16(len(m.ArrayDimensions))
	}

	return lengthInBits
}

func (m *_Variant) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func VariantParse[T Variant](ctx context.Context, theBytes []byte) (T, error) {
	return VariantParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func VariantParseWithBufferProducer[T Variant]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := VariantParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func VariantParseWithBuffer[T Variant](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_Variant{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_Variant) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__variant Variant, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Variant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Variant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	arrayLengthSpecified, err := ReadSimpleField(ctx, "arrayLengthSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLengthSpecified' field"))
	}
	m.ArrayLengthSpecified = arrayLengthSpecified

	arrayDimensionsSpecified, err := ReadSimpleField(ctx, "arrayDimensionsSpecified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensionsSpecified' field"))
	}
	m.ArrayDimensionsSpecified = arrayDimensionsSpecified

	VariantType, err := ReadDiscriminatorField[uint8](ctx, "VariantType", ReadUnsignedByte(readBuffer, uint8(6)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'VariantType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child Variant
	switch {
	case VariantType == uint8(0): // VariantNull
		if _child, err = new(_VariantNull).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantNull for type-switch of Variant")
		}
	case VariantType == uint8(1): // VariantBoolean
		if _child, err = new(_VariantBoolean).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantBoolean for type-switch of Variant")
		}
	case VariantType == uint8(2): // VariantSByte
		if _child, err = new(_VariantSByte).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantSByte for type-switch of Variant")
		}
	case VariantType == uint8(3): // VariantByte
		if _child, err = new(_VariantByte).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantByte for type-switch of Variant")
		}
	case VariantType == uint8(4): // VariantInt16
		if _child, err = new(_VariantInt16).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantInt16 for type-switch of Variant")
		}
	case VariantType == uint8(5): // VariantUInt16
		if _child, err = new(_VariantUInt16).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantUInt16 for type-switch of Variant")
		}
	case VariantType == uint8(6): // VariantInt32
		if _child, err = new(_VariantInt32).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantInt32 for type-switch of Variant")
		}
	case VariantType == uint8(7): // VariantUInt32
		if _child, err = new(_VariantUInt32).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantUInt32 for type-switch of Variant")
		}
	case VariantType == uint8(8): // VariantInt64
		if _child, err = new(_VariantInt64).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantInt64 for type-switch of Variant")
		}
	case VariantType == uint8(9): // VariantUInt64
		if _child, err = new(_VariantUInt64).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantUInt64 for type-switch of Variant")
		}
	case VariantType == uint8(10): // VariantFloat
		if _child, err = new(_VariantFloat).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantFloat for type-switch of Variant")
		}
	case VariantType == uint8(11): // VariantDouble
		if _child, err = new(_VariantDouble).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantDouble for type-switch of Variant")
		}
	case VariantType == uint8(12): // VariantString
		if _child, err = new(_VariantString).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantString for type-switch of Variant")
		}
	case VariantType == uint8(13): // VariantDateTime
		if _child, err = new(_VariantDateTime).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantDateTime for type-switch of Variant")
		}
	case VariantType == uint8(14): // VariantGuid
		if _child, err = new(_VariantGuid).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantGuid for type-switch of Variant")
		}
	case VariantType == uint8(15): // VariantByteString
		if _child, err = new(_VariantByteString).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantByteString for type-switch of Variant")
		}
	case VariantType == uint8(16): // VariantXmlElement
		if _child, err = new(_VariantXmlElement).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantXmlElement for type-switch of Variant")
		}
	case VariantType == uint8(17): // VariantNodeId
		if _child, err = new(_VariantNodeId).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantNodeId for type-switch of Variant")
		}
	case VariantType == uint8(18): // VariantExpandedNodeId
		if _child, err = new(_VariantExpandedNodeId).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantExpandedNodeId for type-switch of Variant")
		}
	case VariantType == uint8(19): // VariantStatusCode
		if _child, err = new(_VariantStatusCode).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantStatusCode for type-switch of Variant")
		}
	case VariantType == uint8(20): // VariantQualifiedName
		if _child, err = new(_VariantQualifiedName).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantQualifiedName for type-switch of Variant")
		}
	case VariantType == uint8(21): // VariantLocalizedText
		if _child, err = new(_VariantLocalizedText).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantLocalizedText for type-switch of Variant")
		}
	case VariantType == uint8(22): // VariantExtensionObject
		if _child, err = new(_VariantExtensionObject).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantExtensionObject for type-switch of Variant")
		}
	case VariantType == uint8(23): // VariantDataValue
		if _child, err = new(_VariantDataValue).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantDataValue for type-switch of Variant")
		}
	case VariantType == uint8(24): // VariantVariant
		if _child, err = new(_VariantVariant).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantVariant for type-switch of Variant")
		}
	case VariantType == uint8(25): // VariantDiagnosticInfo
		if _child, err = new(_VariantDiagnosticInfo).parse(ctx, readBuffer, m, arrayLengthSpecified); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type VariantDiagnosticInfo for type-switch of Variant")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [VariantType=%v, arrayLengthSpecified=%v]", VariantType, arrayLengthSpecified)
	}

	var noOfArrayDimensions *int32
	noOfArrayDimensions, err = ReadOptionalField[int32](ctx, "noOfArrayDimensions", ReadSignedInt(readBuffer, uint8(32)), arrayDimensionsSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfArrayDimensions' field"))
	}
	m.NoOfArrayDimensions = noOfArrayDimensions

	arrayDimensions, err := ReadCountArrayField[bool](ctx, "arrayDimensions", ReadBoolean(readBuffer), uint64(utils.InlineIf(bool((noOfArrayDimensions) == (nil)), func() any { return int32(int32(0)) }, func() any { return int32((*noOfArrayDimensions)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayDimensions' field"))
	}
	m.ArrayDimensions = arrayDimensions

	if closeErr := readBuffer.CloseContext("Variant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Variant")
	}

	return _child, nil
}

func (pm *_Variant) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Variant, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Variant"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Variant")
	}

	if err := WriteSimpleField[bool](ctx, "arrayLengthSpecified", m.GetArrayLengthSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayLengthSpecified' field")
	}

	if err := WriteSimpleField[bool](ctx, "arrayDimensionsSpecified", m.GetArrayDimensionsSpecified(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayDimensionsSpecified' field")
	}

	if err := WriteDiscriminatorField(ctx, "VariantType", m.GetVariantType(), WriteUnsignedByte(writeBuffer, 6)); err != nil {
		return errors.Wrap(err, "Error serializing 'VariantType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteOptionalField[int32](ctx, "noOfArrayDimensions", m.GetNoOfArrayDimensions(), WriteSignedInt(writeBuffer, 32), true); err != nil {
		return errors.Wrap(err, "Error serializing 'noOfArrayDimensions' field")
	}

	if err := WriteSimpleTypeArrayField(ctx, "arrayDimensions", m.GetArrayDimensions(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'arrayDimensions' field")
	}

	if popErr := writeBuffer.PopContext("Variant"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Variant")
	}
	return nil
}

func (m *_Variant) IsVariant() {}
