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

// BACnetPropertyStatesNetworkType is the corresponding interface of BACnetPropertyStatesNetworkType
type BACnetPropertyStatesNetworkType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetNetworkType returns NetworkType (property field)
	GetNetworkType() BACnetNetworkTypeTagged
	// IsBACnetPropertyStatesNetworkType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesNetworkType()
	// CreateBuilder creates a BACnetPropertyStatesNetworkTypeBuilder
	CreateBACnetPropertyStatesNetworkTypeBuilder() BACnetPropertyStatesNetworkTypeBuilder
}

// _BACnetPropertyStatesNetworkType is the data-structure of this message
type _BACnetPropertyStatesNetworkType struct {
	BACnetPropertyStatesContract
	NetworkType BACnetNetworkTypeTagged
}

var _ BACnetPropertyStatesNetworkType = (*_BACnetPropertyStatesNetworkType)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesNetworkType)(nil)

// NewBACnetPropertyStatesNetworkType factory function for _BACnetPropertyStatesNetworkType
func NewBACnetPropertyStatesNetworkType(peekedTagHeader BACnetTagHeader, networkType BACnetNetworkTypeTagged) *_BACnetPropertyStatesNetworkType {
	if networkType == nil {
		panic("networkType of type BACnetNetworkTypeTagged for BACnetPropertyStatesNetworkType must not be nil")
	}
	_result := &_BACnetPropertyStatesNetworkType{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		NetworkType:                  networkType,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesNetworkTypeBuilder is a builder for BACnetPropertyStatesNetworkType
type BACnetPropertyStatesNetworkTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkType BACnetNetworkTypeTagged) BACnetPropertyStatesNetworkTypeBuilder
	// WithNetworkType adds NetworkType (property field)
	WithNetworkType(BACnetNetworkTypeTagged) BACnetPropertyStatesNetworkTypeBuilder
	// WithNetworkTypeBuilder adds NetworkType (property field) which is build by the builder
	WithNetworkTypeBuilder(func(BACnetNetworkTypeTaggedBuilder) BACnetNetworkTypeTaggedBuilder) BACnetPropertyStatesNetworkTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesNetworkType or returns an error if something is wrong
	Build() (BACnetPropertyStatesNetworkType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesNetworkType
}

// NewBACnetPropertyStatesNetworkTypeBuilder() creates a BACnetPropertyStatesNetworkTypeBuilder
func NewBACnetPropertyStatesNetworkTypeBuilder() BACnetPropertyStatesNetworkTypeBuilder {
	return &_BACnetPropertyStatesNetworkTypeBuilder{_BACnetPropertyStatesNetworkType: new(_BACnetPropertyStatesNetworkType)}
}

type _BACnetPropertyStatesNetworkTypeBuilder struct {
	*_BACnetPropertyStatesNetworkType

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesNetworkTypeBuilder) = (*_BACnetPropertyStatesNetworkTypeBuilder)(nil)

func (b *_BACnetPropertyStatesNetworkTypeBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesNetworkType
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) WithMandatoryFields(networkType BACnetNetworkTypeTagged) BACnetPropertyStatesNetworkTypeBuilder {
	return b.WithNetworkType(networkType)
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) WithNetworkType(networkType BACnetNetworkTypeTagged) BACnetPropertyStatesNetworkTypeBuilder {
	b.NetworkType = networkType
	return b
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) WithNetworkTypeBuilder(builderSupplier func(BACnetNetworkTypeTaggedBuilder) BACnetNetworkTypeTaggedBuilder) BACnetPropertyStatesNetworkTypeBuilder {
	builder := builderSupplier(b.NetworkType.CreateBACnetNetworkTypeTaggedBuilder())
	var err error
	b.NetworkType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNetworkTypeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) Build() (BACnetPropertyStatesNetworkType, error) {
	if b.NetworkType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'networkType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesNetworkType.deepCopy(), nil
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) MustBuild() BACnetPropertyStatesNetworkType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesNetworkTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesNetworkTypeBuilder().(*_BACnetPropertyStatesNetworkTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesNetworkTypeBuilder creates a BACnetPropertyStatesNetworkTypeBuilder
func (b *_BACnetPropertyStatesNetworkType) CreateBACnetPropertyStatesNetworkTypeBuilder() BACnetPropertyStatesNetworkTypeBuilder {
	if b == nil {
		return NewBACnetPropertyStatesNetworkTypeBuilder()
	}
	return &_BACnetPropertyStatesNetworkTypeBuilder{_BACnetPropertyStatesNetworkType: b.deepCopy()}
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

func (m *_BACnetPropertyStatesNetworkType) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNetworkType) GetNetworkType() BACnetNetworkTypeTagged {
	return m.NetworkType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNetworkType(structType any) BACnetPropertyStatesNetworkType {
	if casted, ok := structType.(BACnetPropertyStatesNetworkType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNetworkType) GetTypeName() string {
	return "BACnetPropertyStatesNetworkType"
}

func (m *_BACnetPropertyStatesNetworkType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (networkType)
	lengthInBits += m.NetworkType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesNetworkType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesNetworkType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesNetworkType BACnetPropertyStatesNetworkType, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkType, err := ReadSimpleField[BACnetNetworkTypeTagged](ctx, "networkType", ReadComplex[BACnetNetworkTypeTagged](BACnetNetworkTypeTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkType' field"))
	}
	m.NetworkType = networkType

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkType")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesNetworkType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNetworkType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkType")
		}

		if err := WriteSimpleField[BACnetNetworkTypeTagged](ctx, "networkType", m.GetNetworkType(), WriteComplex[BACnetNetworkTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkType")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNetworkType) IsBACnetPropertyStatesNetworkType() {}

func (m *_BACnetPropertyStatesNetworkType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesNetworkType) deepCopy() *_BACnetPropertyStatesNetworkType {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesNetworkTypeCopy := &_BACnetPropertyStatesNetworkType{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetNetworkTypeTagged](m.NetworkType),
	}
	_BACnetPropertyStatesNetworkTypeCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesNetworkTypeCopy
}

func (m *_BACnetPropertyStatesNetworkType) String() string {
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