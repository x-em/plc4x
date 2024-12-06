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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION
    extends InformationObjectWithoutTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION;
  }

  // Properties.
  protected final NameOfFile nof;
  protected final NameOfSection nos;
  protected final SelectAndCallQualifier scq;

  public InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION(
      int address, NameOfFile nof, NameOfSection nos, SelectAndCallQualifier scq) {
    super(address);
    this.nof = nof;
    this.nos = nos;
    this.scq = scq;
  }

  public NameOfFile getNof() {
    return nof;
  }

  public NameOfSection getNos() {
    return nos;
  }

  public SelectAndCallQualifier getScq() {
    return scq;
  }

  @Override
  protected void serializeInformationObjectWithoutTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext(
        "InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION");

    // Simple Field (nof)
    writeSimpleField(
        "nof", nof, writeComplex(writeBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (nos)
    writeSimpleField(
        "nos", nos, writeComplex(writeBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (scq)
    writeSimpleField(
        "scq", scq, writeComplex(writeBuffer), WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext(
        "InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nof)
    lengthInBits += nof.getLengthInBits();

    // Simple field (nos)
    lengthInBits += nos.getLengthInBits();

    // Simple field (scq)
    lengthInBits += scq.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithoutTimeBuilder staticParseInformationObjectWithoutTimeBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext(
        "InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NameOfFile nof =
        readSimpleField(
            "nof",
            readComplex(() -> NameOfFile.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    NameOfSection nos =
        readSimpleField(
            "nos",
            readComplex(() -> NameOfSection.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    SelectAndCallQualifier scq =
        readSimpleField(
            "scq",
            readComplex(() -> SelectAndCallQualifier.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext(
        "InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION");
    // Create the instance
    return new InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTIONBuilderImpl(
        nof, nos, scq);
  }

  public static
  class InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTIONBuilderImpl
      implements InformationObjectWithoutTime.InformationObjectWithoutTimeBuilder {
    private final NameOfFile nof;
    private final NameOfSection nos;
    private final SelectAndCallQualifier scq;

    public
    InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTIONBuilderImpl(
        NameOfFile nof, NameOfSection nos, SelectAndCallQualifier scq) {
      this.nof = nof;
      this.nos = nos;
      this.scq = scq;
    }

    public InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION build(
        int address) {
      InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION
          informationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION =
              new InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION(
                  address, nof, nos, scq);
      return informationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o
        instanceof
        InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION)) {
      return false;
    }
    InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION that =
        (InformationObjectWithoutTime_CALL_DIRECTORY_SELECT_FILE_CALL_FILE_CALL_SECTION) o;
    return (getNof() == that.getNof())
        && (getNos() == that.getNos())
        && (getScq() == that.getScq())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNof(), getNos(), getScq());
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
