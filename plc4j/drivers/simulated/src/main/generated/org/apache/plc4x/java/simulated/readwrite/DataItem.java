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
package org.apache.plc4x.java.simulated.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import java.util.stream.Collectors;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class DataItem {

  private static final Logger LOGGER = LoggerFactory.getLogger(DataItem.class);

  public static PlcValue staticParse(ReadBuffer readBuffer, String dataType, Integer numberOfValues)
      throws ParseException {
    if (EvaluationHelper.equals(dataType, (String) "BOOL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // BOOL
      boolean value = readSimpleField("value", readBoolean(readBuffer));
      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(dataType, (String) "BOOL")) { // List
      List<Boolean> _value = readCountArrayField("value", readBoolean(readBuffer), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (boolean _item : _value) {
        value.add(new PlcBOOL(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "BYTE")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // BYTE
      short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));
      return new PlcBYTE(value);
    } else if (EvaluationHelper.equals(dataType, (String) "BYTE")) { // List
      List<Short> _value =
          readCountArrayField("value", readUnsignedShort(readBuffer, 8), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (short _item : _value) {
        value.add(new PlcUSINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "WORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // WORD
      int value = readSimpleField("value", readUnsignedInt(readBuffer, 16));
      return new PlcWORD(value);
    } else if (EvaluationHelper.equals(dataType, (String) "WORD")) { // List
      List<Integer> _value =
          readCountArrayField("value", readUnsignedInt(readBuffer, 16), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (int _item : _value) {
        value.add(new PlcUINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "DWORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // DWORD
      long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));
      return new PlcDWORD(value);
    } else if (EvaluationHelper.equals(dataType, (String) "DWORD")) { // List
      List<Long> _value =
          readCountArrayField("value", readUnsignedLong(readBuffer, 32), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (long _item : _value) {
        value.add(new PlcUDINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "LWORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LWORD
      BigInteger value = readSimpleField("value", readUnsignedBigInteger(readBuffer, 64));
      return new PlcLWORD(value);
    } else if (EvaluationHelper.equals(dataType, (String) "LWORD")) { // List
      List<BigInteger> _value =
          readCountArrayField("value", readUnsignedBigInteger(readBuffer, 64), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (BigInteger _item : _value) {
        value.add(new PlcULINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "SINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // SINT
      byte value = readSimpleField("value", readSignedByte(readBuffer, 8));
      return new PlcSINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "SINT")) { // List
      List<Byte> _value =
          readCountArrayField("value", readSignedByte(readBuffer, 8), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (byte _item : _value) {
        value.add(new PlcSINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "INT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // INT
      short value = readSimpleField("value", readSignedShort(readBuffer, 16));
      return new PlcINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "INT")) { // List
      List<Short> _value =
          readCountArrayField("value", readSignedShort(readBuffer, 16), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (short _item : _value) {
        value.add(new PlcINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "DINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // DINT
      int value = readSimpleField("value", readSignedInt(readBuffer, 32));
      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "DINT")) { // List
      List<Integer> _value =
          readCountArrayField("value", readSignedInt(readBuffer, 32), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (int _item : _value) {
        value.add(new PlcDINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "LINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LINT
      long value = readSimpleField("value", readSignedLong(readBuffer, 64));
      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "LINT")) { // List
      List<Long> _value =
          readCountArrayField("value", readSignedLong(readBuffer, 64), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (long _item : _value) {
        value.add(new PlcLINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "USINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // USINT
      short value = readSimpleField("value", readUnsignedShort(readBuffer, 8));
      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "USINT")) { // List
      List<Short> _value =
          readCountArrayField("value", readUnsignedShort(readBuffer, 8), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (short _item : _value) {
        value.add(new PlcUSINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "UINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // UINT
      int value = readSimpleField("value", readUnsignedInt(readBuffer, 16));
      return new PlcUINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "UINT")) { // List
      List<Integer> _value =
          readCountArrayField("value", readUnsignedInt(readBuffer, 16), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (int _item : _value) {
        value.add(new PlcUINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "UDINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // UDINT
      long value = readSimpleField("value", readUnsignedLong(readBuffer, 32));
      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "UDINT")) { // List
      List<Long> _value =
          readCountArrayField("value", readUnsignedLong(readBuffer, 32), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (long _item : _value) {
        value.add(new PlcUDINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "ULINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // ULINT
      BigInteger value = readSimpleField("value", readUnsignedBigInteger(readBuffer, 64));
      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataType, (String) "ULINT")) { // List
      List<BigInteger> _value =
          readCountArrayField("value", readUnsignedBigInteger(readBuffer, 64), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (BigInteger _item : _value) {
        value.add(new PlcULINT(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "REAL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // REAL
      float value = readSimpleField("value", readFloat(readBuffer, 32));
      return new PlcREAL(value);
    } else if (EvaluationHelper.equals(dataType, (String) "REAL")) { // List
      List<Float> _value = readCountArrayField("value", readFloat(readBuffer, 32), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (float _item : _value) {
        value.add(new PlcREAL(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "LREAL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LREAL
      double value = readSimpleField("value", readDouble(readBuffer, 64));
      return new PlcLREAL(value);
    } else if (EvaluationHelper.equals(dataType, (String) "LREAL")) { // List
      List<Double> _value =
          readCountArrayField("value", readDouble(readBuffer, 64), numberOfValues);
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (double _item : _value) {
        value.add(new PlcLREAL(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "CHAR")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // CHAR
      String value =
          readSimpleField("value", readString(readBuffer, 8), WithOption.WithEncoding("UTF-8"));
      return new PlcCHAR(value);
    } else if (EvaluationHelper.equals(dataType, (String) "CHAR")) { // List
      List<String> _value =
          readCountArrayField(
              "value", readString(readBuffer, 8), numberOfValues, WithOption.WithEncoding("UTF-8"));
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (String _item : _value) {
        value.add(new PlcSTRING(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "WCHAR")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // WCHAR
      String value =
          readSimpleField("value", readString(readBuffer, 16), WithOption.WithEncoding("UTF-16"));
      return new PlcWCHAR(value);
    } else if (EvaluationHelper.equals(dataType, (String) "WCHAR")) { // List
      List<String> _value =
          readCountArrayField(
              "value",
              readString(readBuffer, 16),
              numberOfValues,
              WithOption.WithEncoding("UTF-16"));
      List<PlcValue> value = new ArrayList<>(_value.size());
      for (String _item : _value) {
        value.add(new PlcSTRING(_item));
      }
      return new PlcList(value);
    } else if (EvaluationHelper.equals(dataType, (String) "STRING")) { // STRING
      String value =
          readSimpleField("value", readString(readBuffer, 255), WithOption.WithEncoding("UTF-8"));
      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataType, (String) "WSTRING")) { // STRING
      String value =
          readSimpleField("value", readString(readBuffer, 255), WithOption.WithEncoding("UTF-16"));
      return new PlcSTRING(value);
    }
    return null;
  }

  public static int getLengthInBytes(PlcValue _value, String dataType, Integer numberOfValues) {
    return (int) Math.ceil((float) getLengthInBits(_value, dataType, numberOfValues) / 8.0);
  }

  public static int getLengthInBits(PlcValue _value, String dataType, Integer numberOfValues) {
    int lengthInBits = 0;
    if (EvaluationHelper.equals(dataType, (String) "BOOL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // BOOL
      // Simple field (value)
      lengthInBits += 1;
    } else if (EvaluationHelper.equals(dataType, (String) "BOOL")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 1 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "BYTE")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // BYTE
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataType, (String) "BYTE")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 8 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "WORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // WORD
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataType, (String) "WORD")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 16 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "DWORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // DWORD
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataType, (String) "DWORD")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 32 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "LWORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LWORD
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataType, (String) "LWORD")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 64 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "SINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // SINT
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataType, (String) "SINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 8 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "INT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // INT
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataType, (String) "INT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 16 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "DINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // DINT
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataType, (String) "DINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 32 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "LINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LINT
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataType, (String) "LINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 64 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "USINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // USINT
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataType, (String) "USINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 8 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "UINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // UINT
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataType, (String) "UINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 16 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "UDINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // UDINT
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataType, (String) "UDINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 32 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "ULINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // ULINT
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataType, (String) "ULINT")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 64 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "REAL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // REAL
      // Simple field (value)
      lengthInBits += 32;
    } else if (EvaluationHelper.equals(dataType, (String) "REAL")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 32 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "LREAL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LREAL
      // Simple field (value)
      lengthInBits += 64;
    } else if (EvaluationHelper.equals(dataType, (String) "LREAL")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 64 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "CHAR")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // CHAR
      // Simple field (value)
      lengthInBits += 8;
    } else if (EvaluationHelper.equals(dataType, (String) "CHAR")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 8 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "WCHAR")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // WCHAR
      // Simple field (value)
      lengthInBits += 16;
    } else if (EvaluationHelper.equals(dataType, (String) "WCHAR")) { // List
      // Array field
      if (_value != null) {
        lengthInBits += 16 * _value.getList().size();
      }
    } else if (EvaluationHelper.equals(dataType, (String) "STRING")) { // STRING
      // Simple field (value)
      lengthInBits += 255;
    } else if (EvaluationHelper.equals(dataType, (String) "WSTRING")) { // STRING
      // Simple field (value)
      lengthInBits += 255;
    }

    return lengthInBits;
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, String dataType, Integer numberOfValues)
      throws SerializationException {
    staticSerialize(writeBuffer, _value, dataType, numberOfValues, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer,
      PlcValue _value,
      String dataType,
      Integer numberOfValues,
      ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(dataType, (String) "BOOL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // BOOL
      // Simple Field (value)
      writeSimpleField("value", (boolean) _value.getBoolean(), writeBoolean(writeBuffer));
    } else if (EvaluationHelper.equals(dataType, (String) "BOOL")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getBoolean).collect(Collectors.toList()),
          writeBoolean(writeBuffer));
    } else if (EvaluationHelper.equals(dataType, (String) "BYTE")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // BYTE
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataType, (String) "BYTE")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getShort).collect(Collectors.toList()),
          writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataType, (String) "WORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // WORD
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataType, (String) "WORD")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getInteger).collect(Collectors.toList()),
          writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataType, (String) "DWORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // DWORD
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "DWORD")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getLong).collect(Collectors.toList()),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "LWORD")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LWORD
      // Simple Field (value)
      writeSimpleField(
          "value", (BigInteger) _value.getBigInteger(), writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "LWORD")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getBigInteger).collect(Collectors.toList()),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "SINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // SINT
      // Simple Field (value)
      writeSimpleField("value", (byte) _value.getByte(), writeSignedByte(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataType, (String) "SINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getByte).collect(Collectors.toList()),
          writeSignedByte(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataType, (String) "INT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // INT
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeSignedShort(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataType, (String) "INT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getShort).collect(Collectors.toList()),
          writeSignedShort(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataType, (String) "DINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // DINT
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeSignedInt(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "DINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getInteger).collect(Collectors.toList()),
          writeSignedInt(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "LINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LINT
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeSignedLong(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "LINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getLong).collect(Collectors.toList()),
          writeSignedLong(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "USINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // USINT
      // Simple Field (value)
      writeSimpleField("value", (short) _value.getShort(), writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataType, (String) "USINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getShort).collect(Collectors.toList()),
          writeUnsignedShort(writeBuffer, 8));
    } else if (EvaluationHelper.equals(dataType, (String) "UINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // UINT
      // Simple Field (value)
      writeSimpleField("value", (int) _value.getInteger(), writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataType, (String) "UINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getInteger).collect(Collectors.toList()),
          writeUnsignedInt(writeBuffer, 16));
    } else if (EvaluationHelper.equals(dataType, (String) "UDINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // UDINT
      // Simple Field (value)
      writeSimpleField("value", (long) _value.getLong(), writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "UDINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getLong).collect(Collectors.toList()),
          writeUnsignedLong(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "ULINT")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // ULINT
      // Simple Field (value)
      writeSimpleField(
          "value", (BigInteger) _value.getBigInteger(), writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "ULINT")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getBigInteger).collect(Collectors.toList()),
          writeUnsignedBigInteger(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "REAL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // REAL
      // Simple Field (value)
      writeSimpleField("value", (float) _value.getFloat(), writeFloat(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "REAL")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getFloat).collect(Collectors.toList()),
          writeFloat(writeBuffer, 32));
    } else if (EvaluationHelper.equals(dataType, (String) "LREAL")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // LREAL
      // Simple Field (value)
      writeSimpleField("value", (double) _value.getDouble(), writeDouble(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "LREAL")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getDouble).collect(Collectors.toList()),
          writeDouble(writeBuffer, 64));
    } else if (EvaluationHelper.equals(dataType, (String) "CHAR")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // CHAR
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 8),
          WithOption.WithEncoding("UTF-8"));
    } else if (EvaluationHelper.equals(dataType, (String) "CHAR")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getString).collect(Collectors.toList()),
          writeString(writeBuffer, 8),
          WithOption.WithEncoding("UTF-8"));
    } else if (EvaluationHelper.equals(dataType, (String) "WCHAR")
        && EvaluationHelper.equals(numberOfValues, (int) 1)) { // WCHAR
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 16),
          WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(dataType, (String) "WCHAR")) { // List
      // Array Field (value)
      writeSimpleTypeArrayField(
          "value",
          _value.getList().stream().map(PlcValue::getString).collect(Collectors.toList()),
          writeString(writeBuffer, 16),
          WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(dataType, (String) "STRING")) { // STRING
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 255),
          WithOption.WithEncoding("UTF-8"));
    } else if (EvaluationHelper.equals(dataType, (String) "WSTRING")) { // STRING
      // Simple Field (value)
      writeSimpleField(
          "value",
          (String) _value.getString(),
          writeString(writeBuffer, 255),
          WithOption.WithEncoding("UTF-16"));
    }
  }
}