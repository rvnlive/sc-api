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
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./lock_unlock_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.LockUnlockApiClient =
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
proto.smartcore.traits.LockUnlockApiPromiseClient =
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
 *   !proto.smartcore.traits.GetLockUnlockRequest,
 *   !proto.smartcore.traits.LockUnlock>}
 */
const methodDescriptor_LockUnlockApi_GetLockUnlock = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LockUnlockApi/GetLockUnlock',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetLockUnlockRequest,
  proto.smartcore.traits.LockUnlock,
  /**
   * @param {!proto.smartcore.traits.GetLockUnlockRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.LockUnlock.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetLockUnlockRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.LockUnlock)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.LockUnlock>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LockUnlockApiClient.prototype.getLockUnlock =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.LockUnlockApi/GetLockUnlock',
      request,
      metadata || {},
      methodDescriptor_LockUnlockApi_GetLockUnlock,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetLockUnlockRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.LockUnlock>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.LockUnlockApiPromiseClient.prototype.getLockUnlock =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.LockUnlockApi/GetLockUnlock',
      request,
      metadata || {},
      methodDescriptor_LockUnlockApi_GetLockUnlock);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateLockUnlockRequest,
 *   !proto.smartcore.traits.LockUnlock>}
 */
const methodDescriptor_LockUnlockApi_UpdateLockUnlock = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LockUnlockApi/UpdateLockUnlock',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateLockUnlockRequest,
  proto.smartcore.traits.LockUnlock,
  /**
   * @param {!proto.smartcore.traits.UpdateLockUnlockRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.LockUnlock.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateLockUnlockRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.LockUnlock)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.LockUnlock>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LockUnlockApiClient.prototype.updateLockUnlock =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.LockUnlockApi/UpdateLockUnlock',
      request,
      metadata || {},
      methodDescriptor_LockUnlockApi_UpdateLockUnlock,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateLockUnlockRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.LockUnlock>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.LockUnlockApiPromiseClient.prototype.updateLockUnlock =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.LockUnlockApi/UpdateLockUnlock',
      request,
      metadata || {},
      methodDescriptor_LockUnlockApi_UpdateLockUnlock);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullLockUnlockRequest,
 *   !proto.smartcore.traits.PullLockUnlockResponse>}
 */
const methodDescriptor_LockUnlockApi_PullLockUnlock = new grpc.web.MethodDescriptor(
  '/smartcore.traits.LockUnlockApi/PullLockUnlock',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullLockUnlockRequest,
  proto.smartcore.traits.PullLockUnlockResponse,
  /**
   * @param {!proto.smartcore.traits.PullLockUnlockRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullLockUnlockResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullLockUnlockRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullLockUnlockResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LockUnlockApiClient.prototype.pullLockUnlock =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.LockUnlockApi/PullLockUnlock',
      request,
      metadata || {},
      methodDescriptor_LockUnlockApi_PullLockUnlock);
};


/**
 * @param {!proto.smartcore.traits.PullLockUnlockRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullLockUnlockResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.LockUnlockApiPromiseClient.prototype.pullLockUnlock =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.LockUnlockApi/PullLockUnlock',
      request,
      metadata || {},
      methodDescriptor_LockUnlockApi_PullLockUnlock);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.LockUnlockInfoClient =
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
proto.smartcore.traits.LockUnlockInfoPromiseClient =
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


module.exports = proto.smartcore.traits;

