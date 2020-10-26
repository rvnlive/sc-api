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


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./occupancy_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.OccupancyApiClient =
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
proto.smartcore.traits.OccupancyApiPromiseClient =
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
 *   !proto.smartcore.traits.GetOccupancyRequest,
 *   !proto.smartcore.traits.Occupancy>}
 */
const methodDescriptor_OccupancyApi_GetOccupancy = new grpc.web.MethodDescriptor(
  '/smartcore.traits.OccupancyApi/GetOccupancy',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetOccupancyRequest,
  proto.smartcore.traits.Occupancy,
  /**
   * @param {!proto.smartcore.traits.GetOccupancyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Occupancy.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.GetOccupancyRequest,
 *   !proto.smartcore.traits.Occupancy>}
 */
const methodInfo_OccupancyApi_GetOccupancy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.Occupancy,
  /**
   * @param {!proto.smartcore.traits.GetOccupancyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Occupancy.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetOccupancyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.Occupancy)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Occupancy>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OccupancyApiClient.prototype.getOccupancy =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.OccupancyApi/GetOccupancy',
      request,
      metadata || {},
      methodDescriptor_OccupancyApi_GetOccupancy,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetOccupancyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Occupancy>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.OccupancyApiPromiseClient.prototype.getOccupancy =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.OccupancyApi/GetOccupancy',
      request,
      metadata || {},
      methodDescriptor_OccupancyApi_GetOccupancy);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullOccupancyRequest,
 *   !proto.smartcore.traits.PullOccupancyResponse>}
 */
const methodDescriptor_OccupancyApi_PullOccupancy = new grpc.web.MethodDescriptor(
  '/smartcore.traits.OccupancyApi/PullOccupancy',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullOccupancyRequest,
  proto.smartcore.traits.PullOccupancyResponse,
  /**
   * @param {!proto.smartcore.traits.PullOccupancyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullOccupancyResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.PullOccupancyRequest,
 *   !proto.smartcore.traits.PullOccupancyResponse>}
 */
const methodInfo_OccupancyApi_PullOccupancy = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.PullOccupancyResponse,
  /**
   * @param {!proto.smartcore.traits.PullOccupancyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullOccupancyResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullOccupancyRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullOccupancyResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OccupancyApiClient.prototype.pullOccupancy =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.OccupancyApi/PullOccupancy',
      request,
      metadata || {},
      methodDescriptor_OccupancyApi_PullOccupancy);
};


/**
 * @param {!proto.smartcore.traits.PullOccupancyRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullOccupancyResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OccupancyApiPromiseClient.prototype.pullOccupancy =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.OccupancyApi/PullOccupancy',
      request,
      metadata || {},
      methodDescriptor_OccupancyApi_PullOccupancy);
};


module.exports = proto.smartcore.traits;

