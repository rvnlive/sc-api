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

var types_change_pb = require('../types/change_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./parent_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.ParentApiClient =
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
proto.smartcore.traits.ParentApiPromiseClient =
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
 *   !proto.smartcore.traits.ListChildrenRequest,
 *   !proto.smartcore.traits.ListChildrenResponse>}
 */
const methodDescriptor_ParentApi_ListChildren = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ParentApi/ListChildren',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ListChildrenRequest,
  proto.smartcore.traits.ListChildrenResponse,
  /**
   * @param {!proto.smartcore.traits.ListChildrenRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ListChildrenResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ListChildrenRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ListChildrenResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ListChildrenResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ParentApiClient.prototype.listChildren =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.ParentApi/ListChildren',
      request,
      metadata || {},
      methodDescriptor_ParentApi_ListChildren,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ListChildrenRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ListChildrenResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.ParentApiPromiseClient.prototype.listChildren =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.ParentApi/ListChildren',
      request,
      metadata || {},
      methodDescriptor_ParentApi_ListChildren);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullChildrenRequest,
 *   !proto.smartcore.traits.PullChildrenResponse>}
 */
const methodDescriptor_ParentApi_PullChildren = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ParentApi/PullChildren',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullChildrenRequest,
  proto.smartcore.traits.PullChildrenResponse,
  /**
   * @param {!proto.smartcore.traits.PullChildrenRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullChildrenResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullChildrenRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullChildrenResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ParentApiClient.prototype.pullChildren =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ParentApi/PullChildren',
      request,
      metadata || {},
      methodDescriptor_ParentApi_PullChildren);
};


/**
 * @param {!proto.smartcore.traits.PullChildrenRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullChildrenResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ParentApiPromiseClient.prototype.pullChildren =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ParentApi/PullChildren',
      request,
      metadata || {},
      methodDescriptor_ParentApi_PullChildren);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.ParentInfoClient =
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
proto.smartcore.traits.ParentInfoPromiseClient =
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

