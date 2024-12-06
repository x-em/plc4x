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

public class NodeIdTwoByte extends NodeIdTypeDefinition implements Message {

  // Accessors for discriminator values.
  public NodeIdType getNodeType() {
    return NodeIdType.nodeIdTypeTwoByte;
  }

  // Properties.
  protected final short id;

  public NodeIdTwoByte(short id) {
    super();
    this.id = id;
  }

  public short getId() {
    return id;
  }

  public String getIdentifier() {
    return String.valueOf(getId());
  }

  public short getNamespace() {
    return (short) (-(1));
  }

  @Override
  protected void serializeNodeIdTypeDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NodeIdTwoByte");

    // Simple Field (id)
    writeSimpleField("id", id, writeUnsignedShort(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    String identifier = getIdentifier();
    writeBuffer.writeVirtual("identifier", identifier);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    short namespace = getNamespace();
    writeBuffer.writeVirtual("namespace", namespace);

    writeBuffer.popContext("NodeIdTwoByte");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NodeIdTwoByte _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (id)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static NodeIdTypeDefinitionBuilder staticParseNodeIdTypeDefinitionBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("NodeIdTwoByte");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short id = readSimpleField("id", readUnsignedShort(readBuffer, 8));
    String identifier = readVirtualField("identifier", String.class, id);
    short namespace = readVirtualField("namespace", short.class, -(1));

    readBuffer.closeContext("NodeIdTwoByte");
    // Create the instance
    return new NodeIdTwoByteBuilderImpl(id);
  }

  public static class NodeIdTwoByteBuilderImpl
      implements NodeIdTypeDefinition.NodeIdTypeDefinitionBuilder {
    private final short id;

    public NodeIdTwoByteBuilderImpl(short id) {
      this.id = id;
    }

    public NodeIdTwoByte build() {
      NodeIdTwoByte nodeIdTwoByte = new NodeIdTwoByte(id);
      return nodeIdTwoByte;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NodeIdTwoByte)) {
      return false;
    }
    NodeIdTwoByte that = (NodeIdTwoByte) o;
    return (getId() == that.getId()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getId());
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