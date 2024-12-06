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

// BrokerDataSetReaderTransportDataType is the corresponding interface of BrokerDataSetReaderTransportDataType
type BrokerDataSetReaderTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetQueueName returns QueueName (property field)
	GetQueueName() PascalString
	// GetResourceUri returns ResourceUri (property field)
	GetResourceUri() PascalString
	// GetAuthenticationProfileUri returns AuthenticationProfileUri (property field)
	GetAuthenticationProfileUri() PascalString
	// GetRequestedDeliveryGuarantee returns RequestedDeliveryGuarantee (property field)
	GetRequestedDeliveryGuarantee() BrokerTransportQualityOfService
	// GetMetaDataQueueName returns MetaDataQueueName (property field)
	GetMetaDataQueueName() PascalString
	// IsBrokerDataSetReaderTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBrokerDataSetReaderTransportDataType()
	// CreateBuilder creates a BrokerDataSetReaderTransportDataTypeBuilder
	CreateBrokerDataSetReaderTransportDataTypeBuilder() BrokerDataSetReaderTransportDataTypeBuilder
}

// _BrokerDataSetReaderTransportDataType is the data-structure of this message
type _BrokerDataSetReaderTransportDataType struct {
	ExtensionObjectDefinitionContract
	QueueName                  PascalString
	ResourceUri                PascalString
	AuthenticationProfileUri   PascalString
	RequestedDeliveryGuarantee BrokerTransportQualityOfService
	MetaDataQueueName          PascalString
}

var _ BrokerDataSetReaderTransportDataType = (*_BrokerDataSetReaderTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrokerDataSetReaderTransportDataType)(nil)

// NewBrokerDataSetReaderTransportDataType factory function for _BrokerDataSetReaderTransportDataType
func NewBrokerDataSetReaderTransportDataType(queueName PascalString, resourceUri PascalString, authenticationProfileUri PascalString, requestedDeliveryGuarantee BrokerTransportQualityOfService, metaDataQueueName PascalString) *_BrokerDataSetReaderTransportDataType {
	if queueName == nil {
		panic("queueName of type PascalString for BrokerDataSetReaderTransportDataType must not be nil")
	}
	if resourceUri == nil {
		panic("resourceUri of type PascalString for BrokerDataSetReaderTransportDataType must not be nil")
	}
	if authenticationProfileUri == nil {
		panic("authenticationProfileUri of type PascalString for BrokerDataSetReaderTransportDataType must not be nil")
	}
	if metaDataQueueName == nil {
		panic("metaDataQueueName of type PascalString for BrokerDataSetReaderTransportDataType must not be nil")
	}
	_result := &_BrokerDataSetReaderTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		QueueName:                         queueName,
		ResourceUri:                       resourceUri,
		AuthenticationProfileUri:          authenticationProfileUri,
		RequestedDeliveryGuarantee:        requestedDeliveryGuarantee,
		MetaDataQueueName:                 metaDataQueueName,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BrokerDataSetReaderTransportDataTypeBuilder is a builder for BrokerDataSetReaderTransportDataType
type BrokerDataSetReaderTransportDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(queueName PascalString, resourceUri PascalString, authenticationProfileUri PascalString, requestedDeliveryGuarantee BrokerTransportQualityOfService, metaDataQueueName PascalString) BrokerDataSetReaderTransportDataTypeBuilder
	// WithQueueName adds QueueName (property field)
	WithQueueName(PascalString) BrokerDataSetReaderTransportDataTypeBuilder
	// WithQueueNameBuilder adds QueueName (property field) which is build by the builder
	WithQueueNameBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder
	// WithResourceUri adds ResourceUri (property field)
	WithResourceUri(PascalString) BrokerDataSetReaderTransportDataTypeBuilder
	// WithResourceUriBuilder adds ResourceUri (property field) which is build by the builder
	WithResourceUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder
	// WithAuthenticationProfileUri adds AuthenticationProfileUri (property field)
	WithAuthenticationProfileUri(PascalString) BrokerDataSetReaderTransportDataTypeBuilder
	// WithAuthenticationProfileUriBuilder adds AuthenticationProfileUri (property field) which is build by the builder
	WithAuthenticationProfileUriBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder
	// WithRequestedDeliveryGuarantee adds RequestedDeliveryGuarantee (property field)
	WithRequestedDeliveryGuarantee(BrokerTransportQualityOfService) BrokerDataSetReaderTransportDataTypeBuilder
	// WithMetaDataQueueName adds MetaDataQueueName (property field)
	WithMetaDataQueueName(PascalString) BrokerDataSetReaderTransportDataTypeBuilder
	// WithMetaDataQueueNameBuilder adds MetaDataQueueName (property field) which is build by the builder
	WithMetaDataQueueNameBuilder(func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the BrokerDataSetReaderTransportDataType or returns an error if something is wrong
	Build() (BrokerDataSetReaderTransportDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BrokerDataSetReaderTransportDataType
}

// NewBrokerDataSetReaderTransportDataTypeBuilder() creates a BrokerDataSetReaderTransportDataTypeBuilder
func NewBrokerDataSetReaderTransportDataTypeBuilder() BrokerDataSetReaderTransportDataTypeBuilder {
	return &_BrokerDataSetReaderTransportDataTypeBuilder{_BrokerDataSetReaderTransportDataType: new(_BrokerDataSetReaderTransportDataType)}
}

type _BrokerDataSetReaderTransportDataTypeBuilder struct {
	*_BrokerDataSetReaderTransportDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (BrokerDataSetReaderTransportDataTypeBuilder) = (*_BrokerDataSetReaderTransportDataTypeBuilder)(nil)

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._BrokerDataSetReaderTransportDataType
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithMandatoryFields(queueName PascalString, resourceUri PascalString, authenticationProfileUri PascalString, requestedDeliveryGuarantee BrokerTransportQualityOfService, metaDataQueueName PascalString) BrokerDataSetReaderTransportDataTypeBuilder {
	return b.WithQueueName(queueName).WithResourceUri(resourceUri).WithAuthenticationProfileUri(authenticationProfileUri).WithRequestedDeliveryGuarantee(requestedDeliveryGuarantee).WithMetaDataQueueName(metaDataQueueName)
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithQueueName(queueName PascalString) BrokerDataSetReaderTransportDataTypeBuilder {
	b.QueueName = queueName
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithQueueNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.QueueName.CreatePascalStringBuilder())
	var err error
	b.QueueName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithResourceUri(resourceUri PascalString) BrokerDataSetReaderTransportDataTypeBuilder {
	b.ResourceUri = resourceUri
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithResourceUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.ResourceUri.CreatePascalStringBuilder())
	var err error
	b.ResourceUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithAuthenticationProfileUri(authenticationProfileUri PascalString) BrokerDataSetReaderTransportDataTypeBuilder {
	b.AuthenticationProfileUri = authenticationProfileUri
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithAuthenticationProfileUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.AuthenticationProfileUri.CreatePascalStringBuilder())
	var err error
	b.AuthenticationProfileUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithRequestedDeliveryGuarantee(requestedDeliveryGuarantee BrokerTransportQualityOfService) BrokerDataSetReaderTransportDataTypeBuilder {
	b.RequestedDeliveryGuarantee = requestedDeliveryGuarantee
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithMetaDataQueueName(metaDataQueueName PascalString) BrokerDataSetReaderTransportDataTypeBuilder {
	b.MetaDataQueueName = metaDataQueueName
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) WithMetaDataQueueNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) BrokerDataSetReaderTransportDataTypeBuilder {
	builder := builderSupplier(b.MetaDataQueueName.CreatePascalStringBuilder())
	var err error
	b.MetaDataQueueName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) Build() (BrokerDataSetReaderTransportDataType, error) {
	if b.QueueName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'queueName' not set"))
	}
	if b.ResourceUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'resourceUri' not set"))
	}
	if b.AuthenticationProfileUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'authenticationProfileUri' not set"))
	}
	if b.MetaDataQueueName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'metaDataQueueName' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BrokerDataSetReaderTransportDataType.deepCopy(), nil
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) MustBuild() BrokerDataSetReaderTransportDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_BrokerDataSetReaderTransportDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateBrokerDataSetReaderTransportDataTypeBuilder().(*_BrokerDataSetReaderTransportDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBrokerDataSetReaderTransportDataTypeBuilder creates a BrokerDataSetReaderTransportDataTypeBuilder
func (b *_BrokerDataSetReaderTransportDataType) CreateBrokerDataSetReaderTransportDataTypeBuilder() BrokerDataSetReaderTransportDataTypeBuilder {
	if b == nil {
		return NewBrokerDataSetReaderTransportDataTypeBuilder()
	}
	return &_BrokerDataSetReaderTransportDataTypeBuilder{_BrokerDataSetReaderTransportDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrokerDataSetReaderTransportDataType) GetExtensionId() int32 {
	return int32(15672)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrokerDataSetReaderTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrokerDataSetReaderTransportDataType) GetQueueName() PascalString {
	return m.QueueName
}

func (m *_BrokerDataSetReaderTransportDataType) GetResourceUri() PascalString {
	return m.ResourceUri
}

func (m *_BrokerDataSetReaderTransportDataType) GetAuthenticationProfileUri() PascalString {
	return m.AuthenticationProfileUri
}

func (m *_BrokerDataSetReaderTransportDataType) GetRequestedDeliveryGuarantee() BrokerTransportQualityOfService {
	return m.RequestedDeliveryGuarantee
}

func (m *_BrokerDataSetReaderTransportDataType) GetMetaDataQueueName() PascalString {
	return m.MetaDataQueueName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBrokerDataSetReaderTransportDataType(structType any) BrokerDataSetReaderTransportDataType {
	if casted, ok := structType.(BrokerDataSetReaderTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*BrokerDataSetReaderTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_BrokerDataSetReaderTransportDataType) GetTypeName() string {
	return "BrokerDataSetReaderTransportDataType"
}

func (m *_BrokerDataSetReaderTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (queueName)
	lengthInBits += m.QueueName.GetLengthInBits(ctx)

	// Simple field (resourceUri)
	lengthInBits += m.ResourceUri.GetLengthInBits(ctx)

	// Simple field (authenticationProfileUri)
	lengthInBits += m.AuthenticationProfileUri.GetLengthInBits(ctx)

	// Simple field (requestedDeliveryGuarantee)
	lengthInBits += 32

	// Simple field (metaDataQueueName)
	lengthInBits += m.MetaDataQueueName.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BrokerDataSetReaderTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrokerDataSetReaderTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__brokerDataSetReaderTransportDataType BrokerDataSetReaderTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrokerDataSetReaderTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrokerDataSetReaderTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	queueName, err := ReadSimpleField[PascalString](ctx, "queueName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'queueName' field"))
	}
	m.QueueName = queueName

	resourceUri, err := ReadSimpleField[PascalString](ctx, "resourceUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resourceUri' field"))
	}
	m.ResourceUri = resourceUri

	authenticationProfileUri, err := ReadSimpleField[PascalString](ctx, "authenticationProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationProfileUri' field"))
	}
	m.AuthenticationProfileUri = authenticationProfileUri

	requestedDeliveryGuarantee, err := ReadEnumField[BrokerTransportQualityOfService](ctx, "requestedDeliveryGuarantee", "BrokerTransportQualityOfService", ReadEnum(BrokerTransportQualityOfServiceByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedDeliveryGuarantee' field"))
	}
	m.RequestedDeliveryGuarantee = requestedDeliveryGuarantee

	metaDataQueueName, err := ReadSimpleField[PascalString](ctx, "metaDataQueueName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'metaDataQueueName' field"))
	}
	m.MetaDataQueueName = metaDataQueueName

	if closeErr := readBuffer.CloseContext("BrokerDataSetReaderTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrokerDataSetReaderTransportDataType")
	}

	return m, nil
}

func (m *_BrokerDataSetReaderTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrokerDataSetReaderTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrokerDataSetReaderTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrokerDataSetReaderTransportDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "queueName", m.GetQueueName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'queueName' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "resourceUri", m.GetResourceUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'resourceUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "authenticationProfileUri", m.GetAuthenticationProfileUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationProfileUri' field")
		}

		if err := WriteSimpleEnumField[BrokerTransportQualityOfService](ctx, "requestedDeliveryGuarantee", "BrokerTransportQualityOfService", m.GetRequestedDeliveryGuarantee(), WriteEnum[BrokerTransportQualityOfService, uint32](BrokerTransportQualityOfService.GetValue, BrokerTransportQualityOfService.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedDeliveryGuarantee' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "metaDataQueueName", m.GetMetaDataQueueName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'metaDataQueueName' field")
		}

		if popErr := writeBuffer.PopContext("BrokerDataSetReaderTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrokerDataSetReaderTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrokerDataSetReaderTransportDataType) IsBrokerDataSetReaderTransportDataType() {}

func (m *_BrokerDataSetReaderTransportDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BrokerDataSetReaderTransportDataType) deepCopy() *_BrokerDataSetReaderTransportDataType {
	if m == nil {
		return nil
	}
	_BrokerDataSetReaderTransportDataTypeCopy := &_BrokerDataSetReaderTransportDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.QueueName),
		utils.DeepCopy[PascalString](m.ResourceUri),
		utils.DeepCopy[PascalString](m.AuthenticationProfileUri),
		m.RequestedDeliveryGuarantee,
		utils.DeepCopy[PascalString](m.MetaDataQueueName),
	}
	_BrokerDataSetReaderTransportDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _BrokerDataSetReaderTransportDataTypeCopy
}

func (m *_BrokerDataSetReaderTransportDataType) String() string {
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
