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

// BACnetTagPayloadDouble is the corresponding interface of BACnetTagPayloadDouble
type BACnetTagPayloadDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetValue returns Value (property field)
	GetValue() float64
	// IsBACnetTagPayloadDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTagPayloadDouble()
	// CreateBuilder creates a BACnetTagPayloadDoubleBuilder
	CreateBACnetTagPayloadDoubleBuilder() BACnetTagPayloadDoubleBuilder
}

// _BACnetTagPayloadDouble is the data-structure of this message
type _BACnetTagPayloadDouble struct {
	Value float64
}

var _ BACnetTagPayloadDouble = (*_BACnetTagPayloadDouble)(nil)

// NewBACnetTagPayloadDouble factory function for _BACnetTagPayloadDouble
func NewBACnetTagPayloadDouble(value float64) *_BACnetTagPayloadDouble {
	return &_BACnetTagPayloadDouble{Value: value}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTagPayloadDoubleBuilder is a builder for BACnetTagPayloadDouble
type BACnetTagPayloadDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value float64) BACnetTagPayloadDoubleBuilder
	// WithValue adds Value (property field)
	WithValue(float64) BACnetTagPayloadDoubleBuilder
	// Build builds the BACnetTagPayloadDouble or returns an error if something is wrong
	Build() (BACnetTagPayloadDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTagPayloadDouble
}

// NewBACnetTagPayloadDoubleBuilder() creates a BACnetTagPayloadDoubleBuilder
func NewBACnetTagPayloadDoubleBuilder() BACnetTagPayloadDoubleBuilder {
	return &_BACnetTagPayloadDoubleBuilder{_BACnetTagPayloadDouble: new(_BACnetTagPayloadDouble)}
}

type _BACnetTagPayloadDoubleBuilder struct {
	*_BACnetTagPayloadDouble

	err *utils.MultiError
}

var _ (BACnetTagPayloadDoubleBuilder) = (*_BACnetTagPayloadDoubleBuilder)(nil)

func (b *_BACnetTagPayloadDoubleBuilder) WithMandatoryFields(value float64) BACnetTagPayloadDoubleBuilder {
	return b.WithValue(value)
}

func (b *_BACnetTagPayloadDoubleBuilder) WithValue(value float64) BACnetTagPayloadDoubleBuilder {
	b.Value = value
	return b
}

func (b *_BACnetTagPayloadDoubleBuilder) Build() (BACnetTagPayloadDouble, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTagPayloadDouble.deepCopy(), nil
}

func (b *_BACnetTagPayloadDoubleBuilder) MustBuild() BACnetTagPayloadDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetTagPayloadDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTagPayloadDoubleBuilder().(*_BACnetTagPayloadDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTagPayloadDoubleBuilder creates a BACnetTagPayloadDoubleBuilder
func (b *_BACnetTagPayloadDouble) CreateBACnetTagPayloadDoubleBuilder() BACnetTagPayloadDoubleBuilder {
	if b == nil {
		return NewBACnetTagPayloadDoubleBuilder()
	}
	return &_BACnetTagPayloadDoubleBuilder{_BACnetTagPayloadDouble: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadDouble) GetValue() float64 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadDouble(structType any) BACnetTagPayloadDouble {
	if casted, ok := structType.(BACnetTagPayloadDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadDouble) GetTypeName() string {
	return "BACnetTagPayloadDouble"
}

func (m *_BACnetTagPayloadDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 64

	return lengthInBits
}

func (m *_BACnetTagPayloadDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadDoubleParse(ctx context.Context, theBytes []byte) (BACnetTagPayloadDouble, error) {
	return BACnetTagPayloadDoubleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadDoubleParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDouble, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDouble, error) {
		return BACnetTagPayloadDoubleParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetTagPayloadDoubleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTagPayloadDouble, error) {
	v, err := (&_BACnetTagPayloadDouble{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetTagPayloadDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetTagPayloadDouble BACnetTagPayloadDouble, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField(ctx, "value", ReadDouble(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadDouble")
	}

	return m, nil
}

func (m *_BACnetTagPayloadDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadDouble"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadDouble")
	}

	if err := WriteSimpleField[float64](ctx, "value", m.GetValue(), WriteDouble(writeBuffer, 64)); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadDouble"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadDouble")
	}
	return nil
}

func (m *_BACnetTagPayloadDouble) IsBACnetTagPayloadDouble() {}

func (m *_BACnetTagPayloadDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTagPayloadDouble) deepCopy() *_BACnetTagPayloadDouble {
	if m == nil {
		return nil
	}
	_BACnetTagPayloadDoubleCopy := &_BACnetTagPayloadDouble{
		m.Value,
	}
	return _BACnetTagPayloadDoubleCopy
}

func (m *_BACnetTagPayloadDouble) String() string {
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