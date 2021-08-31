#include "user_info.hpp"

UserInfo::UserInfo()
: _name(""), _idcard(""), _birthday("")
, _age(0), _gender(Gender::MALE), _weight(0.0), _height(0.0)
{ 

}

UserInfo::~UserInfo() 
{
    _phones.clear();
}

void UserInfo::set_name(const std::string &name)
{
    _name = name;
}
std::string UserInfo::get_name() const
{
    return _name;
}

void UserInfo::set_idcard(const std::string &idcard)
{
    _idcard = idcard;
}
std::string UserInfo::get_idcard() const
{
    return _idcard;
}

void UserInfo::set_birthday(const std::string &birth)
{
    _birthday = birth;
}
std::string UserInfo::get_birthday() const
{
    return _birthday;
}

void UserInfo::set_age(int age)
{
    _age = age;
}
int UserInfo::get_age() const
{
    return _age;
}

void UserInfo::set_gender(Gender g)
{
    _gender = g;
}
Gender UserInfo::get_gender() const
{
    return _gender;
}

void UserInfo::set_weight(float w)
{
    _weight = w;
}
float UserInfo::get_weight() const
{
    return _weight;
}

void UserInfo::set_height(float h)
{
    _height = h;
}
float UserInfo::get_height() const
{
    return _height;
}

PhoneState UserInfo::add_phone(const std::string &phone)
{
    for (int idx = 0; idx < _phones.size(); idx++) {
		if (_phones.at(idx) == phone) {
            return PHONE_EXIST;
        }
	}

    _phones.push_back(phone);
    return PHONE_ADD_SUCCESS;
}
PhoneState UserInfo::delete_phone(const std::string &phone)
{
    for (auto it = _phones.begin(); it != _phones.end(); ) {
        if (*it == phone) {
            it = _phones.erase(it);
            return PHONE_DEL_SUCCESS;
        } else {
            ++it;
        }
    }
    return PHONE_NOT_FOUND;
}

bool UserInfo::has_phone(const std::string &phone)
{
    for (int idx=0; idx<_phones.size(); idx++) {
        if (phone == _phones[idx]) {
            return true;
        }
    }
    return false;
}