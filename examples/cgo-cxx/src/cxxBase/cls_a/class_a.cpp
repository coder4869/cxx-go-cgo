#include "class_a.h"

#include <string>
std::map<std::string, std::function<void(std::string, int)>> ClassA::_func_factories;

ClassA::ClassA(int x, std::string str)
: _x(x), _input_str(str)
{
    _class_name = "ClassA";
}

void ClassA::increase_x()
{
    ++_x;
}

int ClassA::get_x() const 
{
    return _x; 
}

void ClassA::set_x(int x) 
{ 
    _x = x; 
}

const char* ClassA::get_class_name()
{
    return (const char*)_class_name;
}

std::string ClassA::get_input_string(const ClassA& instance) 
{
    return instance._input_str;
}

std::string ClassA::foo () const {
    return "Signature: ()";
}

std::string ClassA::foo (int x) {
    return "Signature: (int x)";
}

std::string ClassA::foo (int x, int y) {
    return "Signature: (int x, int y)";
}

void ClassA::add_func(std::string func_name, int idx)
{
    printf("%s:func_name = %s, idx = %d \n", __FUNCTION__, func_name.c_str(), idx);
    _func_factories[func_name] = [func_name, idx](std::string name, int tag) {
        printf("%s: func_name = %s, idx = %d \n", __FUNCTION__, name.c_str(), tag);
    };
}

void ClassA::call_func(std::string func_name, int idx)
{
    auto iter = _func_factories.find(func_name);
    if (iter != _func_factories.end()) {
         iter->second(func_name, idx);
    } else {
        printf("%s:call back [%s] not found! idx = %d \n", __FUNCTION__, func_name.c_str(), idx);
    }
}

void ClassA::set_vector(const std::vector<float> &fvector)
{
    for(int idx=0; idx<fvector.size(); idx++) {
        printf("vector[%d] = %f\n", idx, fvector[idx]);
    }
}
