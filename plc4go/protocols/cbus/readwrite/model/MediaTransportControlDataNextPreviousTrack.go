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

// MediaTransportControlDataNextPreviousTrack is the corresponding interface of MediaTransportControlDataNextPreviousTrack
type MediaTransportControlDataNextPreviousTrack interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsSetThePreviousTrack returns IsSetThePreviousTrack (virtual field)
	GetIsSetThePreviousTrack() bool
	// GetIsSetTheNextTrack returns IsSetTheNextTrack (virtual field)
	GetIsSetTheNextTrack() bool
	// IsMediaTransportControlDataNextPreviousTrack is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataNextPreviousTrack()
	// CreateBuilder creates a MediaTransportControlDataNextPreviousTrackBuilder
	CreateMediaTransportControlDataNextPreviousTrackBuilder() MediaTransportControlDataNextPreviousTrackBuilder
}

// _MediaTransportControlDataNextPreviousTrack is the data-structure of this message
type _MediaTransportControlDataNextPreviousTrack struct {
	MediaTransportControlDataContract
	Operation byte
}

var _ MediaTransportControlDataNextPreviousTrack = (*_MediaTransportControlDataNextPreviousTrack)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataNextPreviousTrack)(nil)

// NewMediaTransportControlDataNextPreviousTrack factory function for _MediaTransportControlDataNextPreviousTrack
func NewMediaTransportControlDataNextPreviousTrack(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, operation byte) *_MediaTransportControlDataNextPreviousTrack {
	_result := &_MediaTransportControlDataNextPreviousTrack{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		Operation:                         operation,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataNextPreviousTrackBuilder is a builder for MediaTransportControlDataNextPreviousTrack
type MediaTransportControlDataNextPreviousTrackBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(operation byte) MediaTransportControlDataNextPreviousTrackBuilder
	// WithOperation adds Operation (property field)
	WithOperation(byte) MediaTransportControlDataNextPreviousTrackBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() MediaTransportControlDataBuilder
	// Build builds the MediaTransportControlDataNextPreviousTrack or returns an error if something is wrong
	Build() (MediaTransportControlDataNextPreviousTrack, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlDataNextPreviousTrack
}

// NewMediaTransportControlDataNextPreviousTrackBuilder() creates a MediaTransportControlDataNextPreviousTrackBuilder
func NewMediaTransportControlDataNextPreviousTrackBuilder() MediaTransportControlDataNextPreviousTrackBuilder {
	return &_MediaTransportControlDataNextPreviousTrackBuilder{_MediaTransportControlDataNextPreviousTrack: new(_MediaTransportControlDataNextPreviousTrack)}
}

type _MediaTransportControlDataNextPreviousTrackBuilder struct {
	*_MediaTransportControlDataNextPreviousTrack

	parentBuilder *_MediaTransportControlDataBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataNextPreviousTrackBuilder) = (*_MediaTransportControlDataNextPreviousTrackBuilder)(nil)

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) setParent(contract MediaTransportControlDataContract) {
	b.MediaTransportControlDataContract = contract
	contract.(*_MediaTransportControlData)._SubType = b._MediaTransportControlDataNextPreviousTrack
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) WithMandatoryFields(operation byte) MediaTransportControlDataNextPreviousTrackBuilder {
	return b.WithOperation(operation)
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) WithOperation(operation byte) MediaTransportControlDataNextPreviousTrackBuilder {
	b.Operation = operation
	return b
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) Build() (MediaTransportControlDataNextPreviousTrack, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlDataNextPreviousTrack.deepCopy(), nil
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) MustBuild() MediaTransportControlDataNextPreviousTrack {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) Done() MediaTransportControlDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	}
	return b.parentBuilder
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) buildForMediaTransportControlData() (MediaTransportControlData, error) {
	return b.Build()
}

func (b *_MediaTransportControlDataNextPreviousTrackBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataNextPreviousTrackBuilder().(*_MediaTransportControlDataNextPreviousTrackBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataNextPreviousTrackBuilder creates a MediaTransportControlDataNextPreviousTrackBuilder
func (b *_MediaTransportControlDataNextPreviousTrack) CreateMediaTransportControlDataNextPreviousTrackBuilder() MediaTransportControlDataNextPreviousTrackBuilder {
	if b == nil {
		return NewMediaTransportControlDataNextPreviousTrackBuilder()
	}
	return &_MediaTransportControlDataNextPreviousTrackBuilder{_MediaTransportControlDataNextPreviousTrack: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataNextPreviousTrack) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousTrack) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousTrack) GetIsSetThePreviousTrack() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetIsSetTheNextTrack() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataNextPreviousTrack(structType any) MediaTransportControlDataNextPreviousTrack {
	if casted, ok := structType.(MediaTransportControlDataNextPreviousTrack); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataNextPreviousTrack); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetTypeName() string {
	return "MediaTransportControlDataNextPreviousTrack"
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataNextPreviousTrack) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataNextPreviousTrack) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataNextPreviousTrack MediaTransportControlDataNextPreviousTrack, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataNextPreviousTrack"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataNextPreviousTrack")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operation, err := ReadSimpleField(ctx, "operation", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operation' field"))
	}
	m.Operation = operation

	isSetThePreviousTrack, err := ReadVirtualField[bool](ctx, "isSetThePreviousTrack", (*bool)(nil), bool((operation) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetThePreviousTrack' field"))
	}
	_ = isSetThePreviousTrack

	isSetTheNextTrack, err := ReadVirtualField[bool](ctx, "isSetTheNextTrack", (*bool)(nil), bool((operation) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetTheNextTrack' field"))
	}
	_ = isSetTheNextTrack

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataNextPreviousTrack"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataNextPreviousTrack")
	}

	return m, nil
}

func (m *_MediaTransportControlDataNextPreviousTrack) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataNextPreviousTrack) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataNextPreviousTrack"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataNextPreviousTrack")
		}

		if err := WriteSimpleField[byte](ctx, "operation", m.GetOperation(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'operation' field")
		}
		// Virtual field
		isSetThePreviousTrack := m.GetIsSetThePreviousTrack()
		_ = isSetThePreviousTrack
		if _isSetThePreviousTrackErr := writeBuffer.WriteVirtual(ctx, "isSetThePreviousTrack", m.GetIsSetThePreviousTrack()); _isSetThePreviousTrackErr != nil {
			return errors.Wrap(_isSetThePreviousTrackErr, "Error serializing 'isSetThePreviousTrack' field")
		}
		// Virtual field
		isSetTheNextTrack := m.GetIsSetTheNextTrack()
		_ = isSetTheNextTrack
		if _isSetTheNextTrackErr := writeBuffer.WriteVirtual(ctx, "isSetTheNextTrack", m.GetIsSetTheNextTrack()); _isSetTheNextTrackErr != nil {
			return errors.Wrap(_isSetTheNextTrackErr, "Error serializing 'isSetTheNextTrack' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataNextPreviousTrack"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataNextPreviousTrack")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataNextPreviousTrack) IsMediaTransportControlDataNextPreviousTrack() {
}

func (m *_MediaTransportControlDataNextPreviousTrack) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlDataNextPreviousTrack) deepCopy() *_MediaTransportControlDataNextPreviousTrack {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataNextPreviousTrackCopy := &_MediaTransportControlDataNextPreviousTrack{
		m.MediaTransportControlDataContract.(*_MediaTransportControlData).deepCopy(),
		m.Operation,
	}
	_MediaTransportControlDataNextPreviousTrackCopy.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = m
	return _MediaTransportControlDataNextPreviousTrackCopy
}

func (m *_MediaTransportControlDataNextPreviousTrack) String() string {
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
