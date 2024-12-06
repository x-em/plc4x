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

public class HistoryEvent extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 661;
  }

  // Properties.
  protected final List<HistoryEventFieldList> events;

  public HistoryEvent(List<HistoryEventFieldList> events) {
    super();
    this.events = events;
  }

  public List<HistoryEventFieldList> getEvents() {
    return events;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("HistoryEvent");

    // Implicit Field (noOfEvents) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int noOfEvents = (int) ((((getEvents()) == (null)) ? -(1) : COUNT(getEvents())));
    writeImplicitField("noOfEvents", noOfEvents, writeSignedInt(writeBuffer, 32));

    // Array Field (events)
    writeComplexTypeArrayField("events", events, writeBuffer);

    writeBuffer.popContext("HistoryEvent");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    HistoryEvent _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (noOfEvents)
    lengthInBits += 32;

    // Array field
    if (events != null) {
      int i = 0;
      for (HistoryEventFieldList element : events) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= events.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("HistoryEvent");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfEvents = readImplicitField("noOfEvents", readSignedInt(readBuffer, 32));

    List<HistoryEventFieldList> events =
        readCountArrayField(
            "events",
            readComplex(
                () ->
                    (HistoryEventFieldList)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (922)),
                readBuffer),
            noOfEvents);

    readBuffer.closeContext("HistoryEvent");
    // Create the instance
    return new HistoryEventBuilderImpl(events);
  }

  public static class HistoryEventBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final List<HistoryEventFieldList> events;

    public HistoryEventBuilderImpl(List<HistoryEventFieldList> events) {
      this.events = events;
    }

    public HistoryEvent build() {
      HistoryEvent historyEvent = new HistoryEvent(events);
      return historyEvent;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof HistoryEvent)) {
      return false;
    }
    HistoryEvent that = (HistoryEvent) o;
    return (getEvents() == that.getEvents()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getEvents());
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