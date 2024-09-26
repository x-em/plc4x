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

// BACnetFaultParameterFaultOutOfRangeMinNormalValue is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValue
type BACnetFaultParameterFaultOutOfRangeMinNormalValue interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueContract
	BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsBACnetFaultParameterFaultOutOfRangeMinNormalValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMinNormalValue()
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueContract provides a set of functions which can be overwritten by a sub struct
type BACnetFaultParameterFaultOutOfRangeMinNormalValueContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetFaultParameterFaultOutOfRangeMinNormalValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMinNormalValue()
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetFaultParameterFaultOutOfRangeMinNormalValueRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValue is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValue struct {
	_SubType        BACnetFaultParameterFaultOutOfRangeMinNormalValue
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetFaultParameterFaultOutOfRangeMinNormalValueContract = (*_BACnetFaultParameterFaultOutOfRangeMinNormalValue)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValue factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValue
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetFaultParameterFaultOutOfRangeMinNormalValue must not be nil")
	}
	if peekedTagHeader == nil {
		panic("peekedTagHeader of type BACnetTagHeader for BACnetFaultParameterFaultOutOfRangeMinNormalValue must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetFaultParameterFaultOutOfRangeMinNormalValue must not be nil")
	}
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValue(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValue"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueParse[T BACnetFaultParameterFaultOutOfRangeMinNormalValue](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBufferProducer[T BACnetFaultParameterFaultOutOfRangeMinNormalValue](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBuffer[T BACnetFaultParameterFaultOutOfRangeMinNormalValue](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetFaultParameterFaultOutOfRangeMinNormalValue{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
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

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMinNormalValue BACnetFaultParameterFaultOutOfRangeMinNormalValue, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Validation
	if !(bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "only application tags allowed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetFaultParameterFaultOutOfRangeMinNormalValue
	switch {
	case peekedTagNumber == 0x4: // BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
		if _child, err = new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueReal).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultOutOfRangeMinNormalValueReal for type-switch of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
		}
	case peekedTagNumber == 0x2: // BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
		if _child, err = new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned for type-switch of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
		}
	case peekedTagNumber == 0x5: // BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
		if _child, err = new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble for type-switch of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
		}
	case peekedTagNumber == 0x3: // BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
		if _child, err = new(_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger for type-switch of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}

	return _child, nil
}

func (pm *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMinNormalValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) IsBACnetFaultParameterFaultOutOfRangeMinNormalValue() {
}
