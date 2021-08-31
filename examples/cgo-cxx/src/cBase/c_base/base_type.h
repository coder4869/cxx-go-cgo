// int float double long

#ifndef BASE_TYPE_H
#define BASE_TYPE_H

#include <stdio.h>
#include "../macros.h"

#ifdef __cplusplus
extern "C" {
#endif

API_EXPORT(int) base_add(int a, int b);

API_EXPORT(float) base_multi(float a, float b);

API_EXPORT(double) base_sub(double a, double b);

API_EXPORT(long) base_div(long a, long b);

API_EXPORT(int) base_sum(int* ptr, int count);

// 未导出的方法，在js中使用时，会报错 base_no_export is not a function
int base_no_export(int a, int b);

#ifdef __cplusplus
}
#endif

#endif  // BASE_TYPE_H