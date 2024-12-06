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

// MFuncPropCommandReq is the corresponding interface of MFuncPropCommandReq
type MFuncPropCommandReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsMFuncPropCommandReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMFuncPropCommandReq()
	// CreateBuilder creates a MFuncPropCommandReqBuilder
	CreateMFuncPropCommandReqBuilder() MFuncPropCommandReqBuilder
}

// _MFuncPropCommandReq is the data-structure of this message
type _MFuncPropCommandReq struct {
	CEMIContract
}

var _ MFuncPropCommandReq = (*_MFuncPropCommandReq)(nil)
var _ CEMIRequirements = (*_MFuncPropCommandReq)(nil)

// NewMFuncPropCommandReq factory function for _MFuncPropCommandReq
func NewMFuncPropCommandReq(size uint16) *_MFuncPropCommandReq {
	_result := &_MFuncPropCommandReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MFuncPropCommandReqBuilder is a builder for MFuncPropCommandReq
type MFuncPropCommandReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() MFuncPropCommandReqBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CEMIBuilder
	// Build builds the MFuncPropCommandReq or returns an error if something is wrong
	Build() (MFuncPropCommandReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MFuncPropCommandReq
}

// NewMFuncPropCommandReqBuilder() creates a MFuncPropCommandReqBuilder
func NewMFuncPropCommandReqBuilder() MFuncPropCommandReqBuilder {
	return &_MFuncPropCommandReqBuilder{_MFuncPropCommandReq: new(_MFuncPropCommandReq)}
}

type _MFuncPropCommandReqBuilder struct {
	*_MFuncPropCommandReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (MFuncPropCommandReqBuilder) = (*_MFuncPropCommandReqBuilder)(nil)

func (b *_MFuncPropCommandReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
	contract.(*_CEMI)._SubType = b._MFuncPropCommandReq
}

func (b *_MFuncPropCommandReqBuilder) WithMandatoryFields() MFuncPropCommandReqBuilder {
	return b
}

func (b *_MFuncPropCommandReqBuilder) Build() (MFuncPropCommandReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MFuncPropCommandReq.deepCopy(), nil
}

func (b *_MFuncPropCommandReqBuilder) MustBuild() MFuncPropCommandReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MFuncPropCommandReqBuilder) Done() CEMIBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCEMIBuilder().(*_CEMIBuilder)
	}
	return b.parentBuilder
}

func (b *_MFuncPropCommandReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_MFuncPropCommandReqBuilder) DeepCopy() any {
	_copy := b.CreateMFuncPropCommandReqBuilder().(*_MFuncPropCommandReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMFuncPropCommandReqBuilder creates a MFuncPropCommandReqBuilder
func (b *_MFuncPropCommandReq) CreateMFuncPropCommandReqBuilder() MFuncPropCommandReqBuilder {
	if b == nil {
		return NewMFuncPropCommandReqBuilder()
	}
	return &_MFuncPropCommandReqBuilder{_MFuncPropCommandReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropCommandReq) GetMessageCode() uint8 {
	return 0xF8
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropCommandReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastMFuncPropCommandReq(structType any) MFuncPropCommandReq {
	if casted, ok := structType.(MFuncPropCommandReq); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropCommandReq); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropCommandReq) GetTypeName() string {
	return "MFuncPropCommandReq"
}

func (m *_MFuncPropCommandReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_MFuncPropCommandReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MFuncPropCommandReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__mFuncPropCommandReq MFuncPropCommandReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropCommandReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropCommandReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropCommandReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropCommandReq")
	}

	return m, nil
}

func (m *_MFuncPropCommandReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MFuncPropCommandReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCommandReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropCommandReq")
		}

		if popErr := writeBuffer.PopContext("MFuncPropCommandReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropCommandReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MFuncPropCommandReq) IsMFuncPropCommandReq() {}

func (m *_MFuncPropCommandReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MFuncPropCommandReq) deepCopy() *_MFuncPropCommandReq {
	if m == nil {
		return nil
	}
	_MFuncPropCommandReqCopy := &_MFuncPropCommandReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	_MFuncPropCommandReqCopy.CEMIContract.(*_CEMI)._SubType = m
	return _MFuncPropCommandReqCopy
}

func (m *_MFuncPropCommandReq) String() string {
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