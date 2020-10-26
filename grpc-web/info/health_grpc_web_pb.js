/**
 * @fileoverview gRPC-Web generated client stub for smartcore.info
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_connection_pb = require('../types/connection_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.info = require('./health_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.info.HealthClient =
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
proto.smartcore.info.HealthPromiseClient =
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
 *   !proto.smartcore.info.GetHealthStateRequest,
 *   !proto.smartcore.info.HealthState>}
 */
const methodDescriptor_Health_GetHealthState = new grpc.web.MethodDescriptor(
  '/smartcore.info.Health/GetHealthState',
  grpc.web.MethodType.UNARY,
  proto.smartcore.info.GetHealthStateRequest,
  proto.smartcore.info.HealthState,
  /**
   * @param {!proto.smartcore.info.GetHealthStateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.info.HealthState.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.info.GetHealthStateRequest,
 *   !proto.smartcore.info.HealthState>}
 */
const methodInfo_Health_GetHealthState = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.info.HealthState,
  /**
   * @param {!proto.smartcore.info.GetHealthStateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.info.HealthState.deserializeBinary
);


/**
 * @param {!proto.smartcore.info.GetHealthStateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.info.HealthState)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.info.HealthState>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.info.HealthClient.prototype.getHealthState =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.info.Health/GetHealthState',
      request,
      metadata || {},
      methodDescriptor_Health_GetHealthState,
      callback);
};


/**
 * @param {!proto.smartcore.info.GetHealthStateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.info.HealthState>}
 *     Promise that resolves to the response
 */
proto.smartcore.info.HealthPromiseClient.prototype.getHealthState =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.info.Health/GetHealthState',
      request,
      metadata || {},
      methodDescriptor_Health_GetHealthState);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.info.PullHealthStatesRequest,
 *   !proto.smartcore.info.PullHealthStatesResponse>}
 */
const methodDescriptor_Health_PullHealthStates = new grpc.web.MethodDescriptor(
  '/smartcore.info.Health/PullHealthStates',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.info.PullHealthStatesRequest,
  proto.smartcore.info.PullHealthStatesResponse,
  /**
   * @param {!proto.smartcore.info.PullHealthStatesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.info.PullHealthStatesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.info.PullHealthStatesRequest,
 *   !proto.smartcore.info.PullHealthStatesResponse>}
 */
const methodInfo_Health_PullHealthStates = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.info.PullHealthStatesResponse,
  /**
   * @param {!proto.smartcore.info.PullHealthStatesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.info.PullHealthStatesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.info.PullHealthStatesRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.info.PullHealthStatesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.info.HealthClient.prototype.pullHealthStates =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.info.Health/PullHealthStates',
      request,
      metadata || {},
      methodDescriptor_Health_PullHealthStates);
};


/**
 * @param {!proto.smartcore.info.PullHealthStatesRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.info.PullHealthStatesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.info.HealthPromiseClient.prototype.pullHealthStates =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.info.Health/PullHealthStates',
      request,
      metadata || {},
      methodDescriptor_Health_PullHealthStates);
};


module.exports = proto.smartcore.info;
