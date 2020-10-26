import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as types_number_pb from '../types/number_pb';
import * as types_tween_pb from '../types/tween_pb';


export class Temperature extends jspb.Message {
  getValueCelsius(): number;
  setValueCelsius(value: number): Temperature;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Temperature.AsObject;
  static toObject(includeInstance: boolean, msg: Temperature): Temperature.AsObject;
  static serializeBinaryToWriter(message: Temperature, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Temperature;
  static deserializeBinaryFromReader(message: Temperature, reader: jspb.BinaryReader): Temperature;
}

export namespace Temperature {
  export type AsObject = {
    valueCelsius: number,
  }
}

export class Volume extends jspb.Message {
  getGain(): types_number_pb.FloatVar | undefined;
  setGain(value?: types_number_pb.FloatVar): Volume;
  hasGain(): boolean;
  clearGain(): Volume;

  getMute(): boolean;
  setMute(value: boolean): Volume;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Volume.AsObject;
  static toObject(includeInstance: boolean, msg: Volume): Volume.AsObject;
  static serializeBinaryToWriter(message: Volume, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Volume;
  static deserializeBinaryFromReader(message: Volume, reader: jspb.BinaryReader): Volume;
}

export namespace Volume {
  export type AsObject = {
    gain?: types_number_pb.FloatVar.AsObject,
    mute: boolean,
  }
}

export class AudioLevel extends jspb.Message {
  getGain(): number;
  setGain(value: number): AudioLevel;

  getGainTween(): types_tween_pb.Tween | undefined;
  setGainTween(value?: types_tween_pb.Tween): AudioLevel;
  hasGainTween(): boolean;
  clearGainTween(): AudioLevel;

  getTargetGain(): number;
  setTargetGain(value: number): AudioLevel;

  getMuted(): boolean;
  setMuted(value: boolean): AudioLevel;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AudioLevel.AsObject;
  static toObject(includeInstance: boolean, msg: AudioLevel): AudioLevel.AsObject;
  static serializeBinaryToWriter(message: AudioLevel, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AudioLevel;
  static deserializeBinaryFromReader(message: AudioLevel, reader: jspb.BinaryReader): AudioLevel;
}

export namespace AudioLevel {
  export type AsObject = {
    gain: number,
    gainTween?: types_tween_pb.Tween.AsObject,
    targetGain: number,
    muted: boolean,
  }
}

export class AudioLevelChange extends jspb.Message {
  getName(): string;
  setName(value: string): AudioLevelChange;

  getChangeTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setChangeTime(value?: google_protobuf_timestamp_pb.Timestamp): AudioLevelChange;
  hasChangeTime(): boolean;
  clearChangeTime(): AudioLevelChange;

  getLevel(): AudioLevel | undefined;
  setLevel(value?: AudioLevel): AudioLevelChange;
  hasLevel(): boolean;
  clearLevel(): AudioLevelChange;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AudioLevelChange.AsObject;
  static toObject(includeInstance: boolean, msg: AudioLevelChange): AudioLevelChange.AsObject;
  static serializeBinaryToWriter(message: AudioLevelChange, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AudioLevelChange;
  static deserializeBinaryFromReader(message: AudioLevelChange, reader: jspb.BinaryReader): AudioLevelChange;
}

export namespace AudioLevelChange {
  export type AsObject = {
    name: string,
    changeTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    level?: AudioLevel.AsObject,
  }
}

export enum TemperatureUnit { 
  CELSIUS = 0,
  FAHRENHEIT = 1,
  KELVIN = 2,
}
