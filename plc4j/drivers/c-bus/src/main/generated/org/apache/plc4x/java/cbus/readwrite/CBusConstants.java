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

public class CBusConstants implements Message {

  // Constant values.
  public static final Integer CBUSTCPDEFAULTPORT = 10001;

  public CBusConstants() {
    super();
  }

  public int getCbusTcpDefaultPort() {
    return CBUSTCPDEFAULTPORT;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("CBusConstants");

    // Const Field (cbusTcpDefaultPort)
    writeConstField("cbusTcpDefaultPort", CBUSTCPDEFAULTPORT, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("CBusConstants");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CBusConstants _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (cbusTcpDefaultPort)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static CBusConstants staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static CBusConstants staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CBusConstants");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int cbusTcpDefaultPort =
        readConstField(
            "cbusTcpDefaultPort",
            readUnsignedInt(readBuffer, 16),
            CBusConstants.CBUSTCPDEFAULTPORT);

    readBuffer.closeContext("CBusConstants");
    // Create the instance
    CBusConstants _cBusConstants;
    _cBusConstants = new CBusConstants();
    return _cBusConstants;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CBusConstants)) {
      return false;
    }
    CBusConstants that = (CBusConstants) o;
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
