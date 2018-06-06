import {Message} from "google-protobuf";

import {grpc} from "grpc-web-client";
import {Users, Void} from "./../_proto/ag_service_pb";
import {AutograderService} from "./../_proto/ag_service_pb_service";

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://ec2-18-130-15-23.eu-west-2.compute.amazonaws.com:8091" : "http://ec2-18-130-15-23.eu-west-2.compute.amazonaws.com:8090";

export interface IGrpcResult<T> {
  statusCode: number;
  data?: T;
}

function getUsers(): Promise<IGrpcResult<Users>> {
  const usersRequest = new Void();
  return grpcUnary<Users>(AutograderService.GetUsers, usersRequest);
}

function grpcUnary<TReceive extends Message>(method: any, request: any): Promise<IGrpcResult<TReceive>> {
  const requestPromise = new Promise<IGrpcResult<TReceive>>((resolve, reject) => {
      grpc.unary(method, {
          request,
          host,
          onEnd: (res) => {
              const {status, statusMessage, headers, message, trailers} = res;
              const temp: IGrpcResult<TReceive> = {
                  data: message as TReceive,
                  statusCode: status,
              };
              console.log("Result: ", temp.data)
              resolve(temp);

          },
      });
  });
  return requestPromise;
}

var i;
for (i = 0; i < 100; i++){
  setTimeout(function(){
    getUsers();
  }, i * 1000)
}