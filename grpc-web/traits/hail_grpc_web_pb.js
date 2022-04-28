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
proto.smartcore.traits = require('./hail_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.HailApiClient =
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
proto.smartcore.traits.HailApiPromiseClient =
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
 *   !proto.smartcore.traits.CreateHailRequest,
 *   !proto.smartcore.traits.Hail>}
 */
const methodDescriptor_HailApi_CreateHail = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/CreateHail',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.CreateHailRequest,
  proto.smartcore.traits.Hail,
  /**
   * @param {!proto.smartcore.traits.CreateHailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Hail.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.CreateHailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Hail)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Hail>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.createHail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.HailApi/CreateHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_CreateHail,
      callback);
};


/**
 * @param {!proto.smartcore.traits.CreateHailRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Hail>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.createHail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.HailApi/CreateHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_CreateHail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.GetHailRequest,
 *   !proto.smartcore.traits.Hail>}
 */
const methodDescriptor_HailApi_GetHail = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/GetHail',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetHailRequest,
  proto.smartcore.traits.Hail,
  /**
   * @param {!proto.smartcore.traits.GetHailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Hail.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetHailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Hail)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Hail>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.getHail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.HailApi/GetHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_GetHail,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetHailRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Hail>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.getHail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.HailApi/GetHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_GetHail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateHailRequest,
 *   !proto.smartcore.traits.Hail>}
 */
const methodDescriptor_HailApi_UpdateHail = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/UpdateHail',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateHailRequest,
  proto.smartcore.traits.Hail,
  /**
   * @param {!proto.smartcore.traits.UpdateHailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.Hail.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateHailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.Hail)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.Hail>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.updateHail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.HailApi/UpdateHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_UpdateHail,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateHailRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.Hail>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.updateHail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.HailApi/UpdateHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_UpdateHail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.DeleteHailRequest,
 *   !proto.smartcore.traits.DeleteHailResponse>}
 */
const methodDescriptor_HailApi_DeleteHail = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/DeleteHail',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DeleteHailRequest,
  proto.smartcore.traits.DeleteHailResponse,
  /**
   * @param {!proto.smartcore.traits.DeleteHailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.DeleteHailResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DeleteHailRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.DeleteHailResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.DeleteHailResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.deleteHail =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.HailApi/DeleteHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_DeleteHail,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DeleteHailRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.DeleteHailResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.deleteHail =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.HailApi/DeleteHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_DeleteHail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullHailRequest,
 *   !proto.smartcore.traits.PullHailResponse>}
 */
const methodDescriptor_HailApi_PullHail = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/PullHail',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullHailRequest,
  proto.smartcore.traits.PullHailResponse,
  /**
   * @param {!proto.smartcore.traits.PullHailRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullHailResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullHailRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullHailResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.pullHail =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.HailApi/PullHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_PullHail);
};


/**
 * @param {!proto.smartcore.traits.PullHailRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullHailResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.pullHail =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.HailApi/PullHail',
      request,
      metadata || {},
      methodDescriptor_HailApi_PullHail);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ListHailsRequest,
 *   !proto.smartcore.traits.ListHailsResponse>}
 */
const methodDescriptor_HailApi_ListHails = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/ListHails',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ListHailsRequest,
  proto.smartcore.traits.ListHailsResponse,
  /**
   * @param {!proto.smartcore.traits.ListHailsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ListHailsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ListHailsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ListHailsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ListHailsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.listHails =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.HailApi/ListHails',
      request,
      metadata || {},
      methodDescriptor_HailApi_ListHails,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ListHailsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ListHailsResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.listHails =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.HailApi/ListHails',
      request,
      metadata || {},
      methodDescriptor_HailApi_ListHails);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullHailsRequest,
 *   !proto.smartcore.traits.PullHailsResponse>}
 */
const methodDescriptor_HailApi_PullHails = new grpc.web.MethodDescriptor(
  '/smartcore.traits.HailApi/PullHails',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullHailsRequest,
  proto.smartcore.traits.PullHailsResponse,
  /**
   * @param {!proto.smartcore.traits.PullHailsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullHailsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullHailsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullHailsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiClient.prototype.pullHails =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.HailApi/PullHails',
      request,
      metadata || {},
      methodDescriptor_HailApi_PullHails);
};


/**
 * @param {!proto.smartcore.traits.PullHailsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullHailsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.HailApiPromiseClient.prototype.pullHails =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.HailApi/PullHails',
      request,
      metadata || {},
      methodDescriptor_HailApi_PullHails);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.HailInfoClient =
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
proto.smartcore.traits.HailInfoPromiseClient =
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
