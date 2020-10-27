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

var types_info_pb = require('../types/info_pb.js')

var types_time_period_pb = require('../types/time/period_pb.js')

var types_time_unit_pb = require('../types/time/unit_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.traits = require('./booking_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.BookingApiClient =
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
proto.smartcore.traits.BookingApiPromiseClient =
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
 *   !proto.smartcore.traits.ListBookingsRequest,
 *   !proto.smartcore.traits.ListBookingsResponse>}
 */
const methodDescriptor_BookingApi_ListBookings = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingApi/ListBookings',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.ListBookingsRequest,
  proto.smartcore.traits.ListBookingsResponse,
  /**
   * @param {!proto.smartcore.traits.ListBookingsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ListBookingsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.ListBookingsRequest,
 *   !proto.smartcore.traits.ListBookingsResponse>}
 */
const methodInfo_BookingApi_ListBookings = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.ListBookingsResponse,
  /**
   * @param {!proto.smartcore.traits.ListBookingsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.ListBookingsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ListBookingsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.ListBookingsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.ListBookingsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiClient.prototype.listBookings =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.BookingApi/ListBookings',
      request,
      metadata || {},
      methodDescriptor_BookingApi_ListBookings,
      callback);
};


/**
 * @param {!proto.smartcore.traits.ListBookingsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.ListBookingsResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.BookingApiPromiseClient.prototype.listBookings =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.BookingApi/ListBookings',
      request,
      metadata || {},
      methodDescriptor_BookingApi_ListBookings);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.CheckInBookingRequest,
 *   !proto.smartcore.traits.CheckInBookingResponse>}
 */
const methodDescriptor_BookingApi_CheckInBooking = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingApi/CheckInBooking',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.CheckInBookingRequest,
  proto.smartcore.traits.CheckInBookingResponse,
  /**
   * @param {!proto.smartcore.traits.CheckInBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.CheckInBookingResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.CheckInBookingRequest,
 *   !proto.smartcore.traits.CheckInBookingResponse>}
 */
const methodInfo_BookingApi_CheckInBooking = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.CheckInBookingResponse,
  /**
   * @param {!proto.smartcore.traits.CheckInBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.CheckInBookingResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.CheckInBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.CheckInBookingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.CheckInBookingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiClient.prototype.checkInBooking =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.BookingApi/CheckInBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_CheckInBooking,
      callback);
};


/**
 * @param {!proto.smartcore.traits.CheckInBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.CheckInBookingResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.BookingApiPromiseClient.prototype.checkInBooking =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.BookingApi/CheckInBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_CheckInBooking);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.CheckOutBookingRequest,
 *   !proto.smartcore.traits.CheckOutBookingResponse>}
 */
const methodDescriptor_BookingApi_CheckOutBooking = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingApi/CheckOutBooking',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.CheckOutBookingRequest,
  proto.smartcore.traits.CheckOutBookingResponse,
  /**
   * @param {!proto.smartcore.traits.CheckOutBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.CheckOutBookingResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.CheckOutBookingRequest,
 *   !proto.smartcore.traits.CheckOutBookingResponse>}
 */
const methodInfo_BookingApi_CheckOutBooking = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.CheckOutBookingResponse,
  /**
   * @param {!proto.smartcore.traits.CheckOutBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.CheckOutBookingResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.CheckOutBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.CheckOutBookingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.CheckOutBookingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiClient.prototype.checkOutBooking =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.BookingApi/CheckOutBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_CheckOutBooking,
      callback);
};


/**
 * @param {!proto.smartcore.traits.CheckOutBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.CheckOutBookingResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.BookingApiPromiseClient.prototype.checkOutBooking =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.BookingApi/CheckOutBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_CheckOutBooking);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.CreateBookingRequest,
 *   !proto.smartcore.traits.CreateBookingResponse>}
 */
const methodDescriptor_BookingApi_CreateBooking = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingApi/CreateBooking',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.CreateBookingRequest,
  proto.smartcore.traits.CreateBookingResponse,
  /**
   * @param {!proto.smartcore.traits.CreateBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.CreateBookingResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.CreateBookingRequest,
 *   !proto.smartcore.traits.CreateBookingResponse>}
 */
const methodInfo_BookingApi_CreateBooking = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.CreateBookingResponse,
  /**
   * @param {!proto.smartcore.traits.CreateBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.CreateBookingResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.CreateBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.CreateBookingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.CreateBookingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiClient.prototype.createBooking =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.BookingApi/CreateBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_CreateBooking,
      callback);
};


/**
 * @param {!proto.smartcore.traits.CreateBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.CreateBookingResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.BookingApiPromiseClient.prototype.createBooking =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.BookingApi/CreateBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_CreateBooking);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.UpdateBookingRequest,
 *   !proto.smartcore.traits.UpdateBookingResponse>}
 */
const methodDescriptor_BookingApi_UpdateBooking = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingApi/UpdateBooking',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.UpdateBookingRequest,
  proto.smartcore.traits.UpdateBookingResponse,
  /**
   * @param {!proto.smartcore.traits.UpdateBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.UpdateBookingResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.UpdateBookingRequest,
 *   !proto.smartcore.traits.UpdateBookingResponse>}
 */
const methodInfo_BookingApi_UpdateBooking = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.UpdateBookingResponse,
  /**
   * @param {!proto.smartcore.traits.UpdateBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.UpdateBookingResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.UpdateBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.UpdateBookingResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.UpdateBookingResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiClient.prototype.updateBooking =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.BookingApi/UpdateBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_UpdateBooking,
      callback);
};


/**
 * @param {!proto.smartcore.traits.UpdateBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.UpdateBookingResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.BookingApiPromiseClient.prototype.updateBooking =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.BookingApi/UpdateBooking',
      request,
      metadata || {},
      methodDescriptor_BookingApi_UpdateBooking);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.traits.ListBookingsRequest,
 *   !proto.smartcore.traits.PullBookingsResponse>}
 */
const methodDescriptor_BookingApi_PullBookings = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingApi/PullBookings',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.traits.ListBookingsRequest,
  proto.smartcore.traits.PullBookingsResponse,
  /**
   * @param {!proto.smartcore.traits.ListBookingsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullBookingsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.ListBookingsRequest,
 *   !proto.smartcore.traits.PullBookingsResponse>}
 */
const methodInfo_BookingApi_PullBookings = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.PullBookingsResponse,
  /**
   * @param {!proto.smartcore.traits.ListBookingsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.PullBookingsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.ListBookingsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullBookingsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiClient.prototype.pullBookings =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.BookingApi/PullBookings',
      request,
      metadata || {},
      methodDescriptor_BookingApi_PullBookings);
};


/**
 * @param {!proto.smartcore.traits.ListBookingsRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.PullBookingsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingApiPromiseClient.prototype.pullBookings =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.traits.BookingApi/PullBookings',
      request,
      metadata || {},
      methodDescriptor_BookingApi_PullBookings);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.traits.BookingInfoClient =
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
proto.smartcore.traits.BookingInfoPromiseClient =
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
 *   !proto.smartcore.traits.DescribeBookingRequest,
 *   !proto.smartcore.traits.BookingSupport>}
 */
const methodDescriptor_BookingInfo_DescribeBooking = new grpc.web.MethodDescriptor(
  '/smartcore.traits.BookingInfo/DescribeBooking',
  grpc.web.MethodType.UNARY,
  proto.smartcore.traits.DescribeBookingRequest,
  proto.smartcore.traits.BookingSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.BookingSupport.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.traits.DescribeBookingRequest,
 *   !proto.smartcore.traits.BookingSupport>}
 */
const methodInfo_BookingInfo_DescribeBooking = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.traits.BookingSupport,
  /**
   * @param {!proto.smartcore.traits.DescribeBookingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.traits.BookingSupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.traits.DescribeBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.traits.BookingSupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.traits.BookingSupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.traits.BookingInfoClient.prototype.describeBooking =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.traits.BookingInfo/DescribeBooking',
      request,
      metadata || {},
      methodDescriptor_BookingInfo_DescribeBooking,
      callback);
};


/**
 * @param {!proto.smartcore.traits.DescribeBookingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.traits.BookingSupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.traits.BookingInfoPromiseClient.prototype.describeBooking =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.traits.BookingInfo/DescribeBooking',
      request,
      metadata || {},
      methodDescriptor_BookingInfo_DescribeBooking);
};


module.exports = proto.smartcore.traits;

