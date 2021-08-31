#ifndef USER_INFO_ENUM_HPP
#define USER_INFO_ENUM_HPP

enum PhoneState {
    PHONE_ADD_SUCCESS   = 0,
    PHONE_DEL_SUCCESS   = 1,
    PHONE_EXIST         = 2,
    PHONE_NOT_FOUND     = 3,
    PHONE_FORMAT_ERROR  = 4,
    PHONE_LENGTH_ERROR  = 5
};

enum class Gender {
    MALE    = 0,
    FEMALE  = 1
};

#endif // USER_INFO_ENUM_HPP