#ifndef RESPONSE_PB_H
#define RESPONSE_PB_H

#include <string>
#include <vector>

struct Vector2f {
  float x = 1;
  float y = 2;
}

struct Vector2d {
  double x = 1;
  double y = 2;
}

struct ResponseData {
    int         user_id;    // 用户编号
    std::string name;       // 用户名称
    std::string image;      // 用户头像base64
}

class ResponsePB {
public:
    ResponsePB();
    std::string get_pb();

private:
    int         err_no;             // 错误编号
    std::string err_msg;            // 错误描述
    std::vector<ResponseData> data; // 数据列表
};

#endif // RESPONSE_PB_H