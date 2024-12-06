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

// BACnetConstructedDataTrendLogLogBuffer is the corresponding interface of BACnetConstructedDataTrendLogLogBuffer
type BACnetConstructedDataTrendLogLogBuffer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetLogRecord
	// IsBACnetConstructedDataTrendLogLogBuffer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTrendLogLogBuffer()
	// CreateBuilder creates a BACnetConstructedDataTrendLogLogBufferBuilder
	CreateBACnetConstructedDataTrendLogLogBufferBuilder() BACnetConstructedDataTrendLogLogBufferBuilder
}

// _BACnetConstructedDataTrendLogLogBuffer is the data-structure of this message
type _BACnetConstructedDataTrendLogLogBuffer struct {
	BACnetConstructedDataContract
	FloorText []BACnetLogRecord
}

var _ BACnetConstructedDataTrendLogLogBuffer = (*_BACnetConstructedDataTrendLogLogBuffer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTrendLogLogBuffer)(nil)

// NewBACnetConstructedDataTrendLogLogBuffer factory function for _BACnetConstructedDataTrendLogLogBuffer
func NewBACnetConstructedDataTrendLogLogBuffer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, floorText []BACnetLogRecord, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogLogBuffer {
	_result := &_BACnetConstructedDataTrendLogLogBuffer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FloorText:                     floorText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTrendLogLogBufferBuilder is a builder for BACnetConstructedDataTrendLogLogBuffer
type BACnetConstructedDataTrendLogLogBufferBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(floorText []BACnetLogRecord) BACnetConstructedDataTrendLogLogBufferBuilder
	// WithFloorText adds FloorText (property field)
	WithFloorText(...BACnetLogRecord) BACnetConstructedDataTrendLogLogBufferBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataTrendLogLogBuffer or returns an error if something is wrong
	Build() (BACnetConstructedDataTrendLogLogBuffer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTrendLogLogBuffer
}

// NewBACnetConstructedDataTrendLogLogBufferBuilder() creates a BACnetConstructedDataTrendLogLogBufferBuilder
func NewBACnetConstructedDataTrendLogLogBufferBuilder() BACnetConstructedDataTrendLogLogBufferBuilder {
	return &_BACnetConstructedDataTrendLogLogBufferBuilder{_BACnetConstructedDataTrendLogLogBuffer: new(_BACnetConstructedDataTrendLogLogBuffer)}
}

type _BACnetConstructedDataTrendLogLogBufferBuilder struct {
	*_BACnetConstructedDataTrendLogLogBuffer

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTrendLogLogBufferBuilder) = (*_BACnetConstructedDataTrendLogLogBufferBuilder)(nil)

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataTrendLogLogBuffer
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) WithMandatoryFields(floorText []BACnetLogRecord) BACnetConstructedDataTrendLogLogBufferBuilder {
	return b.WithFloorText(floorText...)
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) WithFloorText(floorText ...BACnetLogRecord) BACnetConstructedDataTrendLogLogBufferBuilder {
	b.FloorText = floorText
	return b
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) Build() (BACnetConstructedDataTrendLogLogBuffer, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTrendLogLogBuffer.deepCopy(), nil
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) MustBuild() BACnetConstructedDataTrendLogLogBuffer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTrendLogLogBufferBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTrendLogLogBufferBuilder().(*_BACnetConstructedDataTrendLogLogBufferBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTrendLogLogBufferBuilder creates a BACnetConstructedDataTrendLogLogBufferBuilder
func (b *_BACnetConstructedDataTrendLogLogBuffer) CreateBACnetConstructedDataTrendLogLogBufferBuilder() BACnetConstructedDataTrendLogLogBufferBuilder {
	if b == nil {
		return NewBACnetConstructedDataTrendLogLogBufferBuilder()
	}
	return &_BACnetConstructedDataTrendLogLogBufferBuilder{_BACnetConstructedDataTrendLogLogBuffer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetFloorText() []BACnetLogRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogLogBuffer(structType any) BACnetConstructedDataTrendLogLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataTrendLogLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataTrendLogLogBuffer"
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTrendLogLogBuffer BACnetConstructedDataTrendLogLogBuffer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floorText, err := ReadTerminatedArrayField[BACnetLogRecord](ctx, "floorText", ReadComplex[BACnetLogRecord](BACnetLogRecordParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorText' field"))
	}
	m.FloorText = floorText

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogLogBuffer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogLogBuffer")
		}

		if err := WriteComplexTypeArrayField(ctx, "floorText", m.GetFloorText(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'floorText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogLogBuffer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) IsBACnetConstructedDataTrendLogLogBuffer() {}

func (m *_BACnetConstructedDataTrendLogLogBuffer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) deepCopy() *_BACnetConstructedDataTrendLogLogBuffer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTrendLogLogBufferCopy := &_BACnetConstructedDataTrendLogLogBuffer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetLogRecord, BACnetLogRecord](m.FloorText),
	}
	_BACnetConstructedDataTrendLogLogBufferCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTrendLogLogBufferCopy
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}