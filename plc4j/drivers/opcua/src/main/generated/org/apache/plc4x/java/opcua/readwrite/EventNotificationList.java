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

public class EventNotificationList extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 916;
  }

  // Properties.
  protected final List<EventFieldList> events;

  public EventNotificationList(List<EventFieldList> events) {
    super();
    this.events = events;
  }

  public List<EventFieldList> getEvents() {
    return events;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EventNotificationList");

    // Implicit Field (noOfEvents) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    int noOfEvents = (int) ((((getEvents()) == (null)) ? -(1) : COUNT(getEvents())));
    writeImplicitField("noOfEvents", noOfEvents, writeSignedInt(writeBuffer, 32));

    // Array Field (events)
    writeComplexTypeArrayField("events", events, writeBuffer);

    writeBuffer.popContext("EventNotificationList");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EventNotificationList _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (noOfEvents)
    lengthInBits += 32;

    // Array field
    if (events != null) {
      int i = 0;
      for (EventFieldList element : events) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= events.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("EventNotificationList");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfEvents = readImplicitField("noOfEvents", readSignedInt(readBuffer, 32));

    List<EventFieldList> events =
        readCountArrayField(
            "events",
            readComplex(
                () ->
                    (EventFieldList) ExtensionObjectDefinition.staticParse(readBuffer, (int) (919)),
                readBuffer),
            noOfEvents);

    readBuffer.closeContext("EventNotificationList");
    // Create the instance
    return new EventNotificationListBuilderImpl(events);
  }

  public static class EventNotificationListBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final List<EventFieldList> events;

    public EventNotificationListBuilderImpl(List<EventFieldList> events) {
      this.events = events;
    }

    public EventNotificationList build() {
      EventNotificationList eventNotificationList = new EventNotificationList(events);
      return eventNotificationList;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EventNotificationList)) {
      return false;
    }
    EventNotificationList that = (EventNotificationList) o;
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
