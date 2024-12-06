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

public abstract class Variant implements Message {

  // Abstract accessors for discriminator values.
  public abstract Byte getVariantType();

  // Properties.
  protected final boolean arrayLengthSpecified;
  protected final boolean arrayDimensionsSpecified;
  protected final Integer noOfArrayDimensions;
  protected final List<Boolean> arrayDimensions;

  public Variant(
      boolean arrayLengthSpecified,
      boolean arrayDimensionsSpecified,
      Integer noOfArrayDimensions,
      List<Boolean> arrayDimensions) {
    super();
    this.arrayLengthSpecified = arrayLengthSpecified;
    this.arrayDimensionsSpecified = arrayDimensionsSpecified;
    this.noOfArrayDimensions = noOfArrayDimensions;
    this.arrayDimensions = arrayDimensions;
  }

  public boolean getArrayLengthSpecified() {
    return arrayLengthSpecified;
  }

  public boolean getArrayDimensionsSpecified() {
    return arrayDimensionsSpecified;
  }

  public Integer getNoOfArrayDimensions() {
    return noOfArrayDimensions;
  }

  public List<Boolean> getArrayDimensions() {
    return arrayDimensions;
  }

  protected abstract void serializeVariantChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Variant");

    // Simple Field (arrayLengthSpecified)
    writeSimpleField("arrayLengthSpecified", arrayLengthSpecified, writeBoolean(writeBuffer));

    // Simple Field (arrayDimensionsSpecified)
    writeSimpleField(
        "arrayDimensionsSpecified", arrayDimensionsSpecified, writeBoolean(writeBuffer));

    // Discriminator Field (VariantType) (Used as input to a switch field)
    writeDiscriminatorField("VariantType", getVariantType(), writeUnsignedByte(writeBuffer, 6));

    // Switch field (Serialize the sub-type)
    serializeVariantChild(writeBuffer);

    // Optional Field (noOfArrayDimensions) (Can be skipped, if the value is null)
    writeOptionalField("noOfArrayDimensions", noOfArrayDimensions, writeSignedInt(writeBuffer, 32));

    // Array Field (arrayDimensions)
    writeSimpleTypeArrayField("arrayDimensions", arrayDimensions, writeBoolean(writeBuffer));

    writeBuffer.popContext("Variant");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    Variant _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (arrayLengthSpecified)
    lengthInBits += 1;

    // Simple field (arrayDimensionsSpecified)
    lengthInBits += 1;

    // Discriminator Field (VariantType)
    lengthInBits += 6;

    // Length of sub-type elements will be added by sub-type...

    // Optional Field (noOfArrayDimensions)
    if (noOfArrayDimensions != null) {
      lengthInBits += 32;
    }

    // Array field
    if (arrayDimensions != null) {
      lengthInBits += 1 * arrayDimensions.size();
    }

    return lengthInBits;
  }

  public static Variant staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("Variant");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    boolean arrayLengthSpecified = readSimpleField("arrayLengthSpecified", readBoolean(readBuffer));

    boolean arrayDimensionsSpecified =
        readSimpleField("arrayDimensionsSpecified", readBoolean(readBuffer));

    byte VariantType = readDiscriminatorField("VariantType", readUnsignedByte(readBuffer, 6));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    VariantBuilder builder = null;
    if (EvaluationHelper.equals(VariantType, (byte) 0)) {
      builder = VariantNull.staticParseVariantBuilder(readBuffer);
    } else if (EvaluationHelper.equals(VariantType, (byte) 1)) {
      builder = VariantBoolean.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 2)) {
      builder = VariantSByte.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 3)) {
      builder = VariantByte.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 4)) {
      builder = VariantInt16.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 5)) {
      builder = VariantUInt16.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 6)) {
      builder = VariantInt32.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 7)) {
      builder = VariantUInt32.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 8)) {
      builder = VariantInt64.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 9)) {
      builder = VariantUInt64.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 10)) {
      builder = VariantFloat.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 11)) {
      builder = VariantDouble.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 12)) {
      builder = VariantString.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 13)) {
      builder = VariantDateTime.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 14)) {
      builder = VariantGuid.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 15)) {
      builder = VariantByteString.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 16)) {
      builder = VariantXmlElement.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 17)) {
      builder = VariantNodeId.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 18)) {
      builder = VariantExpandedNodeId.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 19)) {
      builder = VariantStatusCode.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 20)) {
      builder = VariantQualifiedName.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 21)) {
      builder = VariantLocalizedText.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 22)) {
      builder = VariantExtensionObject.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 23)) {
      builder = VariantDataValue.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 24)) {
      builder = VariantVariant.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    } else if (EvaluationHelper.equals(VariantType, (byte) 25)) {
      builder = VariantDiagnosticInfo.staticParseVariantBuilder(readBuffer, arrayLengthSpecified);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "VariantType="
              + VariantType
              + " "
              + "arrayLengthSpecified="
              + arrayLengthSpecified
              + "]");
    }

    Integer noOfArrayDimensions =
        readOptionalField(
            "noOfArrayDimensions", readSignedInt(readBuffer, 32), arrayDimensionsSpecified);

    List<Boolean> arrayDimensions =
        readCountArrayField(
            "arrayDimensions",
            readBoolean(readBuffer),
            (((noOfArrayDimensions) == (null)) ? 0 : noOfArrayDimensions));

    readBuffer.closeContext("Variant");
    // Create the instance
    Variant _variant =
        builder.build(
            arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions);
    return _variant;
  }

  public interface VariantBuilder {
    Variant build(
        boolean arrayLengthSpecified,
        boolean arrayDimensionsSpecified,
        Integer noOfArrayDimensions,
        List<Boolean> arrayDimensions);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Variant)) {
      return false;
    }
    Variant that = (Variant) o;
    return (getArrayLengthSpecified() == that.getArrayLengthSpecified())
        && (getArrayDimensionsSpecified() == that.getArrayDimensionsSpecified())
        && (getNoOfArrayDimensions() == that.getNoOfArrayDimensions())
        && (getArrayDimensions() == that.getArrayDimensions())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getArrayLengthSpecified(),
        getArrayDimensionsSpecified(),
        getNoOfArrayDimensions(),
        getArrayDimensions());
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