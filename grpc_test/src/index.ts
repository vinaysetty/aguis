import {grpc} from "grpc-web-client";
import {Users, Void} from "./../_proto/ag_service_pb";
import {AutograderService} from "./../_proto/ag_service_pb_service";

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://localhost:9091" : "http://localhost:9090";

function getUsers() {
  const getUsersRequest = new Void();
  grpc.unary(AutograderService.GetUsers, {
    request: getUsersRequest,
    host: host,
    onEnd: res => {
      const { status, statusMessage, headers, message, trailers } = res;
      console.log("getBook.onEnd.status", status, statusMessage);
      console.log("getBook.onEnd.headers", headers);
      if (status === grpc.Code.OK && message) {
        console.log("getBook.onEnd.message", message.toObject());
      }
      console.log("getBook.onEnd.trailers", trailers);
    }
  });
}

getUsers();