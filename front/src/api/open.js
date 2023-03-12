import request from "../util/request";

export function getInfo() {
  return request({
    url: `api/info`,
    method: "get",
  });
}

export function uploadPic(id, data) {
  return request({
    url: `api/upload/${id}`,
    method: "post",
    data,
  });
}
