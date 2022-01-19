/**
 * @fileoverview gRPC-Web generated client stub for smartcore.traits
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_info_pb = require('../types/info_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./emergency_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.EmergencyApiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.EmergencyApiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.GetEmergencyRequest,
 *   !proto.smartcore.traits.Emergency>}
 */
const methodDescriptor_EmergencyApi_GetEmergency = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EmergencyApi/GetEmergency',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetEmergencyRequest,
  proto.smartcore.traits.Emergency,
  /**
   * @param {!proto.smartcore.traits.GetEmergencyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Emergency.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetEmergencyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Emergency)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Emergency>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EmergencyApiClient.prototype.getEmergency =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.EmergencyApi/GetEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyApi_GetEmergency,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetEmergencyRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Emergency>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.EmergencyApiPromiseClient.prototype.getEmergency =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.EmergencyApi/GetEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyApi_GetEmergency);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateEmergencyRequest,
 *   !proto.smartcore.traits.Emergency>}
 */
const methodDescriptor_EmergencyApi_UpdateEmergency = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EmergencyApi/UpdateEmergency',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateEmergencyRequest,
  proto.smartcore.traits.Emergency,
  /**
   * @param {!proto.smartcore.traits.UpdateEmergencyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Emergency.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateEmergencyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Emergency)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Emergency>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EmergencyApiClient.prototype.updateEmergency =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.EmergencyApi/UpdateEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyApi_UpdateEmergency,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateEmergencyRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Emergency>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.EmergencyApiPromiseClient.prototype.updateEmergency =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.EmergencyApi/UpdateEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyApi_UpdateEmergency);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullEmergencyRequest,
 *   !proto.smartcore.traits.PullEmergencyResponse>}
 */
const methodDescriptor_EmergencyApi_PullEmergency = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EmergencyApi/PullEmergency',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullEmergencyRequest,
  proto.smartcore.traits.PullEmergencyResponse,
  /**
   * @param {!proto.smartcore.traits.PullEmergencyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullEmergencyResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullEmergencyRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullEmergencyResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EmergencyApiClient.prototype.pullEmergency =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.EmergencyApi/PullEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyApi_PullEmergency);
};


/**
 * @param {!proto.smartcore.traits.PullEmergencyRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullEmergencyResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EmergencyApiPromiseClient.prototype.pullEmergency =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.EmergencyApi/PullEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyApi_PullEmergency);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.EmergencyInfoClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.EmergencyInfoPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.DescribeEmergencyRequest,
 *   !proto.smartcore.traits.EmergencySupport>}
 */
const methodDescriptor_EmergencyInfo_DescribeEmergency = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EmergencyInfo/DescribeEmergency',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DescribeEmergencyRequest,
  proto.smartcore.traits.EmergencySupport,
  /**
   * @param {!proto.smartcore.traits.DescribeEmergencyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.EmergencySupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DescribeEmergencyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.EmergencySupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.EmergencySupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EmergencyInfoClient.prototype.describeEmergency =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.EmergencyInfo/DescribeEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyInfo_DescribeEmergency,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DescribeEmergencyRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.EmergencySupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.EmergencyInfoPromiseClient.prototype.describeEmergency =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.EmergencyInfo/DescribeEmergency',
      request,
      metadata || {},
      methodDescriptor_EmergencyInfo_DescribeEmergency);
};


module.exports = proto.smartcore.traits;

