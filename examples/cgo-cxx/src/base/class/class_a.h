#ifndef CLASS_A_H
#define CLASS_A_H

#include <string>
#include <map>
#include <vector>
#include <functional>

class ClassA {
public:
    ClassA(int x, std::string str);

    void increase_x();

    int get_x() const;
    void set_x(int x);

    const char* get_class_name();

    static std::string get_input_string(const ClassA& instance);

    std::string foo () const;
    std::string foo (int x);
    std::string foo (int x, int y);

    void add_func(std::string func_name, int idx);
    void call_func(std::string func_name, int idx);

    void set_vector(const std::vector<float> &fvector);

private:
  static std::map<std::string, std::function<void(std::string, int)>> _func_factories;
  int _x;
  char * _class_name;
  std::string _input_str;
};

#endif // CLASS_A_H