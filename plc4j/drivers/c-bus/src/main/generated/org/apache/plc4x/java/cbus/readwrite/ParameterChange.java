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

public class ParameterChange implements Message {

  // Constant values.
  public static final Byte SPECIALCHAR1 = 0x3D;
  public static final Byte SPECIALCHAR2 = 0x3D;

  public ParameterChange() {
    super();
  }

  public byte getSpecialChar1() {
    return SPECIALCHAR1;
  }

  public byte getSpecialChar2() {
    return SPECIALCHAR2;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ParameterChange");

    // Const Field (specialChar1)
    writeConstField("specialChar1", SPECIALCHAR1, writeByte(writeBuffer, 8));

    // Const Field (specialChar2)
    writeConstField("specialChar2", SPECIALCHAR2, writeByte(writeBuffer, 8));

    writeBuffer.popContext("ParameterChange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    ParameterChange _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (specialChar1)
    lengthInBits += 8;

    // Const Field (specialChar2)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ParameterChange staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ParameterChange");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte specialChar1 =
        readConstField("specialChar1", readByte(readBuffer, 8), ParameterChange.SPECIALCHAR1);

    byte specialChar2 =
        readConstField("specialChar2", readByte(readBuffer, 8), ParameterChange.SPECIALCHAR2);

    readBuffer.closeContext("ParameterChange");
    // Create the instance
    ParameterChange _parameterChange;
    _parameterChange = new ParameterChange();
    return _parameterChange;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ParameterChange)) {
      return false;
    }
    ParameterChange that = (ParameterChange) o;
    return true;
  }

  @Override
  public int hashCode() {
    return Objects.hash();
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