import * as jspb from "google-protobuf"

export class Http extends jspb.Message {
  getRulesList(): Array<HttpRule>;
  setRulesList(value: Array<HttpRule>): void;
  clearRulesList(): void;
  addRules(value?: HttpRule, index?: number): HttpRule;

  getFullyDecodeReservedExpansion(): boolean;
  setFullyDecodeReservedExpansion(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Http.AsObject;
  static toObject(includeInstance: boolean, msg: Http): Http.AsObject;
  static serializeBinaryToWriter(message: Http, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Http;
  static deserializeBinaryFromReader(message: Http, reader: jspb.BinaryReader): Http;
}

export namespace Http {
  export type AsObject = {
    rulesList: Array<HttpRule.AsObject>,
    fullyDecodeReservedExpansion: boolean,
  }
}

export class HttpRule extends jspb.Message {
  getSelector(): string;
  setSelector(value: string): void;

  getGet(): string;
  setGet(value: string): void;
  hasGet(): boolean;

  getPut(): string;
  setPut(value: string): void;
  hasPut(): boolean;

  getPost(): string;
  setPost(value: string): void;
  hasPost(): boolean;

  getDelete(): string;
  setDelete(value: string): void;
  hasDelete(): boolean;

  getPatch(): string;
  setPatch(value: string): void;
  hasPatch(): boolean;

  getCustom(): CustomHttpPattern | undefined;
  setCustom(value?: CustomHttpPattern): void;
  hasCustom(): boolean;
  clearCustom(): void;
  hasCustom(): boolean;

  getBody(): string;
  setBody(value: string): void;

  getResponseBody(): string;
  setResponseBody(value: string): void;

  getAdditionalBindingsList(): Array<HttpRule>;
  setAdditionalBindingsList(value: Array<HttpRule>): void;
  clearAdditionalBindingsList(): void;
  addAdditionalBindings(value?: HttpRule, index?: number): HttpRule;

  getPatternCase(): HttpRule.PatternCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HttpRule.AsObject;
  static toObject(includeInstance: boolean, msg: HttpRule): HttpRule.AsObject;
  static serializeBinaryToWriter(message: HttpRule, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HttpRule;
  static deserializeBinaryFromReader(message: HttpRule, reader: jspb.BinaryReader): HttpRule;
}

export namespace HttpRule {
  export type AsObject = {
    selector: string,
    get: string,
    put: string,
    post: string,
    pb_delete: string,
    patch: string,
    custom?: CustomHttpPattern.AsObject,
    body: string,
    responseBody: string,
    additionalBindingsList: Array<HttpRule.AsObject>,
  }

  export enum PatternCase { 
    PATTERN_NOT_SET = 0,
    GET = 2,
    PUT = 3,
    POST = 4,
    DELETE = 5,
    PATCH = 6,
    CUSTOM = 8,
  }
}

export class CustomHttpPattern extends jspb.Message {
  getKind(): string;
  setKind(value: string): void;

  getPath(): string;
  setPath(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomHttpPattern.AsObject;
  static toObject(includeInstance: boolean, msg: CustomHttpPattern): CustomHttpPattern.AsObject;
  static serializeBinaryToWriter(message: CustomHttpPattern, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomHttpPattern;
  static deserializeBinaryFromReader(message: CustomHttpPattern, reader: jspb.BinaryReader): CustomHttpPattern;
}

export namespace CustomHttpPattern {
  export type AsObject = {
    kind: string,
    path: string,
  }
}

