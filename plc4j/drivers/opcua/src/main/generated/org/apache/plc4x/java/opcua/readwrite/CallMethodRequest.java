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

public class CallMethodRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 706;
  }

  // Properties.
  protected final NodeId objectId;
  protected final NodeId methodId;
  protected final List<Variant> inputArguments;

  public CallMethodRequest(NodeId objectId, NodeId methodId, List<Variant> inputArguments) {
    super();
    this.objectId = objectId;
    this.methodId = methodId;
    this.inputArguments = inputArguments;
  }

  public NodeId getObjectId() {
    return objectId;
  }

  public NodeId getMethodId() {
    return methodId;
  }

  public List<Variant> getInputArguments() {
    return inputArguments;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CallMethodRequest");

    // Simple Field (objectId)
    writeSimpleField("objectId", objectId, writeComplex(writeBuffer));

    // Simple Field (methodId)
    writeSimpleField("methodId", methodId, writeComplex(writeBuffer));

    // Implicit Field (noOfInputArguments) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfInputArguments =
        (int) ((((getInputArguments()) == (null)) ? -(1) : COUNT(getInputArguments())));
    writeImplicitField("noOfInputArguments", noOfInputArguments, writeSignedInt(writeBuffer, 32));

    // Array Field (inputArguments)
    writeComplexTypeArrayField("inputArguments", inputArguments, writeBuffer);

    writeBuffer.popContext("CallMethodRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CallMethodRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (objectId)
    lengthInBits += objectId.getLengthInBits();

    // Simple field (methodId)
    lengthInBits += methodId.getLengthInBits();

    // Implicit Field (noOfInputArguments)
    lengthInBits += 32;

    // Array field
    if (inputArguments != null) {
      int i = 0;
      for (Variant element : inputArguments) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= inputArguments.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("CallMethodRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId objectId =
        readSimpleField("objectId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    NodeId methodId =
        readSimpleField("methodId", readComplex(() -> NodeId.staticParse(readBuffer), readBuffer));

    int noOfInputArguments = readImplicitField("noOfInputArguments", readSignedInt(readBuffer, 32));

    List<Variant> inputArguments =
        readCountArrayField(
            "inputArguments",
            readComplex(() -> Variant.staticParse(readBuffer), readBuffer),
            noOfInputArguments);

    readBuffer.closeContext("CallMethodRequest");
    // Create the instance
    return new CallMethodRequestBuilderImpl(objectId, methodId, inputArguments);
  }

  public static class CallMethodRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId objectId;
    private final NodeId methodId;
    private final List<Variant> inputArguments;

    public CallMethodRequestBuilderImpl(
        NodeId objectId, NodeId methodId, List<Variant> inputArguments) {
      this.objectId = objectId;
      this.methodId = methodId;
      this.inputArguments = inputArguments;
    }

    public CallMethodRequest build() {
      CallMethodRequest callMethodRequest =
          new CallMethodRequest(objectId, methodId, inputArguments);
      return callMethodRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CallMethodRequest)) {
      return false;
    }
    CallMethodRequest that = (CallMethodRequest) o;
    return (getObjectId() == that.getObjectId())
        && (getMethodId() == that.getMethodId())
        && (getInputArguments() == that.getInputArguments())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getObjectId(), getMethodId(), getInputArguments());
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