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
package org.apache.plc4x.java.cbus.readwrite;

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

public class ErrorReportingSystemCategory implements Message {

  // Properties.
  protected final ErrorReportingSystemCategoryClass systemCategoryClass;
  protected final ErrorReportingSystemCategoryType systemCategoryType;
  protected final ErrorReportingSystemCategoryVariant systemCategoryVariant;

  public ErrorReportingSystemCategory(
      ErrorReportingSystemCategoryClass systemCategoryClass,
      ErrorReportingSystemCategoryType systemCategoryType,
      ErrorReportingSystemCategoryVariant systemCategoryVariant) {
    super();
    this.systemCategoryClass = systemCategoryClass;
    this.systemCategoryType = systemCategoryType;
    this.systemCategoryVariant = systemCategoryVariant;
  }

  public ErrorReportingSystemCategoryClass getSystemCategoryClass() {
    return systemCategoryClass;
  }

  public ErrorReportingSystemCategoryType getSystemCategoryType() {
    return systemCategoryType;
  }

  public ErrorReportingSystemCategoryVariant getSystemCategoryVariant() {
    return systemCategoryVariant;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ErrorReportingSystemCategory");

    // Simple Field (systemCategoryClass)
    writeSimpleEnumField(
        "systemCategoryClass",
        "ErrorReportingSystemCategoryClass",
        systemCategoryClass,
        writeEnum(
            ErrorReportingSystemCategoryClass::getValue,
            ErrorReportingSystemCategoryClass::name,
            writeUnsignedByte(writeBuffer, 4)));

    // Simple Field (systemCategoryType)
    writeSimpleField("systemCategoryType", systemCategoryType, writeComplex(writeBuffer));

    // Simple Field (systemCategoryVariant)
    writeSimpleEnumField(
        "systemCategoryVariant",
        "ErrorReportingSystemCategoryVariant",
        systemCategoryVariant,
        writeEnum(
            ErrorReportingSystemCategoryVariant::getValue,
            ErrorReportingSystemCategoryVariant::name,
            writeUnsignedByte(writeBuffer, 2)));

    writeBuffer.popContext("ErrorReportingSystemCategory");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ErrorReportingSystemCategory _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (systemCategoryClass)
    lengthInBits += 4;

    // Simple field (systemCategoryType)
    lengthInBits += systemCategoryType.getLengthInBits();

    // Simple field (systemCategoryVariant)
    lengthInBits += 2;

    return lengthInBits;
  }

  public static ErrorReportingSystemCategory staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("ErrorReportingSystemCategory");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ErrorReportingSystemCategoryClass systemCategoryClass =
        readEnumField(
            "systemCategoryClass",
            "ErrorReportingSystemCategoryClass",
            readEnum(
                ErrorReportingSystemCategoryClass::enumForValue, readUnsignedByte(readBuffer, 4)));

    ErrorReportingSystemCategoryType systemCategoryType =
        readSimpleField(
            "systemCategoryType",
            readComplex(
                () ->
                    ErrorReportingSystemCategoryType.staticParse(
                        readBuffer, (ErrorReportingSystemCategoryClass) (systemCategoryClass)),
                readBuffer));

    ErrorReportingSystemCategoryVariant systemCategoryVariant =
        readEnumField(
            "systemCategoryVariant",
            "ErrorReportingSystemCategoryVariant",
            readEnum(
                ErrorReportingSystemCategoryVariant::enumForValue,
                readUnsignedByte(readBuffer, 2)));

    readBuffer.closeContext("ErrorReportingSystemCategory");
    // Create the instance
    ErrorReportingSystemCategory _errorReportingSystemCategory;
    _errorReportingSystemCategory =
        new ErrorReportingSystemCategory(
            systemCategoryClass, systemCategoryType, systemCategoryVariant);
    return _errorReportingSystemCategory;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ErrorReportingSystemCategory)) {
      return false;
    }
    ErrorReportingSystemCategory that = (ErrorReportingSystemCategory) o;
    return (getSystemCategoryClass() == that.getSystemCategoryClass())
        && (getSystemCategoryType() == that.getSystemCategoryType())
        && (getSystemCategoryVariant() == that.getSystemCategoryVariant())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getSystemCategoryClass(), getSystemCategoryType(), getSystemCategoryVariant());
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