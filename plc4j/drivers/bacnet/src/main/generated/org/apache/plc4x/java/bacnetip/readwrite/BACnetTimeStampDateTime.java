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

public class BACnetTimeStampDateTime extends BACnetTimeStamp implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetDateTimeEnclosed dateTimeValue;

  public BACnetTimeStampDateTime(
      BACnetTagHeader peekedTagHeader, BACnetDateTimeEnclosed dateTimeValue) {
    super(peekedTagHeader);
    this.dateTimeValue = dateTimeValue;
  }

  public BACnetDateTimeEnclosed getDateTimeValue() {
    return dateTimeValue;
  }

  @Override
  protected void serializeBACnetTimeStampChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetTimeStampDateTime");

    // Simple Field (dateTimeValue)
    writeSimpleField("dateTimeValue", dateTimeValue, writeComplex(writeBuffer));

    writeBuffer.popContext("BACnetTimeStampDateTime");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetTimeStampDateTime _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (dateTimeValue)
    lengthInBits += dateTimeValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetTimeStampBuilder staticParseBACnetTimeStampBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("BACnetTimeStampDateTime");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetDateTimeEnclosed dateTimeValue =
        readSimpleField(
            "dateTimeValue",
            readComplex(
                () -> BACnetDateTimeEnclosed.staticParse(readBuffer, (short) (2)), readBuffer));

    readBuffer.closeContext("BACnetTimeStampDateTime");
    // Create the instance
    return new BACnetTimeStampDateTimeBuilderImpl(dateTimeValue);
  }

  public static class BACnetTimeStampDateTimeBuilderImpl
      implements BACnetTimeStamp.BACnetTimeStampBuilder {
    private final BACnetDateTimeEnclosed dateTimeValue;

    public BACnetTimeStampDateTimeBuilderImpl(BACnetDateTimeEnclosed dateTimeValue) {
      this.dateTimeValue = dateTimeValue;
    }

    public BACnetTimeStampDateTime build(BACnetTagHeader peekedTagHeader) {
      BACnetTimeStampDateTime bACnetTimeStampDateTime =
          new BACnetTimeStampDateTime(peekedTagHeader, dateTimeValue);
      return bACnetTimeStampDateTime;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetTimeStampDateTime)) {
      return false;
    }
    BACnetTimeStampDateTime that = (BACnetTimeStampDateTime) o;
    return (getDateTimeValue() == that.getDateTimeValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDateTimeValue());
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