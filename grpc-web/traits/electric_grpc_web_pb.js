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


var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js')

var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_change_pb = require('../types/change_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./electric_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.ElectricApiClient =
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
proto.smartcore.traits.ElectricApiPromiseClient =
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
 *   !proto.smartcore.traits.GetDemandRequest,
 *   !proto.smartcore.traits.ElectricDemand>}
 */
const methodDescriptor_ElectricApi_GetDemand = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/GetDemand',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetDemandRequest,
  proto.smartcore.traits.ElectricDemand,
  /**
   * @param {!proto.smartcore.traits.GetDemandRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ElectricDemand.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetDemandRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ElectricDemand)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ElectricDemand>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.getDemand =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/GetDemand',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_GetDemand,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetDemandRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ElectricDemand>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.getDemand =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/GetDemand',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_GetDemand);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullDemandRequest,
 *   !proto.smartcore.traits.PullDemandResponse>}
 */
const methodDescriptor_ElectricApi_PullDemand = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/PullDemand',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullDemandRequest,
  proto.smartcore.traits.PullDemandResponse,
  /**
   * @param {!proto.smartcore.traits.PullDemandRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullDemandResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullDemandRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullDemandResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.pullDemand =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ElectricApi/PullDemand',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_PullDemand);
};


/**
 * @param {!proto.smartcore.traits.PullDemandRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullDemandResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.pullDemand =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ElectricApi/PullDemand',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_PullDemand);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.GetActiveModeRequest,
 *   !proto.smartcore.traits.ElectricMode>}
 */
const methodDescriptor_ElectricApi_GetActiveMode = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/GetActiveMode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetActiveModeRequest,
  proto.smartcore.traits.ElectricMode,
  /**
   * @param {!proto.smartcore.traits.GetActiveModeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ElectricMode.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetActiveModeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ElectricMode)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ElectricMode>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.getActiveMode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/GetActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_GetActiveMode,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetActiveModeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ElectricMode>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.getActiveMode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/GetActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_GetActiveMode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateActiveModeRequest,
 *   !proto.smartcore.traits.ElectricMode>}
 */
const methodDescriptor_ElectricApi_UpdateActiveMode = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/UpdateActiveMode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateActiveModeRequest,
  proto.smartcore.traits.ElectricMode,
  /**
   * @param {!proto.smartcore.traits.UpdateActiveModeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ElectricMode.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateActiveModeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ElectricMode)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ElectricMode>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.updateActiveMode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/UpdateActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_UpdateActiveMode,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateActiveModeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ElectricMode>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.updateActiveMode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/UpdateActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_UpdateActiveMode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ClearActiveModeRequest,
 *   !proto.smartcore.traits.ElectricMode>}
 */
const methodDescriptor_ElectricApi_ClearActiveMode = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/ClearActiveMode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ClearActiveModeRequest,
  proto.smartcore.traits.ElectricMode,
  /**
   * @param {!proto.smartcore.traits.ClearActiveModeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ElectricMode.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ClearActiveModeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ElectricMode)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ElectricMode>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.clearActiveMode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/ClearActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_ClearActiveMode,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ClearActiveModeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ElectricMode>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.clearActiveMode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/ClearActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_ClearActiveMode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullActiveModeRequest,
 *   !proto.smartcore.traits.PullActiveModeResponse>}
 */
const methodDescriptor_ElectricApi_PullActiveMode = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/PullActiveMode',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullActiveModeRequest,
  proto.smartcore.traits.PullActiveModeResponse,
  /**
   * @param {!proto.smartcore.traits.PullActiveModeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullActiveModeResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullActiveModeRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullActiveModeResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.pullActiveMode =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ElectricApi/PullActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_PullActiveMode);
};


/**
 * @param {!proto.smartcore.traits.PullActiveModeRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullActiveModeResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.pullActiveMode =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ElectricApi/PullActiveMode',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_PullActiveMode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ListModesRequest,
 *   !proto.smartcore.traits.ListModesResponse>}
 */
const methodDescriptor_ElectricApi_ListModes = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/ListModes',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ListModesRequest,
  proto.smartcore.traits.ListModesResponse,
  /**
   * @param {!proto.smartcore.traits.ListModesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ListModesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ListModesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.ListModesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ListModesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.listModes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/ListModes',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_ListModes,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ListModesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ListModesResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.listModes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.ElectricApi/ListModes',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_ListModes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullModesRequest,
 *   !proto.smartcore.traits.PullModesResponse>}
 */
const methodDescriptor_ElectricApi_PullModes = new grpc.web.MethodDescriptor(
  '/smartcore.traits.ElectricApi/PullModes',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullModesRequest,
  proto.smartcore.traits.PullModesResponse,
  /**
   * @param {!proto.smartcore.traits.PullModesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullModesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullModesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullModesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiClient.prototype.pullModes =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ElectricApi/PullModes',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_PullModes);
};


/**
 * @param {!proto.smartcore.traits.PullModesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullModesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.ElectricApiPromiseClient.prototype.pullModes =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.ElectricApi/PullModes',
      request,
      metadata || {},
      methodDescriptor_ElectricApi_PullModes);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.ElectricInfoClient =
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
proto.smartcore.traits.ElectricInfoPromiseClient =
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

