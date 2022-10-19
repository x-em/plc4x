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
	api "github.com/apache/plc4x/plc4go/pkg/api/values"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/apache/plc4x/plc4go/spi/values"
	"github.com/pkg/errors"
	"time"
)

// Code generated by code-generation. DO NOT EDIT.

func DataItemParse(readBuffer utils.ReadBuffer, dataProtocolId string, stringLength int32) (api.PlcValue, error) {
	readBuffer.PullContext("DataItem")
	switch {
	case dataProtocolId == "IEC61131_BOOL": // BOOL
		// Reserved Field (Just skip the bytes)
		if _, _err := readBuffer.ReadUint8("reserved", 7); _err != nil {
			return nil, errors.Wrap(_err, "Error parsing reserved field")
		}

		// Simple Field (value)
		value, _valueErr := readBuffer.ReadBit("value")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBOOL(value), nil
	case dataProtocolId == "IEC61131_BYTE": // BYTE
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcBYTE(value), nil
	case dataProtocolId == "IEC61131_WORD": // WORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcWORD(value), nil
	case dataProtocolId == "IEC61131_DWORD": // DWORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDWORD(value), nil
	case dataProtocolId == "IEC61131_LWORD": // LWORD
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLWORD(value), nil
	case dataProtocolId == "IEC61131_SINT": // SINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSINT(value), nil
	case dataProtocolId == "IEC61131_USINT": // USINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint8("value", 8)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUSINT(value), nil
	case dataProtocolId == "IEC61131_INT": // INT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcINT(value), nil
	case dataProtocolId == "IEC61131_UINT": // UINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUINT(value), nil
	case dataProtocolId == "IEC61131_DINT": // DINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDINT(value), nil
	case dataProtocolId == "IEC61131_UDINT": // UDINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcUDINT(value), nil
	case dataProtocolId == "IEC61131_LINT": // LINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadInt64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLINT(value), nil
	case dataProtocolId == "IEC61131_ULINT": // ULINT
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcULINT(value), nil
	case dataProtocolId == "IEC61131_REAL": // REAL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcREAL(value), nil
	case dataProtocolId == "IEC61131_LREAL": // LREAL
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadFloat64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLREAL(value), nil
	case dataProtocolId == "IEC61131_CHAR": // CHAR
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32(8), "UTF-8")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcCHAR(value), nil
	case dataProtocolId == "IEC61131_WCHAR": // CHAR
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadString("value", uint32(16), "UTF-16")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcCHAR(value), nil
	case dataProtocolId == "IEC61131_STRING": // STRING
		// Manual Field (value)
		value, _valueErr := ParseS7String(readBuffer, stringLength, "UTF-8")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataProtocolId == "IEC61131_WSTRING": // STRING
		// Manual Field (value)
		value, _valueErr := ParseS7String(readBuffer, stringLength, "UTF-16")
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcSTRING(value), nil
	case dataProtocolId == "IEC61131_TIME": // TIME
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIME(value), nil
	case dataProtocolId == "IEC61131_LTIME": // LTIME
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint64("value", 64)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcLTIME(value), nil
	case dataProtocolId == "IEC61131_DATE": // DATE
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint16("value", 16)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDATE(value), nil
	case dataProtocolId == "IEC61131_TIME_OF_DAY": // TIME_OF_DAY
		// Simple Field (value)
		value, _valueErr := readBuffer.ReadUint32("value", 32)
		if _valueErr != nil {
			return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
		}
		readBuffer.CloseContext("DataItem")
		return values.NewPlcTIME_OF_DAY(value), nil
	case dataProtocolId == "IEC61131_DATE_AND_TIME": // DATE_AND_TIME
		// Simple Field (year)
		year, _yearErr := readBuffer.ReadUint16("year", 16)
		if _yearErr != nil {
			return nil, errors.Wrap(_yearErr, "Error parsing 'year' field")
		}

		// Simple Field (month)
		month, _monthErr := readBuffer.ReadUint8("month", 8)
		if _monthErr != nil {
			return nil, errors.Wrap(_monthErr, "Error parsing 'month' field")
		}

		// Simple Field (day)
		day, _dayErr := readBuffer.ReadUint8("day", 8)
		if _dayErr != nil {
			return nil, errors.Wrap(_dayErr, "Error parsing 'day' field")
		}

		// Simple Field (dayOfWeek)
		_, _dayOfWeekErr := readBuffer.ReadUint8("dayOfWeek", 8)
		if _dayOfWeekErr != nil {
			return nil, errors.Wrap(_dayOfWeekErr, "Error parsing 'dayOfWeek' field")
		}

		// Simple Field (hour)
		hour, _hourErr := readBuffer.ReadUint8("hour", 8)
		if _hourErr != nil {
			return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
		}

		// Simple Field (minutes)
		minutes, _minutesErr := readBuffer.ReadUint8("minutes", 8)
		if _minutesErr != nil {
			return nil, errors.Wrap(_minutesErr, "Error parsing 'minutes' field")
		}

		// Simple Field (seconds)
		seconds, _secondsErr := readBuffer.ReadUint8("seconds", 8)
		if _secondsErr != nil {
			return nil, errors.Wrap(_secondsErr, "Error parsing 'seconds' field")
		}

		// Simple Field (nanos)
		_, _nanosErr := readBuffer.ReadUint32("nanos", 32)
		if _nanosErr != nil {
			return nil, errors.Wrap(_nanosErr, "Error parsing 'nanos' field")
		}
		value := time.Date(int(year), time.Month(month), int(day), int(hour), int(minutes), int(seconds), 0, nil)
		readBuffer.CloseContext("DataItem")
		return values.NewPlcDATE_AND_TIME(value), nil
	}
	// TODO: add more info which type it is actually
	return nil, errors.New("unsupported type")
}

func DataItemSerialize(writeBuffer utils.WriteBuffer, value api.PlcValue, dataProtocolId string, stringLength int32) error {
	m := struct {
		DataProtocolId string
		StringLength   int32
	}{
		DataProtocolId: dataProtocolId,
		StringLength:   stringLength,
	}
	_ = m
	writeBuffer.PushContext("DataItem")
	switch {
	case dataProtocolId == "IEC61131_BOOL": // BOOL
		// Reserved Field (Just skip the bytes)
		if _err := writeBuffer.WriteUint8("reserved", 7, uint8(0x00)); _err != nil {
			return errors.Wrap(_err, "Error serializing reserved field")
		}

		// Simple Field (value)
		if _err := writeBuffer.WriteBit("value", value.GetBool()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_BYTE": // BYTE
		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_WORD": // WORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_DWORD": // DWORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LWORD": // LWORD
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_SINT": // SINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt8("value", 8, value.GetInt8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_USINT": // USINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint8("value", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_INT": // INT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt16("value", 16, value.GetInt16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_UINT": // UINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_DINT": // DINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt32("value", 32, value.GetInt32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_UDINT": // UDINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LINT": // LINT
		// Simple Field (value)
		if _err := writeBuffer.WriteInt64("value", 64, value.GetInt64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_ULINT": // ULINT
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_REAL": // REAL
		// Simple Field (value)
		if _err := writeBuffer.WriteFloat32("value", 32, value.GetFloat32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LREAL": // LREAL
		// Simple Field (value)
		if _err := writeBuffer.WriteFloat64("value", 64, value.GetFloat64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_CHAR": // CHAR
		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint32(8), "UTF-8", value.GetString()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_WCHAR": // CHAR
		// Simple Field (value)
		if _err := writeBuffer.WriteString("value", uint32(16), "UTF-16", value.GetString()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_STRING": // STRING
		// Manual Field (value)
		_valueErr := SerializeS7String(writeBuffer, value, stringLength, "UTF-8")
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_WSTRING": // STRING
		// Manual Field (value)
		_valueErr := SerializeS7String(writeBuffer, value, stringLength, "UTF-16")
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_TIME": // TIME
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_LTIME": // LTIME
		// Simple Field (value)
		if _err := writeBuffer.WriteUint64("value", 64, value.GetUint64()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_DATE": // DATE
		// Simple Field (value)
		if _err := writeBuffer.WriteUint16("value", 16, value.GetUint16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_TIME_OF_DAY": // TIME_OF_DAY
		// Simple Field (value)
		if _err := writeBuffer.WriteUint32("value", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'value' field")
		}
	case dataProtocolId == "IEC61131_DATE_AND_TIME": // DATE_AND_TIME
		// Simple Field (year)
		if _err := writeBuffer.WriteUint16("year", 16, value.GetUint16()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'year' field")
		}

		// Simple Field (month)
		if _err := writeBuffer.WriteUint8("month", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'month' field")
		}

		// Simple Field (day)
		if _err := writeBuffer.WriteUint8("day", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'day' field")
		}

		// Simple Field (dayOfWeek)
		if _err := writeBuffer.WriteUint8("dayOfWeek", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'dayOfWeek' field")
		}

		// Simple Field (hour)
		if _err := writeBuffer.WriteUint8("hour", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'hour' field")
		}

		// Simple Field (minutes)
		if _err := writeBuffer.WriteUint8("minutes", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'minutes' field")
		}

		// Simple Field (seconds)
		if _err := writeBuffer.WriteUint8("seconds", 8, value.GetUint8()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'seconds' field")
		}

		// Simple Field (nanos)
		if _err := writeBuffer.WriteUint32("nanos", 32, value.GetUint32()); _err != nil {
			return errors.Wrap(_err, "Error serializing 'nanos' field")
		}
	default:
		// TODO: add more info which type it is actually
		return errors.New("unsupported type")
	}
	writeBuffer.PopContext("DataItem")
	return nil
}
