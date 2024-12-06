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

// DataSetFieldContentMask is an enum
type DataSetFieldContentMask uint32

type IDataSetFieldContentMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	DataSetFieldContentMask_dataSetFieldContentMaskNone              DataSetFieldContentMask = 0
	DataSetFieldContentMask_dataSetFieldContentMaskStatusCode        DataSetFieldContentMask = 1
	DataSetFieldContentMask_dataSetFieldContentMaskSourceTimestamp   DataSetFieldContentMask = 2
	DataSetFieldContentMask_dataSetFieldContentMaskServerTimestamp   DataSetFieldContentMask = 4
	DataSetFieldContentMask_dataSetFieldContentMaskSourcePicoSeconds DataSetFieldContentMask = 8
	DataSetFieldContentMask_dataSetFieldContentMaskServerPicoSeconds DataSetFieldContentMask = 16
	DataSetFieldContentMask_dataSetFieldContentMaskRawData           DataSetFieldContentMask = 32
)

var DataSetFieldContentMaskValues []DataSetFieldContentMask

func init() {
	_ = errors.New
	DataSetFieldContentMaskValues = []DataSetFieldContentMask{
		DataSetFieldContentMask_dataSetFieldContentMaskNone,
		DataSetFieldContentMask_dataSetFieldContentMaskStatusCode,
		DataSetFieldContentMask_dataSetFieldContentMaskSourceTimestamp,
		DataSetFieldContentMask_dataSetFieldContentMaskServerTimestamp,
		DataSetFieldContentMask_dataSetFieldContentMaskSourcePicoSeconds,
		DataSetFieldContentMask_dataSetFieldContentMaskServerPicoSeconds,
		DataSetFieldContentMask_dataSetFieldContentMaskRawData,
	}
}

func DataSetFieldContentMaskByValue(value uint32) (enum DataSetFieldContentMask, ok bool) {
	switch value {
	case 0:
		return DataSetFieldContentMask_dataSetFieldContentMaskNone, true
	case 1:
		return DataSetFieldContentMask_dataSetFieldContentMaskStatusCode, true
	case 16:
		return DataSetFieldContentMask_dataSetFieldContentMaskServerPicoSeconds, true
	case 2:
		return DataSetFieldContentMask_dataSetFieldContentMaskSourceTimestamp, true
	case 32:
		return DataSetFieldContentMask_dataSetFieldContentMaskRawData, true
	case 4:
		return DataSetFieldContentMask_dataSetFieldContentMaskServerTimestamp, true
	case 8:
		return DataSetFieldContentMask_dataSetFieldContentMaskSourcePicoSeconds, true
	}
	return 0, false
}

func DataSetFieldContentMaskByName(value string) (enum DataSetFieldContentMask, ok bool) {
	switch value {
	case "dataSetFieldContentMaskNone":
		return DataSetFieldContentMask_dataSetFieldContentMaskNone, true
	case "dataSetFieldContentMaskStatusCode":
		return DataSetFieldContentMask_dataSetFieldContentMaskStatusCode, true
	case "dataSetFieldContentMaskServerPicoSeconds":
		return DataSetFieldContentMask_dataSetFieldContentMaskServerPicoSeconds, true
	case "dataSetFieldContentMaskSourceTimestamp":
		return DataSetFieldContentMask_dataSetFieldContentMaskSourceTimestamp, true
	case "dataSetFieldContentMaskRawData":
		return DataSetFieldContentMask_dataSetFieldContentMaskRawData, true
	case "dataSetFieldContentMaskServerTimestamp":
		return DataSetFieldContentMask_dataSetFieldContentMaskServerTimestamp, true
	case "dataSetFieldContentMaskSourcePicoSeconds":
		return DataSetFieldContentMask_dataSetFieldContentMaskSourcePicoSeconds, true
	}
	return 0, false
}

func DataSetFieldContentMaskKnows(value uint32) bool {
	for _, typeValue := range DataSetFieldContentMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDataSetFieldContentMask(structType any) DataSetFieldContentMask {
	castFunc := func(typ any) DataSetFieldContentMask {
		if sDataSetFieldContentMask, ok := typ.(DataSetFieldContentMask); ok {
			return sDataSetFieldContentMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m DataSetFieldContentMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m DataSetFieldContentMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DataSetFieldContentMaskParse(ctx context.Context, theBytes []byte) (DataSetFieldContentMask, error) {
	return DataSetFieldContentMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DataSetFieldContentMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DataSetFieldContentMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("DataSetFieldContentMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading DataSetFieldContentMask")
	}
	if enum, ok := DataSetFieldContentMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for DataSetFieldContentMask")
		return DataSetFieldContentMask(val), nil
	} else {
		return enum, nil
	}
}

func (e DataSetFieldContentMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e DataSetFieldContentMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("DataSetFieldContentMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e DataSetFieldContentMask) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e DataSetFieldContentMask) PLC4XEnumName() string {
	switch e {
	case DataSetFieldContentMask_dataSetFieldContentMaskNone:
		return "dataSetFieldContentMaskNone"
	case DataSetFieldContentMask_dataSetFieldContentMaskStatusCode:
		return "dataSetFieldContentMaskStatusCode"
	case DataSetFieldContentMask_dataSetFieldContentMaskServerPicoSeconds:
		return "dataSetFieldContentMaskServerPicoSeconds"
	case DataSetFieldContentMask_dataSetFieldContentMaskSourceTimestamp:
		return "dataSetFieldContentMaskSourceTimestamp"
	case DataSetFieldContentMask_dataSetFieldContentMaskRawData:
		return "dataSetFieldContentMaskRawData"
	case DataSetFieldContentMask_dataSetFieldContentMaskServerTimestamp:
		return "dataSetFieldContentMaskServerTimestamp"
	case DataSetFieldContentMask_dataSetFieldContentMaskSourcePicoSeconds:
		return "dataSetFieldContentMaskSourcePicoSeconds"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e DataSetFieldContentMask) String() string {
	return e.PLC4XEnumName()
}