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

// BACnetActionList is the corresponding interface of BACnetActionList
type BACnetActionList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetAction returns Action (property field)
	GetAction() []BACnetActionCommand
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetActionList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetActionList()
	// CreateBuilder creates a BACnetActionListBuilder
	CreateBACnetActionListBuilder() BACnetActionListBuilder
}

// _BACnetActionList is the data-structure of this message
type _BACnetActionList struct {
	InnerOpeningTag BACnetOpeningTag
	Action          []BACnetActionCommand
	InnerClosingTag BACnetClosingTag
}

var _ BACnetActionList = (*_BACnetActionList)(nil)

// NewBACnetActionList factory function for _BACnetActionList
func NewBACnetActionList(innerOpeningTag BACnetOpeningTag, action []BACnetActionCommand, innerClosingTag BACnetClosingTag) *_BACnetActionList {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetActionList must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetActionList must not be nil")
	}
	return &_BACnetActionList{InnerOpeningTag: innerOpeningTag, Action: action, InnerClosingTag: innerClosingTag}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetActionListBuilder is a builder for BACnetActionList
type BACnetActionListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, action []BACnetActionCommand, innerClosingTag BACnetClosingTag) BACnetActionListBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetActionListBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetActionListBuilder
	// WithAction adds Action (property field)
	WithAction(...BACnetActionCommand) BACnetActionListBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetActionListBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetActionListBuilder
	// Build builds the BACnetActionList or returns an error if something is wrong
	Build() (BACnetActionList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetActionList
}

// NewBACnetActionListBuilder() creates a BACnetActionListBuilder
func NewBACnetActionListBuilder() BACnetActionListBuilder {
	return &_BACnetActionListBuilder{_BACnetActionList: new(_BACnetActionList)}
}

type _BACnetActionListBuilder struct {
	*_BACnetActionList

	err *utils.MultiError
}

var _ (BACnetActionListBuilder) = (*_BACnetActionListBuilder)(nil)

func (b *_BACnetActionListBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, action []BACnetActionCommand, innerClosingTag BACnetClosingTag) BACnetActionListBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithAction(action...).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetActionListBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetActionListBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetActionListBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetActionListBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetActionListBuilder) WithAction(action ...BACnetActionCommand) BACnetActionListBuilder {
	b.Action = action
	return b
}

func (b *_BACnetActionListBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetActionListBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetActionListBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetActionListBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetActionListBuilder) Build() (BACnetActionList, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetActionList.deepCopy(), nil
}

func (b *_BACnetActionListBuilder) MustBuild() BACnetActionList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetActionListBuilder) DeepCopy() any {
	_copy := b.CreateBACnetActionListBuilder().(*_BACnetActionListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetActionListBuilder creates a BACnetActionListBuilder
func (b *_BACnetActionList) CreateBACnetActionListBuilder() BACnetActionListBuilder {
	if b == nil {
		return NewBACnetActionListBuilder()
	}
	return &_BACnetActionListBuilder{_BACnetActionList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetActionList) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetActionList) GetAction() []BACnetActionCommand {
	return m.Action
}

func (m *_BACnetActionList) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetActionList(structType any) BACnetActionList {
	if casted, ok := structType.(BACnetActionList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetActionList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetActionList) GetTypeName() string {
	return "BACnetActionList"
}

func (m *_BACnetActionList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.Action) > 0 {
		for _, element := range m.Action {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetActionList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetActionListParse(ctx context.Context, theBytes []byte) (BACnetActionList, error) {
	return BACnetActionListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetActionListParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionList, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionList, error) {
		return BACnetActionListParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetActionListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetActionList, error) {
	v, err := (&_BACnetActionList{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetActionList) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetActionList BACnetActionList, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetActionList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetActionList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	action, err := ReadTerminatedArrayField[BACnetActionCommand](ctx, "action", ReadComplex[BACnetActionCommand](BACnetActionCommandParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, 0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'action' field"))
	}
	m.Action = action

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetActionList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetActionList")
	}

	return m, nil
}

func (m *_BACnetActionList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetActionList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetActionList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetActionList")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "action", m.GetAction(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'action' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetActionList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetActionList")
	}
	return nil
}

func (m *_BACnetActionList) IsBACnetActionList() {}

func (m *_BACnetActionList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetActionList) deepCopy() *_BACnetActionList {
	if m == nil {
		return nil
	}
	_BACnetActionListCopy := &_BACnetActionList{
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopySlice[BACnetActionCommand, BACnetActionCommand](m.Action),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	return _BACnetActionListCopy
}

func (m *_BACnetActionList) String() string {
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
