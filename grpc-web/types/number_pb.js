// source: types/number.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

var types_tween_pb = require('../types/tween_pb.js');
goog.object.extend(proto, types_tween_pb);
goog.exportSymbol('proto.smartcore.types.FloatAttributes', null, global);
goog.exportSymbol('proto.smartcore.types.FloatBounds', null, global);
goog.exportSymbol('proto.smartcore.types.Int32Attributes', null, global);
goog.exportSymbol('proto.smartcore.types.Int32Bounds', null, global);
goog.exportSymbol('proto.smartcore.types.InvalidNumberBehaviour', null, global);
goog.exportSymbol('proto.smartcore.types.NumberCapping', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.types.NumberCapping = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.types.NumberCapping, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.types.NumberCapping.displayName = 'proto.smartcore.types.NumberCapping';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.types.Int32Bounds = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.types.Int32Bounds, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.types.Int32Bounds.displayName = 'proto.smartcore.types.Int32Bounds';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.types.Int32Attributes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.types.Int32Attributes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.types.Int32Attributes.displayName = 'proto.smartcore.types.Int32Attributes';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.types.FloatBounds = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.types.FloatBounds, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.types.FloatBounds.displayName = 'proto.smartcore.types.FloatBounds';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.types.FloatAttributes = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.types.FloatAttributes, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.types.FloatAttributes.displayName = 'proto.smartcore.types.FloatAttributes';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.types.NumberCapping.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.types.NumberCapping.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.types.NumberCapping} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.NumberCapping.toObject = function(includeInstance, msg) {
  var f, obj = {
    min: jspb.Message.getFieldWithDefault(msg, 1, 0),
    step: jspb.Message.getFieldWithDefault(msg, 2, 0),
    max: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.types.NumberCapping}
 */
proto.smartcore.types.NumberCapping.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.types.NumberCapping;
  return proto.smartcore.types.NumberCapping.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.types.NumberCapping} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.types.NumberCapping}
 */
proto.smartcore.types.NumberCapping.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.smartcore.types.InvalidNumberBehaviour} */ (reader.readEnum());
      msg.setMin(value);
      break;
    case 2:
      var value = /** @type {!proto.smartcore.types.InvalidNumberBehaviour} */ (reader.readEnum());
      msg.setStep(value);
      break;
    case 3:
      var value = /** @type {!proto.smartcore.types.InvalidNumberBehaviour} */ (reader.readEnum());
      msg.setMax(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.types.NumberCapping.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.types.NumberCapping.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.types.NumberCapping} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.NumberCapping.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMin();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getStep();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getMax();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
};


/**
 * optional InvalidNumberBehaviour min = 1;
 * @return {!proto.smartcore.types.InvalidNumberBehaviour}
 */
proto.smartcore.types.NumberCapping.prototype.getMin = function() {
  return /** @type {!proto.smartcore.types.InvalidNumberBehaviour} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.smartcore.types.InvalidNumberBehaviour} value
 * @return {!proto.smartcore.types.NumberCapping} returns this
 */
proto.smartcore.types.NumberCapping.prototype.setMin = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional InvalidNumberBehaviour step = 2;
 * @return {!proto.smartcore.types.InvalidNumberBehaviour}
 */
proto.smartcore.types.NumberCapping.prototype.getStep = function() {
  return /** @type {!proto.smartcore.types.InvalidNumberBehaviour} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.smartcore.types.InvalidNumberBehaviour} value
 * @return {!proto.smartcore.types.NumberCapping} returns this
 */
proto.smartcore.types.NumberCapping.prototype.setStep = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional InvalidNumberBehaviour max = 3;
 * @return {!proto.smartcore.types.InvalidNumberBehaviour}
 */
proto.smartcore.types.NumberCapping.prototype.getMax = function() {
  return /** @type {!proto.smartcore.types.InvalidNumberBehaviour} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.smartcore.types.InvalidNumberBehaviour} value
 * @return {!proto.smartcore.types.NumberCapping} returns this
 */
proto.smartcore.types.NumberCapping.prototype.setMax = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.types.Int32Bounds.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.types.Int32Bounds.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.types.Int32Bounds} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.Int32Bounds.toObject = function(includeInstance, msg) {
  var f, obj = {
    min: jspb.Message.getFieldWithDefault(msg, 1, 0),
    max: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.types.Int32Bounds}
 */
proto.smartcore.types.Int32Bounds.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.types.Int32Bounds;
  return proto.smartcore.types.Int32Bounds.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.types.Int32Bounds} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.types.Int32Bounds}
 */
proto.smartcore.types.Int32Bounds.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMin(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setMax(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.types.Int32Bounds.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.types.Int32Bounds.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.types.Int32Bounds} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.Int32Bounds.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeInt32(
      2,
      f
    );
  }
};


/**
 * optional int32 min = 1;
 * @return {number}
 */
proto.smartcore.types.Int32Bounds.prototype.getMin = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.smartcore.types.Int32Bounds} returns this
 */
proto.smartcore.types.Int32Bounds.prototype.setMin = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.smartcore.types.Int32Bounds} returns this
 */
proto.smartcore.types.Int32Bounds.prototype.clearMin = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.Int32Bounds.prototype.hasMin = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int32 max = 2;
 * @return {number}
 */
proto.smartcore.types.Int32Bounds.prototype.getMax = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.smartcore.types.Int32Bounds} returns this
 */
proto.smartcore.types.Int32Bounds.prototype.setMax = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.smartcore.types.Int32Bounds} returns this
 */
proto.smartcore.types.Int32Bounds.prototype.clearMax = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.Int32Bounds.prototype.hasMax = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.types.Int32Attributes.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.types.Int32Attributes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.types.Int32Attributes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.Int32Attributes.toObject = function(includeInstance, msg) {
  var f, obj = {
    bounds: (f = msg.getBounds()) && proto.smartcore.types.Int32Bounds.toObject(includeInstance, f),
    step: jspb.Message.getFieldWithDefault(msg, 3, 0),
    supportsDelta: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    rampSupport: jspb.Message.getFieldWithDefault(msg, 5, 0),
    defaultCapping: (f = msg.getDefaultCapping()) && proto.smartcore.types.NumberCapping.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.types.Int32Attributes}
 */
proto.smartcore.types.Int32Attributes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.types.Int32Attributes;
  return proto.smartcore.types.Int32Attributes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.types.Int32Attributes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.types.Int32Attributes}
 */
proto.smartcore.types.Int32Attributes.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.smartcore.types.Int32Bounds;
      reader.readMessage(value,proto.smartcore.types.Int32Bounds.deserializeBinaryFromReader);
      msg.setBounds(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setStep(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSupportsDelta(value);
      break;
    case 5:
      var value = /** @type {!proto.smartcore.types.TweenSupport} */ (reader.readEnum());
      msg.setRampSupport(value);
      break;
    case 6:
      var value = new proto.smartcore.types.NumberCapping;
      reader.readMessage(value,proto.smartcore.types.NumberCapping.deserializeBinaryFromReader);
      msg.setDefaultCapping(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.types.Int32Attributes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.types.Int32Attributes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.types.Int32Attributes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.Int32Attributes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBounds();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.smartcore.types.Int32Bounds.serializeBinaryToWriter
    );
  }
  f = message.getStep();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getSupportsDelta();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getRampSupport();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
  f = message.getDefaultCapping();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.smartcore.types.NumberCapping.serializeBinaryToWriter
    );
  }
};


/**
 * optional Int32Bounds bounds = 1;
 * @return {?proto.smartcore.types.Int32Bounds}
 */
proto.smartcore.types.Int32Attributes.prototype.getBounds = function() {
  return /** @type{?proto.smartcore.types.Int32Bounds} */ (
    jspb.Message.getWrapperField(this, proto.smartcore.types.Int32Bounds, 1));
};


/**
 * @param {?proto.smartcore.types.Int32Bounds|undefined} value
 * @return {!proto.smartcore.types.Int32Attributes} returns this
*/
proto.smartcore.types.Int32Attributes.prototype.setBounds = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.smartcore.types.Int32Attributes} returns this
 */
proto.smartcore.types.Int32Attributes.prototype.clearBounds = function() {
  return this.setBounds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.Int32Attributes.prototype.hasBounds = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int32 step = 3;
 * @return {number}
 */
proto.smartcore.types.Int32Attributes.prototype.getStep = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.smartcore.types.Int32Attributes} returns this
 */
proto.smartcore.types.Int32Attributes.prototype.setStep = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional bool supports_delta = 4;
 * @return {boolean}
 */
proto.smartcore.types.Int32Attributes.prototype.getSupportsDelta = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.smartcore.types.Int32Attributes} returns this
 */
proto.smartcore.types.Int32Attributes.prototype.setSupportsDelta = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional TweenSupport ramp_support = 5;
 * @return {!proto.smartcore.types.TweenSupport}
 */
proto.smartcore.types.Int32Attributes.prototype.getRampSupport = function() {
  return /** @type {!proto.smartcore.types.TweenSupport} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.smartcore.types.TweenSupport} value
 * @return {!proto.smartcore.types.Int32Attributes} returns this
 */
proto.smartcore.types.Int32Attributes.prototype.setRampSupport = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
};


/**
 * optional NumberCapping default_capping = 6;
 * @return {?proto.smartcore.types.NumberCapping}
 */
proto.smartcore.types.Int32Attributes.prototype.getDefaultCapping = function() {
  return /** @type{?proto.smartcore.types.NumberCapping} */ (
    jspb.Message.getWrapperField(this, proto.smartcore.types.NumberCapping, 6));
};


/**
 * @param {?proto.smartcore.types.NumberCapping|undefined} value
 * @return {!proto.smartcore.types.Int32Attributes} returns this
*/
proto.smartcore.types.Int32Attributes.prototype.setDefaultCapping = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.smartcore.types.Int32Attributes} returns this
 */
proto.smartcore.types.Int32Attributes.prototype.clearDefaultCapping = function() {
  return this.setDefaultCapping(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.Int32Attributes.prototype.hasDefaultCapping = function() {
  return jspb.Message.getField(this, 6) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.types.FloatBounds.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.types.FloatBounds.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.types.FloatBounds} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.FloatBounds.toObject = function(includeInstance, msg) {
  var f, obj = {
    min: jspb.Message.getFloatingPointFieldWithDefault(msg, 1, 0.0),
    max: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.types.FloatBounds}
 */
proto.smartcore.types.FloatBounds.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.types.FloatBounds;
  return proto.smartcore.types.FloatBounds.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.types.FloatBounds} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.types.FloatBounds}
 */
proto.smartcore.types.FloatBounds.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setMin(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setMax(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.types.FloatBounds.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.types.FloatBounds.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.types.FloatBounds} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.FloatBounds.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeFloat(
      2,
      f
    );
  }
};


/**
 * optional float min = 1;
 * @return {number}
 */
proto.smartcore.types.FloatBounds.prototype.getMin = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.smartcore.types.FloatBounds} returns this
 */
proto.smartcore.types.FloatBounds.prototype.setMin = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.smartcore.types.FloatBounds} returns this
 */
proto.smartcore.types.FloatBounds.prototype.clearMin = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.FloatBounds.prototype.hasMin = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional float max = 2;
 * @return {number}
 */
proto.smartcore.types.FloatBounds.prototype.getMax = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.smartcore.types.FloatBounds} returns this
 */
proto.smartcore.types.FloatBounds.prototype.setMax = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.smartcore.types.FloatBounds} returns this
 */
proto.smartcore.types.FloatBounds.prototype.clearMax = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.FloatBounds.prototype.hasMax = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.types.FloatAttributes.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.types.FloatAttributes.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.types.FloatAttributes} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.FloatAttributes.toObject = function(includeInstance, msg) {
  var f, obj = {
    bounds: (f = msg.getBounds()) && proto.smartcore.types.FloatBounds.toObject(includeInstance, f),
    step: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0),
    supportsDelta: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
    rampSupport: jspb.Message.getFieldWithDefault(msg, 5, 0),
    defaultCapping: (f = msg.getDefaultCapping()) && proto.smartcore.types.NumberCapping.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.types.FloatAttributes}
 */
proto.smartcore.types.FloatAttributes.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.types.FloatAttributes;
  return proto.smartcore.types.FloatAttributes.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.types.FloatAttributes} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.types.FloatAttributes}
 */
proto.smartcore.types.FloatAttributes.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.smartcore.types.FloatBounds;
      reader.readMessage(value,proto.smartcore.types.FloatBounds.deserializeBinaryFromReader);
      msg.setBounds(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setStep(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setSupportsDelta(value);
      break;
    case 5:
      var value = /** @type {!proto.smartcore.types.TweenSupport} */ (reader.readEnum());
      msg.setRampSupport(value);
      break;
    case 6:
      var value = new proto.smartcore.types.NumberCapping;
      reader.readMessage(value,proto.smartcore.types.NumberCapping.deserializeBinaryFromReader);
      msg.setDefaultCapping(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.types.FloatAttributes.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.types.FloatAttributes.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.types.FloatAttributes} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.types.FloatAttributes.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBounds();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.smartcore.types.FloatBounds.serializeBinaryToWriter
    );
  }
  f = message.getStep();
  if (f !== 0.0) {
    writer.writeFloat(
      3,
      f
    );
  }
  f = message.getSupportsDelta();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getRampSupport();
  if (f !== 0.0) {
    writer.writeEnum(
      5,
      f
    );
  }
  f = message.getDefaultCapping();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.smartcore.types.NumberCapping.serializeBinaryToWriter
    );
  }
};


/**
 * optional FloatBounds bounds = 1;
 * @return {?proto.smartcore.types.FloatBounds}
 */
proto.smartcore.types.FloatAttributes.prototype.getBounds = function() {
  return /** @type{?proto.smartcore.types.FloatBounds} */ (
    jspb.Message.getWrapperField(this, proto.smartcore.types.FloatBounds, 1));
};


/**
 * @param {?proto.smartcore.types.FloatBounds|undefined} value
 * @return {!proto.smartcore.types.FloatAttributes} returns this
*/
proto.smartcore.types.FloatAttributes.prototype.setBounds = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.smartcore.types.FloatAttributes} returns this
 */
proto.smartcore.types.FloatAttributes.prototype.clearBounds = function() {
  return this.setBounds(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.FloatAttributes.prototype.hasBounds = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional float step = 3;
 * @return {number}
 */
proto.smartcore.types.FloatAttributes.prototype.getStep = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.smartcore.types.FloatAttributes} returns this
 */
proto.smartcore.types.FloatAttributes.prototype.setStep = function(value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};


/**
 * optional bool supports_delta = 4;
 * @return {boolean}
 */
proto.smartcore.types.FloatAttributes.prototype.getSupportsDelta = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.smartcore.types.FloatAttributes} returns this
 */
proto.smartcore.types.FloatAttributes.prototype.setSupportsDelta = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional TweenSupport ramp_support = 5;
 * @return {!proto.smartcore.types.TweenSupport}
 */
proto.smartcore.types.FloatAttributes.prototype.getRampSupport = function() {
  return /** @type {!proto.smartcore.types.TweenSupport} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {!proto.smartcore.types.TweenSupport} value
 * @return {!proto.smartcore.types.FloatAttributes} returns this
 */
proto.smartcore.types.FloatAttributes.prototype.setRampSupport = function(value) {
  return jspb.Message.setProto3EnumField(this, 5, value);
};


/**
 * optional NumberCapping default_capping = 6;
 * @return {?proto.smartcore.types.NumberCapping}
 */
proto.smartcore.types.FloatAttributes.prototype.getDefaultCapping = function() {
  return /** @type{?proto.smartcore.types.NumberCapping} */ (
    jspb.Message.getWrapperField(this, proto.smartcore.types.NumberCapping, 6));
};


/**
 * @param {?proto.smartcore.types.NumberCapping|undefined} value
 * @return {!proto.smartcore.types.FloatAttributes} returns this
*/
proto.smartcore.types.FloatAttributes.prototype.setDefaultCapping = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.smartcore.types.FloatAttributes} returns this
 */
proto.smartcore.types.FloatAttributes.prototype.clearDefaultCapping = function() {
  return this.setDefaultCapping(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.types.FloatAttributes.prototype.hasDefaultCapping = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * @enum {number}
 */
proto.smartcore.types.InvalidNumberBehaviour = {
  INVALID_NUMBER_BEHAVIOUR_UNSPECIFIED: 0,
  RESTRICT: 1,
  ERROR: 2,
  ALLOW: 3
};

goog.object.extend(exports, proto.smartcore.types);
