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

public class IdentifyReplyCommandType extends IdentifyReplyCommand implements Message {

  // Accessors for discriminator values.
  public Attribute getAttribute() {
    return Attribute.Type;
  }

  // Properties.
  protected final String unitType;

  // Arguments.
  protected final Byte numBytes;

  public IdentifyReplyCommandType(String unitType, Byte numBytes) {
    super(numBytes);
    this.unitType = unitType;
    this.numBytes = numBytes;
  }

  public String getUnitType() {
    return unitType;
  }

  @Override
  protected void serializeIdentifyReplyCommandChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("IdentifyReplyCommandType");

    // Simple Field (unitType)
    writeSimpleField("unitType", unitType, writeString(writeBuffer, 64));

    writeBuffer.popContext("IdentifyReplyCommandType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    IdentifyReplyCommandType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (unitType)
    lengthInBits += 64;

    return lengthInBits;
  }

  public static IdentifyReplyCommandBuilder staticParseIdentifyReplyCommandBuilder(
      ReadBuffer readBuffer, Attribute attribute, Byte numBytes) throws ParseException {
    readBuffer.pullContext("IdentifyReplyCommandType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    String unitType = readSimpleField("unitType", readString(readBuffer, 64));

    readBuffer.closeContext("IdentifyReplyCommandType");
    // Create the instance
    return new IdentifyReplyCommandTypeBuilderImpl(unitType, numBytes);
  }

  public static class IdentifyReplyCommandTypeBuilderImpl
      implements IdentifyReplyCommand.IdentifyReplyCommandBuilder {
    private final String unitType;
    private final Byte numBytes;

    public IdentifyReplyCommandTypeBuilderImpl(String unitType, Byte numBytes) {
      this.unitType = unitType;
      this.numBytes = numBytes;
    }

    public IdentifyReplyCommandType build(Byte numBytes) {

      IdentifyReplyCommandType identifyReplyCommandType =
          new IdentifyReplyCommandType(unitType, numBytes);
      return identifyReplyCommandType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof IdentifyReplyCommandType)) {
      return false;
    }
    IdentifyReplyCommandType that = (IdentifyReplyCommandType) o;
    return (getUnitType() == that.getUnitType()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getUnitType());
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