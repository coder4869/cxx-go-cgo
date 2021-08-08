#ifndef REQUEST_PB_H
#define REQUEST_PB_H

#include <string>

class RequestPB {
public:
    RequestPB();
    std::string get_pb();
  
private:
    int user_id;    // 用户编号
};

#endif // REQUEST_PB_H