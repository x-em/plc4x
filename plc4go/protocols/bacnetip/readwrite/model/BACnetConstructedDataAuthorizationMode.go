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

// BACnetConstructedDataAuthorizationMode is the corresponding interface of BACnetConstructedDataAuthorizationMode
type BACnetConstructedDataAuthorizationMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAuthorizationMode returns AuthorizationMode (property field)
	GetAuthorizationMode() BACnetAuthorizationModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthorizationModeTagged
	// IsBACnetConstructedDataAuthorizationMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAuthorizationMode()
	// CreateBuilder creates a BACnetConstructedDataAuthorizationModeBuilder
	CreateBACnetConstructedDataAuthorizationModeBuilder() BACnetConstructedDataAuthorizationModeBuilder
}

// _BACnetConstructedDataAuthorizationMode is the data-structure of this message
type _BACnetConstructedDataAuthorizationMode struct {
	BACnetConstructedDataContract
	AuthorizationMode BACnetAuthorizationModeTagged
}

var _ BACnetConstructedDataAuthorizationMode = (*_BACnetConstructedDataAuthorizationMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAuthorizationMode)(nil)

// NewBACnetConstructedDataAuthorizationMode factory function for _BACnetConstructedDataAuthorizationMode
func NewBACnetConstructedDataAuthorizationMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, authorizationMode BACnetAuthorizationModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthorizationMode {
	if authorizationMode == nil {
		panic("authorizationMode of type BACnetAuthorizationModeTagged for BACnetConstructedDataAuthorizationMode must not be nil")
	}
	_result := &_BACnetConstructedDataAuthorizationMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AuthorizationMode:             authorizationMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAuthorizationModeBuilder is a builder for BACnetConstructedDataAuthorizationMode
type BACnetConstructedDataAuthorizationModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(authorizationMode BACnetAuthorizationModeTagged) BACnetConstructedDataAuthorizationModeBuilder
	// WithAuthorizationMode adds AuthorizationMode (property field)
	WithAuthorizationMode(BACnetAuthorizationModeTagged) BACnetConstructedDataAuthorizationModeBuilder
	// WithAuthorizationModeBuilder adds AuthorizationMode (property field) which is build by the builder
	WithAuthorizationModeBuilder(func(BACnetAuthorizationModeTaggedBuilder) BACnetAuthorizationModeTaggedBuilder) BACnetConstructedDataAuthorizationModeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAuthorizationMode or returns an error if something is wrong
	Build() (BACnetConstructedDataAuthorizationMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAuthorizationMode
}

// NewBACnetConstructedDataAuthorizationModeBuilder() creates a BACnetConstructedDataAuthorizationModeBuilder
func NewBACnetConstructedDataAuthorizationModeBuilder() BACnetConstructedDataAuthorizationModeBuilder {
	return &_BACnetConstructedDataAuthorizationModeBuilder{_BACnetConstructedDataAuthorizationMode: new(_BACnetConstructedDataAuthorizationMode)}
}

type _BACnetConstructedDataAuthorizationModeBuilder struct {
	*_BACnetConstructedDataAuthorizationMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAuthorizationModeBuilder) = (*_BACnetConstructedDataAuthorizationModeBuilder)(nil)

func (b *_BACnetConstructedDataAuthorizationModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAuthorizationMode
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) WithMandatoryFields(authorizationMode BACnetAuthorizationModeTagged) BACnetConstructedDataAuthorizationModeBuilder {
	return b.WithAuthorizationMode(authorizationMode)
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) WithAuthorizationMode(authorizationMode BACnetAuthorizationModeTagged) BACnetConstructedDataAuthorizationModeBuilder {
	b.AuthorizationMode = authorizationMode
	return b
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) WithAuthorizationModeBuilder(builderSupplier func(BACnetAuthorizationModeTaggedBuilder) BACnetAuthorizationModeTaggedBuilder) BACnetConstructedDataAuthorizationModeBuilder {
	builder := builderSupplier(b.AuthorizationMode.CreateBACnetAuthorizationModeTaggedBuilder())
	var err error
	b.AuthorizationMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetAuthorizationModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) Build() (BACnetConstructedDataAuthorizationMode, error) {
	if b.AuthorizationMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'authorizationMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAuthorizationMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) MustBuild() BACnetConstructedDataAuthorizationMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAuthorizationModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAuthorizationModeBuilder().(*_BACnetConstructedDataAuthorizationModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAuthorizationModeBuilder creates a BACnetConstructedDataAuthorizationModeBuilder
func (b *_BACnetConstructedDataAuthorizationMode) CreateBACnetConstructedDataAuthorizationModeBuilder() BACnetConstructedDataAuthorizationModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataAuthorizationModeBuilder()
	}
	return &_BACnetConstructedDataAuthorizationModeBuilder{_BACnetConstructedDataAuthorizationMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthorizationMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHORIZATION_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetAuthorizationMode() BACnetAuthorizationModeTagged {
	return m.AuthorizationMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthorizationMode) GetActualValue() BACnetAuthorizationModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAuthorizationModeTagged(m.GetAuthorizationMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthorizationMode(structType any) BACnetConstructedDataAuthorizationMode {
	if casted, ok := structType.(BACnetConstructedDataAuthorizationMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthorizationMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthorizationMode) GetTypeName() string {
	return "BACnetConstructedDataAuthorizationMode"
}

func (m *_BACnetConstructedDataAuthorizationMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (authorizationMode)
	lengthInBits += m.AuthorizationMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthorizationMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAuthorizationMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAuthorizationMode BACnetConstructedDataAuthorizationMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthorizationMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthorizationMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	authorizationMode, err := ReadSimpleField[BACnetAuthorizationModeTagged](ctx, "authorizationMode", ReadComplex[BACnetAuthorizationModeTagged](BACnetAuthorizationModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authorizationMode' field"))
	}
	m.AuthorizationMode = authorizationMode

	actualValue, err := ReadVirtualField[BACnetAuthorizationModeTagged](ctx, "actualValue", (*BACnetAuthorizationModeTagged)(nil), authorizationMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthorizationMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthorizationMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAuthorizationMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthorizationMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthorizationMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthorizationMode")
		}

		if err := WriteSimpleField[BACnetAuthorizationModeTagged](ctx, "authorizationMode", m.GetAuthorizationMode(), WriteComplex[BACnetAuthorizationModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authorizationMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthorizationMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthorizationMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthorizationMode) IsBACnetConstructedDataAuthorizationMode() {}

func (m *_BACnetConstructedDataAuthorizationMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAuthorizationMode) deepCopy() *_BACnetConstructedDataAuthorizationMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAuthorizationModeCopy := &_BACnetConstructedDataAuthorizationMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetAuthorizationModeTagged](m.AuthorizationMode),
	}
	_BACnetConstructedDataAuthorizationModeCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAuthorizationModeCopy
}

func (m *_BACnetConstructedDataAuthorizationMode) String() string {
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