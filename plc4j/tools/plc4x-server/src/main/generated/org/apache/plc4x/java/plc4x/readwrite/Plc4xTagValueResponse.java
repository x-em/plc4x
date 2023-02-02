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
package org.apache.plc4x.java.plc4x.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class Plc4xTagValueResponse implements Message {

  // Properties.
  protected final Plc4xTag tag;
  protected final Plc4xResponseCode responseCode;
  protected final Plc4xValueType valueType;
  protected final PlcValue value;

  public Plc4xTagValueResponse(
      Plc4xTag tag, Plc4xResponseCode responseCode, Plc4xValueType valueType, PlcValue value) {
    super();
    this.tag = tag;
    this.responseCode = responseCode;
    this.valueType = valueType;
    this.value = value;
  }

  public Plc4xTag getTag() {
    return tag;
  }

  public Plc4xResponseCode getResponseCode() {
    return responseCode;
  }

  public Plc4xValueType getValueType() {
    return valueType;
  }

  public PlcValue getValue() {
    return value;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("Plc4xTagValueResponse");

    // Simple Field (tag)
    writeSimpleField("tag", tag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (responseCode)
    writeSimpleEnumField(
        "responseCode",
        "Plc4xResponseCode",
        responseCode,
        new DataWriterEnumDefault<>(
            Plc4xResponseCode::getValue,
            Plc4xResponseCode::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (valueType)
    writeSimpleEnumField(
        "valueType",
        "Plc4xValueType",
        valueType,
        new DataWriterEnumDefault<>(
            Plc4xValueType::getValue, Plc4xValueType::name, writeUnsignedShort(writeBuffer, 8)));

    // Optional Field (value) (Can be skipped, if the value is null)
    writeOptionalField(
        "value",
        value,
        new DataWriterDataIoDefault(
            writeBuffer, (wb, val) -> Plc4xValue.staticSerialize(wb, val, valueType)),
        (getValueType()) != (Plc4xValueType.NULL));

    writeBuffer.popContext("Plc4xTagValueResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Plc4xTagValueResponse _value = this;

    // Simple field (tag)
    lengthInBits += tag.getLengthInBits();

    // Simple field (responseCode)
    lengthInBits += 8;

    // Simple field (valueType)
    lengthInBits += 8;

    // Optional Field (value)
    if (value != null) {
      lengthInBits += Plc4xValue.getLengthInBits(value, valueType);
    }

    return lengthInBits;
  }

  public static Plc4xTagValueResponse staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static Plc4xTagValueResponse staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Plc4xTagValueResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Plc4xTag tag =
        readSimpleField(
            "tag",
            new DataReaderComplexDefault<>(() -> Plc4xTag.staticParse(readBuffer), readBuffer));

    Plc4xResponseCode responseCode =
        readEnumField(
            "responseCode",
            "Plc4xResponseCode",
            new DataReaderEnumDefault<>(
                Plc4xResponseCode::enumForValue, readUnsignedShort(readBuffer, 8)));

    Plc4xValueType valueType =
        readEnumField(
            "valueType",
            "Plc4xValueType",
            new DataReaderEnumDefault<>(
                Plc4xValueType::enumForValue, readUnsignedShort(readBuffer, 8)));

    PlcValue value =
        readOptionalField(
            "value",
            new DataReaderComplexDefault<>(
                () -> Plc4xValue.staticParse(readBuffer, (Plc4xValueType) (valueType)), readBuffer),
            (valueType) != (Plc4xValueType.NULL));

    readBuffer.closeContext("Plc4xTagValueResponse");
    // Create the instance
    Plc4xTagValueResponse _plc4xTagValueResponse;
    _plc4xTagValueResponse = new Plc4xTagValueResponse(tag, responseCode, valueType, value);
    return _plc4xTagValueResponse;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Plc4xTagValueResponse)) {
      return false;
    }
    Plc4xTagValueResponse that = (Plc4xTagValueResponse) o;
    return (getTag() == that.getTag())
        && (getResponseCode() == that.getResponseCode())
        && (getValueType() == that.getValueType())
        && (getValue() == that.getValue())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getTag(), getResponseCode(), getValueType(), getValue());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
