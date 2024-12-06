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

// BACnetFaultParameterFaultExtendedParametersEntryOctetString is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryOctetString
type BACnetFaultParameterFaultExtendedParametersEntryOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
	// IsBACnetFaultParameterFaultExtendedParametersEntryOctetString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryOctetString()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryOctetString is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryOctetString struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	OctetStringValue BACnetApplicationTagOctetString
}

var _ BACnetFaultParameterFaultExtendedParametersEntryOctetString = (*_BACnetFaultParameterFaultExtendedParametersEntryOctetString)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryOctetString)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryOctetString factory function for _BACnetFaultParameterFaultExtendedParametersEntryOctetString
func NewBACnetFaultParameterFaultExtendedParametersEntryOctetString(peekedTagHeader BACnetTagHeader, octetStringValue BACnetApplicationTagOctetString) *_BACnetFaultParameterFaultExtendedParametersEntryOctetString {
	if octetStringValue == nil {
		panic("octetStringValue of type BACnetApplicationTagOctetString for BACnetFaultParameterFaultExtendedParametersEntryOctetString must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryOctetString{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		OctetStringValue: octetStringValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryOctetString
type BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(octetStringValue BACnetApplicationTagOctetString) BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
	// WithOctetStringValue adds OctetStringValue (property field)
	WithOctetStringValue(BACnetApplicationTagOctetString) BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
	// WithOctetStringValueBuilder adds OctetStringValue (property field) which is build by the builder
	WithOctetStringValueBuilder(func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryOctetString or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryOctetString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryOctetString
}

// NewBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder{_BACnetFaultParameterFaultExtendedParametersEntryOctetString: new(_BACnetFaultParameterFaultExtendedParametersEntryOctetString)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryOctetString

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryOctetString
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) WithMandatoryFields(octetStringValue BACnetApplicationTagOctetString) BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder {
	return b.WithOctetStringValue(octetStringValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) WithOctetStringValue(octetStringValue BACnetApplicationTagOctetString) BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder {
	b.OctetStringValue = octetStringValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) WithOctetStringValueBuilder(builderSupplier func(BACnetApplicationTagOctetStringBuilder) BACnetApplicationTagOctetStringBuilder) BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder {
	builder := builderSupplier(b.OctetStringValue.CreateBACnetApplicationTagOctetStringBuilder())
	var err error
	b.OctetStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagOctetStringBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryOctetString, error) {
	if b.OctetStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'octetStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryOctetString.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryOctetString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) CreateBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryOctetStringBuilder{_BACnetFaultParameterFaultExtendedParametersEntryOctetString: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryOctetString(structType any) BACnetFaultParameterFaultExtendedParametersEntryOctetString {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryOctetString"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryOctetString BACnetFaultParameterFaultExtendedParametersEntryOctetString, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octetStringValue, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octetStringValue' field"))
	}
	m.OctetStringValue = octetStringValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", m.GetOctetStringValue(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryOctetString")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) IsBACnetFaultParameterFaultExtendedParametersEntryOctetString() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryOctetString {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryOctetStringCopy := &_BACnetFaultParameterFaultExtendedParametersEntryOctetString{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagOctetString](m.OctetStringValue),
	}
	_BACnetFaultParameterFaultExtendedParametersEntryOctetStringCopy.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryOctetStringCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryOctetString) String() string {
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