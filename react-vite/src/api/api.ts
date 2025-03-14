// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.1
//   protoc               v5.29.3
// source: api.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "buf";

/** Socket */
export enum SocketMessageType {
  SET_NAME = 0,
  CREATE_ROOM = 1,
  JOIN_ROOM = 2,
  LEAVE_ROOM = 3,
  ERROR = 4,
  UNRECOGNIZED = -1,
}

export function socketMessageTypeFromJSON(object: any): SocketMessageType {
  switch (object) {
    case 0:
    case "SET_NAME":
      return SocketMessageType.SET_NAME;
    case 1:
    case "CREATE_ROOM":
      return SocketMessageType.CREATE_ROOM;
    case 2:
    case "JOIN_ROOM":
      return SocketMessageType.JOIN_ROOM;
    case 3:
    case "LEAVE_ROOM":
      return SocketMessageType.LEAVE_ROOM;
    case 4:
    case "ERROR":
      return SocketMessageType.ERROR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SocketMessageType.UNRECOGNIZED;
  }
}

export function socketMessageTypeToJSON(object: SocketMessageType): string {
  switch (object) {
    case SocketMessageType.SET_NAME:
      return "SET_NAME";
    case SocketMessageType.CREATE_ROOM:
      return "CREATE_ROOM";
    case SocketMessageType.JOIN_ROOM:
      return "JOIN_ROOM";
    case SocketMessageType.LEAVE_ROOM:
      return "LEAVE_ROOM";
    case SocketMessageType.ERROR:
      return "ERROR";
    case SocketMessageType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Error {
  message: string;
}

export interface User {
  id: number;
  name: string;
}

export interface Users {
  users: User[];
}

export interface WSMessage {
  type: SocketMessageType;
  setName?: SetNameRequest | undefined;
  createRoom?: CreateRoomRequest | undefined;
  joinRoom?: JoinRoomRequest | undefined;
  error?: Error | undefined;
}

export interface SetNameRequest {
  name: string;
}

export interface CreateRoomRequest {
  name: string;
}

export interface JoinRoomRequest {
  id: string;
}

export interface Room {
  id: string;
  name: string;
}

export interface CreateRoomResponse {
  room: Room | undefined;
}

/** Also API */
export interface GetRooms {
  rooms: Room[];
}

function createBaseError(): Error {
  return { message: "" };
}

export const Error: MessageFns<Error> = {
  encode(message: Error, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Error {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseError();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.message = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Error {
    return { message: isSet(object.message) ? globalThis.String(object.message) : "" };
  },

  toJSON(message: Error): unknown {
    const obj: any = {};
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Error>, I>>(base?: I): Error {
    return Error.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Error>, I>>(object: I): Error {
    const message = createBaseError();
    message.message = object.message ?? "";
    return message;
  },
};

function createBaseUser(): User {
  return { id: 0, name: "" };
}

export const User: MessageFns<User> = {
  encode(message: User, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): User {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.id = reader.int32();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      id: isSet(object.id) ? globalThis.Number(object.id) : 0,
      name: isSet(object.name) ? globalThis.String(object.name) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    if (message.id !== 0) {
      obj.id = Math.round(message.id);
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<User>, I>>(base?: I): User {
    return User.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.id = object.id ?? 0;
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseUsers(): Users {
  return { users: [] };
}

export const Users: MessageFns<Users> = {
  encode(message: Users, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.users) {
      User.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Users {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUsers();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.users.push(User.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Users {
    return { users: globalThis.Array.isArray(object?.users) ? object.users.map((e: any) => User.fromJSON(e)) : [] };
  },

  toJSON(message: Users): unknown {
    const obj: any = {};
    if (message.users?.length) {
      obj.users = message.users.map((e) => User.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Users>, I>>(base?: I): Users {
    return Users.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Users>, I>>(object: I): Users {
    const message = createBaseUsers();
    message.users = object.users?.map((e) => User.fromPartial(e)) || [];
    return message;
  },
};

function createBaseWSMessage(): WSMessage {
  return { type: 0, setName: undefined, createRoom: undefined, joinRoom: undefined, error: undefined };
}

export const WSMessage: MessageFns<WSMessage> = {
  encode(message: WSMessage, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.setName !== undefined) {
      SetNameRequest.encode(message.setName, writer.uint32(18).fork()).join();
    }
    if (message.createRoom !== undefined) {
      CreateRoomRequest.encode(message.createRoom, writer.uint32(26).fork()).join();
    }
    if (message.joinRoom !== undefined) {
      JoinRoomRequest.encode(message.joinRoom, writer.uint32(34).fork()).join();
    }
    if (message.error !== undefined) {
      Error.encode(message.error, writer.uint32(42).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): WSMessage {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWSMessage();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 8) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.setName = SetNameRequest.decode(reader, reader.uint32());
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.createRoom = CreateRoomRequest.decode(reader, reader.uint32());
          continue;
        }
        case 4: {
          if (tag !== 34) {
            break;
          }

          message.joinRoom = JoinRoomRequest.decode(reader, reader.uint32());
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.error = Error.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): WSMessage {
    return {
      type: isSet(object.type) ? socketMessageTypeFromJSON(object.type) : 0,
      setName: isSet(object.setName) ? SetNameRequest.fromJSON(object.setName) : undefined,
      createRoom: isSet(object.createRoom) ? CreateRoomRequest.fromJSON(object.createRoom) : undefined,
      joinRoom: isSet(object.joinRoom) ? JoinRoomRequest.fromJSON(object.joinRoom) : undefined,
      error: isSet(object.error) ? Error.fromJSON(object.error) : undefined,
    };
  },

  toJSON(message: WSMessage): unknown {
    const obj: any = {};
    if (message.type !== 0) {
      obj.type = socketMessageTypeToJSON(message.type);
    }
    if (message.setName !== undefined) {
      obj.setName = SetNameRequest.toJSON(message.setName);
    }
    if (message.createRoom !== undefined) {
      obj.createRoom = CreateRoomRequest.toJSON(message.createRoom);
    }
    if (message.joinRoom !== undefined) {
      obj.joinRoom = JoinRoomRequest.toJSON(message.joinRoom);
    }
    if (message.error !== undefined) {
      obj.error = Error.toJSON(message.error);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<WSMessage>, I>>(base?: I): WSMessage {
    return WSMessage.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<WSMessage>, I>>(object: I): WSMessage {
    const message = createBaseWSMessage();
    message.type = object.type ?? 0;
    message.setName = (object.setName !== undefined && object.setName !== null)
      ? SetNameRequest.fromPartial(object.setName)
      : undefined;
    message.createRoom = (object.createRoom !== undefined && object.createRoom !== null)
      ? CreateRoomRequest.fromPartial(object.createRoom)
      : undefined;
    message.joinRoom = (object.joinRoom !== undefined && object.joinRoom !== null)
      ? JoinRoomRequest.fromPartial(object.joinRoom)
      : undefined;
    message.error = (object.error !== undefined && object.error !== null) ? Error.fromPartial(object.error) : undefined;
    return message;
  },
};

function createBaseSetNameRequest(): SetNameRequest {
  return { name: "" };
}

export const SetNameRequest: MessageFns<SetNameRequest> = {
  encode(message: SetNameRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): SetNameRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSetNameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SetNameRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: SetNameRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SetNameRequest>, I>>(base?: I): SetNameRequest {
    return SetNameRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SetNameRequest>, I>>(object: I): SetNameRequest {
    const message = createBaseSetNameRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseCreateRoomRequest(): CreateRoomRequest {
  return { name: "" };
}

export const CreateRoomRequest: MessageFns<CreateRoomRequest> = {
  encode(message: CreateRoomRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateRoomRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateRoomRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateRoomRequest {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: CreateRoomRequest): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateRoomRequest>, I>>(base?: I): CreateRoomRequest {
    return CreateRoomRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateRoomRequest>, I>>(object: I): CreateRoomRequest {
    const message = createBaseCreateRoomRequest();
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseJoinRoomRequest(): JoinRoomRequest {
  return { id: "" };
}

export const JoinRoomRequest: MessageFns<JoinRoomRequest> = {
  encode(message: JoinRoomRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): JoinRoomRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseJoinRoomRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): JoinRoomRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: JoinRoomRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<JoinRoomRequest>, I>>(base?: I): JoinRoomRequest {
    return JoinRoomRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<JoinRoomRequest>, I>>(object: I): JoinRoomRequest {
    const message = createBaseJoinRoomRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseRoom(): Room {
  return { id: "", name: "" };
}

export const Room: MessageFns<Room> = {
  encode(message: Room, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Room {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoom();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Room {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      name: isSet(object.name) ? globalThis.String(object.name) : "",
    };
  },

  toJSON(message: Room): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Room>, I>>(base?: I): Room {
    return Room.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Room>, I>>(object: I): Room {
    const message = createBaseRoom();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    return message;
  },
};

function createBaseCreateRoomResponse(): CreateRoomResponse {
  return { room: undefined };
}

export const CreateRoomResponse: MessageFns<CreateRoomResponse> = {
  encode(message: CreateRoomResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.room !== undefined) {
      Room.encode(message.room, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateRoomResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateRoomResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.room = Room.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateRoomResponse {
    return { room: isSet(object.room) ? Room.fromJSON(object.room) : undefined };
  },

  toJSON(message: CreateRoomResponse): unknown {
    const obj: any = {};
    if (message.room !== undefined) {
      obj.room = Room.toJSON(message.room);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateRoomResponse>, I>>(base?: I): CreateRoomResponse {
    return CreateRoomResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateRoomResponse>, I>>(object: I): CreateRoomResponse {
    const message = createBaseCreateRoomResponse();
    message.room = (object.room !== undefined && object.room !== null) ? Room.fromPartial(object.room) : undefined;
    return message;
  },
};

function createBaseGetRooms(): GetRooms {
  return { rooms: [] };
}

export const GetRooms: MessageFns<GetRooms> = {
  encode(message: GetRooms, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.rooms) {
      Room.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetRooms {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRooms();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.rooms.push(Room.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRooms {
    return { rooms: globalThis.Array.isArray(object?.rooms) ? object.rooms.map((e: any) => Room.fromJSON(e)) : [] };
  },

  toJSON(message: GetRooms): unknown {
    const obj: any = {};
    if (message.rooms?.length) {
      obj.rooms = message.rooms.map((e) => Room.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRooms>, I>>(base?: I): GetRooms {
    return GetRooms.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetRooms>, I>>(object: I): GetRooms {
    const message = createBaseGetRooms();
    message.rooms = object.rooms?.map((e) => Room.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
