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

// BACnetConstructedDataIntegerValueCOVIncrement is the corresponding interface of BACnetConstructedDataIntegerValueCOVIncrement
type BACnetConstructedDataIntegerValueCOVIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIntegerValueCOVIncrement is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueCOVIncrement()
}

// _BACnetConstructedDataIntegerValueCOVIncrement is the data-structure of this message
type _BACnetConstructedDataIntegerValueCOVIncrement struct {
	BACnetConstructedDataContract
	CovIncrement BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIntegerValueCOVIncrement = (*_BACnetConstructedDataIntegerValueCOVIncrement)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueCOVIncrement)(nil)

// NewBACnetConstructedDataIntegerValueCOVIncrement factory function for _BACnetConstructedDataIntegerValueCOVIncrement
func NewBACnetConstructedDataIntegerValueCOVIncrement(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, covIncrement BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueCOVIncrement {
	if covIncrement == nil {
		panic("covIncrement of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIntegerValueCOVIncrement must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueCOVIncrement{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovIncrement:                  covIncrement,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetCovIncrement() BACnetApplicationTagUnsignedInteger {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueCOVIncrement(structType any) BACnetConstructedDataIntegerValueCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueCOVIncrement"
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueCOVIncrement BACnetConstructedDataIntegerValueCOVIncrement, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covIncrement, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covIncrement", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	m.CovIncrement = covIncrement

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), covIncrement)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueCOVIncrement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueCOVIncrement")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "covIncrement", m.GetCovIncrement(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueCOVIncrement")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) IsBACnetConstructedDataIntegerValueCOVIncrement() {
}

func (m *_BACnetConstructedDataIntegerValueCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
