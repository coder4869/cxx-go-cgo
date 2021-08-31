#ifndef USER_INFO_HPP
#define USER_INFO_HPP

#include <string>
#include <vector>

#include "user_info_enum.hpp"

class UserInfo {

public:
    UserInfo();
    ~UserInfo();
    
    void set_name(const std::string &name);
    std::string get_name() const;

    void set_idcard(const std::string &idcard);
    std::string get_idcard() const;

    void set_birthday(const std::string &birth);
    std::string get_birthday() const;

    void set_age(int age);
    int get_age() const;

    void set_gender(Gender g);
    Gender get_gender() const;

    void set_weight(float w);
    float get_weight() const;

    void set_height(float h);
    float get_height() const;

    PhoneState add_phone(const std::string &phone);
    PhoneState delete_phone(const std::string &phone);
    bool has_phone(const std::string &phone);

private:
    std::string _name;
    std::string _idcard;
    std::string _birthday;
    Gender      _gender;
    int         _age;
    float       _weight;
    float       _height;
    std::vector<std::string> _phones;
};

#endif // USER_INFO_HPP