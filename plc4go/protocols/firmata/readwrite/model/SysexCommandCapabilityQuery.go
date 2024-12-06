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

// SysexCommandCapabilityQuery is the corresponding interface of SysexCommandCapabilityQuery
type SysexCommandCapabilityQuery interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SysexCommand
	// IsSysexCommandCapabilityQuery is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSysexCommandCapabilityQuery()
	// CreateBuilder creates a SysexCommandCapabilityQueryBuilder
	CreateSysexCommandCapabilityQueryBuilder() SysexCommandCapabilityQueryBuilder
}

// _SysexCommandCapabilityQuery is the data-structure of this message
type _SysexCommandCapabilityQuery struct {
	SysexCommandContract
}

var _ SysexCommandCapabilityQuery = (*_SysexCommandCapabilityQuery)(nil)
var _ SysexCommandRequirements = (*_SysexCommandCapabilityQuery)(nil)

// NewSysexCommandCapabilityQuery factory function for _SysexCommandCapabilityQuery
func NewSysexCommandCapabilityQuery() *_SysexCommandCapabilityQuery {
	_result := &_SysexCommandCapabilityQuery{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SysexCommandCapabilityQueryBuilder is a builder for SysexCommandCapabilityQuery
type SysexCommandCapabilityQueryBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SysexCommandCapabilityQueryBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SysexCommandBuilder
	// Build builds the SysexCommandCapabilityQuery or returns an error if something is wrong
	Build() (SysexCommandCapabilityQuery, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SysexCommandCapabilityQuery
}

// NewSysexCommandCapabilityQueryBuilder() creates a SysexCommandCapabilityQueryBuilder
func NewSysexCommandCapabilityQueryBuilder() SysexCommandCapabilityQueryBuilder {
	return &_SysexCommandCapabilityQueryBuilder{_SysexCommandCapabilityQuery: new(_SysexCommandCapabilityQuery)}
}

type _SysexCommandCapabilityQueryBuilder struct {
	*_SysexCommandCapabilityQuery

	parentBuilder *_SysexCommandBuilder

	err *utils.MultiError
}

var _ (SysexCommandCapabilityQueryBuilder) = (*_SysexCommandCapabilityQueryBuilder)(nil)

func (b *_SysexCommandCapabilityQueryBuilder) setParent(contract SysexCommandContract) {
	b.SysexCommandContract = contract
	contract.(*_SysexCommand)._SubType = b._SysexCommandCapabilityQuery
}

func (b *_SysexCommandCapabilityQueryBuilder) WithMandatoryFields() SysexCommandCapabilityQueryBuilder {
	return b
}

func (b *_SysexCommandCapabilityQueryBuilder) Build() (SysexCommandCapabilityQuery, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SysexCommandCapabilityQuery.deepCopy(), nil
}

func (b *_SysexCommandCapabilityQueryBuilder) MustBuild() SysexCommandCapabilityQuery {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SysexCommandCapabilityQueryBuilder) Done() SysexCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSysexCommandBuilder().(*_SysexCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_SysexCommandCapabilityQueryBuilder) buildForSysexCommand() (SysexCommand, error) {
	return b.Build()
}

func (b *_SysexCommandCapabilityQueryBuilder) DeepCopy() any {
	_copy := b.CreateSysexCommandCapabilityQueryBuilder().(*_SysexCommandCapabilityQueryBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSysexCommandCapabilityQueryBuilder creates a SysexCommandCapabilityQueryBuilder
func (b *_SysexCommandCapabilityQuery) CreateSysexCommandCapabilityQueryBuilder() SysexCommandCapabilityQueryBuilder {
	if b == nil {
		return NewSysexCommandCapabilityQueryBuilder()
	}
	return &_SysexCommandCapabilityQueryBuilder{_SysexCommandCapabilityQuery: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandCapabilityQuery) GetCommandType() uint8 {
	return 0x6B
}

func (m *_SysexCommandCapabilityQuery) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandCapabilityQuery) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// Deprecated: use the interface for direct cast
func CastSysexCommandCapabilityQuery(structType any) SysexCommandCapabilityQuery {
	if casted, ok := structType.(SysexCommandCapabilityQuery); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandCapabilityQuery); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandCapabilityQuery) GetTypeName() string {
	return "SysexCommandCapabilityQuery"
}

func (m *_SysexCommandCapabilityQuery) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandCapabilityQuery) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandCapabilityQuery) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandCapabilityQuery SysexCommandCapabilityQuery, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandCapabilityQuery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandCapabilityQuery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandCapabilityQuery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandCapabilityQuery")
	}

	return m, nil
}

func (m *_SysexCommandCapabilityQuery) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandCapabilityQuery) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandCapabilityQuery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandCapabilityQuery")
		}

		if popErr := writeBuffer.PopContext("SysexCommandCapabilityQuery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandCapabilityQuery")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandCapabilityQuery) IsSysexCommandCapabilityQuery() {}

func (m *_SysexCommandCapabilityQuery) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SysexCommandCapabilityQuery) deepCopy() *_SysexCommandCapabilityQuery {
	if m == nil {
		return nil
	}
	_SysexCommandCapabilityQueryCopy := &_SysexCommandCapabilityQuery{
		m.SysexCommandContract.(*_SysexCommand).deepCopy(),
	}
	_SysexCommandCapabilityQueryCopy.SysexCommandContract.(*_SysexCommand)._SubType = m
	return _SysexCommandCapabilityQueryCopy
}

func (m *_SysexCommandCapabilityQuery) String() string {
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