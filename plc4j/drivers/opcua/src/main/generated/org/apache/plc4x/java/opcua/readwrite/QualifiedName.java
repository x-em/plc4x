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
package org.apache.plc4x.java.opcua.readwrite;

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

public class QualifiedName implements Message {

  // Properties.
  protected final int namespaceIndex;
  protected final PascalString name;

  public QualifiedName(int namespaceIndex, PascalString name) {
    super();
    this.namespaceIndex = namespaceIndex;
    this.name = name;
  }

  public int getNamespaceIndex() {
    return namespaceIndex;
  }

  public PascalString getName() {
    return name;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("QualifiedName");

    // Simple Field (namespaceIndex)
    writeSimpleField("namespaceIndex", namespaceIndex, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (name)
    writeSimpleField("name", name, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("QualifiedName");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    QualifiedName _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (namespaceIndex)
    lengthInBits += 16;

    // Simple field (name)
    lengthInBits += name.getLengthInBits();

    return lengthInBits;
  }

  public static QualifiedName staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static QualifiedName staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("QualifiedName");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int namespaceIndex = readSimpleField("namespaceIndex", readUnsignedInt(readBuffer, 16));

    PascalString name =
        readSimpleField(
            "name",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("QualifiedName");
    // Create the instance
    QualifiedName _qualifiedName;
    _qualifiedName = new QualifiedName(namespaceIndex, name);
    return _qualifiedName;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof QualifiedName)) {
      return false;
    }
    QualifiedName that = (QualifiedName) o;
    return (getNamespaceIndex() == that.getNamespaceIndex())
        && (getName() == that.getName())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getNamespaceIndex(), getName());
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
