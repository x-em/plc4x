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

// MediaTransportControlData is the corresponding interface of MediaTransportControlData
type MediaTransportControlData interface {
	MediaTransportControlDataContract
	MediaTransportControlDataRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsMediaTransportControlData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlData()
	// CreateBuilder creates a MediaTransportControlDataBuilder
	CreateMediaTransportControlDataBuilder() MediaTransportControlDataBuilder
}

// MediaTransportControlDataContract provides a set of functions which can be overwritten by a sub struct
type MediaTransportControlDataContract interface {
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() MediaTransportControlCommandTypeContainer
	// GetMediaLinkGroup returns MediaLinkGroup (property field)
	GetMediaLinkGroup() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() MediaTransportControlCommandType
	// IsMediaTransportControlData is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlData()
	// CreateBuilder creates a MediaTransportControlDataBuilder
	CreateMediaTransportControlDataBuilder() MediaTransportControlDataBuilder
}

// MediaTransportControlDataRequirements provides a set of functions which need to be implemented by a sub struct
type MediaTransportControlDataRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetCommandType returns CommandType (discriminator field)
	GetCommandType() MediaTransportControlCommandType
}

// _MediaTransportControlData is the data-structure of this message
type _MediaTransportControlData struct {
	_SubType interface {
		MediaTransportControlDataContract
		MediaTransportControlDataRequirements
	}
	CommandTypeContainer MediaTransportControlCommandTypeContainer
	MediaLinkGroup       byte
}

var _ MediaTransportControlDataContract = (*_MediaTransportControlData)(nil)

// NewMediaTransportControlData factory function for _MediaTransportControlData
func NewMediaTransportControlData(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlData {
	return &_MediaTransportControlData{CommandTypeContainer: commandTypeContainer, MediaLinkGroup: mediaLinkGroup}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MediaTransportControlDataBuilder is a builder for MediaTransportControlData
type MediaTransportControlDataBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) MediaTransportControlDataBuilder
	// WithCommandTypeContainer adds CommandTypeContainer (property field)
	WithCommandTypeContainer(MediaTransportControlCommandTypeContainer) MediaTransportControlDataBuilder
	// WithMediaLinkGroup adds MediaLinkGroup (property field)
	WithMediaLinkGroup(byte) MediaTransportControlDataBuilder
	// AsMediaTransportControlDataStop converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataStop() MediaTransportControlDataStopBuilder
	// AsMediaTransportControlDataPlay converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataPlay() MediaTransportControlDataPlayBuilder
	// AsMediaTransportControlDataPauseResume converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataPauseResume() MediaTransportControlDataPauseResumeBuilder
	// AsMediaTransportControlDataSetCategory converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataSetCategory() MediaTransportControlDataSetCategoryBuilder
	// AsMediaTransportControlDataSetSelection converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataSetSelection() MediaTransportControlDataSetSelectionBuilder
	// AsMediaTransportControlDataSetTrack converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataSetTrack() MediaTransportControlDataSetTrackBuilder
	// AsMediaTransportControlDataShuffleOnOff converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataShuffleOnOff() MediaTransportControlDataShuffleOnOffBuilder
	// AsMediaTransportControlDataRepeatOnOff converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataRepeatOnOff() MediaTransportControlDataRepeatOnOffBuilder
	// AsMediaTransportControlDataNextPreviousCategory converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataNextPreviousCategory() MediaTransportControlDataNextPreviousCategoryBuilder
	// AsMediaTransportControlDataNextPreviousSelection converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataNextPreviousSelection() MediaTransportControlDataNextPreviousSelectionBuilder
	// AsMediaTransportControlDataNextPreviousTrack converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataNextPreviousTrack() MediaTransportControlDataNextPreviousTrackBuilder
	// AsMediaTransportControlDataFastForward converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataFastForward() MediaTransportControlDataFastForwardBuilder
	// AsMediaTransportControlDataRewind converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataRewind() MediaTransportControlDataRewindBuilder
	// AsMediaTransportControlDataSourcePowerControl converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataSourcePowerControl() MediaTransportControlDataSourcePowerControlBuilder
	// AsMediaTransportControlDataTotalTracks converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataTotalTracks() MediaTransportControlDataTotalTracksBuilder
	// AsMediaTransportControlDataStatusRequest converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataStatusRequest() MediaTransportControlDataStatusRequestBuilder
	// AsMediaTransportControlDataEnumerateCategoriesSelectionTracks converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataEnumerateCategoriesSelectionTracks() MediaTransportControlDataEnumerateCategoriesSelectionTracksBuilder
	// AsMediaTransportControlDataEnumerationsSize converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataEnumerationsSize() MediaTransportControlDataEnumerationsSizeBuilder
	// AsMediaTransportControlDataTrackName converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataTrackName() MediaTransportControlDataTrackNameBuilder
	// AsMediaTransportControlDataSelectionName converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataSelectionName() MediaTransportControlDataSelectionNameBuilder
	// AsMediaTransportControlDataCategoryName converts this build to a subType of MediaTransportControlData. It is always possible to return to current builder using Done()
	AsMediaTransportControlDataCategoryName() MediaTransportControlDataCategoryNameBuilder
	// Build builds the MediaTransportControlData or returns an error if something is wrong
	PartialBuild() (MediaTransportControlDataContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() MediaTransportControlDataContract
	// Build builds the MediaTransportControlData or returns an error if something is wrong
	Build() (MediaTransportControlData, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MediaTransportControlData
}

// NewMediaTransportControlDataBuilder() creates a MediaTransportControlDataBuilder
func NewMediaTransportControlDataBuilder() MediaTransportControlDataBuilder {
	return &_MediaTransportControlDataBuilder{_MediaTransportControlData: new(_MediaTransportControlData)}
}

type _MediaTransportControlDataChildBuilder interface {
	utils.Copyable
	setParent(MediaTransportControlDataContract)
	buildForMediaTransportControlData() (MediaTransportControlData, error)
}

type _MediaTransportControlDataBuilder struct {
	*_MediaTransportControlData

	childBuilder _MediaTransportControlDataChildBuilder

	err *utils.MultiError
}

var _ (MediaTransportControlDataBuilder) = (*_MediaTransportControlDataBuilder)(nil)

func (b *_MediaTransportControlDataBuilder) WithMandatoryFields(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) MediaTransportControlDataBuilder {
	return b.WithCommandTypeContainer(commandTypeContainer).WithMediaLinkGroup(mediaLinkGroup)
}

func (b *_MediaTransportControlDataBuilder) WithCommandTypeContainer(commandTypeContainer MediaTransportControlCommandTypeContainer) MediaTransportControlDataBuilder {
	b.CommandTypeContainer = commandTypeContainer
	return b
}

func (b *_MediaTransportControlDataBuilder) WithMediaLinkGroup(mediaLinkGroup byte) MediaTransportControlDataBuilder {
	b.MediaLinkGroup = mediaLinkGroup
	return b
}

func (b *_MediaTransportControlDataBuilder) PartialBuild() (MediaTransportControlDataContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MediaTransportControlData.deepCopy(), nil
}

func (b *_MediaTransportControlDataBuilder) PartialMustBuild() MediaTransportControlDataContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataStop() MediaTransportControlDataStopBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataStopBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataStopBuilder().(*_MediaTransportControlDataStopBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataPlay() MediaTransportControlDataPlayBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataPlayBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataPlayBuilder().(*_MediaTransportControlDataPlayBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataPauseResume() MediaTransportControlDataPauseResumeBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataPauseResumeBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataPauseResumeBuilder().(*_MediaTransportControlDataPauseResumeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataSetCategory() MediaTransportControlDataSetCategoryBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataSetCategoryBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataSetCategoryBuilder().(*_MediaTransportControlDataSetCategoryBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataSetSelection() MediaTransportControlDataSetSelectionBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataSetSelectionBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataSetSelectionBuilder().(*_MediaTransportControlDataSetSelectionBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataSetTrack() MediaTransportControlDataSetTrackBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataSetTrackBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataSetTrackBuilder().(*_MediaTransportControlDataSetTrackBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataShuffleOnOff() MediaTransportControlDataShuffleOnOffBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataShuffleOnOffBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataShuffleOnOffBuilder().(*_MediaTransportControlDataShuffleOnOffBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataRepeatOnOff() MediaTransportControlDataRepeatOnOffBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataRepeatOnOffBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataRepeatOnOffBuilder().(*_MediaTransportControlDataRepeatOnOffBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataNextPreviousCategory() MediaTransportControlDataNextPreviousCategoryBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataNextPreviousCategoryBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataNextPreviousCategoryBuilder().(*_MediaTransportControlDataNextPreviousCategoryBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataNextPreviousSelection() MediaTransportControlDataNextPreviousSelectionBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataNextPreviousSelectionBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataNextPreviousSelectionBuilder().(*_MediaTransportControlDataNextPreviousSelectionBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataNextPreviousTrack() MediaTransportControlDataNextPreviousTrackBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataNextPreviousTrackBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataNextPreviousTrackBuilder().(*_MediaTransportControlDataNextPreviousTrackBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataFastForward() MediaTransportControlDataFastForwardBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataFastForwardBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataFastForwardBuilder().(*_MediaTransportControlDataFastForwardBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataRewind() MediaTransportControlDataRewindBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataRewindBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataRewindBuilder().(*_MediaTransportControlDataRewindBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataSourcePowerControl() MediaTransportControlDataSourcePowerControlBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataSourcePowerControlBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataSourcePowerControlBuilder().(*_MediaTransportControlDataSourcePowerControlBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataTotalTracks() MediaTransportControlDataTotalTracksBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataTotalTracksBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataTotalTracksBuilder().(*_MediaTransportControlDataTotalTracksBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataStatusRequest() MediaTransportControlDataStatusRequestBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataStatusRequestBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataStatusRequestBuilder().(*_MediaTransportControlDataStatusRequestBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataEnumerateCategoriesSelectionTracks() MediaTransportControlDataEnumerateCategoriesSelectionTracksBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataEnumerateCategoriesSelectionTracksBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataEnumerateCategoriesSelectionTracksBuilder().(*_MediaTransportControlDataEnumerateCategoriesSelectionTracksBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataEnumerationsSize() MediaTransportControlDataEnumerationsSizeBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataEnumerationsSizeBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataEnumerationsSizeBuilder().(*_MediaTransportControlDataEnumerationsSizeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataTrackName() MediaTransportControlDataTrackNameBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataTrackNameBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataTrackNameBuilder().(*_MediaTransportControlDataTrackNameBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataSelectionName() MediaTransportControlDataSelectionNameBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataSelectionNameBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataSelectionNameBuilder().(*_MediaTransportControlDataSelectionNameBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) AsMediaTransportControlDataCategoryName() MediaTransportControlDataCategoryNameBuilder {
	if cb, ok := b.childBuilder.(MediaTransportControlDataCategoryNameBuilder); ok {
		return cb
	}
	cb := NewMediaTransportControlDataCategoryNameBuilder().(*_MediaTransportControlDataCategoryNameBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_MediaTransportControlDataBuilder) Build() (MediaTransportControlData, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForMediaTransportControlData()
}

func (b *_MediaTransportControlDataBuilder) MustBuild() MediaTransportControlData {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_MediaTransportControlDataBuilder) DeepCopy() any {
	_copy := b.CreateMediaTransportControlDataBuilder().(*_MediaTransportControlDataBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_MediaTransportControlDataChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMediaTransportControlDataBuilder creates a MediaTransportControlDataBuilder
func (b *_MediaTransportControlData) CreateMediaTransportControlDataBuilder() MediaTransportControlDataBuilder {
	if b == nil {
		return NewMediaTransportControlDataBuilder()
	}
	return &_MediaTransportControlDataBuilder{_MediaTransportControlData: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlData) GetCommandTypeContainer() MediaTransportControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_MediaTransportControlData) GetMediaLinkGroup() byte {
	return m.MediaLinkGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_MediaTransportControlData) GetCommandType() MediaTransportControlCommandType {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return CastMediaTransportControlCommandType(m.GetCommandTypeContainer().CommandType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMediaTransportControlData(structType any) MediaTransportControlData {
	if casted, ok := structType.(MediaTransportControlData); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlData); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlData) GetTypeName() string {
	return "MediaTransportControlData"
}

func (m *_MediaTransportControlData) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (mediaLinkGroup)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlData) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_MediaTransportControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataParse[T MediaTransportControlData](ctx context.Context, theBytes []byte) (T, error) {
	return MediaTransportControlDataParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataParseWithBufferProducer[T MediaTransportControlData]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := MediaTransportControlDataParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func MediaTransportControlDataParseWithBuffer[T MediaTransportControlData](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_MediaTransportControlData{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_MediaTransportControlData) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__mediaTransportControlData MediaTransportControlData, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsMediaTransportControlCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "no command type could be found"})
	}

	commandTypeContainer, err := ReadEnumField[MediaTransportControlCommandTypeContainer](ctx, "commandTypeContainer", "MediaTransportControlCommandTypeContainer", ReadEnum(MediaTransportControlCommandTypeContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandTypeContainer' field"))
	}
	m.CommandTypeContainer = commandTypeContainer

	commandType, err := ReadVirtualField[MediaTransportControlCommandType](ctx, "commandType", (*MediaTransportControlCommandType)(nil), commandTypeContainer.CommandType())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'commandType' field"))
	}
	_ = commandType

	mediaLinkGroup, err := ReadSimpleField(ctx, "mediaLinkGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mediaLinkGroup' field"))
	}
	m.MediaLinkGroup = mediaLinkGroup

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child MediaTransportControlData
	switch {
	case commandType == MediaTransportControlCommandType_STOP: // MediaTransportControlDataStop
		if _child, err = new(_MediaTransportControlDataStop).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataStop for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_PLAY: // MediaTransportControlDataPlay
		if _child, err = new(_MediaTransportControlDataPlay).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataPlay for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_PAUSE_RESUME: // MediaTransportControlDataPauseResume
		if _child, err = new(_MediaTransportControlDataPauseResume).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataPauseResume for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECT_CATEGORY: // MediaTransportControlDataSetCategory
		if _child, err = new(_MediaTransportControlDataSetCategory).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSetCategory for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECT_SELECTION: // MediaTransportControlDataSetSelection
		if _child, err = new(_MediaTransportControlDataSetSelection).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSetSelection for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECT_TRACK: // MediaTransportControlDataSetTrack
		if _child, err = new(_MediaTransportControlDataSetTrack).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSetTrack for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SHUFFLE_ON_OFF: // MediaTransportControlDataShuffleOnOff
		if _child, err = new(_MediaTransportControlDataShuffleOnOff).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataShuffleOnOff for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_REPEAT_ON_OFF: // MediaTransportControlDataRepeatOnOff
		if _child, err = new(_MediaTransportControlDataRepeatOnOff).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataRepeatOnOff for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_CATEGORY: // MediaTransportControlDataNextPreviousCategory
		if _child, err = new(_MediaTransportControlDataNextPreviousCategory).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataNextPreviousCategory for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_SELECTION: // MediaTransportControlDataNextPreviousSelection
		if _child, err = new(_MediaTransportControlDataNextPreviousSelection).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataNextPreviousSelection for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_NEXT_PREVIOUS_TRACK: // MediaTransportControlDataNextPreviousTrack
		if _child, err = new(_MediaTransportControlDataNextPreviousTrack).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataNextPreviousTrack for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_FAST_FORWARD: // MediaTransportControlDataFastForward
		if _child, err = new(_MediaTransportControlDataFastForward).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataFastForward for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_REWIND: // MediaTransportControlDataRewind
		if _child, err = new(_MediaTransportControlDataRewind).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataRewind for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SOURCE_POWER_CONTROL: // MediaTransportControlDataSourcePowerControl
		if _child, err = new(_MediaTransportControlDataSourcePowerControl).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSourcePowerControl for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_TOTAL_TRACKS: // MediaTransportControlDataTotalTracks
		if _child, err = new(_MediaTransportControlDataTotalTracks).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataTotalTracks for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_STATUS_REQUEST: // MediaTransportControlDataStatusRequest
		if _child, err = new(_MediaTransportControlDataStatusRequest).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataStatusRequest for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_ENUMERATE_CATEGORIES_SELECTIONS_TRACKS: // MediaTransportControlDataEnumerateCategoriesSelectionTracks
		if _child, err = new(_MediaTransportControlDataEnumerateCategoriesSelectionTracks).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataEnumerateCategoriesSelectionTracks for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_ENUMERATION_SIZE: // MediaTransportControlDataEnumerationsSize
		if _child, err = new(_MediaTransportControlDataEnumerationsSize).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataEnumerationsSize for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_TRACK_NAME: // MediaTransportControlDataTrackName
		if _child, err = new(_MediaTransportControlDataTrackName).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataTrackName for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_SELECTION_NAME: // MediaTransportControlDataSelectionName
		if _child, err = new(_MediaTransportControlDataSelectionName).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataSelectionName for type-switch of MediaTransportControlData")
		}
	case commandType == MediaTransportControlCommandType_CATEGORY_NAME: // MediaTransportControlDataCategoryName
		if _child, err = new(_MediaTransportControlDataCategoryName).parse(ctx, readBuffer, m, commandTypeContainer); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MediaTransportControlDataCategoryName for type-switch of MediaTransportControlData")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}

	if closeErr := readBuffer.CloseContext("MediaTransportControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlData")
	}

	return _child, nil
}

func (pm *_MediaTransportControlData) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MediaTransportControlData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MediaTransportControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MediaTransportControlData")
	}

	if err := WriteSimpleEnumField[MediaTransportControlCommandTypeContainer](ctx, "commandTypeContainer", "MediaTransportControlCommandTypeContainer", m.GetCommandTypeContainer(), WriteEnum[MediaTransportControlCommandTypeContainer, uint8](MediaTransportControlCommandTypeContainer.GetValue, MediaTransportControlCommandTypeContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
		return errors.Wrap(err, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	if err := WriteSimpleField[byte](ctx, "mediaLinkGroup", m.GetMediaLinkGroup(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'mediaLinkGroup' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MediaTransportControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MediaTransportControlData")
	}
	return nil
}

func (m *_MediaTransportControlData) IsMediaTransportControlData() {}

func (m *_MediaTransportControlData) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MediaTransportControlData) deepCopy() *_MediaTransportControlData {
	if m == nil {
		return nil
	}
	_MediaTransportControlDataCopy := &_MediaTransportControlData{
		nil, // will be set by child
		m.CommandTypeContainer,
		m.MediaLinkGroup,
	}
	return _MediaTransportControlDataCopy
}