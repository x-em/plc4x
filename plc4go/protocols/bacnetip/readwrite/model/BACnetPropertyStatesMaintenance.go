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

// BACnetPropertyStatesMaintenance is the corresponding interface of BACnetPropertyStatesMaintenance
type BACnetPropertyStatesMaintenance interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetMaintenance returns Maintenance (property field)
	GetMaintenance() BACnetMaintenanceTagged
	// IsBACnetPropertyStatesMaintenance is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesMaintenance()
	// CreateBuilder creates a BACnetPropertyStatesMaintenanceBuilder
	CreateBACnetPropertyStatesMaintenanceBuilder() BACnetPropertyStatesMaintenanceBuilder
}

// _BACnetPropertyStatesMaintenance is the data-structure of this message
type _BACnetPropertyStatesMaintenance struct {
	BACnetPropertyStatesContract
	Maintenance BACnetMaintenanceTagged
}

var _ BACnetPropertyStatesMaintenance = (*_BACnetPropertyStatesMaintenance)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesMaintenance)(nil)

// NewBACnetPropertyStatesMaintenance factory function for _BACnetPropertyStatesMaintenance
func NewBACnetPropertyStatesMaintenance(peekedTagHeader BACnetTagHeader, maintenance BACnetMaintenanceTagged) *_BACnetPropertyStatesMaintenance {
	if maintenance == nil {
		panic("maintenance of type BACnetMaintenanceTagged for BACnetPropertyStatesMaintenance must not be nil")
	}
	_result := &_BACnetPropertyStatesMaintenance{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		Maintenance:                  maintenance,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesMaintenanceBuilder is a builder for BACnetPropertyStatesMaintenance
type BACnetPropertyStatesMaintenanceBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maintenance BACnetMaintenanceTagged) BACnetPropertyStatesMaintenanceBuilder
	// WithMaintenance adds Maintenance (property field)
	WithMaintenance(BACnetMaintenanceTagged) BACnetPropertyStatesMaintenanceBuilder
	// WithMaintenanceBuilder adds Maintenance (property field) which is build by the builder
	WithMaintenanceBuilder(func(BACnetMaintenanceTaggedBuilder) BACnetMaintenanceTaggedBuilder) BACnetPropertyStatesMaintenanceBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesMaintenance or returns an error if something is wrong
	Build() (BACnetPropertyStatesMaintenance, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesMaintenance
}

// NewBACnetPropertyStatesMaintenanceBuilder() creates a BACnetPropertyStatesMaintenanceBuilder
func NewBACnetPropertyStatesMaintenanceBuilder() BACnetPropertyStatesMaintenanceBuilder {
	return &_BACnetPropertyStatesMaintenanceBuilder{_BACnetPropertyStatesMaintenance: new(_BACnetPropertyStatesMaintenance)}
}

type _BACnetPropertyStatesMaintenanceBuilder struct {
	*_BACnetPropertyStatesMaintenance

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesMaintenanceBuilder) = (*_BACnetPropertyStatesMaintenanceBuilder)(nil)

func (b *_BACnetPropertyStatesMaintenanceBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesMaintenance
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) WithMandatoryFields(maintenance BACnetMaintenanceTagged) BACnetPropertyStatesMaintenanceBuilder {
	return b.WithMaintenance(maintenance)
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) WithMaintenance(maintenance BACnetMaintenanceTagged) BACnetPropertyStatesMaintenanceBuilder {
	b.Maintenance = maintenance
	return b
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) WithMaintenanceBuilder(builderSupplier func(BACnetMaintenanceTaggedBuilder) BACnetMaintenanceTaggedBuilder) BACnetPropertyStatesMaintenanceBuilder {
	builder := builderSupplier(b.Maintenance.CreateBACnetMaintenanceTaggedBuilder())
	var err error
	b.Maintenance, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetMaintenanceTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) Build() (BACnetPropertyStatesMaintenance, error) {
	if b.Maintenance == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maintenance' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesMaintenance.deepCopy(), nil
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) MustBuild() BACnetPropertyStatesMaintenance {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesMaintenanceBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesMaintenanceBuilder().(*_BACnetPropertyStatesMaintenanceBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesMaintenanceBuilder creates a BACnetPropertyStatesMaintenanceBuilder
func (b *_BACnetPropertyStatesMaintenance) CreateBACnetPropertyStatesMaintenanceBuilder() BACnetPropertyStatesMaintenanceBuilder {
	if b == nil {
		return NewBACnetPropertyStatesMaintenanceBuilder()
	}
	return &_BACnetPropertyStatesMaintenanceBuilder{_BACnetPropertyStatesMaintenance: b.deepCopy()}
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

func (m *_BACnetPropertyStatesMaintenance) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesMaintenance) GetMaintenance() BACnetMaintenanceTagged {
	return m.Maintenance
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesMaintenance(structType any) BACnetPropertyStatesMaintenance {
	if casted, ok := structType.(BACnetPropertyStatesMaintenance); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesMaintenance); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesMaintenance) GetTypeName() string {
	return "BACnetPropertyStatesMaintenance"
}

func (m *_BACnetPropertyStatesMaintenance) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (maintenance)
	lengthInBits += m.Maintenance.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesMaintenance) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesMaintenance) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesMaintenance BACnetPropertyStatesMaintenance, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesMaintenance"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesMaintenance")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maintenance, err := ReadSimpleField[BACnetMaintenanceTagged](ctx, "maintenance", ReadComplex[BACnetMaintenanceTagged](BACnetMaintenanceTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maintenance' field"))
	}
	m.Maintenance = maintenance

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesMaintenance"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesMaintenance")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesMaintenance) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesMaintenance) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesMaintenance"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesMaintenance")
		}

		if err := WriteSimpleField[BACnetMaintenanceTagged](ctx, "maintenance", m.GetMaintenance(), WriteComplex[BACnetMaintenanceTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maintenance' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesMaintenance"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesMaintenance")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesMaintenance) IsBACnetPropertyStatesMaintenance() {}

func (m *_BACnetPropertyStatesMaintenance) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesMaintenance) deepCopy() *_BACnetPropertyStatesMaintenance {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesMaintenanceCopy := &_BACnetPropertyStatesMaintenance{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetMaintenanceTagged](m.Maintenance),
	}
	_BACnetPropertyStatesMaintenanceCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesMaintenanceCopy
}

func (m *_BACnetPropertyStatesMaintenance) String() string {
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