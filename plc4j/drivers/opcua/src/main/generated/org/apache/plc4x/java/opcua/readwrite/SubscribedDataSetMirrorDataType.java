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

public class SubscribedDataSetMirrorDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public Integer getExtensionId() {
    return (int) 15637;
  }

  // Properties.
  protected final PascalString parentNodeName;
  protected final List<RolePermissionType> rolePermissions;

  public SubscribedDataSetMirrorDataType(
      PascalString parentNodeName, List<RolePermissionType> rolePermissions) {
    super();
    this.parentNodeName = parentNodeName;
    this.rolePermissions = rolePermissions;
  }

  public PascalString getParentNodeName() {
    return parentNodeName;
  }

  public List<RolePermissionType> getRolePermissions() {
    return rolePermissions;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SubscribedDataSetMirrorDataType");

    // Simple Field (parentNodeName)
    writeSimpleField("parentNodeName", parentNodeName, writeComplex(writeBuffer));

    // Implicit Field (noOfRolePermissions) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int noOfRolePermissions =
        (int) ((((getRolePermissions()) == (null)) ? -(1) : COUNT(getRolePermissions())));
    writeImplicitField("noOfRolePermissions", noOfRolePermissions, writeSignedInt(writeBuffer, 32));

    // Array Field (rolePermissions)
    writeComplexTypeArrayField("rolePermissions", rolePermissions, writeBuffer);

    writeBuffer.popContext("SubscribedDataSetMirrorDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SubscribedDataSetMirrorDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (parentNodeName)
    lengthInBits += parentNodeName.getLengthInBits();

    // Implicit Field (noOfRolePermissions)
    lengthInBits += 32;

    // Array field
    if (rolePermissions != null) {
      int i = 0;
      for (RolePermissionType element : rolePermissions) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= rolePermissions.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, Integer extensionId) throws ParseException {
    readBuffer.pullContext("SubscribedDataSetMirrorDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalString parentNodeName =
        readSimpleField(
            "parentNodeName", readComplex(() -> PascalString.staticParse(readBuffer), readBuffer));

    int noOfRolePermissions =
        readImplicitField("noOfRolePermissions", readSignedInt(readBuffer, 32));

    List<RolePermissionType> rolePermissions =
        readCountArrayField(
            "rolePermissions",
            readComplex(
                () ->
                    (RolePermissionType)
                        ExtensionObjectDefinition.staticParse(readBuffer, (int) (98)),
                readBuffer),
            noOfRolePermissions);

    readBuffer.closeContext("SubscribedDataSetMirrorDataType");
    // Create the instance
    return new SubscribedDataSetMirrorDataTypeBuilderImpl(parentNodeName, rolePermissions);
  }

  public static class SubscribedDataSetMirrorDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final PascalString parentNodeName;
    private final List<RolePermissionType> rolePermissions;

    public SubscribedDataSetMirrorDataTypeBuilderImpl(
        PascalString parentNodeName, List<RolePermissionType> rolePermissions) {
      this.parentNodeName = parentNodeName;
      this.rolePermissions = rolePermissions;
    }

    public SubscribedDataSetMirrorDataType build() {
      SubscribedDataSetMirrorDataType subscribedDataSetMirrorDataType =
          new SubscribedDataSetMirrorDataType(parentNodeName, rolePermissions);
      return subscribedDataSetMirrorDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SubscribedDataSetMirrorDataType)) {
      return false;
    }
    SubscribedDataSetMirrorDataType that = (SubscribedDataSetMirrorDataType) o;
    return (getParentNodeName() == that.getParentNodeName())
        && (getRolePermissions() == that.getRolePermissions())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getParentNodeName(), getRolePermissions());
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
