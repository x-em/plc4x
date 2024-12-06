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

// JsonNetworkMessageContentMask is an enum
type JsonNetworkMessageContentMask uint32

type IJsonNetworkMessageContentMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNone                 JsonNetworkMessageContentMask = 0
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNetworkMessageHeader JsonNetworkMessageContentMask = 1
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetMessageHeader JsonNetworkMessageContentMask = 2
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskSingleDataSetMessage JsonNetworkMessageContentMask = 4
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskPublisherId          JsonNetworkMessageContentMask = 8
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetClassId       JsonNetworkMessageContentMask = 16
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskReplyTo              JsonNetworkMessageContentMask = 32
	JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskWriterGroupName      JsonNetworkMessageContentMask = 64
)

var JsonNetworkMessageContentMaskValues []JsonNetworkMessageContentMask

func init() {
	_ = errors.New
	JsonNetworkMessageContentMaskValues = []JsonNetworkMessageContentMask{
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNone,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNetworkMessageHeader,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetMessageHeader,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskSingleDataSetMessage,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskPublisherId,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetClassId,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskReplyTo,
		JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskWriterGroupName,
	}
}

func JsonNetworkMessageContentMaskByValue(value uint32) (enum JsonNetworkMessageContentMask, ok bool) {
	switch value {
	case 0:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNone, true
	case 1:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNetworkMessageHeader, true
	case 16:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetClassId, true
	case 2:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetMessageHeader, true
	case 32:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskReplyTo, true
	case 4:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskSingleDataSetMessage, true
	case 64:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskWriterGroupName, true
	case 8:
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskPublisherId, true
	}
	return 0, false
}

func JsonNetworkMessageContentMaskByName(value string) (enum JsonNetworkMessageContentMask, ok bool) {
	switch value {
	case "jsonNetworkMessageContentMaskNone":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNone, true
	case "jsonNetworkMessageContentMaskNetworkMessageHeader":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNetworkMessageHeader, true
	case "jsonNetworkMessageContentMaskDataSetClassId":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetClassId, true
	case "jsonNetworkMessageContentMaskDataSetMessageHeader":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetMessageHeader, true
	case "jsonNetworkMessageContentMaskReplyTo":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskReplyTo, true
	case "jsonNetworkMessageContentMaskSingleDataSetMessage":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskSingleDataSetMessage, true
	case "jsonNetworkMessageContentMaskWriterGroupName":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskWriterGroupName, true
	case "jsonNetworkMessageContentMaskPublisherId":
		return JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskPublisherId, true
	}
	return 0, false
}

func JsonNetworkMessageContentMaskKnows(value uint32) bool {
	for _, typeValue := range JsonNetworkMessageContentMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastJsonNetworkMessageContentMask(structType any) JsonNetworkMessageContentMask {
	castFunc := func(typ any) JsonNetworkMessageContentMask {
		if sJsonNetworkMessageContentMask, ok := typ.(JsonNetworkMessageContentMask); ok {
			return sJsonNetworkMessageContentMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m JsonNetworkMessageContentMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m JsonNetworkMessageContentMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func JsonNetworkMessageContentMaskParse(ctx context.Context, theBytes []byte) (JsonNetworkMessageContentMask, error) {
	return JsonNetworkMessageContentMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func JsonNetworkMessageContentMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (JsonNetworkMessageContentMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("JsonNetworkMessageContentMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading JsonNetworkMessageContentMask")
	}
	if enum, ok := JsonNetworkMessageContentMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for JsonNetworkMessageContentMask")
		return JsonNetworkMessageContentMask(val), nil
	} else {
		return enum, nil
	}
}

func (e JsonNetworkMessageContentMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e JsonNetworkMessageContentMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("JsonNetworkMessageContentMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e JsonNetworkMessageContentMask) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e JsonNetworkMessageContentMask) PLC4XEnumName() string {
	switch e {
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNone:
		return "jsonNetworkMessageContentMaskNone"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskNetworkMessageHeader:
		return "jsonNetworkMessageContentMaskNetworkMessageHeader"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetClassId:
		return "jsonNetworkMessageContentMaskDataSetClassId"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskDataSetMessageHeader:
		return "jsonNetworkMessageContentMaskDataSetMessageHeader"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskReplyTo:
		return "jsonNetworkMessageContentMaskReplyTo"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskSingleDataSetMessage:
		return "jsonNetworkMessageContentMaskSingleDataSetMessage"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskWriterGroupName:
		return "jsonNetworkMessageContentMaskWriterGroupName"
	case JsonNetworkMessageContentMask_jsonNetworkMessageContentMaskPublisherId:
		return "jsonNetworkMessageContentMaskPublisherId"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e JsonNetworkMessageContentMask) String() string {
	return e.PLC4XEnumName()
}