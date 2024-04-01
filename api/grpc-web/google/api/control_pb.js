// source: google/api/control.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_api_policy_pb = require('../../google/api/policy_pb.js');
goog.object.extend(proto, google_api_policy_pb);
goog.exportSymbol('proto.google.api.Control', null, global);
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
proto.google.api.Control = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.google.api.Control.repeatedFields_, null);
};
goog.inherits(proto.google.api.Control, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.google.api.Control.displayName = 'proto.google.api.Control';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.google.api.Control.repeatedFields_ = [4];



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
proto.google.api.Control.prototype.toObject = function(opt_includeInstance) {
  return proto.google.api.Control.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.google.api.Control} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.Control.toObject = function(includeInstance, msg) {
  var f, obj = {
    environment: jspb.Message.getFieldWithDefault(msg, 1, ""),
    methodPoliciesList: jspb.Message.toObjectList(msg.getMethodPoliciesList(),
    google_api_policy_pb.MethodPolicy.toObject, includeInstance)
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
 * @return {!proto.google.api.Control}
 */
proto.google.api.Control.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.google.api.Control;
  return proto.google.api.Control.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.google.api.Control} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.google.api.Control}
 */
proto.google.api.Control.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEnvironment(value);
      break;
    case 4:
      var value = new google_api_policy_pb.MethodPolicy;
      reader.readMessage(value,google_api_policy_pb.MethodPolicy.deserializeBinaryFromReader);
      msg.addMethodPolicies(value);
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
proto.google.api.Control.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.google.api.Control.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.google.api.Control} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.Control.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEnvironment();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getMethodPoliciesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      google_api_policy_pb.MethodPolicy.serializeBinaryToWriter
    );
  }
};


/**
 * optional string environment = 1;
 * @return {string}
 */
proto.google.api.Control.prototype.getEnvironment = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.google.api.Control} returns this
 */
proto.google.api.Control.prototype.setEnvironment = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * repeated MethodPolicy method_policies = 4;
 * @return {!Array<!proto.google.api.MethodPolicy>}
 */
proto.google.api.Control.prototype.getMethodPoliciesList = function() {
  return /** @type{!Array<!proto.google.api.MethodPolicy>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_api_policy_pb.MethodPolicy, 4));
};


/**
 * @param {!Array<!proto.google.api.MethodPolicy>} value
 * @return {!proto.google.api.Control} returns this
*/
proto.google.api.Control.prototype.setMethodPoliciesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.google.api.MethodPolicy=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.api.MethodPolicy}
 */
proto.google.api.Control.prototype.addMethodPolicies = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.google.api.MethodPolicy, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.google.api.Control} returns this
 */
proto.google.api.Control.prototype.clearMethodPoliciesList = function() {
  return this.setMethodPoliciesList([]);
};


goog.object.extend(exports, proto.google.api);
