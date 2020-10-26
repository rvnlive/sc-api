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

var google_protobuf_wrappers_pb = require('google-protobuf/google/protobuf/wrappers_pb.js')

var types_unit_pb = require('../types/unit_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./air_temperature_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.AirTemperatureApiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.AirTemperatureApiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.smartcore.traits.GetAirTemperatureRequest,
 *   !proto.smartcore.traits.AirTemperature>}
 */
const methodDescriptor_AirTemperatureApi_GetAirTemperature = new grpc.web.MethodDescriptor(
  '/smartcore.traits.AirTemperatureApi/GetAirTemperature',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetAirTemperatureRequest,
  proto.smartcore.traits.AirTemperature,
  /**
   * @param {!proto.smartcore.traits.GetAirTemperatureRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.AirTemperature.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.GetAirTemperatureRequest,
 *   !proto.smartcore.traits.AirTemperature>}
 */
const methodInfo_AirTemperatureApi_GetAirTemperature = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.AirTemperature,
  /**
   * @param {!proto.smartcore.traits.GetAirTemperatureRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.AirTemperature.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetAirTemperatureRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.AirTemperature)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.AirTemperature>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.AirTemperatureApiClient.prototype.getAirTemperature =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.AirTemperatureApi/GetAirTemperature',
      request,
      metadata || {},
      methodDescriptor_AirTemperatureApi_GetAirTemperature,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetAirTemperatureRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.AirTemperature>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.AirTemperatureApiPromiseClient.prototype.getAirTemperature =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.AirTemperatureApi/GetAirTemperature',
      request,
      metadata || {},
      methodDescriptor_AirTemperatureApi_GetAirTemperature);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateAirTemperatureRequest,
 *   !proto.smartcore.traits.AirTemperature>}
 */
const methodDescriptor_AirTemperatureApi_UpdateAirTemperature = new grpc.web.MethodDescriptor(
  '/smartcore.traits.AirTemperatureApi/UpdateAirTemperature',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateAirTemperatureRequest,
  proto.smartcore.traits.AirTemperature,
  /**
   * @param {!proto.smartcore.traits.UpdateAirTemperatureRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.AirTemperature.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.UpdateAirTemperatureRequest,
 *   !proto.smartcore.traits.AirTemperature>}
 */
const methodInfo_AirTemperatureApi_UpdateAirTemperature = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.AirTemperature,
  /**
   * @param {!proto.smartcore.traits.UpdateAirTemperatureRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.AirTemperature.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateAirTemperatureRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.AirTemperature)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.AirTemperature>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.AirTemperatureApiClient.prototype.updateAirTemperature =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.AirTemperatureApi/UpdateAirTemperature',
      request,
      metadata || {},
      methodDescriptor_AirTemperatureApi_UpdateAirTemperature,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateAirTemperatureRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.AirTemperature>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.AirTemperatureApiPromiseClient.prototype.updateAirTemperature =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.AirTemperatureApi/UpdateAirTemperature',
      request,
      metadata || {},
      methodDescriptor_AirTemperatureApi_UpdateAirTemperature);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullAirTemperatureRequest,
 *   !proto.smartcore.traits.PullAirTemperatureResponse>}
 */
const methodDescriptor_AirTemperatureApi_PullAirTemperature = new grpc.web.MethodDescriptor(
  '/smartcore.traits.AirTemperatureApi/PullAirTemperature',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullAirTemperatureRequest,
  proto.smartcore.traits.PullAirTemperatureResponse,
  /**
   * @param {!proto.smartcore.traits.PullAirTemperatureRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullAirTemperatureResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.PullAirTemperatureRequest,
 *   !proto.smartcore.traits.PullAirTemperatureResponse>}
 */
const methodInfo_AirTemperatureApi_PullAirTemperature = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.PullAirTemperatureResponse,
  /**
   * @param {!proto.smartcore.traits.PullAirTemperatureRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullAirTemperatureResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullAirTemperatureRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullAirTemperatureResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.AirTemperatureApiClient.prototype.pullAirTemperature =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.AirTemperatureApi/PullAirTemperature',
      request,
      metadata || {},
      methodDescriptor_AirTemperatureApi_PullAirTemperature);
};


/**
 * @param {!proto.smartcore.traits.PullAirTemperatureRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullAirTemperatureResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.AirTemperatureApiPromiseClient.prototype.pullAirTemperature =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.AirTemperatureApi/PullAirTemperature',
      request,
      metadata || {},
      methodDescriptor_AirTemperatureApi_PullAirTemperature);
};


module.exports = proto.smartcore.traits;
