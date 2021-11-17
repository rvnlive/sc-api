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

var types_info_pb = require('../types/info_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./energy_storage_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.EnergyStorageApiClient =
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
proto.smartcore.traits.EnergyStorageApiPromiseClient =
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
 *   !proto.smartcore.traits.GetEnergyLevelRequest,
 *   !proto.smartcore.traits.EnergyLevel>}
 */
const methodDescriptor_EnergyStorageApi_GetEnergyLevel = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EnergyStorageApi/GetEnergyLevel',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.GetEnergyLevelRequest,
  proto.smartcore.traits.EnergyLevel,
  /**
   * @param {!proto.smartcore.traits.GetEnergyLevelRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.EnergyLevel.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.GetEnergyLevelRequest,
 *   !proto.smartcore.traits.EnergyLevel>}
 */
const methodInfo_EnergyStorageApi_GetEnergyLevel = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.EnergyLevel,
  /**
   * @param {!proto.smartcore.traits.GetEnergyLevelRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.EnergyLevel.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.GetEnergyLevelRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.EnergyLevel)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.EnergyLevel>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EnergyStorageApiClient.prototype.getEnergyLevel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.EnergyStorageApi/GetEnergyLevel',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageApi_GetEnergyLevel,
      callback);
};


/**
 * @param {!proto.smartcore.traits.GetEnergyLevelRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.EnergyLevel>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.EnergyStorageApiPromiseClient.prototype.getEnergyLevel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.EnergyStorageApi/GetEnergyLevel',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageApi_GetEnergyLevel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.PullEnergyLevelRequest,
 *   !proto.smartcore.traits.PullEnergyLevelResponse>}
 */
const methodDescriptor_EnergyStorageApi_PullEnergyLevel = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EnergyStorageApi/PullEnergyLevel',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.PullEnergyLevelRequest,
  proto.smartcore.traits.PullEnergyLevelResponse,
  /**
   * @param {!proto.smartcore.traits.PullEnergyLevelRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullEnergyLevelResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.PullEnergyLevelRequest,
 *   !proto.smartcore.traits.PullEnergyLevelResponse>}
 */
const methodInfo_EnergyStorageApi_PullEnergyLevel = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.PullEnergyLevelResponse,
  /**
   * @param {!proto.smartcore.traits.PullEnergyLevelRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullEnergyLevelResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.PullEnergyLevelRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullEnergyLevelResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EnergyStorageApiClient.prototype.pullEnergyLevel =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.EnergyStorageApi/PullEnergyLevel',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageApi_PullEnergyLevel);
};


/**
 * @param {!proto.smartcore.traits.PullEnergyLevelRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullEnergyLevelResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EnergyStorageApiPromiseClient.prototype.pullEnergyLevel =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.EnergyStorageApi/PullEnergyLevel',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageApi_PullEnergyLevel);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ChargeRequest,
 *   !proto.smartcore.traits.ChargeResponse>}
 */
const methodDescriptor_EnergyStorageApi_Charge = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EnergyStorageApi/Charge',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ChargeRequest,
  proto.smartcore.traits.ChargeResponse,
  /**
   * @param {!proto.smartcore.traits.ChargeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ChargeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.ChargeRequest,
 *   !proto.smartcore.traits.ChargeResponse>}
 */
const methodInfo_EnergyStorageApi_Charge = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.ChargeResponse,
  /**
   * @param {!proto.smartcore.traits.ChargeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ChargeResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ChargeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.ChargeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ChargeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EnergyStorageApiClient.prototype.charge =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.EnergyStorageApi/Charge',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageApi_Charge,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ChargeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ChargeResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.EnergyStorageApiPromiseClient.prototype.charge =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.EnergyStorageApi/Charge',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageApi_Charge);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.EnergyStorageInfoClient =
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
proto.smartcore.traits.EnergyStorageInfoPromiseClient =
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
 *   !proto.smartcore.traits.DescribeEnergyLevelRequest,
 *   !proto.smartcore.traits.EnergyLevelSupport>}
 */
const methodDescriptor_EnergyStorageInfo_DescribeEnergyLevel = new grpc.web.MethodDescriptor(
  '/smartcore.traits.EnergyStorageInfo/DescribeEnergyLevel',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DescribeEnergyLevelRequest,
  proto.smartcore.traits.EnergyLevelSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeEnergyLevelRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.EnergyLevelSupport.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.DescribeEnergyLevelRequest,
 *   !proto.smartcore.traits.EnergyLevelSupport>}
 */
const methodInfo_EnergyStorageInfo_DescribeEnergyLevel = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.EnergyLevelSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeEnergyLevelRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.EnergyLevelSupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DescribeEnergyLevelRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.EnergyLevelSupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.EnergyLevelSupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.EnergyStorageInfoClient.prototype.describeEnergyLevel =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.EnergyStorageInfo/DescribeEnergyLevel',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageInfo_DescribeEnergyLevel,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DescribeEnergyLevelRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.EnergyLevelSupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.EnergyStorageInfoPromiseClient.prototype.describeEnergyLevel =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.EnergyStorageInfo/DescribeEnergyLevel',
      request,
      metadata || {},
      methodDescriptor_EnergyStorageInfo_DescribeEnergyLevel);
};


module.exports = proto.smartcore.traits;
