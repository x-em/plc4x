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

// TimeBase is an enum
type TimeBase uint8

type ITimeBase interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TimeBase_B01SEC TimeBase = 0x00
	TimeBase_B1SEC  TimeBase = 0x01
	TimeBase_B10SEC TimeBase = 0x02
)

var TimeBaseValues []TimeBase

func init() {
	_ = errors.New
	TimeBaseValues = []TimeBase{
		TimeBase_B01SEC,
		TimeBase_B1SEC,
		TimeBase_B10SEC,
	}
}

func TimeBaseByValue(value uint8) (enum TimeBase, ok bool) {
	switch value {
	case 0x02:
		return TimeBase_B10SEC, true
	case 0x00:
		return TimeBase_B01SEC, true
	case 0x01:
		return TimeBase_B1SEC, true
	}
	return 0, false
}

func TimeBaseByName(value string) (enum TimeBase, ok bool) {
	switch value {
	case "B10SEC":
		return TimeBase_B10SEC, true
	case "B01SEC":
		return TimeBase_B01SEC, true
	case "B1SEC":
		return TimeBase_B1SEC, true
	}
	return 0, false
}

func TimeBaseKnows(value uint8) bool {
	for _, typeValue := range TimeBaseValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTimeBase(structType any) TimeBase {
	castFunc := func(typ any) TimeBase {
		if sTimeBase, ok := typ.(TimeBase); ok {
			return sTimeBase
		}
		return 0
	}
	return castFunc(structType)
}

func (m TimeBase) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m TimeBase) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TimeBaseParse(ctx context.Context, theBytes []byte) (TimeBase, error) {
	return TimeBaseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TimeBaseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TimeBase, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("TimeBase", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TimeBase")
	}
	if enum, ok := TimeBaseByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TimeBase")
		return TimeBase(val), nil
	} else {
		return enum, nil
	}
}

func (e TimeBase) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TimeBase) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("TimeBase", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e TimeBase) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TimeBase) PLC4XEnumName() string {
	switch e {
	case TimeBase_B10SEC:
		return "B10SEC"
	case TimeBase_B01SEC:
		return "B01SEC"
	case TimeBase_B1SEC:
		return "B1SEC"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e TimeBase) String() string {
	return e.PLC4XEnumName()
}