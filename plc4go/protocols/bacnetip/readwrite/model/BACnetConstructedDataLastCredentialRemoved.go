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

// BACnetConstructedDataLastCredentialRemoved is the corresponding interface of BACnetConstructedDataLastCredentialRemoved
type BACnetConstructedDataLastCredentialRemoved interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastCredentialRemoved returns LastCredentialRemoved (property field)
	GetLastCredentialRemoved() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataLastCredentialRemoved is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLastCredentialRemoved()
}

// _BACnetConstructedDataLastCredentialRemoved is the data-structure of this message
type _BACnetConstructedDataLastCredentialRemoved struct {
	BACnetConstructedDataContract
	LastCredentialRemoved BACnetDeviceObjectReference
}

var _ BACnetConstructedDataLastCredentialRemoved = (*_BACnetConstructedDataLastCredentialRemoved)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastCredentialRemoved)(nil)

// NewBACnetConstructedDataLastCredentialRemoved factory function for _BACnetConstructedDataLastCredentialRemoved
func NewBACnetConstructedDataLastCredentialRemoved(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, lastCredentialRemoved BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCredentialRemoved {
	if lastCredentialRemoved == nil {
		panic("lastCredentialRemoved of type BACnetDeviceObjectReference for BACnetConstructedDataLastCredentialRemoved must not be nil")
	}
	_result := &_BACnetConstructedDataLastCredentialRemoved{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastCredentialRemoved:         lastCredentialRemoved,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCredentialRemoved) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_REMOVED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) GetLastCredentialRemoved() BACnetDeviceObjectReference {
	return m.LastCredentialRemoved
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCredentialRemoved) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetLastCredentialRemoved())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCredentialRemoved(structType any) BACnetConstructedDataLastCredentialRemoved {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialRemoved); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialRemoved); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCredentialRemoved) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialRemoved"
}

func (m *_BACnetConstructedDataLastCredentialRemoved) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastCredentialRemoved)
	lengthInBits += m.LastCredentialRemoved.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCredentialRemoved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastCredentialRemoved) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastCredentialRemoved BACnetConstructedDataLastCredentialRemoved, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialRemoved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCredentialRemoved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastCredentialRemoved, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "lastCredentialRemoved", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastCredentialRemoved' field"))
	}
	m.LastCredentialRemoved = lastCredentialRemoved

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), lastCredentialRemoved)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialRemoved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCredentialRemoved")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastCredentialRemoved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCredentialRemoved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialRemoved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCredentialRemoved")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "lastCredentialRemoved", m.GetLastCredentialRemoved(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastCredentialRemoved' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialRemoved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCredentialRemoved")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCredentialRemoved) IsBACnetConstructedDataLastCredentialRemoved() {
}

func (m *_BACnetConstructedDataLastCredentialRemoved) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
