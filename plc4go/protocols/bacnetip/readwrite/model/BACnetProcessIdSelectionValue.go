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

// BACnetProcessIdSelectionValue is the corresponding interface of BACnetProcessIdSelectionValue
type BACnetProcessIdSelectionValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetProcessIdSelection
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
	// IsBACnetProcessIdSelectionValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetProcessIdSelectionValue()
	// CreateBuilder creates a BACnetProcessIdSelectionValueBuilder
	CreateBACnetProcessIdSelectionValueBuilder() BACnetProcessIdSelectionValueBuilder
}

// _BACnetProcessIdSelectionValue is the data-structure of this message
type _BACnetProcessIdSelectionValue struct {
	BACnetProcessIdSelectionContract
	ProcessIdentifier BACnetApplicationTagUnsignedInteger
}

var _ BACnetProcessIdSelectionValue = (*_BACnetProcessIdSelectionValue)(nil)
var _ BACnetProcessIdSelectionRequirements = (*_BACnetProcessIdSelectionValue)(nil)

// NewBACnetProcessIdSelectionValue factory function for _BACnetProcessIdSelectionValue
func NewBACnetProcessIdSelectionValue(peekedTagHeader BACnetTagHeader, processIdentifier BACnetApplicationTagUnsignedInteger) *_BACnetProcessIdSelectionValue {
	if processIdentifier == nil {
		panic("processIdentifier of type BACnetApplicationTagUnsignedInteger for BACnetProcessIdSelectionValue must not be nil")
	}
	_result := &_BACnetProcessIdSelectionValue{
		BACnetProcessIdSelectionContract: NewBACnetProcessIdSelection(peekedTagHeader),
		ProcessIdentifier:                processIdentifier,
	}
	_result.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetProcessIdSelectionValueBuilder is a builder for BACnetProcessIdSelectionValue
type BACnetProcessIdSelectionValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(processIdentifier BACnetApplicationTagUnsignedInteger) BACnetProcessIdSelectionValueBuilder
	// WithProcessIdentifier adds ProcessIdentifier (property field)
	WithProcessIdentifier(BACnetApplicationTagUnsignedInteger) BACnetProcessIdSelectionValueBuilder
	// WithProcessIdentifierBuilder adds ProcessIdentifier (property field) which is build by the builder
	WithProcessIdentifierBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetProcessIdSelectionValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetProcessIdSelectionBuilder
	// Build builds the BACnetProcessIdSelectionValue or returns an error if something is wrong
	Build() (BACnetProcessIdSelectionValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetProcessIdSelectionValue
}

// NewBACnetProcessIdSelectionValueBuilder() creates a BACnetProcessIdSelectionValueBuilder
func NewBACnetProcessIdSelectionValueBuilder() BACnetProcessIdSelectionValueBuilder {
	return &_BACnetProcessIdSelectionValueBuilder{_BACnetProcessIdSelectionValue: new(_BACnetProcessIdSelectionValue)}
}

type _BACnetProcessIdSelectionValueBuilder struct {
	*_BACnetProcessIdSelectionValue

	parentBuilder *_BACnetProcessIdSelectionBuilder

	err *utils.MultiError
}

var _ (BACnetProcessIdSelectionValueBuilder) = (*_BACnetProcessIdSelectionValueBuilder)(nil)

func (b *_BACnetProcessIdSelectionValueBuilder) setParent(contract BACnetProcessIdSelectionContract) {
	b.BACnetProcessIdSelectionContract = contract
	contract.(*_BACnetProcessIdSelection)._SubType = b._BACnetProcessIdSelectionValue
}

func (b *_BACnetProcessIdSelectionValueBuilder) WithMandatoryFields(processIdentifier BACnetApplicationTagUnsignedInteger) BACnetProcessIdSelectionValueBuilder {
	return b.WithProcessIdentifier(processIdentifier)
}

func (b *_BACnetProcessIdSelectionValueBuilder) WithProcessIdentifier(processIdentifier BACnetApplicationTagUnsignedInteger) BACnetProcessIdSelectionValueBuilder {
	b.ProcessIdentifier = processIdentifier
	return b
}

func (b *_BACnetProcessIdSelectionValueBuilder) WithProcessIdentifierBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetProcessIdSelectionValueBuilder {
	builder := builderSupplier(b.ProcessIdentifier.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetProcessIdSelectionValueBuilder) Build() (BACnetProcessIdSelectionValue, error) {
	if b.ProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'processIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetProcessIdSelectionValue.deepCopy(), nil
}

func (b *_BACnetProcessIdSelectionValueBuilder) MustBuild() BACnetProcessIdSelectionValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetProcessIdSelectionValueBuilder) Done() BACnetProcessIdSelectionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetProcessIdSelectionBuilder().(*_BACnetProcessIdSelectionBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetProcessIdSelectionValueBuilder) buildForBACnetProcessIdSelection() (BACnetProcessIdSelection, error) {
	return b.Build()
}

func (b *_BACnetProcessIdSelectionValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetProcessIdSelectionValueBuilder().(*_BACnetProcessIdSelectionValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetProcessIdSelectionValueBuilder creates a BACnetProcessIdSelectionValueBuilder
func (b *_BACnetProcessIdSelectionValue) CreateBACnetProcessIdSelectionValueBuilder() BACnetProcessIdSelectionValueBuilder {
	if b == nil {
		return NewBACnetProcessIdSelectionValueBuilder()
	}
	return &_BACnetProcessIdSelectionValueBuilder{_BACnetProcessIdSelectionValue: b.deepCopy()}
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

func (m *_BACnetProcessIdSelectionValue) GetParent() BACnetProcessIdSelectionContract {
	return m.BACnetProcessIdSelectionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProcessIdSelectionValue) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetProcessIdSelectionValue(structType any) BACnetProcessIdSelectionValue {
	if casted, ok := structType.(BACnetProcessIdSelectionValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProcessIdSelectionValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProcessIdSelectionValue) GetTypeName() string {
	return "BACnetProcessIdSelectionValue"
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).getLengthInBits(ctx))

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetProcessIdSelectionValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetProcessIdSelectionValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetProcessIdSelection) (__bACnetProcessIdSelectionValue BACnetProcessIdSelectionValue, err error) {
	m.BACnetProcessIdSelectionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProcessIdSelectionValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProcessIdSelectionValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	processIdentifier, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'processIdentifier' field"))
	}
	m.ProcessIdentifier = processIdentifier

	if closeErr := readBuffer.CloseContext("BACnetProcessIdSelectionValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProcessIdSelectionValue")
	}

	return m, nil
}

func (m *_BACnetProcessIdSelectionValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProcessIdSelectionValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetProcessIdSelectionValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetProcessIdSelectionValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "processIdentifier", m.GetProcessIdentifier(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'processIdentifier' field")
		}

		if popErr := writeBuffer.PopContext("BACnetProcessIdSelectionValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetProcessIdSelectionValue")
		}
		return nil
	}
	return m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetProcessIdSelectionValue) IsBACnetProcessIdSelectionValue() {}

func (m *_BACnetProcessIdSelectionValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetProcessIdSelectionValue) deepCopy() *_BACnetProcessIdSelectionValue {
	if m == nil {
		return nil
	}
	_BACnetProcessIdSelectionValueCopy := &_BACnetProcessIdSelectionValue{
		m.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ProcessIdentifier),
	}
	_BACnetProcessIdSelectionValueCopy.BACnetProcessIdSelectionContract.(*_BACnetProcessIdSelection)._SubType = m
	return _BACnetProcessIdSelectionValueCopy
}

func (m *_BACnetProcessIdSelectionValue) String() string {
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