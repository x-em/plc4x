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

// BACnetPropertyStatesNetworkNumberQuality is the corresponding interface of BACnetPropertyStatesNetworkNumberQuality
type BACnetPropertyStatesNetworkNumberQuality interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetNetworkNumberQuality returns NetworkNumberQuality (property field)
	GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged
	// IsBACnetPropertyStatesNetworkNumberQuality is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesNetworkNumberQuality()
	// CreateBuilder creates a BACnetPropertyStatesNetworkNumberQualityBuilder
	CreateBACnetPropertyStatesNetworkNumberQualityBuilder() BACnetPropertyStatesNetworkNumberQualityBuilder
}

// _BACnetPropertyStatesNetworkNumberQuality is the data-structure of this message
type _BACnetPropertyStatesNetworkNumberQuality struct {
	BACnetPropertyStatesContract
	NetworkNumberQuality BACnetNetworkNumberQualityTagged
}

var _ BACnetPropertyStatesNetworkNumberQuality = (*_BACnetPropertyStatesNetworkNumberQuality)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesNetworkNumberQuality)(nil)

// NewBACnetPropertyStatesNetworkNumberQuality factory function for _BACnetPropertyStatesNetworkNumberQuality
func NewBACnetPropertyStatesNetworkNumberQuality(peekedTagHeader BACnetTagHeader, networkNumberQuality BACnetNetworkNumberQualityTagged) *_BACnetPropertyStatesNetworkNumberQuality {
	if networkNumberQuality == nil {
		panic("networkNumberQuality of type BACnetNetworkNumberQualityTagged for BACnetPropertyStatesNetworkNumberQuality must not be nil")
	}
	_result := &_BACnetPropertyStatesNetworkNumberQuality{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		NetworkNumberQuality:         networkNumberQuality,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesNetworkNumberQualityBuilder is a builder for BACnetPropertyStatesNetworkNumberQuality
type BACnetPropertyStatesNetworkNumberQualityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkNumberQuality BACnetNetworkNumberQualityTagged) BACnetPropertyStatesNetworkNumberQualityBuilder
	// WithNetworkNumberQuality adds NetworkNumberQuality (property field)
	WithNetworkNumberQuality(BACnetNetworkNumberQualityTagged) BACnetPropertyStatesNetworkNumberQualityBuilder
	// WithNetworkNumberQualityBuilder adds NetworkNumberQuality (property field) which is build by the builder
	WithNetworkNumberQualityBuilder(func(BACnetNetworkNumberQualityTaggedBuilder) BACnetNetworkNumberQualityTaggedBuilder) BACnetPropertyStatesNetworkNumberQualityBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesNetworkNumberQuality or returns an error if something is wrong
	Build() (BACnetPropertyStatesNetworkNumberQuality, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesNetworkNumberQuality
}

// NewBACnetPropertyStatesNetworkNumberQualityBuilder() creates a BACnetPropertyStatesNetworkNumberQualityBuilder
func NewBACnetPropertyStatesNetworkNumberQualityBuilder() BACnetPropertyStatesNetworkNumberQualityBuilder {
	return &_BACnetPropertyStatesNetworkNumberQualityBuilder{_BACnetPropertyStatesNetworkNumberQuality: new(_BACnetPropertyStatesNetworkNumberQuality)}
}

type _BACnetPropertyStatesNetworkNumberQualityBuilder struct {
	*_BACnetPropertyStatesNetworkNumberQuality

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesNetworkNumberQualityBuilder) = (*_BACnetPropertyStatesNetworkNumberQualityBuilder)(nil)

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesNetworkNumberQuality
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) WithMandatoryFields(networkNumberQuality BACnetNetworkNumberQualityTagged) BACnetPropertyStatesNetworkNumberQualityBuilder {
	return b.WithNetworkNumberQuality(networkNumberQuality)
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) WithNetworkNumberQuality(networkNumberQuality BACnetNetworkNumberQualityTagged) BACnetPropertyStatesNetworkNumberQualityBuilder {
	b.NetworkNumberQuality = networkNumberQuality
	return b
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) WithNetworkNumberQualityBuilder(builderSupplier func(BACnetNetworkNumberQualityTaggedBuilder) BACnetNetworkNumberQualityTaggedBuilder) BACnetPropertyStatesNetworkNumberQualityBuilder {
	builder := builderSupplier(b.NetworkNumberQuality.CreateBACnetNetworkNumberQualityTaggedBuilder())
	var err error
	b.NetworkNumberQuality, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetNetworkNumberQualityTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) Build() (BACnetPropertyStatesNetworkNumberQuality, error) {
	if b.NetworkNumberQuality == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'networkNumberQuality' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesNetworkNumberQuality.deepCopy(), nil
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) MustBuild() BACnetPropertyStatesNetworkNumberQuality {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesNetworkNumberQualityBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesNetworkNumberQualityBuilder().(*_BACnetPropertyStatesNetworkNumberQualityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesNetworkNumberQualityBuilder creates a BACnetPropertyStatesNetworkNumberQualityBuilder
func (b *_BACnetPropertyStatesNetworkNumberQuality) CreateBACnetPropertyStatesNetworkNumberQualityBuilder() BACnetPropertyStatesNetworkNumberQualityBuilder {
	if b == nil {
		return NewBACnetPropertyStatesNetworkNumberQualityBuilder()
	}
	return &_BACnetPropertyStatesNetworkNumberQualityBuilder{_BACnetPropertyStatesNetworkNumberQuality: b.deepCopy()}
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

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged {
	return m.NetworkNumberQuality
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNetworkNumberQuality(structType any) BACnetPropertyStatesNetworkNumberQuality {
	if casted, ok := structType.(BACnetPropertyStatesNetworkNumberQuality); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkNumberQuality); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetTypeName() string {
	return "BACnetPropertyStatesNetworkNumberQuality"
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (networkNumberQuality)
	lengthInBits += m.NetworkNumberQuality.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesNetworkNumberQuality BACnetPropertyStatesNetworkNumberQuality, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkNumberQuality")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkNumberQuality, err := ReadSimpleField[BACnetNetworkNumberQualityTagged](ctx, "networkNumberQuality", ReadComplex[BACnetNetworkNumberQualityTagged](BACnetNetworkNumberQualityTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkNumberQuality' field"))
	}
	m.NetworkNumberQuality = networkNumberQuality

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkNumberQuality")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkNumberQuality")
		}

		if err := WriteSimpleField[BACnetNetworkNumberQualityTagged](ctx, "networkNumberQuality", m.GetNetworkNumberQuality(), WriteComplex[BACnetNetworkNumberQualityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkNumberQuality' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkNumberQuality")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) IsBACnetPropertyStatesNetworkNumberQuality() {}

func (m *_BACnetPropertyStatesNetworkNumberQuality) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) deepCopy() *_BACnetPropertyStatesNetworkNumberQuality {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesNetworkNumberQualityCopy := &_BACnetPropertyStatesNetworkNumberQuality{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetNetworkNumberQualityTagged](m.NetworkNumberQuality),
	}
	_BACnetPropertyStatesNetworkNumberQualityCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesNetworkNumberQualityCopy
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) String() string {
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