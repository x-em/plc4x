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

// BACnetConstructedDataDeployedProfileLocation is the corresponding interface of BACnetConstructedDataDeployedProfileLocation
type BACnetConstructedDataDeployedProfileLocation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDeployedProfileLocation returns DeployedProfileLocation (property field)
	GetDeployedProfileLocation() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataDeployedProfileLocation is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataDeployedProfileLocation()
	// CreateBuilder creates a BACnetConstructedDataDeployedProfileLocationBuilder
	CreateBACnetConstructedDataDeployedProfileLocationBuilder() BACnetConstructedDataDeployedProfileLocationBuilder
}

// _BACnetConstructedDataDeployedProfileLocation is the data-structure of this message
type _BACnetConstructedDataDeployedProfileLocation struct {
	BACnetConstructedDataContract
	DeployedProfileLocation BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataDeployedProfileLocation = (*_BACnetConstructedDataDeployedProfileLocation)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDeployedProfileLocation)(nil)

// NewBACnetConstructedDataDeployedProfileLocation factory function for _BACnetConstructedDataDeployedProfileLocation
func NewBACnetConstructedDataDeployedProfileLocation(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, deployedProfileLocation BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDeployedProfileLocation {
	if deployedProfileLocation == nil {
		panic("deployedProfileLocation of type BACnetApplicationTagCharacterString for BACnetConstructedDataDeployedProfileLocation must not be nil")
	}
	_result := &_BACnetConstructedDataDeployedProfileLocation{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DeployedProfileLocation:       deployedProfileLocation,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataDeployedProfileLocationBuilder is a builder for BACnetConstructedDataDeployedProfileLocation
type BACnetConstructedDataDeployedProfileLocationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deployedProfileLocation BACnetApplicationTagCharacterString) BACnetConstructedDataDeployedProfileLocationBuilder
	// WithDeployedProfileLocation adds DeployedProfileLocation (property field)
	WithDeployedProfileLocation(BACnetApplicationTagCharacterString) BACnetConstructedDataDeployedProfileLocationBuilder
	// WithDeployedProfileLocationBuilder adds DeployedProfileLocation (property field) which is build by the builder
	WithDeployedProfileLocationBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataDeployedProfileLocationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataDeployedProfileLocation or returns an error if something is wrong
	Build() (BACnetConstructedDataDeployedProfileLocation, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataDeployedProfileLocation
}

// NewBACnetConstructedDataDeployedProfileLocationBuilder() creates a BACnetConstructedDataDeployedProfileLocationBuilder
func NewBACnetConstructedDataDeployedProfileLocationBuilder() BACnetConstructedDataDeployedProfileLocationBuilder {
	return &_BACnetConstructedDataDeployedProfileLocationBuilder{_BACnetConstructedDataDeployedProfileLocation: new(_BACnetConstructedDataDeployedProfileLocation)}
}

type _BACnetConstructedDataDeployedProfileLocationBuilder struct {
	*_BACnetConstructedDataDeployedProfileLocation

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataDeployedProfileLocationBuilder) = (*_BACnetConstructedDataDeployedProfileLocationBuilder)(nil)

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataDeployedProfileLocation
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) WithMandatoryFields(deployedProfileLocation BACnetApplicationTagCharacterString) BACnetConstructedDataDeployedProfileLocationBuilder {
	return b.WithDeployedProfileLocation(deployedProfileLocation)
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) WithDeployedProfileLocation(deployedProfileLocation BACnetApplicationTagCharacterString) BACnetConstructedDataDeployedProfileLocationBuilder {
	b.DeployedProfileLocation = deployedProfileLocation
	return b
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) WithDeployedProfileLocationBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataDeployedProfileLocationBuilder {
	builder := builderSupplier(b.DeployedProfileLocation.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.DeployedProfileLocation, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) Build() (BACnetConstructedDataDeployedProfileLocation, error) {
	if b.DeployedProfileLocation == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deployedProfileLocation' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataDeployedProfileLocation.deepCopy(), nil
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) MustBuild() BACnetConstructedDataDeployedProfileLocation {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataDeployedProfileLocationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataDeployedProfileLocationBuilder().(*_BACnetConstructedDataDeployedProfileLocationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataDeployedProfileLocationBuilder creates a BACnetConstructedDataDeployedProfileLocationBuilder
func (b *_BACnetConstructedDataDeployedProfileLocation) CreateBACnetConstructedDataDeployedProfileLocationBuilder() BACnetConstructedDataDeployedProfileLocationBuilder {
	if b == nil {
		return NewBACnetConstructedDataDeployedProfileLocationBuilder()
	}
	return &_BACnetConstructedDataDeployedProfileLocationBuilder{_BACnetConstructedDataDeployedProfileLocation: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEPLOYED_PROFILE_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetDeployedProfileLocation() BACnetApplicationTagCharacterString {
	return m.DeployedProfileLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDeployedProfileLocation) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDeployedProfileLocation())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeployedProfileLocation(structType any) BACnetConstructedDataDeployedProfileLocation {
	if casted, ok := structType.(BACnetConstructedDataDeployedProfileLocation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeployedProfileLocation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetTypeName() string {
	return "BACnetConstructedDataDeployedProfileLocation"
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (deployedProfileLocation)
	lengthInBits += m.DeployedProfileLocation.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDeployedProfileLocation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDeployedProfileLocation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDeployedProfileLocation BACnetConstructedDataDeployedProfileLocation, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeployedProfileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeployedProfileLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deployedProfileLocation, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "deployedProfileLocation", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deployedProfileLocation' field"))
	}
	m.DeployedProfileLocation = deployedProfileLocation

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), deployedProfileLocation)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeployedProfileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeployedProfileLocation")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDeployedProfileLocation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDeployedProfileLocation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeployedProfileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeployedProfileLocation")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "deployedProfileLocation", m.GetDeployedProfileLocation(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deployedProfileLocation' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeployedProfileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeployedProfileLocation")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDeployedProfileLocation) IsBACnetConstructedDataDeployedProfileLocation() {
}

func (m *_BACnetConstructedDataDeployedProfileLocation) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataDeployedProfileLocation) deepCopy() *_BACnetConstructedDataDeployedProfileLocation {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataDeployedProfileLocationCopy := &_BACnetConstructedDataDeployedProfileLocation{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.DeployedProfileLocation),
	}
	_BACnetConstructedDataDeployedProfileLocationCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataDeployedProfileLocationCopy
}

func (m *_BACnetConstructedDataDeployedProfileLocation) String() string {
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