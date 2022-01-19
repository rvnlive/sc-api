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

var types_info_pb = require('../types/info_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./on_off_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.OnOffApiClient =
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
proto.smartcore.traits.OnOffApiPromiseClient =
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
 *   !proto.smartcore.traits.GetOnOffRequest,
 *   !proto.smartcore.traits.OnOff>}
 */
const methodDescriptor_OnOffApi_GetOnOff = new grpc.web.MethodDescriptor(
  '/smartcore.traits.OnOffApi/GetOnOff',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetOnOffRequest,
  proto.smartcore.traits.OnOff,
  /**
   * @param {!proto.smartcore.traits.GetOnOffRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.OnOff.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetOnOffRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.OnOff)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.OnOff>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OnOffApiClient.prototype.getOnOff =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.OnOffApi/GetOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffApi_GetOnOff,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetOnOffRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.OnOff>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.OnOffApiPromiseClient.prototype.getOnOff =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.OnOffApi/GetOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffApi_GetOnOff);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateOnOffRequest,
 *   !proto.smartcore.traits.OnOff>}
 */
const methodDescriptor_OnOffApi_UpdateOnOff = new grpc.web.MethodDescriptor(
  '/smartcore.traits.OnOffApi/UpdateOnOff',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateOnOffRequest,
  proto.smartcore.traits.OnOff,
  /**
   * @param {!proto.smartcore.traits.UpdateOnOffRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.OnOff.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateOnOffRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.OnOff)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.OnOff>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OnOffApiClient.prototype.updateOnOff =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.OnOffApi/UpdateOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffApi_UpdateOnOff,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateOnOffRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.OnOff>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.OnOffApiPromiseClient.prototype.updateOnOff =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.OnOffApi/UpdateOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffApi_UpdateOnOff);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullOnOffRequest,
 *   !proto.smartcore.traits.PullOnOffResponse>}
 */
const methodDescriptor_OnOffApi_PullOnOff = new grpc.web.MethodDescriptor(
  '/smartcore.traits.OnOffApi/PullOnOff',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullOnOffRequest,
  proto.smartcore.traits.PullOnOffResponse,
  /**
   * @param {!proto.smartcore.traits.PullOnOffRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullOnOffResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullOnOffRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullOnOffResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OnOffApiClient.prototype.pullOnOff =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.OnOffApi/PullOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffApi_PullOnOff);
};


/**
 * @param {!proto.smartcore.traits.PullOnOffRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullOnOffResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OnOffApiPromiseClient.prototype.pullOnOff =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.OnOffApi/PullOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffApi_PullOnOff);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.OnOffInfoClient =
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
proto.smartcore.traits.OnOffInfoPromiseClient =
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
 *   !proto.smartcore.traits.DescribeOnOffRequest,
 *   !proto.smartcore.traits.OnOffSupport>}
 */
const methodDescriptor_OnOffInfo_DescribeOnOff = new grpc.web.MethodDescriptor(
  '/smartcore.traits.OnOffInfo/DescribeOnOff',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DescribeOnOffRequest,
  proto.smartcore.traits.OnOffSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeOnOffRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.OnOffSupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DescribeOnOffRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.traits.OnOffSupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.OnOffSupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.OnOffInfoClient.prototype.describeOnOff =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.OnOffInfo/DescribeOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffInfo_DescribeOnOff,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DescribeOnOffRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.OnOffSupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.OnOffInfoPromiseClient.prototype.describeOnOff =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.OnOffInfo/DescribeOnOff',
      request,
      metadata || {},
      methodDescriptor_OnOffInfo_DescribeOnOff);
};


module.exports = proto.smartcore.traits;

