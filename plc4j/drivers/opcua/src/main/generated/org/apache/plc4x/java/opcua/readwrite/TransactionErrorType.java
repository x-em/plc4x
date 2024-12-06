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

public class TransactionErrorType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 32287;
  }

  // Properties.
  protected final NodeId targetId;
  protected final StatusCode error;
  protected final LocalizedText message;

  public TransactionErrorType(NodeId targetId, StatusCode error, LocalizedText message) {
    super();
    this.targetId = targetId;
    this.error = error;
    this.message = message;
  }

  public NodeId getTargetId() {
    return targetId;
  }

  public StatusCode getError() {
    return error;
  }

  public LocalizedText getMessage() {
    return message;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TransactionErrorType");

    // Simple Field (targetId)
    writeSimpleField("targetId", targetId, writeComplex(writeBuffer));

    // Simple Field (error)
    writeSimpleField("error", error, writeComplex(writeBuffer));

    // Simple Field (message)
    writeSimpleField("message", message, writeComplex(writeBuffer));

    writeBuffer.popContext("TransactionErrorType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TransactionErrorType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (targetId)
    lengthInBits += targetId.getLengthInBits();

    // Simple field (error)
    lengthInBits += error.getLengthInBits();

    // Simple field (message)
    lengthInBits += message.getLengthInBits();

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("TransactionErrorType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId targetId =
        readSimpleField("targetId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    StatusCode error =
        readSimpleField("error", readComplex(() -> StatusCode.staticParse(readBuffer), readBuffer));

    LocalizedText message =
        readSimpleField(
            "message", readComplex(() -> LocalizedText.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("TransactionErrorType");
    // Create the instance
    return new TransactionErrorTypeBuilderImpl(targetId, error, message);
  }

  public static class TransactionErrorTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId targetId;
    private final StatusCode error;
    private final LocalizedText message;

    public TransactionErrorTypeBuilderImpl(
        NodeId targetId, StatusCode error, LocalizedText message) {
      this.targetId = targetId;
      this.error = error;
      this.message = message;
    }

    public TransactionErrorType build() {
      TransactionErrorType transactionErrorType =
          new TransactionErrorType(targetId, error, message);
      return transactionErrorType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TransactionErrorType)) {
      return false;
    }
    TransactionErrorType that = (TransactionErrorType) o;
    return (getTargetId() == that.getTargetId())
        && (getError() == that.getError())
        && (getMessage() == that.getMessage())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getTargetId(), getError(), getMessage());
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