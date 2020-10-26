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

var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./count_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.CountApiClient =
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
proto.smartcore.traits.CountApiPromiseClient =
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
 *   !proto.smartcore.traits.GetCountRequest,
 *   !proto.smartcore.traits.Count>}
 */
const methodDescriptor_CountApi_GetCount = new grpc.web.MethodDescriptor(
  '/smartcore.traits.CountApi/GetCount',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetCountRequest,
  proto.smartcore.traits.Count,
  /**
   * @param {!proto.smartcore.traits.GetCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Count.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.GetCountRequest,
 *   !proto.smartcore.traits.Count>}
 */
const methodInfo_CountApi_GetCount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.Count,
  /**
   * @param {!proto.smartcore.traits.GetCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Count.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.Count)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Count>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.CountApiClient.prototype.getCount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.CountApi/GetCount',
      request,
      metadata || {},
      methodDescriptor_CountApi_GetCount,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Count>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.CountApiPromiseClient.prototype.getCount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.CountApi/GetCount',
      request,
      metadata || {},
      methodDescriptor_CountApi_GetCount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ResetCountRequest,
 *   !proto.smartcore.traits.Count>}
 */
const methodDescriptor_CountApi_ResetCount = new grpc.web.MethodDescriptor(
  '/smartcore.traits.CountApi/ResetCount',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ResetCountRequest,
  proto.smartcore.traits.Count,
  /**
   * @param {!proto.smartcore.traits.ResetCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Count.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.ResetCountRequest,
 *   !proto.smartcore.traits.Count>}
 */
const methodInfo_CountApi_ResetCount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.Count,
  /**
   * @param {!proto.smartcore.traits.ResetCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Count.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ResetCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.Count)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Count>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.CountApiClient.prototype.resetCount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.CountApi/ResetCount',
      request,
      metadata || {},
      methodDescriptor_CountApi_ResetCount,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ResetCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Count>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.CountApiPromiseClient.prototype.resetCount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.CountApi/ResetCount',
      request,
      metadata || {},
      methodDescriptor_CountApi_ResetCount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateCountRequest,
 *   !proto.smartcore.traits.Count>}
 */
const methodDescriptor_CountApi_UpdateCount = new grpc.web.MethodDescriptor(
  '/smartcore.traits.CountApi/UpdateCount',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateCountRequest,
  proto.smartcore.traits.Count,
  /**
   * @param {!proto.smartcore.traits.UpdateCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Count.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.UpdateCountRequest,
 *   !proto.smartcore.traits.Count>}
 */
const methodInfo_CountApi_UpdateCount = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.Count,
  /**
   * @param {!proto.smartcore.traits.UpdateCountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Count.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.Count)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Count>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.CountApiClient.prototype.updateCount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.CountApi/UpdateCount',
      request,
      metadata || {},
      methodDescriptor_CountApi_UpdateCount,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateCountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Count>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.CountApiPromiseClient.prototype.updateCount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.CountApi/UpdateCount',
      request,
      metadata || {},
      methodDescriptor_CountApi_UpdateCount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullCountsRequest,
 *   !proto.smartcore.traits.PullCountsResponse>}
 */
const methodDescriptor_CountApi_PullCounts = new grpc.web.MethodDescriptor(
  '/smartcore.traits.CountApi/PullCounts',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullCountsRequest,
  proto.smartcore.traits.PullCountsResponse,
  /**
   * @param {!proto.smartcore.traits.PullCountsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullCountsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.PullCountsRequest,
 *   !proto.smartcore.traits.PullCountsResponse>}
 */
const methodInfo_CountApi_PullCounts = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.PullCountsResponse,
  /**
   * @param {!proto.smartcore.traits.PullCountsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullCountsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullCountsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullCountsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.CountApiClient.prototype.pullCounts =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.CountApi/PullCounts',
      request,
      metadata || {},
      methodDescriptor_CountApi_PullCounts);
};


/**
 * @param {!proto.smartcore.traits.PullCountsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullCountsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.CountApiPromiseClient.prototype.pullCounts =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.CountApi/PullCounts',
      request,
      metadata || {},
      methodDescriptor_CountApi_PullCounts);
};


module.exports = proto.smartcore.traits;
