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

public class CALDataIdentifyReply extends CALData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final Attribute attribute;
  protected final IdentifyReplyCommand identifyReplyCommand;

  // Arguments.
  protected final RequestContext requestContext;

  public CALDataIdentifyReply(
      CALCommandTypeContainer commandTypeContainer,
      CALData additionalData,
      Attribute attribute,
      IdentifyReplyCommand identifyReplyCommand,
      RequestContext requestContext) {
    super(commandTypeContainer, additionalData, requestContext);
    this.attribute = attribute;
    this.identifyReplyCommand = identifyReplyCommand;
    this.requestContext = requestContext;
  }

  public Attribute getAttribute() {
    return attribute;
  }

  public IdentifyReplyCommand getIdentifyReplyCommand() {
    return identifyReplyCommand;
  }

  @Override
  protected void serializeCALDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CALDataIdentifyReply");

    // Simple Field (attribute)
    writeSimpleEnumField(
        "attribute",
        "Attribute",
        attribute,
        writeEnum(Attribute::getValue, Attribute::name, writeUnsignedShort(writeBuffer, 8)));

    // Simple Field (identifyReplyCommand)
    writeSimpleField("identifyReplyCommand", identifyReplyCommand, writeComplex(writeBuffer));

    writeBuffer.popContext("CALDataIdentifyReply");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CALDataIdentifyReply _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (attribute)
    lengthInBits += 8;

    // Simple field (identifyReplyCommand)
    lengthInBits += identifyReplyCommand.getLengthInBits();

    return lengthInBits;
  }

  public static CALDataBuilder staticParseCALDataBuilder(
      ReadBuffer readBuffer,
      CALCommandTypeContainer commandTypeContainer,
      RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("CALDataIdentifyReply");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Attribute attribute =
        readEnumField(
            "attribute",
            "Attribute",
            readEnum(Attribute::enumForValue, readUnsignedShort(readBuffer, 8)));

    IdentifyReplyCommand identifyReplyCommand =
        readSimpleField(
            "identifyReplyCommand",
            readComplex(
                () ->
                    IdentifyReplyCommand.staticParse(
                        readBuffer,
                        (Attribute) (attribute),
                        (byte) ((commandTypeContainer.getNumBytes()) - (1))),
                readBuffer));

    readBuffer.closeContext("CALDataIdentifyReply");
    // Create the instance
    return new CALDataIdentifyReplyBuilderImpl(attribute, identifyReplyCommand, requestContext);
  }

  public static class CALDataIdentifyReplyBuilderImpl implements CALData.CALDataBuilder {
    private final Attribute attribute;
    private final IdentifyReplyCommand identifyReplyCommand;
    private final RequestContext requestContext;

    public CALDataIdentifyReplyBuilderImpl(
        Attribute attribute,
        IdentifyReplyCommand identifyReplyCommand,
        RequestContext requestContext) {
      this.attribute = attribute;
      this.identifyReplyCommand = identifyReplyCommand;
      this.requestContext = requestContext;
    }

    public CALDataIdentifyReply build(
        CALCommandTypeContainer commandTypeContainer,
        CALData additionalData,
        RequestContext requestContext) {
      CALDataIdentifyReply cALDataIdentifyReply =
          new CALDataIdentifyReply(
              commandTypeContainer,
              additionalData,
              attribute,
              identifyReplyCommand,
              requestContext);
      return cALDataIdentifyReply;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CALDataIdentifyReply)) {
      return false;
    }
    CALDataIdentifyReply that = (CALDataIdentifyReply) o;
    return (getAttribute() == that.getAttribute())
        && (getIdentifyReplyCommand() == that.getIdentifyReplyCommand())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getAttribute(), getIdentifyReplyCommand());
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