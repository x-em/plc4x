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

// CBusCommandPointToMultiPoint is the corresponding interface of CBusCommandPointToMultiPoint
type CBusCommandPointToMultiPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CBusCommand
	// GetCommand returns Command (property field)
	GetCommand() CBusPointToMultiPointCommand
	// IsCBusCommandPointToMultiPoint is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCBusCommandPointToMultiPoint()
	// CreateBuilder creates a CBusCommandPointToMultiPointBuilder
	CreateCBusCommandPointToMultiPointBuilder() CBusCommandPointToMultiPointBuilder
}

// _CBusCommandPointToMultiPoint is the data-structure of this message
type _CBusCommandPointToMultiPoint struct {
	CBusCommandContract
	Command CBusPointToMultiPointCommand
}

var _ CBusCommandPointToMultiPoint = (*_CBusCommandPointToMultiPoint)(nil)
var _ CBusCommandRequirements = (*_CBusCommandPointToMultiPoint)(nil)

// NewCBusCommandPointToMultiPoint factory function for _CBusCommandPointToMultiPoint
func NewCBusCommandPointToMultiPoint(header CBusHeader, command CBusPointToMultiPointCommand, cBusOptions CBusOptions) *_CBusCommandPointToMultiPoint {
	if command == nil {
		panic("command of type CBusPointToMultiPointCommand for CBusCommandPointToMultiPoint must not be nil")
	}
	_result := &_CBusCommandPointToMultiPoint{
		CBusCommandContract: NewCBusCommand(header, cBusOptions),
		Command:             command,
	}
	_result.CBusCommandContract.(*_CBusCommand)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CBusCommandPointToMultiPointBuilder is a builder for CBusCommandPointToMultiPoint
type CBusCommandPointToMultiPointBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(command CBusPointToMultiPointCommand) CBusCommandPointToMultiPointBuilder
	// WithCommand adds Command (property field)
	WithCommand(CBusPointToMultiPointCommand) CBusCommandPointToMultiPointBuilder
	// WithCommandBuilder adds Command (property field) which is build by the builder
	WithCommandBuilder(func(CBusPointToMultiPointCommandBuilder) CBusPointToMultiPointCommandBuilder) CBusCommandPointToMultiPointBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CBusCommandBuilder
	// Build builds the CBusCommandPointToMultiPoint or returns an error if something is wrong
	Build() (CBusCommandPointToMultiPoint, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CBusCommandPointToMultiPoint
}

// NewCBusCommandPointToMultiPointBuilder() creates a CBusCommandPointToMultiPointBuilder
func NewCBusCommandPointToMultiPointBuilder() CBusCommandPointToMultiPointBuilder {
	return &_CBusCommandPointToMultiPointBuilder{_CBusCommandPointToMultiPoint: new(_CBusCommandPointToMultiPoint)}
}

type _CBusCommandPointToMultiPointBuilder struct {
	*_CBusCommandPointToMultiPoint

	parentBuilder *_CBusCommandBuilder

	err *utils.MultiError
}

var _ (CBusCommandPointToMultiPointBuilder) = (*_CBusCommandPointToMultiPointBuilder)(nil)

func (b *_CBusCommandPointToMultiPointBuilder) setParent(contract CBusCommandContract) {
	b.CBusCommandContract = contract
	contract.(*_CBusCommand)._SubType = b._CBusCommandPointToMultiPoint
}

func (b *_CBusCommandPointToMultiPointBuilder) WithMandatoryFields(command CBusPointToMultiPointCommand) CBusCommandPointToMultiPointBuilder {
	return b.WithCommand(command)
}

func (b *_CBusCommandPointToMultiPointBuilder) WithCommand(command CBusPointToMultiPointCommand) CBusCommandPointToMultiPointBuilder {
	b.Command = command
	return b
}

func (b *_CBusCommandPointToMultiPointBuilder) WithCommandBuilder(builderSupplier func(CBusPointToMultiPointCommandBuilder) CBusPointToMultiPointCommandBuilder) CBusCommandPointToMultiPointBuilder {
	builder := builderSupplier(b.Command.CreateCBusPointToMultiPointCommandBuilder())
	var err error
	b.Command, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "CBusPointToMultiPointCommandBuilder failed"))
	}
	return b
}

func (b *_CBusCommandPointToMultiPointBuilder) Build() (CBusCommandPointToMultiPoint, error) {
	if b.Command == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'command' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CBusCommandPointToMultiPoint.deepCopy(), nil
}

func (b *_CBusCommandPointToMultiPointBuilder) MustBuild() CBusCommandPointToMultiPoint {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CBusCommandPointToMultiPointBuilder) Done() CBusCommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCBusCommandBuilder().(*_CBusCommandBuilder)
	}
	return b.parentBuilder
}

func (b *_CBusCommandPointToMultiPointBuilder) buildForCBusCommand() (CBusCommand, error) {
	return b.Build()
}

func (b *_CBusCommandPointToMultiPointBuilder) DeepCopy() any {
	_copy := b.CreateCBusCommandPointToMultiPointBuilder().(*_CBusCommandPointToMultiPointBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCBusCommandPointToMultiPointBuilder creates a CBusCommandPointToMultiPointBuilder
func (b *_CBusCommandPointToMultiPoint) CreateCBusCommandPointToMultiPointBuilder() CBusCommandPointToMultiPointBuilder {
	if b == nil {
		return NewCBusCommandPointToMultiPointBuilder()
	}
	return &_CBusCommandPointToMultiPointBuilder{_CBusCommandPointToMultiPoint: b.deepCopy()}
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

func (m *_CBusCommandPointToMultiPoint) GetParent() CBusCommandContract {
	return m.CBusCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandPointToMultiPoint) GetCommand() CBusPointToMultiPointCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCBusCommandPointToMultiPoint(structType any) CBusCommandPointToMultiPoint {
	if casted, ok := structType.(CBusCommandPointToMultiPoint); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandPointToMultiPoint); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandPointToMultiPoint) GetTypeName() string {
	return "CBusCommandPointToMultiPoint"
}

func (m *_CBusCommandPointToMultiPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CBusCommandContract.(*_CBusCommand).getLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusCommandPointToMultiPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CBusCommandPointToMultiPoint) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CBusCommand, cBusOptions CBusOptions) (__cBusCommandPointToMultiPoint CBusCommandPointToMultiPoint, err error) {
	m.CBusCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommandPointToMultiPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandPointToMultiPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	command, err := ReadSimpleField[CBusPointToMultiPointCommand](ctx, "command", ReadComplex[CBusPointToMultiPointCommand](CBusPointToMultiPointCommandParseWithBufferProducer[CBusPointToMultiPointCommand]((CBusOptions)(cBusOptions)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	if closeErr := readBuffer.CloseContext("CBusCommandPointToMultiPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandPointToMultiPoint")
	}

	return m, nil
}

func (m *_CBusCommandPointToMultiPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusCommandPointToMultiPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToMultiPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandPointToMultiPoint")
		}

		if err := WriteSimpleField[CBusPointToMultiPointCommand](ctx, "command", m.GetCommand(), WriteComplex[CBusPointToMultiPointCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToMultiPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandPointToMultiPoint")
		}
		return nil
	}
	return m.CBusCommandContract.(*_CBusCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusCommandPointToMultiPoint) IsCBusCommandPointToMultiPoint() {}

func (m *_CBusCommandPointToMultiPoint) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CBusCommandPointToMultiPoint) deepCopy() *_CBusCommandPointToMultiPoint {
	if m == nil {
		return nil
	}
	_CBusCommandPointToMultiPointCopy := &_CBusCommandPointToMultiPoint{
		m.CBusCommandContract.(*_CBusCommand).deepCopy(),
		utils.DeepCopy[CBusPointToMultiPointCommand](m.Command),
	}
	_CBusCommandPointToMultiPointCopy.CBusCommandContract.(*_CBusCommand)._SubType = m
	return _CBusCommandPointToMultiPointCopy
}

func (m *_CBusCommandPointToMultiPoint) String() string {
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