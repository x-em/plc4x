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

// NLMUpdateKeyDistributionKey is the corresponding interface of NLMUpdateKeyDistributionKey
type NLMUpdateKeyDistributionKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NLM
	// GetKeyRevision returns KeyRevision (property field)
	GetKeyRevision() byte
	// GetKey returns Key (property field)
	GetKey() NLMUpdateKeyUpdateKeyEntry
	// IsNLMUpdateKeyDistributionKey is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNLMUpdateKeyDistributionKey()
	// CreateBuilder creates a NLMUpdateKeyDistributionKeyBuilder
	CreateNLMUpdateKeyDistributionKeyBuilder() NLMUpdateKeyDistributionKeyBuilder
}

// _NLMUpdateKeyDistributionKey is the data-structure of this message
type _NLMUpdateKeyDistributionKey struct {
	NLMContract
	KeyRevision byte
	Key         NLMUpdateKeyUpdateKeyEntry
}

var _ NLMUpdateKeyDistributionKey = (*_NLMUpdateKeyDistributionKey)(nil)
var _ NLMRequirements = (*_NLMUpdateKeyDistributionKey)(nil)

// NewNLMUpdateKeyDistributionKey factory function for _NLMUpdateKeyDistributionKey
func NewNLMUpdateKeyDistributionKey(keyRevision byte, key NLMUpdateKeyUpdateKeyEntry, apduLength uint16) *_NLMUpdateKeyDistributionKey {
	if key == nil {
		panic("key of type NLMUpdateKeyUpdateKeyEntry for NLMUpdateKeyDistributionKey must not be nil")
	}
	_result := &_NLMUpdateKeyDistributionKey{
		NLMContract: NewNLM(apduLength),
		KeyRevision: keyRevision,
		Key:         key,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NLMUpdateKeyDistributionKeyBuilder is a builder for NLMUpdateKeyDistributionKey
type NLMUpdateKeyDistributionKeyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(keyRevision byte, key NLMUpdateKeyUpdateKeyEntry) NLMUpdateKeyDistributionKeyBuilder
	// WithKeyRevision adds KeyRevision (property field)
	WithKeyRevision(byte) NLMUpdateKeyDistributionKeyBuilder
	// WithKey adds Key (property field)
	WithKey(NLMUpdateKeyUpdateKeyEntry) NLMUpdateKeyDistributionKeyBuilder
	// WithKeyBuilder adds Key (property field) which is build by the builder
	WithKeyBuilder(func(NLMUpdateKeyUpdateKeyEntryBuilder) NLMUpdateKeyUpdateKeyEntryBuilder) NLMUpdateKeyDistributionKeyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NLMBuilder
	// Build builds the NLMUpdateKeyDistributionKey or returns an error if something is wrong
	Build() (NLMUpdateKeyDistributionKey, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NLMUpdateKeyDistributionKey
}

// NewNLMUpdateKeyDistributionKeyBuilder() creates a NLMUpdateKeyDistributionKeyBuilder
func NewNLMUpdateKeyDistributionKeyBuilder() NLMUpdateKeyDistributionKeyBuilder {
	return &_NLMUpdateKeyDistributionKeyBuilder{_NLMUpdateKeyDistributionKey: new(_NLMUpdateKeyDistributionKey)}
}

type _NLMUpdateKeyDistributionKeyBuilder struct {
	*_NLMUpdateKeyDistributionKey

	parentBuilder *_NLMBuilder

	err *utils.MultiError
}

var _ (NLMUpdateKeyDistributionKeyBuilder) = (*_NLMUpdateKeyDistributionKeyBuilder)(nil)

func (b *_NLMUpdateKeyDistributionKeyBuilder) setParent(contract NLMContract) {
	b.NLMContract = contract
	contract.(*_NLM)._SubType = b._NLMUpdateKeyDistributionKey
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) WithMandatoryFields(keyRevision byte, key NLMUpdateKeyUpdateKeyEntry) NLMUpdateKeyDistributionKeyBuilder {
	return b.WithKeyRevision(keyRevision).WithKey(key)
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) WithKeyRevision(keyRevision byte) NLMUpdateKeyDistributionKeyBuilder {
	b.KeyRevision = keyRevision
	return b
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) WithKey(key NLMUpdateKeyUpdateKeyEntry) NLMUpdateKeyDistributionKeyBuilder {
	b.Key = key
	return b
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) WithKeyBuilder(builderSupplier func(NLMUpdateKeyUpdateKeyEntryBuilder) NLMUpdateKeyUpdateKeyEntryBuilder) NLMUpdateKeyDistributionKeyBuilder {
	builder := builderSupplier(b.Key.CreateNLMUpdateKeyUpdateKeyEntryBuilder())
	var err error
	b.Key, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NLMUpdateKeyUpdateKeyEntryBuilder failed"))
	}
	return b
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) Build() (NLMUpdateKeyDistributionKey, error) {
	if b.Key == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'key' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NLMUpdateKeyDistributionKey.deepCopy(), nil
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) MustBuild() NLMUpdateKeyDistributionKey {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) Done() NLMBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNLMBuilder().(*_NLMBuilder)
	}
	return b.parentBuilder
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) buildForNLM() (NLM, error) {
	return b.Build()
}

func (b *_NLMUpdateKeyDistributionKeyBuilder) DeepCopy() any {
	_copy := b.CreateNLMUpdateKeyDistributionKeyBuilder().(*_NLMUpdateKeyDistributionKeyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNLMUpdateKeyDistributionKeyBuilder creates a NLMUpdateKeyDistributionKeyBuilder
func (b *_NLMUpdateKeyDistributionKey) CreateNLMUpdateKeyDistributionKeyBuilder() NLMUpdateKeyDistributionKeyBuilder {
	if b == nil {
		return NewNLMUpdateKeyDistributionKeyBuilder()
	}
	return &_NLMUpdateKeyDistributionKeyBuilder{_NLMUpdateKeyDistributionKey: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMUpdateKeyDistributionKey) GetMessageType() uint8 {
	return 0x0F
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMUpdateKeyDistributionKey) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMUpdateKeyDistributionKey) GetKeyRevision() byte {
	return m.KeyRevision
}

func (m *_NLMUpdateKeyDistributionKey) GetKey() NLMUpdateKeyUpdateKeyEntry {
	return m.Key
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNLMUpdateKeyDistributionKey(structType any) NLMUpdateKeyDistributionKey {
	if casted, ok := structType.(NLMUpdateKeyDistributionKey); ok {
		return casted
	}
	if casted, ok := structType.(*NLMUpdateKeyDistributionKey); ok {
		return *casted
	}
	return nil
}

func (m *_NLMUpdateKeyDistributionKey) GetTypeName() string {
	return "NLMUpdateKeyDistributionKey"
}

func (m *_NLMUpdateKeyDistributionKey) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (keyRevision)
	lengthInBits += 8

	// Simple field (key)
	lengthInBits += m.Key.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_NLMUpdateKeyDistributionKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMUpdateKeyDistributionKey) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMUpdateKeyDistributionKey NLMUpdateKeyDistributionKey, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMUpdateKeyDistributionKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMUpdateKeyDistributionKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	keyRevision, err := ReadSimpleField(ctx, "keyRevision", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'keyRevision' field"))
	}
	m.KeyRevision = keyRevision

	key, err := ReadSimpleField[NLMUpdateKeyUpdateKeyEntry](ctx, "key", ReadComplex[NLMUpdateKeyUpdateKeyEntry](NLMUpdateKeyUpdateKeyEntryParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'key' field"))
	}
	m.Key = key

	if closeErr := readBuffer.CloseContext("NLMUpdateKeyDistributionKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMUpdateKeyDistributionKey")
	}

	return m, nil
}

func (m *_NLMUpdateKeyDistributionKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMUpdateKeyDistributionKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMUpdateKeyDistributionKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMUpdateKeyDistributionKey")
		}

		if err := WriteSimpleField[byte](ctx, "keyRevision", m.GetKeyRevision(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'keyRevision' field")
		}

		if err := WriteSimpleField[NLMUpdateKeyUpdateKeyEntry](ctx, "key", m.GetKey(), WriteComplex[NLMUpdateKeyUpdateKeyEntry](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'key' field")
		}

		if popErr := writeBuffer.PopContext("NLMUpdateKeyDistributionKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMUpdateKeyDistributionKey")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMUpdateKeyDistributionKey) IsNLMUpdateKeyDistributionKey() {}

func (m *_NLMUpdateKeyDistributionKey) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NLMUpdateKeyDistributionKey) deepCopy() *_NLMUpdateKeyDistributionKey {
	if m == nil {
		return nil
	}
	_NLMUpdateKeyDistributionKeyCopy := &_NLMUpdateKeyDistributionKey{
		m.NLMContract.(*_NLM).deepCopy(),
		m.KeyRevision,
		utils.DeepCopy[NLMUpdateKeyUpdateKeyEntry](m.Key),
	}
	_NLMUpdateKeyDistributionKeyCopy.NLMContract.(*_NLM)._SubType = m
	return _NLMUpdateKeyDistributionKeyCopy
}

func (m *_NLMUpdateKeyDistributionKey) String() string {
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