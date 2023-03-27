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

public class PublishRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "826";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final int noOfSubscriptionAcknowledgements;
  protected final List<ExtensionObjectDefinition> subscriptionAcknowledgements;

  public PublishRequest(
      ExtensionObjectDefinition requestHeader,
      int noOfSubscriptionAcknowledgements,
      List<ExtensionObjectDefinition> subscriptionAcknowledgements) {
    super();
    this.requestHeader = requestHeader;
    this.noOfSubscriptionAcknowledgements = noOfSubscriptionAcknowledgements;
    this.subscriptionAcknowledgements = subscriptionAcknowledgements;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public int getNoOfSubscriptionAcknowledgements() {
    return noOfSubscriptionAcknowledgements;
  }

  public List<ExtensionObjectDefinition> getSubscriptionAcknowledgements() {
    return subscriptionAcknowledgements;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("PublishRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (noOfSubscriptionAcknowledgements)
    writeSimpleField(
        "noOfSubscriptionAcknowledgements",
        noOfSubscriptionAcknowledgements,
        writeSignedInt(writeBuffer, 32));

    // Array Field (subscriptionAcknowledgements)
    writeComplexTypeArrayField(
        "subscriptionAcknowledgements", subscriptionAcknowledgements, writeBuffer);

    writeBuffer.popContext("PublishRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PublishRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (noOfSubscriptionAcknowledgements)
    lengthInBits += 32;

    // Array field
    if (subscriptionAcknowledgements != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : subscriptionAcknowledgements) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= subscriptionAcknowledgements.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("PublishRequest");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    int noOfSubscriptionAcknowledgements =
        readSimpleField("noOfSubscriptionAcknowledgements", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> subscriptionAcknowledgements =
        readCountArrayField(
            "subscriptionAcknowledgements",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("823")),
                readBuffer),
            noOfSubscriptionAcknowledgements);

    readBuffer.closeContext("PublishRequest");
    // Create the instance
    return new PublishRequestBuilderImpl(
        requestHeader, noOfSubscriptionAcknowledgements, subscriptionAcknowledgements);
  }

  public static class PublishRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final int noOfSubscriptionAcknowledgements;
    private final List<ExtensionObjectDefinition> subscriptionAcknowledgements;

    public PublishRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        int noOfSubscriptionAcknowledgements,
        List<ExtensionObjectDefinition> subscriptionAcknowledgements) {
      this.requestHeader = requestHeader;
      this.noOfSubscriptionAcknowledgements = noOfSubscriptionAcknowledgements;
      this.subscriptionAcknowledgements = subscriptionAcknowledgements;
    }

    public PublishRequest build() {
      PublishRequest publishRequest =
          new PublishRequest(
              requestHeader, noOfSubscriptionAcknowledgements, subscriptionAcknowledgements);
      return publishRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PublishRequest)) {
      return false;
    }
    PublishRequest that = (PublishRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getNoOfSubscriptionAcknowledgements() == that.getNoOfSubscriptionAcknowledgements())
        && (getSubscriptionAcknowledgements() == that.getSubscriptionAcknowledgements())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getNoOfSubscriptionAcknowledgements(),
        getSubscriptionAcknowledgements());
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
