import * as grpcWeb from 'grpc-web';

import * as traits_occupancy_sensor_pb from '../traits/occupancy_sensor_pb';


export class OccupancySensorApiClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getOccupancy(
    request: traits_occupancy_sensor_pb.GetOccupancyRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: traits_occupancy_sensor_pb.Occupancy) => void
  ): grpcWeb.ClientReadableStream<traits_occupancy_sensor_pb.Occupancy>;

  pullOccupancy(
    request: traits_occupancy_sensor_pb.PullOccupancyRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<traits_occupancy_sensor_pb.PullOccupancyResponse>;

}

export class OccupancySensorApiPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  getOccupancy(
    request: traits_occupancy_sensor_pb.GetOccupancyRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<traits_occupancy_sensor_pb.Occupancy>;

  pullOccupancy(
    request: traits_occupancy_sensor_pb.PullOccupancyRequest,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<traits_occupancy_sensor_pb.PullOccupancyResponse>;

}
