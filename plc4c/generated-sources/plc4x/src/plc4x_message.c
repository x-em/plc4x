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

#include <stdio.h>
#include <plc4c/spi/evaluation_helper.h>
#include <plc4c/driver_plc4x_static.h>

#include "plc4x_message.h"

// Code generated by code-generation. DO NOT EDIT.

// Array of discriminator values that match the enum type constants.
// (The order is identical to the enum constants, so we can use the
// enum constant to directly access a given type's discriminator values)
const plc4c_plc4x_read_write_plc4x_message_discriminator plc4c_plc4x_read_write_plc4x_message_discriminators[] = {
  {/* plc4c_plc4x_read_write_plc4x_connect_request */
   .requestType = plc4c_plc4x_read_write_plc4x_request_type_CONNECT_REQUEST },
  {/* plc4c_plc4x_read_write_plc4x_connect_response */
   .requestType = plc4c_plc4x_read_write_plc4x_request_type_CONNECT_RESPONSE },
  {/* plc4c_plc4x_read_write_plc4x_read_request */
   .requestType = plc4c_plc4x_read_write_plc4x_request_type_READ_REQUEST },
  {/* plc4c_plc4x_read_write_plc4x_read_response */
   .requestType = plc4c_plc4x_read_write_plc4x_request_type_READ_RESPONSE },
  {/* plc4c_plc4x_read_write_plc4x_write_request */
   .requestType = plc4c_plc4x_read_write_plc4x_request_type_WRITE_REQUEST },
  {/* plc4c_plc4x_read_write_plc4x_write_response */
   .requestType = plc4c_plc4x_read_write_plc4x_request_type_WRITE_RESPONSE }

};

// Function returning the discriminator values for a given type constant.
plc4c_plc4x_read_write_plc4x_message_discriminator plc4c_plc4x_read_write_plc4x_message_get_discriminator(plc4c_plc4x_read_write_plc4x_message_type type) {
  return plc4c_plc4x_read_write_plc4x_message_discriminators[type];
}

// Create an empty NULL-struct
static const plc4c_plc4x_read_write_plc4x_message plc4c_plc4x_read_write_plc4x_message_null_const;

plc4c_plc4x_read_write_plc4x_message plc4c_plc4x_read_write_plc4x_message_null() {
  return plc4c_plc4x_read_write_plc4x_message_null_const;
}


// Constant values.
static const uint8_t PLC4C_PLC4X_READ_WRITE_PLC4X_MESSAGE_VERSION_const = 0x01;
uint8_t PLC4C_PLC4X_READ_WRITE_PLC4X_MESSAGE_VERSION() {
  return PLC4C_PLC4X_READ_WRITE_PLC4X_MESSAGE_VERSION_const;
}

// Parse function.
plc4c_return_code plc4c_plc4x_read_write_plc4x_message_parse(plc4c_spi_read_buffer* readBuffer, plc4c_plc4x_read_write_plc4x_message** _message) {
  uint16_t startPos = plc4c_spi_read_get_pos(readBuffer);
  plc4c_return_code _res = OK;

  // Allocate enough memory to contain this data structure.
  (*_message) = malloc(sizeof(plc4c_plc4x_read_write_plc4x_message));
  if(*_message == NULL) {
    return NO_MEMORY;
  }

  // Const Field (version)
  uint8_t version = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &version);
  if(_res != OK) {
    return _res;
  }
  if(version != PLC4C_PLC4X_READ_WRITE_PLC4X_MESSAGE_VERSION()) {
    return PARSE_ERROR;
    // throw new ParseException("Expected constant value " + PLC4C_PLC4X_READ_WRITE_PLC4X_MESSAGE_VERSION + " but got " + version);
  }

  // Implicit Field (packetLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint16_t packetLength = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &packetLength);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (requestId)
  uint16_t requestId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &requestId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->request_id = requestId;
  // Discriminator Field (requestType)
  plc4c_plc4x_read_write_plc4x_request_type requestType;
  _res = plc4c_plc4x_read_write_plc4x_request_type_parse(readBuffer, &requestType);
  if(_res != OK) {
    return _res;
  }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
if( requestType == plc4c_plc4x_read_write_plc4x_request_type_CONNECT_REQUEST ) { /* Plc4xConnectRequest */
    (*_message)->_type = plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_connect_request;

  // Implicit Field (connectionStringLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t connectionStringLen = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &connectionStringLen);
  if(_res != OK) {
    return _res;
  }


  // Simple Field (connectionString)
  char* connectionString = "";
  _res = plc4c_spi_read_string(readBuffer, (connectionStringLen) * (8), "UTF-8", (char**) &connectionString);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_connect_request_connection_string = connectionString;
  } else 
if( requestType == plc4c_plc4x_read_write_plc4x_request_type_CONNECT_RESPONSE ) { /* Plc4xConnectResponse */
    (*_message)->_type = plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_connect_response;

  // Simple Field (connectionId)
  uint16_t connectionId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &connectionId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_connect_response_connection_id = connectionId;


  // Simple Field (responseCode)
  plc4c_plc4x_read_write_plc4x_response_code responseCode;
  _res = plc4c_plc4x_read_write_plc4x_response_code_parse(readBuffer, (void*) &responseCode);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_connect_response_response_code = responseCode;
  } else 
if( requestType == plc4c_plc4x_read_write_plc4x_request_type_READ_REQUEST ) { /* Plc4xReadRequest */
    (*_message)->_type = plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_read_request;

  // Simple Field (connectionId)
  uint16_t connectionId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &connectionId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_read_request_connection_id = connectionId;


  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t numTags = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &numTags);
  if(_res != OK) {
    return _res;
  }


  // Array field (tags)
  plc4c_list* tags = NULL;
  plc4c_utils_list_create(&tags);
  if(tags == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) numTags;
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_request* _value = NULL;
      _res = plc4c_plc4x_read_write_plc4x_tag_request_parse(readBuffer, (void*) &_value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(tags, _value);
    }
  }
  (*_message)->plc4x_read_request_tags = tags;
  } else 
if( requestType == plc4c_plc4x_read_write_plc4x_request_type_READ_RESPONSE ) { /* Plc4xReadResponse */
    (*_message)->_type = plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_read_response;

  // Simple Field (connectionId)
  uint16_t connectionId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &connectionId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_read_response_connection_id = connectionId;


  // Simple Field (responseCode)
  plc4c_plc4x_read_write_plc4x_response_code responseCode;
  _res = plc4c_plc4x_read_write_plc4x_response_code_parse(readBuffer, (void*) &responseCode);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_read_response_response_code = responseCode;


  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t numTags = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &numTags);
  if(_res != OK) {
    return _res;
  }


  // Array field (tags)
  plc4c_list* tags = NULL;
  plc4c_utils_list_create(&tags);
  if(tags == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) numTags;
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_value_response* _value = NULL;
      _res = plc4c_plc4x_read_write_plc4x_tag_value_response_parse(readBuffer, (void*) &_value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(tags, _value);
    }
  }
  (*_message)->plc4x_read_response_tags = tags;
  } else 
if( requestType == plc4c_plc4x_read_write_plc4x_request_type_WRITE_REQUEST ) { /* Plc4xWriteRequest */
    (*_message)->_type = plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_write_request;

  // Simple Field (connectionId)
  uint16_t connectionId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &connectionId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_write_request_connection_id = connectionId;


  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t numTags = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &numTags);
  if(_res != OK) {
    return _res;
  }


  // Array field (tags)
  plc4c_list* tags = NULL;
  plc4c_utils_list_create(&tags);
  if(tags == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) numTags;
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_value_request* _value = NULL;
      _res = plc4c_plc4x_read_write_plc4x_tag_value_request_parse(readBuffer, (void*) &_value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(tags, _value);
    }
  }
  (*_message)->plc4x_write_request_tags = tags;
  } else 
if( requestType == plc4c_plc4x_read_write_plc4x_request_type_WRITE_RESPONSE ) { /* Plc4xWriteResponse */
    (*_message)->_type = plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_write_response;

  // Simple Field (connectionId)
  uint16_t connectionId = 0;
  _res = plc4c_spi_read_unsigned_short(readBuffer, 16, (uint16_t*) &connectionId);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_write_response_connection_id = connectionId;


  // Simple Field (responseCode)
  plc4c_plc4x_read_write_plc4x_response_code responseCode;
  _res = plc4c_plc4x_read_write_plc4x_response_code_parse(readBuffer, (void*) &responseCode);
  if(_res != OK) {
    return _res;
  }
  (*_message)->plc4x_write_response_response_code = responseCode;


  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  uint8_t numTags = 0;
  _res = plc4c_spi_read_unsigned_byte(readBuffer, 8, (uint8_t*) &numTags);
  if(_res != OK) {
    return _res;
  }


  // Array field (tags)
  plc4c_list* tags = NULL;
  plc4c_utils_list_create(&tags);
  if(tags == NULL) {
    return NO_MEMORY;
  }
  {
    // Count array
    uint16_t itemCount = (uint16_t) numTags;
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_response* _value = NULL;
      _res = plc4c_plc4x_read_write_plc4x_tag_response_parse(readBuffer, (void*) &_value);
      if(_res != OK) {
        return _res;
      }
      plc4c_utils_list_insert_head_value(tags, _value);
    }
  }
  (*_message)->plc4x_write_response_tags = tags;
  }

  return OK;
}

plc4c_return_code plc4c_plc4x_read_write_plc4x_message_serialize(plc4c_spi_write_buffer* writeBuffer, plc4c_plc4x_read_write_plc4x_message* _message) {
  plc4c_return_code _res = OK;

  // Const Field (version)
  plc4c_spi_write_unsigned_byte(writeBuffer, 8, PLC4C_PLC4X_READ_WRITE_PLC4X_MESSAGE_VERSION());

  // Implicit Field (packetLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, plc4c_plc4x_read_write_plc4x_message_length_in_bytes(_message));
  if(_res != OK) {
    return _res;
  }

  // Simple Field (requestId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->request_id);
  if(_res != OK) {
    return _res;
  }

  // Enumerated Discriminator Field (requestType)
  plc4c_plc4x_read_write_plc4x_request_type _requestType = plc4c_plc4x_read_write_plc4x_message_get_discriminator(_message->_type).requestType;
  _res = plc4c_plc4x_read_write_plc4x_request_type_serialize(writeBuffer, &_requestType);
  if(_res != OK) {
    return _res;
  }

  // Switch Field (Depending on the current type, serialize the subtype elements)
  switch(_message->_type) {
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_connect_request: {

  // Implicit Field (connectionStringLen) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_spi_evaluation_helper_str_len(_message->plc4x_connect_request_connection_string));
  if(_res != OK) {
    return _res;
  }

  // Simple Field (connectionString)
  _res = plc4c_spi_write_string(writeBuffer, (plc4c_spi_evaluation_helper_str_len(_message->plc4x_connect_request_connection_string)) * (8), "UTF-8", _message->plc4x_connect_request_connection_string);
  if(_res != OK) {
    return _res;
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_connect_response: {

  // Simple Field (connectionId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->plc4x_connect_response_connection_id);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (responseCode)
  _res = plc4c_plc4x_read_write_plc4x_response_code_serialize(writeBuffer, &_message->plc4x_connect_response_response_code);
  if(_res != OK) {
    return _res;
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_read_request: {

  // Simple Field (connectionId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->plc4x_read_request_connection_id);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_spi_evaluation_helper_count(_message->plc4x_read_request_tags));
  if(_res != OK) {
    return _res;
  }

  // Array field (tags)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->plc4x_read_request_tags);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_request* _value = (plc4c_plc4x_read_write_plc4x_tag_request*) plc4c_utils_list_get_value(_message->plc4x_read_request_tags, curItem);
      _res = plc4c_plc4x_read_write_plc4x_tag_request_serialize(writeBuffer, (void*) _value);
      if(_res != OK) {
        return _res;
      }
    }
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_read_response: {

  // Simple Field (connectionId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->plc4x_read_response_connection_id);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (responseCode)
  _res = plc4c_plc4x_read_write_plc4x_response_code_serialize(writeBuffer, &_message->plc4x_read_response_response_code);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_spi_evaluation_helper_count(_message->plc4x_read_response_tags));
  if(_res != OK) {
    return _res;
  }

  // Array field (tags)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->plc4x_read_response_tags);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_value_response* _value = (plc4c_plc4x_read_write_plc4x_tag_value_response*) plc4c_utils_list_get_value(_message->plc4x_read_response_tags, curItem);
      _res = plc4c_plc4x_read_write_plc4x_tag_value_response_serialize(writeBuffer, (void*) _value);
      if(_res != OK) {
        return _res;
      }
    }
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_write_request: {

  // Simple Field (connectionId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->plc4x_write_request_connection_id);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_spi_evaluation_helper_count(_message->plc4x_write_request_tags));
  if(_res != OK) {
    return _res;
  }

  // Array field (tags)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->plc4x_write_request_tags);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_value_request* _value = (plc4c_plc4x_read_write_plc4x_tag_value_request*) plc4c_utils_list_get_value(_message->plc4x_write_request_tags, curItem);
      _res = plc4c_plc4x_read_write_plc4x_tag_value_request_serialize(writeBuffer, (void*) _value);
      if(_res != OK) {
        return _res;
      }
    }
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_write_response: {

  // Simple Field (connectionId)
  _res = plc4c_spi_write_unsigned_short(writeBuffer, 16, _message->plc4x_write_response_connection_id);
  if(_res != OK) {
    return _res;
  }

  // Simple Field (responseCode)
  _res = plc4c_plc4x_read_write_plc4x_response_code_serialize(writeBuffer, &_message->plc4x_write_response_response_code);
  if(_res != OK) {
    return _res;
  }

  // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
  _res = plc4c_spi_write_unsigned_byte(writeBuffer, 8, plc4c_spi_evaluation_helper_count(_message->plc4x_write_response_tags));
  if(_res != OK) {
    return _res;
  }

  // Array field (tags)
  {
    uint8_t itemCount = plc4c_utils_list_size(_message->plc4x_write_response_tags);
    for(int curItem = 0; curItem < itemCount; curItem++) {
      bool lastItem = curItem == (itemCount - 1);
      plc4c_plc4x_read_write_plc4x_tag_response* _value = (plc4c_plc4x_read_write_plc4x_tag_response*) plc4c_utils_list_get_value(_message->plc4x_write_response_tags, curItem);
      _res = plc4c_plc4x_read_write_plc4x_tag_response_serialize(writeBuffer, (void*) _value);
      if(_res != OK) {
        return _res;
      }
    }
  }

      break;
    }
  }

  return OK;
}

uint16_t plc4c_plc4x_read_write_plc4x_message_length_in_bytes(plc4c_plc4x_read_write_plc4x_message* _message) {
  return plc4c_plc4x_read_write_plc4x_message_length_in_bits(_message) / 8;
}

uint16_t plc4c_plc4x_read_write_plc4x_message_length_in_bits(plc4c_plc4x_read_write_plc4x_message* _message) {
  uint16_t lengthInBits = 0;

  // Const Field (version)
  lengthInBits += 8;

  // Implicit Field (packetLength)
  lengthInBits += 16;

  // Simple field (requestId)
  lengthInBits += 16;

  // Discriminator Field (requestType)
  lengthInBits += 8;

  // Depending on the current type, add the length of sub-type elements ...
  switch(_message->_type) {
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_connect_request: {

  // Implicit Field (connectionStringLen)
  lengthInBits += 8;


  // Simple field (connectionString)
  lengthInBits +=  (plc4c_spi_evaluation_helper_str_len(_message->plc4x_connect_request_connection_string)) * (8);

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_connect_response: {

  // Simple field (connectionId)
  lengthInBits += 16;


  // Simple field (responseCode)
  lengthInBits += plc4c_plc4x_read_write_plc4x_response_code_length_in_bits(&_message->plc4x_connect_response_response_code);

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_read_request: {

  // Simple field (connectionId)
  lengthInBits += 16;


  // Implicit Field (numTags)
  lengthInBits += 8;


  // Array field
  if(_message->plc4x_read_request_tags != NULL) {
    plc4c_list_element* curElement = _message->plc4x_read_request_tags->tail;
    while (curElement != NULL) {
      lengthInBits += plc4c_plc4x_read_write_plc4x_tag_request_length_in_bits((plc4c_plc4x_read_write_plc4x_tag_request*) curElement->value);
      curElement = curElement->next;
    }
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_read_response: {

  // Simple field (connectionId)
  lengthInBits += 16;


  // Simple field (responseCode)
  lengthInBits += plc4c_plc4x_read_write_plc4x_response_code_length_in_bits(&_message->plc4x_read_response_response_code);


  // Implicit Field (numTags)
  lengthInBits += 8;


  // Array field
  if(_message->plc4x_read_response_tags != NULL) {
    plc4c_list_element* curElement = _message->plc4x_read_response_tags->tail;
    while (curElement != NULL) {
      lengthInBits += plc4c_plc4x_read_write_plc4x_tag_value_response_length_in_bits((plc4c_plc4x_read_write_plc4x_tag_value_response*) curElement->value);
      curElement = curElement->next;
    }
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_write_request: {

  // Simple field (connectionId)
  lengthInBits += 16;


  // Implicit Field (numTags)
  lengthInBits += 8;


  // Array field
  if(_message->plc4x_write_request_tags != NULL) {
    plc4c_list_element* curElement = _message->plc4x_write_request_tags->tail;
    while (curElement != NULL) {
      lengthInBits += plc4c_plc4x_read_write_plc4x_tag_value_request_length_in_bits((plc4c_plc4x_read_write_plc4x_tag_value_request*) curElement->value);
      curElement = curElement->next;
    }
  }

      break;
    }
    case plc4c_plc4x_read_write_plc4x_message_type_plc4c_plc4x_read_write_plc4x_write_response: {

  // Simple field (connectionId)
  lengthInBits += 16;


  // Simple field (responseCode)
  lengthInBits += plc4c_plc4x_read_write_plc4x_response_code_length_in_bits(&_message->plc4x_write_response_response_code);


  // Implicit Field (numTags)
  lengthInBits += 8;


  // Array field
  if(_message->plc4x_write_response_tags != NULL) {
    plc4c_list_element* curElement = _message->plc4x_write_response_tags->tail;
    while (curElement != NULL) {
      lengthInBits += plc4c_plc4x_read_write_plc4x_tag_response_length_in_bits((plc4c_plc4x_read_write_plc4x_tag_response*) curElement->value);
      curElement = curElement->next;
    }
  }

      break;
    }
  }

  return lengthInBits;
}

