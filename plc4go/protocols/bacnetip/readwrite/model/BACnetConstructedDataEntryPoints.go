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

// BACnetConstructedDataEntryPoints is the corresponding interface of BACnetConstructedDataEntryPoints
type BACnetConstructedDataEntryPoints interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEntryPoints returns EntryPoints (property field)
	GetEntryPoints() []BACnetDeviceObjectReference
	// IsBACnetConstructedDataEntryPoints is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEntryPoints()
}

// _BACnetConstructedDataEntryPoints is the data-structure of this message
type _BACnetConstructedDataEntryPoints struct {
	BACnetConstructedDataContract
	EntryPoints []BACnetDeviceObjectReference
}

var _ BACnetConstructedDataEntryPoints = (*_BACnetConstructedDataEntryPoints)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEntryPoints)(nil)

// NewBACnetConstructedDataEntryPoints factory function for _BACnetConstructedDataEntryPoints
func NewBACnetConstructedDataEntryPoints(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, entryPoints []BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEntryPoints {
	_result := &_BACnetConstructedDataEntryPoints{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EntryPoints:                   entryPoints,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEntryPoints) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEntryPoints) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENTRY_POINTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEntryPoints) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEntryPoints) GetEntryPoints() []BACnetDeviceObjectReference {
	return m.EntryPoints
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEntryPoints(structType any) BACnetConstructedDataEntryPoints {
	if casted, ok := structType.(BACnetConstructedDataEntryPoints); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEntryPoints); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEntryPoints) GetTypeName() string {
	return "BACnetConstructedDataEntryPoints"
}

func (m *_BACnetConstructedDataEntryPoints) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.EntryPoints) > 0 {
		for _, element := range m.EntryPoints {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataEntryPoints) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEntryPoints) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEntryPoints BACnetConstructedDataEntryPoints, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEntryPoints"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEntryPoints")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	entryPoints, err := ReadTerminatedArrayField[BACnetDeviceObjectReference](ctx, "entryPoints", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entryPoints' field"))
	}
	m.EntryPoints = entryPoints

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEntryPoints"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEntryPoints")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEntryPoints) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEntryPoints) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEntryPoints"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEntryPoints")
		}

		if err := WriteComplexTypeArrayField(ctx, "entryPoints", m.GetEntryPoints(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'entryPoints' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEntryPoints"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEntryPoints")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEntryPoints) IsBACnetConstructedDataEntryPoints() {}

func (m *_BACnetConstructedDataEntryPoints) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
