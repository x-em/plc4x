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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetLogDataLogDataEntryNullValue extends BACnetLogDataLogDataEntry
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagNull nullValue;

  public BACnetLogDataLogDataEntryNullValue(
      BACnetTagHeader peekedTagHeader, BACnetContextTagNull nullValue) {
    super(peekedTagHeader);
    this.nullValue = nullValue;
  }

  public BACnetContextTagNull getNullValue() {
    return nullValue;
  }

  @Override
  protected void serializeBACnetLogDataLogDataEntryChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetLogDataLogDataEntryNullValue");

    // Simple Field (nullValue)
    writeSimpleField("nullValue", nullValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetLogDataLogDataEntryNullValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetLogDataLogDataEntryNullValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nullValue)
    lengthInBits += nullValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetLogDataLogDataEntryBuilder staticParseBACnetLogDataLogDataEntryBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetLogDataLogDataEntryNullValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagNull nullValue =
        readSimpleField(
            "nullValue",
            readComplex(
                () ->
                    (BACnetContextTagNull)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (6), (BACnetDataType) (BACnetDataType.NULL)),
                readBuffer));

    readBuffer.closeContext("BACnetLogDataLogDataEntryNullValue");
    // Create the instance
    return new BACnetLogDataLogDataEntryNullValueBuilderImpl(nullValue);
  }

  public static class BACnetLogDataLogDataEntryNullValueBuilderImpl
      implements BACnetLogDataLogDataEntry.BACnetLogDataLogDataEntryBuilder {
    private final BACnetContextTagNull nullValue;

    public BACnetLogDataLogDataEntryNullValueBuilderImpl(BACnetContextTagNull nullValue) {
      this.nullValue = nullValue;
    }

    public BACnetLogDataLogDataEntryNullValue build(BACnetTagHeader peekedTagHeader) {
      BACnetLogDataLogDataEntryNullValue bACnetLogDataLogDataEntryNullValue =
          new BACnetLogDataLogDataEntryNullValue(peekedTagHeader, nullValue);
      return bACnetLogDataLogDataEntryNullValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetLogDataLogDataEntryNullValue)) {
      return false;
    }
    BACnetLogDataLogDataEntryNullValue that = (BACnetLogDataLogDataEntryNullValue) o;
    return (getNullValue() == that.getNullValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNullValue());
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