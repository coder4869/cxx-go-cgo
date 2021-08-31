// int float double long

#ifndef ARRAY_TYPE_H
#define ARRAY_TYPE_H

#include <stdio.h>
#include "../macros.h"

#ifdef __cplusplus
extern "C" {
#endif

API_EXPORT(int) array_sum_int(int* ptr, int count);

API_EXPORT(float) array_sum_float(float* ptr, int count);

#ifdef __cplusplus
}
#endif

#endif  // ARRAY_TYPE_H