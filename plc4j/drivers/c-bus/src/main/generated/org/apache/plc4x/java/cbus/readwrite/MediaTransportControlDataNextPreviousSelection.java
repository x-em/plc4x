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

public class MediaTransportControlDataNextPreviousSelection extends MediaTransportControlData
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte operation;

  public MediaTransportControlDataNextPreviousSelection(
      MediaTransportControlCommandTypeContainer commandTypeContainer,
      byte mediaLinkGroup,
      byte operation) {
    super(commandTypeContainer, mediaLinkGroup);
    this.operation = operation;
  }

  public byte getOperation() {
    return operation;
  }

  public boolean getIsSetThePreviousSelection() {
    return (boolean) ((getOperation()) == (0x00));
  }

  public boolean getIsSetTheNextSelection() {
    return (boolean) ((getOperation()) != (0x00));
  }

  @Override
  protected void serializeMediaTransportControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MediaTransportControlDataNextPreviousSelection");

    // Simple Field (operation)
    writeSimpleField("operation", operation, writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isSetThePreviousSelection = getIsSetThePreviousSelection();
    writeBuffer.writeVirtual("isSetThePreviousSelection", isSetThePreviousSelection);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isSetTheNextSelection = getIsSetTheNextSelection();
    writeBuffer.writeVirtual("isSetTheNextSelection", isSetTheNextSelection);

    writeBuffer.popContext("MediaTransportControlDataNextPreviousSelection");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MediaTransportControlDataNextPreviousSelection _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (operation)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static MediaTransportControlDataBuilder staticParseMediaTransportControlDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("MediaTransportControlDataNextPreviousSelection");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte operation = readSimpleField("operation", readByte(readBuffer, 8));
    boolean isSetThePreviousSelection =
        readVirtualField("isSetThePreviousSelection", boolean.class, (operation) == (0x00));
    boolean isSetTheNextSelection =
        readVirtualField("isSetTheNextSelection", boolean.class, (operation) != (0x00));

    readBuffer.closeContext("MediaTransportControlDataNextPreviousSelection");
    // Create the instance
    return new MediaTransportControlDataNextPreviousSelectionBuilderImpl(operation);
  }

  public static class MediaTransportControlDataNextPreviousSelectionBuilderImpl
      implements MediaTransportControlData.MediaTransportControlDataBuilder {
    private final byte operation;

    public MediaTransportControlDataNextPreviousSelectionBuilderImpl(byte operation) {
      this.operation = operation;
    }

    public MediaTransportControlDataNextPreviousSelection build(
        MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup) {
      MediaTransportControlDataNextPreviousSelection
          mediaTransportControlDataNextPreviousSelection =
              new MediaTransportControlDataNextPreviousSelection(
                  commandTypeContainer, mediaLinkGroup, operation);
      return mediaTransportControlDataNextPreviousSelection;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MediaTransportControlDataNextPreviousSelection)) {
      return false;
    }
    MediaTransportControlDataNextPreviousSelection that =
        (MediaTransportControlDataNextPreviousSelection) o;
    return (getOperation() == that.getOperation()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getOperation());
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
