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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataExtLinkWrite is the corresponding interface of ApduDataExtLinkWrite
type ApduDataExtLinkWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduDataExt
	// IsApduDataExtLinkWrite is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduDataExtLinkWrite()
	// CreateBuilder creates a ApduDataExtLinkWriteBuilder
	CreateApduDataExtLinkWriteBuilder() ApduDataExtLinkWriteBuilder
}

// _ApduDataExtLinkWrite is the data-structure of this message
type _ApduDataExtLinkWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtLinkWrite = (*_ApduDataExtLinkWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtLinkWrite)(nil)

// NewApduDataExtLinkWrite factory function for _ApduDataExtLinkWrite
func NewApduDataExtLinkWrite(length uint8) *_ApduDataExtLinkWrite {
	_result := &_ApduDataExtLinkWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduDataExtLinkWriteBuilder is a builder for ApduDataExtLinkWrite
type ApduDataExtLinkWriteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduDataExtLinkWriteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduDataExtBuilder
	// Build builds the ApduDataExtLinkWrite or returns an error if something is wrong
	Build() (ApduDataExtLinkWrite, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduDataExtLinkWrite
}

// NewApduDataExtLinkWriteBuilder() creates a ApduDataExtLinkWriteBuilder
func NewApduDataExtLinkWriteBuilder() ApduDataExtLinkWriteBuilder {
	return &_ApduDataExtLinkWriteBuilder{_ApduDataExtLinkWrite: new(_ApduDataExtLinkWrite)}
}

type _ApduDataExtLinkWriteBuilder struct {
	*_ApduDataExtLinkWrite

	parentBuilder *_ApduDataExtBuilder

	err *utils.MultiError
}

var _ (ApduDataExtLinkWriteBuilder) = (*_ApduDataExtLinkWriteBuilder)(nil)

func (b *_ApduDataExtLinkWriteBuilder) setParent(contract ApduDataExtContract) {
	b.ApduDataExtContract = contract
	contract.(*_ApduDataExt)._SubType = b._ApduDataExtLinkWrite
}

func (b *_ApduDataExtLinkWriteBuilder) WithMandatoryFields() ApduDataExtLinkWriteBuilder {
	return b
}

func (b *_ApduDataExtLinkWriteBuilder) Build() (ApduDataExtLinkWrite, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduDataExtLinkWrite.deepCopy(), nil
}

func (b *_ApduDataExtLinkWriteBuilder) MustBuild() ApduDataExtLinkWrite {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduDataExtLinkWriteBuilder) Done() ApduDataExtBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduDataExtBuilder().(*_ApduDataExtBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduDataExtLinkWriteBuilder) buildForApduDataExt() (ApduDataExt, error) {
	return b.Build()
}

func (b *_ApduDataExtLinkWriteBuilder) DeepCopy() any {
	_copy := b.CreateApduDataExtLinkWriteBuilder().(*_ApduDataExtLinkWriteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduDataExtLinkWriteBuilder creates a ApduDataExtLinkWriteBuilder
func (b *_ApduDataExtLinkWrite) CreateApduDataExtLinkWriteBuilder() ApduDataExtLinkWriteBuilder {
	if b == nil {
		return NewApduDataExtLinkWriteBuilder()
	}
	return &_ApduDataExtLinkWriteBuilder{_ApduDataExtLinkWrite: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtLinkWrite) GetExtApciType() uint8 {
	return 0x27
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtLinkWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// Deprecated: use the interface for direct cast
func CastApduDataExtLinkWrite(structType any) ApduDataExtLinkWrite {
	if casted, ok := structType.(ApduDataExtLinkWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtLinkWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtLinkWrite) GetTypeName() string {
	return "ApduDataExtLinkWrite"
}

func (m *_ApduDataExtLinkWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtLinkWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtLinkWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtLinkWrite ApduDataExtLinkWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtLinkWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtLinkWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtLinkWrite")
	}

	return m, nil
}

func (m *_ApduDataExtLinkWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtLinkWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtLinkWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtLinkWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtLinkWrite) IsApduDataExtLinkWrite() {}

func (m *_ApduDataExtLinkWrite) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduDataExtLinkWrite) deepCopy() *_ApduDataExtLinkWrite {
	if m == nil {
		return nil
	}
	_ApduDataExtLinkWriteCopy := &_ApduDataExtLinkWrite{
		m.ApduDataExtContract.(*_ApduDataExt).deepCopy(),
	}
	_ApduDataExtLinkWriteCopy.ApduDataExtContract.(*_ApduDataExt)._SubType = m
	return _ApduDataExtLinkWriteCopy
}

func (m *_ApduDataExtLinkWrite) String() string {
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