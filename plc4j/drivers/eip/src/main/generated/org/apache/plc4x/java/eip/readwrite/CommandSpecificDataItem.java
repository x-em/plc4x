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
package org.apache.plc4x.java.eip.readwrite;

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

public abstract class CommandSpecificDataItem implements Message {

  // Abstract accessors for discriminator values.
  public abstract Integer getItemType();

  public CommandSpecificDataItem() {
    super();
  }

  protected abstract void serializeCommandSpecificDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CommandSpecificDataItem");

    // Discriminator Field (itemType) (Used as input to a switch field)
    writeDiscriminatorField("itemType", getItemType(), writeUnsignedInt(writeBuffer, 16));

    // Switch field (Serialize the sub-type)
    serializeCommandSpecificDataItemChild(writeBuffer);

    writeBuffer.popContext("CommandSpecificDataItem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    CommandSpecificDataItem _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Discriminator Field (itemType)
    lengthInBits += 16;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static CommandSpecificDataItem staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("CommandSpecificDataItem");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int itemType = readDiscriminatorField("itemType", readUnsignedInt(readBuffer, 16));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    CommandSpecificDataItemBuilder builder = null;
    if (EvaluationHelper.equals(itemType, (int) 0x000C)) {
      builder = CipIdentity.staticParseCommandSpecificDataItemBuilder(readBuffer);
    } else if (EvaluationHelper.equals(itemType, (int) 0x0086)) {
      builder = CipSecurityInformation.staticParseCommandSpecificDataItemBuilder(readBuffer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "itemType="
              + itemType
              + "]");
    }

    readBuffer.closeContext("CommandSpecificDataItem");
    // Create the instance
    CommandSpecificDataItem _commandSpecificDataItem = builder.build();
    return _commandSpecificDataItem;
  }

  public interface CommandSpecificDataItemBuilder {
    CommandSpecificDataItem build();
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CommandSpecificDataItem)) {
      return false;
    }
    CommandSpecificDataItem that = (CommandSpecificDataItem) o;
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
