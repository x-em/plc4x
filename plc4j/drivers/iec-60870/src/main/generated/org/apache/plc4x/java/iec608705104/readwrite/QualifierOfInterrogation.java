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

public class QualifierOfInterrogation implements Message {

  // Properties.
  protected final short qualifierOfCommand;

  public QualifierOfInterrogation(short qualifierOfCommand) {
    super();
    this.qualifierOfCommand = qualifierOfCommand;
  }

  public short getQualifierOfCommand() {
    return qualifierOfCommand;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("QualifierOfInterrogation");

    // Simple Field (qualifierOfCommand)
    writeSimpleField(
        "qualifierOfCommand",
        qualifierOfCommand,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("QualifierOfInterrogation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    QualifierOfInterrogation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (qualifierOfCommand)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static QualifierOfInterrogation staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("QualifierOfInterrogation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short qualifierOfCommand =
        readSimpleField(
            "qualifierOfCommand",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("QualifierOfInterrogation");
    // Create the instance
    QualifierOfInterrogation _qualifierOfInterrogation;
    _qualifierOfInterrogation = new QualifierOfInterrogation(qualifierOfCommand);
    return _qualifierOfInterrogation;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QualifierOfInterrogation)) {
      return false;
    }
    QualifierOfInterrogation that = (QualifierOfInterrogation) o;
    return (getQualifierOfCommand() == that.getQualifierOfCommand()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getQualifierOfCommand());
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
