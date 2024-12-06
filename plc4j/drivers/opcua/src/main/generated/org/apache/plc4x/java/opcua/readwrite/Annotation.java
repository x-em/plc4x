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

public class Annotation extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 893;
  }

  // Properties.
  protected final PascalString message;
  protected final PascalString userName;
  protected final long annotationTime;

  public Annotation(PascalString message, PascalString userName, long annotationTime) {
    super();
    this.message = message;
    this.userName = userName;
    this.annotationTime = annotationTime;
  }

  public PascalString getMessage() {
    return message;
  }

  public PascalString getUserName() {
    return userName;
  }

  public long getAnnotationTime() {
    return annotationTime;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Annotation");

    // Simple Field (message)
    writeSimpleField("message", message, writeComplex(writeBuffer));

    // Simple Field (userName)
    writeSimpleField("userName", userName, writeComplex(writeBuffer));

    // Simple Field (annotationTime)
    writeSimpleField("annotationTime", annotationTime, writeSignedLong(writeBuffer, 64));

    writeBuffer.popContext("Annotation");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Annotation _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (message)
    lengthInBits += message.getLengthInBits();

    // Simple field (userName)
    lengthInBits += userName.getLengthInBits();

    // Simple field (annotationTime)
    lengthInBits += 64;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("Annotation");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString message =
        readSimpleField(
            "message", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    PascalString userName =
        readSimpleField(
            "userName", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    long annotationTime = readSimpleField("annotationTime", readSignedLong(readBuffer, 64));

    readBuffer.closeContext("Annotation");
    // Create the instance
    return new AnnotationBuilderImpl(message, userName, annotationTime);
  }

  public static class AnnotationBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString message;
    private final PascalString userName;
    private final long annotationTime;

    public AnnotationBuilderImpl(PascalString message, PascalString userName, long annotationTime) {
      this.message = message;
      this.userName = userName;
      this.annotationTime = annotationTime;
    }

    public Annotation build() {
      Annotation annotation = new Annotation(message, userName, annotationTime);
      return annotation;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Annotation)) {
      return false;
    }
    Annotation that = (Annotation) o;
    return (getMessage() == that.getMessage())
        && (getUserName() == that.getUserName())
        && (getAnnotationTime() == that.getAnnotationTime())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getMessage(), getUserName(), getAnnotationTime());
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
